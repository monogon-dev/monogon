// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/bringup"
)

func main() {
	bringup.Runnable(func(ctx context.Context) error {
		fmt.Println("_BRINGUP_LAUNCH_SUCCESS_")
		time.Sleep(5 * time.Second)
		unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
		return nil
	}).Run()
}
