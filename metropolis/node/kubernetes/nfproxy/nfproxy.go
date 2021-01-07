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

// Package nfproxy is a Kubernetes Service IP proxy based exclusively on the Linux nftables interface.
// It uses netfilter's NAT capabilities to accept traffic on service IPs and DNAT it to the respective endpoint.
package nfproxy

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sbezverk/nfproxy/pkg/controller"
	"github.com/sbezverk/nfproxy/pkg/nftables"
	"github.com/sbezverk/nfproxy/pkg/proxy"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"

	"source.monogon.dev/metropolis/pkg/supervisor"
)

type Service struct {
	// Traffic in ClusterCIDR is assumed to be originated inside the cluster and will not be SNATed
	ClusterCIDR net.IPNet
	// A Kubernetes ClientSet with read access to endpoints and services
	ClientSet kubernetes.Interface
}

func (s *Service) Run(ctx context.Context) error {
	var ipv4ClusterCIDR string
	var ipv6ClusterCIDR string
	if s.ClusterCIDR.IP.To4() == nil && s.ClusterCIDR.IP.To16() != nil {
		ipv6ClusterCIDR = s.ClusterCIDR.String()
	} else if s.ClusterCIDR.IP.To4() != nil {
		ipv4ClusterCIDR = s.ClusterCIDR.String()
	} else {
		return errors.New("invalid ClusterCIDR")
	}
	nfti, err := nftables.InitNFTables(ipv4ClusterCIDR, ipv6ClusterCIDR)
	if err != nil {
		return fmt.Errorf("failed to initialize nftables with error: %w", err)
	}

	// Create event recorder to report events into K8s
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get local host name with error: %w", err)
	}
	eventBroadcaster := record.NewBroadcaster()
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "nfproxy", Host: hostname})

	// Create new proxy controller with endpoint slices enabled
	// https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/
	nfproxy := proxy.NewProxy(nfti, hostname, recorder, true)

	// Create special informer which doesn't track headless services
	noHeadlessEndpoints, err := labels.NewRequirement(v1.IsHeadlessService, selection.DoesNotExist, nil)
	if err != nil {
		return fmt.Errorf("failed to create Requirement for noHeadlessEndpoints: %w", err)
	}
	labelSelector := labels.NewSelector()
	labelSelector = labelSelector.Add(*noHeadlessEndpoints)

	kubeInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(s.ClientSet, time.Minute*5,
		kubeinformers.WithTweakListOptions(func(options *metav1.ListOptions) {
			options.LabelSelector = labelSelector.String()
		}))

	svcController := controller.NewServiceController(nfproxy, s.ClientSet, kubeInformerFactory.Core().V1().Services())
	ep := controller.NewEndpointSliceController(nfproxy, s.ClientSet, kubeInformerFactory.Discovery().V1beta1().EndpointSlices())
	kubeInformerFactory.Start(ctx.Done())

	if err = svcController.Start(ctx.Done()); err != nil {
		return fmt.Errorf("error running Service controller: %w", err)
	}
	if err = ep.Start(ctx.Done()); err != nil {
		return fmt.Errorf("error running endpoint controller: %w", err)
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
