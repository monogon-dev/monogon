// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package erofs

import (
	"io"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestKernelInterop(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	type testCase struct {
		name     string
		setup    func(w *Writer) error
		validate func(t *testing.T) error
	}

	tests := []testCase{
		{
			name: "SimpleFolder",
			setup: func(w *Writer) error {
				return w.Create(".", &Directory{
					Base:     Base{GID: 123, UID: 124, Permissions: 0753},
					Children: []string{},
				})
			},
			validate: func(t *testing.T) error {
				var stat unix.Stat_t
				if err := unix.Stat("/test", &stat); err != nil {
					t.Errorf("failed to stat output: %v", err)
				}
				require.EqualValues(t, 124, stat.Uid, "wrong Uid")
				require.EqualValues(t, 123, stat.Gid, "wrong Gid")
				require.EqualValues(t, 0753, stat.Mode&^unix.S_IFMT, "wrong mode")
				return nil
			},
		},
		{
			name: "FolderHierarchy",
			setup: func(w *Writer) error {
				if err := w.Create(".", &Directory{
					Base:     Base{GID: 123, UID: 124, Permissions: 0753},
					Children: []string{"subdir"},
				}); err != nil {
					return err
				}
				if err := w.Create("subdir", &Directory{
					Base:     Base{GID: 123, UID: 124, Permissions: 0753},
					Children: []string{},
				}); err != nil {
					return err
				}
				return nil
			},
			validate: func(t *testing.T) error {
				dirInfo, err := os.ReadDir("/test")
				if err != nil {
					t.Fatalf("Failed to read top-level directory: %v", err)
				}
				require.Len(t, dirInfo, 1, "more subdirs than expected")
				require.Equal(t, "subdir", dirInfo[0].Name(), "unexpected subdir")
				require.True(t, dirInfo[0].IsDir(), "subdir not a directory")
				subdirInfo, err := os.ReadDir("/test/subdir")
				assert.NoError(t, err, "cannot read empty subdir")
				require.Len(t, subdirInfo, 0, "unexpected subdirs in empty directory")
				return nil
			},
		},
		{
			name: "SmallFile",
			setup: func(w *Writer) error {
				if err := w.Create(".", &Directory{
					Base:     Base{GID: 123, UID: 123, Permissions: 0755},
					Children: []string{"test.bin"},
				}); err != nil {
					return err
				}
				writer := w.CreateFile("test.bin", &FileMeta{
					Base: Base{GID: 123, UID: 124, Permissions: 0644},
				})
				r := rand.New(rand.NewSource(0)) // Random but deterministic data
				if _, err := io.CopyN(writer, r, 128); err != nil {
					return err
				}
				if err := writer.Close(); err != nil {
					return err
				}
				return nil
			},
			validate: func(t *testing.T) error {
				var stat unix.Stat_t
				err := unix.Stat("/test/test.bin", &stat)
				assert.NoError(t, err, "failed to stat file")
				require.EqualValues(t, 124, stat.Uid, "wrong Uid")
				require.EqualValues(t, 123, stat.Gid, "wrong Gid")
				require.EqualValues(t, 0644, stat.Mode&^unix.S_IFMT, "wrong mode")
				file, err := os.Open("/test/test.bin")
				assert.NoError(t, err, "failed to open test file")
				defer file.Close()
				r := io.LimitReader(rand.New(rand.NewSource(0)), 128) // Random but deterministic data
				expected, _ := io.ReadAll(r)
				actual, err := io.ReadAll(file)
				assert.NoError(t, err, "failed to read test file")
				assert.Equal(t, expected, actual, "content not identical")
				return nil
			},
		},
		{
			name: "Chardev",
			setup: func(w *Writer) error {
				if err := w.Create(".", &Directory{
					Base:     Base{GID: 123, UID: 123, Permissions: 0755},
					Children: []string{"ttyS0"},
				}); err != nil {
					return err
				}
				err := w.Create("ttyS0", &CharacterDevice{
					Base:  Base{GID: 0, UID: 0, Permissions: 0600},
					Major: 4,
					Minor: 64,
				})
				if err != nil {
					return err
				}
				return nil
			},
			validate: func(t *testing.T) error {
				var stat unix.Statx_t
				err := unix.Statx(0, "/test/ttyS0", 0, unix.STATX_ALL, &stat)
				assert.NoError(t, err, "failed to statx file")
				require.EqualValues(t, 0, stat.Uid, "wrong Uid")
				require.EqualValues(t, 0, stat.Gid, "wrong Gid")
				require.EqualValues(t, 0600, stat.Mode&^unix.S_IFMT, "wrong mode")
				require.EqualValues(t, unix.S_IFCHR, stat.Mode&unix.S_IFMT, "wrong file type")
				require.EqualValues(t, 4, stat.Rdev_major, "wrong dev major")
				require.EqualValues(t, 64, stat.Rdev_minor, "wrong dev minor")
				return nil
			},
		},
		{
			name: "LargeFile",
			setup: func(w *Writer) error {
				if err := w.Create(".", &Directory{
					Base:     Base{GID: 123, UID: 123, Permissions: 0755},
					Children: []string{"test.bin"},
				}); err != nil {
					return err
				}
				writer := w.CreateFile("test.bin", &FileMeta{
					Base: Base{GID: 123, UID: 124, Permissions: 0644},
				})
				r := rand.New(rand.NewSource(1)) // Random but deterministic data
				if _, err := io.CopyN(writer, r, 6500); err != nil {
					return err
				}
				if err := writer.Close(); err != nil {
					return err
				}
				return nil
			},
			validate: func(t *testing.T) error {
				var stat unix.Stat_t
				rawContents, err := os.ReadFile("/dev/ram0")
				assert.NoError(t, err, "failed to read test data")
				log.Printf("%x", rawContents)
				err = unix.Stat("/test/test.bin", &stat)
				assert.NoError(t, err, "failed to stat file")
				require.EqualValues(t, 124, stat.Uid, "wrong Uid")
				require.EqualValues(t, 123, stat.Gid, "wrong Gid")
				require.EqualValues(t, 0644, stat.Mode&^unix.S_IFMT, "wrong mode")
				require.EqualValues(t, 6500, stat.Size, "wrong size")
				file, err := os.Open("/test/test.bin")
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
			name: "MultipleMetaBlocks",
			setup: func(w *Writer) error {
				testFileNames := []string{"test1.bin", "test2.bin", "test3.bin"}
				if err := w.Create(".", &Directory{
					Base:     Base{GID: 123, UID: 123, Permissions: 0755},
					Children: testFileNames,
				}); err != nil {
					return err
				}
				for i, fileName := range testFileNames {
					writer := w.CreateFile(fileName, &FileMeta{
						Base: Base{GID: 123, UID: 124, Permissions: 0644},
					})
					r := rand.New(rand.NewSource(int64(i))) // Random but deterministic data
					if _, err := io.CopyN(writer, r, 2053); err != nil {
						return err
					}
					if err := writer.Close(); err != nil {
						return err
					}
				}
				return nil
			},
			validate: func(t *testing.T) error {
				testFileNames := []string{"test1.bin", "test2.bin", "test3.bin"}
				for i, fileName := range testFileNames {
					file, err := os.Open("/test/" + fileName)
					assert.NoError(t, err, "failed to open test file")
					defer file.Close()
					r := io.LimitReader(rand.New(rand.NewSource(int64(i))), 2053) // Random but deterministic data
					expected, _ := io.ReadAll(r)
					actual, err := io.ReadAll(file)
					assert.NoError(t, err, "failed to read test file")
					require.Equal(t, expected, actual, "content not identical")
				}
				return nil
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			file, err := os.OpenFile("/dev/ram0", os.O_WRONLY, 0644)
			if err != nil {
				t.Fatalf("failed to create test image: %v", err)
			}
			defer file.Close()
			w, err := NewWriter(file)
			if err != nil {
				t.Fatalf("failed to initialize EROFS writer: %v", err)
			}
			if err := test.setup(w); err != nil {
				t.Fatalf("setup failed: %v", err)
			}
			if err := w.Close(); err != nil {
				t.Errorf("failed close: %v", err)
			}
			_ = file.Close()
			if err := os.MkdirAll("/test", 0755); err != nil {
				t.Error(err)
			}
			if err := unix.Mount("/dev/ram0", "/test", "erofs", unix.MS_NOEXEC|unix.MS_NODEV, ""); err != nil {
				t.Fatal(err)
			}
			if err := test.validate(t); err != nil {
				t.Errorf("validation failure: %v", err)
			}
			if err := unix.Unmount("/test", 0); err != nil {
				t.Fatalf("failed to unmount: %v", err)
			}
		})

	}
}
