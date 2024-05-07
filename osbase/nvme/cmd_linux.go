//go:build linux

package nvme

import (
	"errors"
	"fmt"
	"math"
	"runtime"
	"unsafe"

	"golang.org/x/sys/unix"
)

// From @linux//include/uapi/linux/nvme_ioctl.h
const (
	nvmeIoctlAdminCmd = 0xC0484E41 // _IOWR('N', 0x41, sizeof cmd)
)

// From @linux//include/uapi/linux/nvme_ioctl.h
type passthruCmd struct {
	// Corresponding to Figure 88
	opcode      uint8
	flags       uint8
	rsvd1       uint16
	nsid        uint32
	cdw2        uint32
	cdw3        uint32
	metadata    uint64
	addr        uint64
	metadataLen uint32
	dataLen     uint32
	cdw10       uint32
	cdw11       uint32
	cdw12       uint32
	cdw13       uint32
	cdw14       uint32
	cdw15       uint32

	// Linux ioctl-specific
	timeoutMs uint32
	result    uint32
}

// RawCommand runs a raw command on the NVMe device.
// Please note that depending on the payload this can be very dangerous and can
// cause data loss or even firmware issues.
func (d *Device) RawCommand(cmd *Command) error {
	conn, err := d.fd.SyscallConn()
	if err != nil {
		return fmt.Errorf("unable to get RawConn: %w", err)
	}
	cmdRaw := passthruCmd{
		opcode:    cmd.Opcode,
		flags:     cmd.Flags,
		nsid:      cmd.NamespaceID,
		cdw2:      cmd.CDW2,
		cdw3:      cmd.CDW3,
		cdw10:     cmd.CDW10,
		cdw11:     cmd.CDW11,
		cdw12:     cmd.CDW12,
		cdw13:     cmd.CDW13,
		cdw14:     cmd.CDW14,
		cdw15:     cmd.CDW15,
		timeoutMs: uint32(cmd.Timeout.Milliseconds()),
	}
	var ioctlPins runtime.Pinner
	defer ioctlPins.Unpin()
	if cmd.Data != nil {
		if len(cmd.Data) > math.MaxUint32 {
			return errors.New("data buffer larger than uint32, this is unsupported")
		}
		ioctlPins.Pin(&cmd.Data[0])
		cmdRaw.dataLen = uint32(len(cmd.Data))
		cmdRaw.addr = uint64(uintptr(unsafe.Pointer(&cmd.Data[0])))
	}
	if cmd.Metadata != nil {
		if len(cmd.Metadata) > math.MaxUint32 {
			return errors.New("metadata buffer larger than uint32, this is unsupported")
		}
		ioctlPins.Pin(&cmd.Metadata[0])
		cmdRaw.metadataLen = uint32(len(cmd.Metadata))
		cmdRaw.metadata = uint64(uintptr(unsafe.Pointer(&cmd.Metadata[0])))
	}
	var errno unix.Errno
	var status uintptr
	err = conn.Control(func(fd uintptr) {
		status, _, errno = unix.Syscall(unix.SYS_IOCTL, fd, nvmeIoctlAdminCmd, uintptr(unsafe.Pointer(&cmdRaw)))
	})
	runtime.KeepAlive(cmdRaw)
	runtime.KeepAlive(cmd.Data)
	runtime.KeepAlive(cmd.Metadata)
	if err != nil {
		return fmt.Errorf("unable to get fd: %w", err)
	}
	if errno != 0 {
		return errno
	}
	var commandErr Error
	commandErr.DoNotRetry = status&(1<<15) != 0            // Bit 31
	commandErr.More = status&(1<<14) != 0                  // Bit 30
	commandErr.StatusCodeType = uint8((status >> 8) & 0x7) // Bits 27:25
	commandErr.StatusCode = uint8(status & 0xff)           // Bits 24:17
	// The only success status is in the generic status code set with value 0
	if commandErr.StatusCodeType != StatusCodeTypeGeneric ||
		commandErr.StatusCode != 0 {
		return commandErr
	}
	return nil
}
