package main

import (
	"context"
	"fmt"

	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/bringup"
)

func main() {
	bringup.Runnable(func(ctx context.Context) error {
		fmt.Println("_BRINGUP_LAUNCH_SUCCESS_")
		unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
		return nil
	}).Run()
}
