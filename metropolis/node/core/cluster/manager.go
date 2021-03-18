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
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/pki"
	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
	ppb "source.monogon.dev/metropolis/proto/private"
)

type managerResult struct {
	node *Node
	err  error
}

type state struct {
	mu sync.RWMutex

	oneway       bool
	stateCluster ClusterState
	stateNode    ppb.Node_FSMState

	configuration *ppb.SealedConfiguration

	result  *managerResult
	waiters []chan *managerResult
}

func (s *state) setResult(node *Node, err error) {
	s.result = &managerResult{
		node: node,
		err:  err,
	}
	for _, w := range s.waiters {
		go func(c chan *managerResult) {
			c <- s.result
		}(w)
	}
	s.waiters = nil
}

type Manager struct {
	storageRoot    *localstorage.Root
	networkService *network.Service

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

func (m *Manager) Wait() (*Node, error) {
	state, unlock := m.lock()

	if state.result != nil {
		unlock()
		return state.result.node, state.result.err
	}

	c := make(chan *managerResult)
	state.waiters = append(state.waiters, c)
	unlock()
	res := <-c
	return res.node, res.err
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

func (m *Manager) bootstrap(ctx context.Context, bootstrap *apb.NodeParameters_ClusterBootstrap) error {
	supervisor.Logger(ctx).Infof("Bootstrapping new cluster, owner public key: %s", hex.EncodeToString(bootstrap.OwnerPublicKey))
	state, unlock := m.lock()
	defer unlock()

	state.configuration = &ppb.SealedConfiguration{}

	// Mount new storage with generated CUK, and save LUK into sealed config proto.
	supervisor.Logger(ctx).Infof("Bootstrapping: mounting new storage...")
	cuk, err := m.storageRoot.Data.MountNew(state.configuration)
	if err != nil {
		return fmt.Errorf("could not make and mount data partition: %w", err)
	}

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate node keypair: %w", err)
	}
	supervisor.Logger(ctx).Infof("Bootstrapping: node public key: %s", hex.EncodeToString([]byte(pub)))

	node := Node{
		clusterUnlockKey: cuk,
		pubkey:           pub,
		state:            ppb.Node_FSM_STATE_UP,
		// TODO(q3k): make this configurable.
		consensusMember:  &NodeRoleConsensusMember{},
		kubernetesWorker: &NodeRoleKubernetesWorker{},
	}

	// Run worker to keep updating /ephemeral/hosts (and thus, /etc/hosts) with
	// our own IP address. This ensures that the node's ID always resolves to
	// its current external IP address.
	supervisor.Run(ctx, "hostsfile", func(ctx context.Context) error {
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		watcher := m.networkService.Watch()
		for {
			status, err := watcher.Get(ctx)
			if err != nil {
				return err
			}
			err = node.ConfigureLocalHostname(ctx, &m.storageRoot.Ephemeral, status.ExternalAddress)
			if err != nil {
				return fmt.Errorf("could not configure hostname: %w", err)
			}
		}
	})

	// Bring up consensus with this node as the only member.
	m.consensus = consensus.New(consensus.Config{
		Data:       &m.storageRoot.Data.Etcd,
		Ephemeral:  &m.storageRoot.Ephemeral.Consensus,
		NewCluster: true,
		Name:       node.ID(),
	})

	supervisor.Logger(ctx).Infof("Bootstrapping: starting consensus...")
	if err := supervisor.Run(ctx, "consensus", m.consensus.Run); err != nil {
		return fmt.Errorf("when starting consensus: %w", err)
	}

	supervisor.Logger(ctx).Info("Bootstrapping: waiting for consensus...")
	if err := m.consensus.WaitReady(ctx); err != nil {
		return fmt.Errorf("consensus service failed to become ready: %w", err)
	}
	supervisor.Logger(ctx).Info("Bootstrapping: consensus ready.")

	kv := m.consensus.KVRoot()
	node.KV = kv

	// Create Metropolis CA and this node's certificate.
	caCertBytes, _, err := PKICA.Ensure(ctx, kv)
	if err != nil {
		return fmt.Errorf("failed to create cluster CA: %w", err)
	}
	nodeCert := PKINamespace.New(PKICA, "", pki.Server([]string{node.ID()}, nil))
	nodeCert.UseExistingKey(priv)
	nodeCertBytes, _, err := nodeCert.Ensure(ctx, kv)
	if err != nil {
		return fmt.Errorf("failed to create node certificate: %w", err)
	}

	if err := m.storageRoot.Data.Node.Credentials.CACertificate.Write(caCertBytes, 0400); err != nil {
		return fmt.Errorf("failed to write CA certificate: %w", err)
	}
	if err := m.storageRoot.Data.Node.Credentials.Certificate.Write(nodeCertBytes, 0400); err != nil {
		return fmt.Errorf("failed to write node certificate: %w", err)
	}
	if err := m.storageRoot.Data.Node.Credentials.Key.Write(priv, 0400); err != nil {
		return fmt.Errorf("failed to write node private key: %w", err)
	}

	// Update our Node obejct in etcd.
	if err := node.Store(ctx, kv); err != nil {
		return fmt.Errorf("failed to store new node in etcd: %w", err)
	}

	state.setResult(&node, nil)

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
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
		supervisor.Logger(ctx).Warningf("Node parameters found both inboth ESP and qemu fwcfg, using the latter")
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
