// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package erofs

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

// uncompressedInodeWriter exposes a io.Write-style interface for a single
// uncompressed inode. It splits the Write-calls into blocks and writes both
// the blocks and inode metadata. It is required to call Close() to ensure
// everything is properly written down before writing another inode.
type uncompressedInodeWriter struct {
	buf               bytes.Buffer
	writer            *Writer
	inode             inodeCompact
	baseBlock         uint32 // baseBlock == 0 implies this inode didn't allocate a block (yet).
	writtenBytes      int
	legacyInodeNumber uint32
	pathname          string
}

func (i *uncompressedInodeWriter) allocateBlock() error {
	bb, err := i.writer.allocateBlocks(1)
	if err != nil {
		return err
	}
	if i.baseBlock == 0 {
		i.baseBlock = bb
	}
	return nil
}

func (i *uncompressedInodeWriter) flush(n int) error {
	if err := i.allocateBlock(); err != nil {
		return err
	}
	slice := i.buf.Next(n)
	if _, err := i.writer.w.Write(slice); err != nil {
		return err
	}
	// Always pad to BlockSize.
	_, err := i.writer.w.Write(make([]byte, BlockSize-len(slice)))
	return err
}

func (i *uncompressedInodeWriter) Write(b []byte) (int, error) {
	i.writtenBytes += len(b)
	if _, err := i.buf.Write(b); err != nil {
		return 0, err
	}
	for i.buf.Len() >= BlockSize {
		if err := i.flush(BlockSize); err != nil {
			return 0, err
		}
	}
	return len(b), nil
}

func (i *uncompressedInodeWriter) Close() error {
	if i.buf.Len() > BlockSize {
		panic("programming error")
	}
	inodeSize := binary.Size(i.inode)
	if i.buf.Len()+inodeSize > BlockSize {
		// Can't fit last part of data inline, write it in its own block.
		if err := i.flush(i.buf.Len()); err != nil {
			return err
		}
	}
	if i.buf.Len() == 0 {
		i.inode.Format = inodeFlatPlain << 1
	} else {
		// Colocate last part of data with inode.
		i.inode.Format = inodeFlatInline << 1
	}
	if i.writtenBytes > math.MaxUint32 {
		return errors.New("inodes bigger than 2^32 need the extended inode format which is unsupported by this library")
	}
	i.inode.Size = uint32(i.writtenBytes)
	if i.baseBlock != 0 {
		i.inode.Union = i.baseBlock
	}
	i.inode.HardlinkCount = 1
	i.inode.InodeNumCompat = i.legacyInodeNumber
	basePos, err := i.writer.allocateMetadata(inodeSize+i.buf.Len(), 32)
	if err != nil {
		return fmt.Errorf("failed to allocate metadata: %w", err)
	}
	i.writer.pathInodeMeta[i.pathname] = &uncompressedInodeMeta{
		nid:          uint64(basePos) / 32,
		ftype:        unixModeToFT(i.inode.Mode),
		blockStart:   int64(i.baseBlock),
		blockLength:  (int64(i.writtenBytes) / BlockSize) * BlockSize,
		inlineStart:  basePos + 32,
		inlineLength: int64(i.buf.Len()),
		writer:       i.writer,
	}
	if err := binary.Write(i.writer.w, binary.LittleEndian, &i.inode); err != nil {
		return err
	}
	if i.inode.Format&(inodeFlatInline<<1) != 0 {
		// Data colocated in inode, if any.
		_, err := i.writer.w.Write(i.buf.Bytes())
		return err
	}
	return nil
}
