// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/sys/unix"

	"source.monogon.dev/osbase/bringup"
)

// Environment variable which tells the takeover binary to run the correct stage
const launchModeEnv = "TAKEOVER_LAUNCH_MODE"

const (
	launchModeTakeover = ""
	launchModeDetached = "DETACHED"
	launchModeInit     = "INIT"
)

func main() {
	switch m := os.Getenv(launchModeEnv); m {
	case launchModeTakeover:
		launchTakeover()
	case launchModeDetached:
		launchDetached()
	case launchModeInit:
		launchInit()
	default:
		panic("unknown launch mode: " + m)
	}
}

func launchTakeover() {
	disk := flag.String("disk", "", "disk to install to without /dev/")
	flag.Parse()
	if disk == nil || *disk == "" {
		log.Fatal("missing target disk")
	}

	nodeParamsRaw, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// try removing /dev/ just to be safe
	diskName := strings.ReplaceAll(*disk, "/dev/", "")
	warns, err := setupTakeover(nodeParamsRaw, diskName)
	if err != nil {
		log.Fatal(err)
	}

	if len(warns) != 0 {
		for _, s := range warns {
			os.Stdout.WriteString(s)
		}
	}

	// Close stdout, we're done responding
	os.Stdout.Close()

	// Start second stage which waits for 5 seconds while performing
	// final cleanup.
	detachedCmd := exec.Command("/proc/self/exe")
	detachedCmd.Env = []string{launchModeEnv + "=" + launchModeDetached}
	if err := detachedCmd.Start(); err != nil {
		log.Fatalf("failed to launch final stage: %v", err)
	}
	// Release the second stage so that the first stage can cleanly terminate.
	if err := detachedCmd.Process.Release(); err != nil {
		log.Fatalf("error releasing final stage process: %v", err)
	}
}

// launchDetached executes the second stage
func launchDetached() {
	// Wait 5 seconds for data to be sent, connections to be closed and
	// syncs to be executed
	time.Sleep(5 * time.Second)
	// Perform kexec, this will not return unless it fails
	err := unix.Reboot(unix.LINUX_REBOOT_CMD_KEXEC)
	msg := "takeover: reboot succeeded, but we're still runing??"
	if err != nil {
		msg = err.Error()
	}
	// We have no standard output/error anymore, if this fails it's
	// just borked. Attempt to dump the error into kmesg for manual
	// debugging.
	kmsg, err := os.OpenFile("/dev/kmsg", os.O_WRONLY, 0)
	if err != nil {
		os.Exit(2)
	}
	kmsg.WriteString(msg)
	kmsg.Close()
	os.Exit(1)
}

func launchInit() {
	bringup.Runnable(takeoverRunnable).Run()
}
