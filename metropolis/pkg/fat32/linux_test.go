package fat32

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/mod/semver"
	"golang.org/x/sys/unix"
)

func TestKernelInterop(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	// ONCHANGE(//third_party/linux): Drop this once we move to a Kernel version
	// newer than 5.19 which will have FAT btime support.
	kernelVersion, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		t.Fatalf("unable to determine kernel version: %v", err)
	}
	haveBtime := semver.Compare("v"+string(kernelVersion), "v5.19.0") >= 0

	type testCase struct {
		name     string
		setup    func(root *Inode) error
		validate func(t *testing.T) error
	}

	// Random timestamp in UTC, divisible by 10ms
	testTimestamp1 := time.Date(2022, 03, 04, 5, 6, 7, 10, time.UTC)
	// Random timestamp in UTC, divisible by 2s
	testTimestamp2 := time.Date(2022, 03, 04, 5, 6, 8, 0, time.UTC)
	// Random timestamp in UTC, divisible by 10ms
	testTimestamp3 := time.Date(2052, 03, 02, 5, 6, 7, 10, time.UTC)
	// Random timestamp in UTC, divisible by 2s
	testTimestamp4 := time.Date(2052, 10, 04, 5, 3, 4, 0, time.UTC)

	testContent1 := "testcontent1"

	tests := []testCase{
		{
			name: "SimpleFolder",
			setup: func(root *Inode) error {
				root.Children = []*Inode{{
					Name:       "testdir",
					Attrs:      AttrDirectory,
					CreateTime: testTimestamp1,
					ModTime:    testTimestamp2,
				}}
				return nil
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
				if !btime.Equal(testTimestamp1) && haveBtime {
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
			setup: func(root *Inode) error {
				root.Children = []*Inode{{
					Name:       "testfile",
					CreateTime: testTimestamp3,
					ModTime:    testTimestamp4,
					Content:    strings.NewReader(testContent1),
				}}
				return nil
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
				if !btime.Equal(testTimestamp3) && haveBtime {
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
			setup: func(i *Inode) error {
				i.Children = []*Inode{{
					Name:       "l1",
					Attrs:      AttrDirectory,
					CreateTime: testTimestamp1,
					ModTime:    testTimestamp2,
					Children: []*Inode{{
						Name:       "l2",
						Attrs:      AttrDirectory,
						CreateTime: testTimestamp1,
						ModTime:    testTimestamp2,
					}},
				}}
				return nil
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
			setup: func(i *Inode) error {
				content := make([]byte, 6500)
				io.ReadFull(rand.New(rand.NewSource(1)), content)
				i.Children = []*Inode{{
					Name:    "test.bin",
					Content: bytes.NewReader(content),
				}}
				return nil
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
			setup: func(i *Inode) error {
				i.Children = []*Inode{{
					Name:    "✨😂", // Really exercise that UTF-16 conversion
					Content: strings.NewReader("😂"),
				}}
				return nil
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
				contents, err := io.ReadAll(file)
				if err != nil {
					t.Errorf("Wrong content: expected %x, got %x", []byte("😂"), contents)
				}
				return nil
			},
		},
		{
			name: "MultipleMetaClusters",
			setup: func(root *Inode) error {
				// Only test up to 2048 files as Linux gets VERY slow if going
				// up to the maximum of approximately 32K
				for i := 0; i < 2048; i++ {
					root.Children = append(root.Children, &Inode{
						Name:    fmt.Sprintf("verylongtestfilename%d", i),
						Content: strings.NewReader("random test content"),
					})
				}
				return nil
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
			rootInode := Inode{
				Attrs: AttrDirectory,
			}
			if err := test.setup(&rootInode); err != nil {
				t.Fatalf("setup failed: %v", err)
			}
			if err := WriteFS(file, rootInode, Options{
				ID:         1234,
				Label:      "KTEST",
				BlockSize:  uint16(blockSize),
				BlockCount: uint32(size / blockSize),
			}); err != nil {
				t.Fatalf("failed to write fileystem: %v", err)
			}
			_ = file.Close()
			if err := os.MkdirAll("/dut", 0755); err != nil {
				t.Error(err)
			}
			// TODO(lorenz): Set CONFIG_FAT_DEFAULT_UTF8 for Monogon Kernel
			if err := unix.Mount("/dev/ram0", "/dut", "vfat", unix.MS_NOEXEC|unix.MS_NODEV, "utf8=1"); err != nil {
				t.Fatal(err)
			}
			defer unix.Unmount("/dut", 0)
			test.validate(t)
		})

	}
}
