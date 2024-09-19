// Package fat32 implements a writer for the FAT32 filesystem.
package fat32

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"math"
	"math/bits"
	"strings"
	"time"
	"unicode/utf16"
)

// This package contains multiple references to the FAT32 specification, called
// Microsoft Extensible Firmware Initiative FAT32 File System Specification
// version 1.03 (just called the spec from now on). You can get it at
// https://download.microsoft.com/download/0/8/4/\
// 084c452b-b772-4fe5-89bb-a0cbf082286a/fatgen103.doc

type Options struct {
	// Size of a logical block on the block device. Needs to be a power of two
	// equal or bigger than 512. If left at zero, defaults to 512.
	BlockSize uint16

	// Number of blocks the filesystem should span. If zero, it will be exactly
	// as large as it needs to be.
	BlockCount uint32

	// Human-readable filesystem label. Maximum 10 bytes (gets cut off), should
	// be uppercase alphanumeric.
	Label string

	// Filesystem identifier. If unset (i.e. left at zero) a random value will
	// be assigned by WriteFS.
	ID uint32
}

// SizedReader is an io.Reader with a known size
type SizedReader interface {
	io.Reader
	Size() int64
}

// Attribute is a bitset of flags set on an inode.
// See also the spec page 24
type Attribute uint8

const (
	// AttrReadOnly marks a file as read-only
	AttrReadOnly Attribute = 0x01
	// AttrHidden indicates that directory listings should not show this file.
	AttrHidden Attribute = 0x02
	// AttrSystem indicates that this is an operating system file.
	AttrSystem Attribute = 0x04
	// AttrDirectory indicates that this is a directory and not a file.
	AttrDirectory Attribute = 0x10
	// AttrArchive canonically indicates that a file has been created/modified
	// since the last backup. Its use in practice is inconsistent.
	AttrArchive Attribute = 0x20
)

// Inode is file or directory on the FAT32 filesystem. Note that the concept
// of an inode doesn't really exist on FAT32, its directories are just special
// files.
type Inode struct {
	// Name of the file or directory (not including its path)
	Name string
	// Time the file or directory was last modified
	ModTime time.Time
	// Time the file or directory was created
	CreateTime time.Time
	// Attributes
	Attrs Attribute
	// Children of this directory (only valid when Attrs has AttrDirectory set)
	Children []*Inode
	// Content of this file
	// Only valid when Attrs doesn't have AttrDirectory set.
	Content SizedReader

	// Filled out on placement and write-out
	startCluster int
	parent       *Inode
	dosName      [11]byte
}

// Number of LFN entries + normal entry (all 32 bytes)
func (i Inode) metaSize() (int64, error) {
	fileNameUTF16 := utf16.Encode([]rune(i.Name))
	// VFAT file names are null-terminated
	fileNameUTF16 = append(fileNameUTF16, 0x00)
	if len(fileNameUTF16) > 255 {
		return 0, errors.New("file name too long, maximum is 255 UTF-16 code points")
	}

	// ⌈len(fileNameUTF16)/codepointsPerEntry⌉
	numEntries := (len(fileNameUTF16) + codepointsPerEntry - 1) / codepointsPerEntry
	return (int64(numEntries) + 1) * 32, nil
}

func lfnChecksum(dosName [11]byte) uint8 {
	var sum uint8
	for _, b := range dosName {
		sum = ((sum & 1) << 7) + (sum >> 1) + b
	}
	return sum
}

// writeMeta writes information about this inode into the contents of the parent
// inode.
func (i Inode) writeMeta(w io.Writer) error {
	fileNameUTF16 := utf16.Encode([]rune(i.Name))
	// VFAT file names are null-terminated
	fileNameUTF16 = append(fileNameUTF16, 0x00)
	if len(fileNameUTF16) > 255 {
		return errors.New("file name too long, maximum is 255 UTF-16 code points")
	}

	// ⌈len(fileNameUTF16)/codepointsPerEntry⌉
	numEntries := (len(fileNameUTF16) + codepointsPerEntry - 1) / codepointsPerEntry
	// Fill up to space in given number of entries with fill code point 0xffff
	fillCodePoints := (numEntries * codepointsPerEntry) - len(fileNameUTF16)
	for j := 0; j < fillCodePoints; j++ {
		fileNameUTF16 = append(fileNameUTF16, 0xffff)
	}

	// Write entries in reverse order
	for j := numEntries; j > 0; j-- {
		// Index of the code point being processed
		cpIdx := (j - 1) * codepointsPerEntry
		var entry lfnEntry
		entry.Checksum = lfnChecksum(i.dosName)
		// Downcast is safe as i <= numEntries <= ⌈255/codepointsPerEntry⌉
		entry.SequenceNumber = uint8(j)
		if j == numEntries {
			entry.SequenceNumber |= lastSequenceNumberFlag
		}
		entry.Attributes = 0x0F
		copy(entry.NamePart1[:], fileNameUTF16[cpIdx:])
		cpIdx += len(entry.NamePart1)
		copy(entry.NamePart2[:], fileNameUTF16[cpIdx:])
		cpIdx += len(entry.NamePart2)
		copy(entry.NamePart3[:], fileNameUTF16[cpIdx:])
		cpIdx += len(entry.NamePart3)

		if err := binary.Write(w, binary.LittleEndian, entry); err != nil {
			return err
		}
	}
	selfSize, err := i.dataSize()
	if err != nil {
		return err
	}
	if selfSize >= 4*1024*1024*1024 {
		return errors.New("single file size exceeds 4GiB which is prohibited in FAT32")
	}
	if i.Attrs&AttrDirectory != 0 {
		selfSize = 0 // Directories don't have an explicit size
	}
	date, t, _ := timeToMsDosTime(i.ModTime)
	if err := binary.Write(w, binary.LittleEndian, &dirEntry{
		DOSName:           i.dosName,
		Attributes:        uint8(i.Attrs),
		FirstClusterHigh:  uint16(i.startCluster >> 16),
		LastWrittenToTime: t,
		LastWrittenToDate: date,
		FirstClusterLow:   uint16(i.startCluster & 0xffff),
		FileSize:          uint32(selfSize),
	}); err != nil {
		return err
	}
	return nil
}

// writeData writes the contents of this inode (including possible metadata
// of its children, but not its children's data)
func (i Inode) writeData(w io.Writer, volumeLabel [11]byte) error {
	if i.Attrs&AttrDirectory != 0 {
		if i.parent == nil {
			if err := binary.Write(w, binary.LittleEndian, &dirEntry{
				DOSName:    volumeLabel,
				Attributes: 0x08, // Volume ID, internal use only
			}); err != nil {
				return err
			}
		} else {
			date, t, _ := timeToMsDosTime(i.ModTime)
			cdate, ctime, ctens := timeToMsDosTime(i.CreateTime)
			if err := binary.Write(w, binary.LittleEndian, &dirEntry{
				DOSName:           [11]byte{'.', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				CreationDate:      cdate,
				CreationTime:      ctime,
				CreationTenMilli:  ctens,
				LastWrittenToTime: t,
				LastWrittenToDate: date,
				Attributes:        uint8(i.Attrs),
				FirstClusterHigh:  uint16(i.startCluster >> 16),
				FirstClusterLow:   uint16(i.startCluster & 0xffff),
			}); err != nil {
				return err
			}
			startCluster := i.parent.startCluster
			if i.parent.parent == nil {
				// Special case: When the dotdot directory points to the root
				// directory, the start cluster is defined to be zero even if
				// it isn't.
				startCluster = 0
			}
			// Time is intentionally taken from this directory, not the parent
			if err := binary.Write(w, binary.LittleEndian, &dirEntry{
				DOSName:           [11]byte{'.', '.', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				LastWrittenToTime: t,
				LastWrittenToDate: date,
				Attributes:        uint8(AttrDirectory),
				FirstClusterHigh:  uint16(startCluster >> 16),
				FirstClusterLow:   uint16(startCluster & 0xffff),
			}); err != nil {
				return err
			}
		}
		err := makeUniqueDOSNames(i.Children)
		if err != nil {
			return err
		}
		for _, c := range i.Children {
			if err := c.writeMeta(w); err != nil {
				return err
			}
		}
	} else {
		if _, err := io.CopyN(w, i.Content, i.Content.Size()); err != nil {
			return err
		}
	}
	return nil
}

func (i Inode) dataSize() (int64, error) {
	if i.Attrs&AttrDirectory != 0 {
		var size int64
		if i.parent != nil {
			// Dot and dotdot directories
			size += 2 * 32
		} else {
			// Volume ID
			size += 1 * 32
		}
		for _, c := range i.Children {
			cs, err := c.metaSize()
			if err != nil {
				return 0, err
			}
			size += cs
		}
		if size > 2*1024*1024 {
			return 0, errors.New("directory contains > 2MiB of metadata which is prohibited in FAT32")
		}
		return size, nil
	} else {
		return i.Content.Size(), nil
	}
}

func (i *Inode) PlaceFile(path string, reader SizedReader) error {
	pathParts := strings.Split(path, "/")
	inodeRef := i
	for j, part := range pathParts {
		var childExists bool
		for _, child := range inodeRef.Children {
			if strings.EqualFold(child.Name, part) {
				inodeRef = child
				childExists = true
				break
			}
		}
		if j == len(pathParts)-1 { // Is last path part (i.e. file name)
			if childExists {
				return &fs.PathError{Path: path, Err: fs.ErrExist, Op: "create"}
			}
			newInode := &Inode{
				Name:    part,
				Content: reader,
			}
			inodeRef.Children = append(inodeRef.Children, newInode)
			return nil
		} else if !childExists {
			newInode := &Inode{
				Name:  part,
				Attrs: AttrDirectory,
			}
			inodeRef.Children = append(inodeRef.Children, newInode)
			inodeRef = newInode
		}
	}
	panic("unreachable")
}

type planningState struct {
	// List of inodes in filesystem layout order
	orderedInodes []*Inode
	// File Allocation Table
	fat []uint32
	// Size of a single cluster in the FAT in bytes
	clusterSize int64
}

// Allocates clusters capable of holding at least b bytes and returns the
// starting cluster index
func (p *planningState) allocBytes(b int64) int {
	// Zero-byte data entries are located at the cluster zero by definition
	// No actual allocation is performed
	if b == 0 {
		return 0
	}
	// Calculate the number of clusters to be allocated
	n := (b + p.clusterSize - 1) / p.clusterSize
	allocStartCluster := len(p.fat)
	for i := int64(0); i < n-1; i++ {
		p.fat = append(p.fat, uint32(len(p.fat)+1))
	}
	p.fat = append(p.fat, fatEOF)
	return allocStartCluster
}

func (i *Inode) placeRecursively(p *planningState) error {
	selfDataSize, err := i.dataSize()
	if err != nil {
		return fmt.Errorf("%s: %w", i.Name, err)
	}
	i.startCluster = p.allocBytes(selfDataSize)
	p.orderedInodes = append(p.orderedInodes, i)
	for _, c := range i.Children {
		c.parent = i
		err = c.placeRecursively(p)
		if err != nil {
			return fmt.Errorf("%s/%w", i.Name, err)
		}
	}
	return nil
}

// WriteFS writes a filesystem described by a root inode and its children to a
// given io.Writer.
func WriteFS(w io.Writer, rootInode Inode, opts Options) error {
	bs, fsi, p, err := prepareFS(&opts, rootInode)
	if err != nil {
		return err
	}

	wb := newBlockWriter(w)

	// Write superblock
	if err := binary.Write(wb, binary.LittleEndian, bs); err != nil {
		return err
	}
	if err := wb.FinishBlock(int64(opts.BlockSize), true); err != nil {
		return err
	}
	if err := binary.Write(wb, binary.LittleEndian, fsi); err != nil {
		return err
	}
	if err := wb.FinishBlock(int64(opts.BlockSize), true); err != nil {
		return err
	}

	block := make([]byte, opts.BlockSize)
	for i := 0; i < 4; i++ {
		if _, err := wb.Write(block); err != nil {
			return err
		}
	}
	// Backup of superblock at block 6
	if err := binary.Write(wb, binary.LittleEndian, bs); err != nil {
		return err
	}
	if err := wb.FinishBlock(int64(opts.BlockSize), true); err != nil {
		return err
	}
	if err := binary.Write(wb, binary.LittleEndian, fsi); err != nil {
		return err
	}
	if err := wb.FinishBlock(int64(opts.BlockSize), true); err != nil {
		return err
	}

	for i := uint8(0); i < bs.NumFATs; i++ {
		if err := binary.Write(wb, binary.LittleEndian, p.fat); err != nil {
			return err
		}
		if err := wb.FinishBlock(int64(opts.BlockSize), true); err != nil {
			return err
		}
	}

	for _, i := range p.orderedInodes {
		if err := i.writeData(wb, bs.Label); err != nil {
			return fmt.Errorf("failed to write inode %q: %w", i.Name, err)
		}
		if err := wb.FinishBlock(int64(opts.BlockSize)*int64(bs.BlocksPerCluster), false); err != nil {
			return err
		}
	}
	// Creatively use block writer to write out all empty data at the end
	if err := wb.FinishBlock(int64(opts.BlockSize)*int64(bs.TotalBlocks), false); err != nil {
		return err
	}
	return nil
}

func prepareFS(opts *Options, rootInode Inode) (*bootSector, *fsinfo, *planningState, error) {
	if opts.BlockSize == 0 {
		opts.BlockSize = 512
	}
	if bits.OnesCount16(opts.BlockSize) != 1 {
		return nil, nil, nil, fmt.Errorf("option BlockSize is not a power of two")
	}
	if opts.BlockSize < 512 {
		return nil, nil, nil, fmt.Errorf("option BlockSize must be at least 512 bytes")
	}
	if opts.ID == 0 {
		var buf [4]byte
		if _, err := rand.Read(buf[:]); err != nil {
			return nil, nil, nil, fmt.Errorf("failed to assign random FAT ID: %w", err)
		}
		opts.ID = binary.BigEndian.Uint32(buf[:])
	}
	if rootInode.Attrs&AttrDirectory == 0 {
		return nil, nil, nil, errors.New("root inode must be a directory (i.e. have AttrDirectory set)")
	}
	bs := bootSector{
		// Assembled x86_32 machine code corresponding to
		// jmp $
		// nop
		// i.e. an infinite loop doing nothing. Nothing created in the last 35
		// years should boot this anyway.
		// TODO(q3k): write a stub
		JmpInstruction: [3]byte{0xEB, 0xFE, 0x90},
		// Identification
		OEMName: [8]byte{'M', 'O', 'N', 'O', 'G', 'O', 'N'},
		ID:      opts.ID,
		// Block geometry
		BlockSize:   opts.BlockSize,
		TotalBlocks: opts.BlockCount,
		// BootSector block + FSInfo Block, backup copy at blocks 6 and 7
		ReservedBlocks: 8,
		// FSInfo block is always in block 1, right after this block
		FSInfoBlock: 1,
		// Start block of the backup of the boot block and FSInfo block
		// De facto this must be 6 as it is only used when the primary
		// boot block is damaged at which point this field can no longer be
		// read.
		BackupStartBlock: 6,
		// A lot of implementations only work with 2, so use that
		NumFATs:          2,
		BlocksPerCluster: 1,
		// Flags and signatures
		MediaCode:     0xf8,
		BootSignature: 0x29,
		Label:         [11]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		Type:          [8]byte{'F', 'A', 'T', '3', '2', ' ', ' ', ' '},
		Signature:     [2]byte{0x55, 0xaa},
	}

	copy(bs.Label[:], opts.Label)

	fsi := fsinfo{
		// Signatures
		LeadSignature:     [4]byte{0x52, 0x52, 0x61, 0x41},
		StructSignature:   [4]byte{0x72, 0x72, 0x41, 0x61},
		TrailingSignature: [2]byte{0x55, 0xAA},

		// This is the unset value which is always legal
		NextFreeCluster: 0xFFFFFFFF,
	}

	p := planningState{
		clusterSize: int64(bs.BlocksPerCluster) * int64(bs.BlockSize),
	}
	if opts.BlockCount != 0 {
		// Preallocate FAT if we know how big it needs to be
		p.fat = make([]uint32, 0, opts.BlockCount/uint32(bs.BlocksPerCluster))
	} else {
		// Preallocate minimum size FAT
		// See the spec page 15 for the origin of this calculation.
		p.fat = make([]uint32, 0, 65525+2)
	}
	// First two clusters are special
	p.fat = append(p.fat, 0x0fffff00|uint32(bs.MediaCode), 0x0fffffff)
	err := rootInode.placeRecursively(&p)
	if err != nil {
		return nil, nil, nil, err
	}

	allocClusters := len(p.fat)
	if allocClusters >= fatMask&math.MaxUint32 {
		return nil, nil, nil, fmt.Errorf("filesystem contains more than 2^28 FAT entries, this is unsupported. Note that this package currently always creates minimal clusters")
	}

	// Fill out FAT to minimum size for FAT32
	for len(p.fat) < 65525+2 {
		p.fat = append(p.fat, fatFree)
	}

	bs.RootClusterNumber = uint32(rootInode.startCluster)

	bs.BlocksPerFAT = uint32(binary.Size(p.fat)+int(opts.BlockSize)-1) / uint32(opts.BlockSize)
	occupiedBlocks := uint32(bs.ReservedBlocks) + (uint32(len(p.fat)-2) * uint32(bs.BlocksPerCluster)) + bs.BlocksPerFAT*uint32(bs.NumFATs)
	if bs.TotalBlocks == 0 {
		bs.TotalBlocks = occupiedBlocks
	} else if bs.TotalBlocks < occupiedBlocks {
		return nil, nil, nil, fmt.Errorf("content (minimum %d blocks) would exceed number of blocks specified (%d blocks)", occupiedBlocks, bs.TotalBlocks)
	} else { // Fixed-size file system with enough space
		blocksToDistribute := bs.TotalBlocks - uint32(bs.ReservedBlocks)
		// Number of data blocks which can be described by one metadata/FAT
		// block. Always an integer because 4 (bytes per uint32) is a divisor of
		// all powers of two equal or bigger than 8 and FAT32 requires a minimum
		// of 512.
		dataBlocksPerFATBlock := (uint32(bs.BlocksPerCluster) * uint32(bs.BlockSize)) / (uint32(binary.Size(p.fat[0])))
		// Split blocksToDistribute between metadata and data so that exactly as
		// much metadata (FAT) exists for describing the amount of data blocks
		// while respecting alignment.
		divisor := dataBlocksPerFATBlock + uint32(bs.NumFATs)
		// 2*blocksPerCluster compensates for the first two "magic" FAT entries
		// which do not have corresponding data.
		bs.BlocksPerFAT = (bs.TotalBlocks + 2*uint32(bs.BlocksPerCluster) + (divisor - 1)) / divisor
		dataBlocks := blocksToDistribute - (uint32(bs.NumFATs) * bs.BlocksPerFAT)
		// Align to full clusters
		dataBlocks -= dataBlocks % uint32(bs.BlocksPerCluster)
		// Magic +2 as the first two entries do not describe data
		for len(p.fat) < (int(dataBlocks)/int(bs.BlocksPerCluster))+2 {
			p.fat = append(p.fat, fatFree)
		}
	}
	fsi.FreeCount = uint32(len(p.fat) - allocClusters)
	if fsi.FreeCount > 1 {
		fsi.NextFreeCluster = uint32(allocClusters) + 1
	}
	return &bs, &fsi, &p, nil
}

// SizeFS returns the number of blocks required to hold the filesystem defined
// by rootInode and opts. This can be used for sizing calculations before
// calling WriteFS.
func SizeFS(rootInode Inode, opts Options) (int64, error) {
	bs, _, _, err := prepareFS(&opts, rootInode)
	if err != nil {
		return 0, err
	}

	return int64(bs.TotalBlocks), nil
}
