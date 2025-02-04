// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package kmod

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// LoadModule loads a kernel module into the kernel.
func LoadModule(file syscall.Conn, params string, flags uintptr) error {
	sc, err := file.SyscallConn()
	if err != nil {
		return fmt.Errorf("failed getting SyscallConn handle: %w", err)
	}
	paramsRaw, err := unix.BytePtrFromString(params)
	if err != nil {
		return errors.New("invalid null byte in params")
	}
	var errNo unix.Errno
	ctrlErr := sc.Control(func(fd uintptr) {
		_, _, errNo = unix.Syscall(unix.SYS_FINIT_MODULE, fd, uintptr(unsafe.Pointer(paramsRaw)), flags)
	})
	if ctrlErr != nil {
		return fmt.Errorf("unable to get control handle: %w", ctrlErr)
	}
	if errNo != unix.Errno(0) {
		return errNo
	}
	return nil
}

// UnloadModule unloads a kernel module from the kernel.
func UnloadModule(name string, flags uintptr) error {
	nameRaw, err := unix.BytePtrFromString(name)
	if err != nil {
		return errors.New("invalid null byte in name")
	}
	_, _, errNo := unix.Syscall(unix.SYS_DELETE_MODULE, uintptr(unsafe.Pointer(nameRaw)), flags, 0)
	if errNo != unix.Errno(0) {
		return errNo
	}
	return nil
}
