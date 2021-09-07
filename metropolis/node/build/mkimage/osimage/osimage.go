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
	"io/ioutil"
	"os"

	diskfs "github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
	"github.com/diskfs/go-diskfs/partition/gpt"

	"source.monogon.dev/metropolis/pkg/efivarfs"
)

const (
	systemPartitionType = gpt.Type("ee96055b-f6d0-4267-8bbb-724b2afea74c")
	SystemVolumeLabel   = "METROPOLIS-SYSTEM"

	dataPartitionType = gpt.Type("9eeec464-6885-414a-b278-4305c51f7966")
	DataVolumeLabel   = "METROPOLIS-NODE-DATA"

	ESPVolumeLabel = "ESP"

	EFIPayloadPath = "/EFI/BOOT/BOOTx64.EFI"
	nodeParamsPath = "/EFI/metropolis/parameters.pb"

	mib = 1024 * 1024
)

// put creates a file on the target filesystem fs and fills it with
// contents read from an io.Reader object src.
func put(fs filesystem.FileSystem, dst string, src io.Reader) error {
	target, err := fs.OpenFile(dst, os.O_CREATE|os.O_RDWR)
	if err != nil {
		return fmt.Errorf("while opening %q: %w", dst, err)
	}

	// If this is streamed (e.g. using io.Copy) it exposes a bug in diskfs, so
	// do it in one go.
	// TODO(mateusz@monogon.tech): Investigate the bug.
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return fmt.Errorf("while reading %q: %w", src, err)
	}
	if _, err := target.Write(data); err != nil {
		return fmt.Errorf("while writing to %q: %w", dst, err)
	}
	return nil
}

// initializeESP creates an ESP filesystem in a partition specified by
// index. It then creates the EFI executable and copies into it contents
// of the reader object exec, which must not be nil. The node parameters
// file is optionally created if params is not nil. initializeESP may return
// an error.
func initializeESP(image *disk.Disk, index int, exec, params io.Reader) error {
	// Create a FAT ESP filesystem inside a partition pointed to by
	// index.
	spec := disk.FilesystemSpec{
		Partition:   index,
		FSType:      filesystem.TypeFat32,
		VolumeLabel: ESPVolumeLabel,
	}
	fs, err := image.CreateFilesystem(spec)
	if err != nil {
		return fmt.Errorf("while creating an ESP filesystem: %w", err)
	}

	// Create the EFI partition structure.
	for _, dir := range []string{"/EFI", "/EFI/BOOT", "/EFI/metropolis"} {
		if err := fs.Mkdir(dir); err != nil {
			return fmt.Errorf("while creating %q: %w", dir, err)
		}
	}

	// Copy the EFI payload to the newly created filesystem.
	if exec == nil {
		return fmt.Errorf("exec must not be nil")
	}
	if err := put(fs, EFIPayloadPath, exec); err != nil {
		return fmt.Errorf("while writing an EFI payload: %w", err)
	}

	if params != nil {
		// Copy Node Parameters into the ESP.
		if err := put(fs, nodeParamsPath, params); err != nil {
			return fmt.Errorf("while writing node parameters: %w", err)
		}
	}
	return nil
}

// zeroSrc is a source of null bytes implementing io.Reader. It acts as a
// helper data type.
type zeroSrc struct{}

// Read implements io.Reader for zeroSrc. It fills b with zero bytes. The
// returned error is always nil.
func (_ zeroSrc) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

// initializeSystemPartition copies system partition contents into a partition
// at index. The remaining partition space is zero-padded. This function may
// return an error.
func initializeSystemPartition(image *disk.Disk, index int, contents io.Reader) error {
	// Check the parameters.
	if contents == nil {
		return fmt.Errorf("system partition contents parameter must not be nil")
	}
	if index <= 0 {
		return fmt.Errorf("partition index must be greater than zero")
	}

	// Get the system partition's size.
	table, err := image.GetPartitionTable()
	if err != nil {
		return fmt.Errorf("while accessing a go-diskfs partition table: %w", err)
	}
	partitions := table.GetPartitions()
	if index > len(partitions) {
		return fmt.Errorf("partition index out of bounds")
	}
	size := partitions[index-1].GetSize()

	// Copy the contents of the Metropolis system image into the system partition
	// at partitionIndex. Zero-pad the remaining space.
	var zero zeroSrc
	sys := io.LimitReader(io.MultiReader(contents, zero), size)
	if _, err := image.WritePartitionContents(index, sys); err != nil {
		return fmt.Errorf("while copying the system partition: %w", err)
	}
	return nil
}

// mibToSectors converts a size expressed in mebibytes to a number of
// sectors needed to store data of that size. sectorSize parameter
// specifies the size of a logical sector.
func mibToSectors(size, sectorSize uint64) uint64 {
	return (size * mib) / sectorSize
}

// PartitionSizeInfo contains parameters used during partition table
// initialization and, in case of image files, space allocation.
type PartitionSizeInfo struct {
	// Size of the EFI System Partition (ESP), in mebibytes. The size must
	// not be zero.
	ESP uint64
	// Size of the Metropolis system partition, in mebibytes. The partition
	// won't be created if the size is zero.
	System uint64
	// Size of the Metropolis data partition, in mebibytes. The partition
	// won't be created if the size is zero. If the image is output to a
	// block device, the partition will be extended to fill the remaining
	// space.
	Data uint64
}

// partitionList stores partition definitions in an ascending order.
type partitionList []*gpt.Partition

// appendPartition puts a new partition at the end of a partitionList,
// automatically calculating its start and end markers based on data in
// the list and the argument psize. A partition type and a name are
// assigned to the partition. The containing image is used to calculate
// sector addresses based on its logical block size.
func (pl *partitionList) appendPartition(image *disk.Disk, ptype gpt.Type, pname string, psize uint64) {
	// Calculate the start and end marker.
	var pstart, pend uint64
	if len(*pl) != 0 {
		pstart = (*pl)[len(*pl)-1].End + 1
	} else {
		pstart = mibToSectors(1, uint64(image.LogicalBlocksize))
	}
	pend = pstart + mibToSectors(psize, uint64(image.LogicalBlocksize)) - 1

	// Put the new partition at the end of the list.
	*pl = append(*pl, &gpt.Partition{
		Type:  ptype,
		Name:  pname,
		Start: pstart,
		End:   pend,
	})
}

// extendLastPartition moves the end marker of the last partition in a
// partitionList to the end of image, assigning all of the remaining free
// space to it. It may return an error.
func (pl *partitionList) extendLastPartition(image *disk.Disk) error {
	if len(*pl) == 0 {
		return fmt.Errorf("no partitions defined")
	}
	if image.Size == 0 {
		return fmt.Errorf("the image size mustn't be zero")
	}
	if image.LogicalBlocksize == 0 {
		return fmt.Errorf("the image's logical block size mustn't be zero")
	}

	// The last 33 blocks are occupied by the Secondary GPT.
	(*pl)[len(*pl)-1].End = uint64(image.Size/image.LogicalBlocksize) - 33
	return nil
}

// initializePartitionTable applies a Metropolis-compatible GPT partition
// table to an image. Logical and physical sector sizes are based on
// block sizes exposed by Disk. Partition extents are defined by the size
// argument. A diskGUID is associated with the partition table. In an event
// the table couldn't have been applied, the function will return an error.
func initializePartitionTable(image *disk.Disk, size *PartitionSizeInfo, diskGUID string) error {
	// Start with preparing a partition list.
	var pl partitionList
	// Create the ESP.
	if size.ESP == 0 {
		return fmt.Errorf("ESP size mustn't be zero")
	}
	pl.appendPartition(image, gpt.EFISystemPartition, ESPVolumeLabel, size.ESP)
	// Create the system partition only if its size is specified.
	if size.System != 0 {
		pl.appendPartition(image, systemPartitionType, SystemVolumeLabel, size.System)
	}
	// Create the data partition only if its size is specified.
	if size.Data != 0 {
		// Don't specify the last partition's length, as it will be extended
		// to fit the image size anyway. size.Data will still be used in the
		// space allocation step.
		pl.appendPartition(image, dataPartitionType, DataVolumeLabel, 0)
		if err := pl.extendLastPartition(image); err != nil {
			return fmt.Errorf("while extending the last partition: %w", err)
		}
	}

	// Build and apply the partition table.
	table := &gpt.Table{
		LogicalSectorSize:  int(image.LogicalBlocksize),
		PhysicalSectorSize: int(image.PhysicalBlocksize),
		ProtectiveMBR:      true,
		GUID:               diskGUID,
		Partitions:         pl,
	}
	if err := image.Partition(table); err != nil {
		// Return the error unwrapped.
		return err
	}
	return nil
}

// Params contains parameters used by Create to build a Metropolis OS
// image.
type Params struct {
	// OutputPath is the path an OS image will be written to. If the path
	// points to an existing block device, the data partition will be
	// extended to fill it entirely. Otherwise, a regular image file will
	// be created at OutputPath. The path must not point to an existing
	// regular file.
	OutputPath string
	// EFIPayload provides contents of the EFI payload file. It must not be
	// nil.
	EFIPayload io.Reader
	// SystemImage provides contents of the Metropolis system partition.
	// If nil, no contents will be copied into the partition.
	SystemImage io.Reader
	// NodeParameters provides contents of the node parameters file. If nil,
	// the node parameters file won't be created in the target ESP
	// filesystem.
	NodeParameters io.Reader
	// DiskGUID is a unique identifier of the image and a part of GPT
	// header. It's optional and can be left blank if the identifier is
	// to be randomly generated. Setting it to a predetermined value can
	// help in implementing reproducible builds.
	DiskGUID string
	// PartitionSize specifies a size for the ESP, Metropolis System and
	// Metropolis data partition.
	PartitionSize PartitionSizeInfo
}

// Create writes a Metropolis OS image to either a newly created regular
// file or a block device. The image is parametrized by the params
// argument. In case a regular file already exists at params.OutputPath,
// the function will fail. It returns nil on success or an error, if one
// did occur.
func Create(params *Params) (*efivarfs.BootEntry, error) {
	// Validate each parameter before use.
	if params.OutputPath == "" {
		return nil, fmt.Errorf("image output path must be set")
	}

	// Learn whether we're creating a new image or writing to an existing
	// block device by stat-ing the output path parameter.
	outInfo, err := os.Stat(params.OutputPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	// Calculate the image size (bytes) by summing up partition sizes
	// (mebibytes).
	minSize := (int64(params.PartitionSize.ESP) +
		int64(params.PartitionSize.System) +
		int64(params.PartitionSize.Data) + 1) * mib
	var diskImg *disk.Disk
	if !os.IsNotExist(err) && outInfo.Mode()&os.ModeDevice != 0 {
		// Open the block device. The data partition size parameter won't
		// matter in this case, as said partition will be extended till the
		// end of device.
		diskImg, err = diskfs.Open(params.OutputPath)
		if err != nil {
			return nil, fmt.Errorf("couldn't open the block device at %q: %w", params.OutputPath, err)
		}
		// Make sure there's at least minSize space available on the block
		// device.
		if minSize > diskImg.Size {
			return nil, fmt.Errorf("not enough space available on the block device at %q", params.OutputPath)
		}
	} else {
		// Attempt to create an image file at params.OutputPath. diskfs.Create
		// will abort in case a file already exists at the given path.
		// Calculate the image size expressed in bytes by summing up declared
		// partition lengths.
		diskImg, err = diskfs.Create(params.OutputPath, minSize, diskfs.Raw)
		if err != nil {
			return nil, fmt.Errorf("couldn't create a disk image at %q: %w", params.OutputPath, err)
		}
	}

	// Go through the initialization steps, starting with applying a
	// partition table according to params.
	if err := initializePartitionTable(diskImg, &params.PartitionSize, params.DiskGUID); err != nil {
		return nil, fmt.Errorf("failed to initialize the partition table: %w", err)
	}
	// The system partition will be created only if its specified size isn't
	// equal to zero, making the initialization step optional as well. In
	// addition, params.SystemImage must be set.
	if params.PartitionSize.System != 0 && params.SystemImage != nil {
		if err := initializeSystemPartition(diskImg, 2, params.SystemImage); err != nil {
			return nil, fmt.Errorf("failed to initialize the system partition: %w", err)
		}
	} else if params.PartitionSize.System == 0 && params.SystemImage != nil {
		// Safeguard against contradicting parameters.
		return nil, fmt.Errorf("the system image parameter was passed while the associated partition size is zero")
	}
	// Attempt to initialize the ESP unconditionally, as it's the only
	// partition guaranteed to be created regardless of params.PartitionSize.
	if err := initializeESP(diskImg, 1, params.EFIPayload, params.NodeParameters); err != nil {
		return nil, fmt.Errorf("failed to initialize the ESP: %w", err)
	}
	// The data partition, even if created, is always left uninitialized.

	// Build an EFI boot entry pointing to the image's ESP. go-diskfs won't let
	// you do that after you close the image.
	t, err := diskImg.GetPartitionTable()
	p := t.GetPartitions()
	esp := (p[0]).(*gpt.Partition)
	be := efivarfs.BootEntry{
		Description:     "Metropolis",
		Path:            EFIPayloadPath,
		PartitionGUID:   esp.GUID,
		PartitionNumber: 1,
		PartitionStart:  esp.Start,
		PartitionSize:   esp.End - esp.Start + 1,
	}
	// Close the image and return the EFI boot entry.
	if err := diskImg.File.Close(); err != nil {
		return nil, fmt.Errorf("failed to finalize image: %w", err)
	}
	return &be, nil
}
