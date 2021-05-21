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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/golang/protobuf/proto"

	pb "source.monogon.dev/build/fietsje/proto"
)

// The Shelf is a combined cache and dependency lockfile, not unlike go.sum. It's
// implemented as a text proto file on disk, and currently stores a single mapping
// of shelfKeys to shelfValues, which are in order a (importpath, version) tuple
// and the `locked` structure of a dependency. The resulting shelf file should be
// commited to the monogon repository. It can be freely deleted to force recreation
// from scratch, which can be useful as there is no garbage collection implemented
// for it. The 'lockfile' aspect of the Shelf is counter-intuitive to what readers
// might be used to from other dependency management systems. It does not lock a
// third-party dependency to a particular version, but only locks a well defined
// version to its checksum. As such, recreating the shelf from scratch should not
// bump any dependencies, unless some upstream-project retagged a release to a
// different VCS commit, or a fietsje user pinned to 'master' instead of a
// particular commit. The effective changes will always be reflected in the
// resulting starlark repository ruleset, which (also being commited to source
// control) can be used as a canary of a version being effectively bumped.

// shelfKey is the key into the shelf map structure.
type shelfKey struct {
	importpath string
	version    string
}

// shelfValue is the entry of a shelf map structure.
type shelfValue struct {
	l *locked
}

// shelf is an in-memory representation of the shelf loaded from disk.
type shelf struct {
	path string
	data map[shelfKey]shelfValue
}

func shelfLoad(path string) (*shelf, error) {
	var data []byte
	var err error

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Creating new shelf file at %q, this run will be slow.", path)
	} else {
		data, err = ioutil.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("could not read shelf: %v", err)
		}
	}
	var shelfProto pb.Shelf
	err = proto.UnmarshalText(string(data), &shelfProto)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal shelf: %v", err)
	}

	res := &shelf{
		path: path,
		data: make(map[shelfKey]shelfValue),
	}

	for _, e := range shelfProto.Entry {
		k := shelfKey{
			importpath: e.ImportPath,
			version:    e.Version,
		}
		v := shelfValue{
			l: &locked{
				bazelName: e.BazelName,
				sum:       e.Sum,
				semver:    e.Semver,
			},
		}
		res.data[k] = v
	}
	return res, nil
}

// get retrieves a given lock entry from the in-memory shelf.
func (s *shelf) get(importpath, version string) *locked {
	res, ok := s.data[shelfKey{importpath: importpath, version: version}]
	if !ok {
		return nil
	}
	return res.l
}

// put stores a given locked entry in memory. This will not be commited to disk
// until .save() is called.
func (s *shelf) put(importpath, version string, l *locked) {
	s.data[shelfKey{importpath: importpath, version: version}] = shelfValue{l: l}
}

// save commits the shelf to disk (to the same location it was loaded from), fully
// overwriting from in-memory data.
func (s *shelf) save() error {
	// Build proto representation of shelf data.
	var shelfProto pb.Shelf
	for k, v := range s.data {
		shelfProto.Entry = append(shelfProto.Entry, &pb.Shelf_Entry{
			ImportPath: k.importpath,
			Version:    k.version,
			BazelName:  v.l.bazelName,
			Sum:        v.l.sum,
			Semver:     v.l.semver,
		})
	}

	// Sort shelf keys by importpath, then by version.
	sort.Slice(shelfProto.Entry, func(i, j int) bool {
		a := shelfProto.Entry[i]
		b := shelfProto.Entry[j]

		if a.ImportPath < b.ImportPath {
			return true
		}
		if a.ImportPath > b.ImportPath {
			return false
		}
		return a.Version < b.Version
	})

	// Make an in-memory representation of the marshaled shelf.
	buf := bytes.NewBuffer(nil)
	err := proto.MarshalText(buf, &shelfProto)
	if err != nil {
		return fmt.Errorf("could not serialize shelf: %v", err)
	}

	// And write it out.
	err = ioutil.WriteFile(s.path, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("could not write shelf: %v", err)
	}

	return nil
}
