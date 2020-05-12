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
	"context"
	"crypto/ed25519"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	schema "git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/pkg/logbuffer"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"

	"git.monogon.dev/source/nexantic.git/core/internal/common/service"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus"
)

type Config struct {
	AdvertiseAddress net.IP
	ServiceIPRange   net.IPNet
	ClusterNet       net.IPNet
}

type Service struct {
	*service.BaseService
	consensusService      *consensus.Service
	logger                *zap.Logger
	apiserverLogs         *logbuffer.LogBuffer
	controllerManagerLogs *logbuffer.LogBuffer
	schedulerLogs         *logbuffer.LogBuffer
	kubeletLogs           *logbuffer.LogBuffer
}

func New(logger *zap.Logger, consensusService *consensus.Service) *Service {
	s := &Service{
		consensusService:      consensusService,
		logger:                logger,
		apiserverLogs:         logbuffer.New(5000, 16384),
		controllerManagerLogs: logbuffer.New(5000, 16384),
		schedulerLogs:         logbuffer.New(5000, 16384),
		kubeletLogs:           logbuffer.New(5000, 16384),
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

// GetComponentLogs grabs logs from various Kubernetes binaries
func (s *Service) GetComponentLogs(component string, n int) ([]string, error) {
	switch component {
	case "apiserver":
		return s.apiserverLogs.ReadLinesTruncated(n, "..."), nil
	case "controller-manager":
		return s.controllerManagerLogs.ReadLinesTruncated(n, "..."), nil
	case "scheduler":
		return s.schedulerLogs.ReadLinesTruncated(n, "..."), nil
	case "kubelet":
		return s.kubeletLogs.ReadLinesTruncated(n, "..."), nil
	default:
		return []string{}, errors.New("component not available")
	}
}

// GetDebugKubeconfig issues a kubeconfig for an arbitrary given identity. Useful for debugging and testing.
func (s *Service) GetDebugKubeconfig(ctx context.Context, request *schema.GetDebugKubeconfigRequest) (*schema.GetDebugKubeconfigResponse, error) {
	idCA, idKeyRaw, err := getCert(s.getKV(), "id-ca")
	idKey := ed25519.PrivateKey(idKeyRaw)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Failed to load ID CA: %v", err)
	}
	debugCert, debugKey, err := issueCertificate(clientCertTemplate(request.Id, request.Groups), idCA, idKey)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Failed to issue certs for kubeconfig: %v\n", err)
	}
	debugKubeconfig, err := makeLocalKubeconfig(idCA, debugCert, debugKey)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Failed to generate kubeconfig: %v", err)
	}
	return &schema.GetDebugKubeconfigResponse{DebugKubeconfig: string(debugKubeconfig)}, nil
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

	masterKubeconfig, err := getSingle(consensusKV, "master.kubeconfig")
	if err != nil {
		return err
	}

	// TODO(lorenz): Once internal/node is part of the supervisor tree, these should all be supervisor runnables
	go func() {
		s.runAPIServer(context.TODO(), *apiserverConfig)
	}()
	go func() {
		s.runControllerManager(context.TODO(), *controllerManagerConfig)
	}()
	go func() {
		s.runScheduler(context.TODO(), *schedulerConfig)
	}()

	go func() {
		if err := s.runKubelet(context.TODO(), &KubeletSpec{}); err != nil {
			fmt.Printf("Failed to launch kubelet: %v\n", err)
		}
	}()

	go func() {
		go runReconciler(context.TODO(), masterKubeconfig, s.logger)
	}()

	return nil
}

func (s *Service) OnStop() error {
	// Requires advanced process management and not necessary for MVP
	return errors.New("Not implemented")
}
