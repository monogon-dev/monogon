package core

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
	"github.com/diskfs/go-diskfs/partition/gpt"
	"google.golang.org/protobuf/proto"
	"source.monogon.dev/metropolis/proto/api"
)

func mibToSectors(size uint64, logicalBlockSize int64) uint64 {
	return (size * 1024 * 1024) / uint64(logicalBlockSize)
}

// Mask for aligning values to 1MiB boundaries. Go complains if you shift
// 1 bits out of the value in a constant so the construction is a bit
// convoluted.
const align1MiBMask = (1<<44 - 1) << 20

const MiB = 1024 * 1024

type MakeInstallerImageArgs struct {
	// Path to either a file or a disk which will contain the installer data.
	TargetPath string

	// Reader for the installer EFI executable. Mandatory.
	Installer     io.Reader
	InstallerSize uint64

	// Optional NodeParameters to be embedded for use by the installer.
	NodeParams *api.NodeParameters

	// Optional Reader for a Metropolis bundle for use by the installer.
	Bundle     io.Reader
	BundleSize uint64
}

// MakeInstallerImage generates an installer disk image containing a GPT
// partition table and a single FAT32 partition with an installer and optionally
// with a bundle and/or Node Parameters.
func MakeInstallerImage(args MakeInstallerImageArgs) error {
	if args.Installer == nil {
		return errors.New("Installer is mandatory")
	}
	if args.InstallerSize == 0 {
		return errors.New("InstallerSize needs to be valid (>0)")
	}
	if args.Bundle != nil && args.BundleSize == 0 {
		return errors.New("if a Bundle is passed BundleSize needs to be valid (>0)")
	}

	var err error
	var nodeParamsRaw []byte
	if args.NodeParams != nil {
		nodeParamsRaw, err = proto.Marshal(args.NodeParams)
		if err != nil {
			return fmt.Errorf("failed to marshal node params: %w", err)
		}
	}

	var img *disk.Disk
	// The following section is a bit ugly, it would technically be nicer to
	// just pack all clusters of the FAT32 together, figure out how many were
	// needed at the end and truncate the partition there. But that would
	// require writing a new FAT32 writer, the effort to do that is in no way
	// proportional to its advantages. So let's just do some conservative
	// calculations on how much space we need and call it a day.

	// ~4MiB FAT32 headers, 1MiB alignment overhead (bitmask drops up to 1MiB),
	// 5% filesystem overhead
	partitionSizeBytes := (uint64(float32(5*MiB+args.BundleSize+args.InstallerSize+uint64(len(nodeParamsRaw))) * 1.05)) & align1MiBMask
	// FAT32 has a minimum partition size of 32MiB, so clamp the lower partition
	// size to a notch more than that.
	minimumSize := uint64(33 * MiB)
	if partitionSizeBytes < minimumSize {
		partitionSizeBytes = minimumSize
	}
	// If creating an image, create it with minimal size, i.e. 1MiB at each
	// end for partitioning metadata and alignment.
	// 1MiB alignment is used as that will essentially guarantee that any
	// partition is aligned to whatever internal block size is used by the
	// storage device. Especially flash-based storage likes to use much bigger
	// blocks than advertised as sectors which can degrade performance when
	// partitions are misaligned.
	calculatedImageBytes := 2*MiB + partitionSizeBytes

	if _, err = os.Stat(args.TargetPath); os.IsNotExist(err) {
		img, err = diskfs.Create(args.TargetPath, int64(calculatedImageBytes), diskfs.Raw)
	} else {
		img, err = diskfs.Open(args.TargetPath)
	}
	if err != nil {
		return fmt.Errorf("failed to create/open target: %w", err)
	}
	defer img.File.Close()
	// This has an edge case where the data would technically fit but our 5%
	// overhead are too conservative. But it is very rare and I don't really
	// trust diskfs to generate good errors when it overflows so we'll just
	// reject early.
	if uint64(img.Size) < calculatedImageBytes {
		return errors.New("target too small, data won't fit")
	}

	gptTable := &gpt.Table{
		LogicalSectorSize:  int(img.LogicalBlocksize),
		PhysicalSectorSize: int(img.PhysicalBlocksize),
		ProtectiveMBR:      true,
		Partitions: []*gpt.Partition{
			{
				Type:  gpt.EFISystemPartition,
				Name:  "MetropolisInstaller",
				Start: mibToSectors(1, img.LogicalBlocksize),
				Size:  partitionSizeBytes,
			},
		},
	}
	if err := img.Partition(gptTable); err != nil {
		return fmt.Errorf("failed to partition target: %w", err)
	}
	fs, err := img.CreateFilesystem(disk.FilesystemSpec{Partition: 1, FSType: filesystem.TypeFat32, VolumeLabel: "METRO_INST"})
	if err != nil {
		return fmt.Errorf("failed to create target filesystem: %w", err)
	}

	// Create EFI partition structure.
	for _, dir := range []string{"/EFI", "/EFI/BOOT", "/EFI/metropolis-installer"} {
		if err := fs.Mkdir(dir); err != nil {
			panic(err)
		}
	}
	// This needs to be a "Removable Media" according to the UEFI Specification
	// V2.9 Section 3.5.1.1. This file is booted by any compliant UEFI firmware
	// in absence of another bootable boot entry.
	installerFile, err := fs.OpenFile("/EFI/BOOT/BOOTx64.EFI", os.O_CREATE|os.O_RDWR)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(installerFile, args.Installer); err != nil {
		return fmt.Errorf("failed to copy installer file: %w", err)
	}
	if args.NodeParams != nil {
		nodeParamsFile, err := fs.OpenFile("/EFI/metropolis-installer/nodeparams.pb", os.O_CREATE|os.O_RDWR)
		if err != nil {
			panic(err)
		}
		_, err = nodeParamsFile.Write(nodeParamsRaw)
		if err != nil {
			return fmt.Errorf("failed to write node params: %w", err)
		}
	}
	if args.Bundle != nil {
		bundleFile, err := fs.OpenFile("/EFI/metropolis-installer/bundle.bin", os.O_CREATE|os.O_RDWR)
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(bundleFile, args.Bundle); err != nil {
			return fmt.Errorf("failed to copy bundle: %w", err)
		}
	}
	return nil
}
