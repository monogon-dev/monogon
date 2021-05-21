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
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rekby/gpt"
	"golang.org/x/sys/unix"

	"source.monogon.dev/metropolis/pkg/sysfs"
)

var (
	// EFIPartitionType is the standardized partition type value for the EFI
	// ESP partition. The human readable GUID is
	// C12A7328-F81F-11D2-BA4B-00A0C93EC93B.
	EFIPartitionType = gpt.PartType{0x28, 0x73, 0x2a, 0xc1, 0x1f, 0xf8, 0xd2, 0x11, 0xba, 0x4b, 0x00, 0xa0, 0xc9, 0x3e, 0xc9, 0x3b}

	// NodeDataPartitionType is the partition type value for a Metropolis Node
	// data partition. The human-readable GUID is
	// 9eeec464-6885-414a-b278-4305c51f7966.
	NodeDataPartitionType = gpt.PartType{0x64, 0xc4, 0xee, 0x9e, 0x85, 0x68, 0x4a, 0x41, 0xb2, 0x78, 0x43, 0x05, 0xc5, 0x1f, 0x79, 0x66}
)

const (
	ESPDevicePath     = "/dev/esp"
	NodeDataCryptPath = "/dev/data-crypt"
)

// MakeBlockDevices looks for the ESP and the node data partition and maps them
// to ESPDevicePath and NodeDataCryptPath respectively. This doesn't fail if it
// doesn't find the partitions, only if something goes catastrophically wrong.
func MakeBlockDevices(ctx context.Context) error {
	blockdevNames, err := ioutil.ReadDir("/sys/class/block")
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
			blkdev, err := os.Open(devNodeName)
			if err != nil {
				return fmt.Errorf("failed to open block device %v: %w", devNodeName, err)
			}
			defer blkdev.Close()
			blockSize, err := unix.IoctlGetUint32(int(blkdev.Fd()), unix.BLKSSZGET)
			if err != nil {
				continue // This is not a regular block device
			}
			blkdev.Seek(int64(blockSize), 0)
			table, err := gpt.ReadTable(blkdev, uint64(blockSize))
			if err != nil {
				// Probably just not a GPT-partitioned disk
				continue
			}
			for partNumber, part := range table.Partitions {
				if part.Type == EFIPartitionType {
					err := unix.Mknod(ESPDevicePath, 0600|unix.S_IFBLK, int(unix.Mkdev(uint32(majorDev), uint32(partNumber+1))))
					if err != nil && !os.IsExist(err) {
						return fmt.Errorf("failed to create device node for ESP partition: %w", err)
					}
				}
				if part.Type == NodeDataPartitionType {
					err := unix.Mknod(NodeDataCryptPath, 0600|unix.S_IFBLK, int(unix.Mkdev(uint32(majorDev), uint32(partNumber+1))))
					if err != nil && !os.IsExist(err) {
						return fmt.Errorf("failed to create device node for Metropolis node encrypted data partition: %w", err)
					}
				}
			}
		}
	}
	return nil
}
