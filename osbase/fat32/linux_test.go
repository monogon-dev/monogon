// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package fat32

import (
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/structfs"
)

func TestKernelInterop(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	type testCase struct {
		name     string
		setup    func() structfs.Tree
		validate func(t *testing.T) error
	}

	// Random timestamp in UTC, divisible by 10ms
	testTimestamp1 := time.Date(2022, 03, 04, 5, 6, 7, 10_000_000, time.UTC)
	// Random timestamp in UTC, divisible by 2s
	testTimestamp2 := time.Date(2022, 03, 04, 5, 6, 8, 0, time.UTC)
	// Random timestamp in UTC, divisible by 10ms
	testTimestamp3 := time.Date(2052, 03, 02, 5, 6, 7, 10_000_000, time.UTC)
	// Random timestamp in UTC, divisible by 2s
	testTimestamp4 := time.Date(2052, 10, 04, 5, 3, 4, 0, time.UTC)

	testContent1 := "testcontent1"

	tests := []testCase{
		{
			name: "SimpleFolder",
			setup: func() structfs.Tree {
				return structfs.Tree{{
					Name:    "testdir",
					Mode:    fs.ModeDir,
					ModTime: testTimestamp2,
					Sys: &DirEntrySys{
						CreateTime: testTimestamp1,
					},
				}}
			},
			validate: func(t *testing.T) error {
				var stat unix.Statx_t
				if err := unix.Statx(0, "/dut/testdir", 0, unix.STATX_TYPE|unix.STATX_MTIME|unix.STATX_BTIME, &stat); err != nil {
					availableFiles, err := os.ReadDir("/dut")
					var availableFileNames []string
					for _, f := range availableFiles {
						availableFileNames = append(availableFileNames, f.Name())
					}
					if err != nil {
						t.Fatalf("Failed to list filesystem root directory: %v", err)
					}
					t.Fatalf("Failed to stat output: %v (available: %v)", err, strings.Join(availableFileNames, ", "))
				}
				if stat.Mode&unix.S_IFDIR == 0 {
					t.Errorf("testdir is expected to be a directory, but has mode %v", stat.Mode)
				}
				btime := time.Unix(stat.Btime.Sec, int64(stat.Btime.Nsec))
				if !btime.Equal(testTimestamp1) {
					t.Errorf("testdir btime expected %v, got %v", testTimestamp1, btime)
				}
				mtime := time.Unix(stat.Mtime.Sec, int64(stat.Mtime.Nsec))
				if !mtime.Equal(testTimestamp2) {
					t.Errorf("testdir mtime expected %v, got %v", testTimestamp2, mtime)
				}
				return nil
			},
		},
		{
			name: "SimpleFile",
			setup: func() structfs.Tree {
				return structfs.Tree{{
					Name:    "testfile",
					ModTime: testTimestamp4,
					Sys: &DirEntrySys{
						CreateTime: testTimestamp3,
					},
					Content: structfs.Bytes(testContent1),
				}}
			},
			validate: func(t *testing.T) error {
				var stat unix.Statx_t
				if err := unix.Statx(0, "/dut/testfile", 0, unix.STATX_TYPE|unix.STATX_MTIME|unix.STATX_BTIME, &stat); err != nil {
					t.Fatalf("failed to stat output: %v", err)
				}
				if stat.Mode&unix.S_IFREG == 0 {
					t.Errorf("testfile is expected to be a file, but has mode %v", stat.Mode)
				}
				btime := time.Unix(stat.Btime.Sec, int64(stat.Btime.Nsec))
				if !btime.Equal(testTimestamp3) {
					t.Errorf("testfile ctime expected %v, got %v", testTimestamp3, btime)
				}
				mtime := time.Unix(stat.Mtime.Sec, int64(stat.Mtime.Nsec))
				if !mtime.Equal(testTimestamp4) {
					t.Errorf("testfile mtime expected %v, got %v", testTimestamp3, mtime)
				}
				contents, err := os.ReadFile("/dut/testfile")
				if err != nil {
					t.Fatalf("failed to read back test file: %v", err)
				}
				if string(contents) != testContent1 {
					t.Errorf("testfile contains %x, got %x", contents, []byte(testContent1))
				}
				return nil
			},
		},
		{
			name: "FolderHierarchy",
			setup: func() structfs.Tree {
				return structfs.Tree{{
					Name:    "l1",
					Mode:    fs.ModeDir,
					ModTime: testTimestamp2,
					Sys: &DirEntrySys{
						CreateTime: testTimestamp1,
					},
					Children: structfs.Tree{{
						Name:    "l2",
						Mode:    fs.ModeDir,
						ModTime: testTimestamp2,
						Sys: &DirEntrySys{
							CreateTime: testTimestamp1,
						},
					}},
				}}
			},
			validate: func(t *testing.T) error {
				dirInfo, err := os.ReadDir("/dut/l1")
				if err != nil {
					t.Fatalf("Failed to read top-level directory: %v", err)
				}
				require.Len(t, dirInfo, 1, "more subdirs than expected")
				require.Equal(t, "l2", dirInfo[0].Name(), "unexpected subdir")
				require.True(t, dirInfo[0].IsDir(), "l1 not a directory")
				subdirInfo, err := os.ReadDir("/dut/l1/l2")
				assert.NoError(t, err, "cannot read empty subdir")
				require.Len(t, subdirInfo, 0, "unexpected subdirs in empty directory")
				return nil
			},
		},
		{
			name: "LargeFile",
			setup: func() structfs.Tree {
				content := make([]byte, 6500)
				io.ReadFull(rand.New(rand.NewSource(1)), content)
				return structfs.Tree{{
					Name:    "test.bin",
					Content: structfs.Bytes(content),
				}}
			},
			validate: func(t *testing.T) error {
				var stat unix.Stat_t
				err := unix.Stat("/dut/test.bin", &stat)
				assert.NoError(t, err, "failed to stat file")
				require.EqualValues(t, 6500, stat.Size, "wrong size")
				file, err := os.Open("/dut/test.bin")
				assert.NoError(t, err, "failed to open test file")
				defer file.Close()
				r := io.LimitReader(rand.New(rand.NewSource(1)), 6500) // Random but deterministic data
				expected, _ := io.ReadAll(r)
				actual, err := io.ReadAll(file)
				assert.NoError(t, err, "failed to read test file")
				assert.Equal(t, expected, actual, "content not identical")
				return nil
			},
		},
		{
			name: "Unicode",
			setup: func() structfs.Tree {
				return structfs.Tree{{
					Name:    "✨😂", // Really exercise that UTF-16 conversion
					Content: structfs.Bytes("😂"),
				}}
			},
			validate: func(t *testing.T) error {
				file, err := os.Open("/dut/✨😂")
				if err != nil {
					availableFiles, err := os.ReadDir("/dut")
					var availableFileNames []string
					for _, f := range availableFiles {
						availableFileNames = append(availableFileNames, f.Name())
					}
					if err != nil {
						t.Fatalf("Failed to list filesystem root directory: %v", err)
					}
					t.Fatalf("Failed to open unicode file: %v (available files: %v)", err, strings.Join(availableFileNames, ", "))
				}
				defer file.Close()
				expected := []byte("😂")
				actual, err := io.ReadAll(file)
				assert.NoError(t, err, "failed to read test file")
				assert.Equal(t, expected, actual, "content not identical")
				return nil
			},
		},
		{
			name: "MultipleMetaClusters",
			setup: func() structfs.Tree {
				// Only test up to 2048 files as Linux gets VERY slow if going
				// up to the maximum of approximately 32K
				var root structfs.Tree
				for i := 0; i < 2048; i++ {
					root = append(root, &structfs.Node{
						Name:    fmt.Sprintf("verylongtestfilename%d", i),
						Content: structfs.Bytes("random test content"),
					})
				}
				return root
			},
			validate: func(t *testing.T) error {
				files, err := os.ReadDir("/dut")
				if err != nil {
					t.Errorf("failed to list directory: %v", err)
				}
				if len(files) != 2048 {
					t.Errorf("wrong number of files: expected %d, got %d", 2048, len(files))
				}
				return nil
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file, err := os.OpenFile("/dev/ram0", os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				t.Fatalf("failed to create test image: %v", err)
			}
			size, err := unix.IoctlGetInt(int(file.Fd()), unix.BLKGETSIZE64)
			if err != nil {
				t.Fatalf("failed to get ramdisk size: %v", err)
			}
			blockSize, err := unix.IoctlGetInt(int(file.Fd()), unix.BLKBSZGET)
			if err != nil {
				t.Fatalf("failed to get ramdisk block size: %v", err)
			}
			defer file.Close()
			root := test.setup()
			if err := WriteFS(file, root, Options{
				ID:         1234,
				Label:      "KTEST",
				BlockSize:  uint16(blockSize),
				BlockCount: uint32(size / blockSize),
			}); err != nil {
				t.Fatalf("failed to write fileystem: %v", err)
			}
			_ = file.Close()
			if err := os.MkdirAll("/dut", 0755); err != nil {
				t.Fatal(err)
			}
			// TODO(lorenz): Set CONFIG_FAT_DEFAULT_UTF8 for Monogon Kernel
			if err := unix.Mount("/dev/ram0", "/dut", "vfat", unix.MS_NOEXEC|unix.MS_NODEV, "utf8=1"); err != nil {
				t.Fatal(err)
			}
			defer unix.Unmount("/dut", 0)
			if err := test.validate(t); err != nil {
				t.Fatal(err)
			}
		})

	}
}
