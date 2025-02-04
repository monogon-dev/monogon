// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package scsi

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"runtime"
	"unsafe"

	"golang.org/x/sys/unix"
)

// RawCommand issues a raw command against the device.
func (d *Device) RawCommand(c *CommandDataBuffer) error {
	cdb, err := c.Bytes()
	if err != nil {
		return fmt.Errorf("error encoding CDB: %w", err)
	}
	conn, err := d.fd.SyscallConn()
	if err != nil {
		return fmt.Errorf("unable to get RawConn: %w", err)
	}
	var dxferDir int32
	switch c.DataTransferDirection {
	case DataTransferNone:
		dxferDir = SG_DXFER_NONE
	case DataTransferFromDevice:
		dxferDir = SG_DXFER_FROM_DEV
	case DataTransferToDevice:
		dxferDir = SG_DXFER_TO_DEV
	case DataTransferBidirectional:
		dxferDir = SG_DXFER_TO_FROM_DEV
	default:
		return errors.New("invalid DataTransferDirection")
	}
	var timeout uint32
	if c.Timeout.Milliseconds() > math.MaxUint32 {
		timeout = math.MaxUint32
	}
	if len(c.Data) > math.MaxUint32 {
		return errors.New("payload larger than 2^32 bytes, unable to issue")
	}
	if len(cdb) > math.MaxUint8 {
		return errors.New("CDB larger than 2^8 bytes, unable to issue")
	}
	var senseBuf [32]byte

	var ioctlPins runtime.Pinner
	ioctlPins.Pin(&c.Data[0])
	ioctlPins.Pin(&cdb[0])
	ioctlPins.Pin(&senseBuf[0])
	defer ioctlPins.Unpin()

	cmdRaw := sgIOHdr{
		Interface_id:    'S',
		Dxfer_direction: dxferDir,
		Dxfer_len:       uint32(len(c.Data)),
		Dxferp:          uintptr(unsafe.Pointer(&c.Data[0])),
		Cmd_len:         uint8(len(cdb)),
		Cmdp:            uintptr(unsafe.Pointer(&cdb[0])),
		Mx_sb_len:       uint8(len(senseBuf)),
		Sbp:             uintptr(unsafe.Pointer(&senseBuf[0])),
		Timeout:         timeout,
	}
	var errno unix.Errno
	err = conn.Control(func(fd uintptr) {
		_, _, errno = unix.Syscall(unix.SYS_IOCTL, fd, SG_IO, uintptr(unsafe.Pointer(&cmdRaw)))
	})
	runtime.KeepAlive(cmdRaw)
	runtime.KeepAlive(c.Data)
	runtime.KeepAlive(senseBuf)
	runtime.KeepAlive(cdb)
	if err != nil {
		return fmt.Errorf("unable to get fd: %w", err)
	}
	if errno != 0 {
		return errno
	}
	if cmdRaw.Masked_status != 0 {
		if senseBuf[0] == 0x70 || senseBuf[0] == 0x71 {
			err := &FixedError{
				Deferred:    senseBuf[0] == 0x71,
				SenseKey:    SenseKey(senseBuf[2] & 0b1111),
				Information: binary.BigEndian.Uint32(senseBuf[3:7]),
			}
			length := int(senseBuf[7])
			if length >= 4 {
				err.CommandSpecificInformation = binary.BigEndian.Uint32(senseBuf[8:12])
				if length >= 6 {
					err.AdditionalSenseCode = AdditionalSenseCode(uint16(senseBuf[12])<<8 | uint16(senseBuf[13]))
				}
			}
			return err
		}
		return &UnknownError{
			RawSenseData: senseBuf[:],
		}
	}
	return nil
}
