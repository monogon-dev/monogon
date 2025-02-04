// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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

	"golang.org/x/sys/unix"
)

// Lazily-initialized file descriptor for the control device /dev/loop-control
// (singleton)
var (
	mutex         sync.Mutex
	loopControlFd *os.File
)

const (
	// LOOP_MAJOR from @linux//include/uapi/linux:major.h
	loopMajor = 7
)

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
		devNum, err := unix.IoctlRetInt(int(loopControlFd.Fd()), unix.LOOP_CTL_GET_FREE)
		if err != nil {
			return nil, fmt.Errorf("failed to allocate loop device: %w", os.NewSyscallError("ioctl(LOOP_CTL_GET_FREE)", err))
		}
		dev, err := os.OpenFile(fmt.Sprintf("/dev/loop%v", devNum), os.O_RDWR|os.O_EXCL, 0)
		if errors.Is(err, unix.EBUSY) {
			// We have lost the race, get a new device
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("failed to open newly-allocated loop device: %w", err)
		}

		var config unix.LoopConfig
		config.Fd = uint32(f.Fd())
		config.Size = c.BlockSize
		config.Info.Flags = c.Flags
		config.Info.Offset = c.Offset
		config.Info.Sizelimit = c.SizeLimit

		err = unix.IoctlLoopConfigure(int(dev.Fd()), &config)
		if errors.Is(err, unix.EBUSY) {
			// We have lost the race, get a new device
			continue
		}
		if err != nil {
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
	loopInfo, err := unix.IoctlLoopGetStatus64(int(potentialDevice.Fd()))
	if err == nil {
		return &Device{dev: potentialDevice, num: loopInfo.Number}, nil
	}
	potentialDevice.Close()
	if errors.Is(err, unix.EINVAL) {
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
