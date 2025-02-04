// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package fsxattrs

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

type FSXAttrFlag uint32

// Defined in uapi/linux/fs.h
const (
	FlagRealtime        FSXAttrFlag = 0x00000001
	FlagPreallocated    FSXAttrFlag = 0x00000002
	FlagImmutable       FSXAttrFlag = 0x00000008
	FlagAppend          FSXAttrFlag = 0x00000010
	FlagSync            FSXAttrFlag = 0x00000020
	FlagNoATime         FSXAttrFlag = 0x00000040
	FlagNoDump          FSXAttrFlag = 0x00000080
	FlagRealtimeInherit FSXAttrFlag = 0x00000100
	FlagProjectInherit  FSXAttrFlag = 0x00000200
	FlagNoSymlinks      FSXAttrFlag = 0x00000400
	FlagExtentSize      FSXAttrFlag = 0x00000800
	FlagNoDefragment    FSXAttrFlag = 0x00002000
	FlagFilestream      FSXAttrFlag = 0x00004000
	FlagDAX             FSXAttrFlag = 0x00008000
	FlagCOWExtentSize   FSXAttrFlag = 0x00010000
	FlagHasAttribute    FSXAttrFlag = 0x80000000
)

// FS_IOC_FSGETXATTR and FS_IOC_FSSETXATTR are defined in uapi/linux/fs.h
// and normally would be imported from x/sys/unix. Since they don't exist
// there define them here for now.
const (
	FS_IOC_FSGETXATTR = 0x801c581f
	FS_IOC_FSSETXATTR = 0x401c5820
)

type FSXAttrs struct {
	Flags         FSXAttrFlag
	ExtentSize    uint32
	ExtentCount   uint32
	ProjectID     uint32
	CoWExtentSize uint32
	_pad          [8]byte
}

func Get(file *os.File) (*FSXAttrs, error) {
	var attrs FSXAttrs
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, file.Fd(), FS_IOC_FSGETXATTR, uintptr(unsafe.Pointer(&attrs)))
	if errno != 0 {
		return nil, fmt.Errorf("failed to execute getFSXAttrs: %w", errno)
	}
	return &attrs, nil
}

func Set(file *os.File, attrs *FSXAttrs) error {
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, file.Fd(), FS_IOC_FSSETXATTR, uintptr(unsafe.Pointer(attrs)))
	if errno != 0 {
		return fmt.Errorf("failed to execute setFSXAttrs: %w", errno)
	}
	return nil
}
