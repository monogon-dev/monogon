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

package localstorage

import (
	"crypto/rand"
	"fmt"
	"os/exec"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/node/core/localstorage/crypt"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/pkg/tpm"
	ppb "source.monogon.dev/metropolis/proto/private"
)

var keySize uint16 = 256 / 8

// MountData mounts the node data partition with the given global unlock key.
// It automatically unseals the local unlock key from the TPM.
func (d *DataDirectory) MountExisting(config *ppb.SealedConfiguration, clusterUnlockKey []byte) error {
	d.flagLock.Lock()
	defer d.flagLock.Unlock()

	if !d.canMount {
		return fmt.Errorf("cannot mount yet (root not ready?)")
	}
	if d.mounted {
		return fmt.Errorf("already mounted")
	}
	d.mounted = true

	key := make([]byte, keySize)
	for i := uint16(0); i < keySize; i++ {
		key[i] = config.NodeUnlockKey[i] ^ clusterUnlockKey[i]
	}

	if err := crypt.CryptMap("data", crypt.NodeDataCryptPath, key); err != nil {
		return err
	}
	if err := d.mount(); err != nil {
		return err
	}
	return nil
}

// InitializeData initializes the node data partition and returns the node and
// cluster unlock keys. It seals the local portion into the TPM. This is a
// potentially slow operation since it touches the whole partition.
func (d *DataDirectory) MountNew(config *ppb.SealedConfiguration) ([]byte, error) {
	d.flagLock.Lock()
	defer d.flagLock.Unlock()
	if !d.canMount {
		return nil, fmt.Errorf("cannot mount yet (root not ready?)")
	}
	if d.mounted {
		return nil, fmt.Errorf("already mounted")
	}
	d.mounted = true

	var nodeUnlockKey, globalUnlockKey []byte
	var err error
	if tpm.IsInitialized() {
		nodeUnlockKey, err = tpm.GenerateSafeKey(keySize)
	} else {
		nodeUnlockKey = make([]byte, keySize)
		_, err = rand.Read(nodeUnlockKey)
	}
	if err != nil {
		return nil, fmt.Errorf("generating local unlock key: %w", err)
	}
	if tpm.IsInitialized() {
		globalUnlockKey, err = tpm.GenerateSafeKey(keySize)
	} else {
		globalUnlockKey = make([]byte, keySize)
		_, err = rand.Read(globalUnlockKey)
	}
	if err != nil {
		return nil, fmt.Errorf("generating global unlock key: %w", err)
	}

	// The actual key is generated by XORing together the nodeUnlockKey and the
	// globalUnlockKey This provides us with a mathematical guarantee that the
	// resulting key cannot be recovered whithout knowledge of both parts.
	key := make([]byte, keySize)
	for i := uint16(0); i < keySize; i++ {
		key[i] = nodeUnlockKey[i] ^ globalUnlockKey[i]
	}

	if err := crypt.CryptInit("data", crypt.NodeDataCryptPath, key); err != nil {
		return nil, fmt.Errorf("initializing encrypted block device: %w", err)
	}
	mkfsCmd := exec.Command("/bin/mkfs.xfs", "-qKf", "/dev/data")
	if _, err := mkfsCmd.Output(); err != nil {
		return nil, fmt.Errorf("formatting encrypted block device: %w", err)
	}

	if err := d.mount(); err != nil {
		return nil, fmt.Errorf("mounting: %w", err)
	}

	// TODO(q3k): do this automatically?
	for _, d := range []declarative.DirectoryPlacement{
		d.Etcd, d.Etcd.Data, d.Etcd.PeerPKI,
		d.Containerd,
		d.Kubernetes,
		d.Kubernetes.Kubelet, d.Kubernetes.Kubelet.Plugins, d.Kubernetes.Kubelet.PluginsRegistry,
		d.Kubernetes.ClusterNetworking,
		d.Node,
		d.Node.Credentials,
		d.Volumes,
	} {
		err := d.MkdirAll(0700)
		if err != nil {
			return nil, fmt.Errorf("creating directory failed: %w", err)
		}
	}

	config.NodeUnlockKey = nodeUnlockKey

	return globalUnlockKey, nil
}

func (d *DataDirectory) mount() error {
	// TODO(T965): MS_NODEV should definitely be set on the data partition, but as long as the kubelet root
	// is on there, we can't do it.
	if err := unix.Mount("/dev/data", d.FullPath(), "xfs", unix.MS_NOEXEC, "pquota"); err != nil {
		return fmt.Errorf("mounting data directory: %w", err)
	}
	return nil
}
