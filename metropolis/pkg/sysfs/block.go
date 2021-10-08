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

// Implementation included in this file was written with the aim of easing
// integration with the interface exposed at /sys/class/block. It assumes sysfs
// is already mounted at /sys.
package sysfs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// PartUUIDMap returns a mapping between partition UUIDs and block device
// names based on information exposed by uevent. UUID keys of the returned
// map are represented as lowercase strings.
func PartUUIDMap() (map[string]string, error) {
	m := make(map[string]string)
	// Get a list of block device symlinks from sysfs.
	const blkDirPath = "/sys/class/block"
	blkDevs, err := os.ReadDir(blkDirPath)
	if err != nil {
		return m, fmt.Errorf("couldn't read %q: %w", blkDirPath, err)
	}
	// Iterate over block device symlinks present in blkDevs, creating a mapping
	// in m for each device with both PARTUUID and DEVNAME keys present in their
	// respective uevent files.
	for _, devInfo := range blkDevs {
		// Read the uevent file and transform it into a string->string map.
		kv, err := ReadUevents(filepath.Join(blkDirPath, devInfo.Name(), "uevent"))
		if err != nil {
			return m, fmt.Errorf("while reading uevents: %w", err)
		}
		// Check that the required keys are present in the map.
		if uuid, name := kv["PARTUUID"], kv["DEVNAME"]; uuid != "" && name != "" {
			m[uuid] = name
		}
	}
	return m, nil
}

// DeviceByPartUUID returns a block device name, given its corresponding
// partition UUID.
func DeviceByPartUUID(uuid string) (string, error) {
	pm, err := PartUUIDMap()
	if err != nil {
		return "", err
	}
	if bdev, ok := pm[strings.ToLower(uuid)]; ok {
		return bdev, nil
	}
	return "", fmt.Errorf("couldn't find a block device matching the partition UUID %q", uuid)
}

// ParentBlockDevice transforms the block device name of a partition, eg
// "sda1", to the name of the block device hosting it, eg "sda".
func ParentBlockDevice(dev string) (string, error) {
	// Build a path pointing to a sysfs block device symlink.
	partLink := filepath.Join("/sys/class/block", dev)
	// Read the symlink at partLink. This should leave us with a path of the form
	// (...)/sda/sdaN.
	linkTgt, err := os.Readlink(partLink)
	if err != nil {
		return "", fmt.Errorf("couldn't read the block device symlink at %q: %w", partLink, err)
	}
	// Remove the last element from the path, leaving us with a path pointing to
	// the block device containting the installer partition, of the form
	// (...)/sda.
	devPath := filepath.Dir(linkTgt)
	// Get the last element of the path, leaving us with just the block device
	// name, eg sda
	devName := filepath.Base(devPath)
	return devName, nil
}

// PartitionBlockDevice returns the name of a block device associated with the
// partition at index in the containing block device dev, eg "nvme0n1pN" for
// "nvme0n1" or "sdaN" for "sda".
func PartitionBlockDevice(dev string, index int) (string, error) {
	dp := filepath.Join("/sys/class/block", dev)
	dir, err := os.ReadDir(dp)
	if err != nil {
		return "", err
	}
	for _, info := range dir {
		// Skip non-directories
		if !info.IsDir() {
			continue
		}
		// Check whether the directory contains a file named 'partition'. If that's
		// the case, read the partition index from it and compare it with the one
		// supplied as a function parameter. If they're equal, return the directory
		// name.
		istr, err := os.ReadFile(filepath.Join(dp, info.Name(), "partition"))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", err
		}
		// istr holds a newline-terminated ASCII-encoded decimal number.
		pi, err := strconv.Atoi(strings.TrimSuffix(string(istr), "\n"))
		if err != nil {
			return "", fmt.Errorf("failed to parse partition index: %w", err)
		}
		if pi == index {
			return info.Name(), nil
		}
	}
	return "", fmt.Errorf("couldn't find partition %d of %q", index, dev)
}
