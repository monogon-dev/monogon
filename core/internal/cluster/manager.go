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
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"

	"git.monogon.dev/source/nexantic.git/core/internal/common/supervisor"
	"git.monogon.dev/source/nexantic.git/core/internal/consensus"
	"git.monogon.dev/source/nexantic.git/core/internal/localstorage"
	"git.monogon.dev/source/nexantic.git/core/internal/network"
)

// Manager is a finite state machine that joins this node (ie., Smalltown instance running on a virtual/physical machine)
// into a Smalltown cluster (ie. group of nodes that act as a single control plane for Smalltown services). It does that
// by bringing up all required operating-system level components, including mounting the local filesystem, bringing up
// a consensus (etcd) server/client, ...
//
// The Manager runs as a single-shot Runnable. It will attempt to progress its state from the initial state (New) to
// either Running (meaning that the node is now part of a cluster), or Failed (meaning that the node couldn't become
// part of a cluster). It is not restartable, as it mutates quite a bit of implicit operating-system level state (like
// filesystem mounts).  As such, it's difficult to recover reliably from failures, and since these failures indicate
// some high issues with the cluster configuration/state, a failure requires a full kernel reboot to retry (or fix/
// reconfigure the node).
//
// Currently, the Manager only supports one flow for bringing up a Node: by creating a new cluster. As such, it's
// missing the following flows:
//  - joining a new node into an already running cluster
//  - restarting a node into an already existing cluster
//  - restarting a node into an already running cluster (ie. full reboot of whole cluster)
//
type Manager struct {
	storageRoot    *localstorage.Root
	networkService *network.Service

	// stateLock locks all state* variables.
	stateLock sync.RWMutex
	// state is the FSM state of the Manager.
	state State
	// stateRunningNode is the Node that this Manager got from joining a cluster. It's only valid if the Manager is
	// Running.
	stateRunningNode *Node
	// stateWaiters is a list of channels that wish to be notified (by sending true or false) for when the Manager
	// reaches a final state (Running or Failed respectively).
	stateWaiters []chan bool

	// consensus is the spawned etcd/consensus service, if the Manager brought up a Node that should run one.
	consensus *consensus.Service
}

// NewManager creates a new cluster Manager. The given localstorage Root must be places, but not yet started (and will
// be started as the Manager makes progress). The given network Service must already be running.
func NewManager(storageRoot *localstorage.Root, networkService *network.Service) *Manager {
	return &Manager{
		storageRoot:    storageRoot,
		networkService: networkService,
	}
}

// State is the state of the Manager finite state machine.
type State int

const (
	// StateNew is the initial state of the Manager. It decides how to go about joining or creating a cluster.
	StateNew State = iota
	// StateCreatingCluster is when the Manager attempts to create a new cluster - this happens when a node is started
	// with no EnrolmentConfig.
	StateCreatingCluster
	// StateRunning is when the Manager successfully got the node to be part of a cluster. stateRunningNode is valid.
	StateRunning
	// StateFailed is when the Manager failed to ge the node to be part of a cluster.
	StateFailed
)

func (s State) String() string {
	switch s {
	case StateNew:
		return "New"
	case StateCreatingCluster:
		return "CreatingCluster"
	case StateRunning:
		return "Running"
	case StateFailed:
		return "Failed"
	default:
		return "UNKNOWN"
	}
}

// allowedTransition describes all allowed state transitions (map[From][]To).
var allowedTransitions = map[State][]State{
	StateNew:             {StateCreatingCluster},
	StateCreatingCluster: {StateRunning, StateFailed},
}

// allowed returns whether a transition from a state to another state is allowed (ie. is defined in allowedTransitions).
func (m *Manager) allowed(from, to State) bool {
	for _, allowed := range allowedTransitions[from] {
		if to == allowed {
			return true
		}
	}
	return false
}

// next moves the Manager finite state machine from its current state to `n`, or to Failed if the transition is not
// allowed.
func (m *Manager) next(ctx context.Context, n State) {
	m.stateLock.Lock()
	defer m.stateLock.Unlock()

	if !m.allowed(m.state, n) {
		supervisor.Logger(ctx).Error("Attempted invalid enrolment state transition, failing enrolment",
			zap.String("from", m.state.String()), zap.String("to", m.state.String()))
		m.state = StateFailed
		return
	}

	supervisor.Logger(ctx).Info("Enrolment state change",
		zap.String("from", m.state.String()), zap.String("to", n.String()))

	m.state = n
}

// State returns the state of the Manager. It's safe to call this from any goroutine.
func (m *Manager) State() State {
	m.stateLock.RLock()
	defer m.stateLock.RUnlock()
	return m.state
}

// WaitFinished waits until the Manager FSM reaches Running or Failed, and returns true if the FSM is Running. It's
// safe to call this from any goroutine.
func (m *Manager) WaitFinished() (success bool) {
	m.stateLock.Lock()
	switch m.state {
	case StateFailed:
		m.stateLock.Unlock()
		return false
	case StateRunning:
		m.stateLock.Unlock()
		return true
	}

	C := make(chan bool)
	m.stateWaiters = append(m.stateWaiters, C)
	m.stateLock.Unlock()
	return <-C
}

// wakeWaiters wakes any WaitFinished waiters and lets them know about the current state of the Manager.
// The stateLock must already been taken, and the state must have been set in the same critical section (otherwise
// this can cause a race condition).
func (m *Manager) wakeWaiters() {
	state := m.state
	waiters := m.stateWaiters
	m.stateWaiters = nil

	for _, waiter := range waiters {
		go func(w chan bool) {
			w <- state == StateRunning
		}(waiter)
	}
}

// Run is the runnable of the Manager, to be started using the Supervisor. It is one-shot, and should not be restarted.
func (m *Manager) Run(ctx context.Context) error {
	if state := m.State(); state != StateNew {
		supervisor.Logger(ctx).Error("Manager started with non-New state, failing", zap.String("state", state.String()))
		m.stateLock.Lock()
		m.state = StateFailed
		m.wakeWaiters()
		m.stateLock.Unlock()
		return nil
	}

	var err error
	bo := backoff.NewExponentialBackOff()
	for {
		done := false
		state := m.State()
		switch state {
		case StateNew:
			err = m.stateNew(ctx)
		case StateCreatingCluster:
			err = m.stateCreatingCluster(ctx)
		default:
			done = true
			break
		}

		if err != nil || done {
			break
		}

		if state == m.State() && !m.allowed(state, m.State()) {
			supervisor.Logger(ctx).Error("Enrolment got stuck, failing", zap.String("state", m.state.String()))
			m.stateLock.Lock()
			m.state = StateFailed
			m.stateLock.Unlock()
		} else {
			bo.Reset()
		}
	}

	m.stateLock.Lock()
	state := m.state
	if state != StateRunning {
		supervisor.Logger(ctx).Error("Enrolment failed", zap.Error(err), zap.String("state", m.state.String()))
	} else {
		supervisor.Logger(ctx).Info("Enrolment successful!")
	}
	m.wakeWaiters()
	m.stateLock.Unlock()

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}

// stateNew is called when a Manager is New. It makes the decision on how to join this node into a cluster.
func (m *Manager) stateNew(ctx context.Context) error {
	supervisor.Logger(ctx).Info("Starting enrolment process...")

	// Check for presence of EnrolmentConfig on ESP or in qemu firmware variables.
	var configRaw []byte
	configRaw, err := m.storageRoot.ESP.Enrolment.Read()
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("could not read local enrolment file: %w", err)
	} else if err != nil {
		configRaw, err = ioutil.ReadFile("/sys/firmware/qemu_fw_cfg/by_name/com.nexantic.smalltown/enrolment.pb/raw")
		if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("could not read firmware enrolment file: %w", err)
		}
	}

	// If no enrolment file exists, we create a new cluster.
	if configRaw == nil {
		m.next(ctx, StateCreatingCluster)
		return nil
	}

	// Enrolment file exists, this is not yet implemented (need to enroll into or join existing cluster).
	return fmt.Errorf("unimplemented join/enroll")
}

// stateCreatingCluster is called when the Manager has decided to create a new cluster.
//
// The process to create a new cluster is as follows:
//   - wait for IP address
//   - initialize new data partition, by generating local and cluster unlock keys (the local unlock key is saved to
//     the ESP, while the cluster unlock key is returned)
//   - create a new node certificate and Node (with new given cluster unlock key)
//   - start up a new etcd cluster, with this node being the only member
//   - save the new Node to the new etcd cluster (thereby saving the node's cluster unlock key to etcd)
func (m *Manager) stateCreatingCluster(ctx context.Context) error {
	logger := supervisor.Logger(ctx)
	logger.Info("Creating new cluster: waiting for IP address...")
	ip, err := m.networkService.GetIP(ctx, true)
	if err != nil {
		return fmt.Errorf("when getting IP address: %w", err)
	}
	logger.Info("Creating new cluster: got IP address", zap.String("address", ip.String()))

	logger.Info("Creating new cluster: initializing storage...")
	cuk, err := m.storageRoot.Data.MountNew(&m.storageRoot.ESP.LocalUnlock)
	if err != nil {
		return fmt.Errorf("when making new data partition: %w", err)
	}
	logger.Info("Creating new cluster: storage initialized")

	// Create certificate for node.
	cert, err := m.storageRoot.Data.Node.EnsureSelfSigned(localstorage.CertificateForNode)
	if err != nil {
		return fmt.Errorf("failed to create new node certificate: %w", err)
	}

	node := NewNode(cuk, *ip, *cert.Leaf)

	m.consensus = consensus.New(consensus.Config{
		Data:           &m.storageRoot.Data.Etcd,
		Ephemeral:      &m.storageRoot.Ephemeral.Consensus,
		NewCluster:     true,
		Name:           node.ID(),
		InitialCluster: ip.String(),
		ExternalHost:   ip.String(),
		ListenHost:     ip.String(),
	})
	if err := supervisor.Run(ctx, "consensus", m.consensus.Run); err != nil {
		return fmt.Errorf("when starting consensus: %w", err)
	}

	// TODO(q3k): make timeout configurable?
	ctxT, ctxC := context.WithTimeout(ctx, 5*time.Second)
	defer ctxC()

	supervisor.Logger(ctx).Info("Creating new cluster: waiting for consensus...")
	if err := m.consensus.WaitReady(ctxT); err != nil {
		return fmt.Errorf("consensus service failed to become ready: %w", err)
	}

	// Configure node to be a consensus member and kubernetes worker. In the future, different nodes will have
	// different roles, but for now they're all symmetrical.
	_, consensusName, err := m.consensus.MemberInfo(ctx)
	if err != nil {
		return fmt.Errorf("could not get consensus MemberInfo: %w", err)
	}
	if err := node.MakeConsensusMember(consensusName); err != nil {
		return fmt.Errorf("could not make new node into consensus member: %w", err)
	}
	if err := node.MakeKubernetesWorker(node.ID()); err != nil {
		return fmt.Errorf("could not make new node into kubernetes worker: %w", err)
	}

	// Save node into etcd.
	supervisor.Logger(ctx).Info("Creating new cluster: storing first node...")
	if err := node.Store(ctx, m.consensus.KV("cluster", "enrolment")); err != nil {
		return fmt.Errorf("could not save new node: %w", err)
	}

	m.stateLock.Lock()
	m.stateRunningNode = node
	m.stateLock.Unlock()

	m.next(ctx, StateRunning)
	return nil
}

// Node returns the Node that the Manager brought into a cluster, or nil if the Manager is not Running.
// This is safe to call from any goroutine.
func (m *Manager) Node() *Node {
	m.stateLock.Lock()
	defer m.stateLock.Unlock()
	if m.state != StateRunning {
		return nil
	}
	return m.stateRunningNode
}

// ConsensusKV returns a namespaced etcd KV client, or nil if the Manager is not Running.
// This is safe to call from any goroutine.
func (m *Manager) ConsensusKV(module, space string) clientv3.KV {
	m.stateLock.Lock()
	defer m.stateLock.Unlock()
	if m.state != StateRunning {
		return nil
	}
	if m.stateRunningNode.ConsensusMember() == nil {
		// TODO(q3k): in this case, we should return a client to etcd even though this
		// node is not a member of consensus. For now, all nodes are consensus members.
		return nil
	}
	return m.consensus.KV(module, space)
}

// ConsensusKVRoot returns a non-namespaced etcd KV client, or nil if the Manager is not Running.
// This is safe to call from any goroutine.
func (m *Manager) ConsensusKVRoot() clientv3.KV {
	m.stateLock.Lock()
	defer m.stateLock.Unlock()
	if m.state != StateRunning {
		return nil
	}
	if m.stateRunningNode.ConsensusMember() == nil {
		// TODO(q3k): in this case, we should return a client to etcd even though this
		// node is not a member of consensus. For now, all nodes are consensus members.
		return nil
	}
	return m.consensus.KVRoot()
}

// ConsensusCluster returns an etcd Cluster client, or nil if the Manager is not Running.
// This is safe to call from any goroutine.
func (m *Manager) ConsensusCluster() clientv3.Cluster {
	m.stateLock.Lock()
	defer m.stateLock.Unlock()
	if m.state != StateRunning {
		return nil
	}
	if m.stateRunningNode.ConsensusMember() == nil {
		// TODO(q3k): in this case, we should return a client to etcd even though this
		// node is not a member of consensus. For now, all nodes are consensus members.
		return nil
	}
	return m.consensus.Cluster()
}
