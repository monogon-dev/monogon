// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package erofs

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

// These test that the specified structures serialize to the same number of
// bytes as the ones in the EROFS kernel module.

func TestSuperblockSize(t *testing.T) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, &superblock{}); err != nil {
		t.Fatalf("failed to write superblock: %v", err)
	}
	assert.Equal(t, 128, buf.Len())
}

func TestDirectoryEntrySize(t *testing.T) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, &directoryEntryRaw{}); err != nil {
		t.Fatalf("failed to write directory entry: %v", err)
	}
	assert.Equal(t, 12, buf.Len())
}

func TestInodeCompactSize(t *testing.T) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, &inodeCompact{}); err != nil {
		t.Fatalf("failed to write compact inode: %v", err)
	}
	assert.Equal(t, 32, buf.Len())
}
