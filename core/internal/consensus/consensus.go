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

package consensus

import (
	"context"
	"fmt"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/namespace"
	"go.etcd.io/etcd/embed"
	"go.etcd.io/etcd/etcdserver/api/membership"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/proxy/grpcproxy/adapter"
	"go.uber.org/zap"
	"net/url"
	"os"
	"strings"
)

const (
	DefaultClusterToken = "SIGNOS"
	DefaultLogger       = "zap"
)

type (
	Service struct {
		*common.BaseService

		etcd  *embed.Etcd
		kv    clientv3.KV
		ready bool

		config *Config
	}

	Config struct {
		Name           string
		DataDir        string
		InitialCluster string
		NewCluster     bool

		ExternalHost string
		ListenHost   string
		ListenPort   uint16
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
	consensusServer.BaseService = common.NewBaseService("consensus", logger, consensusServer)

	return consensusServer, nil
}

func (s *Service) OnStart() error {
	if s.config == nil {
		return errors.New("config for consensus is nil")
	}

	cfg := embed.NewConfig()

	// Reset LCUrls because we don't want to expose any client
	cfg.LCUrls = nil

	apURL, err := url.Parse(fmt.Sprintf("http://%s:%d", s.config.ExternalHost, s.config.ListenPort))
	if err != nil {
		return errors.Wrap(err, "invalid external_host or listen_port")
	}

	lpURL, err := url.Parse(fmt.Sprintf("http://%s:%d", s.config.ListenHost, s.config.ListenPort))
	if err != nil {
		return errors.Wrap(err, "invalid listen_host or listen_port")
	}
	cfg.APUrls = []url.URL{*apURL}
	cfg.LPUrls = []url.URL{*lpURL}
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

	server, err := embed.StartEtcd(cfg)
	if err != nil {
		return err
	}
	s.etcd = server

	// Override the logger
	//*server.GetLogger() = *s.Logger.With(zap.String("component", "etcd"))

	go func() {
		s.Logger.Info("waiting for etcd to become ready")
		<-s.etcd.Server.ReadyNotify()
		s.ready = true
		s.Logger.Info("etcd is now ready")
	}()

	// Inject kv client
	s.kv = clientv3.NewKVFromKVClient(adapter.KvServerToKvClient(s.etcd.Server), nil)

	return nil
}

func (s *Service) OnStop() error {
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
	return s.ready
}

// AddMember adds a new etcd member to the cluster
func (s *Service) AddMember(ctx context.Context, name string, url string) (uint64, error) {
	urls, err := types.NewURLs([]string{url})
	if err != nil {
		return 0, err
	}

	member := membership.NewMember(name, urls, DefaultClusterToken, nil)

	_, err = s.etcd.Server.AddMember(ctx, *member)
	if err != nil {
		return 0, err
	}

	return uint64(member.ID), nil
}

// RemoveMember removes a member from the etcd cluster
func (s *Service) RemoveMember(ctx context.Context, id uint64) error {
	_, err := s.etcd.Server.RemoveMember(ctx, id)
	return err
}

// Health returns the current cluster health
func (s *Service) Health() {
}

// GetConfig returns the current consensus config
func (s *Service) GetConfig() Config {
	return *s.config
}

// SetConfig sets the consensus config. Changes are only applied when the service is restarted.
func (s *Service) SetConfig(config Config) {
	s.config = &config
}

// GetInitialClusterString returns the InitialCluster string that can be used to bootstrap a consensus node
func (s *Service) GetInitialClusterString() string {
	members := s.etcd.Server.Cluster().Members()
	clusterString := strings.Builder{}

	for i, m := range members {
		if i != 0 {
			clusterString.WriteString(",")
		}
		clusterString.WriteString(m.Name)
		clusterString.WriteString("=")
		clusterString.WriteString(m.PickPeerURL())
	}

	return clusterString.String()
}

// GetNodes returns a list of consensus nodes
func (s *Service) GetNodes() []Member {
	members := s.etcd.Server.Cluster().Members()
	cMembers := make([]Member, len(members))
	for i, m := range members {
		cMembers[i] = Member{
			ID:      uint64(m.ID),
			Name:    m.Name,
			Address: m.PickPeerURL(),
			Synced:  !m.IsLearner,
		}
	}

	return cMembers
}

func (s *Service) GetStore(module, space string) clientv3.KV {
	return namespace.NewKV(s.kv, fmt.Sprintf("%s:%s", module, space))
}
