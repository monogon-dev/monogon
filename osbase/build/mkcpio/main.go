package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/cavaliergopher/cpio"
	"github.com/klauspost/compress/zstd"
	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/build/fsspec"
)

var (
	outPath = flag.String("out", "", "Output file path")
)

type placeEnum int

const (
	// placeNone implies that currently nothing is placed at that path.
	// Can be overridden by everything.
	placeNone placeEnum = 0
	// placeDirImplicit means that there is currently a implied directory
	// at the given path. It can be overridden by (and only by) an explicit
	// directory.
	placeDirImplicit placeEnum = 1
	// placeDirExplicit means that there is an explicit (i.e. specified by
	// the FSSpec) directory at the given path. Nothing else can override
	// this.
	placeDirExplicit placeEnum = 2
	// placeNonDir means that there is a file-type resource (i.e a file, symlink
	// or special_file) at the given path. Nothing else can override this.
	placeNonDir placeEnum = 3
)

// place represents the state a given canonical path is in during metadata
// construction. Its zero value is { State: placeNone, Inode: nil }.
type place struct {
	State placeEnum
	// Inode contains one of the types inside an FSSpec (e.g. *fsspec.File)
	Inode interface{}
}

// Usage: -out <out-path.cpio.zst> fsspec-path...
func main() {
	flag.Parse()
	outFile, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("Failed to open CPIO output file: %v", err)
	}
	defer outFile.Close()
	compressedOut, err := zstd.NewWriter(outFile)
	if err != nil {
		log.Fatalf("While initializing zstd writer: %v", err)
	}
	defer compressedOut.Close()
	cpioWriter := cpio.NewWriter(compressedOut)
	defer cpioWriter.Close()

	spec, err := fsspec.ReadMergeSpecs(flag.Args())
	if err != nil {
		log.Fatalf("failed to load specs: %v", err)
	}

	// Map of paths to metadata for validation & implicit directory injection
	places := make(map[string]place)

	// The idea behind this machinery is that we try to place all files and
	// directories into a map while creating the required parent directories
	// on-the-fly as implicit directories. Overriding an implicit directory
	// with an explicit one is allowed thus the actual order in which this
	// structure is created does not matter. All non-directories cannot be
	// overridden anyways so their insertion order does not matter.
	// This also has the job of validating the FSSpec structure, ensuring that
	// there are no duplicate paths and that there is nothing placed below a
	// non-directory.
	var placeInode func(p string, isDir bool, inode interface{})
	placeInode = func(p string, isDir bool, inode interface{}) {
		cleanPath := path.Clean(p)
		if !isDir {
			if places[cleanPath].State != placeNone {
				log.Fatalf("Invalid FSSpec: Duplicate Inode at %q", cleanPath)
			}
			places[cleanPath] = place{
				State: placeNonDir,
				Inode: inode,
			}
		} else {
			switch places[cleanPath].State {
			case placeNone:
				if inode != nil {
					places[cleanPath] = place{
						State: placeDirExplicit,
						Inode: inode,
					}
				} else {
					places[cleanPath] = place{
						State: placeDirImplicit,
						Inode: &fsspec.Directory{Path: cleanPath, Mode: 0555},
					}
				}
			case placeDirImplicit:
				if inode != nil {
					places[cleanPath] = place{
						State: placeDirExplicit,
						Inode: inode,
					}
				}
			case placeDirExplicit:
				if inode != nil {
					log.Fatalf("Invalid FSSpec: Conflicting explicit directories at %v", cleanPath)
				}
			case placeNonDir:
				log.Fatalf("Invalid FSSpec: Trying to place inode below non-directory at #{cleanPath}")
			default:
				panic("unhandled placeEnum value")
			}
		}
		parentPath, _ := path.Split(p)
		parentPath = path.Clean(parentPath)
		if parentPath == "/" || parentPath == p {
			return
		}
		placeInode(parentPath, true, nil)
	}
	for _, d := range spec.Directory {
		placeInode(d.Path, true, d)
	}
	for _, f := range spec.File {
		placeInode(f.Path, false, f)
	}
	for _, s := range spec.SymbolicLink {
		placeInode(s.Path, false, s)
	}
	for _, s := range spec.SpecialFile {
		placeInode(s.Path, false, s)
	}

	var writeOrder []string
	for path := range places {
		writeOrder = append(writeOrder, path)
	}
	// Sorting a list of normalized paths representing a tree gives us Depth-
	// first search (DFS) order which is the correct order for writing archives.
	// This also makes the output reproducible.
	sort.Strings(writeOrder)

	for _, path := range writeOrder {
		place := places[path]
		switch i := place.Inode.(type) {
		case *fsspec.File:
			inF, err := os.Open(i.SourcePath)
			if err != nil {
				log.Fatalf("Failed to open source path for file %q: %v", i.Path, err)
			}
			inFStat, err := inF.Stat()
			if err != nil {
				log.Fatalf("Failed to stat source path for file %q: %v", i.Path, err)
			}
			if err := cpioWriter.WriteHeader(&cpio.Header{
				Mode: cpio.FileMode(i.Mode),
				Name: strings.TrimPrefix(i.Path, "/"),
				Size: inFStat.Size(),
			}); err != nil {
				log.Fatalf("Failed to write cpio header for file %q: %v", i.Path, err)
			}
			if n, err := io.Copy(cpioWriter, inF); err != nil || n != inFStat.Size() {
				log.Fatalf("Failed to copy file %q into cpio: %v", i.SourcePath, err)
			}
			inF.Close()
		case *fsspec.Directory:
			if err := cpioWriter.WriteHeader(&cpio.Header{
				Mode: cpio.FileMode(i.Mode) | cpio.TypeDir,
				Name: strings.TrimPrefix(i.Path, "/"),
			}); err != nil {
				log.Fatalf("Failed to write cpio header for directory %q: %v", i.Path, err)
			}
		case *fsspec.SymbolicLink:
			if err := cpioWriter.WriteHeader(&cpio.Header{
				// Symlinks are 0777 by definition (from man 7 symlink on Linux)
				Mode: 0777 | cpio.TypeSymlink,
				Name: strings.TrimPrefix(i.Path, "/"),
				Size: int64(len(i.TargetPath)),
			}); err != nil {
				log.Fatalf("Failed to write cpio header for symlink %q: %v", i.Path, err)
			}
			if _, err := cpioWriter.Write([]byte(i.TargetPath)); err != nil {
				log.Fatalf("Failed to write cpio symlink %q: %v", i.Path, err)
			}
		case *fsspec.SpecialFile:
			mode := cpio.FileMode(i.Mode)
			switch i.Type {
			case fsspec.SpecialFile_TYPE_CHARACTER_DEV:
				mode |= cpio.TypeChar
			case fsspec.SpecialFile_TYPE_BLOCK_DEV:
				mode |= cpio.TypeBlock
			case fsspec.SpecialFile_TYPE_FIFO:
				mode |= cpio.TypeFifo
			}

			if err := cpioWriter.WriteHeader(&cpio.Header{
				Mode:     mode,
				Name:     strings.TrimPrefix(i.Path, "/"),
				DeviceID: int(unix.Mkdev(i.Major, i.Minor)),
			}); err != nil {
				log.Fatalf("Failed to write CPIO header for special file %q: %v", i.Path, err)
			}
		default:
			panic("inode type not handled")
		}
	}
}
