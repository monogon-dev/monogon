package blkio

import (
	"fmt"
	"io"
	"os"
)

type ReaderWithSize struct {
	io.Reader
	size int64
}

// SizedReader is an io.Reader with a known size
type SizedReader interface {
	io.Reader
	Size() int64
}

// NewSizedReader returns a SizedReader given a reader and a size.
// The returned SizedReader is a ReaderWithSize.
func NewSizedReader(r io.Reader, size int64) SizedReader {
	return &ReaderWithSize{r, size}
}

func (r *ReaderWithSize) Size() int64 {
	return r.size
}

// LazyFileReader implements a SizedReader which opens a file on first read
// and closes it again after the reader has reached EOF.
type LazyFileReader struct {
	name string
	size int64
	f    *os.File
	done bool
}

func (r *LazyFileReader) init() error {
	f, err := os.Open(r.name)
	if err != nil {
		return fmt.Errorf("failed to open file for reading: %w", err)
	}
	r.f = f
	return nil
}

func (r *LazyFileReader) Size() int64 {
	return r.size
}

func (r *LazyFileReader) Read(b []byte) (n int, err error) {
	if r.done {
		return 0, io.EOF
	}
	if r.f == nil {
		if err = r.init(); err != nil {
			return
		}
	}
	n, err = r.f.Read(b)
	if err == io.EOF {
		r.done = true
		r.f.Close()
	}
	return
}

func (r *LazyFileReader) Close() {
	r.done = true
	r.f.Close()
}

func NewFileReader(name string) (*LazyFileReader, error) {
	info, err := os.Stat(name)
	if err != nil {
		return nil, fmt.Errorf("failed to stat: %w", err)
	}
	return &LazyFileReader{
		size: info.Size(),
		name: name,
	}, nil
}
