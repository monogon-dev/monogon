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

package crypt

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/node/core/update"
	"source.monogon.dev/metropolis/pkg/blockdev"
	"source.monogon.dev/metropolis/pkg/efivarfs"
	"source.monogon.dev/metropolis/pkg/gpt"
	"source.monogon.dev/metropolis/pkg/supervisor"
	"source.monogon.dev/metropolis/pkg/sysfs"
)

// NodeDataPartitionType is the partition type value for a Metropolis Node
// data partition.
var NodeDataPartitionType = uuid.MustParse("9eeec464-6885-414a-b278-4305c51f7966")

var (
	SystemAType = uuid.MustParse("ee96054b-f6d0-4267-aaaa-724b2afea74c")
	SystemBType = uuid.MustParse("ee96054b-f6d0-4267-bbbb-724b2afea74c")
)

const (
	ESPDevicePath     = "/dev/esp"
	NodeDataRawPath   = "/dev/data-raw"
	SystemADevicePath = "/dev/system-a"
	SystemBDevicePath = "/dev/system-b"
)

// nodePathForPartitionType returns the device node path
// for a given partition type.
func nodePathForPartitionType(t uuid.UUID) string {
	switch t {
	case gpt.PartitionTypeEFISystem:
		return ESPDevicePath
	case NodeDataPartitionType:
		return NodeDataRawPath
	case SystemAType:
		return SystemADevicePath
	case SystemBType:
		return SystemBDevicePath
	}
	return ""
}

// MakeBlockDevices looks for the ESP and the node data partition and maps them
// to ESPDevicePath and NodeDataCryptPath respectively. This doesn't fail if it
// doesn't find the partitions, only if something goes catastrophically wrong.
func MakeBlockDevices(ctx context.Context, updateSvc *update.Service) error {
	espUUID, err := efivarfs.ReadLoaderDevicePartUUID()
	if err != nil {
		supervisor.Logger(ctx).Warningf("No EFI variable for the loader device partition UUID present")
	}

	blockDevs, err := os.ReadDir("/sys/class/block")
	if err != nil {
		return fmt.Errorf("failed to read sysfs block class: %w", err)
	}

	for _, blockDev := range blockDevs {
		if err := handleBlockDevice(blockDev.Name(), blockDevs, espUUID, updateSvc); err != nil {
			supervisor.Logger(ctx).Errorf("Failed to create block device %s: %w", blockDev.Name(), err)
		}
	}

	return nil
}

// handleBlockDevice reads the uevent data and continues to iterate over all
// partitions to create all required device nodes.
func handleBlockDevice(diskBlockDev string, blockDevs []os.DirEntry, espUUID uuid.UUID, updateSvc *update.Service) error {
	data, err := readUEvent(diskBlockDev)
	if err != nil {
		return err
	}

	// We only care about disks, skip all other dev types.
	if data["DEVTYPE"] != "disk" {
		return nil
	}

	blkdev, err := blockdev.Open(fmt.Sprintf("/dev/%v", data["DEVNAME"]))
	if err != nil {
		return fmt.Errorf("failed to open block device: %w", err)
	}
	defer blkdev.Close()

	table, err := gpt.Read(blkdev)
	if err != nil {
		return nil // Probably just not a GPT-partitioned disk
	}

	skipDisk := false
	if espUUID != uuid.Nil {
		// If we know where we booted from, ignore all disks which do
		// not contain this partition.
		skipDisk = true
		for _, part := range table.Partitions {
			if part.ID == espUUID {
				skipDisk = false
				break
			}
		}
	}
	if skipDisk {
		return nil
	}

	seenTypes := make(map[uuid.UUID]bool)
	for _, dev := range blockDevs {
		if err := handlePartition(diskBlockDev, dev.Name(), table, seenTypes, updateSvc); err != nil {
			return fmt.Errorf("when creating partition %s: %w", dev.Name(), err)
		}
	}

	return nil
}

func handlePartition(diskBlockDev string, partBlockDev string, table *gpt.Table, seenTypes map[uuid.UUID]bool, updateSvc *update.Service) error {
	// Skip all blockdev that dont share the same name/prefix,
	// also skip the blockdev itself.
	if !strings.HasPrefix(partBlockDev, diskBlockDev) || partBlockDev == diskBlockDev {
		return nil
	}

	data, err := readUEvent(partBlockDev)
	if err != nil {
		return err
	}

	// We only care about partitions, skip all other dev types.
	if data["DEVTYPE"] != "partition" {
		return nil
	}

	pi, err := data.readPartitionInfo()
	if err != nil {
		return err
	}

	part := table.Partitions[pi.partNumber-1]

	if part.Type == gpt.PartitionTypeEFISystem {
		updateSvc.ProvideESP("/esp", uint32(pi.partNumber), part)
	}

	nodePath := nodePathForPartitionType(part.Type)
	if nodePath == "" {
		// Ignore partitions with an unknown type.
		return nil
	}

	if seenTypes[part.Type] {
		return fmt.Errorf("node for this type (%s) already created/multiple partitions found", part.Type.String())
	}
	seenTypes[part.Type] = true

	if err := pi.makeDeviceNode(nodePath); err != nil {
		return fmt.Errorf("when creating partition node: %w", err)
	}

	return nil
}

type partInfo struct {
	major, minor, partNumber int
}

// validateDeviceNode tries to open a device node and validates that it
// has the expected major and minor device numbers. If the path does non exist,
// no error will be returned.
func (pi partInfo) validateDeviceNode(path string) error {
	var s unix.Stat_t
	if err := unix.Stat(path, &s); err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("inspecting device node %q: %w", path, err)
	}

	if unix.Major(s.Rdev) != uint32(pi.major) || unix.Minor(s.Rdev) != uint32(pi.minor) {
		return fmt.Errorf("device node %q exists for different device %d:%d", path, unix.Major(s.Rdev), unix.Minor(s.Rdev))
	}

	return nil
}

// makeDeviceNode creates the device node at the given path based on the
// major and minor device number. If the device node already exists and points
// to the same device, no error will be returned.
func (pi partInfo) makeDeviceNode(path string) error {
	if err := pi.validateDeviceNode(path); err != nil {
		return err
	}

	err := unix.Mknod(path, 0600|unix.S_IFBLK, int(unix.Mkdev(uint32(pi.major), uint32(pi.minor))))
	if err != nil {
		return fmt.Errorf("create device node %q: %w", path, err)
	}
	return nil
}

func readUEvent(blockName string) (blockUEvent, error) {
	data, err := sysfs.ReadUevents(filepath.Join("/sys/class/block", blockName, "uevent"))
	if err != nil {
		return nil, fmt.Errorf("when reading uevent: %w", err)
	}
	return data, nil
}

type blockUEvent map[string]string

func (b blockUEvent) readUdevKeyInteger(key string) (int, error) {
	if _, ok := b[key]; !ok {
		return 0, fmt.Errorf("missing udev value %s", key)
	}

	v, err := strconv.Atoi(b[key])
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %w", key, err)
	}

	return v, nil
}

// readPartitionInfo parses all fields for partInfo from a blockUEvent.
func (b blockUEvent) readPartitionInfo() (pi partInfo, err error) {
	pi.major, err = b.readUdevKeyInteger("MAJOR")
	if err != nil {
		return
	}

	pi.minor, err = b.readUdevKeyInteger("MINOR")
	if err != nil {
		return
	}

	pi.partNumber, err = b.readUdevKeyInteger("PARTN")
	if err != nil {
		return
	}

	return
}
