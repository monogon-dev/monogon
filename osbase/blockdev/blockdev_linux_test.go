// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package blockdev

import (
	"io"
	"os"
	"testing"

	"source.monogon.dev/osbase/loop"
)

const loopBlockSize = 1024
const loopBlockCount = 8

func TestLoopDevice(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	underlying, err := os.CreateTemp("/tmp", "")
	if err != nil {
		t.Fatalf("CreateTemp failed: %v", err)
	}
	defer os.Remove(underlying.Name())

	_, err = underlying.Write(make([]byte, loopBlockSize*loopBlockCount))
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	loopDev, err := loop.Create(underlying, loop.Config{
		BlockSize: loopBlockSize,
	})
	if err != nil {
		t.Fatalf("loop.Create failed: %v", err)
	}
	defer loopDev.Remove()

	devPath, err := loopDev.DevPath()
	if err != nil {
		t.Fatalf("loopDev.DevPath failed: %v", err)
	}

	loopDev.Close()
	blk, err := Open(devPath)
	if err != nil {
		t.Fatalf("Failed to open loop device: %v", err)
	}
	defer blk.Close()

	ValidateBlockDev(t, blk, loopBlockCount, loopBlockSize, loopBlockSize)
}

const fileBlockSize = 1024
const fileBlockCount = 8

func TestFile(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	blk, err := CreateFile("/tmp/testfile", fileBlockSize, fileBlockCount)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	defer os.Remove("/tmp/testfile")
	defer blk.Close()

	ValidateBlockDev(t, blk, fileBlockCount, fileBlockSize, fileBlockSize)

	// ReadFromAt
	srcFile, err := os.Create("/tmp/copysrc")
	if err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}
	defer os.Remove("/tmp/copysrc")
	defer srcFile.Close()
	var size int64 = fileBlockSize * fileBlockCount
	readFromAtTests := []struct {
		name   string
		offset int64
		data   string
		limit  int64
		ok     bool
	}{
		{"empty start", 0, "", -1, true},
		{"empty end", size, "", -1, true},
		{"normal", 3, "abcdef", -1, true},
		{"limited", 3, "abcdef", 4, true},
		{"large limit", 3, "abcdef", size, true},
		{"ends at the end", size - 4, "abcd", -1, true},
		{"ends past the end", size - 4, "abcde", -1, false},
		{"ends past the end with limit", size - 4, "abcde", 10, false},
		{"offset negative", -1, "abc", -1, false},
		{"starts at the end", size, "abc", -1, false},
		{"starts past the end", size + 4, "abc", -1, false},
	}
	for _, tt := range readFromAtTests {
		t.Run("readFromAt "+tt.name, func(t *testing.T) {
			checkBlockDevOp(t, blk, func(content []byte) {
				// Prepare source file
				err = srcFile.Truncate(0)
				if err != nil {
					t.Fatalf("Failed to truncate source file: %v", err)
				}
				_, err = srcFile.WriteAt([]byte("123"+tt.data), 0)
				if err != nil {
					t.Fatalf("Failed to write source file: %v", err)
				}
				_, err = srcFile.Seek(3, io.SeekStart)
				if err != nil {
					t.Fatalf("Failed to seek source file: %v", err)
				}

				// Do ReadFromAt
				r := io.Reader(srcFile)
				lr := &io.LimitedReader{R: srcFile, N: tt.limit}
				if tt.limit != -1 {
					r = lr
				}
				n, err := blk.ReadFromAt(r, tt.offset)
				if (err == nil) != tt.ok {
					t.Errorf("expected error %v, got %v", tt.ok, err)
				}
				expectedN := 0
				if tt.offset >= 0 && tt.offset < size {
					c := content[tt.offset:]
					if tt.limit != -1 && tt.limit < int64(len(c)) {
						c = c[:tt.limit]
					}
					expectedN = copy(c, tt.data)
				}
				if n != int64(expectedN) {
					t.Errorf("got n = %d, expected %d; err: %v", n, expectedN, err)
				}

				// Check new offset
				newOffset, err := srcFile.Seek(0, io.SeekCurrent)
				if err != nil {
					t.Fatalf("Failed to get source file position: %v", err)
				}
				newOffset -= 3
				minOffset := n
				maxOffset := n
				if !tt.ok {
					maxOffset = int64(len(tt.data))
					if tt.limit != -1 {
						maxOffset = min(maxOffset, tt.limit)
					}
				}
				if minOffset > newOffset || newOffset > maxOffset {
					t.Errorf("Got newOffset = %d, expected between %d and %d", newOffset, minOffset, maxOffset)
				}
				remaining := tt.limit - newOffset
				if tt.limit != -1 && lr.N != remaining {
					t.Errorf("Got lr.N = %d, expected %d", lr.N, remaining)
				}
			})
		})
	}
}
