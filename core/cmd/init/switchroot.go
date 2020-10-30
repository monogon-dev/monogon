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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"git.monogon.dev/source/nexantic.git/core/pkg/logtree"

	"golang.org/x/sys/unix"
)

// switchRoot moves the root from initramfs into a tmpfs
// This is necessary because you cannot pivot_root from a initramfs (and runsc wants to do that).
// In the future, we should instead use something like squashfs instead of an initramfs and just nuke this.
func switchRoot(log logtree.LeveledLogger) error {
	// We detect the need to remount to tmpfs over env vars.
	// The first run of /init (from initramfs) will not have this var, and will be re-exec'd from a new tmpfs root with
	// that variable set.
	witness := "SIGNOS_REMOUNTED"

	// If the witness env var is found in the environment, it means we are ready to go.
	environ := os.Environ()
	for _, env := range environ {
		if strings.HasPrefix(env, witness+"=") {
			log.Info("Smalltown running in tmpfs root")
			return nil
		}
	}

	// Otherwise, we need to remount to a tmpfs.
	environ = append(environ, witness+"=yes")
	log.Info("Smalltown running in initramfs, remounting to tmpfs...")

	// Make note of all directories we have to make and files that we have to copy.
	paths := []string{}
	dirs := []string{}
	err := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == "/" {
			return nil
		}
		// /dev is prepopulated by the initramfs, skip that. The target root uses devtmpfs.
		if path == "/dev" || strings.HasPrefix(path, "/dev/") {
			return nil
		}

		if info.IsDir() {
			dirs = append(dirs, path)
		} else {
			paths = append(paths, path)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("could not list root files: %w", err)
	}

	log.Info("Copying paths to tmpfs:")
	for _, p := range paths {
		log.Infof(" - %s", p)
	}

	// Make new root at /mnt
	if err := os.Mkdir("/mnt", 0755); err != nil {
		return fmt.Errorf("could not make /mnt: %w", err)
	}
	// And mount a tmpfs on it
	if err := unix.Mount("tmpfs", "/mnt", "tmpfs", 0, ""); err != nil {
		return fmt.Errorf("could not mount tmpfs on /mnt: %w", err)
	}

	// Make all directories. Since filepath.Walk is lexicographically ordered, we don't need to ensure that the parent
	// exists.
	for _, src := range dirs {
		stat, err := os.Stat(src)
		if err != nil {
			return fmt.Errorf("Stat(%q): %w", src, err)
		}
		dst := "/mnt" + src
		err = os.Mkdir(dst, stat.Mode())
		if err != nil {
			return fmt.Errorf("Mkdir(%q): %w", dst, err)
		}
	}

	// Move all files over. Parent directories will exist by now.
	for _, src := range paths {
		stat, err := os.Stat(src)
		if err != nil {
			return fmt.Errorf("Stat(%q): %w", src, err)
		}
		dst := "/mnt" + src

		// Copy file.
		sfd, err := os.Open(src)
		if err != nil {
			return fmt.Errorf("Open(%q): %w", src, err)
		}
		dfd, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, stat.Mode())
		if err != nil {
			sfd.Close()
			return fmt.Errorf("OpenFile(%q): %w", dst, err)
		}
		_, err = io.Copy(dfd, sfd)

		sfd.Close()
		dfd.Close()
		if err != nil {
			return fmt.Errorf("Copying %q failed: %w", src, err)
		}

		// Remove the old file.
		err = unix.Unlink(src)
		if err != nil {
			return fmt.Errorf("Unlink(%q): %w", src, err)
		}
	}

	// Set up target filesystems.
	for _, el := range []struct {
		dir   string
		fs    string
		flags uintptr
	}{
		{"/sys", "sysfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/proc", "proc", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/dev", "devtmpfs", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/dev/pts", "devpts", unix.MS_NOEXEC | unix.MS_NOSUID},
	} {
		if err := os.Mkdir("/mnt"+el.dir, 0755); err != nil {
			return fmt.Errorf("could not make /mnt%s: %w", el.dir, err)
		}
		if err := unix.Mount(el.fs, "/mnt"+el.dir, el.fs, el.flags, ""); err != nil {
			return fmt.Errorf("could not mount %s on /mnt%s: %w", el.fs, el.dir, err)
		}
	}

	// Mount all available CGroups for v1 (v2 uses a single unified hierarchy and is not supported by our runtimes yet)
	if unix.Mount("tmpfs", "/mnt/sys/fs/cgroup", "tmpfs", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV, ""); err != nil {
		panic(err)
	}
	cgroupsRaw, err := ioutil.ReadFile("/mnt/proc/cgroups")
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
		if err := os.Mkdir("/mnt/sys/fs/cgroup/"+cgroupName, 0755); err != nil {
			panic(err)
		}
		if err := unix.Mount("cgroup", "/mnt/sys/fs/cgroup/"+cgroupName, "cgroup", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV, cgroupName); err != nil {
			panic(err)
		}
	}

	// Enable hierarchical memory accounting
	useMemoryHierarchy, err := os.OpenFile("/mnt/sys/fs/cgroup/memory/memory.use_hierarchy", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	if _, err := useMemoryHierarchy.WriteString("1"); err != nil {
		panic(err)
	}
	useMemoryHierarchy.Close()

	// Chroot to new root.
	// This is adapted from util-linux's switch_root.
	err = os.Chdir("/mnt")
	if err != nil {
		return fmt.Errorf("could not chdir to /mnt: %w", err)
	}
	err = syscall.Mount("/mnt", "/", "", syscall.MS_MOVE, "")
	if err != nil {
		return fmt.Errorf("could not remount /mnt to /: %w", err)
	}
	err = syscall.Chroot(".")
	if err != nil {
		return fmt.Errorf("could not chroot to new root: %w", err)
	}

	// Re-exec into new init with new environment
	return unix.Exec("/init", os.Args, environ)
}
