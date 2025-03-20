// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// mkimage is a tool to generate node disk images.
// It can be used both to initialize block devices and to create image
// files.
//
// The tool takes a path to an EFI payload (--efi), a path to a abloader
// payload (--abloader) and a path to a system image (--system) as its only
// required inputs. In addition, an output path must be supplied (--out).
// Node parameters file path (--node_parameters) may also be supplied, in
// which case the file will be copied to the EFI system partition.
// Partition sizes are fixed and may be overridden by command line flags.
package main

import (
	_ "embed"
	"flag"
	"log"
	"os"

	"source.monogon.dev/osbase/blockdev"
	"source.monogon.dev/osbase/build/mkimage/osimage"
	"source.monogon.dev/osbase/structfs"
)

func main() {
	// Fill in the image parameters based on flags.
	var (
		efiPayload          string
		systemImage         string
		abLoaderPayload     string
		biosBootCodePayload string
		nodeParams          string
		outputPath          string
		diskUUID            string
		cfg                 osimage.Params
	)
	flag.StringVar(&efiPayload, "efi", "", "Path to the UEFI payload used")
	flag.StringVar(&systemImage, "system", "", "Path to the system partition image used")
	flag.StringVar(&abLoaderPayload, "abloader", "", "Path to the abloader payload used")
	flag.StringVar(&biosBootCodePayload, "bios_bootcode", "", "Optional path to the BIOS bootcode which gets placed at the start of the first block of the image. Limited to 440 bytes, padding is not required. It is only used by legacy BIOS boot.")
	flag.StringVar(&nodeParams, "node_parameters", "", "Path to Node Parameters to be written to the ESP (default: don't write Node Parameters)")
	flag.StringVar(&outputPath, "out", "", "Path to the resulting disk image or block device")
	flag.Int64Var(&cfg.PartitionSize.Data, "data_partition_size", 2048, "Override the data partition size (default 2048 MiB). Used only when generating image files.")
	flag.Int64Var(&cfg.PartitionSize.ESP, "esp_partition_size", 128, "Override the ESP partition size (default: 128MiB)")
	flag.Int64Var(&cfg.PartitionSize.System, "system_partition_size", 1024, "Override the System partition size (default: 1024MiB)")
	flag.StringVar(&diskUUID, "GUID", "", "Disk GUID marked in the resulting image's partition table (default: randomly generated)")
	flag.Parse()

	// Open the input files for osimage.Create, fill in reader objects and
	// metadata in osimage.Params.
	// Start with the EFI Payload the OS will boot from.
	var err error
	cfg.EFIPayload, err = structfs.OSPathBlob(efiPayload)
	if err != nil {
		log.Fatalf("while opening the EFI payload at %q: %v", efiPayload, err)
	}

	cfg.ABLoader, err = structfs.OSPathBlob(abLoaderPayload)
	if err != nil {
		log.Fatalf("while opening the abloader payload at %q: %v", abLoaderPayload, err)
	}

	// Attempt to open the system image if its path is set. In case the path
	// isn't set, the system partition will still be created, but no
	// contents will be written into it.
	if systemImage != "" {
		cfg.SystemImage, err = structfs.OSPathBlob(systemImage)
		if err != nil {
			log.Fatalf("while opening the system image at %q: %v", systemImage, err)
		}
	}

	// Attempt to open the node parameters file if its path is set.
	if nodeParams != "" {
		np, err := structfs.OSPathBlob(nodeParams)
		if err != nil {
			log.Fatalf("while opening node parameters at %q: %v", nodeParams, err)
		}
		cfg.NodeParameters = np
	}

	if biosBootCodePayload != "" {
		bp, err := os.ReadFile(biosBootCodePayload)
		if err != nil {
			log.Fatalf("while opening BIOS bootcode at %q: %v", biosBootCodePayload, err)
		}
		cfg.BIOSBootCode = bp
	}

	// TODO(#254): Build and use dynamically-grown block devices
	cfg.Output, err = blockdev.CreateFile(outputPath, 512, 10*1024*1024)
	if err != nil {
		panic(err)
	}

	// Write the parametrized OS image.
	if _, err := osimage.Write(&cfg); err != nil {
		log.Fatalf("while creating a Metropolis OS image: %v", err)
	}
}
