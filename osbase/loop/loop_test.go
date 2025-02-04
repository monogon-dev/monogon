// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package loop

import (
	"encoding/binary"
	"io"
	"math"
	"os"
	"syscall"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

// Write a test file with a very specific pattern (increasing little-endian 16
// bit unsigned integers) to detect offset correctness. File is always 128KiB
// large (2^16 * 2 bytes).
func makeTestFile() *os.File {
	f, err := os.CreateTemp("/tmp", "")
	if err != nil {
		panic(err)
	}
	for i := 0; i <= math.MaxUint16; i++ {
		if err := binary.Write(f, binary.LittleEndian, uint16(i)); err != nil {
			panic(err)
		}
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		panic(err)
	}
	return f
}

func getBlkdevSize(f *os.File) (size uint64) {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), unix.BLKGETSIZE64, uintptr(unsafe.Pointer(&size))); err != 0 {
		panic(err)
	}
	return
}

func getOffsetFromContent(dev *Device) (firstIndex uint16) {
	if err := binary.Read(dev.dev, binary.LittleEndian, &firstIndex); err != nil {
		panic(err)
	}
	firstIndex *= 2 // 2 bytes per index
	return
}

func setupCreate(t *testing.T, config Config) *Device {
	f := makeTestFile()
	dev, err := Create(f, config)
	defer f.Close()
	assert.NoError(t, err)
	t.Cleanup(func() {
		if dev != nil {
			dev.Remove()
		}
		os.Remove(f.Name())
	})
	if dev == nil {
		t.FailNow()
	}
	return dev
}

func TestDeviceAccessors(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	dev := setupCreate(t, Config{})

	devPath, err := dev.DevPath()
	assert.NoError(t, err)
	require.Equal(t, "/dev/loop0", devPath)

	var stat unix.Stat_t
	assert.NoError(t, unix.Stat("/dev/loop0", &stat))
	devNum, err := dev.Dev()
	assert.NoError(t, err)
	require.Equal(t, stat.Rdev, devNum)

	backingFile, err := dev.BackingFilePath()
	assert.NoError(t, err)
	// The filename of the temporary file is not available in this context, but
	// we know that the file needs to be in /tmp, which should be a good-enough
	// test.
	assert.Contains(t, backingFile, "/tmp/")
}

func TestCreate(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	t.Parallel()
	tests := []struct {
		name     string
		config   Config
		validate func(t *testing.T, dev *Device)
	}{
		{"NoOpts", Config{}, func(t *testing.T, dev *Device) {
			require.Equal(t, uint64(128*1024), getBlkdevSize(dev.dev))
			require.Equal(t, uint16(0), getOffsetFromContent(dev))

			_, err := dev.dev.WriteString("test")
			assert.NoError(t, err)
		}},
		{"DirectIO", Config{Flags: FlagDirectIO}, func(t *testing.T, dev *Device) {
			require.Equal(t, uint64(128*1024), getBlkdevSize(dev.dev))

			_, err := dev.dev.WriteString("test")
			assert.NoError(t, err)
		}},
		{"ReadOnly", Config{Flags: FlagReadOnly}, func(t *testing.T, dev *Device) {
			_, err := dev.dev.WriteString("test")
			assert.Error(t, err)
		}},
		{"Mapping", Config{BlockSize: 512, SizeLimit: 2048, Offset: 4096}, func(t *testing.T, dev *Device) {
			assert.Equal(t, uint16(4096), getOffsetFromContent(dev))
			assert.Equal(t, uint64(2048), getBlkdevSize(dev.dev))
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dev := setupCreate(t, test.config)
			test.validate(t, dev)
			assert.NoError(t, dev.Remove())
		})
	}
}

func TestOpenBadDevice(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	dev, err := Open("/dev/null")
	require.Error(t, err)
	if dev != nil { // Prevent leaks in case this test fails
		dev.Close()
	}
}

func TestOpen(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	f := makeTestFile()
	defer os.Remove(f.Name())
	defer f.Close()
	dev, err := Create(f, Config{})
	assert.NoError(t, err)
	path, err := dev.DevPath()
	assert.NoError(t, err)
	assert.NoError(t, dev.Close())
	reopenedDev, err := Open(path)
	assert.NoError(t, err)
	defer reopenedDev.Remove()
	reopenedDevPath, err := reopenedDev.DevPath()
	assert.NoError(t, err)
	require.Equal(t, path, reopenedDevPath) // Still needs to be the same device
}

func TestResize(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	f, err := os.CreateTemp("/tmp", "")
	assert.NoError(t, err)
	empty1K := make([]byte, 1024)
	for i := 0; i < 64; i++ {
		_, err := f.Write(empty1K)
		assert.NoError(t, err)
	}
	dev, err := Create(f, Config{})
	assert.NoError(t, err)
	require.Equal(t, uint64(64*1024), getBlkdevSize(dev.dev))
	for i := 0; i < 32; i++ {
		_, err := f.Write(empty1K)
		assert.NoError(t, err)
	}
	assert.NoError(t, f.Sync())
	assert.NoError(t, dev.RefreshSize())
	require.Equal(t, uint64(96*1024), getBlkdevSize(dev.dev))
}
