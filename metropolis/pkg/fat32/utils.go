package fat32

import (
	"fmt"
	"io"
	"time"
)

// Wraps a writer and provides support for writing padding up to a specified
// alignment.
// TODO(lorenz): Implement WriterTo when w implements it to allow for copy
// offload
type blockWriter struct {
	w io.Writer
	n int64
}

func newBlockWriter(w io.Writer) *blockWriter {
	return &blockWriter{w: w}
}

func (b *blockWriter) Write(p []byte) (n int, err error) {
	n, err = b.w.Write(p)
	b.n += int64(n)
	return
}

func (b *blockWriter) FinishBlock(alignment int64, mustZero bool) (err error) {
	requiredBytes := (alignment - (b.n % alignment)) % alignment
	if requiredBytes == 0 {
		return nil
	}
	// Do not actually write out zeroes if not necessary
	if s, ok := b.w.(io.Seeker); ok && !mustZero {
		if _, err := s.Seek(requiredBytes-1, io.SeekCurrent); err != nil {
			return fmt.Errorf("failed to seek to create hole for empty block: %w", err)
		}
		if _, err := b.w.Write([]byte{0x00}); err != nil {
			return fmt.Errorf("failed to write last byte to create hole: %w", err)
		}
		b.n += requiredBytes
		return
	}
	emptyBuf := make([]byte, 1*1024*1024)
	for requiredBytes > 0 {
		curBlockBytes := requiredBytes
		if curBlockBytes > int64(len(emptyBuf)) {
			curBlockBytes = int64(len(emptyBuf))
		}
		_, err = b.Write(emptyBuf[:curBlockBytes])
		if err != nil {
			return
		}
		requiredBytes -= curBlockBytes
	}
	return
}

// timeToMsDosTime converts a time.Time to an MS-DOS date and time.
// The resolution is 2s with fTime and 10ms if fTenMils is also used.
// See: http://msdn.microsoft.com/en-us/library/ms724274(v=VS.85).aspx
func timeToMsDosTime(t time.Time) (fDate uint16, fTime uint16, fTenMils uint8) {
	t = t.In(time.UTC)
	if t.Year() < 1980 {
		t = time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	if t.Year() > 2107 {
		t = time.Date(2107, 12, 31, 23, 59, 59, 0, time.UTC)
	}
	fDate = uint16(t.Day() + int(t.Month())<<5 + (t.Year()-1980)<<9)
	fTime = uint16(t.Second()/2 + t.Minute()<<5 + t.Hour()<<11)
	fTenMils = uint8(t.Nanosecond()/1e7 + (t.Second()%2)*100)
	return
}
