// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package kmod

import (
	"errors"
	"os"
	"testing"
)

func TestManagerIntegration(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}
	mgr, err := NewManagerFromPath("/lib/modules")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("LoadExampleModule", func(t *testing.T) {
		if err := mgr.LoadModule("r8169"); err != nil {
			t.Error(err)
		}
		if _, err := os.Stat("/sys/module/r8169"); err != nil {
			t.Error("module load returned success, but module not in sysfs")
		}
	})
	t.Run("LoadNonexistentModule", func(t *testing.T) {
		err := mgr.LoadModule("definitelynomodule")
		var notFoundErr *ErrNotFound
		if !errors.As(err, &notFoundErr) {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
	})
	t.Run("LoadModuleTwice", func(t *testing.T) {
		if err := mgr.LoadModule("r8169"); err != nil {
			t.Error(err)
		}
	})
	// TODO(lorenz): Should test loading dependencies here, but we currently
	// have none in the kernel config and I'm not about to build another kernel
	// just for this.
	t.Run("LoadDeviceModule", func(t *testing.T) {
		if err := mgr.LoadModulesForDevice("pci:v00008086d00001591sv00001043sd000085F0bc02sc00i00"); err != nil {
			t.Error(err)
		}
		if _, err := os.Stat("/sys/module/ice"); err != nil {
			t.Error("module load returned success, but module not in sysfs")
		}
	})
}
