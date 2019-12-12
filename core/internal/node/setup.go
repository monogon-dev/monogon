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
	"context"

	"git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"errors"

	"go.uber.org/zap"
)

var (
	ErrConsensusAlreadyProvisioned = status.Error(codes.FailedPrecondition, "consensus is already provisioned; make sure the data folder is empty")
	ErrAlreadySetup                = status.Error(codes.FailedPrecondition, "node is already set up")
	ErrNotInJoinMode               = status.Error(codes.FailedPrecondition, "node is not in the cluster join mode")
	ErrTrustNotInitialized         = status.Error(codes.FailedPrecondition, "trust backend not initialized")
	ErrStorageNotInitialized       = status.Error(codes.FailedPrecondition, "storage not initialized")
)

func (s *SmalltownNode) CurrentState() common.SmalltownState {
	return s.state
}

// InitializeNode contains functionality that needs to be executed regardless of what the node does
// later on
func (s *SmalltownNode) InitializeNode() (*api.NewNodeInfo, string, error) {
	globalUnlockKey, err := s.Storage.InitializeData()
	if err != nil {
		return nil, "", err
	}

	nodeIP := s.Network.GetIP(true)

	nodeCert, nodeID, err := s.generateNodeID()
	if err != nil {
		return nil, "", err
	}

	return &api.NewNodeInfo{
		EnrolmentConfig: s.enrolmentConfig,
		Ip:              []byte(*nodeIP),
		IdCert:          nodeCert,
		GlobalUnlockKey: globalUnlockKey,
	}, nodeID, nil
}

func (s *SmalltownNode) JoinCluster(context context.Context, req *api.JoinClusterRequest) (*api.JoinClusterResponse, error) {
	if s.state != common.StateEnrollMode {
		return nil, ErrNotInJoinMode
	}

	s.logger.Info("Joining Consenus")

	config := s.Consensus.GetConfig()
	config.Name = s.hostname
	config.InitialCluster = "default" // Clusters can't cross-join anyways due to cryptography
	s.Consensus.SetConfig(config)
	var err error
	if err != nil {
		s.logger.Warn("Invalid JoinCluster request", zap.Error(err))
		return nil, errors.New("invalid join request")
	}
	if err := s.Consensus.WriteCertificateFiles(req.Certs); err != nil {
		return nil, err
	}

	// Start consensus
	err = s.Consensus.Start()
	if err != nil {
		return nil, err
	}

	s.state = common.StateJoined

	s.logger.Info("Joined cluster. Node is now syncing.")

	return &api.JoinClusterResponse{}, nil
}
