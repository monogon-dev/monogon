// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileargs

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

// DefaultSize is the default size limit for FileArgs
const DefaultSize = 4 * 1024 * 1024

// TempDirectory is the directory where FileArgs will mount the actual files
// to. Defaults to os.TempDir() but can be globally overridden by the
// application before any FileArgs are used.
var TempDirectory = os.TempDir()

type FileArgs struct {
	path      string
	lastError error
}

// New initializes a new set of file-based arguments. Remember to call Close()
// if you're done using it, otherwise this leaks memory and mounts.
func New() (*FileArgs, error) {
	return NewWithSize(DefaultSize)
}

// NewWthSize is the same as new, but with a custom size limit. Please be aware
// that this data cannot be swapped out and using a size limit that's too high
// can deadlock your kernel.
func NewWithSize(size uint64) (*FileArgs, error) {
	randomNameRaw := make([]byte, 128/8)
	if _, err := io.ReadFull(rand.Reader, randomNameRaw); err != nil {
		return nil, err
	}
	tmpPath := filepath.Join(TempDirectory, hex.EncodeToString(randomNameRaw))
	if err := os.MkdirAll(tmpPath, 0700); err != nil {
		return nil, err
	}
	// This uses ramfs instead of tmpfs because we never want to swap this for
	// security reasons
	if err := unix.Mount("none", tmpPath, "ramfs", unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV, fmt.Sprintf("size=%v", size)); err != nil {
		return nil, err
	}
	return &FileArgs{
		path: tmpPath,
	}, nil
}

// ArgPath returns the path of the temporary file for this argument. It names
// the temporary file according to name.
func (f *FileArgs) ArgPath(name string, content []byte) string {
	if f.lastError != nil {
		return ""
	}

	path := filepath.Join(f.path, name)

	if err := ioutil.WriteFile(path, content, 0600); err != nil {
		f.lastError = err
		return ""
	}

	return path
}

// FileOpt returns a full option with the temporary file name already filled
// in. Example:
//
// option := FileOpt("--testopt", "test.txt", []byte("hello"))
// option == "--testopt=/tmp/daf8ed.../test.txt"
func (f *FileArgs) FileOpt(optName, fileName string, content []byte) string {
	return fmt.Sprintf("%v=%v", optName, f.ArgPath(fileName, content))
}

func (f *FileArgs) Error() error {
	return f.lastError
}

func (f *FileArgs) Close() error {
	if err := unix.Unmount(f.path, 0); err != nil {
		return err
	}
	return os.Remove(f.path)
}
