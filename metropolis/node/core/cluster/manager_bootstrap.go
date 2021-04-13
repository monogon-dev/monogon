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
	"fmt"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/pkg/pki"
	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
	ppb "source.monogon.dev/metropolis/proto/private"
)

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

