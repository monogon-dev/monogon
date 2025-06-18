// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package blockdev

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrNotBlockDevice = errors.New("not a block device")

// options aggregates all open options for all platforms.
// If these were defined per-platform selecting the right ones per platform
// would require multiple per-platform files at each call site.
type options struct {
	readOnly  bool
	direct    bool
	exclusive bool
}

func (o *options) collect(opts []Option) {
	for _, f := range opts {
		f(o)
	}
}

func (o *options) genericFlags() int {
	if o.readOnly {
		return os.O_RDONLY
	} else {
		return os.O_RDWR
	}
}

type Option func(*options)

// WithReadonly opens the block device read-only. Any write calls will fail.
// Passed as an option to Open.
func WithReadonly(o *options) {
	o.readOnly = true
}

// WithDirect opens the block device bypassing any caching by the kernel.
// Note that additional alignment requirements might be imposed by the
// underlying device.
// Unsupported on non-Linux currently, will return an error.
func WithDirect(o *options) {
	o.direct = true
}

// WithExclusive tries to acquire a pseudo-exclusive lock (only with other
// exclusive FDs) over the block device.
// Unsupported on non-Linux currently, will return an error.
func WithExclusive(o *options) {
	o.exclusive = true
}

// BlockDev represents a generic block device made up of equally-sized blocks.
// All offsets and intervals are expressed in bytes and must be aligned to
// BlockSize and are recommended to be aligned to OptimalBlockSize if feasible.
// Unless stated otherwise, intervals are inclusive-exclusive, i.e. the
// start byte is included but the end byte is not.
type BlockDev interface {
	io.ReaderAt
	io.WriterAt

	// BlockCount returns the number of blocks on the block device or -1 if it
	// is an image with an undefined size.
	BlockCount() int64

	// BlockSize returns the block size of the block device in bytes. This must
	// be a power of two and is commonly (but not always) either 512 or 4096.
	BlockSize() int64

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

	// Sync commits the current contents to stable storage.
	Sync() error
}

// ReaderFromAt is similar to [io.ReaderFrom], except that the write starts at
// offset off instead of using the file offset.
type ReaderFromAt interface {
	ReadFromAt(r io.Reader, off int64) (n int64, err error)
}

// writerOnly wraps an [io.Writer] and hides all methods other than Write
// (such as ReadFrom).
type writerOnly struct {
	io.Writer
}

// genericReadFromAt is a generic implementation which does not use b.ReadFromAt
// to prevent recursive calls.
func genericReadFromAt(b BlockDev, r io.Reader, off int64) (int64, error) {
	w := &writerOnly{Writer: &ReadWriteSeeker{b: b, currPos: off}}
	return io.Copy(w, r)
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

func (s *ReadWriteSeeker) ReadFrom(r io.Reader) (n int64, err error) {
	rfa, rfaOK := s.b.(ReaderFromAt)
	if !rfaOK {
		w := &writerOnly{Writer: s}
		return io.Copy(w, r)
	}
	n, err = rfa.ReadFromAt(r, s.currPos)
	s.currPos += n
	return
}

func (s *ReadWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, errors.New("Seek: invalid whence")
	case io.SeekStart:
	case io.SeekCurrent:
		offset += s.currPos
	case io.SeekEnd:
		offset += s.b.BlockCount() * s.b.BlockSize()
	}
	if offset < 0 {
		return 0, errors.New("Seek: invalid offset")
	}
	s.currPos = offset
	return s.currPos, nil
}

var ErrOutOfBounds = errors.New("write out of bounds")

// NewSection returns a new Section, implementing BlockDev over that subset
// of blocks. The interval is inclusive-exclusive.
func NewSection(b BlockDev, startBlock, endBlock int64) (*Section, error) {
	if startBlock < 0 {
		return nil, fmt.Errorf("invalid range: startBlock (%d) negative", startBlock)
	}
	if startBlock > endBlock {
		return nil, fmt.Errorf("invalid range: startBlock (%d) bigger than endBlock (%d)", startBlock, endBlock)
	}
	if endBlock > b.BlockCount() {
		return nil, fmt.Errorf("endBlock (%d) out of range (%d)", endBlock, b.BlockCount())
	}
	return &Section{
		b:          b,
		startBlock: startBlock,
		endBlock:   endBlock,
	}, nil
}

// Section implements BlockDev on a slice of another BlockDev given a startBlock
// and endBlock.
type Section struct {
	b                    BlockDev
	startBlock, endBlock int64
}

func (s *Section) ReadAt(p []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, errors.New("blockdev.Section.ReadAt: negative offset")
	}
	bOff := off + (s.startBlock * s.b.BlockSize())
	bytesToEnd := (s.endBlock * s.b.BlockSize()) - bOff
	if bytesToEnd < 0 {
		return 0, io.EOF
	}
	if bytesToEnd < int64(len(p)) {
		n, err := s.b.ReadAt(p[:bytesToEnd], bOff)
		if err == nil {
			err = io.EOF
		}
		return n, err
	}
	return s.b.ReadAt(p, bOff)
}

func (s *Section) WriteAt(p []byte, off int64) (n int, err error) {
	bOff := off + (s.startBlock * s.b.BlockSize())
	bytesToEnd := (s.endBlock * s.b.BlockSize()) - bOff
	if off < 0 || bytesToEnd < 0 {
		return 0, ErrOutOfBounds
	}
	if bytesToEnd < int64(len(p)) {
		n, err := s.b.WriteAt(p[:bytesToEnd], bOff)
		if err != nil {
			// If an error happened, prioritize that error
			return n, err
		}
		// Otherwise, return ErrOutOfBounds as even short writes must return an
		// error.
		return n, ErrOutOfBounds
	}
	return s.b.WriteAt(p, bOff)
}

func (s *Section) ReadFromAt(r io.Reader, off int64) (n int64, err error) {
	rfa, rfaOK := s.b.(ReaderFromAt)
	if !rfaOK {
		return genericReadFromAt(s, r, off)
	}
	bOff := off + (s.startBlock * s.b.BlockSize())
	bytesToEnd := (s.endBlock * s.b.BlockSize()) - bOff
	if off < 0 || bytesToEnd < 0 {
		return 0, ErrOutOfBounds
	}
	ur := r
	lr, lrOK := r.(*io.LimitedReader)
	if lrOK {
		if bytesToEnd >= lr.N {
			return rfa.ReadFromAt(r, bOff)
		}
		ur = lr.R
	}
	n, err = rfa.ReadFromAt(io.LimitReader(ur, bytesToEnd), bOff)
	if lrOK {
		lr.N -= n
	}
	if err == nil && n == bytesToEnd {
		// Return an error if we have not reached EOF.
		moreN, moreErr := io.CopyN(io.Discard, r, 1)
		if moreN != 0 {
			err = ErrOutOfBounds
		} else if moreErr != io.EOF {
			err = moreErr
		}
	}
	return
}

func (s *Section) BlockCount() int64 {
	return s.endBlock - s.startBlock
}

func (s *Section) BlockSize() int64 {
	return s.b.BlockSize()
}

func (s *Section) OptimalBlockSize() int64 {
	return s.b.OptimalBlockSize()
}

func (s *Section) Discard(startByte, endByte int64) error {
	if err := validAlignedRange(s, startByte, endByte); err != nil {
		return err
	}
	offset := s.startBlock * s.b.BlockSize()
	return s.b.Discard(offset+startByte, offset+endByte)
}

func (s *Section) Zero(startByte, endByte int64) error {
	if err := validAlignedRange(s, startByte, endByte); err != nil {
		return err
	}
	offset := s.startBlock * s.b.BlockSize()
	return s.b.Zero(offset+startByte, offset+endByte)
}

func (s *Section) Sync() error {
	return s.b.Sync()
}

func validAlignedRange(b BlockDev, startByte, endByte int64) error {
	if startByte < 0 {
		return fmt.Errorf("invalid range: startByte (%d) negative", startByte)
	}
	if startByte > endByte {
		return fmt.Errorf("invalid range: startByte (%d) bigger than endByte (%d)", startByte, endByte)
	}
	devLen := b.BlockCount() * b.BlockSize()
	if endByte > devLen {
		return fmt.Errorf("endByte (%d) out of range (%d)", endByte, devLen)
	}
	if startByte%b.BlockSize() != 0 {
		return fmt.Errorf("startByte (%d) needs to be aligned to block size (%d)", startByte, b.BlockSize())
	}
	if endByte%b.BlockSize() != 0 {
		return fmt.Errorf("endByte (%d) needs to be aligned to block size (%d)", endByte, b.BlockSize())
	}
	return nil
}

// GenericZero implements software-based zeroing. This can be used to implement
// Zero when no acceleration is available or desired.
func GenericZero(b BlockDev, startByte, endByte int64) error {
	if err := validAlignedRange(b, startByte, endByte); err != nil {
		return err
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
