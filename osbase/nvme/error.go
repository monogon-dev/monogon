package nvme

import "fmt"

// Figure 31 in the spec
var genericStatusCodeDesc = map[uint8]string{
	0x00: "successful completion",
	0x01: "invalid command opcode",
	0x02: "invalid field in command",
	0x03: "command ID conflict",
	0x04: "data transfer error",
	0x05: "command aborted due power loss notification",
	0x06: "internal error",
	0x07: "command abort requested",
	0x08: "command abort due to SQ deletion",
	0x09: "command abort due to failed fused command",
	0x0a: "command abort due to missing fused command",
	0x0b: "invalid namespace or format",
	0x0c: "command sequence error",
	0x0d: "invalid SGL segment descriptor",
	0x0e: "invalid number of SGL descriptors",
	0x0f: "data SGL length invalid",
	0x10: "metadata SGL length invalid",
	0x11: "SGL descriptor type invalid",
	0x12: "invalid use of controller memory buffer",
	0x13: "PRP offset invalid",
	0x14: "atomic write unit exceeded",
	0x15: "operation denied",
	0x16: "SGL offset invalid",
	0x18: "host identifer inconsistent format",
	0x19: "keep alive timeout expired",
	0x1a: "keep alive timeout invalid",
	0x1b: "command aborted due to preempt and abort",
	0x1c: "sanitize failed",
	0x1d: "sanitize in progress",
	0x1e: "SGL data block granularity invalid",
	0x1f: "command not supported for queue in CMB",

	// Figure 32
	0x80: "LBA out of range",
	0x81: "capacity exceeded",
	0x82: "namespace not ready",
	0x83: "reservation conflict",
	0x84: "format in progress",
}

// Figure 33 in the spec
var commandSpecificStatusCodeDesc = map[uint8]string{
	0x00: "completion queue invalid",
	0x01: "invalid queue identifier",
	0x02: "invalid queue size",
	0x03: "abort command limit exceeded",
	0x05: "asynchronous event request limit exceeded",
	0x06: "invalid firmware slot",
	0x07: "invalid firmware image",
	0x08: "invalid interrupt vector",
	0x09: "invalid log page",
	0x0a: "invalid format",
	0x0b: "firmware activation requires conventional reset",
	0x0c: "invalid queue deletion",
	0x0d: "feature identifier not saveable",
	0x0e: "feature not changeable",
	0x0f: "feature not namespace-specific",
	0x10: "firmware activation requires NVM subsystem reset",
	0x11: "firmware activation requires reset",
	0x12: "firmware activation requires maximum time violation",
	0x13: "firmware activation prohibited",
	0x14: "overlapping range",
	0x15: "namespace insufficient capacity",
	0x16: "namespace identifier unavailable",
	0x18: "namespace already attached",
	0x19: "namespace is private",
	0x1a: "namespace is not attached",
	0x1b: "thin provisioning not supported",
	0x1c: "controller list invalid",
	0x1d: "device self-test in progress",
	0x1e: "boot partition write prohibited",
	0x1f: "invalid controller identifier",
	0x20: "invalid secondary controller state",
	0x21: "invalid number of controller resources",
	0x22: "invalid resource identifier",

	// Figure 34
	0x80: "conflicting attributes",
	0x81: "invalid protection information",
	0x82: "attempted to write to read-only range",
}

// Figure 36
var mediaAndDataIntegrityStatusCodeDesc = map[uint8]string{
	0x80: "write fault",
	0x81: "unrecovered read error",
	0x82: "end-to-end guard check error",
	0x83: "end-to-end application tag check error",
	0x84: "end-to-end reference tag check error",
	0x85: "compare failure",
	0x86: "access denied",
	0x87: "deallocated or unwritten logical block",
}

const (
	StatusCodeTypeGeneric               = 0x0
	StatusCodeTypeCommandSpecific       = 0x1
	StatusCodeTypeMediaAndDataIntegrity = 0x2
)

// Error represents an error returned by the NVMe device in the form of a
// NVMe Status Field (see also Figure 29 in the spec).
type Error struct {
	DoNotRetry     bool
	More           bool
	StatusCodeType uint8
	StatusCode     uint8
}

func (e Error) Error() string {
	switch e.StatusCodeType {
	case StatusCodeTypeGeneric:
		if errStr, ok := genericStatusCodeDesc[e.StatusCode]; ok {
			return errStr
		}
		return fmt.Sprintf("unknown error with generic code 0x%x", e.StatusCode)
	case StatusCodeTypeCommandSpecific:
		if errStr, ok := commandSpecificStatusCodeDesc[e.StatusCode]; ok {
			return errStr
		}
		return fmt.Sprintf("unknown error with command-specific code 0x%x", e.StatusCode)
	case StatusCodeTypeMediaAndDataIntegrity:
		if errStr, ok := mediaAndDataIntegrityStatusCodeDesc[e.StatusCode]; ok {
			return errStr
		}
		return fmt.Sprintf("unknown error with media and data integrity code 0x%x", e.StatusCode)
	default:
		return fmt.Sprintf("unknown error with unknown type 0x%x and code 0x%x", e.StatusCodeType, e.StatusCode)
	}
}
