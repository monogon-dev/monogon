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

package main

import (
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	if err := unix.Mount("devtmpfs", "/dev", "devtmpfs", 0, ""); err != nil {
		panic(err)
	}
	testPort, err := os.OpenFile("/dev/vport1p1", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	testPort.WriteString("test123")
	testPort.Close()
	unix.Reboot(unix.LINUX_REBOOT_CMD_POWER_OFF)
}
