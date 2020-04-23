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

// ktest is a test launcher for running tests inside a custom kernel and passes the results
// back out.
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	kernelPath = flag.String("kernel-path", "", "Path of the Kernel ELF file")
	initrdPath = flag.String("initrd-path", "", "Path of the initrd image")
	cmdline    = flag.String("cmdline", "", "Additional kernel command line options")
)

func main() {
	flag.Parse()

	// Create a temporary socket for passing data (currently only exit code)
	// TODO: Land https://patchwork.ozlabs.org/project/qemu-devel/patch/1357671226-11334-1-git-send-email-alexander_barabash@mentor.com/
	tmpDir := os.TempDir()
	token := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, token); err != nil {
		log.Fatal(err)
	}

	socketPath := filepath.Join(tmpDir, fmt.Sprintf("qemu-io-%x", token))
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	defer os.Remove(socketPath)

	// Start a QEMU microvm (https://github.com/qemu/qemu/blob/master/docs/microvm.rst) with only
	// a RNG and two character devices (one for console, one for OOB communication) attached.
	cmd := exec.Command("qemu-system-x86_64", "-nodefaults", "-no-user-config", "-nographic", "-no-reboot",
		"-accel", "kvm", "-cpu", "host",
		"-M", "microvm,x-option-roms=off,pic=off,pit=off,rtc=off,isa-serial=off",
		"-kernel", *kernelPath,
		"-append", "reboot=t console=hvc0 quiet "+*cmdline,
		"-initrd", *initrdPath,
		"-device", "virtio-rng-device,max-bytes=1024,period=1000",
		"-device", "virtio-serial-device,max_ports=2",
		"-chardev", "stdio,id=con0", "-device", "virtconsole,chardev=con0",
		"-chardev", "socket,id=io,path="+socketPath, "-device", "virtserialport,chardev=io",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	exitCodeChan := make(chan uint8, 1)

	go func() {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		returnCode := make([]byte, 1)
		if _, err := conn.Read(returnCode); err != nil && err != io.EOF {
			log.Fatalf("Failed to read socket: %v", err)
		}
		exitCodeChan <- returnCode[0]
	}()

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run QEMU: %v", err)
	}
	select {
	case exitCode := <-exitCodeChan:
		os.Exit(int(exitCode))
	default:
		log.Printf("Failed to get an error code back")
		os.Exit(1)
	}
}
