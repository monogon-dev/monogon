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
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/clusternet"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/nfproxy"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/pki"
	"git.monogon.dev/source/nexantic.git/core/internal/kubernetes/reconciler"
	"git.monogon.dev/source/nexantic.git/core/internal/localstorage"
	"git.monogon.dev/source/nexantic.git/core/internal/network/dns"
	"git.monogon.dev/source/nexantic.git/core/pkg/logbuffer"
	apb "git.monogon.dev/source/nexantic.git/core/proto/api"
)

type Config struct {
	AdvertiseAddress net.IP
	ServiceIPRange   net.IPNet
	ClusterNet       net.IPNet

	KPKI                    *pki.KubernetesPKI
	Root                    *localstorage.Root
	CorednsRegistrationChan chan *dns.ExtraDirective
}

type state struct {
	apiserverLogs         *logbuffer.LogBuffer
	controllerManagerLogs *logbuffer.LogBuffer
	schedulerLogs         *logbuffer.LogBuffer
	kubeletLogs           *logbuffer.LogBuffer
}

type Service struct {
	c       Config
	stateMu sync.Mutex
	state   *state
}

func New(c Config) *Service {
	s := &Service{
		c: c,
	}
	return s
}

func (s *Service) getState() *state {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	return s.state
}

func (s *Service) Run(ctx context.Context) error {
	st := &state{
		apiserverLogs:         logbuffer.New(5000, 16384),
		controllerManagerLogs: logbuffer.New(5000, 16384),
		schedulerLogs:         logbuffer.New(5000, 16384),
		kubeletLogs:           logbuffer.New(5000, 16384),
	}
	s.stateMu.Lock()
	s.state = st
	s.stateMu.Unlock()

	controllerManagerConfig, err := getPKIControllerManagerConfig(ctx, s.c.KPKI)
	if err != nil {
		return fmt.Errorf("could not generate controller manager pki config: %w", err)
	}
	controllerManagerConfig.clusterNet = s.c.ClusterNet
	schedulerConfig, err := getPKISchedulerConfig(ctx, s.c.KPKI)
	if err != nil {
		return fmt.Errorf("could not generate scheduler pki config: %w", err)
	}

	masterKubeconfig, err := s.c.KPKI.Kubeconfig(ctx, pki.Master)
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

	dnsHostIP := s.c.AdvertiseAddress // TODO: Which IP to use

	apiserver := &apiserverService{
		KPKI:                        s.c.KPKI,
		AdvertiseAddress:            s.c.AdvertiseAddress,
		ServiceIPRange:              s.c.ServiceIPRange,
		Output:                      st.apiserverLogs,
		EphemeralConsensusDirectory: &s.c.Root.Ephemeral.Consensus,
	}

	kubelet := kubeletService{
		NodeName:           hostname,
		ClusterDNS:         []net.IP{dnsHostIP},
		KubeletDirectory:   &s.c.Root.Data.Kubernetes.Kubelet,
		EphemeralDirectory: &s.c.Root.Ephemeral,
		Output:             st.kubeletLogs,
		KPKI:               s.c.KPKI,
	}

	csiPlugin := csiPluginServer{
		KubeletDirectory: &s.c.Root.Data.Kubernetes.Kubelet,
		VolumesDirectory: &s.c.Root.Data.Volumes,
	}

	csiProvisioner := csiProvisionerServer{
		NodeName:         hostname,
		Kubernetes:       clientSet,
		InformerFactory:  informerFactory,
		VolumesDirectory: &s.c.Root.Data.Volumes,
	}

	clusternet := clusternet.Service{
		NodeName:        hostname,
		Kubernetes:      clientSet,
		ClusterNet:      s.c.ClusterNet,
		InformerFactory: informerFactory,
		DataDirectory:   &s.c.Root.Data.Kubernetes.ClusterNetworking,
	}

	nfproxy := nfproxy.Service{
		ClusterCIDR: s.c.ClusterNet,
		ClientSet:   clientSet,
	}

	for _, sub := range []struct {
		name     string
		runnable supervisor.Runnable
	}{
		{"apiserver", apiserver.Run},
		{"controller-manager", runControllerManager(*controllerManagerConfig, st.controllerManagerLogs)},
		{"scheduler", runScheduler(*schedulerConfig, st.schedulerLogs)},
		{"kubelet", kubelet.Run},
		{"reconciler", reconciler.Run(clientSet)},
		{"csi-plugin", csiPlugin.Run},
		{"csi-provisioner", csiProvisioner.Run},
		{"clusternet", clusternet.Run},
		{"nfproxy", nfproxy.Run},
	} {
		err := supervisor.Run(ctx, sub.name, sub.runnable)
		if err != nil {
			return fmt.Errorf("could not run sub-service %q: %w", sub.name, err)
		}
	}

	supervisor.Logger(ctx).Info("Registering K8s CoreDNS")
	clusterDNSDirective := dns.NewKubernetesDirective("cluster.local", masterKubeconfig)
	s.c.CorednsRegistrationChan <- clusterDNSDirective

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	s.c.CorednsRegistrationChan <- dns.CancelDirective(clusterDNSDirective)
	return nil
}

// GetComponentLogs grabs logs from various Kubernetes binaries
func (s *Service) GetComponentLogs(component string, n int) ([]string, error) {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	if s.state == nil {
		return nil, errors.New("kubernetes not started yet")
	}

	switch component {
	case "apiserver":
		return s.state.apiserverLogs.ReadLinesTruncated(n, "..."), nil
	case "controller-manager":
		return s.state.controllerManagerLogs.ReadLinesTruncated(n, "..."), nil
	case "scheduler":
		return s.state.schedulerLogs.ReadLinesTruncated(n, "..."), nil
	case "kubelet":
		return s.state.kubeletLogs.ReadLinesTruncated(n, "..."), nil
	default:
		return nil, errors.New("component not available")
	}
}

// GetDebugKubeconfig issues a kubeconfig for an arbitrary given identity. Useful for debugging and testing.
func (s *Service) GetDebugKubeconfig(ctx context.Context, request *apb.GetDebugKubeconfigRequest) (*apb.GetDebugKubeconfigResponse, error) {
	ca := s.c.KPKI.Certificates[pki.IdCA]
	debugKubeconfig, err := pki.New(ca, "", pki.Client(request.Id, request.Groups)).Kubeconfig(ctx, s.c.KPKI.KV)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Failed to generate kubeconfig: %v", err)
	}
	return &apb.GetDebugKubeconfigResponse{DebugKubeconfig: string(debugKubeconfig)}, nil
}
