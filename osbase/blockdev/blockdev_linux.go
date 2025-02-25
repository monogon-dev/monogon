// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package blockdev

import (
	"errors"
	"fmt"
	"io"
	"math/bits"
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Device struct {
	backend    *os.File
	rawConn    syscall.RawConn
	blockSize  int64
	blockCount int64
}

func (d *Device) ReadAt(p []byte, off int64) (n int, err error) {
	size := d.blockSize * d.blockCount
	if off > size {
		return 0, io.EOF
	}
	if int64(len(p)) > size-off {
		n, err = d.backend.ReadAt(p[:size-off], off)
		if err == nil {
			err = io.EOF
		}
		return
	}
	return d.backend.ReadAt(p, off)
}

func (d *Device) WriteAt(p []byte, off int64) (n int, err error) {
	size := d.blockSize * d.blockCount
	if off > size {
		return 0, ErrOutOfBounds
	}
	if int64(len(p)) > size-off {
		n, err = d.backend.WriteAt(p[:size-off], off)
		if err == nil {
			err = ErrOutOfBounds
		}
		return
	}
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
	if err := validAlignedRange(d, startByte, endByte); err != nil {
		return err
	}
	if startByte == endByte {
		return nil
	}
	var args [2]uint64
	var err unix.Errno
	args[0] = uint64(startByte)
	args[1] = uint64(endByte - startByte)
	if ctrlErr := d.rawConn.Control(func(fd uintptr) {
		_, _, err = unix.Syscall(unix.SYS_IOCTL, fd, unix.BLKDISCARD, uintptr(unsafe.Pointer(&args[0])))
	}); ctrlErr != nil {
		return ctrlErr
	}
	if err == unix.EOPNOTSUPP {
		return errors.ErrUnsupported
	}
	if err != unix.Errno(0) {
		return fmt.Errorf("failed to discard: %w", err)
	}
	return nil
}

func (d *Device) Zero(startByte int64, endByte int64) error {
	if err := validAlignedRange(d, startByte, endByte); err != nil {
		return err
	}
	if startByte == endByte {
		return nil
	}
	var args [2]uint64
	var err error
	args[0] = uint64(startByte)
	args[1] = uint64(endByte - startByte)
	ctrlErr := d.rawConn.Control(func(fd uintptr) {
		// Attempts to leverage discard guarantees to provide extremely quick
		// metadata-only zeroing.
		err = unix.Fallocate(int(fd), unix.FALLOC_FL_PUNCH_HOLE|unix.FALLOC_FL_KEEP_SIZE, startByte, endByte-startByte)
		if errors.Is(err, unix.EOPNOTSUPP) {
			// Tries Write Same and friends and then just falls back to writing
			// zeroes.
			_, _, errNo := unix.Syscall(unix.SYS_IOCTL, fd, unix.BLKZEROOUT, uintptr(unsafe.Pointer(&args[0])))
			if errNo == unix.Errno(0) {
				err = nil
			} else {
				err = errNo
			}
		}
	})
	if ctrlErr != nil {
		return ctrlErr
	}
	if err != nil {
		return fmt.Errorf("failed to zero out: %w", err)
	}
	return nil
}

func (d *Device) Sync() error {
	return d.backend.Sync()
}

// RefreshPartitionTable refreshes the kernel's view of the partition table
// after changes made from userspace.
func (d *Device) RefreshPartitionTable() error {
	var err unix.Errno
	if ctrlErr := d.rawConn.Control(func(fd uintptr) {
		_, _, err = unix.Syscall(unix.SYS_IOCTL, fd, unix.BLKRRPART, 0)
	}); ctrlErr != nil {
		return ctrlErr
	}
	if err != unix.Errno(0) {
		return fmt.Errorf("ioctl(BLKRRPART): %w", err)
	}
	return nil
}

// ResizePartition updates the start and length of one partition in the kernel.
// This can be used as an alternative to RefreshPartitionTable, which cannot
// be used if any partition on this device is currently mounted.
func (d *Device) ResizePartition(partitionNo int32, startByte, lengthBytes int64) error {
	var ioctlPins runtime.Pinner
	defer ioctlPins.Unpin()

	partition := unix.BlkpgPartition{
		Start:  startByte,
		Length: lengthBytes,
		Pno:    partitionNo,
	}
	ioctlPins.Pin(&partition)
	arg := unix.BlkpgIoctlArg{
		Op:      unix.BLKPG_RESIZE_PARTITION,
		Datalen: int32(unsafe.Sizeof(partition)),
		Data:    (*byte)(unsafe.Pointer(&partition)),
	}

	var err unix.Errno
	if ctrlErr := d.rawConn.Control(func(fd uintptr) {
		_, _, err = unix.Syscall(unix.SYS_IOCTL, fd, unix.BLKPG, uintptr(unsafe.Pointer(&arg)))
	}); ctrlErr != nil {
		return ctrlErr
	}
	if err != unix.Errno(0) {
		return fmt.Errorf("ioctl(BLKPG): %w", err)
	}
	return nil
}

// Open opens a block device given a path to its inode.
func Open(path string, opts ...Option) (*Device, error) {
	var o options
	o.collect(opts)
	flags := o.genericFlags()
	if o.direct {
		flags |= unix.O_DIRECT
	}
	if o.exclusive {
		flags |= unix.O_EXCL
	}

	outFile, err := os.OpenFile(path, flags, 0640)
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
	var blockSize uint32
	outFileC.Control(func(fd uintptr) {
		blockSize, err = unix.IoctlGetUint32(int(fd), unix.BLKSSZGET)
	})
	if errors.Is(err, unix.ENOTTY) || errors.Is(err, unix.EINVAL) {
		return nil, ErrNotBlockDevice
	} else if err != nil {
		return nil, fmt.Errorf("when querying disk block size: %w", err)
	}

	var sizeBytes uint64
	var getSizeErr syscall.Errno
	outFileC.Control(func(fd uintptr) {
		_, _, getSizeErr = unix.Syscall(unix.SYS_IOCTL, fd, unix.BLKGETSIZE64, uintptr(unsafe.Pointer(&sizeBytes)))
	})

	if getSizeErr != unix.Errno(0) {
		return nil, fmt.Errorf("when querying disk block count: %w", err)
	}
	if sizeBytes%uint64(blockSize) != 0 {
		return nil, fmt.Errorf("block device size is not an integer multiple of its block size (%d %% %d = %d)", sizeBytes, blockSize, sizeBytes%uint64(blockSize))
	}
	return &Device{
		backend:    handle,
		rawConn:    outFileC,
		blockSize:  int64(blockSize),
		blockCount: int64(sizeBytes) / int64(blockSize),
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
		return nil, fmt.Errorf("blockSize must be at least 512 bytes")
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
	size := d.blockSize * d.blockCount
	if off > size {
		return 0, io.EOF
	}
	if int64(len(p)) > size-off {
		n, err = d.backend.ReadAt(p[:size-off], off)
		if err == nil {
			err = io.EOF
		}
		return
	}
	return d.backend.ReadAt(p, off)
}

func (d *File) WriteAt(p []byte, off int64) (n int, err error) {
	size := d.blockSize * d.blockCount
	if off > size {
		return 0, ErrOutOfBounds
	}
	if int64(len(p)) > size-off {
		n, err = d.backend.WriteAt(p[:size-off], off)
		if err == nil {
			err = ErrOutOfBounds
		}
		return
	}
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
	if err := validAlignedRange(d, startByte, endByte); err != nil {
		return err
	}
	if startByte == endByte {
		return nil
	}
	var err error
	if ctrlErr := d.rawConn.Control(func(fd uintptr) {
		// There is FALLOC_FL_NO_HIDE_STALE, but it's not implemented by
		// any filesystem right now, so let's not attempt it for the time being.
		err = unix.Fallocate(int(fd), unix.FALLOC_FL_PUNCH_HOLE|unix.FALLOC_FL_KEEP_SIZE, startByte, endByte-startByte)
	}); ctrlErr != nil {
		return ctrlErr
	}
	if errors.Is(err, unix.EOPNOTSUPP) {
		return errors.ErrUnsupported
	}
	if err != nil {
		return fmt.Errorf("failed to discard: %w", err)
	}
	return nil
}

func (d *File) Zero(startByte int64, endByte int64) error {
	if err := validAlignedRange(d, startByte, endByte); err != nil {
		return err
	}
	if startByte == endByte {
		return nil
	}
	var err error
	if ctrlErr := d.rawConn.Control(func(fd uintptr) {
		// Tell the filesystem to punch out the given blocks.
		err = unix.Fallocate(int(fd), unix.FALLOC_FL_PUNCH_HOLE|unix.FALLOC_FL_KEEP_SIZE, startByte, endByte-startByte)
	}); ctrlErr != nil {
		return ctrlErr
	}
	// If unsupported or the syscall is not available (for example in a sandbox)
	// fall back to the generic software implementation.
	if errors.Is(err, unix.EOPNOTSUPP) || errors.Is(err, unix.ENOSYS) {
		return GenericZero(d, startByte, endByte)
	}
	if err != nil {
		return fmt.Errorf("failed to zero out: %w", err)
	}
	return nil
}

func (d *File) Sync() error {
	return d.backend.Sync()
}
