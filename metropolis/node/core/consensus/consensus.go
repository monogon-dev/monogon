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

// Package consensus implements a managed etcd cluster member service, with a self-hosted CA system for issuing peer
// certificates. Currently each Metropolis node runs an etcd member, and connects to the etcd member locally over a
// domain socket.
//
// The service supports two modes of startup:
//  - initializing a new cluster, by bootstrapping the CA in memory, starting a cluster, committing the CA to etcd
//    afterwards, and saving the new node's certificate to local storage
//  - joining an existing cluster, using certificates from local storage and loading the CA from etcd. This flow is also
//    used when the node joins a cluster for the first time (then the certificates required must be provisioned
//    externally before starting the consensus service).
//
// Regardless of how the etcd member service was started, the resulting running service is further managed and used
// in the same way.
//
package consensus

import (
	"context"
	"encoding/pem"
	"fmt"
	"net/url"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/embed"
	"go.uber.org/atomic"

	node "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/consensus/ca"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

const (
	DefaultClusterToken = "METROPOLIS"
	DefaultLogger       = "zap"
)

// Service is the etcd cluster member service.
type Service struct {
	// The configuration with which the service was started. This is immutable.
	config *Config

	// stateMu guards state. This is locked internally on public methods of Service that require access to state. The
	// state might be recreated on service restart.
	stateMu sync.Mutex
	state   *state
}

// state is the runtime state of a running etcd member.
type state struct {
	etcd  *embed.Etcd
	ready atomic.Bool

	ca *ca.CA
	// cl is an etcd client that loops back to the localy running etcd server. This runs over the Client unix domain
	// socket that etcd starts.
	cl *clientv3.Client
}

type Config struct {
	// Data directory (persistent, encrypted storage) for etcd.
	Data *localstorage.DataEtcdDirectory
	// Ephemeral directory for etcd.
	Ephemeral *localstorage.EphemeralConsensusDirectory

	// Name is the cluster name. This must be the same amongst all etcd members within one cluster.
	Name string
	// NewCluster selects whether the etcd member will start a new cluster and bootstrap a CA and the first member
	// certificate, or load existing PKI certificates from disk.
	NewCluster bool
	// Port is the port at which this cluster member will listen for other members. If zero, defaults to the global
	// Metropolis setting.
	Port int

	// externalHost is used by tests to override the address at which etcd should listen for peer connections.
	externalHost string
}

func New(config Config) *Service {
	return &Service{
		config: &config,
	}
}

// configure transforms the service configuration into an embedded etcd configuration. This is pure and side effect
// free.
func (s *Service) configure(ctx context.Context) (*embed.Config, error) {
	if err := s.config.Ephemeral.MkdirAll(0700); err != nil {
		return nil, fmt.Errorf("failed to create ephemeral directory: %w", err)
	}
	if err := s.config.Data.MkdirAll(0700); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	if s.config.Name == "" {
		return nil, fmt.Errorf("Name not set")
	}
	port := s.config.Port
	if port == 0 {
		port = node.ConsensusPort
	}

	cfg := embed.NewConfig()

	cfg.Name = s.config.Name
	cfg.Dir = s.config.Data.Data.FullPath()
	cfg.InitialClusterToken = DefaultClusterToken

	cfg.PeerTLSInfo.CertFile = s.config.Data.PeerPKI.Certificate.FullPath()
	cfg.PeerTLSInfo.KeyFile = s.config.Data.PeerPKI.Key.FullPath()
	cfg.PeerTLSInfo.TrustedCAFile = s.config.Data.PeerPKI.CACertificate.FullPath()
	cfg.PeerTLSInfo.ClientCertAuth = true
	cfg.PeerTLSInfo.CRLFile = s.config.Data.PeerCRL.FullPath()

	cfg.LCUrls = []url.URL{{
		Scheme: "unix",
		Path:   s.config.Ephemeral.ClientSocket.FullPath() + ":0",
	}}
	cfg.ACUrls = []url.URL{}
	cfg.LPUrls = []url.URL{{
		Scheme: "https",
		Host:   fmt.Sprintf("[::]:%d", port),
	}}

	// Always listen on the address pointed to by our name - unless running in
	// tests, where we can't control our hostname easily.
	externalHost := fmt.Sprintf("%s:%d", s.config.Name, port)
	if s.config.externalHost != "" {
		externalHost = fmt.Sprintf("%s:%d", s.config.externalHost, port)
	}
	cfg.APUrls = []url.URL{{
		Scheme: "https",
		Host:   externalHost,
	}}

	if s.config.NewCluster {
		cfg.ClusterState = "new"
		cfg.InitialCluster = cfg.InitialClusterFromName(cfg.Name)
	} else {
		cfg.ClusterState = "existing"
	}

	// TODO(q3k): pipe logs from etcd to supervisor.RawLogger via a file.
	cfg.Logger = DefaultLogger
	cfg.LogOutputs = []string{"stderr"}

	return cfg, nil
}

// Run is a Supervisor runnable that starts the etcd member service. It will become healthy once the member joins the
// cluster successfully.
func (s *Service) Run(ctx context.Context) error {
	st := &state{
		ready: *atomic.NewBool(false),
	}
	s.stateMu.Lock()
	s.state = st
	s.stateMu.Unlock()

	if s.config.NewCluster {
		// Create certificate if absent. It can only be present if we attempt
		// to re-start the service in NewCluster after a failure. This can
		// happen if etcd crashed or failed to start up before (eg. because of
		// networking not having settled yet).
		absent, err := s.config.Data.PeerPKI.AllAbsent()
		if err != nil {
			return fmt.Errorf("checking certificate existence: %w", err)
		}

		if absent {
			// Generate CA, keep in memory, write it down in etcd later.
			st.ca, err = ca.New("Metropolis etcd peer Root CA")
			if err != nil {
				return fmt.Errorf("when creating new cluster's peer CA: %w", err)
			}

			cert, key, err := st.ca.Issue(ctx, nil, s.config.Name)
			if err != nil {
				return fmt.Errorf("when issuing new cluster's first certificate: %w", err)
			}

			if err := s.config.Data.PeerPKI.MkdirAll(0700); err != nil {
				return fmt.Errorf("when creating PKI directory: %w", err)
			}
			if err := s.config.Data.PeerPKI.CACertificate.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: st.ca.CACertRaw}), 0600); err != nil {
				return fmt.Errorf("when writing CA certificate to disk: %w", err)
			}
			if err := s.config.Data.PeerPKI.Certificate.Write(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert}), 0600); err != nil {
				return fmt.Errorf("when writing certificate to disk: %w", err)
			}
			if err := s.config.Data.PeerPKI.Key.Write(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key}), 0600); err != nil {
				return fmt.Errorf("when writing certificate to disk: %w", err)
			}
		}
	}

	// Expect certificate to be present on disk.
	present, err := s.config.Data.PeerPKI.AllExist()
	if err != nil {
		return fmt.Errorf("checking certificate existence: %w", err)
	}
	if !present {
		return fmt.Errorf("etcd starting without fully ready certificates - aborted NewCluster or corrupted local storage?")
	}

	cfg, err := s.configure(ctx)
	if err != nil {
		return fmt.Errorf("when configuring etcd: %w", err)
	}

	server, err := embed.StartEtcd(cfg)
	keep := false
	defer func() {
		if !keep && server != nil {
			server.Close()
		}
	}()
	if err != nil {
		return fmt.Errorf("failed to start etcd: %w", err)
	}
	st.etcd = server

	okay := true
	select {
	case <-st.etcd.Server.ReadyNotify():
	case <-ctx.Done():
		okay = false
	}

	if !okay {
		supervisor.Logger(ctx).Info("context done, aborting wait")
		return ctx.Err()
	}

	socket := s.config.Ephemeral.ClientSocket.FullPath()
	cl, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("unix://%s:0", socket)},
		DialTimeout: time.Second,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to new etcd instance: %w", err)
	}
	st.cl = cl

	if s.config.NewCluster {
		if st.ca == nil {
			panic("peerCA has not been generated")
		}

		// Save new CA into etcd.
		err = st.ca.Save(ctx, cl.KV)
		if err != nil {
			return fmt.Errorf("failed to save new CA to etcd: %w", err)
		}
	} else {
		// Load existing CA from etcd.
		st.ca, err = ca.Load(ctx, cl.KV)
		if err != nil {
			return fmt.Errorf("failed to load CA from etcd: %w", err)
		}
	}

	// Start CRL watcher.
	if err := supervisor.Run(ctx, "crl", s.watchCRL); err != nil {
		return fmt.Errorf("failed to start CRL watcher: %w", err)
	}
	// Start autopromoter.
	if err := supervisor.Run(ctx, "autopromoter", s.autopromoter); err != nil {
		return fmt.Errorf("failed to start autopromoter: %w", err)
	}

	supervisor.Logger(ctx).Info("etcd is now ready")
	keep = true
	st.ready.Store(true)
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	<-ctx.Done()
	st.etcd.Close()
	return ctx.Err()
}

// watchCRL is a sub-runnable of the etcd cluster member service that updates the on-local-storage CRL to match the
// newest available version in etcd.
func (s *Service) watchCRL(ctx context.Context) error {
	s.stateMu.Lock()
	cl := s.state.cl
	ca := s.state.ca
	s.stateMu.Unlock()

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	for e := range ca.WaitCRLChange(ctx, cl.KV, cl.Watcher) {
		if e.Err != nil {
			return fmt.Errorf("watching CRL: %w", e.Err)
		}

		if err := s.config.Data.PeerCRL.Write(e.CRL, 0600); err != nil {
			return fmt.Errorf("saving CRL: %w", err)
		}
	}

	// unreachable
	return nil
}

func (s *Service) autopromoter(ctx context.Context) error {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	autopromote := func() {
		s.stateMu.Lock()
		st := s.state
		s.stateMu.Unlock()

		if st.etcd.Server.Leader() != st.etcd.Server.ID() {
			return
		}

		for _, member := range st.etcd.Server.Cluster().Members() {
			if !member.IsLearner {
				continue
			}

			// We always call PromoteMember since the metadata necessary to decide if we should is private.
			// Luckily etcd already does sanity checks internally and will refuse to promote nodes that aren't
			// connected or are still behind on transactions.
			if _, err := st.etcd.Server.PromoteMember(ctx, uint64(member.ID)); err != nil {
				supervisor.Logger(ctx).Infof("Failed to promote consensus node %s: %v", member.Name, err)
			} else {
				supervisor.Logger(ctx).Infof("Promoted new consensus node %s", member.Name)
			}
		}
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			autopromote()
		}
	}
}

// IsReady returns whether etcd is ready and synced
func (s *Service) IsReady() bool {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	if s.state == nil {
		return false
	}
	return s.state.ready.Load()
}

func (s *Service) WaitReady(ctx context.Context) error {
	// TODO(q3k): reimplement the atomic ready flag as an event synchronization mechanism
	if s.IsReady() {
		return nil
	}
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			if s.IsReady() {
				return nil
			}
		}
	}
}

func (s *Service) Client() client.Namespaced {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	// 'namespaced' is the root of all namespaced clients within the etcd K/V
	// store, with further paths in a colon-separated format, eg.:
	//   namespaced:example/
	//   namespaced:foo:bar:baz/
	client, err := client.NewLocal(s.state.cl).Sub("namespaced")
	if err != nil {
		// This error can only happen due to a malformed path, which is
		// constant. Thus, this is a programming error and we panic.
		panic(fmt.Errorf("Could not get consensus etcd client: %v", err))
	}
	return client
}

func (s *Service) Cluster() clientv3.Cluster {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	return s.state.cl.Cluster
}

// MemberInfo returns information about this etcd cluster member: its ID and name. This will block until this
// information is available (ie. the cluster status is Ready).
func (s *Service) MemberInfo(ctx context.Context) (id uint64, name string, err error) {
	if err = s.WaitReady(ctx); err != nil {
		err = fmt.Errorf("when waiting for cluster readiness: %w", err)
		return
	}

	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	id = uint64(s.state.etcd.Server.ID())
	name = s.config.Name
	return
}
