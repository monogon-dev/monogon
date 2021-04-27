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

package cluster

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
	ppb "source.monogon.dev/metropolis/proto/private"
)

// Status is returned to Cluster clients (ie., node code) on Manager.Watch/.Get.
type Status struct {
	// State is the current state of the cluster, as seen by the node.
	State ClusterState
	// Node is the configuration of this node in the cluster.
	Node *Node

	consensusClient client.Namespaced
}

// ConsensusUser is the to-level user of an etcd client in Metropolis node
// code. These need to be defined ahead of time in an Go 'enum', and different
// ConsensusUsers should not be shared by different codepaths.
type ConsensusUser string

const (
	ConsensusUserKubernetesPKI ConsensusUser = "kubernetes-pki"
)

// ConsensusClient returns an etcd/consensus client for a given ConsensusUser.
func (s *Status) ConsensusClient(user ConsensusUser) (client.Namespaced, error) {
	// Ensure that we already are connected to etcd and are in a state in which we
	// should be handing out cluster connectivity.
	if s.consensusClient == nil {
		return nil, fmt.Errorf("not connected")
	}
	switch s.State {
	case ClusterHome:
	case ClusterSplit:
		return nil, fmt.Errorf("refusing connection with cluster state %v", s.State)
	default:
	}

	// Ensure only defined 'applications' are used to prevent programmer error and
	// casting to ConsensusUser from an arbitrary string.
	switch user {
	case ConsensusUserKubernetesPKI:
	default:
		return nil, fmt.Errorf("unknown ConsensusUser %q", user)
	}
	client, err := s.consensusClient.Sub(string(user))
	if err != nil {
		return nil, fmt.Errorf("retrieving subclient failed: %w", err)
	}
	return client, nil
}

type state struct {
	mu sync.RWMutex

	oneway       bool
	stateCluster ClusterState
	stateNode    ppb.Node_FSMState

	configuration *ppb.SealedConfiguration
}

type Watcher struct {
	event.Watcher
}

func (w *Watcher) Get(ctx context.Context) (*Status, error) {
	val, err := w.Watcher.Get(ctx)
	if err != nil {
		return nil, err
	}
	status := val.(Status)
	return &status, err
}

// GetHome waits until the cluster, from the point of view of this node, is in
// the ClusterHome state. This can be used to wait for the cluster manager to
// 'settle', before clients start more node services.
func (w *Watcher) GetHome(ctx context.Context) (*Status, error) {
	for {
		status, err := w.Get(ctx)
		if err != nil {
			return nil, err
		}
		switch status.State {
		case ClusterHome:
			return status, nil
		case ClusterDisowning:
			return nil, fmt.Errorf("the cluster has disowned this node")
		}
	}
}

func (m *Manager) Watch() Watcher {
	return Watcher{
		Watcher: m.status.Watch(),
	}
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

		state: state{
			stateCluster: ClusterUnknown,
			stateNode:    ppb.Node_FSM_STATE_INVALID,
		},
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
	bytes, err := ioutil.ReadFile("/sys/firmware/qemu_fw_cfg/by_name/dev.monogon.metropolis/parameters.pb/raw")
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

// Node returns the Node that the Manager brought into a cluster, or nil if the
// Manager is not Running.  This is safe to call from any goroutine.
func (m *Manager) Node() *Node {
	return nil
}
