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

// FS_IOC_FSGETXATTR/FS_IOC_FSSETXATTR are defined in uapi/linux/fs.h
const FS_IOC_FSGETXATTR = 0x801c581f
const FS_IOC_FSSETXATTR = 0x401c5820

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
		return nil, fmt.Errorf("failed to execute getFSXAttrs: %v", errno)
	}
	return &attrs, nil
}

func Set(file *os.File, attrs *FSXAttrs) error {
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, file.Fd(), FS_IOC_FSSETXATTR, uintptr(unsafe.Pointer(attrs)))
	if errno != 0 {
		return fmt.Errorf("failed to execute setFSXAttrs: %v", errno)
	}
	return nil
}
