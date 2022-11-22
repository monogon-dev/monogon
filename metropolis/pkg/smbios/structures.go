package smbios

import (
	"time"
)

const (
	structTypeBIOSInformation      = 0
	structTypeSystemInformation    = 1
	structTypeBaseboardInformation = 2
	structTypeSystemSlot           = 9
	structTypeMemoryDevice         = 17
)

// BIOSInformationRaw contains decoded data from the BIOS Information structure
// (SMBIOS Type 0). See Table 6 in the specification for detailed documentation
// about the individual fields. Note that structure versions 2.1 and 2.2 are
// "invented" here as both characteristics extensions bytes were optional
// between 2.0 and 2.4.
type BIOSInformationRaw struct {
	Handle                                 uint16
	StructureVersion                       Version
	Vendor                                 string
	BIOSVersion                            string
	BIOSStartingAddressSegment             uint16
	BIOSReleaseDate                        string
	BIOSROMSize                            uint8
	BIOSCharacteristics                    uint64
	BIOSCharacteristicsExtensionByte1      uint8 `smbios_min_ver:"2.1"`
	BIOSCharacteristicsExtensionByte2      uint8 `smbios_min_ver:"2.2"`
	SystemBIOSMajorRelease                 uint8 `smbios_min_ver:"2.4"`
	SystemBIOSMinorRelease                 uint8
	EmbeddedControllerFirmwareMajorRelease uint8
	EmbeddedControllerFirmwareMinorRelease uint8
	ExtendedBIOSROMSize                    uint16 `smbios_min_ver:"3.1"`
}

// ROMSizeBytes returns the ROM size in bytes
func (rb *BIOSInformationRaw) ROMSizeBytes() uint64 {
	if rb.StructureVersion.AtLeast(3, 1) && rb.BIOSROMSize == 0xFF {
		// Top 2 bits are SI prefix (starting at mega, i.e. 1024^2), lower 14
		// are value. x*1024^n => x << log2(1024)*n => x << 10*n
		return uint64(rb.ExtendedBIOSROMSize&0x3fff) << 10 * uint64(rb.ExtendedBIOSROMSize&0xc00+2)
	} else {
		// (n+1) * 64KiB
		return (uint64(rb.BIOSROMSize) + 1) * (64 * 1024)
	}
}

// ReleaseDate returns the release date of the BIOS as a time.Time value.
func (rb *BIOSInformationRaw) ReleaseDate() (time.Time, error) {
	return time.Parse("01/02/2006", rb.BIOSReleaseDate)
}

// SystemInformationRaw contains decoded data from the System Information
// structure (SMBIOS Type 1). See Table 10 in the specification for detailed
// documentation about the individual fields.
type SystemInformationRaw struct {
	Handle           uint16
	StructureVersion Version
	Manufacturer     string
	ProductName      string
	Version          string
	SerialNumber     string
	UUID             [16]byte `smbios_min_ver:"2.1"`
	WakeupType       uint8
	SKUNumber        string `smbios_min_ver:"2.4"`
	Family           string
}

// BaseboardInformationRaw contains decoded data from the BIOS Information
// structure (SMBIOS Type 3). See Table 13 in the specification for detailed
// documentation about the individual fields.
type BaseboardInformationRaw struct {
	Handle                         uint16
	StructureVersion               Version
	Manufacturer                   string
	Product                        string
	Version                        string
	SerialNumber                   string
	AssetTag                       string `smbios_min_ver:"2.1"`
	FeatureFlags                   uint8
	LocationInChassis              string
	ChassisHandle                  uint16
	BoardType                      uint8
	NumberOfContainedObjectHandles uint8
	ContainedObjectHandles         []uint16 `smbios_repeat:"NumberOfContainedObjectHandles"`
}

// SystemSlotRaw contains decoded data from the System Slot structure
// (SMBIOS Type 9). See Table 44 in the specification for detailed documentation
// about the individual fields.
type SystemSlotRaw struct {
	Handle               uint16
	StructureVersion     Version
	SlotDesignation      string
	SlotType             uint8
	SlotDataBusWidth     uint8
	CurrentUsage         uint8
	SlotLength           uint8
	SlotID               uint16
	SlotCharacteristics1 uint8
	SlotCharacteristics2 uint8  `smbios_min_ver:"2.1"`
	SegmentGroupNumber   uint16 `smbios_min_ver:"2.6"`
	BusNumber            uint8
	DeviceFunctionNumber uint8
	DataBusWidth         uint8 `smbios_min_ver:"3.2"`
	PeerGroupingCount    uint8
	PeerGroups           []SystemSlotPeerRaw `smbios_repeat:"PeerGroupingCount"`
	SlotInformation      uint8               `smbios_min_ver:"3.4"`
	SlotPhysicalWidth    uint8
	SlotPitch            uint16
	SlotHeight           uint8 `smbios_min_ver:"3.5"`
}

type SystemSlotPeerRaw struct {
	SegmentGroupNumber   uint16
	BusNumber            uint8
	DeviceFunctionNumber uint8
	DataBusWidth         uint8
}

// MemoryDeviceRaw contains decoded data from the BIOS Information structure
// (SMBIOS Type 17). See Table 76 in the specification for detailed
// documentation about the individual fields.
type MemoryDeviceRaw struct {
	Handle                                  uint16
	StructureVersion                        Version
	PhysicalMemoryArrayHandle               uint16 `smbios_min_ver:"2.1"`
	MemoryErrorInformationHandle            uint16
	TotalWidth                              uint16
	DataWidth                               uint16
	Size                                    uint16
	FormFactor                              uint8
	DeviceSet                               uint8
	DeviceLocator                           string
	BankLocator                             string
	MemoryType                              uint8
	TypeDetail                              uint16
	Speed                                   uint16 `smbios_min_ver:"2.3"`
	Manufacturer                            string
	SerialNumber                            string
	AssetTag                                string
	PartNumber                              string
	Attributes                              uint8  `smbios_min_ver:"2.6"`
	ExtendedSize                            uint32 `smbios_min_ver:"2.7"`
	ConfiguredMemorySpeed                   uint16
	MinimumVoltage                          uint16 `smbios_min_ver:"2.8"`
	MaximumVoltage                          uint16
	ConfiguredVoltage                       uint16
	MemoryTechnology                        uint8 `smbios_min_ver:"3.2"`
	MemoryOperatingModeCapability           uint16
	FirmwareVersion                         uint8
	ModuleManufacturerID                    uint16
	ModuleProductID                         uint16
	MemorySubsystemControllerManufacturerID uint16
	MemorySubsystemControllerProductID      uint16
	NonVolatileSize                         uint64
	VolatileSize                            uint64
	CacheSize                               uint64
	LogicalSize                             uint64
	ExtendedSpeed                           uint32 `smbios_min_ver:"3.3"`
	ExtendedConfiguredMemorySpeed           uint32
}

func (md *MemoryDeviceRaw) SizeBytes() (uint64, bool) {
	if md.Size == 0 || md.Size == 0xFFFF {
		// Device unpopulated / unknown memory, return ok false
		return 0, false
	}
	if md.Size == 0x7FFF && md.StructureVersion.AtLeast(2, 7) {
		// Bit 31 is reserved, rest is memory size in MiB
		return uint64(md.ExtendedSize&0x7FFFFFFF) * (1024 * 1024), true
	}
	// Bit 15 flips between KiB and MiB, rest is size
	return uint64(md.Size&0x7FFF) << 10 * uint64(md.Size&0x8000+1), true
}
