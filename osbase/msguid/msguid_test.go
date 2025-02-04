// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package msguid

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestRoundTrip(t *testing.T) {
	cases := []struct {
		name     string
		uuid     string
		expected [16]byte
	}{
		{
			"WikipediaExample1",
			"00112233-4455-6677-8899-AABBCCDDEEFF",
			[16]byte{
				0x33, 0x22, 0x11, 0x00, 0x55, 0x44, 0x77, 0x66,
				0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			origUUID := uuid.MustParse(c.uuid)
			got := From(origUUID)
			diff := cmp.Diff(c.expected, got)
			if diff != "" {
				t.Fatalf("To(%q) returned unexpected result: %v", origUUID, diff)
			}
			back := To(got)
			diff2 := cmp.Diff(origUUID, back)
			if diff2 != "" {
				t.Errorf("From(To(%q)) did not return original value: %v", origUUID, diff2)
			}
		})
	}
}
