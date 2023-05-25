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

// Package crypt implements block device (eg. disk) encryption and authentication
// using dm-crypt and dm-integrity.
//
// Encryption using dm-crypt is implemented using AES (either in GCM or XTS mode,
// depending on whether authentication is enabled).
//
// Authentication using dm-integrity provides per-sector integrity protection which
// guards against accidental and malicious bit flips in the underlying storage,
// but does nor protect against individual sectors (or the entire disk) being
// rolled back.
//
// The same key is used for both authentication and encryption. The key must be
// exactly 256 bits long.
//
// When initializing or mapping a device, a name must be provided. This name will
// be used as the device-mapper target name if the device will have a
// device-mapper set up, and will also form the base of any intermediary target
// names used. Thus, it must be unique per data store.

package crypt

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Mode of block device encryption and/or authentication, if any. See the
// package-level documentation for information about how encryption and
// authentication is implemented and what guarantees they provide.
type Mode string

// ModeEncryptedAuthenticated means the block device will first be authenticated
// using dm-integrity, then encrypted using dm-crypt.
//
// A key needs to be provided when initializing and mapping a block device.
const ModeEncryptedAuthenticated Mode = "encrypted+authenticated"

// ModeEncrypted means the device will be encrypted using dm-crypt, but will not
// be authenticated.
//
// A key needs to be provided when initializing and mapping a block device.
const ModeEncrypted Mode = "encrypted"

// ModeAuthenticated means the device will be authenticated using dm-integrity,
// but will not be encrypted.
//
// A key needs to be provided when initializing and mapping a block device.
const ModeAuthenticated Mode = "authenticated"

// ModeInsecure means the device will be neither authenticated nor encrypted.
//
// A key must not be provided, or must be exactly zero bytes long.
const ModeInsecure Mode = "insecure"

func (m Mode) encrypted() bool {
	switch m {
	case ModeEncryptedAuthenticated, ModeEncrypted:
		return true
	case ModeInsecure, ModeAuthenticated:
		return false
	}
	panic("invalid mode " + m)
}

func (m Mode) authenticated() bool {
	switch m {
	case ModeEncryptedAuthenticated, ModeAuthenticated:
		return true
	case ModeEncrypted, ModeInsecure:
		return false
	}
	panic("invalid mode " + m)
}

// getSizeBytes returns the size of a block device in bytes.
func getSizeBytes(path string) (uint64, error) {
	blkdev, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed to open block device: %w", err)
	}
	defer blkdev.Close()

	var sizeBytes uint64
	_, _, err = unix.Syscall(unix.SYS_IOCTL, blkdev.Fd(), unix.BLKGETSIZE64, uintptr(unsafe.Pointer(&sizeBytes)))
	if err != unix.Errno(0) {
		return 0, fmt.Errorf("failed to get device size: %w", err)
	}
	return sizeBytes, nil
}

// getBlockSize returns the size of a block device's sector in bytes.
func getBlockSize(path string) (uint32, error) {
	blkdev, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed to open block device: %w", err)
	}
	defer blkdev.Close()

	blockSize, err := unix.IoctlGetUint32(int(blkdev.Fd()), unix.BLKSSZGET)
	if err != nil {
		return 0, fmt.Errorf("BLKSSZGET: %w", err)
	}
	return blockSize, nil
}

// Map sets up an underlying block device (at path 'underlying') for access.
// Depending on the given mode, authentication/integrity device-mapper targets
// will be set up, and the top-level new block device path will be returned.
//
// The given name will be used as a base for the device mapper targets created,
// and is used to uniquely identify this particular mapping setup. The same name
// must then be used to unmap the device.
//
// If an error occurs during Map, cleanup will be attempted and an error will be
// returned.
//
// The encryption key must be exactly 32 bytes / 256 bits long when
// authentication and/or encryption is enabled, and nil / 0 bytes long when
// insecure mode is used.
//
// Note: a successful Map does not necessarily mean the underlying device is
// ready to access. Integrity errors or data corruption might mean accesses to
// the newly mapped device will fail. The caller is responsible for catching
// these conditions.
func Map(name string, underlying string, encryptionKey []byte, mode Mode) (string, error) {
	return map_(name, underlying, encryptionKey, mode, true)
}

// map_ is the internal implementation of Map, which also allows
// enabling/disabling the integrity journal.
//
// This would be called map, but map is a reserved keyword in Go.
func map_(name string, underlying string, encryptionKey []byte, mode Mode, enableJournal bool) (string, error) {
	// Verify key length.
	switch mode {
	case ModeInsecure:
		if len(encryptionKey) != 0 {
			return "", fmt.Errorf("can't use key in insecure mode")
		}
	default:
		if len(encryptionKey) != 32 {
			return "", fmt.Errorf("key must be exactly 32 bytes / 256 bits")
		}
	}

	device := underlying
	if mode.authenticated() {
		var err error
		device, err = mapIntegrity(name, device, enableJournal)
		if err != nil {
			return "", err
		}
	}

	if mode.encrypted() {
		var err error
		device, err = mapEncryption(name, device, encryptionKey, mode.authenticated())
		if err != nil {
			unmapIntegrity(name)
			return "", err
		}
	}

	return device, nil
}

// Unmap tears down all block devices related to the named mapping. The given
// name and mode must match the name and mode used when mapping and/or
// initializing the disk.
func Unmap(name string, mode Mode) error {
	if mode.encrypted() {
		if err := unmapEncryption(name); err != nil {
			return err
		}
	}
	if mode.authenticated() {
		if err := unmapIntegrity(name); err != nil {
			return err
		}
	}
	return nil
}

// Init sets up encryption/authentication as defined by mode on an underlying
// block device path. After initialization, the setup/mapping is preserved and
// the path of the resulting top-level block device is returned.
//
// Any existing data present on the underlying storage will be ignored. If
// authentication is enabled, the underlying storage will also be fully
// overwritten.
//
// The given name will be used as a base for the device mapper targets created,
// and is used to uniquely identify this particular mapping setup. The same name
// must then be used to unmap the device.
//
// The encryption key must be exactly 32 bytes / 256 bits long when
// authentication and/or encryption is enabled, and nil / 0 bytes long when
// insecure mode is used.
func Init(name, underlying string, encryptionKey []byte, mode Mode) (string, error) {
	// If using an authenticated mode, we'll do an initial map with journaling
	// enabled to speed up the initial zeroing, then remap it with journaling.
	// Otherwise, we immediately map with journaling enabled and don't remap.
	initWithJournal := true
	if mode.authenticated() {
		if err := initializeIntegrity(name, underlying); err != nil {
			return "", err
		}
		initWithJournal = false
	}

	device, err := map_(name, underlying, encryptionKey, mode, initWithJournal)
	if err != nil {
		return "", fmt.Errorf("initial mount failed: %w", err)
	}

	// Zero out device if authentication is enabled.
	if mode.authenticated() {
		blockSize, err := getBlockSize(device)
		if err != nil {
			return "", err
		}

		blkdev, err := os.OpenFile(device, unix.O_DIRECT|os.O_WRONLY, 0000)
		if err != nil {
			return "", fmt.Errorf("failed to open new device for zeroing: %w", err)
		}

		// Use a multiple of the block size to make the initial zeroing faster.
		zeroedBuf := make([]byte, blockSize*256)
		for {
			_, err := blkdev.Write(zeroedBuf)
			if errors.Is(err, syscall.ENOSPC) {
				break
			}
			if err != nil {
				blkdev.Close()
				return "", fmt.Errorf("failed to zero-initalize new device: %w", err)
			}
		}
		if err := blkdev.Close(); err != nil {
			return "", fmt.Errorf("failed to close initialized device: %w", err)
		}
	}

	// Remap with journaling if needed.
	if !initWithJournal {
		if err := Unmap(name, mode); err != nil {
			return "", fmt.Errorf("failed to unmap temporary encrypted block device: %w", err)
		}

		device, err = map_(name, underlying, encryptionKey, mode, true)
		if err != nil {
			return "", fmt.Errorf("failed to map initialized encrypted device: %w", err)
		}
	}
	return device, nil
}
