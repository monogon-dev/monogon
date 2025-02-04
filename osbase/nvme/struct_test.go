// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package nvme

import (
	"encoding/binary"
	"testing"
)

// TestStruct tests if the struct passed to Linux's ioctl has the ABI-specified
// size.
func TestStruct(t *testing.T) {
	passthruCmdSize := binary.Size(passthruCmd{})
	if passthruCmdSize != 72 {
		t.Errorf("passthroughCmd is %d bytes, expected 72", passthruCmdSize)
	}
}
