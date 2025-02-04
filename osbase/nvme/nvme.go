// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package nvme provides methods and data structures for issuing commands to
// device speaking the NVMe protocol.
// This package is written against the NVMe Specification Revision 1.3 and
// all references to figures or other parts of the spec refer to this version.
package nvme

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
)

// Device is a handle for a NVMe device.
type Device struct {
	fd syscall.Conn
}

// NewFromFd creates a new NVMe device handle from a system handle.
func NewFromFd(fd syscall.Conn) (*Device, error) {
	d := &Device{fd: fd}
	// There is no good way to validate that a file descriptor indeed points to
	// a NVMe device. For future compatibility let this return an error so that
	// code is already prepared to handle it.
	return d, nil
}

// Open opens a new NVMe device handle from a device path (like /dev/nvme0).
func Open(path string) (*Device, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open path: %w", err)
	}
	return NewFromFd(f)
}

// Close closes the NVMe device handle. It returns an error if the handle was
// not created by Open. Please close the handle passed to NewFromFd yourself
// in that case.
func (d *Device) Close() error {
	if f, ok := d.fd.(*os.File); ok {
		return f.Close()
	} else {
		return errors.New("unable to close device not opened via Open, please close it yourself")
	}
}

const (
	// GlobalNamespace is the namespace ID for operations not on a specific
	// namespace.
	GlobalNamespace = 0xffffffff
)

// Command represents a generic NVMe command. Only use this if the command
// you need is not already wrapped by this library.
type Command struct {
	Opcode                                   uint8
	Flags                                    uint8
	NamespaceID                              uint32
	CDW2, CDW3                               uint32
	Metadata                                 []byte
	Data                                     []byte
	CDW10, CDW11, CDW12, CDW13, CDW14, CDW15 uint32
	Timeout                                  time.Duration
}

func (d *Device) GetLogPage(ns uint32, logPageIdentifier uint8, logSpecificField uint8, logPageOffset uint64, pageBuf []byte) error {
	numberOfDwords := len(pageBuf) / 4
	return d.RawCommand(&Command{
		Opcode:      0x02,
		NamespaceID: ns,
		Data:        pageBuf,
		CDW10:       uint32(logPageIdentifier) | uint32(logSpecificField&0xF)<<8 | uint32(numberOfDwords)<<16, // TODO: RAE
		CDW11:       uint32(numberOfDwords >> 16 & 0xffff),
		CDW12:       uint32(logPageOffset & 0xffffffff),
		CDW13:       uint32(logPageOffset >> 32),
	})
}
