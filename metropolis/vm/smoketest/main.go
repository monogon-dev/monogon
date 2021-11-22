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

// This is a small smoke test which will run in a container on top of Metropolis
// Kubernetes. It exercises Metropolis' KVM device plugin,
package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	testSocket, err := net.Listen("unix", "@metropolis/vm/smoketest")
	if err != nil {
		panic(err)
	}

	testResultChan := make(chan bool)
	go func() {
		conn, err := testSocket.Accept()
		if err != nil {
			panic(err)
		}
		testValue, _ := io.ReadAll(conn)
		if bytes.Equal(testValue, []byte("test123")) {
			testResultChan <- true
		} else {
			testResultChan <- false
		}
	}()

	baseArgs := []string{"-nodefaults", "-no-user-config", "-nographic", "-no-reboot",
		"-accel", "kvm", "-cpu", "host",
		// TODO(lorenz): This explicitly doesn't use our own qboot because it cannot be built in a musl configuration.
		// This will be fixed once we have a proper multi-target toolchain.
		"-bios", "external/qemu/pc-bios/qboot.rom",
		"-M", "microvm,x-option-roms=off,pic=off,pit=off,rtc=off,isa-serial=off",
		"-kernel", "metropolis/test/ktest/linux-testing.elf",
		"-append", "reboot=t console=hvc0 quiet",
		"-initrd", "metropolis/vm/smoketest/initramfs.lz4",
		"-device", "virtio-rng-device,max-bytes=1024,period=1000",
		"-device", "virtio-serial-device,max_ports=16",
		"-chardev", "stdio,id=con0", "-device", "virtconsole,chardev=con0",
		"-chardev", "socket,id=test,path=metropolis/vm/smoketest,abstract=on",
		"-device", "virtserialport,chardev=test",
	}
	qemuCmd := exec.Command("external/qemu/qemu-x86_64-softmmu", baseArgs...)
	qemuCmd.Stdout = os.Stdout
	qemuCmd.Stderr = os.Stderr
	if err := qemuCmd.Run(); err != nil {
		log.Fatalf("running QEMU failed: %v", err)
	}
	testResult := <-testResultChan
	if testResult {
		return
	} else {
		os.Exit(1)
	}
}
