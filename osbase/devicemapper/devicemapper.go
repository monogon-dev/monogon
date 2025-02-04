// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// package devicemapper is a thin wrapper for the devicemapper ioctl API.
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/dm-ioctl.h
package devicemapper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
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

// stringToDelimitedBuf copies src to dst and returns an error if len(src) >
// len(dst), or when the string contains a null byte.
func stringToDelimitedBuf(dst []byte, src string) error {
	if len(src) > len(dst)-1 {
		return fmt.Errorf("string longer than target buffer (%v > %v)", len(src), len(dst)-1)
	}
	for i := 0; i < len(src); i++ {
		if src[i] == 0x00 {
			return errors.New("string contains null byte, this is unsupported by DM")
		}
		dst[i] = src[i]
	}
	return nil
}

// marshalParams marshals a list of strings into a single string according to
// the rules in the kernel-side decoder. Strings with null bytes or only
// whitespace characters cannot be encoded and will return an errors.
func marshalParams(params []string) (string, error) {
	var strb strings.Builder
	for _, param := range params {
		var hasNonWhitespace bool
		for i := 0; i < len(param); i++ {
			b := param[i]
			if b == 0x00 {
				return "", errors.New("parameter with null bytes cannot be encoded")
			}
			isWhitespace := ctypeLookup[b]&_S != 0
			if !isWhitespace {
				hasNonWhitespace = true
			}
			if isWhitespace || b == '\\' {
				strb.WriteByte('\\')
			}
			strb.WriteByte(b)
		}
		if !hasNonWhitespace {
			return "", errors.New("parameter with only whitespace cannot be encoded")
		}
		strb.WriteByte(' ')
	}
	return strb.String(), nil
}

var ctrlFile *os.File
var ctrlFileError error
var ctrlFileOnce sync.Once

func initCtrlFile() {
	ctrlFile, ctrlFileError = os.Open("/dev/mapper/control")
	if os.IsNotExist(ctrlFileError) {
		_ = os.MkdirAll("/dev/mapper", 0755)
		ctrlFileError = unix.Mknod("/dev/mapper/control", unix.S_IFCHR|0600, int(unix.Mkdev(10, 236)))
		if ctrlFileError != nil {
			ctrlFileError = fmt.Errorf("devicemapper control device doesn't exist and can't be mknod()ed: %w", ctrlFileError)
			return
		}
		ctrlFile, ctrlFileError = os.Open("/dev/mapper/control")
	}
	if ctrlFileError != nil {
		ctrlFileError = fmt.Errorf("failed to open devicemapper control device: %w", ctrlFileError)
	}
}

func GetVersion() (Version, error) {
	req := newReq()
	ctrlFileOnce.Do(initCtrlFile)
	if ctrlFileError != nil {
		return Version{}, ctrlFileError
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, ctrlFile.Fd(), DM_VERSION_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
		return Version{}, err
	}
	return req.Version, nil
}

func CreateDevice(name string) (uint64, error) {
	req := newReq()
	if err := stringToDelimitedBuf(req.Name[:], name); err != nil {
		return 0, err
	}
	ctrlFileOnce.Do(initCtrlFile)
	if ctrlFileError != nil {
		return 0, ctrlFileError
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, ctrlFile.Fd(), DM_DEV_CREATE_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
		return 0, err
	}
	return req.Dev, nil
}

func RemoveDevice(name string) error {
	req := newReq()
	if err := stringToDelimitedBuf(req.Name[:], name); err != nil {
		return err
	}
	ctrlFileOnce.Do(initCtrlFile)
	if ctrlFileError != nil {
		return ctrlFileError
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, ctrlFile.Fd(), DM_DEV_REMOVE_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
		return err
	}
	runtime.KeepAlive(req)
	return nil
}

// Target represents a byte region inside a devicemapper table for a given
// device provided by a given target implementation.
type Target struct {
	// StartSector is the first sector (defined as being 512 bytes long) this
	// target covers.
	StartSector uint64
	// Length is the number of sectors (defined as being 512 bytes long) this
	// target covers, starting from StartSector.
	Length uint64
	// Type is the type of target handling this byte region.
	// Types implemented by the Linux kernel can be found at
	// @linux//drivers/md/... by looking for dm_register_target() calls.
	Type string
	// Parameters are additional parameters specific to the target type.
	// Note that null bytes and parameters consisting only of whitespace
	// characters cannot be encoded and will return an error.
	Parameters []string
}

func LoadTable(name string, readOnly bool, targets []Target) error {
	req := newReq()
	if err := stringToDelimitedBuf(req.Name[:], name); err != nil {
		return err
	}
	var data bytes.Buffer
	for _, target := range targets {
		encodedParams, err := marshalParams(target.Parameters)
		if err != nil {
			return fmt.Errorf("cannot encode parameters: %w", err)
		}
		// Gives the size of the spec and the null-terminated params aligned to 8 bytes
		padding := len(encodedParams) % 8
		targetSize := uint32(int(unsafe.Sizeof(DMTargetSpec{})) + (len(encodedParams) + 1) + padding)

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
		data.WriteString(encodedParams)
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
	if readOnly {
		req.Flags = DM_READONLY_FLAG
	}
	ctrlFileOnce.Do(initCtrlFile)
	if ctrlFileError != nil {
		return ctrlFileError
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, ctrlFile.Fd(), DM_TABLE_LOAD_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
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
	ctrlFileOnce.Do(initCtrlFile)
	if ctrlFileError != nil {
		return ctrlFileError
	}
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, ctrlFile.Fd(), DM_DEV_SUSPEND_CMD, uintptr(unsafe.Pointer(&req))); err != 0 {
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

func CreateActiveDevice(name string, readOnly bool, targets []Target) (uint64, error) {
	dev, err := CreateDevice(name)
	if err != nil {
		return 0, fmt.Errorf("DM_DEV_CREATE failed: %w", err)
	}
	if err := LoadTable(name, readOnly, targets); err != nil {
		_ = RemoveDevice(name)
		return 0, fmt.Errorf("DM_TABLE_LOAD failed: %w", err)
	}
	if err := Resume(name); err != nil {
		_ = RemoveDevice(name)
		return 0, fmt.Errorf("DM_DEV_SUSPEND failed: %w", err)
	}
	return dev, nil
}
