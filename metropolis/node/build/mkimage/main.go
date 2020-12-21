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

package main

// mkimage is a tool to generate a Metropolis node disk image containing the
// given EFI payload, and optionally, a given external initramfs image and
// enrolment credentials.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	diskfs "github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
	"github.com/diskfs/go-diskfs/partition/gpt"
)

var NodeDataPartition gpt.Type = gpt.Type("9eeec464-6885-414a-b278-4305c51f7966")

var (
	flagEFI                  string
	flagOut                  string
	flagInitramfs            string
	flagEnrolmentCredentials string
	flagDataPartitionSize    uint64
	flagESPPartitionSize     uint64
)

func mibToSectors(size uint64) uint64 {
	return (size * 1024 * 1024) / 512
}

func main() {
	flag.StringVar(&flagEFI, "efi", "", "UEFI payload")
	flag.StringVar(&flagOut, "out", "", "Output disk image")
	flag.StringVar(&flagInitramfs, "initramfs", "", "External initramfs [optional]")
	flag.StringVar(&flagEnrolmentCredentials, "enrolment_credentials", "", "Enrolment credentials [optional]")
	flag.Uint64Var(&flagDataPartitionSize, "data_partition_size", 2048, "Override the data partition size (default 2048 MiB)")
	flag.Uint64Var(&flagESPPartitionSize, "esp_partition_size", 512, "Override the ESP partition size (default: 512MiB)")
	flag.Parse()

	if flagEFI == "" || flagOut == "" {
		log.Fatalf("efi and initramfs must be set")
	}

	_ = os.Remove(flagOut)
	diskImg, err := diskfs.Create(flagOut, 3*1024*1024*1024, diskfs.Raw)
	if err != nil {
		log.Fatalf("diskfs.Create(%q): %v", flagOut, err)
	}

	table := &gpt.Table{
		// This is appropriate at least for virtio disks. Might need to be adjusted for real ones.
		LogicalSectorSize:  512,
		PhysicalSectorSize: 512,
		ProtectiveMBR:      true,
		Partitions: []*gpt.Partition{
			{
				Type:  gpt.EFISystemPartition,
				Name:  "ESP",
				Start: mibToSectors(1),
				End:   mibToSectors(flagESPPartitionSize) - 1,
			},
			{
				Type:  NodeDataPartition,
				Name:  "METROPOLIS-NODE-DATA",
				Start: mibToSectors(flagESPPartitionSize),
				End:   mibToSectors(flagESPPartitionSize+flagDataPartitionSize) - 1,
			},
		},
	}
	if err := diskImg.Partition(table); err != nil {
		log.Fatalf("Failed to apply partition table: %v", err)
	}

	fs, err := diskImg.CreateFilesystem(disk.FilesystemSpec{Partition: 1, FSType: filesystem.TypeFat32, VolumeLabel: "ESP"})
	if err != nil {
		log.Fatalf("Failed to create filesystem: %v", err)
	}

	// Create EFI partition structure.
	for _, dir := range []string{"/EFI", "/EFI/BOOT", "/EFI/metropolis"} {
		if err := fs.Mkdir(dir); err != nil {
			log.Fatalf("Mkdir(%q): %v", dir, err)
		}
	}

	put(fs, flagEFI, "/EFI/BOOT/BOOTX64.EFI")

	if flagInitramfs != "" {
		put(fs, flagInitramfs, "/EFI/metropolis/initramfs.cpio.lz4")
	}

	if flagEnrolmentCredentials != "" {
		put(fs, flagEnrolmentCredentials, "/EFI/metropolis/enrolment.pb")
	}

	if err := diskImg.File.Close(); err != nil {
		log.Fatalf("Failed to finalize image: %v", err)
	}
	log.Printf("Success! You can now boot %v", flagOut)
}

// put copies a file from the host filesystem into the target image.
func put(fs filesystem.FileSystem, src, dst string) {
	target, err := fs.OpenFile(dst, os.O_CREATE|os.O_RDWR)
	if err != nil {
		log.Fatalf("fs.OpenFile(%q): %v", dst, err)
	}
	source, err := os.Open(src)
	if err != nil {
		log.Fatalf("os.Open(%q): %v", src, err)
	}
	defer source.Close()
	// If this is streamed (e.g. using io.Copy) it exposes a bug in diskfs, so do it in one go.
	data, err := ioutil.ReadAll(source)
	if err != nil {
		log.Fatalf("Reading %q: %v", src, err)
	}
	if _, err := target.Write(data); err != nil {
		fmt.Printf("writing file %q: %v", dst, err)
		os.Exit(1)
	}
}
