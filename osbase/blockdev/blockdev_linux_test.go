// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package blockdev

import (
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

	ValidateBlockDev(t, blk, fileBlockCount, fileBlockSize, fileBlockSize)
}
