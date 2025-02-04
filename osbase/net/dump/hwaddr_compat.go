// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package netdump

import (
	"errors"
	"fmt"
	"net"
	"runtime"
	"unsafe"

	"golang.org/x/sys/unix"
)

// @linux//include/uapi/linux:if.h
type ifreq struct {
	ifname [16]byte
	data   uintptr
}

// @linux//include/uapi/linux:ethtool.h ethtool_perm_addr
type ethtoolPermAddr struct {
	Cmd  uint32
	Size uint32
	// Make this an array for memory layout reasons (see
	// comment on the kernel struct)
	Data [32]byte
}

var errNoPermenentHWAddr = errors.New("no permanent hardware address available")

func isAllZeroes(data []byte) bool {
	for _, b := range data {
		if b != 0 {
			return false
		}
	}
	return true
}

// Get permanent hardware address on Linux kernels older than 5.6. On newer
// kernels this is available via normal netlink. Returns errNoPermanentHWAddr
// in case no such address is available.
func getPermanentHWAddrLegacy(ifName string) (net.HardwareAddr, error) {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		fd, err = unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_GENERIC)
		if err != nil {
			return nil, err
		}
	}
	defer unix.Close(fd)

	var ioctlPins runtime.Pinner
	defer ioctlPins.Unpin()

	var data ethtoolPermAddr
	data.Cmd = unix.ETHTOOL_GPERMADDR
	data.Size = uint32(len(data.Data))

	var req ifreq
	copy(req.ifname[:], ifName)
	ioctlPins.Pin(&data)
	req.data = uintptr(unsafe.Pointer(&data))
	for {
		_, _, err := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), unix.SIOCETHTOOL, uintptr(unsafe.Pointer(&req)))
		if err == unix.EINTR {
			continue
		}
		if err != 0 {
			return nil, fmt.Errorf("ioctl(SIOETHTOOL) failed: %w", err)
		}
		break
	}
	runtime.KeepAlive(req)
	runtime.KeepAlive(data)
	// This kernel API is rather old and can indicate the absence of a permanent
	// hardware MAC in two ways: a size of zero (in case the driver does not
	// implement a permanent hardware address at all) or an all-zero value in
	// case the driver has support for returning one but hasn't populated it.
	if data.Size == 0 || isAllZeroes(data.Data[:data.Size]) {
		return nil, errNoPermenentHWAddr
	}
	return data.Data[:data.Size], nil
}
