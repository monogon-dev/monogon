package blockdev

import (
	"errors"
	"fmt"
	"io"
	"math/bits"
)

// Memory is a memory-backed implementation of BlockDev. It is optimal
// for testing and temporary use, as it is fast and platform-independent.
type Memory struct {
	blockSize  int64
	blockCount int64
	data       []byte
}

// NewMemory returns a new memory-backed block device with the given geometry.
func NewMemory(blockSize, blockCount int64) (*Memory, error) {
	if blockSize <= 0 {
		return nil, errors.New("block size cannot be zero or negative")
	}
	if bits.OnesCount64(uint64(blockSize)) > 1 {
		return nil, fmt.Errorf("block size must be a power of two (got %d)", blockSize)
	}
	if blockCount < 0 {
		return nil, errors.New("block count cannot be negative")
	}
	return &Memory{
		blockSize:  blockSize,
		blockCount: blockCount,
		data:       make([]byte, blockSize*blockCount),
	}, nil
}

// MustNewMemory works exactly like NewMemory, but panics when NewMemory would
// return an error. Intended for use in tests.
func MustNewMemory(blockSize, blockCount int64) *Memory {
	m, err := NewMemory(blockSize, blockCount)
	if err != nil {
		panic(err)
	}
	return m
}

func (m *Memory) ReadAt(p []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, errors.New("blockdev.Memory.ReadAt: negative offset")
	}
	if off > int64(len(m.data)) {
		return 0, io.EOF
	}
	// TODO: Alignment checks?
	n = copy(p, m.data[off:])
	if n < len(p) {
		err = io.EOF
	}
	return
}

func (m *Memory) WriteAt(p []byte, off int64) (n int, err error) {
	if off < 0 || off > int64(len(m.data)) {
		return 0, ErrOutOfBounds
	}
	// TODO: Alignment checks?
	n = copy(m.data[off:], p)
	if n < len(p) {
		err = ErrOutOfBounds
	}
	return
}

func (m *Memory) BlockCount() int64 {
	return m.blockCount
}

func (m *Memory) BlockSize() int64 {
	return m.blockSize
}

func (m *Memory) OptimalBlockSize() int64 {
	return m.blockSize
}

func (m *Memory) Discard(startByte, endByte int64) error {
	if err := validAlignedRange(m, startByte, endByte); err != nil {
		return err
	}
	for i := startByte; i < endByte; i++ {
		// Intentionally don't set to zero as Discard doesn't guarantee
		// any specific contents. Call Zero if you need this.
		m.data[i] = 0xaa
	}
	return nil
}

func (m *Memory) Zero(startByte, endByte int64) error {
	if err := validAlignedRange(m, startByte, endByte); err != nil {
		return err
	}
	for i := startByte; i < endByte; i++ {
		m.data[i] = 0x00
	}
	return nil
}

func (m *Memory) Sync() error {
	return nil
}
