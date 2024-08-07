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

	metroctl "source.monogon.dev/metropolis/cli/metroctl/core"
	mlaunch "source.monogon.dev/metropolis/test/launch"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	cl, err := mlaunch.LaunchCluster(ctx, mlaunch.ClusterOptions{
		NumNodes:        3,
		NodeLogsToFiles: true,
	})
	if err != nil {
		log.Fatalf("LaunchCluster: %v", err)
	}

	mpath, err := mlaunch.MetroctlRunfilePath()
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

	configName := "launch-cluster"
	if err := metroctl.InstallKubeletConfig(ctx, mpath, cl.ConnectOptions(), configName, apiservers[0]); err != nil {
		log.Fatalf("InstallKubeletConfig: %v", err)
	}

	log.Printf("Launch: Cluster running!")
	log.Printf("  To access cluster use: metroctl %s", cl.MetroctlFlags())
	log.Printf("  Or use this handy wrapper: %s", wpath)
	log.Printf("  To access Kubernetes, use kubectl --context=%s", configName)

	<-ctx.Done()
	cl.Close()
}
