package gpt

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// See UEFI Specification 2.9 Table 5-3
type mbr struct {
	BootCode         [440]byte
	DiskSignature    [4]byte
	_                [2]byte
	PartitionRecords [4]mbrPartitionRecord
	Signature        [2]byte
}

// See UEFI Specification 2.9 Table 5-4
type mbrPartitionRecord struct {
	BootIndicator byte
	StartingCHS   [3]byte
	Type          byte
	EndingCHS     [3]byte
	StartingBlock uint32
	SizeInBlocks  uint32
}

var mbrSignature = [2]byte{0x55, 0xaa}

func makeProtectiveMBR(w io.Writer, blockCount int64, bootCode []byte) error {
	var representedBlockCount = uint32(math.MaxUint32)
	if blockCount < math.MaxUint32 {
		representedBlockCount = uint32(blockCount)
	}
	m := mbr{
		DiskSignature: [4]byte{0, 0, 0, 0},
		PartitionRecords: [4]mbrPartitionRecord{
			{
				StartingCHS:   toCHS(1),
				Type:          0xEE, // Table/Protective MBR
				StartingBlock: 1,
				SizeInBlocks:  representedBlockCount,
				EndingCHS:     toCHS(blockCount + 1),
			},
			{},
			{},
			{},
		},
		Signature: mbrSignature,
	}
	if len(bootCode) > len(m.BootCode) {
		return fmt.Errorf("BootCode is %d bytes, can only store %d", len(bootCode), len(m.BootCode))
	}
	copy(m.BootCode[:], bootCode)
	if err := binary.Write(w, binary.LittleEndian, &m); err != nil {
		return fmt.Errorf("failed to write MBR: %w", err)
	}
	return nil
}

// toCHS converts a LBA to a "logical" CHS, i.e. what a legacy BIOS 13h
// interface would use. This has nothing to do with the actual CHS geometry
// which depends on the disk and interface used.
func toCHS(lba int64) (chs [3]byte) {
	const maxCylinders = (1 << 10) - 1
	const maxHeadsPerCylinder = (1 << 8) - 1
	const maxSectorsPerTrack = (1 << 6) - 2 // Sector is 1-based
	cylinder := lba / (maxHeadsPerCylinder * maxSectorsPerTrack)
	head := (lba / maxSectorsPerTrack) % maxHeadsPerCylinder
	sector := (lba % maxSectorsPerTrack) + 1
	if cylinder > maxCylinders {
		cylinder = maxCylinders
		head = maxHeadsPerCylinder
		sector = maxSectorsPerTrack + 1
	}
	chs[0] = uint8(head)
	chs[1] = uint8(sector)
	chs[1] |= uint8(cylinder>>2) & 0xc0
	chs[2] = uint8(cylinder)
	return
}
