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
	"crypto/subtle"
	"encoding/hex"
	"fmt"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/network/hostsfile"
	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	ppb "source.monogon.dev/metropolis/proto/private"
)

func (m *Manager) bootstrap(ctx context.Context, bootstrap *apb.NodeParameters_ClusterBootstrap) error {
	supervisor.Logger(ctx).Infof("Bootstrapping new cluster, owner public key: %s", hex.EncodeToString(bootstrap.OwnerPublicKey))
	state, unlock := m.lock()
	defer unlock()

	ownerKey := bootstrap.OwnerPublicKey

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

	node := curator.NewNodeForBootstrap(cuk, pub)

	// Run worker to keep updating /ephemeral/hosts (and thus, /etc/hosts) with
	// our own IP address. This ensures that the node's ID always resolves to
	// its current external IP address.
	// TODO(q3k): move this out into roleserver.
	hostsfileSvc := hostsfile.Service{
		Config: hostsfile.Config{
			NodePublicKey: pub,
			Network:       m.networkService,
			Ephemeral:     &m.storageRoot.Ephemeral,
		},
	}
	if err := supervisor.Run(ctx, "hostsfile", hostsfileSvc.Run); err != nil {
		return err
	}

	// Bring up consensus with this node as the only member.
	m.consensus = consensus.New(consensus.Config{
		Data:           &m.storageRoot.Data.Etcd,
		Ephemeral:      &m.storageRoot.Ephemeral.Consensus,
		NodePrivateKey: priv,
	})

	supervisor.Logger(ctx).Infof("Bootstrapping: starting consensus...")
	if err := supervisor.Run(ctx, "consensus", m.consensus.Run); err != nil {
		return fmt.Errorf("when starting consensus: %w", err)
	}

	var ckv client.Namespaced
	cw := m.consensus.Watch()
	for {
		st, err := cw.Get(ctx)
		if err != nil {
			return fmt.Errorf("when waiting for consensus status: %w", err)
		}
		if !st.Running() {
			continue
		}
		ckv, err = st.CuratorClient()
		if err != nil {
			return fmt.Errorf("when retrieving curator client")
		}
		break
	}

	node.EnableKubernetesWorker()
	caCertBytes, nodeCertBytes, err := curator.BootstrapNodeFinish(ctx, ckv, &node, ownerKey)
	if err != nil {
		return fmt.Errorf("failed to finish bootstrap: %w", err)
	}

	creds, err := identity.NewNodeCredentials(priv, nodeCertBytes, caCertBytes)
	if err != nil {
		return fmt.Errorf("failed to use newly bootstrapped node credentials: %w", err)
	}

	// Overly cautious check: ensure that the credentials are for the public key
	// we've generated.
	if want, got := pub, []byte(creds.PublicKey()); subtle.ConstantTimeCompare(want, got) != 1 {
		return fmt.Errorf("newly bootstrapped node credentials emitted for wrong public key")
	}

	if err := creds.Save(&m.storageRoot.Data.Node.Credentials); err != nil {
		return fmt.Errorf("failed to write node credentials: %w", err)
	}

	m.status.Set(Status{
		State:       cpb.ClusterState_CLUSTER_STATE_HOME,
		Consensus:   m.consensus,
		Credentials: creds,
	})

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
