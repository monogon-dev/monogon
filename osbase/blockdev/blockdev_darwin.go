// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build darwin

package blockdev

import (
	"errors"
	"fmt"
	"math/bits"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

// TODO(lorenz): Upstream these to x/sys/unix.
const (
	DKIOCGETBLOCKSIZE  = 0x40046418
	DKIOCGETBLOCKCOUNT = 0x40086419
)

type Device struct {
	backend    *os.File
	rawConn    syscall.RawConn
	blockSize  int64
	blockCount int64
}

func (d *Device) ReadAt(p []byte, off int64) (n int, err error) {
	return d.backend.ReadAt(p, off)
}

func (d *Device) WriteAt(p []byte, off int64) (n int, err error) {
	return d.backend.WriteAt(p, off)
}

func (d *Device) Close() error {
	return d.backend.Close()
}

func (d *Device) BlockCount() int64 {
	return d.blockCount
}

func (d *Device) BlockSize() int64 {
	return d.blockSize
}

func (d *Device) OptimalBlockSize() int64 {
	return d.blockSize
}

func (d *Device) Discard(startByte int64, endByte int64) error {
	// Can be implemented using DKIOCUNMAP, but needs x/sys/unix extension.
	// Not mandatory, so this is fine for now.
	return errors.ErrUnsupported
}

func (d *Device) Zero(startByte int64, endByte int64) error {
	// It doesn't look like MacOS even has any zeroing acceleration, so just
	// use the generic one.
	return GenericZero(d, startByte, endByte)
}

func (d *Device) Sync() error {
	return d.backend.Sync()
}

// Open opens a block device given a path to its inode.
func Open(path string) (*Device, error) {
	outFile, err := os.OpenFile(path, os.O_RDWR, 0640)
	if err != nil {
		return nil, fmt.Errorf("failed to open block device: %w", err)
	}
	return FromFileHandle(outFile)
}

// FromFileHandle creates a blockdev from a device handle. The device handle is
// not duplicated, closing the returned Device will close it. If the handle is
// not a block device, i.e does not implement block device ioctls, an error is
// returned.
func FromFileHandle(handle *os.File) (*Device, error) {
	outFileC, err := handle.SyscallConn()
	if err != nil {
		return nil, fmt.Errorf("error getting SyscallConn: %w", err)
	}
	var blockSize int
	outFileC.Control(func(fd uintptr) {
		blockSize, err = unix.IoctlGetInt(int(fd), DKIOCGETBLOCKSIZE)
	})
	if errors.Is(err, unix.ENOTTY) || errors.Is(err, unix.EINVAL) {
		return nil, ErrNotBlockDevice
	} else if err != nil {
		return nil, fmt.Errorf("when querying disk block size: %w", err)
	}

	var blockCount int
	var getSizeErr error
	outFileC.Control(func(fd uintptr) {
		blockCount, getSizeErr = unix.IoctlGetInt(int(fd), DKIOCGETBLOCKCOUNT)
	})

	if getSizeErr != nil {
		return nil, fmt.Errorf("when querying disk block count: %w", err)
	}
	return &Device{
		backend:    handle,
		rawConn:    outFileC,
		blockSize:  int64(blockSize),
		blockCount: int64(blockCount),
	}, nil
}

type File struct {
	backend    *os.File
	rawConn    syscall.RawConn
	blockSize  int64
	blockCount int64
}

func CreateFile(name string, blockSize int64, blockCount int64) (*File, error) {
	if blockSize < 512 {
		return nil, fmt.Errorf("blockSize must be bigger than 512 bytes")
	}
	if bits.OnesCount64(uint64(blockSize)) != 1 {
		return nil, fmt.Errorf("blockSize must be a power of two")
	}
	out, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("when creating backing file: %w", err)
	}
	rawConn, err := out.SyscallConn()
	if err != nil {
		return nil, fmt.Errorf("unable to get SyscallConn: %w", err)
	}
	return &File{
		backend:    out,
		blockSize:  blockSize,
		rawConn:    rawConn,
		blockCount: blockCount,
	}, nil
}

func (d *File) ReadAt(p []byte, off int64) (n int, err error) {
	return d.backend.ReadAt(p, off)
}

func (d *File) WriteAt(p []byte, off int64) (n int, err error) {
	return d.backend.WriteAt(p, off)
}

func (d *File) Close() error {
	return d.backend.Close()
}

func (d *File) BlockCount() int64 {
	return d.blockCount
}

func (d *File) BlockSize() int64 {
	return d.blockSize
}

func (d *File) OptimalBlockSize() int64 {
	return d.blockSize
}

func (d *File) Discard(startByte int64, endByte int64) error {
	// Can be supported in the future via fnctl.
	return errors.ErrUnsupported
}

func (d *File) Zero(startByte int64, endByte int64) error {
	// Can possibly be accelerated in the future via fnctl.
	return GenericZero(d, startByte, endByte)
}

func (d *File) Sync() error {
	return d.backend.Sync()
}
