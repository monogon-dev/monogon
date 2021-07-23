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

// This package implements a subset of the veritysetup tool from cryptsetup,
// which is a userland tool to interact with dm-verity devices/images. It was
// rewritten to provide the minimum of functionality needed for Metropolis
// without having to package, link against and maintain the original C
// veritysetup tool.
//
// dm-verity is a Linux device mapper target that allows integrity verification of
// a read-only block device. The block device whose integrity should be checked
// (the 'data device') must be first processed by a tool like veritysetup (or this
// tool, mkverity) to generate a hash device and root hash.
// The original data device, hash device and root hash are then set up as a device
// mapper target, and any read performed from the data device through the verity
// target will be verified for integrity by Linux using the hash device and root
// hash.
//
// Internally, the hash device is a Merkle tree of all the bytes in the data
// device, layed out as layers of 'hash blocks'. Starting with data bytes, layers
// are built recursively, with each layer's output hash blocks becoming the next
// layer's data input, ending with the single root hash.
//
// For more information about the internals, see the Linux and cryptsetup
// upstream code:
//
// https://gitlab.com/cryptsetup/cryptsetup/wikis/DMVerity#verity-superblock-format
package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

// veritySuperblock represents data layout inside of a dm-verity hash block
// device superblock. It follows a preexisting verity implementation:
//
// https://gitlab.com/cryptsetup/cryptsetup/wikis/DMVerity#verity-superblock-format
type veritySuperblock struct {
	// signature is the magic signature of a verity hash device superblock,
	// "verity\0\0".
	signature [8]uint8
	// version specifies a superblock format. This structure describes version
	// '1'.
	version uint32
	// hashType defaults to '1' outside Chrome OS, according to scarce dm-verity
	// documentation.
	hashType uint32
	// uuid contains a UUID of the hash device.
	uuid [16]uint8
	// algorithm stores an ASCII-encoded name of the hash function used.
	algorithm [32]uint8

	// dataBlockSize specifies a size of a single data device block, in bytes.
	dataBlockSize uint32
	// hashBlockSize specifies a size of a single hash device block, in bytes.
	hashBlockSize uint32
	// dataBlocks contains a count of blocks available on the data device.
	dataBlocks uint64

	// saltSize encodes the size of hash block salt, up to the maximum of 256 bytes.
	saltSize uint16

	// _pad1 is a zeroed space prepending the salt; unused.
	_pad1 [6]uint8
	// exactly saltSize bytes of salt are prepended to data blocks before hashing.
	salt [256]uint8
	// _pad2 is a zeroed space after the salt; unused.
	_pad2 [168]uint8
}

// divideAndRoundup performs an integer division and returns a rounded up
// result. Useful in calculating block counts.
func divideAndRoundup(a, b uint64) uint64 {
	r := a / b
	if a%b != 0 {
		r++
	}
	return r
}

// newSuperblock builds a dm-verity hash device superblock based on the size
// of data image file reachable through dataImagePath.
// It returns either a fully initialized veritySuperblock, or an
// initialization error.
func newSuperblock(dataImagePath string) (*veritySuperblock, error) {
	// This implementation only handles SHA256-based verity hash images
	// with a specific 4096-byte block size.
	// Block sizes can be updated by adjusting the struct literal below.
	// A change of a hashing algorithm would require a refactor of
	// saltedDigest, and references to sha256.Size.
	//
	// Fill in the defaults (compare with veritySuperblock definition).
	sb := veritySuperblock{
		signature:     [8]uint8{'v', 'e', 'r', 'i', 't', 'y', 0, 0},
		version:       1,
		hashType:      1,
		algorithm:     [32]uint8{'s', 'h', 'a', '2', '5', '6'},
		saltSize:      256,
		dataBlockSize: 4096,
		hashBlockSize: 4096,
	}

	// Get the data image size and compute the data block count.
	ds, err := os.Stat(dataImagePath)
	if err != nil {
		return nil, fmt.Errorf("while stat-ing data device: %w", err)
	}
	if !ds.Mode().IsRegular() {
		return nil, fmt.Errorf("this program only accepts regular files")
	}
	sb.dataBlocks = divideAndRoundup(uint64(ds.Size()), uint64(sb.dataBlockSize))

	// Fill in the superblock UUID and cryptographic salt.
	if _, err := rand.Read(sb.uuid[:]); err != nil {
		return nil, fmt.Errorf("when generating UUID: %w", err)
	}
	if _, err := rand.Read(sb.salt[:]); err != nil {
		return nil, fmt.Errorf("when generating salt: %w", err)
	}

	return &sb, nil
}

// saltedDigest computes and returns a SHA256 sum of a block prepended
// with a Superblock-defined salt.
func (sb *veritySuperblock) saltedDigest(data []byte) (digest [sha256.Size]byte) {
	h := sha256.New()
	h.Write(sb.salt[:int(sb.saltSize)])
	h.Write(data)
	copy(digest[:], h.Sum(nil))
	return
}

// dataBlocksPerHashBlock returns the count of hash operation outputs that
// fit in a hash device block. This is also the amount of data device
// blocks it takes to populate a hash device block.
func (sb *veritySuperblock) dataBlocksPerHashBlock() uint64 {
	return uint64(sb.hashBlockSize) / sha256.Size
}

// computeHashBlock reads at most sb.dataBlocksPerHashBlock blocks from
// the given reader object, returning a padded hash block of length
// defined by sb.hashBlockSize, and an error, if encountered.
// In case a non-nil block is returned, it's guaranteed to contain at
// least one hash. An io.EOF signals that there is no more to be read
// from 'r'.
func (sb *veritySuperblock) computeHashBlock(r io.Reader) ([]byte, error) {
	// Preallocate a whole hash block.
	hblk := bytes.NewBuffer(make([]byte, 0, sb.hashBlockSize))

	// For every data block, compute a hash and place it in hblk. Continue
	// till EOF.
	for b := uint64(0); b < sb.dataBlocksPerHashBlock(); b++ {
		dbuf := make([]byte, sb.dataBlockSize)
		// Attempt to read enough data blocks to make a complete hash block.
		n, err := io.ReadFull(r, dbuf)
		// If any data was read, make a hash and add it to the hash buffer.
		if n != 0 {
			hash := sb.saltedDigest(dbuf)
			hblk.Write(hash[:])
		}
		// Handle the read errors.
		switch err {
		case nil:
		case io.ErrUnexpectedEOF, io.EOF:
			// io.ReadFull returns io.ErrUnexpectedEOF after a partial read,
			// and io.EOF if no bytes were read. In both cases it's possible
			// to end up with a partially filled hash block.
			if hblk.Len() != 0 {
				// Return a zero-padded hash block if any hashes were written
				// to it, and signal that no more blocks can be built.
				res := hblk.Bytes()
				return res[:cap(res)], io.EOF
			}
			// Return nil if the block doesn't contain any hashes.
			return nil, io.EOF
		default:
			// Wrap unhandled read errors.
			return nil, fmt.Errorf("while computing a hash block: %w", err)
		}
	}
	// Return a completely filled hash block.
	res := hblk.Bytes()
	return res[:cap(res)], nil
}

// writeSuperblock writes a verity superblock to a given writer object.
// It returns a write error, if encountered.
func (sb *veritySuperblock) writeSuperblock(w io.Writer) error {
	// Write the superblock.
	if err := binary.Write(w, binary.LittleEndian, sb); err != nil {
		return fmt.Errorf("while writing a header: %w", err)
	}

	// Get the padding size by substracting current offset from a hash block
	// size.
	co := binary.Size(sb)
	pbc := int(sb.hashBlockSize) - co
	if pbc <= 0 {
		return fmt.Errorf("hash device block size smaller than dm-verity superblock")
	}

	// Write the padding bytes at the end of the block.
	if _, err := w.Write(bytes.Repeat([]byte{0}, pbc)); err != nil {
		return fmt.Errorf("while writing padding: %w", err)
	}
	return nil
}

// computeLevelZero produces the base level of a hash tree. It's the only
// level calculated based on raw input from the data image.
// It returns a byte slice containing one or more hash blocks, depending
// on sb.dataBlocks and sb.hashBlockSize, or an error. The returned slice
// length is guaranteed to be a multiple of sb.hashBlockSize if no error
// is returned.
// BUG(mz): Current implementation requires a 1/128th of the data image
// size to be allocatable on the heap.
func (sb *veritySuperblock) computeLevel(r io.Reader) ([]byte, error) {
	// hbuf will store all the computed hash blocks.
	var hbuf bytes.Buffer
	// Compute one or more hash blocks, reading all data available in the
	// 'r' reader object, and write them into hbuf.
	for {
		hblk, err := sb.computeHashBlock(r)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("while building a hash tree level: %w", err)
		}
		if hblk != nil {
			_, err := hbuf.Write(hblk)
			if err != nil {
				return nil, fmt.Errorf("while writing to hash block buffer: %w", err)
			}
		}
		if err == io.EOF {
			break
		}
	}
	return hbuf.Bytes(), nil
}

// computeHashTree builds a complete hash tree based on the given reader
// object. Levels are appended to resulting hashTree from bottom to top.
// It returns a verity hash tree, a verity root hash, and an error, if
// encountered.
func (sb *veritySuperblock) computeHashTree(r io.Reader) ([][]byte, []byte, error) {
	// First, hash contents of the data image. This will result in a bottom
	// level of the hash tree.
	var hashTree [][]byte
	lz, err := sb.computeLevel(r)
	if err != nil {
		return nil, nil, fmt.Errorf("while computing the base level: %w", err)
	}
	hashTree = append(hashTree, lz)

	// Other levels are built by hashing the hash blocks comprising a level
	// below.
	for {
		// Create the next level by hashing the previous one.
		pl := hashTree[len(hashTree)-1]
		nl, err := sb.computeLevel(bytes.NewReader(pl))
		if err != nil {
			return nil, nil, fmt.Errorf("while computing a level: %w", err)
		}
		// Append the resulting next level to a tree.
		hashTree = append(hashTree, nl)

		if len(nl) == int(sb.hashBlockSize) {
			// The last level to compute has a size of exactly one hash block.
			// That's the root level. Its hash serves as a cryptographic root of
			// trust and is returned separately.
			rootHash := sb.saltedDigest(nl)
			return hashTree, rootHash[:], nil
		}
	}
}

// writeHashTree writes a verity-formatted hash tree to the given writer
// object. Compare with computeHashTree.
// It returns the count of bytes written and a write error, if encountered.
func (sb *veritySuperblock) writeHashTree(w io.Writer, treeLevels [][]byte) error {
	// Write the hash tree levels from top to bottom.
	for l := len(treeLevels) - 1; l >= 0; l-- {
		level := treeLevels[l]
		// Call w.Write until a whole level is written.
		for len(level) != 0 {
			n, err := w.Write(level)
			if err != nil && err != io.ErrShortWrite {
				return fmt.Errorf("while writing a level: %w", err)
			}
			level = level[n:]
		}
	}
	return nil
}

// createHashImage creates a complete dm-verity hash image at
// hashImagePath. Contents of the file at dataImagePath are accessed
// read-only, hashed and written to the hash image in the process.
// It returns a verity root hash and an error, if encountered.
func createHashImage(dataImagePath, hashImagePath string) ([]byte, error) {
	// Inspect the data image and build a verity superblock based on its size.
	sb, err := newSuperblock(dataImagePath)
	if err != nil {
		return nil, fmt.Errorf("while building a superblock: %w", err)
	}

	// Open the data image for reading.
	dataImage, err := os.Open(dataImagePath)
	if err != nil {
		return nil, fmt.Errorf("while opening the data image: %w", err)
	}
	defer dataImage.Close()

	// Create an empty hash image file.
	hashImage, err := os.OpenFile(hashImagePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("while opening the hash image for writing: %w", err)
	}
	defer hashImage.Close()

	// Write the superblock to the hash image.
	if err = sb.writeSuperblock(hashImage); err != nil {
		return nil, fmt.Errorf("while writing the superblock: %w", err)
	}

	// Compute a verity hash tree by hashing contents of the data image. Then,
	// write it to the hash image.
	treeLevels, rootHash, err := sb.computeHashTree(dataImage)
	if err != nil {
		return nil, fmt.Errorf("while building a hash tree: %w", err)
	}
	if err = sb.writeHashTree(hashImage, treeLevels); err != nil {
		return nil, fmt.Errorf("while writing a hash tree: %w", err)
	}

	// Return a verity root hash, serving as a root of trust.
	return rootHash, nil
}

// usage prints program usage information.
func usage(executable string) {
	fmt.Println("Usage: ", executable, " format <data image> <hash image>")
}

func main() {
	// Process the command line arguments maintaining a partial
	// compatibility with veritysetup.
	if len(os.Args) != 4 {
		usage(os.Args[0])
		os.Exit(2)
	}
	command := os.Args[1]
	dataImagePath := os.Args[2]
	hashImagePath := os.Args[3]

	switch command {
	case "format":
		rootHash, err := createHashImage(dataImagePath, hashImagePath)
		if err != nil {
			log.Fatal(err)
		}
		// The output differs from the original veritysetup utility in that hash
		// isn't prepended by "Root hash: " string. It's left this way to
		// facilitate machine processing.
		fmt.Printf("%x", rootHash)
	default:
		usage(os.Args[0])
		os.Exit(2)
	}
}
