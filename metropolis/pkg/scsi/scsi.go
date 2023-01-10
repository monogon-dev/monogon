// INCITS 502 Revision 19 / SPC-5 R19
package scsi

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"
)

// Device is a handle for a SCSI device
type Device struct {
	fd syscall.Conn
}

// NewFromFd creates a new SCSI device handle from a system handle.
func NewFromFd(fd syscall.Conn) (*Device, error) {
	d := &Device{fd: fd}
	// There is no good way to validate that a file descriptor indeed points to
	// a SCSI device. For future compatibility let this return an error so that
	// code is already prepared to handle it.
	return d, nil
}

// Open creates a new SCSI device handle from a device path (like /dev/sda).
func Open(path string) (*Device, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open path: %w", err)
	}
	return NewFromFd(f)
}

// Close closes the SCSI device handle if opened by Open()
func (d *Device) Close() error {
	if f, ok := d.fd.(*os.File); ok {
		return f.Close()
	} else {
		return errors.New("unable to close device not opened via Open, please close it yourself")
	}
}

type DataTransferDirection uint8

const (
	DataTransferNone DataTransferDirection = iota
	DataTransferToDevice
	DataTransferFromDevice
	DataTransferBidirectional
)

type OperationCode uint8

const (
	InquiryOp        OperationCode = 0x12
	ReadDefectDataOp OperationCode = 0x37
	LogSenseOp       OperationCode = 0x4d
)

// CommandDataBuffer represents a command
type CommandDataBuffer struct {
	// OperationCode contains the code of the command to be called
	OperationCode OperationCode
	// Request contains the OperationCode-specific request parameters
	Request []byte
	// ServiceAction can (for certain CDB encodings) contain an additional
	// qualification for the OperationCode.
	ServiceAction *uint8
	// Control contains common CDB metadata
	Control uint8
	// DataTransferDirection contains the direction(s) of the data transfer(s)
	// to be made.
	DataTransferDirection DataTransferDirection
	// Data contains the data to be transferred. If data needs to be received
	// from the device, a buffer needs to be provided here.
	Data []byte
	// Timeout can contain an optional timeout (0 = no timeout) for the command
	Timeout time.Duration
}

// Bytes returns the raw CDB to be sent to the device
func (c *CommandDataBuffer) Bytes() ([]byte, error) {
	// Table 24
	switch {
	case c.OperationCode < 0x20:
		// Use CDB6 as defined in Table 3
		if c.ServiceAction != nil {
			return nil, errors.New("ServiceAction field not available in CDB6")
		}
		if len(c.Request) != 4 {
			return nil, fmt.Errorf("CDB6 request size is %d bytes, needs to be 4 bytes without LengthField", len(c.Request))
		}

		outBuf := make([]byte, 6)
		outBuf[0] = uint8(c.OperationCode)

		copy(outBuf[1:5], c.Request)
		outBuf[5] = c.Control
		return outBuf, nil
	case c.OperationCode < 0x60:
		// Use CDB10 as defined in Table 5
		if len(c.Request) != 8 {
			return nil, fmt.Errorf("CDB10 request size is %d bytes, needs to be 4 bytes", len(c.Request))
		}

		outBuf := make([]byte, 10)
		outBuf[0] = uint8(c.OperationCode)
		copy(outBuf[1:9], c.Request)
		if c.ServiceAction != nil {
			outBuf[1] |= *c.ServiceAction & 0b11111
		}
		outBuf[9] = c.Control
		return outBuf, nil
	case c.OperationCode < 0x7e:
		return nil, errors.New("OperationCode is reserved")
	case c.OperationCode == 0x7e:
		// Use variable extended
		return nil, errors.New("variable extended CDBs are unimplemented")
	case c.OperationCode == 0x7f:
		// Use variable
		return nil, errors.New("variable CDBs are unimplemented")
	case c.OperationCode < 0xa0:
		// Use CDB16 as defined in Table 13
		if len(c.Request) != 14 {
			return nil, fmt.Errorf("CDB16 request size is %d bytes, needs to be 14 bytes", len(c.Request))
		}

		outBuf := make([]byte, 16)
		outBuf[0] = uint8(c.OperationCode)
		copy(outBuf[1:15], c.Request)
		if c.ServiceAction != nil {
			outBuf[1] |= *c.ServiceAction & 0b11111
		}
		outBuf[15] = c.Control
		return outBuf, nil
	case c.OperationCode < 0xc0:
		// Use CDB12 as defined in Table 7
		if len(c.Request) != 10 {
			return nil, fmt.Errorf("CDB12 request size is %d bytes, needs to be 10 bytes", len(c.Request))
		}

		outBuf := make([]byte, 12)
		outBuf[0] = uint8(c.OperationCode)
		copy(outBuf[1:11], c.Request)
		if c.ServiceAction != nil {
			outBuf[1] |= *c.ServiceAction & 0b11111
		}
		outBuf[11] = c.Control
		return outBuf, nil
	default:
		return nil, errors.New("unable to encode CDB for given OperationCode")
	}
}

// SenseKey represents the top-level status code of a SCSI sense response.
type SenseKey uint8

const (
	NoSense        SenseKey = 0x0
	RecoveredError SenseKey = 0x1
	NotReady       SenseKey = 0x2
	MediumError    SenseKey = 0x3
	HardwareError  SenseKey = 0x4
	IllegalRequest SenseKey = 0x5
	UnitAttention  SenseKey = 0x6
	DataProtect    SenseKey = 0x7
	BlankCheck     SenseKey = 0x8
	VendorSpecific SenseKey = 0x9
	CopyAborted    SenseKey = 0xa
	AbortedCommand SenseKey = 0xb
	VolumeOverflow SenseKey = 0xd
	Miscompare     SenseKey = 0xe
	Completed      SenseKey = 0xf
)

var senseKeyDesc = map[SenseKey]string{
	NoSense:        "no sense information",
	RecoveredError: "recovered error",
	NotReady:       "not ready",
	MediumError:    "medium error",
	HardwareError:  "hardware error",
	IllegalRequest: "illegal request",
	UnitAttention:  "unit attention",
	DataProtect:    "data protected",
	BlankCheck:     "blank check failed",
	VendorSpecific: "vendor-specific error",
	CopyAborted:    "third-party copy aborted",
	AbortedCommand: "command aborted",
	VolumeOverflow: "volume overflow",
	Miscompare:     "miscompare",
	Completed:      "completed",
}

func (s SenseKey) String() string {
	if str, ok := senseKeyDesc[s]; ok {
		return str
	}
	return fmt.Sprintf("sense key %xh", uint8(s))
}

// AdditionalSenseCode contains the additional sense key and qualifier in one
// 16-bit value. The high 8 bits are the sense key, the bottom 8 bits the
// qualifier.
type AdditionalSenseCode uint16

// ASK returns the raw Additional Sense Key
func (a AdditionalSenseCode) ASK() uint8 {
	return uint8(a >> 8)
}

// ASKQ returns the raw Additional Sense Key Qualifier
func (a AdditionalSenseCode) ASKQ() uint8 {
	return uint8(a & 0xFF)
}

// IsKey checks if the ASK portion of a is the same as the ASK portion of b.
func (a AdditionalSenseCode) IsKey(b AdditionalSenseCode) bool {
	return a.ASK() == b.ASK()
}

// String returns the textual representation of this ASK
func (s AdditionalSenseCode) String() string {
	if str, ok := additionalSenseCodeDesc[s]; ok {
		return str
	}
	return fmt.Sprintf("unknown additional sense code %xh %xh", s.ASK(), s.ASKQ())
}

// FixedError is one type of error returned by a SCSI CHECK_CONDITION.
// See also Table 48 in the standard.
type FixedError struct {
	Deferred                   bool
	SenseKey                   SenseKey
	Information                uint32
	CommandSpecificInformation uint32
	AdditionalSenseCode        AdditionalSenseCode
}

func (e FixedError) Error() string {
	if e.AdditionalSenseCode == 0 {
		return fmt.Sprintf("%v", e.SenseKey)
	}
	return fmt.Sprintf("%v: %v", e.SenseKey, e.AdditionalSenseCode)

}

// UnknownError is a type of error returned by SCSI which is not understood by this
// library. This can be a vendor-specific or future error.
type UnknownError struct {
	RawSenseData []byte
}

func (e *UnknownError) Error() string {
	return fmt.Sprintf("unknown SCSI error, raw sense data follows: %x", e.RawSenseData)
}
