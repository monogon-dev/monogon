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
	"context"
	"fmt"
	"os/exec"

	"source.monogon.dev/metropolis/node/"
	"source.monogon.dev/metropolis/node/core/network"
)

// initializeDebugger attaches Delve to ourselves and exposes it on common.DebuggerPort
// This is coupled to compilation_mode=dbg because otherwise Delve doesn't have the necessary DWARF debug info
func initializeDebugger(networkSvc *network.Service) {
	go func() {
		// This is intentionally delayed until network becomes available since Delve for some reason connects to itself
		// and in early-boot no network interface is available to do that through. Also external access isn't possible
		// early on anyways.
		networkSvc.GetIP(context.Background(), true)
		dlvCmd := exec.Command("/dlv", "--headless=true", fmt.Sprintf("--listen=:%v", node.DebuggerPort),
			"--accept-multiclient", "--only-same-user=false", "attach", "--continue", "1", "/init")
		if err := dlvCmd.Start(); err != nil {
			panic(err)
		}
	}()
}
