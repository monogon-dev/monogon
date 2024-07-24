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

package main

import (
	"fmt"
	"os"

	"github.com/opencontainers/runc/libcontainer/cgroups"
	"golang.org/x/sys/unix"
)

// setupMounts sets up basic mounts like sysfs, procfs, devtmpfs and cgroups.
// This should be called early during init as a lot of processes depend on this
// being available.
func setupMounts() error {
	// Set up target filesystems.
	for _, el := range []struct {
		dir   string
		fs    string
		flags uintptr
	}{
		{"/sys", "sysfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/sys/kernel/tracing", "tracefs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/sys/firmware/efi/efivars", "efivarfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/sys/fs/pstore", "pstore", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		// Nothing in Metropolis currently requires BPF maps,
		// but some privileged customer applications do.
		{"/sys/fs/bpf", "bpf", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/proc", "proc", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/dev", "devtmpfs", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/dev/pts", "devpts", unix.MS_NOEXEC | unix.MS_NOSUID},
		// Nothing in Metropolis currently uses /dev/shm, but it's required
		// by containerd when the host IPC namespace is shared, which
		// is required by "kubectl debug node/" and specific customer applications.
		// https://github.com/monogon-dev/monogon/issues/305.
		{"/dev/shm", "tmpfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
	} {
		if err := os.MkdirAll(el.dir, 0755); err != nil {
			return fmt.Errorf("could not make %s: %w", el.dir, err)
		}
		if err := unix.Mount(el.fs, el.dir, el.fs, el.flags, ""); err != nil {
			return fmt.Errorf("could not mount %s on %s: %w", el.fs, el.dir, err)
		}
	}

	if err := unix.Mount("cgroup2", "/sys/fs/cgroup", "cgroup2", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV, "nsdelegate,memory_recursiveprot"); err != nil {
		panic(err)
	}
	// Create main cgroup "everything" and move ourselves into it.
	if err := os.Mkdir("/sys/fs/cgroup/everything", 0755); err != nil {
		panic(err)
	}
	if err := cgroups.WriteCgroupProc("/sys/fs/cgroup/everything", os.Getpid()); err != nil {
		panic(err)
	}
	return nil
}
