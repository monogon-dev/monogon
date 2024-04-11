package blockdev

import (
	"errors"
	"fmt"
	"io"
)

var ErrNotBlockDevice = errors.New("not a block device")

// BlockDev represents a generic block device made up of equally-sized blocks.
// All offsets and intervals are expressed in bytes and must be aligned to
// BlockSize and are recommended to be aligned to OptimalBlockSize if feasible.
// Unless stated otherwise, intervals are inclusive-exclusive, i.e. the
// start byte is included but the end byte is not.
type BlockDev interface {
	io.ReaderAt
	io.WriterAt
	// BlockSize returns the block size of the block device in bytes. This must
	// be a power of two and is commonly (but not always) either 512 or 4096.
	BlockSize() int64

	// BlockCount returns the number of blocks on the block device or -1 if it
	// is an image with an undefined size.
	BlockCount() int64

	// OptimalBlockSize returns the optimal block size in bytes for aligning
	// to as well as issuing I/O. IO operations with block sizes below this
	// one might incur read-write overhead. This is the larger of the physical
	// block size and a device-reported value if available.
	OptimalBlockSize() int64

	// Discard discards a continuous set of blocks. Discarding means the
	// underlying device gets notified that the data in these blocks is no
	// longer needed. This can improve performance of the device device (as it
	// no longer needs to preserve the unused data) as well as bulk erase
	// operations. This command is advisory and not all implementations support
	// it. The contents of discarded blocks are implementation-defined.
	Discard(startByte int64, endByte int64) error

	// Zero zeroes a continouous set of blocks. On certain implementations this
	// can be significantly faster than just calling Write with zeroes.
	Zero(startByte, endByte int64) error
}

func NewRWS(b BlockDev) *ReadWriteSeeker {
	return &ReadWriteSeeker{b: b}
}

// ReadWriteSeeker provides an adapter implementing ReadWriteSeeker on top of
// a blockdev.
type ReadWriteSeeker struct {
	b       BlockDev
	currPos int64
}

func (s *ReadWriteSeeker) Read(p []byte) (n int, err error) {
	n, err = s.b.ReadAt(p, s.currPos)
	s.currPos += int64(n)
	return
}

func (s *ReadWriteSeeker) Write(p []byte) (n int, err error) {
	n, err = s.b.WriteAt(p, s.currPos)
	s.currPos += int64(n)
	return
}

func (s *ReadWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekCurrent:
		s.currPos += offset
	case io.SeekStart:
		s.currPos = offset
	case io.SeekEnd:
		s.currPos = (s.b.BlockCount() * s.b.BlockSize()) - offset
	}
	return s.currPos, nil
}

var ErrOutOfBounds = errors.New("write out of bounds")

// NewSection returns a new Section, implementing BlockDev over that subset
// of blocks. The interval is inclusive-exclusive.
func NewSection(b BlockDev, startBlock, endBlock int64) *Section {
	return &Section{
		b:          b,
		startBlock: startBlock,
		endBlock:   endBlock,
	}
}

// Section implements BlockDev on a slice of another BlockDev given a startBlock
// and endBlock.
type Section struct {
	b                    BlockDev
	startBlock, endBlock int64
}

func (s *Section) ReadAt(p []byte, off int64) (n int, err error) {
	bOff := off + (s.startBlock * s.b.BlockSize())
	bytesToEnd := (s.endBlock * s.b.BlockSize()) - bOff
	if bytesToEnd <= 0 {
		return 0, io.EOF
	}
	if bytesToEnd < int64(len(p)) {
		return s.b.ReadAt(p[:bytesToEnd], bOff)
	}
	return s.b.ReadAt(p, bOff)
}

func (s *Section) WriteAt(p []byte, off int64) (n int, err error) {
	bOff := off + (s.startBlock * s.b.BlockSize())
	bytesToEnd := (s.endBlock * s.b.BlockSize()) - bOff
	if bytesToEnd <= 0 {
		return 0, ErrOutOfBounds
	}
	if bytesToEnd < int64(len(p)) {
		n, err := s.b.WriteAt(p[:bytesToEnd], off+(s.startBlock*s.b.BlockSize()))
		if err != nil {
			// If an error happened, prioritize that error
			return n, err
		}
		// Otherwise, return ErrOutOfBounds as even short writes must return an
		// error.
		return n, ErrOutOfBounds
	}
	return s.b.WriteAt(p, off+(s.startBlock*s.b.BlockSize()))
}

func (s *Section) BlockCount() int64 {
	return s.endBlock - s.startBlock
}

func (s *Section) BlockSize() int64 {
	return s.b.BlockSize()
}

func (s *Section) inRange(startByte, endByte int64) error {
	if startByte > endByte {
		return fmt.Errorf("invalid range: startByte (%d) bigger than endByte (%d)", startByte, endByte)
	}
	sectionLen := s.BlockCount() * s.BlockSize()
	if startByte >= sectionLen {
		return fmt.Errorf("startByte (%d) out of range (%d)", startByte, sectionLen)
	}
	if endByte > sectionLen {
		return fmt.Errorf("endBlock (%d) out of range (%d)", endByte, sectionLen)
	}
	return nil
}

func (s *Section) Discard(startByte, endByte int64) error {
	if err := s.inRange(startByte, endByte); err != nil {
		return err
	}
	offset := s.startBlock * s.b.BlockSize()
	return s.b.Discard(offset+startByte, offset+endByte)
}

func (s *Section) OptimalBlockSize() int64 {
	return s.b.OptimalBlockSize()
}

func (s *Section) Zero(startByte, endByte int64) error {
	if err := s.inRange(startByte, endByte); err != nil {
		return err
	}
	offset := s.startBlock * s.b.BlockSize()
	return s.b.Zero(offset+startByte, offset+endByte)
}

// GenericZero implements software-based zeroing. This can be used to implement
// Zero when no acceleration is available or desired.
func GenericZero(b BlockDev, startByte, endByte int64) error {
	if startByte%b.BlockSize() != 0 {
		return fmt.Errorf("startByte (%d) needs to be aligned to block size (%d)", startByte, b.BlockSize())
	}
	if endByte%b.BlockSize() != 0 {
		return fmt.Errorf("endByte (%d) needs to be aligned to block size (%d)", endByte, b.BlockSize())
	}
	// Choose buffer size close to 16MiB or the range to be zeroed, whatever
	// is smaller.
	bufSizeTarget := int64(16 * 1024 * 1024)
	if endByte-startByte < bufSizeTarget {
		bufSizeTarget = endByte - startByte
	}
	bufSize := (bufSizeTarget / b.BlockSize()) * b.BlockSize()
	buf := make([]byte, bufSize)
	for i := startByte; i < endByte; i += bufSize {
		if endByte-i < bufSize {
			buf = buf[:endByte-i]
		}
		if _, err := b.WriteAt(buf, i); err != nil {
			return fmt.Errorf("while writing zeroes: %w", err)
		}
	}
	return nil
}
