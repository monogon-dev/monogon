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
	"context"
	"fmt"
	"os"

	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/node/core/localstorage/crypt"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
)

func (r *Root) Start(ctx context.Context) error {
	r.Data.flagLock.Lock()
	defer r.Data.flagLock.Unlock()
	if r.Data.canMount {
		return fmt.Errorf("cannot re-start root storage")
	}
	// TODO(q3k): turn this into an Ensure call
	err := crypt.MakeBlockDevices(ctx)
	if err != nil {
		return fmt.Errorf("MakeBlockDevices: %w", err)
	}

	if err := unix.Mount(crypt.ESPDevicePath, r.ESP.FullPath(), "vfat", unix.MS_NOEXEC|unix.MS_NODEV|unix.MS_SYNCHRONOUS, ""); err != nil {
		return fmt.Errorf("mounting ESP partition: %w", err)
	}

	r.Data.canMount = true

	if err := unix.Mount("tmpfs", r.Tmp.FullPath(), "tmpfs", unix.MS_NOEXEC|unix.MS_NODEV, ""); err != nil {
		return fmt.Errorf("mounting /tmp: %w", err)
	}

	if err := unix.Mount("tmpfs", r.Ephemeral.FullPath(), "tmpfs", unix.MS_NODEV, ""); err != nil {
		return fmt.Errorf("mounting /ephemeral: %v", err)
	}

	if err := unix.Mount("tmpfs", r.Run.FullPath(), "tmpfs", unix.MS_NOEXEC|unix.MS_NODEV, ""); err != nil {
		return fmt.Errorf("mounting /run: %w", err)
	}

	// TODO(q3k): do this automatically?
	for _, d := range []declarative.DirectoryPlacement{
		r.Ephemeral.Consensus,
		r.Ephemeral.Containerd, r.Ephemeral.Containerd.Tmp, r.Ephemeral.Containerd.RunSC, r.Ephemeral.Containerd.IPAM,
		r.Ephemeral.FlexvolumePlugins,
	} {
		err := d.MkdirAll(0700)
		if err != nil {
			return fmt.Errorf("creating directory failed: %w", err)
		}
	}

	for _, d := range []declarative.DirectoryPlacement{
		r.Ephemeral, r.Ephemeral.Containerd, r.Ephemeral.Containerd.Tmp,
	} {
		if err := os.Chmod(d.FullPath(), 0755); err != nil {
			return fmt.Errorf("failed to chmod containerd tmp path: %w", err)
		}
	}

	return nil
}
