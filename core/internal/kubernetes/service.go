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
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	schema "git.monogon.dev/source/nexantic.git/core/generated/api"
	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/clusternet"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/pki"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/reconciler"
	"git.monogon.dev/source/nexantic.git/core/internal/storage"
	"git.monogon.dev/source/nexantic.git/core/pkg/logbuffer"
)

type Config struct {
	AdvertiseAddress net.IP
	ServiceIPRange   net.IPNet
	ClusterNet       net.IPNet
}

type Service struct {
	consensusService *consensus.Service
	storageService   *storage.Manager
	logger           *zap.Logger

	apiserverLogs         *logbuffer.LogBuffer
	controllerManagerLogs *logbuffer.LogBuffer
	schedulerLogs         *logbuffer.LogBuffer
	kubeletLogs           *logbuffer.LogBuffer

	kpki *pki.KubernetesPKI
}

func New(logger *zap.Logger, consensusService *consensus.Service, storageService *storage.Manager) *Service {
	s := &Service{
		consensusService: consensusService,
		storageService:   storageService,
		logger:           logger,

		apiserverLogs:         logbuffer.New(5000, 16384),
		controllerManagerLogs: logbuffer.New(5000, 16384),
		schedulerLogs:         logbuffer.New(5000, 16384),
		kubeletLogs:           logbuffer.New(5000, 16384),

		kpki: pki.NewKubernetes(logger.Named("pki")),
	}

	return s
}

func (s *Service) getKV() clientv3.KV {
	return s.consensusService.GetStore("kubernetes", "")
}

func (s *Service) NewCluster() error {
	// TODO(q3k): this needs to be passed by the caller.
	ctx := context.TODO()
	return s.kpki.EnsureAll(ctx, s.getKV())
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
	if !s.consensusService.IsReady() {
		return nil, status.Error(codes.Unavailable, "Consensus not ready yet")
	}
	ca := s.kpki.Certificates[pki.IdCA]
	debugKubeconfig, err := pki.New(ca, "", pki.Client(request.Id, request.Groups)).Kubeconfig(ctx, s.getKV())
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Failed to generate kubeconfig: %v", err)
	}
	return &schema.GetDebugKubeconfigResponse{DebugKubeconfig: string(debugKubeconfig)}, nil
}

func (s *Service) Start() error {
	// TODO(lorenz): This creates a new supervision tree, it should instead attach to the root one. But for that SmalltownNode needs
	// to be ported to supervisor.
	supervisor.New(context.TODO(), s.logger, s.Run())
	return nil
}

func (s *Service) Run() supervisor.Runnable {
	return func(ctx context.Context) error {
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
		apiserverConfig, err := getPKIApiserverConfig(ctx, consensusKV, s.kpki)
		if err != nil {
			return fmt.Errorf("could not generate apiserver pki config: %w", err)
		}
		apiserverConfig.advertiseAddress = config.AdvertiseAddress
		apiserverConfig.serviceIPRange = config.ServiceIPRange
		controllerManagerConfig, err := getPKIControllerManagerConfig(ctx, consensusKV, s.kpki)
		if err != nil {
			return fmt.Errorf("could not generate controller manager pki config: %w", err)
		}
		controllerManagerConfig.clusterNet = config.ClusterNet
		schedulerConfig, err := getPKISchedulerConfig(ctx, consensusKV, s.kpki)
		if err != nil {
			return fmt.Errorf("could not generate scheduler pki config: %w", err)
		}

		masterKubeconfig, err := s.kpki.Kubeconfig(ctx, pki.Master, consensusKV)
		if err != nil {
			return fmt.Errorf("could not generate master kubeconfig: %w", err)
		}

		rawClientConfig, err := clientcmd.NewClientConfigFromBytes(masterKubeconfig)
		if err != nil {
			return fmt.Errorf("could not generate kubernetes client config: %w", err)
		}

		clientConfig, err := rawClientConfig.ClientConfig()
		clientSet, err := kubernetes.NewForConfig(clientConfig)
		if err != nil {
			return fmt.Errorf("could not generate kubernetes client: %w", err)
		}

		informerFactory := informers.NewSharedInformerFactory(clientSet, 5*time.Minute)

		hostname, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("failed to get hostname: %w", err)
		}
		err = createKubeletConfig(ctx, consensusKV, s.kpki, hostname)
		if err != nil {
			return fmt.Errorf("could not created kubelet config: %w", err)
		}

		key, err := clusternet.EnsureOnDiskKey()
		if err != nil {
			return fmt.Errorf("failed to ensure cluster key: %w", err)
		}

		for _, sub := range []struct {
			name     string
			runnable supervisor.Runnable
		}{
			{"apiserver", runAPIServer(*apiserverConfig, s.apiserverLogs)},
			{"controller-manager", runControllerManager(*controllerManagerConfig, s.controllerManagerLogs)},
			{"scheduler", runScheduler(*schedulerConfig, s.schedulerLogs)},
			{"kubelet", runKubelet(&KubeletSpec{}, s.kubeletLogs)},
			{"reconciler", reconciler.Run(clientSet)},
			{"csi-plugin", runCSIPlugin(s.storageService)},
			{"pv-provisioner", runCSIProvisioner(s.storageService, clientSet, informerFactory)},
			{"clusternet", clusternet.Run(informerFactory, clusterNet, clientSet, key)},
		} {
			err := supervisor.Run(ctx, sub.name, sub.runnable)
			if err != nil {
				return fmt.Errorf("could not run sub-service %q: %w", sub.name, err)
			}
		}

		supervisor.Signal(ctx, supervisor.SignalHealthy)
		supervisor.Signal(ctx, supervisor.SignalDone)
		return nil
	}
}
