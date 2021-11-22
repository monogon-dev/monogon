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

// cluster implements low-level clustering logic, especially logic regarding to
// bootstrapping, registering into and joining a cluster. Its goal is to provide
// the rest of the node code with the following:
//  - A mounted plaintext storage.
//  - Node credentials/identity.
//  - A locally running etcd server if the node is supposed to run one, and a
//    client connection to that etcd cluster if so.
//  - The state of the cluster as seen by the node, to enable code to respond to
//    node lifecycle changes.
package cluster

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
	ppb "source.monogon.dev/metropolis/proto/private"
)

type state struct {
	mu sync.RWMutex

	oneway bool

	configuration *ppb.SealedConfiguration
}

type Manager struct {
	storageRoot    *localstorage.Root
	networkService *network.Service
	status         memory.Value

	state

	// consensus is the spawned etcd/consensus service, if the Manager brought
	// up a Node that should run one.
	consensus *consensus.Service
}

// NewManager creates a new cluster Manager. The given localstorage Root must
// be places, but not yet started (and will be started as the Manager makes
// progress). The given network Service must already be running.
func NewManager(storageRoot *localstorage.Root, networkService *network.Service) *Manager {
	return &Manager{
		storageRoot:    storageRoot,
		networkService: networkService,

		state: state{},
	}
}

func (m *Manager) lock() (*state, func()) {
	m.mu.Lock()
	return &m.state, m.mu.Unlock
}

func (m *Manager) rlock() (*state, func()) {
	m.mu.RLock()
	return &m.state, m.mu.RUnlock
}

// Run is the runnable of the Manager, to be started using the Supervisor. It
// is one-shot, and should not be restarted.
func (m *Manager) Run(ctx context.Context) error {
	state, unlock := m.lock()
	if state.oneway {
		unlock()
		// TODO(q3k): restart the entire system if this happens
		return fmt.Errorf("cannot restart cluster manager")
	}
	state.oneway = true
	unlock()

	configuration, err := m.storageRoot.ESP.SealedConfiguration.Unseal()
	if err == nil {
		supervisor.Logger(ctx).Info("Sealed configuration present. attempting to join cluster")
		return m.join(ctx, configuration)
	}

	if !errors.Is(err, localstorage.ErrNoSealed) {
		return fmt.Errorf("unexpected sealed config error: %w", err)
	}

	supervisor.Logger(ctx).Info("No sealed configuration, looking for node parameters")

	params, err := m.nodeParams(ctx)
	if err != nil {
		return fmt.Errorf("no parameters available: %w", err)
	}

	switch inner := params.Cluster.(type) {
	case *apb.NodeParameters_ClusterBootstrap_:
		return m.bootstrap(ctx, inner.ClusterBootstrap)
	case *apb.NodeParameters_ClusterRegister_:
		return m.register(ctx, inner.ClusterRegister)
	default:
		return fmt.Errorf("node parameters misconfigured: neither cluster_bootstrap nor cluster_register set")
	}
}

func (m *Manager) register(ctx context.Context, bootstrap *apb.NodeParameters_ClusterRegister) error {
	return fmt.Errorf("unimplemented")
}

func (m *Manager) nodeParamsFWCFG(ctx context.Context) (*apb.NodeParameters, error) {
	bytes, err := os.ReadFile("/sys/firmware/qemu_fw_cfg/by_name/dev.monogon.metropolis/parameters.pb/raw")
	if err != nil {
		return nil, fmt.Errorf("could not read firmware enrolment file: %w", err)
	}

	config := apb.NodeParameters{}
	err = proto.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}

	return &config, nil
}

func (m *Manager) nodeParams(ctx context.Context) (*apb.NodeParameters, error) {
	// Retrieve node parameters from qemu's fwcfg interface or ESP.
	// TODO(q3k): probably abstract this away and implement per platform/build/...
	paramsFWCFG, err := m.nodeParamsFWCFG(ctx)
	if err != nil {
		supervisor.Logger(ctx).Warningf("Could not retrieve node parameters from qemu fwcfg: %v", err)
		paramsFWCFG = nil
	} else {
		supervisor.Logger(ctx).Infof("Retrieved node parameters from qemu fwcfg")
	}
	paramsESP, err := m.storageRoot.ESP.NodeParameters.Unmarshal()
	if err != nil {
		supervisor.Logger(ctx).Warningf("Could not retrieve node parameters from ESP: %v", err)
		paramsESP = nil
	} else {
		supervisor.Logger(ctx).Infof("Retrieved node parameters from ESP")
	}
	if paramsFWCFG == nil && paramsESP == nil {
		return nil, fmt.Errorf("could not find node parameters in ESP or qemu fwcfg")
	}
	if paramsFWCFG != nil && paramsESP != nil {
		supervisor.Logger(ctx).Warningf("Node parameters found both in both ESP and qemu fwcfg, using the latter")
		return paramsFWCFG, nil
	} else if paramsFWCFG != nil {
		return paramsFWCFG, nil
	} else {
		return paramsESP, nil
	}
}

func (m *Manager) join(ctx context.Context, cfg *ppb.SealedConfiguration) error {
	return fmt.Errorf("unimplemented")
}
