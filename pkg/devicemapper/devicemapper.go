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

package devicemapper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/yalue/native_endian"
	"golang.org/x/sys/unix"
)

type DMIoctl struct {
	Version     Version
	DataSize    uint32
	DataStart   uint32
	TargetCount uint32
	OpenCount   int32
	Flags       uint32
	EventNumber uint32
	_padding1   uint32
	Dev         uint64
	Name        [128]byte
	UUID        [129]byte
	_padding2   [7]byte
	Data        [16384]byte
}

type DMTargetSpec struct {
	SectorStart uint64
	Length      uint64
	Status      int32
	Next        uint32
	TargetType  [16]byte
}

type DMTargetDeps struct {
	Count   uint32
	Padding uint32
	Dev     []uint64
}

type DMNameList struct {
	Dev  uint64
	Next uint32
	Name []byte
}

type DMTargetVersions struct {
	Next    uint32
	Version [3]uint32
}

type DMTargetMessage struct {
	Sector  uint64
	Message []byte
}

type Version [3]uint32

const (
	/* Top level cmds */
	DM_VERSION_CMD uintptr = (0xc138fd << 8) + iota
	DM_REMOVE_ALL_CMD
	DM_LIST_DEVICES_CMD

	/* device level cmds */
	DM_DEV_CREATE_CMD
	DM_DEV_REMOVE_CMD
	DM_DEV_RENAME_CMD
	DM_DEV_SUSPEND_CMD
	DM_DEV_STATUS_CMD
	DM_DEV_WAIT_CMD

	/* Table level cmds */
	DM_TABLE_LOAD_CMD
	DM_TABLE_CLEAR_CMD
	DM_TABLE_DEPS_CMD
	DM_TABLE_STATUS_CMD

	/* Added later */
	DM_LIST_VERSIONS_CMD
	DM_TARGET_MSG_CMD
	DM_DEV_SET_GEOMETRY_CMD
	DM_DEV_ARM_POLL_CMD
)

const (
	DM_READONLY_FLAG       = 1 << 0 /* In/Out */
	DM_SUSPEND_FLAG        = 1 << 1 /* In/Out */
	DM_PERSISTENT_DEV_FLAG = 1 << 3 /* In */
)

const baseDataSize = uint32(unsafe.Sizeof(DMIoctl{})) - 16384

func newReq() DMIoctl {
	return DMIoctl{
		Version:   Version{4, 0, 0},
		DataSize:  baseDataSize,
		DataStart: baseDataSize,
	}
}

func stringToDelimitedBuf(dst []byte, src string) error {
	if len(src) > len(dst)-1 {
		return fmt.Errorf("String longer than target buffer (%v > %v)", len(src), len(dst)-1)
	}
	for i := 0; i < len(src); i++ {
		if src[i] == 0x00 {
			return errors.New("String contains null byte, this is unsupported by DM")
		}
		dst[i] = src[i]
	}
	return nil
}

var fd uintptr

func getFd() (uintptr, error) {
	if fd == 0 {
		f, err := os.Open("/dev/mapper/control")
		if os.IsNotExist(err) {
			os.MkdirAll("/dev/mapper", 0755)
			if err := unix.Mknod("/dev/mapper/control", unix.S_IFCHR|0600, int(unix.Mkdev(10, 236))); err != nil {
				return 0, err
			}
			f, err = os.Open("/dev/mapper/control")
			if err != nil {
				return 0, err
			}
		} else if err != nil {
			return 0, err
		}
		fd = f.Fd()
		return f.Fd(), nil
	}
	return fd, nil
}

func GetVersion() (Version, error) {
	req := newReq()
	fd, err := getFd()
	if err != nil {
		return Version{}, err
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, fd, DM_VERSION_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
		return Version{}, err
	}
	return req.Version, nil
}

func CreateDevice(name string) (uint64, error) {
	req := newReq()
	if err := stringToDelimitedBuf(req.Name[:], name); err != nil {
		return 0, err
	}
	fd, err := getFd()
	if err != nil {
		return 0, err
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, fd, DM_DEV_CREATE_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
		return 0, err
	}
	return req.Dev, nil
}

func RemoveDevice(name string) error {
	req := newReq()
	if err := stringToDelimitedBuf(req.Name[:], name); err != nil {
		return err
	}
	fd, err := getFd()
	if err != nil {
		return err
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, fd, DM_DEV_REMOVE_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
		return err
	}
	runtime.KeepAlive(req)
	return nil
}

type Target struct {
	StartSector uint64
	Length      uint64
	Type        string
	Parameters  string
}

func LoadTable(name string, targets []Target) error {
	req := newReq()
	if err := stringToDelimitedBuf(req.Name[:], name); err != nil {
		return err
	}
	var data bytes.Buffer
	for _, target := range targets {
		// Gives the size of the spec and the null-terminated params aligned to 8 bytes
		padding := len(target.Parameters) % 8
		targetSize := uint32(int(unsafe.Sizeof(DMTargetSpec{})) + (len(target.Parameters) + 1) + padding)

		targetSpec := DMTargetSpec{
			SectorStart: target.StartSector,
			Length:      target.Length,
			Next:        targetSize,
		}
		if err := stringToDelimitedBuf(targetSpec.TargetType[:], target.Type); err != nil {
			return err
		}
		if err := binary.Write(&data, native_endian.NativeEndian(), &targetSpec); err != nil {
			panic(err)
		}
		data.WriteString(target.Parameters)
		data.WriteByte(0x00)
		for i := 0; i < padding; i++ {
			data.WriteByte(0x00)
		}
	}
	req.TargetCount = uint32(len(targets))
	if data.Len() >= 16384 {
		return errors.New("table too large for allocated memory")
	}
	req.DataSize = baseDataSize + uint32(data.Len())
	copy(req.Data[:], data.Bytes())
	fd, err := getFd()
	if err != nil {
		return err
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, fd, DM_TABLE_LOAD_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
		return err
	}
	runtime.KeepAlive(req)
	return nil
}

func suspendResume(name string, suspend bool) error {
	req := newReq()
	if err := stringToDelimitedBuf(req.Name[:], name); err != nil {
		return err
	}
	if suspend {
		req.Flags = DM_SUSPEND_FLAG
	}
	fd, err := getFd()
	if err != nil {
		return err
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, fd, DM_DEV_SUSPEND_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
		return err
	}
	runtime.KeepAlive(req)
	return nil
}

func Suspend(name string) error {
	return suspendResume(name, true)
}
func Resume(name string) error {
	return suspendResume(name, false)
}

func CreateActiveDevice(name string, targets []Target) (uint64, error) {
	dev, err := CreateDevice(name)
	if err != nil {
		return 0, fmt.Errorf("DM_DEV_CREATE failed: %w", err)
	}
	if err := LoadTable(name, targets); err != nil {
		RemoveDevice(name)
		return 0, fmt.Errorf("DM_TABLE_LOAD failed: %w", err)
	}
	if err := Resume(name); err != nil {
		RemoveDevice(name)
		return 0, fmt.Errorf("DM_DEV_SUSPEND failed: %w", err)
	}
	return dev, nil
}
