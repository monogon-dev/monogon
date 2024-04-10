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

// Adapted from https://github.com/dims/hyperkube

package main

import (
	goflag "flag"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	_ "k8s.io/component-base/metrics/prometheus/restclient" // for client metric registration
	_ "k8s.io/component-base/metrics/prometheus/version"    // for version metric registration
	kubeapiserver "k8s.io/kubernetes/cmd/kube-apiserver/app"
	kubecontrollermanager "k8s.io/kubernetes/cmd/kube-controller-manager/app"
	kubescheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
	kubelet "k8s.io/kubernetes/cmd/kubelet/app"
)

func main() {
	hyperkubeCommand, allCommandFns := NewHyperKubeCommand()

	// TODO: once we switch everything over to Cobra commands, we can go back
	// to calling cliflag.InitFlags() (by removing its pflag.Parse() call). For
	// now, we have to set the normalize func and add the go flag set by hand.
	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	// cliflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	basename := filepath.Base(os.Args[0])
	if err := commandFor(basename, hyperkubeCommand, allCommandFns).Execute(); err != nil {
		os.Exit(1)
	}
}

func commandFor(basename string, defaultCommand *cobra.Command, commands []func() *cobra.Command) *cobra.Command {
	for _, commandFn := range commands {
		command := commandFn()
		if command.Name() == basename {
			return command
		}
		for _, alias := range command.Aliases {
			if alias == basename {
				return command
			}
		}
	}

	return defaultCommand
}

// NewHyperKubeCommand is the entry point for hyperkube
func NewHyperKubeCommand() (*cobra.Command, []func() *cobra.Command) {
	// these have to be functions since the command is polymorphic. Cobra wants
	// you to be top level command to get executed
	apiserver := func() *cobra.Command { return kubeapiserver.NewAPIServerCommand() }
	controller := func() *cobra.Command { return kubecontrollermanager.NewControllerManagerCommand() }
	scheduler := func() *cobra.Command { return kubescheduler.NewSchedulerCommand() }
	kubelet := func() *cobra.Command { return kubelet.NewKubeletCommand() }

	commandFns := []func() *cobra.Command{
		apiserver,
		controller,
		scheduler,
		kubelet,
	}

	cmd := &cobra.Command{
		Use:   "kube",
		Short: "Combines all Kubernetes components in a single binary",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	for i := range commandFns {
		cmd.AddCommand(commandFns[i]())
	}

	return cmd, commandFns
}
