// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package nvme

import (
	"fmt"
	"runtime"
)

func (d *Device) RawCommand(cmd *Command) error {
	return fmt.Errorf("NVMe command interface unimplemented for %v", runtime.GOOS)
}
