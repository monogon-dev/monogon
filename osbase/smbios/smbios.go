// Package smbios implements parsing of SMBIOS data structures.
// SMBIOS data is commonly populated by platform firmware to convey various
// metadata (including name, vendor, slots and serial numbers) about the
// platform to the operating system.
// The SMBIOS standard is maintained by DMTF and available at
// https://www.dmtf.org/sites/default/files/standards/documents/
// DSP0134_3.6.0.pdf. The rest of this package just refers to it as "the
// standard".
package smbios

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// See spec section 6.1.2
type structureHeader struct {
	// Types 128 through 256 are reserved for OEM and system-specific use.
	Type uint8
	// Length of the structure including this header, excluding the string
	// set.
	Length uint8
	// Unique handle for this structure.
	Handle uint16
}

type Structure struct {
	Type             uint8
	Handle           uint16
	FormattedSection []byte
	Strings          []string
}

// Table represents a decoded SMBIOS table consisting of its structures.
// A few known structures are parsed if present, the rest is put into
// Structures unparsed.
type Table struct {
	BIOSInformationRaw       *BIOSInformationRaw
	SystemInformationRaw     *SystemInformationRaw
	BaseboardsInformationRaw []*BaseboardInformationRaw
	SystemSlotsRaw           []*SystemSlotRaw
	MemoryDevicesRaw         []*MemoryDeviceRaw

	Structures []Structure
}

const (
	structTypeInactive   = 126
	structTypeEndOfTable = 127
)

func Unmarshal(table *bufio.Reader) (*Table, error) {
	var tbl Table
	for {
		var structHdr structureHeader
		if err := binary.Read(table, binary.LittleEndian, &structHdr); err != nil {
			if err == io.EOF {
				// Be tolerant of EOFs on structure boundaries even though
				// the EOT marker is specified as a type 127 structure.
				break
			}
			return nil, fmt.Errorf("unable to read structure header: %w", err)
		}
		if int(structHdr.Length) < binary.Size(structHdr) {
			return nil, fmt.Errorf("invalid structure: header length (%d) smaller than header", structHdr.Length)
		}
		if structHdr.Type == structTypeEndOfTable {
			break
		}
		var s Structure
		s.Type = structHdr.Type
		s.Handle = structHdr.Handle
		s.FormattedSection = make([]byte, structHdr.Length-uint8(binary.Size(structHdr)))
		if _, err := io.ReadFull(table, s.FormattedSection); err != nil {
			return nil, fmt.Errorf("error while reading structure (handle %d) contents: %w", structHdr.Handle, err)
		}
		// Read string-set
		for {
			str, err := table.ReadString(0x00)
			if err != nil {
				return nil, fmt.Errorf("error while reading string table (handle %d): %w", structHdr.Handle, err)
			}
			// Remove trailing null byte
			str = strings.TrimSuffix(str, "\x00")
			// Don't populate a zero-length first string if the string-set is
			// empty.
			if len(str) != 0 {
				s.Strings = append(s.Strings, str)
			}
			maybeTerminator, err := table.ReadByte()
			if err != nil {
				return nil, fmt.Errorf("error while reading string table (handle %d): %w", structHdr.Handle, err)
			}
			if maybeTerminator == 0 {
				// We have a valid string-set terminator, exit the loop
				break
			}
			// The next byte was not a terminator, put it back
			if err := table.UnreadByte(); err != nil {
				panic(err) // Cannot happen operationally
			}
		}
		switch structHdr.Type {
		case structTypeInactive:
			continue
		case structTypeBIOSInformation:
			var biosInfo BIOSInformationRaw
			if err := UnmarshalStructureRaw(s, &biosInfo); err != nil {
				return nil, fmt.Errorf("failed unmarshaling BIOS Information: %w", err)
			}
			tbl.BIOSInformationRaw = &biosInfo
		case structTypeSystemInformation:
			var systemInfo SystemInformationRaw
			if err := UnmarshalStructureRaw(s, &systemInfo); err != nil {
				return nil, fmt.Errorf("failed unmarshaling System Information: %w", err)
			}
			tbl.SystemInformationRaw = &systemInfo
		case structTypeBaseboardInformation:
			var baseboardInfo BaseboardInformationRaw
			if err := UnmarshalStructureRaw(s, &baseboardInfo); err != nil {
				return nil, fmt.Errorf("failed unmarshaling Baseboard Information: %w", err)
			}
			tbl.BaseboardsInformationRaw = append(tbl.BaseboardsInformationRaw, &baseboardInfo)
		case structTypeSystemSlot:
			var sysSlot SystemSlotRaw
			if err := UnmarshalStructureRaw(s, &sysSlot); err != nil {
				return nil, fmt.Errorf("failed unmarshaling System Slot: %w", err)
			}
			tbl.SystemSlotsRaw = append(tbl.SystemSlotsRaw, &sysSlot)
		case structTypeMemoryDevice:
			var memoryDev MemoryDeviceRaw
			if err := UnmarshalStructureRaw(s, &memoryDev); err != nil {
				return nil, fmt.Errorf("failed unmarshaling Memory Device: %w", err)
			}
			tbl.MemoryDevicesRaw = append(tbl.MemoryDevicesRaw, &memoryDev)
		default:
			// Just pass through the raw structure
			tbl.Structures = append(tbl.Structures, s)
		}
	}
	return &tbl, nil
}

// Version contains a two-part version number consisting of a major and minor
// version. This is a common structure in SMBIOS.
type Version struct {
	Major uint8
	Minor uint8
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

// AtLeast returns true if the version in v is at least the given version.
func (v *Version) AtLeast(major, minor uint8) bool {
	if v.Major > major {
		return true
	}
	return v.Major == major && v.Minor >= minor
}

// UnmarshalStructureRaw unmarshals a SMBIOS structure into a Go struct which
// has some constraints. The first two fields need to be a `uint16 handle` and
// a `StructureVersion Version` field. After that any number of fields may
// follow as long as they are either of type `string` (which will be looked up
// in the string table) or readable by binary.Read. To determine the structure
// version, the smbios_min_vers struct tag needs to be put on the first field
// of a newer structure version. The version implicitly starts with 2.0.
// The version determined is written to the second target struct field.
// Fields which do not have a fixed size need to be typed as a slice and tagged
// with smbios_repeat set to the name of the field containing the count. The
// count field itself needs to be some width of uint.
func UnmarshalStructureRaw(rawStruct Structure, target any) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Pointer {
		return errors.New("target needs to be a pointer")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("target needs to be a pointer to a struct")
	}
	v.Field(0).SetUint(uint64(rawStruct.Handle))
	r := bytes.NewReader(rawStruct.FormattedSection)
	completedVersion := Version{Major: 0, Minor: 0}
	parsingVersion := Version{Major: 2, Minor: 0}
	numFields := v.NumField()
	hasAborted := false
	for i := 2; i < numFields; i++ {
		fieldType := v.Type().Field(i)
		if minVer := fieldType.Tag.Get("smbios_min_ver"); minVer != "" {
			var ver Version
			if _, err := fmt.Sscanf(minVer, "%d.%d", &ver.Major, &ver.Minor); err != nil {
				panic(fmt.Sprintf("invalid smbios_min_ver tag in %v: %v", fieldType.Name, err))
			}
			completedVersion = parsingVersion
			parsingVersion = ver
		}
		f := v.Field(i)

		if repeat := fieldType.Tag.Get("smbios_repeat"); repeat != "" {
			repeatCountField := v.FieldByName(repeat)
			if !repeatCountField.IsValid() {
				panic(fmt.Sprintf("invalid smbios_repeat tag in %v: no such field %q", fieldType.Name, repeat))
			}
			if !repeatCountField.CanUint() {
				panic(fmt.Sprintf("invalid smbios_repeat tag in %v: referenced field %q is not uint-compatible", fieldType.Name, repeat))
			}
			if f.Kind() != reflect.Slice {
				panic(fmt.Sprintf("cannot repeat a field (%q) which is not a slice", fieldType.Name))
			}
			if repeatCountField.Uint() > 65536 {
				return fmt.Errorf("refusing to read a field repeated more than 65536 times (given %d times)", repeatCountField.Uint())
			}
			repeatCount := int(repeatCountField.Uint())
			f.Set(reflect.MakeSlice(f.Type(), repeatCount, repeatCount))
			for j := 0; j < repeatCount; j++ {
				fs := f.Index(j)
				err := unmarshalField(&rawStruct, fs, r)
				if errors.Is(err, io.EOF) {
					hasAborted = true
					break
				} else if err != nil {
					return fmt.Errorf("error unmarshaling field %q: %w", fieldType.Name, err)
				}
			}
		}
		err := unmarshalField(&rawStruct, f, r)
		if errors.Is(err, io.EOF) {
			hasAborted = true
			break
		} else if err != nil {
			return fmt.Errorf("error unmarshaling field %q: %w", fieldType.Name, err)
		}
	}
	if !hasAborted {
		completedVersion = parsingVersion
	}
	if completedVersion.Major == 0 {
		return fmt.Errorf("structure's formatted section (%d bytes) is smaller than its minimal size", len(rawStruct.FormattedSection))
	}
	v.Field(1).Set(reflect.ValueOf(completedVersion))
	return nil
}

func unmarshalField(rawStruct *Structure, field reflect.Value, r *bytes.Reader) error {
	if field.Kind() == reflect.String {
		var stringTableIdx uint8
		err := binary.Read(r, binary.LittleEndian, &stringTableIdx)
		if err != nil {
			return err
		}
		if stringTableIdx == 0 {
			return nil
		}
		if int(stringTableIdx)-1 >= len(rawStruct.Strings) {
			return fmt.Errorf("string index (%d) bigger than string table (%q)", stringTableIdx-1, rawStruct.Strings)
		}
		field.SetString(rawStruct.Strings[stringTableIdx-1])
		return nil
	}
	return binary.Read(r, binary.LittleEndian, field.Addr().Interface())
}
