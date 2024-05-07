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

func (m *Memory) ReadAt(p []byte, off int64) (int, error) {
	devSize := m.blockSize * m.blockCount
	if off > devSize {
		return 0, io.EOF
	}
	// TODO: Alignment checks?
	copy(p, m.data[off:])
	n := len(m.data[off:])
	if n < len(p) {
		return n, io.EOF
	}
	return len(p), nil
}

func (m *Memory) WriteAt(p []byte, off int64) (int, error) {
	devSize := m.blockSize * m.blockCount
	if off > devSize {
		return 0, io.EOF
	}
	// TODO: Alignment checks?
	copy(m.data[off:], p)
	n := len(m.data[off:])
	if n < len(p) {
		return n, io.EOF
	}
	return len(p), nil
}

func (m *Memory) BlockSize() int64 {
	return m.blockSize
}

func (m *Memory) BlockCount() int64 {
	return m.blockCount
}

func (m *Memory) OptimalBlockSize() int64 {
	return m.blockSize
}

func (m *Memory) validRange(startByte, endByte int64) error {
	if startByte > endByte {
		return fmt.Errorf("startByte (%d) larger than endByte (%d), invalid interval", startByte, endByte)
	}
	devSize := m.blockSize * m.blockCount
	if startByte >= devSize || startByte < 0 {
		return fmt.Errorf("startByte (%d) out of range (0-%d)", endByte, devSize)
	}
	if endByte > devSize || endByte < 0 {
		return fmt.Errorf("endByte (%d) out of range (0-%d)", endByte, devSize)
	}
	// Alignment check works for powers of two by looking at every bit below
	// the bit set in the block size.
	if startByte&(m.blockSize-1) != 0 {
		return fmt.Errorf("startByte (%d) is not aligned to blockSize (%d)", startByte, m.blockSize)
	}
	if endByte&(m.blockSize-1) != 0 {
		return fmt.Errorf("endByte (%d) is not aligned to blockSize (%d)", startByte, m.blockSize)
	}
	return nil
}

func (m *Memory) Discard(startByte, endByte int64) error {
	if err := m.validRange(startByte, endByte); err != nil {
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
	if err := m.validRange(startByte, endByte); err != nil {
		return err
	}
	for i := startByte; i < endByte; i++ {
		m.data[i] = 0x00
	}
	return nil
}
