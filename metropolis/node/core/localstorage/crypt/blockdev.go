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
	"unsafe"

	"github.com/google/uuid"
	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/pkg/efivarfs"
	"source.monogon.dev/metropolis/pkg/gpt"
	"source.monogon.dev/metropolis/pkg/supervisor"
	"source.monogon.dev/metropolis/pkg/sysfs"
)

// NodeDataPartitionType is the partition type value for a Metropolis Node
// data partition.
var NodeDataPartitionType = uuid.MustParse("9eeec464-6885-414a-b278-4305c51f7966")

const (
	ESPDevicePath   = "/dev/esp"
	NodeDataRawPath = "/dev/data-raw"
)

// MakeBlockDevices looks for the ESP and the node data partition and maps them
// to ESPDevicePath and NodeDataCryptPath respectively. This doesn't fail if it
// doesn't find the partitions, only if something goes catastrophically wrong.
func MakeBlockDevices(ctx context.Context) error {
	espUUID, err := efivarfs.ReadLoaderDevicePartUUID()
	if err != nil {
		supervisor.Logger(ctx).Warningf("No EFI variable for the loader device partition UUID present")
	}
	blockdevNames, err := os.ReadDir("/sys/class/block")
	if err != nil {
		return fmt.Errorf("failed to read sysfs block class: %w", err)
	}
	for _, blockdevName := range blockdevNames {
		ueventData, err := sysfs.ReadUevents(filepath.Join("/sys/class/block", blockdevName.Name(), "uevent"))
		if err != nil {
			return fmt.Errorf("failed to read uevent for block device %v: %w", blockdevName.Name(), err)
		}
		if ueventData["DEVTYPE"] == "disk" {
			majorDev, err := strconv.Atoi(ueventData["MAJOR"])
			if err != nil {
				return fmt.Errorf("failed to convert uevent: %w", err)
			}
			devNodeName := fmt.Sprintf("/dev/%v", ueventData["DEVNAME"])
			// TODO(lorenz): This extraction code is all a bit hairy, will get
			// replaced by blockdev shortly.
			blkdev, err := os.Open(devNodeName)
			if err != nil {
				return fmt.Errorf("failed to open block device %v: %w", devNodeName, err)
			}
			defer blkdev.Close()
			blockSize, err := unix.IoctlGetUint32(int(blkdev.Fd()), unix.BLKSSZGET)
			if err != nil {
				continue // This is not a regular block device
			}
			var sizeBytes uint64
			_, _, err = unix.Syscall(unix.SYS_IOCTL, blkdev.Fd(), unix.BLKGETSIZE64, uintptr(unsafe.Pointer(&sizeBytes)))
			if err != unix.Errno(0) {
				return fmt.Errorf("failed to get device size: %w", err)
			}
			blkdev.Seek(int64(blockSize), 0)
			table, err := gpt.Read(blkdev, int64(blockSize), int64(sizeBytes)/int64(blockSize))
			if err != nil {
				// Probably just not a GPT-partitioned disk
				continue
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
				continue
			}
			seenTypes := make(map[uuid.UUID]bool)
			for partNumber, part := range table.Partitions {
				if seenTypes[part.Type] {
					return fmt.Errorf("failed to create device node for %s (%s): node for this type already created/multiple partitions found", part.ID.String(), part.Type.String())
				}
				if part.Type == gpt.PartitionTypeEFISystem {
					seenTypes[part.Type] = true
					err := unix.Mknod(ESPDevicePath, 0600|unix.S_IFBLK, int(unix.Mkdev(uint32(majorDev), uint32(partNumber+1))))
					if err != nil && !os.IsExist(err) {
						return fmt.Errorf("failed to create device node for ESP partition: %w", err)
					}
				}
				if part.Type == NodeDataPartitionType {
					seenTypes[part.Type] = true
					err := unix.Mknod(NodeDataRawPath, 0600|unix.S_IFBLK, int(unix.Mkdev(uint32(majorDev), uint32(partNumber+1))))
					if err != nil && !os.IsExist(err) {
						return fmt.Errorf("failed to create device node for Metropolis node encrypted data partition: %w", err)
					}
				}
			}
		}
	}
	return nil
}
