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
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/core/network/dns"
	"source.monogon.dev/metropolis/node/kubernetes/authproxy"
	"source.monogon.dev/metropolis/node/kubernetes/metricsproxy"
	"source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/node/kubernetes/reconciler"
	"source.monogon.dev/metropolis/pkg/supervisor"

	apb "source.monogon.dev/metropolis/proto/api"
)

type ConfigController struct {
	ServiceIPRange net.IPNet
	ClusterNet     net.IPNet
	ClusterDomain  string

	KPKI    *pki.PKI
	Root    *localstorage.Root
	Network *network.Service
	Node    *identity.NodeCredentials
}

type Controller struct {
	c ConfigController
}

func NewController(c ConfigController) *Controller {
	s := &Controller{
		c: c,
	}
	return s
}

func (s *Controller) Run(ctx context.Context) error {
	controllerManagerConfig, err := getPKIControllerManagerConfig(ctx, s.c.KPKI)
	if err != nil {
		return fmt.Errorf("could not generate controller manager pki config: %w", err)
	}
	controllerManagerConfig.clusterNet = s.c.ClusterNet
	schedulerConfig, err := getPKISchedulerConfig(ctx, s.c.KPKI)
	if err != nil {
		return fmt.Errorf("could not generate scheduler pki config: %w", err)
	}

	masterKubeconfig, err := s.c.KPKI.Kubeconfig(ctx, pki.Master, pki.KubernetesAPIEndpointForController)
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

	// Sub-runnable which starts all parts of Kubernetes that depend on the
	// machine's external IP address. If it changes, the runnable will exit.
	// TODO(q3k): test this
	supervisor.Run(ctx, "networked", func(ctx context.Context) error {
		networkWatch := s.c.Network.Watch()
		defer networkWatch.Close()

		var status *network.Status

		supervisor.Logger(ctx).Info("Waiting for node networking...")
		for status == nil || status.ExternalAddress == nil {
			status, err = networkWatch.Get(ctx)
			if err != nil {
				return fmt.Errorf("failed to get network status: %w", err)
			}
		}
		address := status.ExternalAddress
		supervisor.Logger(ctx).Info("Node has active networking, starting apiserver/kubelet")

		apiserver := &apiserverService{
			KPKI:                        s.c.KPKI,
			AdvertiseAddress:            address,
			ServiceIPRange:              s.c.ServiceIPRange,
			EphemeralConsensusDirectory: &s.c.Root.Ephemeral.Consensus,
		}

		err := supervisor.RunGroup(ctx, map[string]supervisor.Runnable{
			"apiserver": apiserver.Run,
		})
		if err != nil {
			return fmt.Errorf("when starting apiserver/kubelet: %w", err)
		}

		supervisor.Signal(ctx, supervisor.SignalHealthy)

		for status.ExternalAddress.Equal(address) {
			status, err = networkWatch.Get(ctx)
			if err != nil {
				return fmt.Errorf("when watching for network changes: %w", err)
			}
		}
		return fmt.Errorf("network configuration changed (%s -> %s)", address.String(), status.ExternalAddress.String())
	})

	// Before we start anything else, make sure reconciliation passes at least once.
	// This makes the initial startup of a cluster much cleaner as we don't end up
	// starting the scheduler/controller-manager/etc just to get them to immediately
	// fail and back off with 'unauthorized'.
	startLogging := time.Now().Add(2 * time.Second)
	supervisor.Logger(ctx).Infof("Performing initial resource reconciliation...")
	for {
		err := reconciler.ReconcileAll(ctx, clientSet)
		if err == nil {
			supervisor.Logger(ctx).Infof("Initial resource reconciliation succeeded.")
			break
		}
		if time.Now().After(startLogging) {
			supervisor.Logger(ctx).Errorf("Still couldn't do initial reconciliation: %v", err)
			startLogging = time.Now().Add(10 * time.Second)
		}
		time.Sleep(100 * time.Millisecond)
	}

	authProxy := authproxy.Service{
		KPKI: s.c.KPKI,
		Node: s.c.Node,
	}

	metricsProxy := metricsproxy.Service{
		KPKI: s.c.KPKI,
	}

	for _, sub := range []struct {
		name     string
		runnable supervisor.Runnable
	}{
		{"controller-manager", runControllerManager(*controllerManagerConfig)},
		{"scheduler", runScheduler(*schedulerConfig)},
		{"reconciler", reconciler.Maintain(clientSet)},
		{"authproxy", authProxy.Run},
		{"metricsproxy", metricsProxy.Run},
	} {
		err := supervisor.Run(ctx, sub.name, sub.runnable)
		if err != nil {
			return fmt.Errorf("could not run sub-service %q: %w", sub.name, err)
		}
	}

	supervisor.Logger(ctx).Info("Registering K8s CoreDNS")
	clusterDNSDirective := dns.NewKubernetesDirective(s.c.ClusterDomain, masterKubeconfig)
	s.c.Network.ConfigureDNS(clusterDNSDirective)

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	s.c.Network.ConfigureDNS(dns.CancelDirective(clusterDNSDirective))
	return nil
}

// GetDebugKubeconfig issues a kubeconfig for an arbitrary given identity.
// Useful for debugging and testing.
func (s *Controller) GetDebugKubeconfig(ctx context.Context, request *apb.GetDebugKubeconfigRequest) (*apb.GetDebugKubeconfigResponse, error) {
	client, err := s.c.KPKI.VolatileClient(ctx, request.Id, request.Groups)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Failed to get volatile client certificate: %v", err)
	}
	kubeconfig, err := pki.Kubeconfig(ctx, s.c.KPKI.KV, client, pki.KubernetesAPIEndpointForController)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Failed to generate kubeconfig: %v", err)
	}
	return &apb.GetDebugKubeconfigResponse{DebugKubeconfig: string(kubeconfig)}, nil
}
