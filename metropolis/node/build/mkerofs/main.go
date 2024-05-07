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

// mkerofs takes a specification in the form of a prototext file (see fsspec
// next to this) and assembles an EROFS filesystem according to it. The output
// is fully reproducible.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"source.monogon.dev/metropolis/node/build/fsspec"
	"source.monogon.dev/osbase/erofs"
)

func (spec *entrySpec) writeRecursive(w *erofs.Writer, pathname string) {
	switch inode := spec.data.Type.(type) {
	case *fsspec.Inode_Directory:
		// Sort children for reproducibility
		var sortedChildren []string
		for name := range spec.children {
			sortedChildren = append(sortedChildren, name)
		}
		sort.Strings(sortedChildren)

		err := w.Create(pathname, &erofs.Directory{
			Base: erofs.Base{
				Permissions: uint16(inode.Directory.Mode),
				UID:         uint16(inode.Directory.Uid),
				GID:         uint16(inode.Directory.Gid),
			},
			Children: sortedChildren,
		})
		if err != nil {
			log.Fatalf("failed to write directory: %s", err)
		}
		for _, name := range sortedChildren {
			spec.children[name].writeRecursive(w, path.Join(pathname, name))
		}
	case *fsspec.Inode_File:
		iw := w.CreateFile(pathname, &erofs.FileMeta{
			Base: erofs.Base{
				Permissions: uint16(inode.File.Mode),
				UID:         uint16(inode.File.Uid),
				GID:         uint16(inode.File.Gid),
			},
		})

		sourceFile, err := os.Open(inode.File.SourcePath)
		if err != nil {
			log.Fatalf("failed to open source file %s: %s", inode.File.SourcePath, err)
		}

		_, err = io.Copy(iw, sourceFile)
		if err != nil {
			log.Fatalf("failed to copy file into filesystem: %s", err)
		}
		sourceFile.Close()
		if err := iw.Close(); err != nil {
			log.Fatalf("failed to close target file: %s", err)
		}
	case *fsspec.Inode_SymbolicLink:
		err := w.Create(pathname, &erofs.SymbolicLink{
			Base: erofs.Base{
				Permissions: 0777, // Nominal, Linux forces that mode anyways, see symlink(7)
			},
			Target: inode.SymbolicLink.TargetPath,
		})
		if err != nil {
			log.Fatalf("failed to create symbolic link: %s", err)
		}
	case *fsspec.Inode_SpecialFile:
		err := fmt.Errorf("unimplemented special file type %s", inode.SpecialFile.Type)
		base := erofs.Base{
			Permissions: uint16(inode.SpecialFile.Mode),
			UID:         uint16(inode.SpecialFile.Uid),
			GID:         uint16(inode.SpecialFile.Gid),
		}
		switch inode.SpecialFile.Type {
		case fsspec.SpecialFile_FIFO:
			err = w.Create(pathname, &erofs.FIFO{
				Base: base,
			})
		case fsspec.SpecialFile_CHARACTER_DEV:
			err = w.Create(pathname, &erofs.CharacterDevice{
				Base:  base,
				Major: inode.SpecialFile.Major,
				Minor: inode.SpecialFile.Minor,
			})
		case fsspec.SpecialFile_BLOCK_DEV:
			err = w.Create(pathname, &erofs.BlockDevice{
				Base:  base,
				Major: inode.SpecialFile.Major,
				Minor: inode.SpecialFile.Minor,
			})
		}
		if err != nil {
			log.Fatalf("failed to make special file: %v", err)
		}
	}
}

// entrySpec is a recursive structure representing the filesystem tree
type entrySpec struct {
	data     fsspec.Inode
	children map[string]*entrySpec
}

// pathRef gets the entrySpec at the leaf of the given path, inferring
// directories if necessary
func (spec *entrySpec) pathRef(p string) *entrySpec {
	// This block gets a path array starting at the root of the filesystem. The
	// root folder is the zero-length array.
	pathParts := strings.Split(path.Clean("./"+p), "/")
	if pathParts[0] == "." {
		pathParts = pathParts[1:]
	}

	entryRef := spec
	for _, part := range pathParts {
		childRef, ok := entryRef.children[part]
		if !ok {
			childRef = &entrySpec{
				data:     fsspec.Inode{Type: &fsspec.Inode_Directory{Directory: &fsspec.Directory{Mode: 0555}}},
				children: make(map[string]*entrySpec),
			}
			entryRef.children[part] = childRef
		}
		entryRef = childRef
	}
	return entryRef
}

var (
	outPath = flag.String("out", "", "Output file path")
)

func main() {
	flag.Parse()

	spec, err := fsspec.ReadMergeSpecs(flag.Args())
	if err != nil {
		log.Fatalf("failed to load specs: %v", err)
	}

	var fsRoot = &entrySpec{
		data:     fsspec.Inode{Type: &fsspec.Inode_Directory{Directory: &fsspec.Directory{Mode: 0555}}},
		children: make(map[string]*entrySpec),
	}

	for _, dir := range spec.Directory {
		entryRef := fsRoot.pathRef(dir.Path)
		entryRef.data.Type = &fsspec.Inode_Directory{Directory: dir}
	}

	for _, file := range spec.File {
		entryRef := fsRoot.pathRef(file.Path)
		entryRef.data.Type = &fsspec.Inode_File{File: file}
	}

	for _, symlink := range spec.SymbolicLink {
		entryRef := fsRoot.pathRef(symlink.Path)
		entryRef.data.Type = &fsspec.Inode_SymbolicLink{SymbolicLink: symlink}
	}

	for _, specialFile := range spec.SpecialFile {
		entryRef := fsRoot.pathRef(specialFile.Path)
		entryRef.data.Type = &fsspec.Inode_SpecialFile{SpecialFile: specialFile}
	}

	fs, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("failed to open output file: %v", err)
	}
	writer, err := erofs.NewWriter(fs)
	if err != nil {
		log.Fatalf("failed to initialize EROFS writer: %v", err)
	}

	fsRoot.writeRecursive(writer, ".")

	if err := writer.Close(); err != nil {
		panic(err)
	}
	if err := fs.Close(); err != nil {
		panic(err)
	}
}
