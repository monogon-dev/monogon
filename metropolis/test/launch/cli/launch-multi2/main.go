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

	metroctl "source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	"source.monogon.dev/metropolis/test/launch/cluster"
)

func main() {
	ctx := clicontext.WithInterrupt(context.Background())
	cl, err := cluster.LaunchCluster(ctx, cluster.ClusterOptions{
		NumNodes: 2,
	})
	if err != nil {
		log.Fatalf("LaunchCluster: %v", err)
	}

	mpath, err := cluster.MetroctlRunfilePath()
	if err != nil {
		log.Fatalf("MetroctlRunfilePath: %v", err)
	}
	wpath, err := cl.MakeMetroctlWrapper()
	if err != nil {
		log.Fatalf("MakeWrapper: %v", err)
	}

	apiservers, err := cl.KubernetesControllerNodeAddresses(ctx)
	if err != nil {
		log.Fatalf("Could not get Kubernetes controller nodes: %v", err)
	}
	if len(apiservers) < 1 {
		log.Fatalf("Cluster has no Kubernetes controller nodes")
	}

	configName := "launch-multi2"
	if err := metroctl.InstallKubeletConfig(mpath, cl.ConnectOptions(), configName, apiservers[0]); err != nil {
		log.Fatalf("InstallKubeletConfig: %v", err)
	}

	log.Printf("Launch: Cluster running!")
	log.Printf("  To access cluster use: metroctl %s", cl.MetroctlFlags())
	log.Printf("  Or use this handy wrapper: %s", wpath)
	log.Printf("  To access Kubernetes, use kubectl --context=%s", configName)

	<-ctx.Done()
	cl.Close()
}
