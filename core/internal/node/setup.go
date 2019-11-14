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
	"git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"git.monogon.dev/source/nexantic.git/core/internal/storage"

	"errors"

	"go.uber.org/zap"
)

var (
	ErrConsensusAlreadyProvisioned = errors.New("consensus is already provisioned; make sure the data folder is empty")
	ErrAlreadySetup                = errors.New("node is already set up")
	ErrNotInJoinMode               = errors.New("node is not in the cluster join mode")
	ErrTrustNotInitialized         = errors.New("trust backend not initialized")
	ErrStorageNotInitialized       = errors.New("storage not initialized")
)

func (s *SmalltownNode) CurrentState() common.SmalltownState {
	return s.state
}

func (s *SmalltownNode) GetJoinClusterToken() string {
	return s.joinToken
}

func (s *SmalltownNode) SetupNewCluster() error {
	if s.state == common.StateConfigured {
		return ErrAlreadySetup
	}
	dataPath, err := s.Storage.GetPathInPlace(storage.PlaceData, "etcd")
	if err == storage.ErrNotInitialized {
		return ErrStorageNotInitialized
	} else if err != nil {
		return err
	}

	s.logger.Info("Setting up a new cluster")
	s.logger.Info("Provisioning consensus")

	// Make sure etcd is not yet provisioned
	if s.Consensus.IsProvisioned() {
		return ErrConsensusAlreadyProvisioned
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

	// Change system state
	s.state = common.StateConfigured

	s.logger.Info("New Cluster set up. Node is now fully operational")

	return nil
}

func (s *SmalltownNode) EnterJoinClusterMode() error {
	if s.state == common.StateConfigured {
		return ErrAlreadySetup
	}
	s.state = common.StateClusterJoinMode

	s.logger.Info("Node is now in the cluster join mode")

	return nil
}

func (s *SmalltownNode) JoinCluster(clusterString string, certs *api.ConsensusCertificates) error {
	if s.state != common.StateClusterJoinMode {
		return ErrNotInJoinMode
	}

	s.logger.Info("Joining cluster", zap.String("cluster", clusterString))

	err := s.SetupBackend()
	if err != nil {
		return err
	}

	config := s.Consensus.GetConfig()
	config.Name = s.hostname
	config.InitialCluster = clusterString
	s.Consensus.SetConfig(config)
	if err := s.Consensus.WriteCertificateFiles(certs); err != nil {
		return err
	}

	// Start consensus
	err = s.Consensus.Start()
	if err != nil {
		return err
	}

	s.state = common.StateConfigured

	s.logger.Info("Joined cluster. Node is now syncing.")

	return nil
}
