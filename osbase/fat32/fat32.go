// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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
	"time"
	"unicode/utf16"

	"source.monogon.dev/osbase/structfs"
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

	// Human-readable filesystem label. Maximum 11 bytes (gets cut off), should
	// be uppercase alphanumeric.
	Label string

	// Filesystem identifier. If unset (i.e. left at zero) a random value will
	// be assigned by WriteFS.
	ID uint32
}

// Attribute is a bitset of flags set on a directory entry.
// See also the spec page 24
type Attribute uint8

const (
	// AttrReadOnly marks a file as read-only
	AttrReadOnly Attribute = 0x01
	// AttrHidden indicates that directory listings should not show this file.
	AttrHidden Attribute = 0x02
	// AttrSystem indicates that this is an operating system file.
	AttrSystem Attribute = 0x04
	// attrVolumeID indicates that this is a special directory entry which
	// contains the volume label.
	attrVolumeID Attribute = 0x08
	// attrDirectory indicates that this is a directory and not a file.
	attrDirectory Attribute = 0x10
	// AttrArchive canonically indicates that a file has been created/modified
	// since the last backup. Its use in practice is inconsistent.
	AttrArchive Attribute = 0x20
)

// DirEntrySys contains additional directory entry fields which are specific to
// FAT32. To set these fields, the Sys field of a [structfs.Node] can be set to
// a pointer to this or to a struct which embeds it.
type DirEntrySys struct {
	// Time the file or directory was created
	CreateTime time.Time
	// Attributes
	Attrs Attribute
}

func (d *DirEntrySys) FAT32() *DirEntrySys {
	return d
}

// DirEntrySysAccessor is used to access [DirEntrySys] instead of directly type
// asserting the struct, to allow for embedding.
type DirEntrySysAccessor interface {
	FAT32() *DirEntrySys
}

// node is a file or directory on the FAT32 filesystem. It wraps a
// [structfs.Node] and holds additional fields which are filled during planning.
type node struct {
	*structfs.Node
	dosName      [11]byte
	createTime   time.Time
	attrs        Attribute
	parent       *node
	children     []*node
	size         uint32
	startCluster int
}

// Number of LFN entries + normal entry (all 32 bytes)
func (i node) metaSize() (int64, error) {
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

// writeMeta writes information about this node into the contents of the parent
// node.
func (i node) writeMeta(w io.Writer) error {
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
	selfSize := i.size
	if i.attrs&attrDirectory != 0 {
		selfSize = 0 // Directories don't have an explicit size
	}
	date, t, _ := timeToMsDosTime(i.ModTime)
	cdate, ctime, ctens := timeToMsDosTime(i.createTime)
	if err := binary.Write(w, binary.LittleEndian, &dirEntry{
		DOSName:           i.dosName,
		Attributes:        uint8(i.attrs),
		CreationTenMilli:  ctens,
		CreationTime:      ctime,
		CreationDate:      cdate,
		FirstClusterHigh:  uint16(i.startCluster >> 16),
		LastWrittenToTime: t,
		LastWrittenToDate: date,
		FirstClusterLow:   uint16(i.startCluster & 0xffff),
		FileSize:          selfSize,
	}); err != nil {
		return err
	}
	return nil
}

// writeData writes the contents of this node (including possible metadata
// of its children, but not its children's data)
func (i node) writeData(w io.Writer, volumeLabel [11]byte) error {
	if i.attrs&attrDirectory != 0 {
		if i.parent == nil {
			if err := binary.Write(w, binary.LittleEndian, &dirEntry{
				DOSName:    volumeLabel,
				Attributes: uint8(attrVolumeID),
			}); err != nil {
				return err
			}
		} else {
			date, t, _ := timeToMsDosTime(i.ModTime)
			cdate, ctime, ctens := timeToMsDosTime(i.createTime)
			if err := binary.Write(w, binary.LittleEndian, &dirEntry{
				DOSName:           [11]byte{'.', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				CreationDate:      cdate,
				CreationTime:      ctime,
				CreationTenMilli:  ctens,
				LastWrittenToTime: t,
				LastWrittenToDate: date,
				Attributes:        uint8(i.attrs),
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
				CreationDate:      cdate,
				CreationTime:      ctime,
				CreationTenMilli:  ctens,
				LastWrittenToTime: t,
				LastWrittenToDate: date,
				Attributes:        uint8(attrDirectory),
				FirstClusterHigh:  uint16(startCluster >> 16),
				FirstClusterLow:   uint16(startCluster & 0xffff),
			}); err != nil {
				return err
			}
		}
		for _, c := range i.children {
			if err := c.writeMeta(w); err != nil {
				return err
			}
		}
	} else {
		content, err := i.Content.Open()
		if err != nil {
			return err
		}
		defer content.Close()
		if _, err := io.CopyN(w, content, int64(i.size)); err != nil {
			return err
		}
	}
	return nil
}

func (i node) dirSize() (uint32, error) {
	var size int64
	if i.parent != nil {
		// Dot and dotdot directories
		size += 2 * 32
	} else {
		// Volume ID
		size += 1 * 32
	}
	for _, c := range i.children {
		cs, err := c.metaSize()
		if err != nil {
			return 0, err
		}
		size += cs
	}
	if size > 2*1024*1024 {
		return 0, errors.New("directory contains > 2MiB of metadata which is prohibited in FAT32")
	}
	return uint32(size), nil
}

type planningState struct {
	// List of nodes in filesystem layout order
	orderedNodes []*node
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

func (i *node) placeRecursively(p *planningState) error {
	if i.Mode.IsDir() {
		for _, c := range i.Node.Children {
			node := &node{
				Node:       c,
				createTime: c.ModTime,
				parent:     i,
			}
			if sys, ok := c.Sys.(DirEntrySysAccessor); ok {
				sys := sys.FAT32()
				node.attrs = sys.Attrs & (AttrReadOnly | AttrHidden | AttrSystem | AttrArchive)
				if !sys.CreateTime.IsZero() {
					node.createTime = sys.CreateTime
				}
			}
			switch {
			case c.Mode.IsRegular():
				size := c.Content.Size()
				if size < 0 {
					return fmt.Errorf("%s: negative file size", c.Name)
				}
				if size >= 4*1024*1024*1024 {
					return fmt.Errorf("%s: single file size exceeds 4GiB which is prohibited in FAT32", c.Name)
				}
				node.size = uint32(size)
				if len(c.Children) != 0 {
					return fmt.Errorf("%s: file cannot have children", c.Name)
				}
			case c.Mode.IsDir():
				node.attrs |= attrDirectory
			default:
				return fmt.Errorf("%s: unsupported file type %s", c.Name, c.Mode.Type().String())
			}
			i.children = append(i.children, node)
		}
		err := makeUniqueDOSNames(i.children)
		if err != nil {
			return err
		}
		i.size, err = i.dirSize()
		if err != nil {
			return fmt.Errorf("%s: %w", i.Name, err)
		}
	}
	i.startCluster = p.allocBytes(int64(i.size))
	p.orderedNodes = append(p.orderedNodes, i)
	for _, c := range i.children {
		err := c.placeRecursively(p)
		if err != nil {
			return fmt.Errorf("%s/%w", i.Name, err)
		}
	}
	return nil
}

// WriteFS writes a filesystem described by a tree to a given io.Writer.
func WriteFS(w io.Writer, root structfs.Tree, opts Options) error {
	bs, fsi, p, err := prepareFS(&opts, root)
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

	for _, i := range p.orderedNodes {
		if err := i.writeData(wb, bs.Label); err != nil {
			return fmt.Errorf("failed to write contents of %q: %w", i.Name, err)
		}
		if err := wb.FinishBlock(int64(opts.BlockSize)*int64(bs.BlocksPerCluster), true); err != nil {
			return err
		}
	}
	// Creatively use block writer to write out all empty data at the end
	if err := wb.FinishBlock(int64(opts.BlockSize)*int64(bs.TotalBlocks), false); err != nil {
		return err
	}
	return nil
}

func prepareFS(opts *Options, root structfs.Tree) (*bootSector, *fsinfo, *planningState, error) {
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
	bs := bootSector{
		// Assembled x86_32 machine code corresponding to
		// jmp $
		// nop
		// i.e. an infinite loop doing nothing. Nothing created in the last 35
		// years should boot this anyway.
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
	rootNode := &node{
		Node: &structfs.Node{
			Mode:     fs.ModeDir,
			Children: root,
		},
		attrs: attrDirectory,
	}
	err := rootNode.placeRecursively(&p)
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

	bs.RootClusterNumber = uint32(rootNode.startCluster)

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
// by root and opts. This can be used for sizing calculations before calling
// WriteFS.
func SizeFS(root structfs.Tree, opts Options) (int64, error) {
	bs, _, _, err := prepareFS(&opts, root)
	if err != nil {
		return 0, err
	}

	return int64(bs.TotalBlocks), nil
}
