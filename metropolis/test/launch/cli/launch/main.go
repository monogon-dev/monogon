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
	"log"
	"os"
	"os/signal"
	"path/filepath"

	apb "source.monogon.dev/metropolis/proto/api"

	mlaunch "source.monogon.dev/metropolis/test/launch"
	"source.monogon.dev/osbase/test/launch"
)

func main() {
	// Create the launch directory.
	ld, err := os.MkdirTemp(os.Getenv("TEST_TMPDIR"), "node_state*")
	if err != nil {
		log.Fatalf("couldn't create a launch directory: %v", err)
	}
	defer os.RemoveAll(ld)
	// Create the socket directory. Since using TEST_TMPDIR will often result in
	// paths too long to place UNIX sockets at, we'll use the LSB temporary
	// directory.
	sd, err := os.MkdirTemp("/tmp", "node_sock*")
	if err != nil {
		log.Fatalf("couldn't create a socket directory: %v", err)
	}
	defer os.RemoveAll(sd)

	var ports []uint16
	for _, p := range mlaunch.NodePorts {
		ports = append(ports, uint16(p))
	}
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	doneC := make(chan error)
	tpmf, err := mlaunch.NewTPMFactory(filepath.Join(ld, "tpm"))
	if err != nil {
		log.Fatalf("NewTPMFactory: %v", err)
	}
	err = mlaunch.LaunchNode(ctx, ld, sd, tpmf, &mlaunch.NodeOptions{
		Name:       "test-node",
		Ports:      launch.IdentityPortMap(ports),
		SerialPort: os.Stdout,
		RunVNC:     true,
		NodeParameters: &apb.NodeParameters{
			Cluster: &apb.NodeParameters_ClusterBootstrap_{
				ClusterBootstrap: mlaunch.InsecureClusterBootstrap,
			},
		},
	}, doneC)
	if err != nil {
		log.Fatalf("LaunchNode: %v", err)
	}
	err = <-doneC
	if err != nil {
		log.Fatalf("Node returned: %v", err)
	}
}
