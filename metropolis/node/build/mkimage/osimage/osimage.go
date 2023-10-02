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

// This package provides self-contained implementation used to generate
// Metropolis disk images.
package osimage

import (
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"

	"source.monogon.dev/metropolis/pkg/blockdev"
	"source.monogon.dev/metropolis/pkg/efivarfs"
	"source.monogon.dev/metropolis/pkg/fat32"
	"source.monogon.dev/metropolis/pkg/gpt"
)

var (
	SystemAType = uuid.MustParse("ee96054b-f6d0-4267-aaaa-724b2afea74c")
	SystemBType = uuid.MustParse("ee96054b-f6d0-4267-bbbb-724b2afea74c")

	DataType = uuid.MustParse("9eeec464-6885-414a-b278-4305c51f7966")
)

const (
	SystemALabel = "METROPOLIS-SYSTEM-A"
	SystemBLabel = "METROPOLIS-SYSTEM-B"
	DataLabel    = "METROPOLIS-NODE-DATA"
	ESPLabel     = "ESP"

	EFIPayloadPath = "/EFI/BOOT/BOOTx64.EFI"
	EFIBootAPath   = "/EFI/metropolis/boot-a.efi"
	EFIBootBPath   = "/EFI/metropolis/boot-b.efi"
	nodeParamsPath = "metropolis/parameters.pb"
)

// PartitionSizeInfo contains parameters used during partition table
// initialization and, in case of image files, space allocation.
type PartitionSizeInfo struct {
	// Size of the EFI System Partition (ESP), in mebibytes. The size must
	// not be zero.
	ESP int64
	// Size of the Metropolis system partition, in mebibytes. The partition
	// won't be created if the size is zero.
	System int64
	// Size of the Metropolis data partition, in mebibytes. The partition
	// won't be created if the size is zero. If the image is output to a
	// block device, the partition will be extended to fill the remaining
	// space.
	Data int64
}

// Params contains parameters used by Create to build a Metropolis OS
// image.
type Params struct {
	// Output is the block device to which the OS image is written.
	Output blockdev.BlockDev
	// ABLoader provides the A/B loader which then loads the EFI loader for the
	// correct slot.
	ABLoader fat32.SizedReader
	// EFIPayload provides contents of the EFI payload file. It must not be
	// nil. This gets put into boot slot A.
	EFIPayload fat32.SizedReader
	// SystemImage provides contents of the Metropolis system partition.
	// If nil, no contents will be copied into the partition.
	SystemImage io.Reader
	// NodeParameters provides contents of the node parameters file. If nil,
	// the node parameters file won't be created in the target ESP
	// filesystem.
	NodeParameters fat32.SizedReader
	// DiskGUID is a unique identifier of the image and a part of Table
	// header. It's optional and can be left blank if the identifier is
	// to be randomly generated. Setting it to a predetermined value can
	// help in implementing reproducible builds.
	DiskGUID uuid.UUID
	// PartitionSize specifies a size for the ESP, Metropolis System and
	// Metropolis data partition.
	PartitionSize PartitionSizeInfo
}

const Mi = 1024 * 1024

// Create writes a Metropolis OS image to a block device.
func Create(params *Params) (*efivarfs.LoadOption, error) {
	// Discard the entire device, we're going to write new data over it.
	// Ignore errors, this is only advisory.
	params.Output.Discard(0, params.Output.BlockCount())

	tbl, err := gpt.New(params.Output)
	if err != nil {
		return nil, fmt.Errorf("invalid block device: %w", err)
	}
	tbl.ID = params.DiskGUID
	esp := gpt.Partition{
		Type: gpt.PartitionTypeEFISystem,
		Name: ESPLabel,
	}
	if err := tbl.AddPartition(&esp, params.PartitionSize.ESP*Mi); err != nil {
		return nil, fmt.Errorf("failed to allocate ESP: %w", err)
	}

	rootInode := fat32.Inode{
		Attrs: fat32.AttrDirectory,
	}
	if err := rootInode.PlaceFile(strings.TrimPrefix(EFIBootAPath, "/"), params.EFIPayload); err != nil {
		return nil, err
	}
	// Place the A/B loader at the EFI bootloader autodiscovery path.
	if err := rootInode.PlaceFile(strings.TrimPrefix(EFIPayloadPath, "/"), params.ABLoader); err != nil {
		return nil, err
	}
	if params.NodeParameters != nil {
		if err := rootInode.PlaceFile(nodeParamsPath, params.NodeParameters); err != nil {
			return nil, err
		}
	}
	if err := fat32.WriteFS(blockdev.NewRWS(esp), rootInode, fat32.Options{
		BlockSize:  uint16(esp.BlockSize()),
		BlockCount: uint32(esp.BlockCount()),
		Label:      "MNGN_BOOT",
	}); err != nil {
		return nil, fmt.Errorf("failed to write FAT32: %w", err)
	}

	// Create the system partition only if its size is specified.
	if params.PartitionSize.System != 0 && params.SystemImage != nil {
		systemPartitionA := gpt.Partition{
			Type: SystemAType,
			Name: SystemALabel,
		}
		if err := tbl.AddPartition(&systemPartitionA, params.PartitionSize.System*Mi); err != nil {
			return nil, fmt.Errorf("failed to allocate system partition A: %w", err)
		}
		if _, err := io.Copy(blockdev.NewRWS(systemPartitionA), params.SystemImage); err != nil {
			return nil, fmt.Errorf("failed to write system partition A: %w", err)
		}
		systemPartitionB := gpt.Partition{
			Type: SystemBType,
			Name: SystemBLabel,
		}
		if err := tbl.AddPartition(&systemPartitionB, params.PartitionSize.System*Mi); err != nil {
			return nil, fmt.Errorf("failed to allocate system partition B: %w", err)
		}
	} else if params.PartitionSize.System == 0 && params.SystemImage != nil {
		// Safeguard against contradicting parameters.
		return nil, fmt.Errorf("the system image parameter was passed while the associated partition size is zero")
	}
	// Create the data partition only if its size is specified.
	if params.PartitionSize.Data != 0 {
		dataPartition := gpt.Partition{
			Type: DataType,
			Name: DataLabel,
		}
		if err := tbl.AddPartition(&dataPartition, -1); err != nil {
			return nil, fmt.Errorf("failed to allocate data partition: %w", err)
		}
	}

	if err := tbl.Write(); err != nil {
		return nil, fmt.Errorf("failed to write Table: %w", err)
	}

	// Build an EFI boot entry pointing to the image's ESP.
	return &efivarfs.LoadOption{
		Description: "Metropolis",
		FilePath: efivarfs.DevicePath{
			&efivarfs.HardDrivePath{
				PartitionNumber:     1,
				PartitionStartBlock: esp.FirstBlock,
				PartitionSizeBlocks: esp.SizeBlocks(),
				PartitionMatch: efivarfs.PartitionGPT{
					PartitionUUID: esp.ID,
				},
			},
			efivarfs.FilePath(EFIPayloadPath),
		},
	}, nil
}
