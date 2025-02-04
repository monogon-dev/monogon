// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package declarative

import (
	"fmt"
	"os"
	"sync"

	"golang.org/x/sys/unix"
)

// FSRoot is a root of a storage backend that resides on the local filesystem.
type FSRoot struct {
	// The local path at which the declarative directory structure is located
	// (eg. "/").
	root string
}

type FSPlacement struct {
	root      *FSRoot
	path      string
	writeLock sync.Mutex
}

func (f *FSPlacement) FullPath() string {
	return f.path
}

func (f *FSPlacement) RootRef() interface{} {
	return f.root
}

func (f *FSPlacement) Exists() (bool, error) {
	_, err := os.Stat(f.FullPath())
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (f *FSPlacement) Read() ([]byte, error) {
	return os.ReadFile(f.FullPath())
}

// Write performs an atomic file write, via a temporary file.
func (f *FSPlacement) Write(d []byte, mode os.FileMode) error {
	f.writeLock.Lock()
	defer f.writeLock.Unlock()

	// TODO(q3k): ensure that these do not collide with an existing sibling file, or generate this suffix randomly.
	tmp := f.FullPath() + ".__metropolis_tmp"
	defer os.Remove(tmp)

	tf, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE, mode)
	if err != nil {
		return fmt.Errorf("temporary file open failed: %w", err)
	}
	defer tf.Close()
	if _, err := tf.Write(d); err != nil {
		return fmt.Errorf("temporary file write failed: %w", err)
	}

	// Fsync the source file to guarantee that write is durable. Per Theodore Ts'o:
	//
	// > data=ordered only guarantees the avoidance of stale data (e.g., the previous
	// > contents of a data block showing up after a crash, where the previous data
	// > could be someone's love letters, medical records, etc.). Without the fsync(2)
	// > a zero-length file is a valid and possible outcome after the rename.
	if err := tf.Sync(); err != nil {
		return fmt.Errorf("temporary file sync failed: %w", err)
	}
	if err := tf.Close(); err != nil {
		return fmt.Errorf("temporary file close failed: %w", err)
	}

	if err := unix.Rename(tmp, f.FullPath()); err != nil {
		return fmt.Errorf("renaming target file failed: %w", err)
	}

	return nil
}

func (f *FSPlacement) MkdirAll(perm os.FileMode) error {
	return os.MkdirAll(f.FullPath(), perm)
}

// PlaceFS takes a pointer to a Directory or a pointer to a structure embedding
// Directory and places it at a given filesystem root. From this point on the
// given structure pointer has valid Placement interfaces.
func PlaceFS(dd interface{}, root string) error {
	r := &FSRoot{root}
	pathFor := func(parent, this string) string {
		var np string
		switch {
		case parent == "" && this == "":
			np = "/"
		case parent == "/":
			np = "/" + this
		default:
			np = fmt.Sprintf("%s/%s", parent, this)
		}
		return np
	}
	dp := func(parent, this string) DirectoryPlacement {
		np := pathFor(parent, this)
		return &FSPlacement{path: np, root: r}
	}
	fp := func(parent, this string) FilePlacement {
		np := pathFor(parent, this)
		return &FSPlacement{path: np, root: r}
	}
	err := place(dd, r.root, "", dp, fp)
	if err != nil {
		return fmt.Errorf("could not place: %w", err)
	}
	return nil
}
