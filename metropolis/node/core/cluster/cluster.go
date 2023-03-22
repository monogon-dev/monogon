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

// Package cluster implements low-level clustering logic, especially logic
// regarding to bootstrapping, registering into and joining a cluster. Its goal
// is to provide the rest of the node code with the following:
//   - A mounted plaintext storage.
//   - Node credentials/identity.
//   - A locally running etcd server if the node is supposed to run one, and a
//     client connection to that etcd cluster if so.
//   - The state of the cluster as seen by the node, to enable code to respond to
//     node lifecycle changes.
package cluster

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/core/roleserve"
	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

type Manager struct {
	storageRoot    *localstorage.Root
	networkService *network.Service
	roleServer     *roleserve.Service
	nodeParams     *apb.NodeParameters
	haveTPM        bool

	oneway chan struct{}
}

// NewManager creates a new cluster Manager. The given localstorage Root must
// be places, but not yet started (and will be started as the Manager makes
// progress). The given network Service must already be running.
func NewManager(storageRoot *localstorage.Root, networkService *network.Service, rs *roleserve.Service, nodeParams *apb.NodeParameters, haveTPM bool) *Manager {
	return &Manager{
		storageRoot:    storageRoot,
		networkService: networkService,
		roleServer:     rs,
		nodeParams:     nodeParams,
		haveTPM:        haveTPM,
		oneway:         make(chan struct{}),
	}
}

// Run is the runnable of the Manager, to be started using the Supervisor. It
// is one-shot, and should not be restarted.
func (m *Manager) Run(ctx context.Context) error {
	select {
	case <-m.oneway:
		return fmt.Errorf("cannot restart cluster manager")
	default:
	}
	close(m.oneway)

	// Try sealed configuration first.
	configuration, err := m.storageRoot.ESP.Metropolis.SealedConfiguration.Unseal()
	if err == nil {
		supervisor.Logger(ctx).Info("Sealed configuration present. attempting to join cluster")

		// Read Cluster Directory and unmarshal it. Since the node is already
		// registered with the cluster, the directory won't be bootstrapped from
		// Node Parameters.
		cd, err := m.storageRoot.ESP.Metropolis.ClusterDirectory.Unmarshal()
		if err != nil {
			return fmt.Errorf("while reading cluster directory: %w", err)
		}
		return m.join(ctx, configuration, cd, true)
	}

	if !errors.Is(err, localstorage.ErrNoSealed) {
		return fmt.Errorf("unexpected sealed config error: %w", err)
	}

	configuration, err = m.storageRoot.ESP.Metropolis.SealedConfiguration.ReadUnsafe()
	if err == nil {
		supervisor.Logger(ctx).Info("Non-sealed configuration present. attempting to join cluster")

		// Read Cluster Directory and unmarshal it. Since the node is already
		// registered with the cluster, the directory won't be bootstrapped from
		// Node Parameters.
		cd, err := m.storageRoot.ESP.Metropolis.ClusterDirectory.Unmarshal()
		if err != nil {
			return fmt.Errorf("while reading cluster directory: %w", err)
		}
		return m.join(ctx, configuration, cd, false)
	}

	supervisor.Logger(ctx).Info("No sealed configuration, looking for node parameters")

	switch inner := m.nodeParams.Cluster.(type) {
	case *apb.NodeParameters_ClusterBootstrap_:
		err = m.bootstrap(ctx, inner.ClusterBootstrap)
	case *apb.NodeParameters_ClusterRegister_:
		err = m.register(ctx, inner.ClusterRegister)
	default:
		err = fmt.Errorf("node parameters misconfigured: neither cluster_bootstrap nor cluster_register set")
	}

	if err == nil {
		supervisor.Logger(ctx).Info("Cluster enrolment done.")
		return nil
	}
	return err
}

// logClusterDirectory verbosely logs the whole Cluster Directory passed to it.
func logClusterDirectory(ctx context.Context, cd *cpb.ClusterDirectory) {
	for _, node := range cd.Nodes {
		var addresses []string
		for _, add := range node.Addresses {
			addresses = append(addresses, add.Host)
		}
		supervisor.Logger(ctx).Infof("    Addresses: %s", strings.Join(addresses, ","))
	}
}
