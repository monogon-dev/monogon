// TestOS is a tiny "operating system" which is packaged the exact same way as
// an actual Metropolis node but only outputs a single flag before exiting.
// It's used for decoupling the installer tests from the Metropolis Node code.
package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	fmt.Println("TestOS launched successfully! _TESTOS_LAUNCH_SUCCESS_")
	unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
}
