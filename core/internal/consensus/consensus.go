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

// package consensus manages the embedded etcd cluster.
package consensus

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/binary"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/namespace"
	"go.etcd.io/etcd/embed"
	"go.etcd.io/etcd/etcdserver/api/membership"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/proxy/grpcproxy/adapter"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sys/unix"

	"git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"git.monogon.dev/source/nexantic.git/core/internal/common/service"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus/ca"
)

const (
	DefaultClusterToken = "SIGNOS"
	DefaultLogger       = "zap"
)

const (
	CAPath      = "ca.pem"
	CertPath    = "cert.pem"
	KeyPath     = "cert-key.pem"
	CRLPath     = "ca-crl.der"
	CRLSwapPath = "ca-crl.der.swp"
)

const (
	LocalListenerURL = "unix:///consensus/listener.sock:0"
)

type (
	Service struct {
		*service.BaseService

		etcd  *embed.Etcd
		kv    clientv3.KV
		ready atomic.Bool

		// bootstrapCA and bootstrapCert cache the etcd cluster CA data during bootstrap.
		bootstrapCA   *ca.CA
		bootstrapCert []byte

		watchCRLTicker *time.Ticker
		lastCRL        []byte

		config *Config
	}

	Config struct {
		Name           string
		DataDir        string
		InitialCluster string
		NewCluster     bool
		ExternalHost   string
		ListenHost     string
	}

	Member struct {
		ID      uint64
		Name    string
		Address string
		Synced  bool
	}
)

func NewConsensusService(config Config, logger *zap.Logger) (*Service, error) {
	consensusServer := &Service{
		config: &config,
	}
	consensusServer.BaseService = service.NewBaseService("consensus", logger, consensusServer)

	return consensusServer, nil
}

func peerURL(host string) url.URL {
	return url.URL{Scheme: "https", Host: fmt.Sprintf("%s:%d", host, common.ConsensusPort)}
}

func (s *Service) OnStart() error {
	// See: https://godoc.org/github.com/coreos/etcd/embed#Config

	if s.config == nil {
		return errors.New("config for consensus is nil")
	}

	cfg := embed.NewConfig()

	cfg.PeerTLSInfo.CertFile = filepath.Join(s.config.DataDir, CertPath)
	cfg.PeerTLSInfo.KeyFile = filepath.Join(s.config.DataDir, KeyPath)
	cfg.PeerTLSInfo.TrustedCAFile = filepath.Join(s.config.DataDir, CAPath)
	cfg.PeerTLSInfo.ClientCertAuth = true
	cfg.PeerTLSInfo.CRLFile = filepath.Join(s.config.DataDir, CRLPath)

	lastCRL, err := ioutil.ReadFile(cfg.PeerTLSInfo.CRLFile)
	if err != nil {
		return fmt.Errorf("failed to read etcd CRL: %w", err)
	}
	s.lastCRL = lastCRL

	// Expose etcd to local processes
	if err := os.MkdirAll("/consensus", 0700); err != nil {
		return fmt.Errorf("Failed to create consensus runtime state directory: %w", err)
	}
	listenerURL, err := url.Parse(LocalListenerURL)
	if err != nil {
		panic(err)
	}
	cfg.LCUrls = []url.URL{*listenerURL}

	cfg.APUrls = []url.URL{peerURL(s.config.ExternalHost)}
	cfg.LPUrls = []url.URL{peerURL(s.config.ListenHost)}
	cfg.ACUrls = []url.URL{}

	cfg.Dir = s.config.DataDir
	cfg.InitialClusterToken = DefaultClusterToken
	cfg.Name = s.config.Name

	// Only relevant if creating or joining a cluster; otherwise settings will be ignored
	if s.config.NewCluster {
		cfg.ClusterState = "new"
		cfg.InitialCluster = cfg.InitialClusterFromName(cfg.Name)
	} else if s.config.InitialCluster != "" {
		cfg.ClusterState = "existing"
		cfg.InitialCluster = s.config.InitialCluster
	}

	cfg.Logger = DefaultLogger
	cfg.ZapLoggerBuilder = embed.NewZapCoreLoggerBuilder(
		s.Logger.With(zap.String("component", "etcd")).WithOptions(zap.IncreaseLevel(zapcore.WarnLevel)),
		s.Logger.Core(),
		nil,
	)

	server, err := embed.StartEtcd(cfg)
	if err != nil {
		return err
	}
	s.etcd = server

	// Override the logger
	//*server.GetLogger() = *s.Logger.With(zap.String("component", "etcd"))
	// TODO(leo): can we uncomment this?

	go func() {
		s.Logger.Info("waiting for etcd to become ready")
		<-s.etcd.Server.ReadyNotify()
		s.ready.Store(true)
		s.Logger.Info("etcd is now ready")
	}()

	// Inject kv client
	s.kv = clientv3.NewKVFromKVClient(adapter.KvServerToKvClient(s.etcd.Server), nil)

	// Start CRL watcher
	go s.watchCRL()
	ctx := context.TODO()
	go s.autoPromote(ctx)

	return nil
}

// WriteCertificateFiles writes the given node certificate data to local storage
// such that it can be used by the embedded etcd server.
// Unfortunately, we cannot pass the certificates directly to etcd.
func (s *Service) WriteCertificateFiles(certs *api.ConsensusCertificates) error {
	if err := ioutil.WriteFile(filepath.Join(s.config.DataDir, CRLPath), certs.Crl, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(s.config.DataDir, CertPath),
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certs.Cert}), 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(s.config.DataDir, KeyPath),
		pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: certs.Key}), 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(s.config.DataDir, CAPath),
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certs.Ca}), 0600); err != nil {
		return err
	}
	return nil
}

// PrecreateCA generates the etcd cluster certificate authority and writes it to local storage.
func (s *Service) PrecreateCA(extIP net.IP) error {
	// Provision an etcd CA
	etcdRootCA, err := ca.New("Smalltown etcd Root CA")
	if err != nil {
		return err
	}
	cert, privkey, err := etcdRootCA.IssueCertificate(s.config.ExternalHost, extIP)
	if err != nil {
		return fmt.Errorf("failed to self-issue a certificate: %w", err)
	}
	if err := os.MkdirAll(s.config.DataDir, 0700); err != nil {
		return fmt.Errorf("failed to create consensus data dir: %w", err)
	}
	// Preserve certificate for later injection
	s.bootstrapCert = cert
	if err := s.WriteCertificateFiles(&api.ConsensusCertificates{
		Ca:   etcdRootCA.CACertRaw,
		Crl:  etcdRootCA.CRLRaw,
		Cert: cert,
		Key:  privkey,
	}); err != nil {
		return fmt.Errorf("failed to setup certificates: %w", err)
	}
	s.bootstrapCA = etcdRootCA
	return nil
}

const (
	caPathEtcd    = "/etcd-ca/ca.der"
	caKeyPathEtcd = "/etcd-ca/ca-key.der"
	crlPathEtcd   = "/etcd-ca/crl.der"

	// This prefix stores the individual certs the etcd CA has issued.
	certPrefixEtcd = "/etcd-ca/certs"
)

// InjectCA copies the CA from data cached during PrecreateCA to etcd.
// Requires a previous call to PrecreateCA.
func (s *Service) InjectCA() error {
	if s.bootstrapCA == nil || s.bootstrapCert == nil {
		panic("bootstrapCA or bootstrapCert are nil - missing PrecreateCA call?")
	}
	if _, err := s.kv.Put(context.Background(), caPathEtcd, string(s.bootstrapCA.CACertRaw)); err != nil {
		return err
	}
	// TODO(lorenz): Should be wrapped by the master key
	if _, err := s.kv.Put(context.Background(), caKeyPathEtcd, string([]byte(*s.bootstrapCA.PrivateKey))); err != nil {
		return err
	}
	if _, err := s.kv.Put(context.Background(), crlPathEtcd, string(s.bootstrapCA.CRLRaw)); err != nil {
		return err
	}
	certVal, err := x509.ParseCertificate(s.bootstrapCert)
	if err != nil {
		return err
	}
	serial := hex.EncodeToString(certVal.SerialNumber.Bytes())
	if _, err := s.kv.Put(context.Background(), path.Join(certPrefixEtcd, serial), string(s.bootstrapCert)); err != nil {
		return fmt.Errorf("failed to persist certificate: %w", err)
	}
	// Clear out bootstrap CA after injecting
	s.bootstrapCA = nil
	s.bootstrapCert = []byte{}
	return nil
}

func (s *Service) etcdGetSingle(path string) ([]byte, int64, error) {
	res, err := s.kv.Get(context.Background(), path)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to get key from etcd: %w", err)
	}
	if len(res.Kvs) != 1 {
		return nil, -1, errors.New("key not available or multiple keys returned")
	}
	return res.Kvs[0].Value, res.Kvs[0].ModRevision, nil
}

func (s *Service) getCAFromEtcd() (*ca.CA, int64, error) {
	// TODO: Technically this could be done in a single request, but it's more logic
	caCert, _, err := s.etcdGetSingle(caPathEtcd)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to get CA certificate from etcd: %w", err)
	}
	caKey, _, err := s.etcdGetSingle(caKeyPathEtcd)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to get CA key from etcd: %w", err)
	}
	// TODO: Unwrap CA key once wrapping is implemented
	crl, crlRevision, err := s.etcdGetSingle(crlPathEtcd)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to get CRL from etcd: %w", err)
	}
	idCA, err := ca.FromCertificates(caCert, caKey, crl)
	if err != nil {
		return nil, -1, fmt.Errorf("failed to take CA online: %w", err)
	}
	return idCA, crlRevision, nil
}

// ProvisionMember sets up and returns provisioning data to join another node into the consensus.
// It issues PKI material, creates a static cluster bootstrap specification string (known as initial-cluster in etcd)
// and adds the new node as a learner (non-voting) member to the cluster. Once the new node has caught up with the
// cluster it is automatically promoted to a voting member by the autoPromote process.
func (s *Service) ProvisionMember(name string, ip net.IP) (*api.ConsensusCertificates, string, error) {
	idCA, _, err := s.getCAFromEtcd()
	if err != nil {
		return nil, "", fmt.Errorf("failed to get consensus CA: %w", err)
	}
	cert, key, err := idCA.IssueCertificate(name, ip)
	if err != nil {
		return nil, "", fmt.Errorf("failed to issue certificate: %w", err)
	}
	certVal, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse just-issued consensus cert: %w", err)
	}
	serial := hex.EncodeToString(certVal.SerialNumber.Bytes())
	if _, err := s.kv.Put(context.Background(), path.Join(certPrefixEtcd, serial), string(cert)); err != nil {
		// We issued a certificate, but failed to persist it. Return an error and forget it ever happened.
		return nil, "", fmt.Errorf("failed to persist certificate: %w", err)
	}

	currentMembers := s.etcd.Server.Cluster().Members()
	var memberStrs []string
	for _, member := range currentMembers {
		memberStrs = append(memberStrs, fmt.Sprintf("%v=%v", member.Name, member.PickPeerURL()))
	}
	apURL := peerURL(ip.String())
	memberStrs = append(memberStrs, fmt.Sprintf("%s=%s", name, apURL.String()))

	pubKeyPrefix, err := common.IDKeyPrefixFromName(name)
	if err != nil {
		return nil, "", fmt.Errorf("invalid new node name: %v", err)
	}

	crl, _, err := s.etcdGetSingle(crlPathEtcd)

	_, err = s.etcd.Server.AddMember(context.Background(), membership.Member{
		RaftAttributes: membership.RaftAttributes{
			PeerURLs:  types.URLs{apURL}.StringSlice(),
			IsLearner: true,
		},
		Attributes: membership.Attributes{Name: name},
		ID:         types.ID(binary.BigEndian.Uint64(pubKeyPrefix[:8])),
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to provision member: %w", err)
	}
	return &api.ConsensusCertificates{
		Ca:   idCA.CACertRaw,
		Cert: cert,
		Crl:  crl,
		Key:  key,
	}, strings.Join(memberStrs, ","), nil
}

func (s *Service) RevokeCertificate(hostname string) error {
	rand.Seed(time.Now().UnixNano())
	for {
		idCA, crlRevision, err := s.getCAFromEtcd()
		if err != nil {
			return err
		}
		allIssuedCerts, err := s.kv.Get(context.Background(), certPrefixEtcd, clientv3.WithPrefix())
		for _, cert := range allIssuedCerts.Kvs {
			certVal, err := x509.ParseCertificate(cert.Value)
			if err != nil {
				s.Logger.Error("Failed to parse previously issued certificate, this is a security risk", zap.Error(err))
				continue
			}
			for _, dnsName := range certVal.DNSNames {
				if dnsName == hostname {
					// Revoke this
					if err := idCA.Revoke(certVal.SerialNumber); err != nil {
						// We need to fail if any single revocation fails otherwise outer applications
						// have no chance of calling this safely
						return err
					}
				}
			}
		}
		// TODO(leo): this needs a test
		cmp := clientv3.Compare(clientv3.ModRevision(crlPathEtcd), "=", crlRevision)
		op := clientv3.OpPut(crlPathEtcd, string(idCA.CRLRaw))
		res, err := s.kv.Txn(context.Background()).If(cmp).Then(op).Commit()
		if err != nil {
			return fmt.Errorf("failed to persist new CRL in etcd: %w", err)
		}
		if res.Succeeded { // Transaction has succeeded
			break
		}
		// Sleep a random duration between 0 and 300ms to reduce serialization failures
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	}
	return nil
}

func (s *Service) watchCRL() {
	// TODO(lorenz): Change etcd client to WatchableKV and make this an actual watch
	// This needs changes in more places, so leaving it now
	s.watchCRLTicker = time.NewTicker(30 * time.Second)
	for range s.watchCRLTicker.C {
		crl, _, err := s.etcdGetSingle(crlPathEtcd)
		if err != nil {
			s.Logger.Warn("Failed to check for new CRL", zap.Error(err))
			continue
		}
		// This is cryptographic material but not secret, so no constant time compare necessary here
		if !bytes.Equal(crl, s.lastCRL) {
			if err := ioutil.WriteFile(filepath.Join(s.config.DataDir, CRLSwapPath), crl, 0600); err != nil {
				s.Logger.Warn("Failed to write updated CRL", zap.Error(err))
			}
			// This uses unix.Rename to guarantee a particular atomic update behavior
			if err := unix.Rename(filepath.Join(s.config.DataDir, CRLSwapPath), filepath.Join(s.config.DataDir, CRLPath)); err != nil {
				s.Logger.Warn("Failed to atomically swap updated CRL", zap.Error(err))
			}
		}
	}
}

// autoPromote automatically promotes learning (non-voting) members to voting members. etcd currently lacks auto-promote
// capabilities (https://github.com/etcd-io/etcd/issues/10537) so we need to do this ourselves.
func (s *Service) autoPromote(ctx context.Context) {
	promoteTicker := time.NewTicker(5 * time.Second)
	go func() {
		<-ctx.Done()
		promoteTicker.Stop()
	}()
	for range promoteTicker.C {
		if s.etcd.Server.Leader() != s.etcd.Server.ID() {
			continue
		}
		for _, member := range s.etcd.Server.Cluster().Members() {
			if member.IsLearner {
				// We always call PromoteMember since the metadata necessary to decide if we should is private.
				// Luckily etcd already does sanity checks internally and will refuse to promote nodes that aren't
				// connected or are still behind on transactions.
				if _, err := s.etcd.Server.PromoteMember(context.Background(), uint64(member.ID)); err != nil {
					s.Logger.Info("Failed to promote consensus node", zap.String("node", member.Name), zap.Error(err))
				}
				s.Logger.Info("Promoted new consensus node", zap.String("node", member.Name))
			}
		}
	}
}

func (s *Service) OnStop() error {
	s.watchCRLTicker.Stop()
	s.etcd.Close()

	return nil
}

// IsProvisioned returns whether the node has been setup before and etcd has a data directory
func (s *Service) IsProvisioned() bool {
	_, err := os.Stat(s.config.DataDir)

	return !os.IsNotExist(err)
}

// IsReady returns whether etcd is ready and synced
func (s *Service) IsReady() bool {
	return s.ready.Load()
}

// GetConfig returns the current consensus config
func (s *Service) GetConfig() Config {
	return *s.config
}

// SetConfig sets the consensus config. Changes are only applied when the service is restarted.
func (s *Service) SetConfig(config Config) {
	s.config = &config
}

func (s *Service) GetStore(module, space string) clientv3.KV {
	return namespace.NewKV(s.kv, fmt.Sprintf("%s:%s", module, space))
}
