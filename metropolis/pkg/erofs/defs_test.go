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
