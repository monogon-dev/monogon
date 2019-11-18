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
	"flag"

	"git.monogon.dev/source/nexantic.git/core/internal/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes"
	"git.monogon.dev/source/nexantic.git/core/internal/storage"
	"os"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type (
	SmalltownNode struct {
		Api        *api.Server
		Consensus  *consensus.Service
		Storage    *storage.Manager
		Kubernetes *kubernetes.Service

		logger    *zap.Logger
		state     common.SmalltownState
		joinToken string
		hostname  string
	}
)

func NewSmalltownNode(logger *zap.Logger, apiPort, consensusPort uint16) (*SmalltownNode, error) {
	flag.Parse()
	logger.Info("Creating Smalltown node")

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	storageManager, err := storage.Initialize(logger.With(zap.String("component", "storage")))
	if err != nil {
		logger.Error("Failed to initialize storage manager", zap.Error(err))
		return nil, err
	}

	consensusService, err := consensus.NewConsensusService(consensus.Config{
		Name:         hostname,
		ListenPort:   consensusPort,
		ListenHost:   "0.0.0.0",
		ExternalHost: "10.0.2.15", // TODO: Once Multi-Node setups are actually used, this needs to be corrected
	}, logger.With(zap.String("module", "consensus")))
	if err != nil {
		return nil, err
	}

	s := &SmalltownNode{
		Consensus: consensusService,
		Storage:   storageManager,
		logger:    logger,
		hostname:  hostname,
	}

	apiService, err := api.NewApiServer(&api.Config{
		Port: apiPort,
	}, logger.With(zap.String("module", "api")), s, s.Consensus)
	if err != nil {
		return nil, err
	}

	s.Api = apiService

	s.Kubernetes = kubernetes.New(logger.With(zap.String("module", "kubernetes")), consensusService)

	logger.Info("Created SmalltownNode")

	return s, nil
}

func (s *SmalltownNode) Start() error {
	s.logger.Info("Starting Smalltown node")

	if s.Consensus.IsProvisioned() {
		s.logger.Info("Consensus is provisioned")
		err := s.startFull()
		if err != nil {
			return err
		}
	} else {
		s.logger.Info("Consensus is not provisioned, starting provisioning...")
		err := s.startForSetup()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SmalltownNode) startForSetup() error {
	s.logger.Info("Initializing subsystems for setup mode")
	s.state = common.StateSetupMode
	s.joinToken = uuid.New().String()

	err := s.Api.Start()
	if err != nil {
		s.logger.Error("Failed to start the API service", zap.Error(err))
		return err
	}

	return nil
}

func (s *SmalltownNode) startFull() error {
	s.logger.Info("Initializing subsystems for production")
	s.state = common.StateConfigured

	err := s.SetupBackend()
	if err != nil {
		return err
	}

	err = s.Consensus.Start()
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

func (s *SmalltownNode) SetupBackend() error {
	s.logger.Debug("Creating trust backend")

	return nil
}
