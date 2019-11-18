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

package kubernetes

import (
	"errors"
	"net"

	"git.monogon.dev/source/nexantic.git/core/internal/common/service"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

type Config struct {
	AdvertiseAddress net.IP
	ServiceIPRange   net.IPNet
	ClusterNet       net.IPNet
}

type Service struct {
	*service.BaseService
	consensusService *consensus.Service
	logger           *zap.Logger
}

func New(logger *zap.Logger, consensusService *consensus.Service) *Service {
	s := &Service{
		consensusService: consensusService,
		logger:           logger,
	}
	s.BaseService = service.NewBaseService("kubernetes", logger, s)
	return s
}

func (s *Service) getKV() clientv3.KV {
	return s.consensusService.GetStore("kubernetes", "")
}

func (s *Service) NewCluster() error {
	return newCluster(s.getKV())
}

func (s *Service) OnStart() error {
	config := Config{
		AdvertiseAddress: net.IP{10, 0, 2, 15}, // Depends on networking
		ServiceIPRange: net.IPNet{ // TODO: Decide if configurable / final value
			IP:   net.IP{192, 168, 188, 0},
			Mask: net.IPMask{0xff, 0xff, 0xff, 0x00}, // /24, but Go stores as a literal mask
		},
		ClusterNet: net.IPNet{
			IP:   net.IP{192, 168, 188, 0},
			Mask: net.IPMask{0xff, 0xff, 0xfd, 0x00}, // /22
		},
	}
	consensusKV := s.getKV()
	apiserverConfig, err := getPKIApiserverConfig(consensusKV)
	if err != nil {
		return err
	}
	apiserverConfig.advertiseAddress = config.AdvertiseAddress
	apiserverConfig.serviceIPRange = config.ServiceIPRange
	controllerManagerConfig, err := getPKIControllerManagerConfig(consensusKV)
	if err != nil {
		return err
	}
	controllerManagerConfig.clusterNet = config.ClusterNet
	schedulerConfig, err := getPKISchedulerConfig(consensusKV)
	if err != nil {
		return err
	}

	go func() {
		runAPIServer(*apiserverConfig)
	}()
	go func() {
		runControllerManager(*controllerManagerConfig)
	}()
	go func() {
		runScheduler(*schedulerConfig)
	}()

	return nil
}

func (s *Service) OnStop() error {
	// Requires advanced process management and not necessary for MVP
	return errors.New("Not implemented")
}
