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

/*
Copyright 2014 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This is the entry point for our multicall Kubernetes binary. It can act as
// any of the Kubernetes components we use depending on its first argument.
// This saves us a bunch of duplicated code and thus system partition size as
// a large amount of library code is shared between all of the Kubernetes
// components.
//
// As this is not intended by the K8s developers the Cobra setup is unusual
// in that even the command structs are only created on-demand and not
// registered with AddCommand. This is done as Kubernetes performs one-off
// global setup inside their NewXYZCommand functions, for example for signal
// handling and their global registries.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/component-base/cli"
	_ "k8s.io/component-base/metrics/prometheus/restclient" // for client metric registration
	_ "k8s.io/component-base/metrics/prometheus/version"    // for version metric registration
	kubeapiserver "k8s.io/kubernetes/cmd/kube-apiserver/app"
	kubecontrollermanager "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	kubescheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
	kubelet "k8s.io/kubernetes/cmd/kubelet/app"
)

// Map of subcommand to Cobra command generator for all subcommands
var subcommands = map[string]func() *cobra.Command{
	"kube-apiserver":          kubeapiserver.NewAPIServerCommand,
	"kube-controller-manager": kubecontrollermanager.NewControllerManagerCommand,
	"kube-scheduler":          func() *cobra.Command { return kubescheduler.NewSchedulerCommand() },
	"kubelet":                 kubelet.NewKubeletCommand,
}

func main() {
	if len(os.Args) < 2 || subcommands[os.Args[1]] == nil {
		fmt.Fprintf(os.Stderr, "Unknown subcommand\n")
	} else {
		cmdGen := subcommands[os.Args[1]]
		cmd := cmdGen()
		// Strip first argument as it has already been consumed
		cmd.SetArgs(os.Args[2:])
		os.Exit(cli.Run(cmd))
	}
}
