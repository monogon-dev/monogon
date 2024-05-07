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

package fsquota

import (
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

// withinTolerance is a helper for asserting that a value is within a certain
// percentage of the expected value. The tolerance is specified as a float
// between 0 (exact match) and 1 (between 0 and twice the expected value).
func withinTolerance(t *testing.T, expected uint64, actual uint64, tolerance float64, name string) {
	t.Helper()
	delta := uint64(math.Round(float64(expected) * tolerance))
	lowerBound := expected - delta
	upperBound := expected + delta
	if actual < lowerBound {
		t.Errorf("Value %v (%v) is too low, expected between %v and %v", name, actual, lowerBound, upperBound)
	}
	if actual > upperBound {
		t.Errorf("Value %v (%v) is too high, expected between %v and %v", name, actual, lowerBound, upperBound)
	}
}

func TestBasic(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	mkfsCmd := exec.Command("/mkfs.xfs", "-qf", "/dev/ram0")
	if _, err := mkfsCmd.Output(); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir("/test", 0755); err != nil {
		t.Error(err)
	}

	if err := unix.Mount("/dev/ram0", "/test", "xfs", unix.MS_NOEXEC|unix.MS_NODEV, "prjquota"); err != nil {
		t.Fatal(err)
	}
	defer unix.Unmount("/test", 0)
	defer os.RemoveAll("/test")
	t.Run("SetQuota", func(t *testing.T) {
		defer func() {
			os.RemoveAll("/test/set")
		}()
		if err := os.Mkdir("/test/set", 0755); err != nil {
			t.Fatal(err)
		}
		if err := SetQuota("/test/set", 1024*1024, 100); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("SetQuotaAndExhaust", func(t *testing.T) {
		defer func() {
			os.RemoveAll("/test/sizequota")
		}()
		if err := os.Mkdir("/test/sizequota", 0755); err != nil {
			t.Fatal(err)
		}
		const bytesQuota = 1024 * 1024 // 1MiB
		if err := SetQuota("/test/sizequota", bytesQuota, 0); err != nil {
			t.Fatal(err)
		}
		testfile, err := os.Create("/test/sizequota/testfile")
		if err != nil {
			t.Fatal(err)
		}
		testdata := make([]byte, 1024)
		var bytesWritten int
		for {
			n, err := testfile.Write(testdata)
			if err != nil {
				var pathErr *os.PathError
				if errors.As(err, &pathErr) && errors.Is(pathErr.Err, syscall.ENOSPC) {
					// Running out of space is the only acceptable error to continue execution
					break
				}
				t.Fatal(err)
			}
			bytesWritten += n
		}
		if bytesWritten > bytesQuota {
			t.Errorf("Wrote %v bytes, quota is only %v bytes", bytesWritten, bytesQuota)
		}
	})
	t.Run("GetQuotaReadbackAndUtilization", func(t *testing.T) {
		defer func() {
			os.RemoveAll("/test/readback")
		}()
		if err := os.Mkdir("/test/readback", 0755); err != nil {
			t.Fatal(err)
		}
		const bytesQuota = 1024 * 1024 // 1MiB
		const inodesQuota = 100
		if err := SetQuota("/test/readback", bytesQuota, inodesQuota); err != nil {
			t.Fatal(err)
		}
		sizeFileData := make([]byte, 512*1024)
		if err := os.WriteFile("/test/readback/512kfile", sizeFileData, 0644); err != nil {
			t.Fatal(err)
		}

		quotaUtil, err := GetQuota("/test/readback")
		if err != nil {
			t.Fatal(err)
		}
		require.Equal(t, uint64(bytesQuota), quotaUtil.Bytes, "bytes quota readback incorrect")
		require.Equal(t, uint64(inodesQuota), quotaUtil.Inodes, "inodes quota readback incorrect")

		// Give 10% tolerance for quota used values to account for metadata
		// overhead and internal structures that are also in there. If it's out
		// by more than that it's an issue anyways.
		withinTolerance(t, uint64(len(sizeFileData)), quotaUtil.BytesUsed, 0.1, "BytesUsed")

		// Write 50 inodes for a total of 51 (with the 512K file)
		for i := 0; i < 50; i++ {
			if err := os.WriteFile(fmt.Sprintf("/test/readback/ifile%v", i), []byte("test"), 0644); err != nil {
				t.Fatal(err)
			}
		}

		quotaUtil, err = GetQuota("/test/readback")
		if err != nil {
			t.Fatal(err)
		}

		withinTolerance(t, 51, quotaUtil.InodesUsed, 0.1, "InodesUsed")
	})
}
