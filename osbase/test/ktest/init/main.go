// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// ktestinit is an init designed to run inside a lightweight VM for running
// tests in there.  It performs basic platform initialization like mounting
// kernel filesystems and launches the test executable at /tester, passes the
// exit code back out over the control socket to ktest and then terminates the
// default VM kernel.
package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func mountInit() error {
	for _, el := range []struct {
		dir   string
		fs    string
		flags uintptr
	}{
		{"/sys", "sysfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/sys/kernel/debug", "debugfs", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/proc", "proc", unix.MS_NOEXEC | unix.MS_NOSUID | unix.MS_NODEV},
		{"/dev", "devtmpfs", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/dev/pts", "devpts", unix.MS_NOEXEC | unix.MS_NOSUID},
		{"/tmp", "tmpfs", 0},
	} {
		if err := os.Mkdir(el.dir, 0755); err != nil && !os.IsExist(err) {
			return fmt.Errorf("could not make %s: %w", el.dir, err)
		}
		if err := unix.Mount(el.fs, el.dir, el.fs, el.flags, ""); err != nil {
			return fmt.Errorf("could not mount %s on %s: %w", el.fs, el.dir, err)
		}
	}
	return nil
}

func main() {
	if err := mountInit(); err != nil {
		panic(err)
	}

	// First virtual serial is always stdout, second is control
	ioConn, err := os.OpenFile("/dev/vport1p1", os.O_RDWR, 0)
	if err != nil {
		fmt.Printf("Failed to open communication device: %v\n", err)
		return
	}
	cmd := exec.Command("/tester", "-test.v")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = append(cmd.Env, "IN_KTEST=true")
	if err := cmd.Run(); err != nil {
		var exerr *exec.ExitError
		if errors.As(err, &exerr) {
			if _, err := ioConn.Write([]byte{uint8(exerr.ExitCode())}); err != nil {
				panic(err)
			}
		}
		fmt.Printf("Failed to execute tests (tests didn't run): %v", err)
	} else {
		ioConn.Write([]byte{0})
	}
	ioConn.Close()

	unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART)
}
