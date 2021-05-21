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

// This file contains definitions coming from the in-Kernel implementation of
// the EROFS filesystem.  All definitions come from @linux//fs/erofs:erofs_fs.h
// unless stated otherwise.

// Magic contains the 4 magic bytes starting at position 1024 identifying an
// EROFS filesystem.  Defined in @linux//include/uapi/linux/magic.h
// EROFS_SUPER_MAGIC_V1
var Magic = [4]byte{0xe2, 0xe1, 0xf5, 0xe0}

const blockSizeBits = 12
const BlockSize = 1 << blockSizeBits

// Defined in @linux//include/linux:fs_types.h starting at FT_UNKNOWN
const (
	fileTypeUnknown = iota
	fileTypeRegularFile
	fileTypeDirectory
	fileTypeCharacterDevice
	fileTypeBlockDevice
	fileTypeFIFO
	fileTypeSocket
	fileTypeSymbolicLink
)

// Anonymous enum starting at EROFS_INODE_FLAT_PLAIN
const (
	inodeFlatPlain             = 0
	inodeFlatCompressionLegacy = 1
	inodeFlatInline            = 2
	inodeFlatCompression       = 3
)

// struct erofs_dirent
type directoryEntryRaw struct {
	NodeNumber      uint64
	NameStartOffset uint16
	FileType        uint8
	Reserved        uint8
}

// struct erofs_super_block
type superblock struct {
	Magic                [4]byte
	Checksum             uint32
	FeatureCompat        uint32
	BlockSizeBits        uint8
	Reserved0            uint8
	RootNodeNumber       uint16
	TotalInodes          uint64
	BuildTimeSeconds     uint64
	BuildTimeNanoseconds uint32
	Blocks               uint32
	MetaStartAddr        uint32
	SharedXattrStartAddr uint32
	UUID                 [16]byte
	VolumeName           [16]byte
	FeaturesIncompatible uint32
	Reserved1            [44]byte
}

// struct erofs_inode_compact
type inodeCompact struct {
	Format         uint16
	XattrCount     uint16
	Mode           uint16
	HardlinkCount  uint16
	Size           uint32
	Reserved0      uint32
	Union          uint32
	InodeNumCompat uint32
	UID            uint16
	GID            uint16
	Reserved1      uint32
}

// Anonymous enum starting at Z_EROFS_VLE_CLUSTER_TYPE_PLAIN
const (
	vleClusterTypePlain = iota << 12
	vleClusterTypeHead
	vleClusterTypeNonhead
)
