// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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
		{"/proc", "proc", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/dev", "devtmpfs", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/dev/pts", "devpts", unix.MS_NOEXEC | unix.MS_NOSUID},
		// Nothing in Metropolis currently uses /dev/shm, but it's required
		// by containerd when the host IPC namespace is shared, which
		// is required by "kubectl debug node/" and specific customer applications.
		// https://github.com/monogon/monogon/issues/305.
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
