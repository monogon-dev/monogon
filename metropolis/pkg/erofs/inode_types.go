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
	"math"
	"path"
	"sort"

	"golang.org/x/sys/unix"
)

// Inode specifies an interface that all inodes that can be written to an EROFS
// filesystem implement.
type Inode interface {
	inode() *inodeCompact
}

// Base contains generic inode metadata independent from the specific inode
// type.
type Base struct {
	Permissions uint16
	UID, GID    uint16
}

func (b *Base) baseInode(fileType uint16) *inodeCompact {
	return &inodeCompact{
		UID:  b.UID,
		GID:  b.GID,
		Mode: b.Permissions | fileType,
	}
}

// Directory represents a directory inode. The Children property contains the
// directories' direct children (just the name, not the full path).
type Directory struct {
	Base
	Children []string
}

func (d *Directory) inode() *inodeCompact {
	return d.baseInode(unix.S_IFDIR)
}

func (d *Directory) writeTo(w *uncompressedInodeWriter) error {
	// children is d.Children with appended backrefs (. and ..), copied to not
	// pollute source
	children := make([]string, len(d.Children))
	copy(children, d.Children)
	children = append(children, ".", "..")
	sort.Strings(children)

	nameStartOffset := binary.Size(directoryEntryRaw{}) * len(children)
	var rawEntries []directoryEntryRaw
	for _, ent := range children {
		if nameStartOffset > math.MaxUint16 {
			return errors.New("directory name offset out of range, too many or too big entries")
		}
		var entData directoryEntryRaw
		entData.NameStartOffset = uint16(nameStartOffset)
		rawEntries = append(rawEntries, entData)
		nameStartOffset += len(ent)
	}
	for i, ent := range rawEntries {
		targetPath := path.Join(w.pathname, children[i])
		if targetPath == ".." {
			targetPath = "."
		}
		w.writer.fixDirectoryEntry[targetPath] = append(w.writer.fixDirectoryEntry[targetPath], direntFixupLocation{
			path:       w.pathname,
			entryIndex: uint16(i),
		})
		if err := binary.Write(w, binary.LittleEndian, ent); err != nil {
			return fmt.Errorf("failed to write dirent: %w", err)
		}
	}
	for _, childName := range children {
		if _, err := w.Write([]byte(childName)); err != nil {
			return fmt.Errorf("failed to write dirent name: %w", err)
		}
	}
	return nil
}

// CharacterDevice represents a Unix character device inode with major and
// minor numbers.
type CharacterDevice struct {
	Base
	Major uint32
	Minor uint32
}

func (c *CharacterDevice) inode() *inodeCompact {
	i := c.baseInode(unix.S_IFCHR)
	i.Union = uint32(unix.Mkdev(c.Major, c.Minor))
	return i
}

// CharacterDevice represents a Unix block device inode with major and minor
// numbers.
type BlockDevice struct {
	Base
	Major uint32
	Minor uint32
}

func (b *BlockDevice) inode() *inodeCompact {
	i := b.baseInode(unix.S_IFBLK)
	i.Union = uint32(unix.Mkdev(b.Major, b.Minor))
	return i
}

// FIFO represents a Unix FIFO inode.
type FIFO struct {
	Base
}

func (f *FIFO) inode() *inodeCompact {
	return f.baseInode(unix.S_IFIFO)
}

// Socket represents a Unix socket inode.
type Socket struct {
	Base
}

func (s *Socket) inode() *inodeCompact {
	return s.baseInode(unix.S_IFSOCK)
}

// SymbolicLink represents a symbolic link/symlink to another inode. Target is
// the literal string target of the symlink.
type SymbolicLink struct {
	Base
	Target string
}

func (s *SymbolicLink) inode() *inodeCompact {
	return s.baseInode(unix.S_IFLNK)
}

func (s *SymbolicLink) writeTo(w io.Writer) error {
	_, err := w.Write([]byte(s.Target))
	return err
}

// FileMeta represents the metadata of a regular file. In this case the
// contents are written to a Writer returned by the CreateFile function on the
// EROFS Writer and not included in the structure itself.
type FileMeta struct {
	Base
}

func (f *FileMeta) inode() *inodeCompact {
	return f.baseInode(unix.S_IFREG)
}
