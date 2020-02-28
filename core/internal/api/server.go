// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"math/big"
	"net"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"git.monogon.dev/source/nexantic.git/core/generated/api"
	schema "git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"git.monogon.dev/source/nexantic.git/core/internal/common/service"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus"
)

type (
	Server struct {
		*service.BaseService

		grpcServer         *grpc.Server
		externalGrpcServer *grpc.Server

		consensusService *consensus.Service

		config *Config
	}

	Config struct {
	}
)

var (
	// From RFC 5280 Section 4.1.2.5
	unknownNotAfter = time.Unix(253402300799, 0)
)

func NewApiServer(config *Config, logger *zap.Logger, consensusService *consensus.Service) (*Server, error) {
	s := &Server{
		config:           config,
		consensusService: consensusService,
	}

	s.BaseService = service.NewBaseService("api", logger, s)

	return s, nil
}

func (s *Server) getStore() clientv3.KV {
	// Cannot be moved to initialization because an internal reference will be nil
	return s.consensusService.GetStore("api", "")
}

// BootstrapNewClusterHook creates the necessary key material for the API Servers and stores it in
// the consensus service. It also creates a node entry for the initial node.
func (s *Server) BootstrapNewClusterHook(initNodeReq *api.NewNodeInfo) error {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("Failed to generate serial number: %w", err)
	}

	pubKey, privKeyRaw, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	privkey, err := x509.MarshalPKCS8PrivateKey(privKeyRaw)
	if err != nil {
		return err
	}

	// This has no SANs because it authenticates by public key, not by name
	masterCert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: "Smalltown Master",
		},
		IsCA:                  false,
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              unknownNotAfter,
		// Certificate is used both as server & client
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}
	cert, err := x509.CreateCertificate(rand.Reader, masterCert, masterCert, pubKey, privKeyRaw)
	if err != nil {
		return err
	}
	store := s.getStore()
	if _, err := store.Put(context.Background(), "master.der", string(cert)); err != nil {
		return err
	}
	if _, err := store.Put(context.Background(), "master-key.der", string(privkey)); err != nil {
		return err
	}

	// TODO: Further integrity providers need to be plumbed in here
	node, err := s.TPM2BootstrapNode(initNodeReq)
	if err != nil {
		return err
	}

	if err := s.registerNewNode(node); err != nil {
		return err
	}
	return nil
}

// GetMasterCert gets the master certificate in X.509 DER form
// This is mainly used to issue enrolment configs
func (s *Server) GetMasterCert() ([]byte, error) {
	store := s.getStore()
	res, err := store.Get(context.Background(), "master.der")
	if err != nil {
		return []byte{}, err
	}
	if len(res.Kvs) != 1 {
		return []byte{}, errors.New("master certificate not found")
	}
	certRaw := res.Kvs[0].Value
	return certRaw, nil
}

// TODO(lorenz): Move consensus/certificate interaction into a utility, is now duplicated too often
func (s *Server) loadMasterCert() (*tls.Certificate, error) {

	store := s.getStore()
	var tlsCert tls.Certificate
	res, err := store.Get(context.Background(), "master.der")
	if err != nil {
		return nil, err
	}
	if len(res.Kvs) != 1 {
		return nil, errors.New("master certificate not found")
	}
	certRaw := res.Kvs[0].Value

	tlsCert.Certificate = append(tlsCert.Certificate, certRaw)
	tlsCert.Leaf, err = x509.ParseCertificate(certRaw)

	res, err = store.Get(context.Background(), "master-key.der")
	if err != nil {
		return nil, err
	}
	if len(res.Kvs) != 1 {
		return nil, errors.New("master certificate not found")
	}
	keyRaw := res.Kvs[0].Value
	key, err := x509.ParsePKCS8PrivateKey(keyRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to load master private key: %w", err)
	}
	edKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("invalid private key")
	}
	tlsCert.PrivateKey = edKey
	return &tlsCert, nil
}

func (s *Server) OnStart() error {
	masterListenHost := fmt.Sprintf(":%d", common.MasterServicePort)
	lis, err := net.Listen("tcp", masterListenHost)
	if err != nil {
		s.Logger.Fatal("failed to listen", zap.Error(err))
	}

	externalListeneHost := fmt.Sprintf(":%d", common.ExternalServicePort)
	externalListener, err := net.Listen("tcp", externalListeneHost)
	if err != nil {
		s.Logger.Fatal("failed to listen", zap.Error(err))
	}

	masterCert, err := s.loadMasterCert()
	if err != nil {
		s.Logger.Error("Failed to load Master Service Key Material: %w", zap.Error(err))
		return err
	}

	masterTransportCredentials := credentials.NewServerTLSFromCert(masterCert)

	masterGrpcServer := grpc.NewServer(grpc.Creds(masterTransportCredentials))
	clusterManagementGrpcServer := grpc.NewServer()
	schema.RegisterClusterManagementServer(clusterManagementGrpcServer, s)
	schema.RegisterNodeManagementServiceServer(masterGrpcServer, s)

	reflection.Register(masterGrpcServer)

	s.grpcServer = masterGrpcServer
	s.externalGrpcServer = clusterManagementGrpcServer

	go func() {
		err = s.grpcServer.Serve(lis)
		s.Logger.Error("API server failed", zap.Error(err))
	}()

	go func() {
		err = s.externalGrpcServer.Serve(externalListener)
		s.Logger.Error("API server failed", zap.Error(err))
	}()

	s.Logger.Info("gRPC listening", zap.String("host", masterListenHost))

	return nil
}

func (s *Server) OnStop() error {
	s.grpcServer.Stop()
	s.externalGrpcServer.Stop()

	return nil
}
