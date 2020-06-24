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

package node

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	apipb "git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus"
	"git.monogon.dev/source/nexantic.git/core/internal/containerd"
	"git.monogon.dev/source/nexantic.git/core/internal/integrity/tpm2"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes"
	"git.monogon.dev/source/nexantic.git/core/internal/network"
	"git.monogon.dev/source/nexantic.git/core/internal/storage"
)

var (
	// From RFC 5280 Section 4.1.2.5
	unknownNotAfter = time.Unix(253402300799, 0)
)

type (
	SmalltownNode struct {
		Api        *api.Server
		Consensus  *consensus.Service
		Storage    *storage.Manager
		Kubernetes *kubernetes.Service
		Containerd *containerd.Service
		Network    *network.Service

		logger          *zap.Logger
		state           common.SmalltownState
		hostname        string
		enrolmentConfig *apipb.EnrolmentConfig

		debugServer *grpc.Server
	}
)

func NewSmalltownNode(logger *zap.Logger, ntwk *network.Service, strg *storage.Manager) (*SmalltownNode, error) {
	flag.Parse()
	logger.Info("Creating Smalltown node")
	ctx := context.Background()

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// Wait for IP adddress...
	ctxT, ctxTC := context.WithTimeout(ctx, time.Second*10)
	defer ctxTC()
	externalIP, err := ntwk.GetIP(ctxT, true)
	if err != nil {
		logger.Panic("Could not get IP address", zap.Error(err))
	}

	// Important to know if the GetIP above hangs
	logger.Info("Node has IP", zap.String("ip", externalIP.String()))

	consensusService, err := consensus.NewConsensusService(consensus.Config{
		Name:         hostname,
		ListenHost:   "0.0.0.0",
		ExternalHost: externalIP.String(),
	}, logger.With(zap.String("module", "consensus")))
	if err != nil {
		return nil, err
	}

	containerdService, err := containerd.New()
	if err != nil {
		return nil, err
	}

	s := &SmalltownNode{
		Consensus:  consensusService,
		Containerd: containerdService,
		Storage:    strg,
		Network:    ntwk,
		logger:     logger,
		hostname:   hostname,
	}

	apiService, err := api.NewApiServer(&api.Config{}, logger.With(zap.String("module", "api")), s.Consensus)
	if err != nil {
		return nil, err
	}

	s.Api = apiService

	s.Kubernetes = kubernetes.New(logger.With(zap.String("module", "kubernetes")), consensusService, strg)

	s.debugServer = grpc.NewServer()
	apipb.RegisterNodeDebugServiceServer(s.debugServer, s)

	logger.Info("Created SmalltownNode")

	return s, nil
}

func (s *SmalltownNode) Start(ctx context.Context) error {
	s.logger.Info("Starting Smalltown node")

	s.startDebugSvc()

	// TODO(lorenz): Abstracting enrolment sounds like a good idea, but ends up being painful
	// because of things like storage access. I'm keeping it this way until the more complex
	// enrolment procedures are fleshed out. This is also a bit panic()-happy, but there is really
	// no good way out of an invalid enrolment configuration.
	enrolmentPath, err := s.Storage.GetPathInPlace(storage.PlaceESP, "enrolment.pb")
	if err != nil {
		s.logger.Panic("ESP configuration partition not available", zap.Error(err))
	}
	enrolmentConfigRaw, err := ioutil.ReadFile(enrolmentPath)
	if os.IsNotExist(err) {
		enrolmentConfigRaw, err = ioutil.ReadFile("/sys/firmware/qemu_fw_cfg/by_name/com.nexantic.smalltown/enrolment.pb/raw")
	}
	if err == nil {
		// We have an enrolment file, let's check its contents
		var enrolmentConfig apipb.EnrolmentConfig
		if err := proto.Unmarshal(enrolmentConfigRaw, &enrolmentConfig); err != nil {
			s.logger.Panic("Invalid enrolment configuration provided", zap.Error(err))
		}
		s.enrolmentConfig = &enrolmentConfig
		// The enrolment secret is only zeroed after
		if len(enrolmentConfig.EnrolmentSecret) == 0 {
			return s.startFull()
		}
		return s.startEnrolling(ctx)
	} else if os.IsNotExist(err) {
		// This is ok like this, once a new cluster has been set up the initial node also generates
		// its own enrolment config
		return s.startForSetup(ctx)
	}
	// Unknown error reading enrolment config (disk issues/invalid configuration format/...)
	s.logger.Panic("Invalid enrolment configuration provided", zap.Error(err))
	panic("Unreachable")
}

func (s *SmalltownNode) startDebugSvc() {
	debugListenHost := fmt.Sprintf(":%v", common.DebugServicePort)
	debugListener, err := net.Listen("tcp", debugListenHost)
	if err != nil {
		s.logger.Fatal("failed to listen", zap.Error(err))
	}

	go func() {
		if err := s.debugServer.Serve(debugListener); err != nil {
			s.logger.Fatal("failed to serve", zap.Error(err))
		}
	}()
}

func (s *SmalltownNode) initHostname() error {
	if err := unix.Sethostname([]byte(s.hostname)); err != nil {
		return err
	}
	if err := ioutil.WriteFile("/etc/hosts", []byte(fmt.Sprintf("%v %v", "127.0.0.1", s.hostname)), 0644); err != nil {
		return err
	}
	return ioutil.WriteFile("/etc/machine-id", []byte(strings.TrimPrefix(s.hostname, "smalltown-")), 0644)
}

func (s *SmalltownNode) startEnrolling(ctx context.Context) error {
	s.logger.Info("Initializing subsystems for enrolment")
	s.state = common.StateEnrollMode

	nodeInfo, nodeID, err := s.InitializeNode(ctx)
	if err != nil {
		return err
	}

	s.hostname = nodeID
	if err := s.initHostname(); err != nil {
		return err
	}

	// We only support TPM2 at the moment, any abstractions here would be premature
	trustAgent := tpm2.TPM2Agent{}

	initializeOp := func() error {
		if err := trustAgent.Initialize(*nodeInfo, *s.enrolmentConfig); err != nil {
			s.logger.Warn("Failed to initialize integrity backend", zap.Error(err))
			return err
		}
		return nil
	}

	if err := backoff.Retry(initializeOp, getIntegrityBackoff()); err != nil {
		panic("invariant violated: integrity initialization retry can never fail")
	}

	enrolmentPath, err := s.Storage.GetPathInPlace(storage.PlaceESP, "enrolment.pb")
	if err != nil {
		panic(err)
	}

	s.enrolmentConfig.EnrolmentSecret = []byte{}
	s.enrolmentConfig.NodeId = nodeID

	enrolmentConfigRaw, err := proto.Marshal(s.enrolmentConfig)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(enrolmentPath, enrolmentConfigRaw, 0600); err != nil {
		return err
	}
	s.logger.Info("Node successfully enrolled")

	return nil
}

func (s *SmalltownNode) startForSetup(ctx context.Context) error {
	s.logger.Info("Setting up a new cluster")
	initData, nodeID, err := s.InitializeNode(ctx)
	if err != nil {
		return err
	}
	s.hostname = nodeID
	if err := s.initHostname(); err != nil {
		return err
	}

	if err := s.initNodeAPI(); err != nil {
		return err
	}

	// TODO: Use supervisor.Run for this
	go s.Containerd.Run()(context.TODO())

	dataPath, err := s.Storage.GetPathInPlace(storage.PlaceData, "etcd")
	if err != nil {
		return err
	}

	// Spin up etcd
	config := s.Consensus.GetConfig()
	config.NewCluster = true
	config.Name = s.hostname
	config.DataDir = dataPath
	s.Consensus.SetConfig(config)

	// Generate the cluster CA and store it to local storage.
	if err := s.Consensus.PrecreateCA(); err != nil {
		return err
	}

	err = s.Consensus.Start()
	if err != nil {
		return err
	}

	// Now that the cluster is up and running, we can persist the CA to the cluster.
	if err := s.Consensus.InjectCA(); err != nil {
		return err
	}

	if err := s.Api.BootstrapNewClusterHook(initData); err != nil {
		return err
	}

	if err := s.Kubernetes.NewCluster(); err != nil {
		return err
	}

	if err := s.Kubernetes.Start(); err != nil {
		return err
	}

	if err := s.Api.Start(); err != nil {
		s.logger.Error("Failed to start the API service", zap.Error(err))
		return err
	}

	enrolmentPath, err := s.Storage.GetPathInPlace(storage.PlaceESP, "enrolment.pb")
	if err != nil {
		panic(err)
	}

	masterCert, err := s.Api.GetMasterCert()
	if err != nil {
		return err
	}

	ip, err := s.Network.GetIP(ctx, true)
	if err != nil {
		return fmt.Errorf("could not get node IP: %v", err)
	}
	enrolmentConfig := &apipb.EnrolmentConfig{
		EnrolmentSecret: []byte{}, // First node is always already enrolled
		MastersCert:     masterCert,
		MasterIps:       [][]byte{[]byte(*ip)},
		NodeId:          nodeID,
	}
	enrolmentConfigRaw, err := proto.Marshal(enrolmentConfig)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(enrolmentPath, enrolmentConfigRaw, 0600); err != nil {
		return err
	}
	masterCertFingerprint := sha512.Sum512_256(masterCert)
	s.logger.Info("New Smalltown cluster successfully bootstrapped", zap.Binary("fingerprint", masterCertFingerprint[:]))

	return nil
}

func (s *SmalltownNode) generateNodeID() ([]byte, string, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 127)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return []byte{}, "", fmt.Errorf("Failed to generate serial number: %w", err)
	}

	pubKey, privKeyRaw, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return []byte{}, "", err
	}
	privkey, err := x509.MarshalPKCS8PrivateKey(privKeyRaw)
	if err != nil {
		return []byte{}, "", err
	}

	nodeKeyPath, err := s.Storage.GetPathInPlace(storage.PlaceData, "node-key.der")
	if err != nil {
		return []byte{}, "", err
	}

	if err := ioutil.WriteFile(nodeKeyPath, privkey, 0600); err != nil {
		return []byte{}, "", fmt.Errorf("failed to write node key: %w", err)
	}

	name := "smalltown-" + hex.EncodeToString([]byte(pubKey[:16]))

	// This has no SANs because it authenticates by public key, not by name
	nodeCert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			// We identify nodes by their ID public keys (not hashed since a strong hash is longer and serves no benefit)
			CommonName: name,
		},
		IsCA:                  false,
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              unknownNotAfter,
		// Certificate is used both as server & client
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}
	cert, err := x509.CreateCertificate(rand.Reader, nodeCert, nodeCert, pubKey, privKeyRaw)
	if err != nil {
		return []byte{}, "", err
	}

	nodeCertPath, err := s.Storage.GetPathInPlace(storage.PlaceData, "node.der")
	if err != nil {
		return []byte{}, "", err
	}

	if err := ioutil.WriteFile(nodeCertPath, cert, 0600); err != nil {
		return []byte{}, "", fmt.Errorf("failed to write node cert: %w", err)
	}
	return cert, name, nil
}

func (s *SmalltownNode) initNodeAPI() error {
	certPath, err := s.Storage.GetPathInPlace(storage.PlaceData, "node.der")
	if err != nil {
		s.logger.Panic("Invariant violated: Data is available once this is called")
	}
	keyPath, err := s.Storage.GetPathInPlace(storage.PlaceData, "node-key.der")
	if err != nil {
		s.logger.Panic("Invariant violated: Data is available once this is called")
	}

	certRaw, err := ioutil.ReadFile(certPath)
	if err != nil {
		return err
	}
	privKeyRaw, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return err
	}

	var nodeID tls.Certificate

	cert, err := x509.ParseCertificate(certRaw)
	if err != nil {
		return err
	}

	privKey, err := x509.ParsePKCS8PrivateKey(privKeyRaw)
	if err != nil {
		return err
	}

	nodeID.Certificate = [][]byte{certRaw}
	nodeID.PrivateKey = privKey
	nodeID.Leaf = cert

	secureTransport := &tls.Config{
		Certificates:       []tls.Certificate{nodeID},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		InsecureSkipVerify: true,
		// Critical function, please review any changes with care
		// TODO(lorenz): Actively check that this actually provides the security guarantees that we need
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			for _, cert := range rawCerts {
				// X.509 certificates in DER can be compared like this since DER has a unique representation
				// for each certificate.
				if bytes.Equal(cert, s.enrolmentConfig.MastersCert) {
					return nil
				}
			}
			return errors.New("failed to find authorized NMS certificate")
		},
		MinVersion: tls.VersionTLS13,
	}
	secureTransportCreds := credentials.NewTLS(secureTransport)

	masterListenHost := fmt.Sprintf(":%d", common.NodeServicePort)
	lis, err := net.Listen("tcp", masterListenHost)
	if err != nil {
		s.logger.Fatal("failed to listen", zap.Error(err))
	}

	nodeGRPCServer := grpc.NewServer(grpc.Creds(secureTransportCreds))
	apipb.RegisterNodeServiceServer(nodeGRPCServer, s)
	go func() {
		if err := nodeGRPCServer.Serve(lis); err != nil {
			panic(err) // Can only happen during initialization and is always fatal
		}
	}()
	return nil
}

func getIntegrityBackoff() *backoff.ExponentialBackOff {
	unlockBackoff := backoff.NewExponentialBackOff()
	unlockBackoff.MaxElapsedTime = time.Duration(0)
	unlockBackoff.InitialInterval = 5 * time.Second
	unlockBackoff.MaxInterval = 5 * time.Minute
	return unlockBackoff
}

func (s *SmalltownNode) startFull() error {
	s.logger.Info("Initializing subsystems for production")
	s.state = common.StateJoined

	s.hostname = s.enrolmentConfig.NodeId
	if err := s.initHostname(); err != nil {
		return err
	}

	trustAgent := tpm2.TPM2Agent{}
	unlockOp := func() error {
		unlockKey, err := trustAgent.Unlock(*s.enrolmentConfig)
		if err != nil {
			s.logger.Warn("Failed to unlock", zap.Error(err))
			return err
		}
		if err := s.Storage.MountData(unlockKey); err != nil {
			s.logger.Panic("Failed to mount storage", zap.Error(err))
			return err
		}
		return nil
	}

	if err := backoff.Retry(unlockOp, getIntegrityBackoff()); err != nil {
		s.logger.Panic("Invariant violated: Unlock retry can never fail")
	}

	s.initNodeAPI()

	// TODO: Use supervisor.Run for this
	go s.Containerd.Run()(context.TODO())

	err := s.Consensus.Start()
	if err != nil {
		return err
	}

	err = s.Api.Start()
	if err != nil {
		s.logger.Error("Failed to start the API service", zap.Error(err))
		return err
	}

	err = s.Kubernetes.Start()
	if err != nil {
		s.logger.Error("Failed to start the Kubernetes Service", zap.Error(err))
	}

	return nil
}

func (s *SmalltownNode) Stop() error {
	s.logger.Info("Stopping Smalltown node")
	return nil
}
