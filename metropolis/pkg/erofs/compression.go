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

// This file contains compression-related functions.
// TODO(lorenz): Fully implement compression. These are currently unused.

import "encoding/binary"

// mapHeader is a legacy but still-used advisory structure at the start of a
// compressed VLE block. It contains constant values as annotated.
type mapHeader struct {
	Reserved      uint32 // 0
	Advise        uint16 // 1
	AlgorithmType uint8  // 0
	ClusterBits   uint8  // 0
}

// encodeSmallVLEBlock encodes two VLE extents into a 8 byte block.
func encodeSmallVLEBlock(vals [2]uint16, blkaddr uint32) [8]byte {
	var out [8]byte
	binary.LittleEndian.PutUint16(out[0:2], vals[0])
	binary.LittleEndian.PutUint16(out[2:4], vals[1])
	binary.LittleEndian.PutUint32(out[4:8], blkaddr)
	return out
}

// encodeBigVLEBlock encodes 16 VLE extents into a 32 byte block.
func encodeBigVLEBlock(vals [16]uint16, blkaddr uint32) [32]byte {
	var out [32]byte
	for i, val := range vals {
		if val > 1<<14 {
			panic("value is bigger than 14 bits, cannot encode")
		}
		// Writes packed 14 bit unsigned integers
		pos := i * 14
		bitStartPos := pos % 8
		byteStartPos := pos / 8
		out[byteStartPos] = out[byteStartPos]&((1<<bitStartPos)-1) | uint8(val<<bitStartPos)
		out[byteStartPos+1] = uint8(val >> (8 - bitStartPos))
		out[byteStartPos+2] = uint8(val >> (16 - bitStartPos))
	}
	binary.LittleEndian.PutUint32(out[28:32], blkaddr)
	return out
}
