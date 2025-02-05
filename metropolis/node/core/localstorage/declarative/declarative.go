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
	"reflect"
	"strings"
)

// Directory represents the intent of existence of a directory in a
// hierarchical filesystem (simplified to a tree).  This structure can be
// embedded and still be interpreted as a Directory for purposes of use within
// this library. Any inner fields of such an embedding structure that are in
// turn (embedded) Directories or files will be treated as children in the
// intent expressed by this Directory. All contained directory fields must have
// a `dir:"name"` struct tag that names them, and all contained file fields
// must have a `file:"name"` struct tag.
//
// Creation and management of the directory at runtime is left to the
// implementing code. However, the DirectoryPlacement implementation (set as
// the directory is placed onto a backing store) facilitates this management
// (by exposing methods that mutate the backing store).
type Directory struct {
	DirectoryPlacement
}

// File represents the intent of existence of a file. files are usually child
// structures in types that embed Directory.  File can also be embedded in
// another structure, and this embedding type will still be interpreted as a
// File for purposes of use within this library.
//
// As with Directory, the runtime management of a File in a backing store is
// left to the implementing code, and the embedded FilePlacement interface
// facilitates access to the backing store.
type File struct {
	FilePlacement
}

// unpackDirectory takes a pointer to Directory or a pointer to a structure
// embedding Directory, and returns a reflection Value that refers to the
// passed structure itself (not its pointer) and a plain Go pointer to the
// (embedded) Directory.
func unpackDirectory(d interface{}) (*reflect.Value, *Directory, error) {
	td := reflect.TypeOf(d)
	if td.Kind() != reflect.Ptr {
		return nil, nil, fmt.Errorf("wanted a pointer, got %v", td.Kind())
	}

	var dir *Directory
	id := reflect.ValueOf(d).Elem()
	tid := id.Type()
	switch {
	case tid.Name() == reflect.TypeOf(Directory{}).Name():
		dir = id.Addr().Interface().(*Directory)
	case id.FieldByName("Directory").IsValid():
		dir = id.FieldByName("Directory").Addr().Interface().(*Directory)
	default:
		return nil, nil, fmt.Errorf("not a Directory or embedding Directory (%v)", id.Type().String())
	}
	return &id, dir, nil
}

// unpackFile takes a pointer to a File or a pointer to a structure embedding
// File, and returns a reflection Value that refers to the passed structure
// itself (not its pointer) and a plain Go pointer to the (embedded) File.
func unpackFile(f interface{}) (*reflect.Value, *File, error) {
	tf := reflect.TypeOf(f)
	if tf.Kind() != reflect.Ptr {
		return nil, nil, fmt.Errorf("wanted a pointer, got %v", tf.Kind())
	}

	var fil *File
	id := reflect.ValueOf(f).Elem()
	tid := id.Type()
	switch {
	case tid.Name() == reflect.TypeOf(File{}).Name():
		fil = id.Addr().Interface().(*File)
	case id.FieldByName("File").IsValid():
		fil = id.FieldByName("File").Addr().Interface().(*File)
	default:
		return nil, nil, fmt.Errorf("not a File or embedding File (%v)", tid.String())
	}
	return &id, fil, nil

}

// subdirs takes a pointer to a Directory or pointer to a structure embedding
// Directory, and returns a pair of pointers to Directory-like structures
// contained within that directory with corresponding names (based on struct
// tags).
func subdirs(d interface{}) ([]namedDirectory, error) {
	s, _, err := unpackDirectory(d)
	if err != nil {
		return nil, fmt.Errorf("argument could not be parsed as *Directory: %w", err)
	}

	var res []namedDirectory
	for i := 0; i < s.NumField(); i++ {
		tf := s.Type().Field(i)
		dirTag := tf.Tag.Get("dir")
		if dirTag == "" {
			continue
		}
		sf := s.Field(i)
		res = append(res, namedDirectory{dirTag, sf.Addr().Interface()})
	}
	return res, nil
}

type namedDirectory struct {
	name      string
	directory interface{}
}

// files takes a pointer to a File or pointer to a structure embedding File,
// and returns a pair of pointers to Directory-like structures contained within
// that directory with corresponding names (based on struct tags).
func files(d interface{}) ([]namedFile, error) {
	s, _, err := unpackDirectory(d)
	if err != nil {
		return nil, fmt.Errorf("argument could not be parsed as *Directory: %w", err)
	}

	var res []namedFile
	for i := 0; i < s.NumField(); i++ {
		tf := s.Type().Field(i)
		fileTag := tf.Tag.Get("file")
		if fileTag == "" {
			continue
		}
		_, f, err := unpackFile(s.Field(i).Addr().Interface())
		if err != nil {
			return nil, fmt.Errorf("file %q could not be parsed as *File: %w", tf.Name, err)
		}
		res = append(res, namedFile{fileTag, f})
	}
	return res, nil
}

type namedFile struct {
	name string
	file *File
}

// Validate checks that a given pointer to a Directory or pointer to a
// structure containing Directory does not contain any programmer errors in its
// definition:
//   - all subdirectories/files must be named
//   - all subdirectory/file names within a directory must be unique
//   - all subdirectory/file names within a directory must not contain the '/'
//     character (as it is a common path delimiter)
func Validate(d interface{}) error {
	names := make(map[string]bool)

	subs, err := subdirs(d)
	if err != nil {
		return fmt.Errorf("could not get subdirectories: %w", err)
	}

	for _, nd := range subs {
		if nd.name == "" {
			return fmt.Errorf("subdirectory with empty name")
		}
		if strings.Contains(nd.name, "/") {
			return fmt.Errorf("subdirectory with invalid path: %q", nd.name)
		}
		if names[nd.name] {
			return fmt.Errorf("subdirectory with duplicate name: %q", nd.name)
		}
		names[nd.name] = true

		err := Validate(nd.directory)
		if err != nil {
			return fmt.Errorf("%s: %w", nd.name, err)
		}
	}

	filelist, err := files(d)
	if err != nil {
		return fmt.Errorf("could not get files: %w", err)
	}

	for _, nf := range filelist {
		if nf.name == "" {
			return fmt.Errorf("file with empty name")
		}
		if strings.Contains(nf.name, "/") {
			return fmt.Errorf("file with invalid path: %q", nf.name)
		}
		if names[nf.name] {
			return fmt.Errorf("file with duplicate name: %q", nf.name)
		}
		names[nf.name] = true
	}
	return nil
}
