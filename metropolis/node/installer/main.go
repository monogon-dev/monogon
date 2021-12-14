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

// Installer creates a Metropolis image at a suitable block device based on the
// installer bundle present in the installation medium's ESP, after which it
// reboots. It's meant to be used as an init process.
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/sys/unix"
	"source.monogon.dev/metropolis/node/build/mkimage/osimage"
	"source.monogon.dev/metropolis/pkg/efivarfs"
	"source.monogon.dev/metropolis/pkg/sysfs"
)

const mib = 1024 * 1024

// mountPseudoFS mounts efivarfs, devtmpfs and sysfs, used by the installer in
// the block device discovery stage.
func mountPseudoFS() error {
	for _, m := range []struct {
		dir   string
		fs    string
		flags uintptr
	}{
		{"/sys", "sysfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{efivarfs.Path, "efivarfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/dev", "devtmpfs", unix.MS_NOEXEC | unix.MS_NOSUID},
	} {
		if err := unix.Mkdir(m.dir, 0700); err != nil && !os.IsExist(err) {
			return fmt.Errorf("couldn't create the mountpoint at %q: %w", m.dir, err)
		}
		if err := unix.Mount(m.fs, m.dir, m.fs, m.flags, ""); err != nil {
			return fmt.Errorf("couldn't mount %q at %q: %w", m.fs, m.dir, err)
		}
	}
	return nil
}

// mountInstallerESP mounts the filesystem the installer was loaded from based
// on espPath, which must point to the appropriate partition block device. The
// filesystem is mounted at /installer.
func mountInstallerESP(espPath string) error {
	// Create the mountpoint.
	if err := unix.Mkdir("/installer", 0700); err != nil {
		return fmt.Errorf("couldn't create the installer mountpoint: %w", err)
	}
	// Mount the filesystem.
	if err := unix.Mount(espPath, "/installer", "vfat", unix.MS_NOEXEC|unix.MS_RDONLY, ""); err != nil {
		return fmt.Errorf("couldn't mount the installer ESP (%q -> %q): %w", espPath, "/installer", err)
	}
	return nil
}

// findInstallableBlockDevices returns names of all the block devices suitable
// for hosting a Metropolis installation, limited by the size expressed in
// bytes minSize. The install medium espDev will be excluded from the result.
func findInstallableBlockDevices(espDev string, minSize uint64) ([]string, error) {
	// Use the partition's name to find and return the name of its parent
	// device. It will be excluded from the list of suitable target devices.
	srcDev, err := sysfs.ParentBlockDevice(espDev)
	// Build the exclusion list containing forbidden handle prefixes.
	exclude := []string{"dm-", "zram", "ram", "loop", srcDev}

	// Get the block device handles by looking up directory contents.
	const blkDirPath = "/sys/class/block"
	blkDevs, err := os.ReadDir(blkDirPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't read %q: %w", blkDirPath, err)
	}
	// Iterate over the handles, skipping any block device that either points to
	// a partition, matches the exclusion list, or is smaller than minSize.
	var suitable []string
probeLoop:
	for _, devInfo := range blkDevs {
		// Skip devices according to the exclusion list.
		for _, prefix := range exclude {
			if strings.HasPrefix(devInfo.Name(), prefix) {
				continue probeLoop
			}
		}

		// Skip partition symlinks.
		if _, err := os.Stat(filepath.Join(blkDirPath, devInfo.Name(), "partition")); err == nil {
			continue
		} else if !os.IsNotExist(err) {
			return nil, fmt.Errorf("while probing sysfs: %w", err)
		}

		// Skip devices of insufficient size.
		devPath := filepath.Join("/dev", devInfo.Name())
		dev, err := os.Open(devPath)
		if err != nil {
			return nil, fmt.Errorf("couldn't open a block device at %q: %w", devPath, err)
		}
		size, err := unix.IoctlGetInt(int(dev.Fd()), unix.BLKGETSIZE64)
		dev.Close()
		if err != nil {
			return nil, fmt.Errorf("couldn't probe the size of %q: %w", devPath, err)
		}
		if uint64(size) < minSize {
			continue
		}

		suitable = append(suitable, devInfo.Name())
	}
	return suitable, nil
}

// rereadPartitionTable causes the kernel to read the partition table present
// in the block device at blkdevPath. It may return an error.
func rereadPartitionTable(blkdevPath string) error {
	dev, err := os.Open(blkdevPath)
	if err != nil {
		return fmt.Errorf("couldn't open the block device at %q: %w", blkdevPath, err)
	}
	defer dev.Close()
	ret, err := unix.IoctlRetInt(int(dev.Fd()), unix.BLKRRPART)
	if err != nil {
		return fmt.Errorf("while doing an ioctl: %w", err)
	}
	if syscall.Errno(ret) == unix.EINVAL {
		return fmt.Errorf("got an EINVAL from BLKRRPART ioctl")
	}
	return nil
}

// initializeSystemPartition writes image contents to the node's system
// partition using the block device abstraction layer as opposed to slower
// go-diskfs. tgtBlkdev must contain a path pointing to the block device
// associated with the system partition. It may return an error.
func initializeSystemPartition(image io.Reader, tgtBlkdev string) error {
	// Check that tgtBlkdev points at an actual block device.
	info, err := os.Stat(tgtBlkdev)
	if err != nil {
		return fmt.Errorf("couldn't stat the system partition at %q: %w", tgtBlkdev, err)
	}
	if info.Mode()&os.ModeDevice == 0 {
		return fmt.Errorf("system partition path %q doesn't point to a block device", tgtBlkdev)
	}

	// Get the system partition's file descriptor.
	sys, err := os.OpenFile(tgtBlkdev, os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("couldn't open the system partition at %q: %w", tgtBlkdev, err)
	}
	defer sys.Close()
	// Copy the system partition contents. Use a bigger buffer to optimize disk
	// writes.
	buf := make([]byte, mib)
	if _, err := io.CopyBuffer(sys, image, buf); err != nil {
		return fmt.Errorf("couldn't copy partition contents: %w", err)
	}
	return nil
}

func main() {
	// Mount sysfs, devtmpfs and efivarfs.
	if err := mountPseudoFS(); err != nil {
		log.Fatalf("while mounting pseudo-filesystems: %v", err)
	}
	// Read the installer ESP UUID from efivarfs.
	espUuid, err := efivarfs.ReadLoaderDevicePartUUID()
	if err != nil {
		log.Fatalf("while reading the installer ESP UUID: %v", err)
	}
	// Look up the installer partition based on espUuid.
	espDev, err := sysfs.DeviceByPartUUID(espUuid)
	espPath := filepath.Join("/dev", espDev)
	if err != nil {
		log.Fatalf("while resolving the installer device handle: %v", err)
	}
	// Mount the installer partition. The installer bundle will be read from it.
	if err := mountInstallerESP(espPath); err != nil {
		log.Fatalf("while mounting the installer ESP: %v", err)
	}

	nodeParameters, err := os.Open("/installer/metropolis-installer/nodeparams.pb")
	if err != nil {
		log.Fatalf("failed to open node parameters from ESP: %v", err)
	}

	// TODO(lorenz): Replace with proper bundles
	bundle, err := zip.OpenReader("/installer/metropolis-installer/bundle.bin")
	if err != nil {
		log.Fatalf("failed to open node bundle from ESP: %v", err)
	}
	defer bundle.Close()
	efiPayload, err := bundle.Open("kernel_efi.efi")
	if err != nil {
		log.Fatalf("Cannot open EFI payload in bundle: %v", err)
	}
	defer efiPayload.Close()
	systemImage, err := bundle.Open("rootfs.img")
	if err != nil {
		log.Fatalf("Cannot open system image in bundle: %v", err)
	}
	defer systemImage.Close()

	// Build the osimage parameters.
	installParams := osimage.Params{
		PartitionSize: osimage.PartitionSizeInfo{
			// ESP is the size of the node ESP partition, expressed in mebibytes.
			ESP: 128,
			// System is the size of the node system partition, expressed in
			// mebibytes.
			System: 4096,
			// Data must be nonzero in order for the data partition to be created.
			// osimage will extend the data partition to fill all the available space
			// whenever it's writing to block devices, such as now.
			Data: 128,
		},
		// Due to a bug in go-diskfs causing slow writes, SystemImage is explicitly
		// marked unused here, as system partition contents will be written using
		// a workaround below instead.
		// TODO(mateusz@monogon.tech): Address that bug either by patching go-diskfs
		// or rewriting osimage.
		SystemImage: nil,

		EFIPayload:     efiPayload,
		NodeParameters: nodeParameters,
	}
	// Calculate the minimum target size based on the installation parameters.
	minSize := uint64((installParams.PartitionSize.ESP +
		installParams.PartitionSize.System +
		installParams.PartitionSize.Data + 1) * mib)

	// Look for suitable block devices, given the minimum size.
	blkDevs, err := findInstallableBlockDevices(espDev, minSize)
	if err != nil {
		log.Fatal(err)
	}
	if len(blkDevs) == 0 {
		log.Fatal("couldn't find a suitable block device.")
	}
	// Set the first suitable block device found as the installation target.
	tgtBlkdevName := blkDevs[0]
	// Update the osimage parameters with a path pointing at the target device.
	tgtBlkdevPath := filepath.Join("/dev", tgtBlkdevName)
	installParams.OutputPath = tgtBlkdevPath

	// Use osimage to partition the target block device and set up its ESP.
	// Create will return an EFI boot entry on success.
	log.Printf("Installing to %s\n", tgtBlkdevPath)
	be, err := osimage.Create(&installParams)
	if err != nil {
		log.Fatalf("while installing: %v", err)
	}
	// The target device's partition table has just been updated. Re-read it to
	// make the node system partition reachable through /dev.
	if err := rereadPartitionTable(tgtBlkdevPath); err != nil {
		log.Fatalf("while re-reading the partition table of %q: %v", tgtBlkdevPath, err)
	}
	// Look up the node's system partition path to be later used in the
	// initialization step. It's always the second partition, right after
	// the ESP.
	sysBlkdevName, err := sysfs.PartitionBlockDevice(tgtBlkdevName, 2)
	if err != nil {
		log.Fatalf("while looking up the system partition: %v", err)
	}
	sysBlkdevPath := filepath.Join("/dev", sysBlkdevName)
	// Copy the system partition contents.
	if err := initializeSystemPartition(systemImage, sysBlkdevPath); err != nil {
		log.Fatalf("while initializing the system partition at %q: %v", sysBlkdevPath, err)
	}

	// Create an EFI boot entry for Metropolis.
	en, err := efivarfs.CreateBootEntry(be)
	if err != nil {
		log.Fatalf("while creating a boot entry: %v", err)
	}
	// Erase the preexisting boot order, leaving Metropolis as the only option.
	if err := efivarfs.SetBootOrder(&efivarfs.BootOrder{uint16(en)}); err != nil {
		log.Fatalf("while adjusting the boot order: %v", err)
	}

	// Reboot.
	unix.Sync()
	log.Print("Installation completed. Rebooting.")
	unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART)
}
