package scsi

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
)

// Inquiry queries the device for various metadata about its identity and
// supported features.
func (d Device) Inquiry() (*InquiryData, error) {
	data := make([]byte, 96)
	var req [4]byte
	binary.BigEndian.PutUint16(req[2:4], uint16(len(data)))
	if err := d.RawCommand(&CommandDataBuffer{
		OperationCode:         InquiryOp,
		Request:               req[:],
		Data:                  data,
		DataTransferDirection: DataTransferFromDevice,
	}); err != nil {
		return nil, fmt.Errorf("error during INQUIRY: %w", err)
	}
	resLen := int64(data[4]) + 5
	// Use LimitReader to not have to deal with out-of-bounds slices
	rawReader := io.LimitReader(bytes.NewReader(data), resLen)
	var raw inquiryDataRaw
	if err := binary.Read(rawReader, binary.BigEndian, &raw); err != nil {
		if errors.Is(err, io.ErrUnexpectedEOF) {
			return nil, fmt.Errorf("response to INQUIRY is smaller than %d bytes, very old or broken device", binary.Size(raw))
		}
		panic(err) // Read from memory, shouldn't be possible to hit
	}

	var res InquiryData
	res.PeriperalQualifier = (raw.PeripheralData >> 5) & 0b111
	res.PeripheralDeviceType = DeviceType(raw.PeripheralData & 0b11111)
	res.RemovableMedium = (raw.Flags1 & 1 << 0) != 0
	res.LogicalUnitConglomerate = (raw.Flags1 & 1 << 1) != 0
	res.CommandSetVersion = Version(raw.Version)
	res.NormalACASupported = (raw.Flags2 & 1 << 5) != 0
	res.HistoricalSupport = (raw.Flags2 & 1 << 4) != 0
	res.ResponseDataFormat = raw.Flags2 & 0b1111
	res.SCCSupported = (raw.Flags3 & 1 << 7) != 0
	res.TargetPortGroupSupport = (raw.Flags3 >> 4) & 0b11
	res.ThirdPartyCopySupport = (raw.Flags3 & 1 << 3) != 0
	res.HasProtectionInfo = (raw.Flags3 & 1 << 0) != 0
	res.HasEnclosureServices = (raw.Flags4 & 1 << 6) != 0
	res.VendorFeature1 = (raw.Flags4 & 1 << 5) != 0
	res.HasMultipleSCSIPorts = (raw.Flags4 & 1 << 4) != 0
	res.CmdQueue = (raw.Flags5 & 1 << 1) != 0
	res.VendorFeature2 = (raw.Flags5 & 1 << 0) != 0
	res.Vendor = string(bytes.TrimRight(raw.Vendor[:], " "))
	res.Product = string(bytes.TrimRight(raw.Product[:], " "))
	res.ProductRevisionLevel = string(bytes.TrimRight(raw.ProductRevisionLevel[:], " "))

	// Read rest conditionally, as it might not be present on every device
	var vendorSpecific bytes.Buffer
	_, err := io.CopyN(&vendorSpecific, rawReader, 20)
	res.VendorSpecific = vendorSpecific.Bytes()
	if err == io.EOF {
		return &res, nil
	}
	if err != nil {
		panic(err) // Mem2Mem copy, can't really happen
	}
	var padding [2]byte
	if _, err := io.ReadFull(rawReader, padding[:]); err != nil {
		if errors.Is(err, io.ErrUnexpectedEOF) {
			return &res, nil
		}
	}
	for i := 0; i < 8; i++ {
		var versionDesc uint16
		if err := binary.Read(rawReader, binary.BigEndian, &versionDesc); err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
				return &res, nil
			}
		}
		res.VersionDescriptors = append(res.VersionDescriptors, versionDesc)
	}

	return &res, nil
}

// Table 148, only first 36 mandatory bytes
type inquiryDataRaw struct {
	PeripheralData       uint8
	Flags1               uint8
	Version              uint8
	Flags2               uint8
	AdditionalLength     uint8 // n-4
	Flags3               uint8
	Flags4               uint8
	Flags5               uint8
	Vendor               [8]byte
	Product              [16]byte
	ProductRevisionLevel [4]byte
}

// DeviceType represents a SCSI peripheral device type, which
// can be used to determine the command set to use to control
// the device. See Table 150 in the standard.
type DeviceType uint8

const (
	TypeBlockDevice              DeviceType = 0x00
	TypeSequentialAccessDevice   DeviceType = 0x01
	TypeProcessor                DeviceType = 0x03
	TypeOpticalDrive             DeviceType = 0x05
	TypeOpticalMemory            DeviceType = 0x07
	TypeMediaChanger             DeviceType = 0x08
	TypeArrayController          DeviceType = 0x0c
	TypeEncloseServices          DeviceType = 0x0d
	TypeOpticalCardRWDevice      DeviceType = 0x0f
	TypeObjectStorageDevice      DeviceType = 0x11
	TypeAutomationDriveInterface DeviceType = 0x12
	TypeZonedBlockDevice         DeviceType = 0x14
	TypeUnknownDevice            DeviceType = 0x1f
)

var deviceTypeDesc = map[DeviceType]string{
	TypeBlockDevice:              "Block Device",
	TypeSequentialAccessDevice:   "Sequential Access Device",
	TypeProcessor:                "Processor",
	TypeOpticalDrive:             "Optical Drive",
	TypeOpticalMemory:            "Optical Memory",
	TypeMediaChanger:             "Media Changer",
	TypeArrayController:          "Array Controller",
	TypeEncloseServices:          "Enclosure Services",
	TypeOpticalCardRWDevice:      "Optical Card reader/writer device",
	TypeObjectStorageDevice:      "Object-based Storage Device",
	TypeAutomationDriveInterface: "Automation/Drive Interface",
	TypeZonedBlockDevice:         "Zoned Block Device",
	TypeUnknownDevice:            "Unknown or no device",
}

func (d DeviceType) String() string {
	if str, ok := deviceTypeDesc[d]; ok {
		return str
	}
	return fmt.Sprintf("unknown device type %xh", uint8(d))
}

// Version represents a specific standardized version of the SCSI
// primary command set (SPC). The enum values are sorted, so
// for example version >= SPC3 is true for SPC-3 and all later
// standards. See table 151 in the standard.
type Version uint8

const (
	SPC1 = 0x03
	SPC2 = 0x04
	SPC3 = 0x05
	SPC4 = 0x06
	SPC5 = 0x07
)

var versionDesc = map[Version]string{
	SPC1: "SPC-1 (INCITS 301-1997)",
	SPC2: "SPC-2 (INCITS 351-2001)",
	SPC3: "SPC-3 (INCITS 408-2005)",
	SPC4: "SPC-4 (INCITS 513-2015)",
	SPC5: "SPC-5 (INCITS 502-2019)",
}

func (v Version) String() string {
	if str, ok := versionDesc[v]; ok {
		return str
	}
	return fmt.Sprintf("unknown version %xh", uint8(v))
}

// InquiryData contains data returned by the INQUIRY command.
type InquiryData struct {
	PeriperalQualifier      uint8
	PeripheralDeviceType    DeviceType
	RemovableMedium         bool
	LogicalUnitConglomerate bool
	CommandSetVersion       Version
	NormalACASupported      bool
	HistoricalSupport       bool
	ResponseDataFormat      uint8
	SCCSupported            bool
	TargetPortGroupSupport  uint8
	ThirdPartyCopySupport   bool
	HasProtectionInfo       bool
	HasEnclosureServices    bool
	VendorFeature1          bool
	HasMultipleSCSIPorts    bool
	CmdQueue                bool
	VendorFeature2          bool
	Vendor                  string
	Product                 string
	ProductRevisionLevel    string
	VendorSpecific          []byte
	VersionDescriptors      []uint16
}

// Table 498
type VPDPageCode uint8

const (
	SupportedVPDs                      VPDPageCode = 0x00
	UnitSerialNumberVPD                VPDPageCode = 0x80
	DeviceIdentificationVPD            VPDPageCode = 0x83
	SoftwareInterfaceIdentificationVPD VPDPageCode = 0x84
	ManagementNetworkAddressesVPD      VPDPageCode = 0x85
	ExtendedINQUIRYDataVPD             VPDPageCode = 0x86
	ModePagePolicyVPD                  VPDPageCode = 0x87
	SCSIPortsVPD                       VPDPageCode = 0x88
	ATAInformationVPD                  VPDPageCode = 0x89
	PowerConditionVPD                  VPDPageCode = 0x8a
	DeviceConstituentsVPD              VPDPageCode = 0x8b
)

var vpdPageCodeDesc = map[VPDPageCode]string{
	SupportedVPDs:                      "Supported VPD Pages",
	UnitSerialNumberVPD:                "Unit Serial Number",
	DeviceIdentificationVPD:            "Device Identification",
	SoftwareInterfaceIdentificationVPD: "Software Interface Identification",
	ManagementNetworkAddressesVPD:      "Management Network Addresses",
	ExtendedINQUIRYDataVPD:             "Extended INQUIRY Data",
	ModePagePolicyVPD:                  "Mode Page Policy",
	SCSIPortsVPD:                       "SCSI Ports",
	ATAInformationVPD:                  "ATA Information",
	PowerConditionVPD:                  "Power Condition",
	DeviceConstituentsVPD:              "Device Constituents",
}

func (v VPDPageCode) String() string {
	if str, ok := vpdPageCodeDesc[v]; ok {
		return str
	}
	return fmt.Sprintf("Page %xh", uint8(v))
}

// InquiryVPD requests a specified Vital Product Description Page from the
// device. If the size of the page is known in advance, initialSize should be
// set to a non-zero value to make the query more efficient.
func (d *Device) InquiryVPD(pageCode VPDPageCode, initialSize uint16) ([]byte, error) {
	var bufferSize uint16 = 254
	if initialSize > 0 {
		bufferSize = initialSize
	}
	for {
		data := make([]byte, bufferSize)
		var req [4]byte
		req[0] = 0b1 // Enable Vital Product Data
		req[1] = uint8(pageCode)
		binary.BigEndian.PutUint16(req[2:4], uint16(len(data)))
		if err := d.RawCommand(&CommandDataBuffer{
			OperationCode:         InquiryOp,
			Request:               req[:],
			Data:                  data,
			DataTransferDirection: DataTransferFromDevice,
		}); err != nil {
			return nil, fmt.Errorf("error during INQUIRY VPD: %w", err)
		}
		if data[1] != uint8(pageCode) {
			return nil, fmt.Errorf("requested VPD page %x, got %x", pageCode, data[1])
		}
		pageLength := binary.BigEndian.Uint16(data[2:4])
		if pageLength > math.MaxUint16-4 {
			// Guard against uint16 overflows, this cannot be requested anyway
			return nil, fmt.Errorf("device VPD page is too long (%d bytes)", pageLength)
		}
		if pageLength > uint16(len(data)-4) {
			bufferSize = pageLength + 4
			continue
		}
		return data[4 : pageLength+4], nil
	}
}

// SupportedVPDPages returns the list of supported vital product data pages
// supported by the device.
func (d *Device) SupportedVPDPages() (map[VPDPageCode]bool, error) {
	res, err := d.InquiryVPD(SupportedVPDs, 0)
	if err != nil {
		return nil, err
	}
	supportedPages := make(map[VPDPageCode]bool)
	for _, p := range res {
		supportedPages[VPDPageCode(p)] = true
	}
	return supportedPages, nil
}

// UnitSerialNumber returns the serial number of the device. Only available if
// UnitSerialNumberVPD is a supported VPD page.
func (d *Device) UnitSerialNumber() (string, error) {
	serial, err := d.InquiryVPD(UnitSerialNumberVPD, 0)
	if err != nil {
		return "", err
	}
	return string(bytes.Trim(serial, " \x00")), nil
}
