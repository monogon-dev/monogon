// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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
	"source.monogon.dev/osbase/blockdev"
	"source.monogon.dev/osbase/efivarfs"
	"source.monogon.dev/osbase/gpt"
	"source.monogon.dev/osbase/supervisor"
	"source.monogon.dev/osbase/sysfs"
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
		//nolint:returnerrcheck
		return nil // Probably just not a GPT-partitioned disk
	}

	skipDisk := false
	if espUUID != uuid.Nil {
		// If we know where we booted from, ignore all disks which do
		// not contain this partition.
		skipDisk = true
		for _, part := range table.Partitions {
			if part.IsUnused() {
				continue
			}
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

// GrowPartition grows the GPT partition corresponding to the given block device
// path, by adding all free space immediately following the partition to the
// partition. The main use for this are virtual machines which are launched from
// an image which is smaller than the virtual disk.
func GrowPartition(partitionPath string) error {
	// Obtain sysfs path of the partition.
	var partS unix.Stat_t
	if err := unix.Stat(partitionPath, &partS); err != nil {
		return fmt.Errorf("inspecting partition %q: %w", partitionPath, err)
	}
	if partS.Mode&unix.S_IFMT != unix.S_IFBLK {
		return fmt.Errorf("%q is not a block device", partitionPath)
	}
	partitionSysPath := fmt.Sprintf("/sys/dev/block/%d:%d", unix.Major(partS.Rdev), unix.Minor(partS.Rdev))

	// Obtain partition number.
	partitionUevent, err := sysfs.ReadUevents(partitionSysPath + "/uevent")
	if err != nil {
		return fmt.Errorf("when reading uevent: %w", err)
	}
	if partitionUevent["DEVTYPE"] != "partition" {
		return fmt.Errorf("%q is not a partition", partitionPath)
	}
	partitionInfo, err := blockUEvent(partitionUevent).readPartitionInfo()
	if err != nil {
		return fmt.Errorf("when reading partition info: %w", err)
	}

	// Obtain disk info. Note that partitionSysPath is a symlink, and the ..
	// leads to the parent of the symlink target.
	diskUevent, err := sysfs.ReadUevents(partitionSysPath + "/../uevent")
	if err != nil {
		return fmt.Errorf("when reading uevent: %w", err)
	}
	if diskUevent["DEVTYPE"] != "disk" {
		return fmt.Errorf("parent of %q is not a disk", partitionPath)
	}

	// Open the disk device and read the GPT.
	blkdev, err := blockdev.Open(fmt.Sprintf("/dev/%v", diskUevent["DEVNAME"]))
	if err != nil {
		return fmt.Errorf("failed to open block device: %w", err)
	}
	defer blkdev.Close()

	table, err := gpt.Read(blkdev)
	if err != nil {
		//nolint:returnerrcheck
		return nil // Probably just not a GPT-partitioned disk
	}

	if partitionInfo.partNumber-1 > len(table.Partitions) {
		return fmt.Errorf("partition number %d out of bounds", partitionInfo.partNumber)
	}
	tableEntry := table.Partitions[partitionInfo.partNumber-1]
	if tableEntry.IsUnused() {
		return fmt.Errorf("partition %d is unused in the GPT", partitionInfo.partNumber)
	}

	freeSpaces, overlap, err := table.GetFreeSpaces()
	if err != nil {
		return err
	}
	var freeSpace [2]int64
	for _, sp := range freeSpaces {
		if sp[0] == int64(tableEntry.LastBlock)+1 {
			freeSpace = sp
			break
		}
	}
	if freeSpace[0] == 0 {
		// There is no free space after the partition, cannot grow.
		return nil
	}

	// Align the new partition end to 1 MiB to make sure that we do not constantly
	// overwrite hardware blocks containing the alternate GPT at the end,
	// increasing the possibility of corruption.
	alignBlocks := max(1*1024*1024/blkdev.BlockSize(), 1)
	alignedEnd := freeSpace[1] / alignBlocks * alignBlocks
	if alignedEnd <= freeSpace[0] {
		// There is no free space after aligning.
		return nil
	}

	if overlap {
		return fmt.Errorf("found overlap in GPT partitions, refusing to grow partition")
	}

	// We found free space after the partition to grow into.
	tableEntry.LastBlock = uint64(alignedEnd - 1)
	err = table.Write()
	if err != nil {
		return fmt.Errorf("failed to write GPT with grown partition: %w", err)
	}

	// Tell the kernel about the new partition size.
	err = blkdev.ResizePartition(
		int32(partitionInfo.partNumber),
		int64(tableEntry.FirstBlock)*blkdev.BlockSize(),
		int64(tableEntry.SizeBlocks())*blkdev.BlockSize(),
	)
	if err != nil {
		return fmt.Errorf("failed to resize partition in kernel after growing partition: %w", err)
	}

	partBlkdev, err := blockdev.Open(partitionPath)
	if err != nil {
		return fmt.Errorf("failed to open partition as block device: %w", err)
	}
	defer partBlkdev.Close()

	// Discard the newly allocated space. Do this on the partition instead of the
	// whole device, because that works more reliably. When using the whole
	// device, discard may fail if there are dirty pages, but when using the
	// partition, discard can take an exclusive lock, which avoids this problem.
	//
	// Ignore errors, this is only advisory.
	partBlkdev.Discard(
		(freeSpace[0]-int64(tableEntry.FirstBlock))*blkdev.BlockSize(),
		int64(tableEntry.SizeBlocks())*blkdev.BlockSize(),
	)

	return nil
}
