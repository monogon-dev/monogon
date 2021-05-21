// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package erofs

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"path"

	"golang.org/x/sys/unix"
)

// Writer writes a new EROFS filesystem.
type Writer struct {
	w io.WriteSeeker
	// fixDirectoryEntry contains for each referenced path where it is
	// referenced from. Since self-references are required anyways (for the "."
	// and ".." entries) we let the user write files in any order and just
	// point the directory entries to the right target nid and file type on
	// Close().
	fixDirectoryEntry map[string][]direntFixupLocation
	pathInodeMeta     map[string]*uncompressedInodeMeta
	// legacyInodeIndex stores the next legacy (32-bit) inode to be allocated.
	// 64 bit inodes are automatically calculated by EROFS on mount.
	legacyInodeIndex    uint32
	blockAllocatorIndex uint32
	metadataBlocksFree  metadataBlocksMeta
}

// NewWriter creates a new EROFS filesystem writer. The given WriteSeeker needs
// to be at the start.
func NewWriter(w io.WriteSeeker) (*Writer, error) {
	erofsWriter := &Writer{
		w:                 w,
		fixDirectoryEntry: make(map[string][]direntFixupLocation),
		pathInodeMeta:     make(map[string]*uncompressedInodeMeta),
	}
	_, err := erofsWriter.allocateMetadata(1024+binary.Size(&superblock{}), 0)
	if err != nil {
		return nil, fmt.Errorf("cannot allocate first metadata block: %w", err)
	}
	if _, err := erofsWriter.w.Write(make([]byte, 1024)); err != nil { // Padding
		return nil, fmt.Errorf("failed to write initial padding: %w", err)
	}
	if err := binary.Write(erofsWriter.w, binary.LittleEndian, &superblock{
		Magic:         Magic,
		BlockSizeBits: blockSizeBits,
		// 1024 (padding) + 128 (superblock) / 32, not eligible for fixup as
		// different int size
		RootNodeNumber: 36,
	}); err != nil {
		return nil, fmt.Errorf("failed to write superblock: %w", err)
	}
	return erofsWriter, nil
}

// allocateMetadata allocates metadata space of size bytes with a given
// alignment and seeks to the first byte of the newly-allocated metadata space.
// It also returns the position of that first byte.
func (w *Writer) allocateMetadata(size int, alignment uint16) (int64, error) {
	if size > BlockSize {
		panic("cannot allocate a metadata object bigger than BlockSize bytes")
	}
	sizeU16 := uint16(size)
	pos, ok := w.metadataBlocksFree.findBlock(sizeU16, 32)
	if !ok {
		blockNumber, err := w.allocateBlocks(1)
		if err != nil {
			return 0, fmt.Errorf("failed to allocate additional metadata space: %w", err)
		}
		w.metadataBlocksFree = append(w.metadataBlocksFree, metadataBlockMeta{blockNumber: blockNumber, freeBytes: BlockSize - sizeU16})
		if _, err := w.w.Write(make([]byte, BlockSize)); err != nil {
			return 0, fmt.Errorf("failed to write metadata: %w", err)
		}
		pos = int64(blockNumber) * BlockSize // Always aligned to BlockSize, bigger alignments are unsupported anyways
	}
	if _, err := w.w.Seek(pos, io.SeekStart); err != nil {
		return 0, fmt.Errorf("cannot seek to existing metadata nid, likely misaligned meta write")
	}
	return pos, nil
}

// allocateBlocks allocates n new BlockSize-sized block and seeks to the
// beginning of the first newly-allocated block.  It also returns the first
// newly-allocated block number.  The caller is expected to write these blocks
// completely before calling allocateBlocks again.
func (w *Writer) allocateBlocks(n uint32) (uint32, error) {
	if _, err := w.w.Seek(int64(w.blockAllocatorIndex)*BlockSize, io.SeekStart); err != nil {
		return 0, fmt.Errorf("cannot seek to end of last block, check write alignment: %w", err)
	}
	firstBlock := w.blockAllocatorIndex
	w.blockAllocatorIndex += n
	return firstBlock, nil
}

func (w *Writer) create(pathname string, inode Inode) *uncompressedInodeWriter {
	i := &uncompressedInodeWriter{
		writer:            w,
		inode:             *inode.inode(),
		legacyInodeNumber: w.legacyInodeIndex,
		pathname:          path.Clean(pathname),
	}
	w.legacyInodeIndex++
	return i
}

// CreateFile adds a new file to the EROFS. It returns a WriteCloser to which
// the file contents should be written and which then needs to be closed. The
// last writer obtained by calling CreateFile() needs to be closed first before
// opening a new one. The given pathname needs to be referenced by a directory
// created using Create(), otherwise it will not be accessible.
func (w *Writer) CreateFile(pathname string, meta *FileMeta) io.WriteCloser {
	return w.create(pathname, meta)
}

// Create adds a new non-file inode to the EROFS. This includes directories,
// device nodes, symlinks and FIFOs.  The first call to Create() needs to be
// with pathname "." and a directory inode.  The given pathname needs to be
// referenced by a directory, otherwise it will not be accessible (with the
// exception of the directory ".").
func (w *Writer) Create(pathname string, inode Inode) error {
	iw := w.create(pathname, inode)
	switch i := inode.(type) {
	case *Directory:
		if err := i.writeTo(iw); err != nil {
			return fmt.Errorf("failed to write directory contents: %w", err)
		}
	case *SymbolicLink:
		if err := i.writeTo(iw); err != nil {
			return fmt.Errorf("failed to write symbolic link contents: %w", err)
		}
	}
	return iw.Close()
}

// Close finishes writing an EROFS filesystem. Errors by this function need to
// be handled as they indicate if the written filesystem is consistent (i.e.
// there are no directory entries pointing to nonexistent inodes).
func (w *Writer) Close() error {
	for targetPath, entries := range w.fixDirectoryEntry {
		for _, entry := range entries {
			targetMeta, ok := w.pathInodeMeta[targetPath]
			if !ok {
				return fmt.Errorf("failed to link filesystem tree: dangling reference to %v", targetPath)
			}
			if err := direntFixup(w.pathInodeMeta[entry.path], int64(entry.entryIndex), targetMeta); err != nil {
				return err
			}
		}
	}
	return nil
}

// uncompressedInodeMeta tracks enough metadata about a written inode to be
// able to point dirents to it and to provide a WriteSeeker into the inode
// itself.
type uncompressedInodeMeta struct {
	nid   uint64
	ftype uint8

	// Physical placement metdata
	blockStart   int64
	blockLength  int64
	inlineStart  int64
	inlineLength int64

	writer        *Writer
	currentOffset int64
}

func (a *uncompressedInodeMeta) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekCurrent:
		break
	case io.SeekStart:
		a.currentOffset = 0
	case io.SeekEnd:
		a.currentOffset = a.blockLength + a.inlineLength
	}
	a.currentOffset += offset
	return a.currentOffset, nil
}

func (a *uncompressedInodeMeta) Write(p []byte) (int, error) {
	if a.currentOffset < a.blockLength {
		// TODO(lorenz): Handle the special case where a directory inode is
		// spread across multiple blocks (depending on other factors this
		// occurs around ~200 direct children).
		return 0, errors.New("relocating dirents in multi-block directory inodes is unimplemented")
	}
	if _, err := a.writer.w.Seek(a.inlineStart+a.currentOffset, io.SeekStart); err != nil {
		return 0, err
	}
	a.currentOffset += int64(len(p))
	return a.writer.w.Write(p)
}

type direntFixupLocation struct {
	path       string
	entryIndex uint16
}

// direntFixup overrides nid and file type from the path the dirent is pointing
// to. The given iw is expected to be at the start of the dirent inode to be
// fixed up.
func direntFixup(iw io.WriteSeeker, entryIndex int64, meta *uncompressedInodeMeta) error {
	if _, err := iw.Seek(entryIndex*12, io.SeekStart); err != nil {
		return fmt.Errorf("failed to seek to dirent: %w", err)
	}
	if err := binary.Write(iw, binary.LittleEndian, meta.nid); err != nil {
		return fmt.Errorf("failed to write nid: %w", err)
	}
	if _, err := iw.Seek(2, io.SeekCurrent); err != nil { // Skip NameStartOffset
		return fmt.Errorf("failed to seek to dirent: %w", err)
	}
	if err := binary.Write(iw, binary.LittleEndian, meta.ftype); err != nil {
		return fmt.Errorf("failed to write ftype: %w", err)
	}
	return nil
}

type metadataBlockMeta struct {
	blockNumber uint32
	freeBytes   uint16
}

// metadataBlocksMeta contains metadata about all metadata blocks, most
// importantly the amount of free bytes in each block. This is not a map for
// reproducibility (map ordering).
type metadataBlocksMeta []metadataBlockMeta

// findBlock returns the absolute position where `size` bytes with the
// specified alignment can still fit.  If there is not enough space in any
// metadata block it returns false as the second return value.
func (m metadataBlocksMeta) findBlock(size uint16, alignment uint16) (int64, bool) {
	for i, blockMeta := range m {
		freeBytesAligned := blockMeta.freeBytes - (blockMeta.freeBytes % alignment)
		if freeBytesAligned > size {
			m[i] = metadataBlockMeta{
				blockNumber: blockMeta.blockNumber,
				freeBytes:   freeBytesAligned - size,
			}
			pos := int64(blockMeta.blockNumber+1)*BlockSize - int64(freeBytesAligned)
			return pos, true
		}
	}
	return 0, false
}

var unixModeToFTMap = map[uint16]uint8{
	unix.S_IFREG:  fileTypeRegularFile,
	unix.S_IFDIR:  fileTypeDirectory,
	unix.S_IFCHR:  fileTypeCharacterDevice,
	unix.S_IFBLK:  fileTypeBlockDevice,
	unix.S_IFIFO:  fileTypeFIFO,
	unix.S_IFSOCK: fileTypeSocket,
	unix.S_IFLNK:  fileTypeSymbolicLink,
}

// unixModeToFT maps a Unix file type to an EROFS file type.
func unixModeToFT(mode uint16) uint8 {
	return unixModeToFTMap[mode&unix.S_IFMT]
}
