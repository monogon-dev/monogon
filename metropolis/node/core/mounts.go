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
	"strings"

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
		{"/proc", "proc", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/dev", "devtmpfs", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/dev/pts", "devpts", unix.MS_NOEXEC | unix.MS_NOSUID},
	} {
		if err := os.MkdirAll(el.dir, 0755); err != nil {
			return fmt.Errorf("could not make %s: %w", el.dir, err)
		}
		if err := unix.Mount(el.fs, el.dir, el.fs, el.flags, ""); err != nil {
			return fmt.Errorf("could not mount %s on %s: %w", el.fs, el.dir, err)
		}
	}

	// Mount all available CGroups for v1 (v2 uses a single unified hierarchy
	// and is not supported by our runtimes yet)
	if err := unix.Mount("tmpfs", "/sys/fs/cgroup", "tmpfs", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV, ""); err != nil {
		panic(err)
	}
	cgroupsRaw, err := os.ReadFile("/proc/cgroups")
	if err != nil {
		panic(err)
	}

	cgroupLines := strings.Split(string(cgroupsRaw), "\n")
	for _, cgroupLine := range cgroupLines {
		if cgroupLine == "" || strings.HasPrefix(cgroupLine, "#") {
			continue
		}
		cgroupParts := strings.Split(cgroupLine, "\t")
		cgroupName := cgroupParts[0]
		if err := os.Mkdir("/sys/fs/cgroup/"+cgroupName, 0755); err != nil {
			panic(err)
		}
		if err := unix.Mount("cgroup", "/sys/fs/cgroup/"+cgroupName, "cgroup", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV, cgroupName); err != nil {
			panic(err)
		}
	}

	// Enable hierarchical memory accounting
	useMemoryHierarchy, err := os.OpenFile("/sys/fs/cgroup/memory/memory.use_hierarchy", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	if _, err := useMemoryHierarchy.WriteString("1"); err != nil {
		panic(err)
	}
	useMemoryHierarchy.Close()
	return nil
}
