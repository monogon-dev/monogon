// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// ktest is a test launcher for running tests inside a custom kernel and passes
// the results back out.
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"time"

	"source.monogon.dev/osbase/test/qemu"
)

var (
	kernelPath = flag.String("kernel-path", "", "Path of the Kernel ELF file")
	initrdPath = flag.String("initrd-path", "", "Path of the initrd image")
	cmdline    = flag.String("cmdline", "", "Additional kernel command line options")
)

func main() {
	flag.Parse()

	hostFeedbackConn, vmFeedbackConn, err := qemu.NewSocketPair()
	if err != nil {
		log.Fatalf("Failed to create socket pair: %v", err)
	}

	exitCodeChan := make(chan uint8, 1)

	go func() {
		defer hostFeedbackConn.Close()

		returnCode := make([]byte, 1)
		if _, err := io.ReadFull(hostFeedbackConn, returnCode); err != nil {
			log.Fatalf("Failed to read socket: %v", err)
		}
		exitCodeChan <- returnCode[0]
	}()

	if err := qemu.RunMicroVM(context.Background(), &qemu.MicroVMOptions{
		Name:                        "ktest",
		KernelPath:                  *kernelPath,
		InitramfsPath:               *initrdPath,
		Cmdline:                     *cmdline,
		SerialPort:                  os.Stdout,
		ExtraChardevs:               []*os.File{vmFeedbackConn},
		DisableHostNetworkInterface: true,
	}); err != nil {
		log.Fatalf("Failed to run ktest VM: %v", err)
	}

	select {
	case exitCode := <-exitCodeChan:
		os.Exit(int(exitCode))
	case <-time.After(1 * time.Second):
		log.Fatal("Failed to get an error code back (test runtime probably crashed)")
	}
}
