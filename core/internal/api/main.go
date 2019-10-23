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
	"fmt"
	schema "git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus"
	"github.com/casbin/casbin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type (
	Server struct {
		*common.BaseService

		ruleEnforcer *casbin.Enforcer
		setupService common.SetupService
		grpcServer   *grpc.Server

		consensusService *consensus.Service

		config *Config
	}

	Config struct {
		Port uint16
	}
)

func NewApiServer(config *Config, logger *zap.Logger, setupService common.SetupService, consensusService *consensus.Service) (*Server, error) {
	s := &Server{
		config:           config,
		setupService:     setupService,
		consensusService: consensusService,
	}

	s.BaseService = common.NewBaseService("api", logger, s)

	grpcServer := grpc.NewServer()
	schema.RegisterClusterManagementServer(grpcServer, s)
	schema.RegisterSetupServiceServer(grpcServer, s)

	s.grpcServer = grpcServer

	return s, nil
}

func (s *Server) OnStart() error {
	listenHost := fmt.Sprintf(":%d", s.config.Port)
	lis, err := net.Listen("tcp", listenHost)
	if err != nil {
		s.Logger.Fatal("failed to listen", zap.Error(err))
	}

	go func() {
		err = s.grpcServer.Serve(lis)
		s.Logger.Error("API server failed", zap.Error(err))
	}()

	s.Logger.Info("GRPC listening", zap.String("host", listenHost))

	return nil
}

func (s *Server) OnStop() error {
	s.grpcServer.Stop()

	return nil
}
