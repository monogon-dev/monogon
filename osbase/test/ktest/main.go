// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	"source.monogon.dev/osbase/test/launch"
)

var (
	kernelPath = flag.String("kernel-path", "", "Path of the Kernel ELF file")
	initrdPath = flag.String("initrd-path", "", "Path of the initrd image")
	cmdline    = flag.String("cmdline", "", "Additional kernel command line options")
)

func main() {
	flag.Parse()

	hostFeedbackConn, vmFeedbackConn, err := launch.NewSocketPair()
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

	if err := launch.RunMicroVM(context.Background(), &launch.MicroVMOptions{
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
