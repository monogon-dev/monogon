// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package structfs

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"os"
)

// Blob is a binary large object, a read-only sequence of bytes of a known size.
type Blob interface {
	Open() (io.ReadCloser, error)
	Size() int64
}

// Bytes implements [Blob] for a byte slice.
type Bytes []byte

func (b Bytes) Open() (io.ReadCloser, error) {
	return &bytesReadCloser{*bytes.NewReader(b)}, nil
}

func (b Bytes) Size() int64 {
	return int64(len(b))
}

type bytesReadCloser struct {
	bytes.Reader
}

func (*bytesReadCloser) Close() error {
	return nil
}

var errNotRegular = errors.New("not a regular file")

// OSPathBlob creates a [Blob] for an OS path.
func OSPathBlob(path string) (Blob, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, &fs.PathError{Op: "blob", Path: path, Err: errNotRegular}
	}
	b := &osPathBlob{
		path: path,
		size: info.Size(),
	}
	return b, nil
}

type osPathBlob struct {
	path string
	size int64
}

func (b *osPathBlob) Open() (io.ReadCloser, error) {
	return os.Open(b.path)
}

func (b *osPathBlob) Size() int64 {
	return b.size
}
