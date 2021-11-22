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

// Package loop implements an interface to configure Linux loop devices.
//
// This package requires Linux 5.8 or higher because it uses the newer
// LOOP_CONFIGURE ioctl, which is better-behaved and twice as fast as the old
// approach. It doesn't support all of the cryptloop functionality as it has
// been superseded by dm-crypt and has known vulnerabilities. It also doesn't
// support on-the-fly reconfiguration of loop devices as this is rather
// unusual, works only under very specific circumstances and would make the API
// less clean.
package loop

import (
	"errors"
	"fmt"
	"math/bits"
	"os"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Lazily-initialized file descriptor for the control device /dev/loop-control
// (singleton)
var (
	mutex         sync.Mutex
	loopControlFd *os.File
)

const (
	// LOOP_CONFIGURE from @linux//include/uapi/linux:loop.h
	loopConfigure = 0x4C0A
	// LOOP_MAJOR from @linux//include/uapi/linux:major.h
	loopMajor = 7
)

// struct loop_config from @linux//include/uapi/linux:loop.h
type loopConfig struct {
	fd uint32
	// blockSize is a power of 2 between 512 and os.Getpagesize(), defaults
	// reasonably
	blockSize uint32
	info      loopInfo64
	_reserved [64]byte
}

// struct loop_info64 from @linux//include/uapi/linux:loop.h
type loopInfo64 struct {
	device         uint64
	inode          uint64
	rdevice        uint64
	offset         uint64 // used
	sizeLimit      uint64 // used
	number         uint32
	encryptType    uint32
	encryptKeySize uint32
	flags          uint32   // Flags from Flag constant
	filename       [64]byte // used
	cryptname      [64]byte
	encryptkey     [32]byte
	init           [2]uint64
}

type Config struct {
	// Block size of the loop device in bytes. Power of 2 between 512 and page
	// size.  Zero defaults to an reasonable block size.
	BlockSize uint32
	// Combination of flags from the Flag constants in this package.
	Flags uint32
	// Offset in bytes from the start of the file to the first byte of the
	// device. Usually zero.
	Offset uint64
	// Maximum size of the loop device in bytes. Zero defaults to the whole
	// file.
	SizeLimit uint64
}

func (c *Config) validate() error {
	// Additional validation because of inconsistent kernel-side enforcement
	if c.BlockSize != 0 {
		if c.BlockSize < 512 || c.BlockSize > uint32(os.Getpagesize()) || bits.OnesCount32(c.BlockSize) > 1 {
			return errors.New("BlockSize needs to be a power of two between 512 bytes and the OS page size")
		}
	}
	return nil
}

// ensureFds lazily initializes control devices
func ensureFds() (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	if loopControlFd != nil {
		return
	}
	loopControlFd, err = os.Open("/dev/loop-control")
	return
}

// Device represents a loop device.
type Device struct {
	num uint32
	dev *os.File

	closed bool
}

// All from @linux//include/uapi/linux:loop.h
const (
	// Makes the loop device read-only even if the backing file is read-write.
	FlagReadOnly = 1
	// Unbinds the backing file as soon as the last user is gone. Useful for
	// unbinding after unmount.
	FlagAutoclear = 4
	// Enables kernel-side partition scanning on the loop device. Needed if you
	// want to access specific partitions on a loop device.
	FlagPartscan = 8
	// Enables direct IO for the loop device, bypassing caches and buffer
	// copying.
	FlagDirectIO = 16
)

// Create creates a new loop device backed with the given file.
func Create(f *os.File, c Config) (*Device, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if err := ensureFds(); err != nil {
		return nil, fmt.Errorf("failed to access loop control device: %w", err)
	}
	for {
		devNum, _, errno := syscall.Syscall(unix.SYS_IOCTL, loopControlFd.Fd(), unix.LOOP_CTL_GET_FREE, 0)
		if errno != unix.Errno(0) {
			return nil, fmt.Errorf("failed to allocate loop device: %w", os.NewSyscallError("ioctl(LOOP_CTL_GET_FREE)", errno))
		}
		dev, err := os.OpenFile(fmt.Sprintf("/dev/loop%v", devNum), os.O_RDWR|os.O_EXCL, 0)
		if pe, ok := err.(*os.PathError); ok {
			if pe.Err == unix.EBUSY {
				// We have lost the race, get a new device
				continue
			}
		}
		if err != nil {
			return nil, fmt.Errorf("failed to open newly-allocated loop device: %w", err)
		}

		var config loopConfig
		config.fd = uint32(f.Fd())
		config.blockSize = c.BlockSize
		config.info.flags = c.Flags
		config.info.offset = c.Offset
		config.info.sizeLimit = c.SizeLimit

		if _, _, err := syscall.Syscall(unix.SYS_IOCTL, dev.Fd(), loopConfigure, uintptr(unsafe.Pointer(&config))); err != 0 {
			if err == unix.EBUSY {
				// We have lost the race, get a new device
				continue
			}
			return nil, os.NewSyscallError("ioctl(LOOP_CONFIGURE)", err)
		}
		return &Device{dev: dev, num: uint32(devNum)}, nil
	}
}

// Open opens a loop device at the given path. It returns an error if the path
// is not a loop device.
func Open(path string) (*Device, error) {
	potentialDevice, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open device: %w", err)
	}
	var loopInfo loopInfo64
	_, _, err = syscall.Syscall(unix.SYS_IOCTL, potentialDevice.Fd(), unix.LOOP_GET_STATUS64, uintptr(unsafe.Pointer(&loopInfo)))
	if err == syscall.Errno(0) {
		return &Device{dev: potentialDevice, num: loopInfo.number}, nil
	}
	potentialDevice.Close()
	if err == syscall.EINVAL {
		return nil, errors.New("not a loop device")
	}
	return nil, fmt.Errorf("failed to determine state of potential loop device: %w", err)
}

func (d *Device) ensureOpen() error {
	if d.closed {
		return errors.New("device is closed")
	}
	return nil
}

// DevPath returns the canonical path of this loop device in /dev.
func (d *Device) DevPath() (string, error) {
	if err := d.ensureOpen(); err != nil {
		return "", err
	}
	return fmt.Sprintf("/dev/loop%d", d.num), nil
}

// Dev returns the Linux device ID of the loop device.
func (d *Device) Dev() (uint64, error) {
	if err := d.ensureOpen(); err != nil {
		return 0, err
	}
	return unix.Mkdev(loopMajor, d.num), nil
}

// BackingFilePath returns the path of the backing file
func (d *Device) BackingFilePath() (string, error) {
	backingFile, err := os.ReadFile(fmt.Sprintf("/sys/block/loop%d/loop/backing_file", d.num))
	if err != nil {
		return "", fmt.Errorf("failed to get backing file path: %w", err)
	}
	return string(backingFile), err
}

// RefreshSize recalculates the size of the loop device based on the config and
// the size of the backing file.
func (d *Device) RefreshSize() error {
	if err := d.ensureOpen(); err != nil {
		return err
	}
	return unix.IoctlSetInt(int(d.dev.Fd()), unix.LOOP_SET_CAPACITY, 0)
}

// Close closes all file descriptors open to the device. Does not remove the
// device itself or alter its configuration.
func (d *Device) Close() error {
	if err := d.ensureOpen(); err != nil {
		return err
	}
	d.closed = true
	return d.dev.Close()
}

// Remove removes the loop device.
func (d *Device) Remove() error {
	if err := d.ensureOpen(); err != nil {
		return err
	}
	err := unix.IoctlSetInt(int(d.dev.Fd()), unix.LOOP_CLR_FD, 0)
	if err != nil {
		return err
	}
	if err := d.Close(); err != nil {
		return fmt.Errorf("failed to close device: %w", err)
	}
	if err := unix.IoctlSetInt(int(loopControlFd.Fd()), unix.LOOP_CTL_REMOVE, int(d.num)); err != nil {
		return err
	}
	return nil
}
