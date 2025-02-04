// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import "source.monogon.dev/metropolis/node/core/network"

// initializeDebugger does nothing in a non-debug build
func initializeDebugger(*network.Service) {
}
