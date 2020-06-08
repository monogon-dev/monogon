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

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	diskfs "github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
	"github.com/diskfs/go-diskfs/partition/gpt"
)

var SmalltownDataPartition gpt.Type = gpt.Type("9eeec464-6885-414a-b278-4305c51f7966")

var (
	efiPayloadPath           = flag.String("efi", "", "UEFI payload")
	outputPath               = flag.String("out", "", "Output disk image")
	initramfsPath            = flag.String("initramfs", "", "External initramfs [optional]")
	enrolmentCredentialsPath = flag.String("enrolment-credentials", "", "Enrolment credentials [optional]")
	dataPartitionSizeMiB     = flag.Uint64("data-partition-size", 2048, "Override the data partition size (default 2048 MiB)")
	espPartitionSizeMiB      = flag.Uint64("esp-partition-size", 512, "Override the ESP partition size (default: 512MiB)")
)

func mibToSectors(size uint64) uint64 {
	return (size * 1024 * 1024) / 512
}

func main() {
	flag.Parse()
	if *efiPayloadPath == "" || *outputPath == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}

	_ = os.Remove(*outputPath)
	diskImg, err := diskfs.Create(*outputPath, 3*1024*1024*1024, diskfs.Raw)
	if err != nil {
		fmt.Printf("Failed to create disk: %v", err)
		os.Exit(1)
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
				End:   mibToSectors(*espPartitionSizeMiB) - 1,
			},
			{
				Type:  SmalltownDataPartition,
				Name:  "SIGNOS-DATA",
				Start: mibToSectors(*espPartitionSizeMiB),
				End:   mibToSectors(*espPartitionSizeMiB+*dataPartitionSizeMiB) - 1,
			},
		},
	}
	if err := diskImg.Partition(table); err != nil {
		fmt.Printf("Failed to apply partition table: %v", err)
		os.Exit(1)
	}

	fs, err := diskImg.CreateFilesystem(disk.FilesystemSpec{Partition: 1, FSType: filesystem.TypeFat32, VolumeLabel: "ESP"})
	if err != nil {
		fmt.Printf("Failed to create filesystem: %v", err)
		os.Exit(1)
	}
	if err := fs.Mkdir("/EFI"); err != nil {
		panic(err)
	}
	if err := fs.Mkdir("/EFI/BOOT"); err != nil {
		panic(err)
	}
	if err := fs.Mkdir("/EFI/smalltown"); err != nil {
		panic(err)
	}
	efiPayload, err := fs.OpenFile("/EFI/BOOT/BOOTX64.EFI", os.O_CREATE|os.O_RDWR)
	if err != nil {
		fmt.Printf("Failed to open EFI payload for writing: %v", err)
		os.Exit(1)
	}
	efiPayloadSrc, err := os.Open(*efiPayloadPath)
	if err != nil {
		fmt.Printf("Failed to open EFI payload for reading: %v", err)
		os.Exit(1)
	}
	defer efiPayloadSrc.Close()
	// If this is streamed (e.g. using io.Copy) it exposes a bug in diskfs, so do it in one go.
	efiPayloadFull, err := ioutil.ReadAll(efiPayloadSrc)
	if err != nil {
		panic(err)
	}
	if _, err := efiPayload.Write(efiPayloadFull); err != nil {
		fmt.Printf("Failed to write EFI payload: %v", err)
		os.Exit(1)
	}
	initramfs, err := fs.OpenFile("/EFI/smalltown/initramfs.cpio.lz4", os.O_CREATE|os.O_RDWR)
	if err != nil {
		fmt.Printf("Failed to open initramfs for writing: %v", err)
		os.Exit(1)
	}
	// If we have more than two arguments, the second one is the initramfs
	if *initramfsPath != "" {
		initramfsSrc, err := os.Open(*initramfsPath)
		if err != nil {
			fmt.Printf("Failed to open initramfs for reading: %v", err)
			os.Exit(1)
		}
		initramfsFull, err := ioutil.ReadAll(initramfsSrc)
		if err != nil {
			fmt.Printf("Failed to read initramfs: %v", err)
			os.Exit(1)
		}
		if _, err := initramfs.Write(initramfsFull); err != nil {
			fmt.Printf("Failed to write initramfs: %v", err)
			os.Exit(1)
		}
	}
	if *enrolmentCredentialsPath != "" {
		enrolmentCredentials, err := fs.OpenFile("/EFI/smalltown/enrolment.pb", os.O_CREATE|os.O_RDWR)
		enrolmentCredentialsSrc, err := os.Open(*enrolmentCredentialsPath)
		if err != nil {
			fmt.Printf("Failed to open enrolment credentials for reading: %v", err)
			os.Exit(1)
		}
		enrolmentCredentialsFull, err := ioutil.ReadAll(enrolmentCredentialsSrc)
		if err != nil {
			fmt.Printf("Failed to read enrolment credentials: %v", err)
			os.Exit(1)
		}
		if _, err := enrolmentCredentials.Write(enrolmentCredentialsFull); err != nil {
			fmt.Printf("Failed to write enrolment credentials")
			os.Exit(1)
		}
	}
	if err := diskImg.File.Close(); err != nil {
		fmt.Printf("Failed to write image: %v", err)
		os.Exit(1)
	}
	fmt.Println("Success! You can now boot smalltown.img")
}
