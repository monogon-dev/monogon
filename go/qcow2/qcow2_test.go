// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package qcow2

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestGenerate exercises the Generate function for a variety of image sizes.
func TestGenerate(t *testing.T) {
	// Test all orders of magnitude from 1KiB to 1PiB.
	for i := 20; i < 50; i++ {
		t.Run(fmt.Sprintf("%d", 1<<i), func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "test.qcow2")

			f, err := os.Create(path)
			if err != nil {
				t.Fatalf("Could not create test image file: %v", err)
			}
			if err := Generate(f, GenerateWithFileSize(1<<i)); err != nil {
				t.Fatalf("Generate(%d bytes): %v", 1<<i, err)
			}
			if err := f.Close(); err != nil {
				t.Fatalf("Close: %v", err)
			}

			cmd := exec.Command("qemu-img", "check", path)
			if err := cmd.Run(); err != nil {
				t.Fatalf("qemu-img check failed: %v", err)
			}
		})
	}
}
