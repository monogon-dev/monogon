// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package fat32

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"

	"source.monogon.dev/osbase/structfs"
)

var (
	// These are filled by bazel at linking time with the canonical path of
	// their corresponding file. Inside the init function we resolve it
	// with the rules_go runfiles package to the real path.
	xFsckPath string
)

func init() {
	if os.Getenv("IN_KTEST") == "true" {
		return
	}

	var err error
	for _, path := range []*string{
		&xFsckPath,
	} {
		*path, err = runfiles.Rlocation(*path)
		if err != nil {
			panic(err)
		}
	}
}

func testWithFsck(t *testing.T, root structfs.Tree, opts Options) {
	t.Helper()
	testFile, err := os.CreateTemp("", "fat32-fsck-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile.Name())
	sizeBlocks, err := SizeFS(root, opts)
	if err != nil {
		t.Fatalf("failed to calculate size: %v", err)
	}
	sizeBytes := sizeBlocks * int64(opts.BlockSize)

	// Fill the file with random bytes before writing the FS.
	_, err = io.CopyN(testFile, rand.New(rand.NewSource(sizeBytes)), sizeBytes)
	if err != nil {
		t.Fatalf("write failed: %v", err)
	}
	_, err = testFile.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatalf("seek failed: %v", err)
	}

	if err := WriteFS(testFile, root, opts); err != nil {
		t.Fatalf("failed to write test FS: %v", err)
	}
	// Run fsck non-interactively (-n), disallow spaces in short file names (-S)
	// as well as perform deep verification (-V)
	// If the file system is OK (i.e. fsck does not want to fix it) it returns
	// 0, otherwise 1.
	fsckCmd := exec.Command(xFsckPath, "-n", "-S", "-V", testFile.Name())
	result, err := fsckCmd.CombinedOutput()
	if err != nil {
		t.Errorf("fsck failed: %v", string(result))
	}
}

func TestBasicFsck(t *testing.T) {
	if os.Getenv("IN_KTEST") == "true" {
		t.Skip("In ktest")
	}
	var largeString strings.Builder
	for i := 0; i < 16384; i++ {
		fmt.Fprintf(&largeString, "part%d", i)
	}
	// Test both common block sizes (512 and 4096 bytes) as well as the largest
	// supported one (32K)
	for _, blockSize := range []uint16{512, 4096, 32768} {
		for _, fixed := range []string{"", "Fixed"} {
			t.Run(fmt.Sprintf("BlockSize%d%v", blockSize, fixed), func(t *testing.T) {
				var root structfs.Tree
				files := []struct {
					name    string
					path    string
					content string
				}{
					{"FileInRoot", "test1.txt", "test1 content"},
					{"LongFileInRoot", "verylongtest1.txt", "test1 content long"},
					{"LongPath", "test1/test2/test3/test4/longdirname.ext/hello", "long path test content"},
					{"LargeFile", "test1/largefile.txt", largeString.String()},
				}
				for _, c := range files {
					err := root.PlaceFile(c.path, structfs.Bytes(c.content))
					if err != nil {
						t.Errorf("failed to place file: %v", err)
					}
				}
				opts := Options{ID: 1234, Label: "TEST", BlockSize: blockSize}
				if fixed == "Fixed" {
					// Use a block count that is slightly higher than the minimum
					opts.BlockCount = 67000
				}
				testWithFsck(t, root, opts)
			})
		}
	}
}

func TestLotsOfFilesFsck(t *testing.T) {
	if os.Getenv("IN_KTEST") == "true" {
		t.Skip("In ktest")
	}
	var root structfs.Tree
	for i := 0; i < (32*1024)-2; i++ {
		root = append(root, &structfs.Node{
			Name:    fmt.Sprintf("test%d", i),
			Content: structfs.Bytes("random test content"),
			// Add a random ModTime
			ModTime: time.Date(2022, 03, 04, 5, 6, 7, 8, time.UTC),
			Sys: &DirEntrySys{
				// Add some random attributes
				Attrs: AttrHidden | AttrSystem,
			},
		})
	}
	testWithFsck(t, root, Options{ID: 1234, Label: "TEST"})
}
