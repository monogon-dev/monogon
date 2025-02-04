// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package fat32

const (
	// FAT32 entries are only 28 bits
	fatMask = 0x0fffffff
	// Free entries are 0
	fatFree = 0x0
	// Entry at the end of a cluster chain
	fatEOF = 0x0ffffff8
)

// FAT32 Boot Sector and BIOS Parameter Block. This structure is 512 bytes long,
// even if the logical block size is longer. The rest should be filled up with
// zeroes.
type bootSector struct {
	// Jump instruction to boot code.
	JmpInstruction [3]byte
	// Creator name. "MSWIN4.1" recommended for compatibility.
	OEMName [8]byte
	// Count of bytes per block (i.e. logical block size)
	// Must be one of 512, 1024, 2048 or 4096
	BlockSize uint16
	// Number of blocks per allocation unit (cluster).
	// Must be a power of 2 that is greater than 0.
	BlocksPerCluster uint8
	// Number of reserved blocks in the reserved region of the volume starting
	// at the first block of the volume. This field must not be 0.
	ReservedBlocks uint16
	// The count of FAT data structures on the volume. This field should always
	// contain the value of 2 for any FAT volume of any type.
	NumFATs uint8
	_       [4]byte
	// Legacy value for media determination, must be 0xf8.
	MediaCode uint8
	_         [2]byte
	// Number of sectors per track for 0x13 interrupts.
	SectorsPerTrack uint16
	// Number of heads for 0x13 interrupts.
	NumHeads uint16
	// Count of hidden blocks preceding the partition that contains this FAT
	// volume.
	HiddenBlocks uint32
	// Total count of blocks on the volume.
	TotalBlocks uint32
	// Count of blocks per FAT.
	BlocksPerFAT uint32
	// Flags for FAT32
	Flags uint16
	_     [2]byte
	// Cluster number of the first cluster of the root directory. Usually 2.
	RootClusterNumber uint32
	// Block number of the FSINFO structure in the reserved area.
	FSInfoBlock uint16
	// Block number of the copy of the boot record in the reserved area.
	BackupStartBlock uint16
	_                [12]byte
	// Drive number for 0x13 interrupts.
	DriveNumber   uint8
	_             [1]byte
	BootSignature uint8
	// ID of this filesystem
	ID uint32
	// Human-readable label of this filesystem, padded with spaces (0x20)
	Label [11]byte
	// Always set to ASCII "FAT32    "
	Type [8]byte
	_    [420]byte
	// Always 0x55, 0xAA
	Signature [2]byte
}

// Special block (usually at block 1) containing additional metadata,
// specifically the number of free clusters and the next free cluster.
// Always 512 bytes, rest of the block should be padded with zeroes.
type fsinfo struct {
	// Validates that this is an FSINFO block. Always 0x52, 0x52, 0x61, 0x41
	LeadSignature [4]byte
	_             [480]byte
	// Another signature. Always 0x72, 0x72, 0x41, 0x61
	StructSignature [4]byte
	// Last known number of free clusters on the volume.
	FreeCount uint32
	// Next free cluster hint. All 1's is interpreted as undefined.
	NextFreeCluster uint32
	_               [14]byte
	// One more signature. Always 0x55, 0xAA.
	TrailingSignature [2]byte
}

// Directory entry
type dirEntry struct {
	// DOS 8.3 file name.
	DOSName [11]byte
	// Attribtes of the file or directory, 0x0f reserved to mark entry as a
	// LFN entry (see lfnEntry below)
	Attributes        uint8
	_                 byte
	CreationTenMilli  uint8 // Actually 10ms units, 0-199 range
	CreationTime      uint16
	CreationDate      uint16
	_                 [2]byte
	FirstClusterHigh  uint16
	LastWrittenToTime uint16
	LastWrittenToDate uint16
	FirstClusterLow   uint16
	FileSize          uint32
}

const (
	// lastSequenceNumberFlag is logically-ORed with the sequence number of the
	// last Long File Name entry to mark it as such.
	lastSequenceNumberFlag = 0x40
	// codepointsPerEntry is the number of UTF-16 codepoints that fit into a
	// single Long File Name entry.
	codepointsPerEntry = 5 + 6 + 2
)

// VFAT long file name prepended entry
type lfnEntry struct {
	SequenceNumber uint8
	// First 5 UTF-16 code units
	NamePart1 [5]uint16
	// Attributes (must be 0x0f)
	Attributes uint8
	_          byte
	// Checksum of the 8.3 name.
	Checksum uint8
	// Next 6 UTF-16 code units
	NamePart2 [6]uint16
	_         [2]byte
	// Next 2 UTF-16 code units
	NamePart3 [2]uint16
}
