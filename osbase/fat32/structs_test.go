// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package fat32

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestStructureSizes(t *testing.T) {
	cases := []struct {
		StructInstance interface{}
		ExpectedSize   int
	}{
		{bootSector{}, 512},
		{fsinfo{}, 512},
		{dirEntry{}, 32},
		{lfnEntry{}, 32},
	}
	for _, c := range cases {
		t.Run(reflect.TypeOf(c.StructInstance).String(), func(t *testing.T) {
			actualSize := binary.Size(c.StructInstance)
			if actualSize != c.ExpectedSize {
				t.Errorf("Expected %d bytes, got %d", c.ExpectedSize, actualSize)
			}
		})
	}
}
