// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package blockdev

import (
	"bytes"
	"io"
	"os"
	"slices"
	"testing"
)

func errIsNil(err error) bool {
	return err == nil
}
func errIsEOF(err error) bool {
	return err == io.EOF
}
func errIsReadFailed(err error) bool {
	return err != nil && err != io.EOF
}

// ValidateBlockDev tests all methods of the BlockDev interface. This way, all
// implementations can be tested without duplicating test code. This also
// ensures that all implementations behave consistently.
func ValidateBlockDev(t *testing.T, blk BlockDev, blockCount, blockSize, optimalBlockSize int64) {
	if blk.BlockCount() != blockCount {
		t.Errorf("Expected block count %d, got %d", blockCount, blk.BlockCount())
	}
	if blk.BlockSize() != blockSize {
		t.Errorf("Expected block size %d, got %d", blockSize, blk.BlockSize())
	}
	if blk.OptimalBlockSize() != optimalBlockSize {
		t.Errorf("Expected optimal block size %d, got %d", optimalBlockSize, blk.OptimalBlockSize())
	}
	size := blockCount * blockSize

	// ReadAt
	checkBlockDevOp(t, blk, func(content []byte) {
		normalErr := errIsNil
		if size < 3+8 {
			normalErr = errIsEOF
		}
		readAtTests := []struct {
			name        string
			offset, len int64
			expectedErr func(error) bool
		}{
			{"empty start", 0, 0, errIsNil},
			{"empty end", size, 0, errIsNil},
			{"normal", 3, 8, normalErr},
			{"ends past the end", 1, size, errIsEOF},
			{"offset negative", -1, 2, errIsReadFailed},
			{"starts at the end", size, 8, errIsEOF},
			{"starts past the end", size + 4, 8, errIsEOF},
		}
		for _, tt := range readAtTests {
			t.Run("read "+tt.name, func(t *testing.T) {
				buf := make([]byte, tt.len)
				n, err := blk.ReadAt(buf, tt.offset)
				if !tt.expectedErr(err) {
					t.Errorf("unexpected error %v", err)
				}
				expected := []byte{}
				if tt.offset >= 0 && tt.offset <= size {
					expected = content[tt.offset:min(tt.offset+tt.len, size)]
				}
				if n != len(expected) {
					t.Errorf("got n = %d, expected %d", n, len(expected))
				}
				if !slices.Equal(buf[:n], expected) {
					t.Errorf("read unexpected data")
				}
			})
		}
	})

	// WriteAt
	writeAtTests := []struct {
		name   string
		offset int64
		data   string
		ok     bool
	}{
		{"empty start", 0, "", true},
		{"empty end", size, "", true},
		{"normal", 3, "abcdef", size >= 9},
		{"ends at the end", size - 4, "abcd", size >= 4},
		{"ends past the end", size - 4, "abcde", false},
		{"offset negative", -1, "abc", false},
		{"starts at the end", size, "abc", false},
		{"starts past the end", size + 4, "abc", false},
	}
	for _, tt := range writeAtTests {
		t.Run("write "+tt.name, func(t *testing.T) {
			checkBlockDevOp(t, blk, func(content []byte) {
				n, err := blk.WriteAt([]byte(tt.data), tt.offset)
				if (err == nil) != tt.ok {
					t.Errorf("expected error %v, got %v", tt.ok, err)
				}
				expectedN := 0
				if tt.offset >= 0 && tt.offset < size {
					expectedN = copy(content[tt.offset:], tt.data)
				}
				if n != expectedN {
					t.Errorf("got n = %d, expected %d; err: %v", n, expectedN, err)
				}
			})
		})
	}

	// Zero
	zeroTests := []struct {
		name       string
		start, end int64
		ok         bool
	}{
		{"empty range start", 0, 0, true},
		{"empty range end", size, size, true},
		{"full", 0, size, true},
		{"partial", blockSize, 3 * blockSize, blockCount >= 3},
		{"negative start", -blockSize, blockSize, false},
		{"start after end", 2 * blockSize, blockSize, false},
		{"unaligned start", 1, blockSize, false},
		{"unaligned end", 0, 1, false},
	}
	for _, tt := range zeroTests {
		t.Run("zero "+tt.name, func(t *testing.T) {
			checkBlockDevOp(t, blk, func(content []byte) {
				err := blk.Zero(tt.start, tt.end)
				if (err == nil) != tt.ok {
					t.Errorf("expected error %v, got %v", tt.ok, err)
				}
				if tt.ok {
					for i := tt.start; i < tt.end; i++ {
						content[i] = 0
					}
				}
			})
		})
	}

	// Discard
	for _, tt := range zeroTests {
		t.Run("discard "+tt.name, func(t *testing.T) {
			checkBlockDevOp(t, blk, func(content []byte) {
				err := blk.Discard(tt.start, tt.end)
				if (err == nil) != tt.ok {
					t.Errorf("expected error %v, got %v", tt.ok, err)
				}
				if tt.ok {
					n, err := blk.ReadAt(content[tt.start:tt.end], tt.start)
					if n != int(tt.end-tt.start) {
						t.Errorf("read returned %d, %v", n, err)
					}
				}
			})
		})
	}

	// Sync
	checkBlockDevOp(t, blk, func(content []byte) {
		err := blk.Sync()
		if err != nil {
			t.Errorf("Sync failed: %v", err)
		}
	})
}

// checkBlockDevOp is a helper for testing operations on a blockdev. It fills
// the blockdev with a pattern, then calls f with a slice containing the
// pattern, and afterwards reads the blockdev to compare against the expected
// content. f should modify the slice to the content expected afterwards.
func checkBlockDevOp(t *testing.T, blk BlockDev, f func(content []byte)) {
	t.Helper()

	testContent := make([]byte, blk.BlockCount()*blk.BlockSize())
	for i := range testContent {
		testContent[i] = '1' + byte(i%9)
	}
	n, err := blk.WriteAt(testContent, 0)
	if n != len(testContent) || err != nil {
		t.Fatalf("WriteAt = %d, %v; expected %d, nil", n, err, len(testContent))
	}
	f(testContent)
	afterContent := make([]byte, len(testContent))
	n, err = blk.ReadAt(afterContent, 0)
	if n != len(afterContent) || (err != nil && err != io.EOF) {
		t.Fatalf("ReadAt = %d, %v; expected %d, (nil or EOF)", n, err, len(afterContent))
	}
	if !slices.Equal(afterContent, testContent) {
		t.Errorf("Unexpected content after operation")
	}
}

func TestReadWriteSeeker_Seek(t *testing.T) {
	// Verifies that NewRWS's Seeker behaves like bytes.NewReader
	br := bytes.NewReader([]byte("foobar"))
	m := MustNewMemory(2, 3)
	rws := NewRWS(m)
	n, err := rws.Write([]byte("foobar"))
	if n != 6 || err != nil {
		t.Errorf("Write = %v, %v; want 6, nil", n, err)
	}

	for _, whence := range []int{io.SeekStart, io.SeekCurrent, io.SeekEnd} {
		for offset := int64(-7); offset <= 7; offset++ {
			brOff, brErr := br.Seek(offset, whence)
			rwsOff, rwsErr := rws.Seek(offset, whence)
			if (brErr != nil) != (rwsErr != nil) || brOff != rwsOff {
				t.Errorf("For whence %d, offset %d: bytes.Reader.Seek = (%v, %v) != ReadWriteSeeker.Seek = (%v, %v)",
					whence, offset, brOff, brErr, rwsOff, rwsErr)
			}
		}
	}

	// And verify we can just seek past the end and get an EOF
	got, err := rws.Seek(100, io.SeekStart)
	if err != nil || got != 100 {
		t.Errorf("Seek = %v, %v; want 100, nil", got, err)
	}

	n, err = rws.Read(make([]byte, 10))
	if n != 0 || err != io.EOF {
		t.Errorf("Read = %v, %v; want 0, EOF", n, err)
	}
}

func TestNewSection(t *testing.T) {
	tests := []struct {
		name         string
		blockSize    int64
		count        int64
		startBlock   int64
		endBlock     int64
		ok           bool
		sectionCount int64
	}{
		{name: "empty underlying", blockSize: 8, count: 0, startBlock: 0, endBlock: 0, ok: true, sectionCount: 0},
		{name: "empty section", blockSize: 8, count: 5, startBlock: 2, endBlock: 2, ok: true, sectionCount: 0},
		{name: "partial section", blockSize: 8, count: 15, startBlock: 1, endBlock: 11, ok: true, sectionCount: 10},
		{name: "full section", blockSize: 8, count: 15, startBlock: 0, endBlock: 15, ok: true, sectionCount: 15},
		{name: "negative start", blockSize: 8, count: 15, startBlock: -1, endBlock: 11, ok: false},
		{name: "start after end", blockSize: 8, count: 15, startBlock: 6, endBlock: 5, ok: false},
		{name: "end out of bounds", blockSize: 8, count: 15, startBlock: 6, endBlock: 16, ok: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MustNewMemory(tt.blockSize, tt.count)
			s, err := NewSection(m, tt.startBlock, tt.endBlock)
			if (err == nil) != tt.ok {
				t.Errorf("NewSection: expected %v, got %v", tt.ok, err)
			}
			if err == nil {
				checkBlockDevOp(t, m, func(content []byte) {
					ValidateBlockDev(t, s, tt.sectionCount, tt.blockSize, tt.blockSize)

					// Check that content outside the section has not changed.
					start := tt.startBlock * tt.blockSize
					end := tt.endBlock * tt.blockSize
					n, err := m.ReadAt(content[start:end], start)
					if n != int(end-start) {
						t.Errorf("read returned %d, %v", n, err)
					}
				})
			}
		})
	}
}

type MemoryWithGenericZero struct {
	*Memory
}

func (m *MemoryWithGenericZero) Zero(startByte, endByte int64) error {
	return GenericZero(m, startByte, endByte)
}

func TestGenericZero(t *testing.T) {
	if os.Getenv("IN_KTEST") == "true" {
		t.Skip("In ktest")
	}
	// Use size larger than the 16 MiB buffer size in GenericZero.
	blockSize := int64(512)
	blockCount := int64(35 * 1024)
	m, err := NewMemory(blockSize, blockCount)
	if err != nil {
		t.Errorf("NewMemory: %v", err)
	}
	b := &MemoryWithGenericZero{m}
	if err == nil {
		ValidateBlockDev(t, b, blockCount, blockSize, blockSize)
	}
}

func TestNewMemory(t *testing.T) {
	tests := []struct {
		name      string
		blockSize int64
		count     int64
		ok        bool
	}{
		{name: "normal", blockSize: 64, count: 9, ok: true},
		{name: "count 0", blockSize: 8, count: 0, ok: true},
		{name: "count negative", blockSize: 8, count: -1, ok: false},
		{name: "blockSize not a power of 2", blockSize: 9, count: 5, ok: false},
		{name: "blockSize 0", blockSize: 0, count: 5, ok: false},
		{name: "blockSize negative", blockSize: -1, count: 5, ok: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMemory(tt.blockSize, tt.count)
			if (err == nil) != tt.ok {
				t.Errorf("NewMemory: expected %v, got %v", tt.ok, err)
			}
			if err == nil {
				ValidateBlockDev(t, m, tt.count, tt.blockSize, tt.blockSize)
			}
		})
	}
}
