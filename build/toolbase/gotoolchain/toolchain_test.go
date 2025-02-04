// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package gotoolchain

import (
	"os"
	"os/exec"
	"path"
	"testing"
)

func TestGoToolRuns(t *testing.T) {
	cmd := exec.Command(Go, "version")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to run `go version`: %q, %v", string(out), err)
	}
}

func TestGorootContainsRoot(t *testing.T) {
	rootfile := path.Join(Root, "ROOT")
	if _, err := os.Stat(rootfile); err != nil {
		t.Fatalf("ROOT not found in %s", Root)
	}
}
