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

// Package quotactl implements a low-level wrapper around the modern portion of
// Linux's quotactl() syscall. See the fsquota package for a nicer interface to
// the most common part of this API.
package quotactl

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type QuotaType uint

const (
	QuotaTypeUser QuotaType = iota
	QuotaTypeGroup
	QuotaTypeProject
)

const (
	Q_SYNC uint = ((0x800001 + iota) << 8)
	Q_QUOTAON
	Q_QUOTAOFF
	Q_GETFMT
	Q_GETINFO
	Q_SETINFO
	Q_GETQUOTA
	Q_SETQUOTA
	Q_GETNEXTQUOTA
)

const (
	FlagBLimitsValid = 1 << iota
	FlagSpaceValid
	FlagILimitsValid
	FlagInodesValid
	FlagBTimeValid
	FlagITimeValid
)

type DQInfo struct {
	Bgrace uint64
	Igrace uint64
	Flags  uint32
	Valid  uint32
}

type Quota struct {
	BHardLimit uint64 // Both Byte limits are prescaled by 1024 (so are in KiB), but CurSpace is in B
	BSoftLimit uint64
	CurSpace   uint64
	IHardLimit uint64
	ISoftLimit uint64
	CurInodes  uint64
	BTime      uint64
	ITime      uint64
	Valid      uint32
}

type NextDQBlk struct {
	HardLimitBytes  uint64
	SoftLimitBytes  uint64
	CurrentBytes    uint64
	HardLimitInodes uint64
	SoftLimitInodes uint64
	CurrentInodes   uint64
	BTime           uint64
	ITime           uint64
	Valid           uint32
	ID              uint32
}

type QuotaFormat uint32

// Collected from quota_format_type structs
const (
	// QuotaFormatNone is a special case where all quota information is
	// stored inside filesystem metadata and thus requires no quotaFilePath.
	QuotaFormatNone   QuotaFormat = 0
	QuotaFormatVFSOld QuotaFormat = 1
	QuotaFormatVFSV0  QuotaFormat = 2
	QuotaFormatOCFS2  QuotaFormat = 3
	QuotaFormatVFSV1  QuotaFormat = 4
)

// QuotaOn turns quota accounting and enforcement on
func QuotaOn(device string, qtype QuotaType, quotaFormat QuotaFormat, quotaFilePath string) error {
	devArg, err := unix.BytePtrFromString(device)
	if err != nil {
		return err
	}
	pathArg, err := unix.BytePtrFromString(quotaFilePath)
	if err != nil {
		return err
	}
	_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_QUOTAON|uint(qtype)), uintptr(unsafe.Pointer(devArg)), uintptr(quotaFormat), uintptr(unsafe.Pointer(pathArg)), 0, 0)
	if err != unix.Errno(0) {
		return err
	}
	return nil
}

// QuotaOff turns quotas off
func QuotaOff(device string, qtype QuotaType) error {
	devArg, err := unix.BytePtrFromString(device)
	if err != nil {
		return err
	}
	_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_QUOTAOFF|uint(qtype)), uintptr(unsafe.Pointer(devArg)), 0, 0, 0, 0)
	if err != unix.Errno(0) {
		return err
	}
	return nil
}

// GetFmt gets the quota format used on given filesystem
func GetFmt(device string, qtype QuotaType) (QuotaFormat, error) {
	var fmt uint32
	devArg, err := unix.BytePtrFromString(device)
	if err != nil {
		return 0, err
	}
	_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_GETFMT|uint(qtype)), uintptr(unsafe.Pointer(devArg)), 0, uintptr(unsafe.Pointer(&fmt)), 0, 0)
	if err != unix.Errno(0) {
		return 0, err
	}
	return QuotaFormat(fmt), nil
}

// GetInfo gets information about quota files
func GetInfo(device string, qtype QuotaType) (*DQInfo, error) {
	var info DQInfo
	devArg, err := unix.BytePtrFromString(device)
	if err != nil {
		return nil, err
	}
	_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_GETINFO|uint(qtype)), uintptr(unsafe.Pointer(devArg)), 0, uintptr(unsafe.Pointer(&info)), 0, 0)
	if err != unix.Errno(0) {
		return nil, err
	}
	return &info, nil
}

// SetInfo sets information about quota files
func SetInfo(device string, qtype QuotaType, info *DQInfo) error {
	devArg, err := unix.BytePtrFromString(device)
	if err != nil {
		return err
	}
	_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_SETINFO|uint(qtype)), uintptr(unsafe.Pointer(devArg)), 0, uintptr(unsafe.Pointer(info)), 0, 0)
	if err != unix.Errno(0) {
		return err
	}
	return nil
}

// GetQuota gets user quota structure
func GetQuota(device string, qtype QuotaType, id uint32) (*Quota, error) {
	var info Quota
	devArg, err := unix.BytePtrFromString(device)
	if err != nil {
		return nil, err
	}
	_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_GETQUOTA|uint(qtype)), uintptr(unsafe.Pointer(devArg)), uintptr(id), uintptr(unsafe.Pointer(&info)), 0, 0)
	if err != unix.Errno(0) {
		return nil, err
	}
	return &info, nil
}

// GetNextQuota gets disk limits and usage > ID
func GetNextQuota(device string, qtype QuotaType, id uint32) (*NextDQBlk, error) {
	var info NextDQBlk
	devArg, err := unix.BytePtrFromString(device)
	if err != nil {
		return nil, err
	}
	_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_GETNEXTQUOTA|uint(qtype)), uintptr(unsafe.Pointer(devArg)), uintptr(id), uintptr(unsafe.Pointer(&info)), 0, 0)
	if err != unix.Errno(0) {
		return nil, err
	}
	return &info, nil
}

// SetQuota sets the given quota
func SetQuota(device string, qtype QuotaType, id uint32, quota *Quota) error {
	devArg, err := unix.BytePtrFromString(device)
	if err != nil {
		return err
	}
	_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_SETQUOTA|uint(qtype)), uintptr(unsafe.Pointer(devArg)), uintptr(id), uintptr(unsafe.Pointer(quota)), 0, 0)
	if err != unix.Errno(0) {
		return fmt.Errorf("failed to set quota: %w", err)
	}
	return nil
}

// Sync syncs disk copy of filesystems quotas. If device is empty it syncs all
// filesystems.
func Sync(device string) error {
	if device != "" {
		devArg, err := unix.BytePtrFromString(device)
		if err != nil {
			return err
		}
		_, _, err = unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_SYNC), uintptr(unsafe.Pointer(devArg)), 0, 0, 0, 0)
		if err != unix.Errno(0) {
			return err
		}
	} else {
		_, _, err := unix.Syscall6(unix.SYS_QUOTACTL, uintptr(Q_SYNC), 0, 0, 0, 0, 0)
		if err != unix.Errno(0) {
			return err
		}
	}
	return nil
}
