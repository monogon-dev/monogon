package core

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"

	"google.golang.org/protobuf/proto"

	"source.monogon.dev/metropolis/pkg/blockdev"
	"source.monogon.dev/metropolis/pkg/fat32"
	"source.monogon.dev/metropolis/pkg/gpt"
	"source.monogon.dev/metropolis/proto/api"
)

type MakeInstallerImageArgs struct {
	// Path to either a file or a disk which will contain the installer data.
	TargetPath string

	// Reader for the installer EFI executable. Mandatory.
	Installer fat32.SizedReader

	// Optional NodeParameters to be embedded for use by the installer.
	NodeParams *api.NodeParameters

	// Optional Reader for a Metropolis bundle for use by the installer.
	Bundle fat32.SizedReader
}

// MakeInstallerImage generates an installer disk image containing a Table
// partition table and a single FAT32 partition with an installer and optionally
// with a bundle and/or Node Parameters.
func MakeInstallerImage(args MakeInstallerImageArgs) error {
	if args.Installer == nil {
		return errors.New("Installer is mandatory")
	}

	espRoot := fat32.Inode{Attrs: fat32.AttrDirectory}

	// This needs to be a "Removable Media" according to the UEFI Specification
	// V2.9 Section 3.5.1.1. This file is booted by any compliant UEFI firmware
	// in absence of another bootable boot entry.
	if err := espRoot.PlaceFile("EFI/BOOT/BOOTx64.EFI", args.Installer); err != nil {
		return err
	}

	if args.NodeParams != nil {
		nodeParamsRaw, err := proto.Marshal(args.NodeParams)
		if err != nil {
			return fmt.Errorf("failed to marshal node params: %w", err)
		}
		if err := espRoot.PlaceFile("metropolis-installer/nodeparams.pb", bytes.NewReader(nodeParamsRaw)); err != nil {
			return err
		}
	}
	if args.Bundle != nil {
		if err := espRoot.PlaceFile("metropolis-installer/bundle.bin", args.Bundle); err != nil {
			return err
		}
	}
	var targetDev blockdev.BlockDev
	var err error
	targetDev, err = blockdev.Open(args.TargetPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			targetDev, err = blockdev.CreateFile(args.TargetPath, 512, 1024*1024+4096)
		}
		if err != nil {
			return fmt.Errorf("unable to open target device: %w", err)
		}
	}
	partTable, err := gpt.New(targetDev)
	if err != nil {
		return fmt.Errorf("target device has invalid geometry: %w", err)
	}
	esp := gpt.Partition{
		Type: gpt.PartitionTypeEFISystem,
		Name: "MetropolisInstaller",
	}
	fatOpts := fat32.Options{Label: "METRO_INST"}
	// TODO(#254): Build and use dynamically-grown block devices
	var espSize int64 = 512 * 1024 * 1024
	if err := partTable.AddPartition(&esp, espSize); err != nil {
		return fmt.Errorf("unable to create partition layout: %w", err)
	}
	if esp.BlockSize() > math.MaxUint16 {
		return fmt.Errorf("block size (%d) too large for FAT32", esp.BlockSize())
	}
	fatOpts.BlockSize = uint16(esp.BlockSize())
	fatOpts.BlockCount = uint32(esp.BlockCount())
	if err := fat32.WriteFS(blockdev.NewRWS(esp), espRoot, fatOpts); err != nil {
		return fmt.Errorf("failed to write FAT32: %w", err)
	}
	if err := partTable.Write(); err != nil {
		return fmt.Errorf("unable to write partition table: %w", err)
	}
	return nil
}
