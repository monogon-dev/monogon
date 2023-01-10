package scsi

// Written against SBC-4
// Contains SCSI block device specific commands.

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

// ReadDefectDataLBA reads the primary (manufacturer) and/or grown defect list
// in LBA format. This is commonly used on SSDs and generally returns an error
// on spinning drives.
func (d *Device) ReadDefectDataLBA(plist, glist bool) ([]uint64, error) {
	data := make([]byte, 4096)
	var req [8]byte
	if plist {
		req[1] |= 1 << 4
	}
	if glist {
		req[1] |= 1 << 3
	}
	defectListFormat := 0b011
	req[1] |= byte(defectListFormat)
	binary.BigEndian.PutUint16(req[6:8], uint16(len(data)))
	if err := d.RawCommand(&CommandDataBuffer{
		OperationCode:         ReadDefectDataOp,
		Request:               req[:],
		Data:                  data,
		DataTransferDirection: DataTransferFromDevice,
	}); err != nil {
		var fixedErr *FixedError
		if errors.As(err, &fixedErr) && fixedErr.SenseKey == RecoveredError && fixedErr.AdditionalSenseCode == DefectListNotFound {
			return nil, fmt.Errorf("error during LOG SENSE: unsupported defect list format, device returned %03bb", data[1]&0b111)
		}
		return nil, fmt.Errorf("error during LOG SENSE: %w", err)
	}
	if data[1]&0b111 != byte(defectListFormat) {
		return nil, fmt.Errorf("device returned wrong defect list format, requested %03bb, got %03bb", defectListFormat, data[1]&0b111)
	}
	defectListLength := binary.BigEndian.Uint16(data[2:4])
	if defectListLength%8 != 0 {
		return nil, errors.New("returned defect list not divisible by array item size")
	}
	res := make([]uint64, defectListLength/8)
	if err := binary.Read(bytes.NewReader(data[4:]), binary.BigEndian, &res); err != nil {
		panic(err)
	}
	return res, nil
}

const (
	// AllSectors is a magic sector number indicating that it applies to all
	// sectors on the track.
	AllSectors = math.MaxUint16
)

// PhysicalSectorFormatAddress represents a physical sector (or the the whole
// track if SectorNumber == AllSectors) on a spinning hard drive.
type PhysicalSectorFormatAddress struct {
	CylinderNumber              uint32
	HeadNumber                  uint8
	SectorNumber                uint32
	MultiAddressDescriptorStart bool
}

func parseExtendedPhysicalSectorFormatAddress(buf []byte) (p PhysicalSectorFormatAddress) {
	p.CylinderNumber = uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])
	p.HeadNumber = buf[3]
	p.MultiAddressDescriptorStart = buf[4]&(1<<7) != 0
	p.SectorNumber = uint32(buf[4]&0b1111)<<24 | uint32(buf[5])<<16 | uint32(buf[6])<<8 | uint32(buf[7])
	return
}

func parsePhysicalSectorFormatAddress(buf []byte) (p PhysicalSectorFormatAddress) {
	p.CylinderNumber = uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])
	p.HeadNumber = buf[3]
	p.SectorNumber = binary.BigEndian.Uint32(buf[4:8])
	return
}

// ReadDefectDataPhysical reads the primary (manufacturer) and/or grown defect
// list in physical format.
// This is only defined for spinning drives, returning an error on SSDs.
func (d *Device) ReadDefectDataPhysical(plist, glist bool) ([]PhysicalSectorFormatAddress, error) {
	data := make([]byte, 4096)
	var req [8]byte
	if plist {
		req[1] |= 1 << 4
	}
	if glist {
		req[1] |= 1 << 3
	}
	defectListFormat := 0b101
	req[1] |= byte(defectListFormat)
	binary.BigEndian.PutUint16(req[6:8], uint16(len(data)))
	if err := d.RawCommand(&CommandDataBuffer{
		OperationCode:         ReadDefectDataOp,
		Request:               req[:],
		Data:                  data,
		DataTransferDirection: DataTransferFromDevice,
	}); err != nil {
		var fixedErr *FixedError
		if errors.As(err, &fixedErr) && fixedErr.SenseKey == RecoveredError && fixedErr.AdditionalSenseCode == DefectListNotFound {
			return nil, fmt.Errorf("error during LOG SENSE: unsupported defect list format, device returned %03bb", data[1]&0b111)
		}
		return nil, fmt.Errorf("error during LOG SENSE: %w", err)
	}
	if data[1]&0b111 != byte(defectListFormat) {
		return nil, fmt.Errorf("device returned wrong defect list format, requested %03bb, got %03bb", defectListFormat, data[1]&0b111)
	}
	defectListLength := binary.BigEndian.Uint16(data[2:4])
	if defectListLength%8 != 0 {
		return nil, errors.New("returned defect list not divisible by array item size")
	}
	if len(data) < int(defectListLength)+4 {
		return nil, errors.New("returned defect list longer than buffer")
	}
	res := make([]PhysicalSectorFormatAddress, defectListLength/8)
	data = data[4:]
	for i := 0; i < int(defectListLength)/8; i++ {
		res[i] = parsePhysicalSectorFormatAddress(data[i*8 : (i+1)*8])
	}
	return res, nil
}

type SolidStateMediaHealth struct {
	// PercentageUsedEnduranceIndicator is a value which represents a
	// vendor-specific wear estimate of the solid state medium.
	// A new device starts at 0, at 100 the device is considered end-of-life.
	// Values up to 255 are possible.
	PercentageUsedEnduranceIndicator uint8
}

// SolidStateMediaHealth reports parameters about the health of the solid-state
// media of a SCSI block device.
func (d *Device) SolidStateMediaHealth() (*SolidStateMediaHealth, error) {
	raw, err := d.LogSenseParameters(LogSenseRequest{PageCode: 0x11})
	if err != nil {
		return nil, err
	}
	if len(raw[0x1]) == 0 {
		return nil, errors.New("mandatory parameter 0001h missing")
	}
	param1 := raw[0x01][0]
	if len(param1.Data) < 4 {
		return nil, errors.New("parameter 0001h too short")
	}
	return &SolidStateMediaHealth{
		PercentageUsedEnduranceIndicator: param1.Data[3],
	}, nil
}
