package nvme

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

// Figure 109
type identifyData struct {
	// Controller Capabilities and Features
	PCIVendorID                 uint16
	PCISubsystemVendorID        uint16
	SerialNumber                [20]byte
	ModelNumber                 [40]byte
	FirmwareRevision            [8]byte
	RecommendedArbitrationBurst uint8
	IEEEOUI                     [3]byte
	CMIC                        uint8
	MaximumDataTransferSize     uint8
	ControllerID                uint16
	Version                     uint32
	RuntimeD3ResumeLatency      uint32
	RuntimeD3EntryLatency       uint32
	OAES                        uint32
	CTRATT                      uint32
	_                           [12]byte
	FRUGUID                     [16]byte
	_                           [128]byte
	// Admin Command Set Attributes & Optional Controller Capabilities
	OACS                                uint16
	AbortCommandLimit                   uint8
	AsynchronousEventRequestLimit       uint8
	FRMW                                uint8
	LPA                                 uint8
	ErrorLogPageEntries                 uint8
	NumberOfPowerStatesSupport          uint8
	AdminVendorSpecificCmdConfig        uint8
	AutonomousPowerStateTransitionAttrs uint8
	WarningCompositeTempThreshold       uint16
	CriticalCompositeTempThreshold      uint16
	MaximumTimeForFirmwareActivation    uint16
	HostMemoryBufferPreferredSize       uint32
	HostMemoryBufferMinimumSize         uint32
	TotalNVMCapacity                    uint128le
	UnallocatedNVMCapacity              uint128le
	ReplyProtectedMemoryBlockSupport    uint32
	ExtendedDeviceSelfTestTime          uint16
	DeviceSelfTestOptions               uint8
	FirmwareUpdateGranularity           uint8
	KeepAliveSupport                    uint16
	HostControlledThermalMgmtAttrs      uint16
	MinimumThermalMgmntTemp             uint16
	MaximumThermalMgmntTemp             uint16
	SanitizeCapabilities                uint32
	_                                   [180]byte
	// NVM Command Set Attributes
	SubmissionQueueEntrySize       uint8
	CompletionQueueEntrySize       uint8
	MaximumOutstandingCommands     uint16
	NumberOfNamespaces             uint32
	OptionalNVMCommandSupport      uint16
	FusedOperationSupport          uint16
	FormatNVMAttributes            uint8
	VolatileWriteCache             uint8
	AtomicWriteUnitNormal          uint16
	AtomicWriteUnitPowerFail       uint16
	NVMVendorSepcificCommandConfig uint8
	AtomicCompareAndWriteUnit      uint16
	_                              [2]byte
	SGLSupport                     uint32
	_                              [228]byte
	NVMSubsystemNVMeQualifiedName  [256]byte
	_                              [1024]byte
	// Power State Descriptors
	PowerStateDescriptors [32][32]byte
}

// IdentifyData contains various identifying information about a NVMe
// controller. Because the actual data structure is very large, currently not
// all fields are exposed as properly-typed individual fields. If you need
// a new field, please add it to this structure.
type IdentifyData struct {
	// PCIVendorID contains the company vendor identifier assigned by the PCI
	// SIG.
	PCIVendorID uint16
	// PCISubsystemVendorID contains the company vendor identifier that is
	// assigned by the PCI SIG for the subsystem.
	PCISubsystemVendorID uint16
	// SerialNumber contains the serial number for the NVM subsystem that is
	// assigned by the vendor.
	SerialNumber string
	// ModelNumber contains the model number for the NVM subsystem that is
	// assigned by the vendor.
	ModelNumber string
	// FirmwareRevision contains the currently active firmware revision for the
	// NVM subsystem.
	FirmwareRevision string
	// IEEEOUI contains the Organization Unique Identifier for the controller
	// vendor as assigned by the IEEE.
	IEEEOUI [3]byte

	// IsPCIVirtualFunction indicates if the controller is a virtual controller
	// as part of a PCI virtual function.
	IsPCIVirtualFunction bool

	// SpecVersionMajor/Minor contain the version of the NVMe specification the
	// controller supports. Only mandatory from spec version 1.2 onwards.
	SpecVersionMajor uint16
	SpecVersionMinor uint8

	// FRUGloballyUniqueIdentifier contains a 128-bit value that is globally
	// unique for a given Field Replaceable Unit (FRU). Contains all-zeroes if
	// unavailable.
	FRUGloballyUniqueIdentifier [16]byte
	// VirtualizationManagementSupported indicates if the controller
	// supports the Virtualization Management command.
	VirtualizationManagementSupported bool
	// NVMeMISupported indicates if the controller supports the NVMe-MI
	// Send and Receive commands.
	NVMeMISupported bool
	// DirectivesSupported indicates if the controller supports the
	// Directive Send and Receive commands.
	DirectivesSupported bool
	// SelfTestSupported indicates if the controller supports the Device Self-
	// test command.
	SelfTestSupported bool
	// NamespaceManagementSupported indicates if the controller supports the
	// Namespace Management and Attachment commands.
	NamespaceManagementSupported bool
	// FirmwareUpdateSupported indicates if the controller supports the
	// Firmware Commit and Image Download commands.
	FirmwareUpdateSupported bool
	// FormattingSupported indicates if the controller supports the Format
	// command.
	FormattingSupported bool
	// SecuritySupported indicates if the controller supports the Security Send
	// and Receive commands.
	SecuritySupported bool

	// TotalNVMCapacity contains the total NVM capacity in bytes in the NVM
	// subsystem. This can be 0 on devices without NamespaceManagementSupported.
	TotalNVMCapacity *big.Int
	// UnallocatedNVMCapacity contains the unallocated NVM capacity in bytes in
	// the NVM subsystem. This can be 0 on devices without
	// NamespaceManagementSupported.
	UnallocatedNVMCapacity *big.Int

	// MaximumNumberOfNamespace defines the maximum number of namespaces
	// supported by the controller.
	MaximumNumberOfNamespaces uint32
}

func (d *Device) Identify() (*IdentifyData, error) {
	var resp [4096]byte

	if err := d.RawCommand(&Command{
		Opcode: 0x06,
		Data:   resp[:],
		CDW10:  1,
	}); err != nil {
		return nil, fmt.Errorf("Identify command failed: %w", err)
	}
	var raw identifyData
	binary.Read(bytes.NewReader(resp[:]), binary.LittleEndian, &raw)

	var res IdentifyData
	res.PCIVendorID = raw.PCIVendorID
	res.PCISubsystemVendorID = raw.PCISubsystemVendorID
	res.SerialNumber = string(bytes.TrimRight(raw.SerialNumber[:], " "))
	res.ModelNumber = string(bytes.TrimRight(raw.ModelNumber[:], " "))
	res.FirmwareRevision = string(bytes.TrimRight(raw.FirmwareRevision[:], " "))
	// OUIs are traditionally big-endian, but NVMe exposes them in little-endian
	res.IEEEOUI[0], res.IEEEOUI[1], res.IEEEOUI[2] = raw.IEEEOUI[2], raw.IEEEOUI[1], raw.IEEEOUI[0]
	res.IsPCIVirtualFunction = raw.CMIC&(1<<2) != 0
	res.SpecVersionMajor = uint16(raw.Version >> 16)
	res.SpecVersionMinor = uint8((raw.Version >> 8) & 0xFF)
	res.FRUGloballyUniqueIdentifier = raw.FRUGUID
	res.VirtualizationManagementSupported = raw.OACS&(1<<7) != 0
	res.NVMeMISupported = raw.OACS&(1<<6) != 0
	res.DirectivesSupported = raw.OACS&(1<<5) != 0
	res.SelfTestSupported = raw.OACS&(1<<4) != 0
	res.NamespaceManagementSupported = raw.OACS&(1<<3) != 0
	res.FirmwareUpdateSupported = raw.OACS&(1<<2) != 0
	res.FormattingSupported = raw.OACS&(1<<1) != 0
	res.SecuritySupported = raw.OACS&(1<<0) != 0

	res.TotalNVMCapacity = raw.TotalNVMCapacity.BigInt()
	res.UnallocatedNVMCapacity = raw.UnallocatedNVMCapacity.BigInt()
	res.MaximumNumberOfNamespaces = raw.NumberOfNamespaces
	return &res, nil
}
