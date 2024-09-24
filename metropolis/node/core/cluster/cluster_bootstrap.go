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
	"time"

	"google.golang.org/protobuf/proto"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/roleserve"
	"source.monogon.dev/osbase/supervisor"

	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	ppb "source.monogon.dev/metropolis/proto/private"
)

func (m *Manager) bootstrap(ctx context.Context, bootstrap *apb.NodeParameters_ClusterBootstrap) error {
	supervisor.Logger(ctx).Infof("Bootstrapping new cluster, owner public key: %s", hex.EncodeToString(bootstrap.OwnerPublicKey))

	var cc *curator.Cluster

	if bootstrap.InitialClusterConfiguration == nil {
		supervisor.Logger(ctx).Infof("No initial cluster configuration provided, using defaults.")
		cc = curator.DefaultClusterConfiguration()
	} else {
		var err error
		cc, err = curator.ClusterConfigurationFromInitial(bootstrap.InitialClusterConfiguration)
		if err != nil {
			return fmt.Errorf("invalid initial cluster configuration: %w", err)
		}
	}

	tpmUsage, err := cc.NodeTPMUsage(m.haveTPM)
	if err != nil {
		return fmt.Errorf("cannot bootstrap cluster: %w", err)
	}

	storageSecurity, err := cc.NodeStorageSecurity()
	if err != nil {
		return fmt.Errorf("cannot bootstrap cluster: %w", err)
	}

	supervisor.Logger(ctx).Infof("TPM: cluster policy: %s, node: %s", cc.TPMMode, tpmUsage)
	supervisor.Logger(ctx).Infof("Storage Security: cluster policy: %s, node: %s", cc.StorageSecurityPolicy, storageSecurity)

	ownerKey := bootstrap.OwnerPublicKey
	var configuration ppb.SealedConfiguration

	// Mount new storage with generated CUK, and save NUK into sealed config proto.
	supervisor.Logger(ctx).Infof("Bootstrapping: mounting new storage...")
	storageDone := make(chan struct{})
	go func() {
		t := time.NewTicker(5 * time.Second)
		defer t.Stop()
		select {
		case <-storageDone:
			return
		case <-t.C:
			supervisor.Logger(ctx).Infof("Bootstrapping: still waiting for storage....")
		}
	}()
	cuk, err := m.storageRoot.Data.MountNew(&configuration, storageSecurity)
	close(storageDone)
	if err != nil {
		return fmt.Errorf("could not make and mount data partition: %w", err)
	}
	nuk := configuration.NodeUnlockKey

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate node keypair: %w", err)
	}
	id := identity.NodeID(pub)
	supervisor.Logger(ctx).Infof("Bootstrapping: node public key: %s", hex.EncodeToString(pub))

	jpub, jpriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate join keypair: %w", err)
	}
	supervisor.Logger(ctx).Infof("Bootstrapping: node public join key: %s", hex.EncodeToString(jpub))

	directory := &cpb.ClusterDirectory{
		Nodes: []*cpb.ClusterDirectory_Node{
			{
				Id: id,
				Addresses: []*cpb.ClusterDirectory_Node_Address{
					{
						Host: "127.0.0.1",
					},
				},
			},
		},
	}
	cdirRaw, err := proto.Marshal(directory)
	if err != nil {
		return fmt.Errorf("couldn't marshal ClusterDirectory: %w", err)
	}
	if err = m.storageRoot.ESP.Metropolis.ClusterDirectory.Write(cdirRaw, 0644); err != nil {
		return fmt.Errorf("writing cluster directory failed: %w", err)
	}

	sc := ppb.SealedConfiguration{
		NodeUnlockKey: nuk,
		JoinKey:       jpriv,
		// No ClusterCA yet, that's added by the roleserver after it finishes curator
		// bootstrap.
		ClusterCa:       nil,
		StorageSecurity: storageSecurity,
	}
	if err = m.storageRoot.ESP.Metropolis.SealedConfiguration.SealSecureBoot(&sc, tpmUsage); err != nil {
		return fmt.Errorf("writing sealed configuration failed: %w", err)
	}
	supervisor.Logger(ctx).Infof("Saved bootstrapped node's credentials.")

	labels := make(map[string]string)
	if l := bootstrap.Labels; l != nil {
		if nlabels := len(l.Pairs); nlabels > common.MaxLabelsPerNode {
			supervisor.Logger(ctx).Warningf("Too many labels (%d, limit %d), truncating...", nlabels, common.MaxLabelsPerNode)
			l.Pairs = l.Pairs[:common.MaxLabelsPerNode]
		}
		for _, pair := range l.Pairs {
			if err := common.ValidateLabelKey(pair.Key); err != nil {
				supervisor.Logger(ctx).Warningf("Skipping label %q/%q: key invalid: %v", pair.Key, pair.Value, err)
				continue
			}
			if err := common.ValidateLabelValue(pair.Value); err != nil {
				supervisor.Logger(ctx).Warningf("Skipping label %q/%q: value invalid: %v", pair.Key, pair.Value, err)
				continue
			}
			if _, ok := labels[pair.Key]; ok {
				supervisor.Logger(ctx).Warningf("Label %q/%q: repeated key, overwriting previous value", pair.Key, pair.Value)
			}
			labels[pair.Key] = pair.Value
		}
	}

	bd := roleserve.BootstrapData{}
	bd.Node.ID = id
	bd.Node.PrivateKey = priv
	bd.Node.ClusterUnlockKey = cuk
	bd.Node.NodeUnlockKey = nuk
	bd.Node.JoinKey = jpriv
	bd.Node.TPMUsage = tpmUsage
	bd.Node.Labels = labels
	bd.Cluster.InitialOwnerKey = ownerKey
	bd.Cluster.Configuration = cc
	m.roleServer.ProvideBootstrapData(&bd)

	if err := m.updateService.MarkBootSuccessful(); err != nil {
		supervisor.Logger(ctx).Errorf("Failed to mark boot as successful: %v", err)
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	supervisor.Signal(ctx, supervisor.SignalDone)
	return nil
}
