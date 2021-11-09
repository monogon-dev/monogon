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

// This package runs the installer image in a VM provided with an empty block
// device. It then examines the installer console output and the blok device to
// determine whether the installation process completed without issue.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"testing"

	diskfs "github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/partition/gpt"
	"source.monogon.dev/metropolis/proto/api"

	mctl "source.monogon.dev/metropolis/cli/metroctl/core"
	osimage "source.monogon.dev/metropolis/node/build/mkimage/osimage"
)

const (
	InstallerEFIPayload = "metropolis/node/installer/kernel.efi"
	TestOSBundle        = "metropolis/test/installer/testos/testos_bundle.zip"
	InstallerImage      = "metropolis/test/installer/installer.img"
	NodeStorage         = "metropolis/test/installer/stor.img"
)

// runQemu starts a QEMU process and waits for it to finish. args is
// concatenated to the list of predefined default arguments. It returns true if
// expectedOutput is found in the serial port output. It may return an error.
func runQemu(args []string, expectedOutput string) (bool, error) {
	// Prepare the default parameter list.
	defaultArgs := []string{
		"-machine", "q35", "-accel", "kvm", "-nographic", "-nodefaults",
		"-m", "512",
		"-smp", "2",
		"-cpu", "host",
		"-drive", "if=pflash,format=raw,readonly,file=external/edk2/OVMF_CODE.fd",
		"-drive", "if=pflash,format=raw,snapshot=on,file=external/edk2/OVMF_VARS.fd",
		"-serial", "stdio",
		"-no-reboot",
	}
	// Join the parameter lists and prepare the Qemu command, but don't run it
	// just yet.
	qemuArgs := append(defaultArgs, args...)
	qemuCmd := exec.Command("external/qemu/qemu-x86_64-softmmu", qemuArgs...)

	// Copy the stdout and stderr output so that it could be matched against
	// expectedOutput later.
	var outBuf, errBuf bytes.Buffer
	outWriter := bufio.NewWriter(&outBuf)
	errWriter := bufio.NewWriter(&errBuf)
	qemuCmd.Stdout = io.MultiWriter(os.Stdout, outWriter)
	qemuCmd.Stderr = io.MultiWriter(os.Stderr, errWriter)
	if err := qemuCmd.Run(); err != nil {
		return false, fmt.Errorf("couldn't start QEMU: %w", err)
	}
	outWriter.Flush()
	errWriter.Flush()

	// Try matching against expectedOutput and return the result.
	result := bytes.Contains(outBuf.Bytes(), []byte(expectedOutput)) ||
		bytes.Contains(errBuf.Bytes(), []byte(expectedOutput))
	return result, nil
}

// runQemuWithInstaller starts a QEMU process and waits for it to finish. args is
// concatenated to the list of predefined default arguments. It returns true if
// expectedOutput is found in the serial port output. It may return an error.
func runQemuWithInstaller(args []string, expectedOutput string) (bool, error) {
	args = append(args, "-drive", "if=virtio,format=raw,snapshot=on,cache=unsafe,file="+InstallerImage)
	return runQemu(args, expectedOutput)
}

// getStorage creates a sparse file, given a size expressed in mebibytes, and
// returns a path to that file. It may return an error.
func getStorage(size int64) (string, error) {
	image, err := os.Create(NodeStorage)
	if err != nil {
		return "", fmt.Errorf("couldn't create the block device image at %q: %w", NodeStorage, err)
	}
	if err := syscall.Ftruncate(int(image.Fd()), size*1024*1024); err != nil {
		return "", fmt.Errorf("couldn't resize the block device image at %q: %w", NodeStorage, err)
	}
	image.Close()
	return NodeStorage, nil
}

// qemuDriveParam returns QEMU parameters required to run it with a
// raw-format image at path.
func qemuDriveParam(path string) []string {
	return []string{"-drive", "if=virtio,format=raw,snapshot=off,cache=unsafe,file=" + path}
}

// checkEspContents verifies the presence of the EFI payload inside of image's
// first partition. It returns nil on success.
func checkEspContents(image *disk.Disk) error {
	// Get the ESP.
	fs, err := image.GetFilesystem(1)
	if err != nil {
		return fmt.Errorf("couldn't read the installer ESP: %w", err)
	}
	// Make sure the EFI payload exists by attempting to open it.
	efiPayload, err := fs.OpenFile(osimage.EFIPayloadPath, os.O_RDONLY)
	if err != nil {
		return fmt.Errorf("couldn't open the installer's EFI Payload at %q: %w", osimage.EFIPayloadPath, err)
	}
	efiPayload.Close()
	return nil
}

func TestMain(m *testing.M) {
	// Build the installer image with metroctl, given the EFI executable
	// generated by Metropolis buildsystem. This mimics standard usage of
	// metroctl CLI.
	installer, err := os.Open(InstallerEFIPayload)
	if err != nil {
		log.Fatalf("Couldn't open the installer EFI executable at %q: %s", InstallerEFIPayload, err.Error())
	}
	info, err := installer.Stat()
	if err != nil {
		log.Fatalf("Couldn't stat the installer EFI executable: %s", err.Error())
	}
	bundle, err := os.Open(TestOSBundle)
	if err != nil {
		log.Fatalf("failed to open TestOS bundle: %v", err)
	}
	bundleStat, err := bundle.Stat()
	if err != nil {
		log.Fatalf("failed to stat() TestOS bundle: %v", err)
	}
	iargs := mctl.MakeInstallerImageArgs{
		Installer:     installer,
		InstallerSize: uint64(info.Size()),
		TargetPath:    InstallerImage,
		NodeParams:    &api.NodeParameters{},
		Bundle:        bundle,
		BundleSize:    uint64(bundleStat.Size()),
	}
	if err := mctl.MakeInstallerImage(iargs); err != nil {
		log.Fatalf("Couldn't create the installer image at %q: %s", InstallerImage, err.Error())
	}
	// With common dependencies set up, run the tests.
	code := m.Run()
	// Clean up.
	os.Remove(InstallerImage)
	os.Exit(code)
}

func TestInstallerImage(t *testing.T) {
	// This test examines the installer image, making sure that the GPT and the
	// ESP contents are in order.
	image, err := diskfs.OpenWithMode(InstallerImage, diskfs.ReadOnly)
	if err != nil {
		t.Errorf("Couldn't open the installer image at %q: %s", InstallerImage, err.Error())
	}
	// Verify that GPT exists.
	ti, err := image.GetPartitionTable()
	if ti.Type() != "gpt" {
		t.Error("Couldn't verify that the installer image contains a GPT.")
	}
	// Check that the first partition is likely to be a valid ESP.
	pi := ti.GetPartitions()
	esp := (pi[0]).(*gpt.Partition)
	if esp.Start == 0 || esp.End == 0 {
		t.Error("The installer's ESP GPT entry looks off.")
	}
	// Verify that the image contains only one partition.
	second := (pi[1]).(*gpt.Partition)
	if second.Name != "" || second.Start != 0 || second.End != 0 {
		t.Error("It appears the installer image contains more than one partition.")
	}
	// Verify the ESP contents.
	if err := checkEspContents(image); err != nil {
		t.Error(err.Error())
	}
}

func TestNoBlockDevices(t *testing.T) {
	// No block devices are passed to QEMU aside from the install medium. Expect
	// the installer to fail at the device probe stage rather than attempting to
	// use the medium as the target device.
	expectedOutput := "couldn't find a suitable block device"
	result, err := runQemuWithInstaller(nil, expectedOutput)
	if err != nil {
		t.Error(err.Error())
	}
	if result != true {
		t.Errorf("QEMU didn't produce the expected output %q", expectedOutput)
	}
}

func TestBlockDeviceTooSmall(t *testing.T) {
	// Prepare the block device the installer will install to. This time the
	// target device is too small to host a Metropolis installation.
	imagePath, err := getStorage(64)
	defer os.Remove(imagePath)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Run QEMU. Expect the installer to fail with a predefined error string.
	expectedOutput := "couldn't find a suitable block device"
	result, err := runQemuWithInstaller(qemuDriveParam(imagePath), expectedOutput)
	if err != nil {
		t.Error(err.Error())
	}
	if result != true {
		t.Errorf("QEMU didn't produce the expected output %q", expectedOutput)
	}
}

func TestInstall(t *testing.T) {
	// Prepare the block device image the installer will install to.
	storagePath, err := getStorage(4096 + 128 + 128 + 1)
	defer os.Remove(storagePath)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Run QEMU. Expect the installer to succeed.
	expectedOutput := "Installation completed"
	result, err := runQemuWithInstaller(qemuDriveParam(storagePath), expectedOutput)
	if err != nil {
		t.Error(err.Error())
	}
	if result != true {
		t.Errorf("QEMU didn't produce the expected output %q", expectedOutput)
	}

	// Verify the resulting node image. Check whether the node GPT was created.
	storage, err := diskfs.OpenWithMode(storagePath, diskfs.ReadOnly)
	if err != nil {
		t.Errorf("Couldn't open the resulting node image at %q: %s", storagePath, err.Error())
	}
	// Verify that GPT exists.
	ti, err := storage.GetPartitionTable()
	if ti.Type() != "gpt" {
		t.Error("Couldn't verify that the resulting node image contains a GPT.")
	}
	// Check that the first partition is likely to be a valid ESP.
	pi := ti.GetPartitions()
	esp := (pi[0]).(*gpt.Partition)
	if esp.Name != osimage.ESPVolumeLabel || esp.Start == 0 || esp.End == 0 {
		t.Error("The node's ESP GPT entry looks off.")
	}
	// Verify the system partition's GPT entry.
	system := (pi[1]).(*gpt.Partition)
	if system.Name != osimage.SystemVolumeLabel || system.Start == 0 || system.End == 0 {
		t.Error("The node's system partition GPT entry looks off.")
	}
	// Verify the data partition's GPT entry.
	data := (pi[2]).(*gpt.Partition)
	if data.Name != osimage.DataVolumeLabel || data.Start == 0 || data.End == 0 {
		t.Errorf("The node's data partition GPT entry looks off.")
	}
	// Verify that there are no more partitions.
	fourth := (pi[3]).(*gpt.Partition)
	if fourth.Name != "" || fourth.Start != 0 || fourth.End != 0 {
		t.Error("The resulting node image contains more partitions than expected.")
	}
	// Verify the ESP contents.
	if err := checkEspContents(storage); err != nil {
		t.Error(err.Error())
	}
	// Run QEMU again. Expect TestOS to launch successfully.
	expectedOutput = "_TESTOS_LAUNCH_SUCCESS_"
	result, err = runQemu(qemuDriveParam(storagePath), expectedOutput)
	if err != nil {
		t.Error(err.Error())
	}
	if result != true {
		t.Errorf("QEMU didn't produce the expected output %q", expectedOutput)
	}
}
