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

	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	apb "source.monogon.dev/metropolis/proto/api"
	"source.monogon.dev/metropolis/test/launch"
	"source.monogon.dev/metropolis/test/launch/cluster"
)

func main() {
	ctx := clicontext.WithInterrupt(context.Background())
	err := cluster.LaunchNode(ctx, cluster.NodeOptions{
		Ports:      launch.IdentityPortMap(cluster.NodePorts),
		SerialPort: os.Stdout,
		NodeParameters: &apb.NodeParameters{
			Cluster: &apb.NodeParameters_ClusterBootstrap_{
				ClusterBootstrap: cluster.InsecureClusterBootstrap,
			},
		},
	})
	if err != nil {
		log.Fatalf("LaunchNode: %v", err)
	}
}
