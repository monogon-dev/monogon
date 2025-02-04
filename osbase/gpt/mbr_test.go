// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package gpt

import "testing"

func TestToCHS(t *testing.T) {
	cases := []struct {
		name        string
		lba         int64
		expectedCHS [3]byte
	}{
		{ // See UEFI Specification 2.9 Table 5-4 StartingCHS
			name:        "One",
			lba:         1,
			expectedCHS: [3]byte{0x00, 0x02, 0x00},
		},
		{
			name:        "TooBig",
			lba:         (1023 * 255 * 63) + 1,
			expectedCHS: [3]byte{0xff, 0xff, 0xff},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			chs := toCHS(c.lba)
			if chs != c.expectedCHS {
				t.Errorf("expected %x, got %x", c.expectedCHS, chs)
			}
		})
	}
}
