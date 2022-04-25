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

	"source.monogon.dev/metropolis/pkg/supervisor"
	apb "source.monogon.dev/metropolis/proto/api"
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
	nuk := state.configuration.NodeUnlockKey

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate node keypair: %w", err)
	}
	supervisor.Logger(ctx).Infof("Bootstrapping: node public key: %s", hex.EncodeToString([]byte(pub)))

	jpub, jpriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate join keypair: %w", err)
	}
	supervisor.Logger(ctx).Infof("Bootstrapping: node public join key: %s", hex.EncodeToString([]byte(jpub)))

	m.roleServer.ProvideBootstrapData(priv, ownerKey, cuk, nuk, jpriv)

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
