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

package declarative

import (
	"fmt"
	"os"
)

// A declarative Directory/File tree is an abstract definition until it's 'placed' on a backing file system.
// By convention, all abstract definitions of hierarchies are stored as copiable structs, and only turned to pointers
// when placed (ie., implementations like PlaceFS takes a *Directory, but Root as a declarative definition is defined as
// non-pointer).

// Placement is an interface available on Placed Files and Directories. All *Placement interfaces on Files/Directories
// are only available on placed trees - eg., after a PlaceFS call. This is unfortunately not typesafe, callers need to
// either be sure about placement, or check the interface for null.
type Placement interface {
	FullPath() string
	RootRef() interface{}
}

// FilePlacement is an interface available on Placed Files. It is implemented by different placement backends, and
// set on all files during placement by a given backend.
type FilePlacement interface {
	Placement
	Exists() (bool, error)
	Read() ([]byte, error)
	Write([]byte, os.FileMode) error
}

// DirectoryPlacement is an interface available on Placed Directories. It is implemented by different placement
// backends, and set on all directories during placement by a given backend.
type DirectoryPlacement interface {
	Placement
	// MkdirAll creates this directory and all its parents on backing stores that have a physical directory
	// structure.
	MkdirAll(file os.FileMode) error
}

// DirectoryPlacer is a placement backend-defined function that, given the path returned by the parent of a directory,
// and the path to a directory, returns a DirectoryPlacement implementation for this directory. The new placement's
// path (via .FullPath()) will be used for placement of directories/files within the new directory.
type DirectoryPlacer func(parent, this string) DirectoryPlacement

// FilePlacer is analogous to DirectoryPlacer, but for files.
type FilePlacer func(parent, this string) FilePlacement

// place recursively places a pointer to a Directory or pointer to a structure embedding Directory into a given backend,
// by calling DirectoryPlacer and FilePlacer where appropriate. This is done recursively across a declarative tree until
// all children are placed.
func place(d interface{}, parent, this string, dpl DirectoryPlacer, fpl FilePlacer) error {
	_, dir, err := unpackDirectory(d)
	if err != nil {
		return err
	}

	if dir.DirectoryPlacement != nil {
		return fmt.Errorf("already placed")
	}
	dir.DirectoryPlacement = dpl(parent, this)

	dirlist, err := subdirs(d)
	if err != nil {
		return fmt.Errorf("could not list subdirectories: %w", err)
	}
	for _, nd := range dirlist {
		err := place(nd.directory, dir.FullPath(), nd.name, dpl, fpl)
		if err != nil {
			return fmt.Errorf("%v: %w", nd.name, err)
		}
	}
	filelist, err := files(d)
	if err != nil {
		return fmt.Errorf("could not list files: %w", err)
	}
	for _, nf := range filelist {
		nf.file.FilePlacement = fpl(dir.FullPath(), nf.name)
	}
	return nil
}
