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
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"git.monogon.dev/source/nexantic.git/core/pkg/logtree"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/kubectl/pkg/cmd/plugin"
	"k8s.io/kubectl/pkg/util/logs"
	"k8s.io/kubernetes/pkg/kubectl/cmd"

	apb "git.monogon.dev/source/nexantic.git/core/proto/api"
)

func main() {
	ctx := context.Background()
	// Hardcode localhost since this should never be used to interface with a production node because of missing
	// encryption & authentication
	grpcClient, err := grpc.Dial("localhost:7837", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to dial debug service (is it running): %v\n", err)
	}
	debugClient := apb.NewNodeDebugServiceClient(grpcClient)
	if len(os.Args) < 2 {
		fmt.Println("Please specify a subcommand")
		os.Exit(1)
	}

	logsCmd := flag.NewFlagSet("logs", flag.ExitOnError)
	logsTailN := logsCmd.Int("tail", -1, "Get last n lines (-1 = whole buffer, 0 = disable)")
	logsStream := logsCmd.Bool("follow", false, "Stream log entries live from the system")
	logsRecursive := logsCmd.Bool("recursive", false, "Get entries from entire DN subtree")
	logsCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s %s [options] dn\n", os.Args[0], os.Args[1])
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "Example:\n  %s %s --tail 5 --follow init\n", os.Args[0], os.Args[1])
	}
	goldenticketCmd := flag.NewFlagSet("goldenticket", flag.ExitOnError)
	conditionCmd := flag.NewFlagSet("condition", flag.ExitOnError)
	conditionCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s %s [options] component_path\n", os.Args[0], os.Args[1])
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "Example:\n  %s %s IPAssigned\n", os.Args[0], os.Args[1])
	}

	switch os.Args[1] {
	case "logs":
		logsCmd.Parse(os.Args[2:])
		dn := logsCmd.Arg(0)
		req := &apb.GetLogsRequest{
			Dn:          dn,
			BacklogMode: apb.GetLogsRequest_BACKLOG_DISABLE,
			StreamMode:  apb.GetLogsRequest_STREAM_DISABLE,
			Filters:     nil,
		}

		switch *logsTailN {
		case 0:
		case -1:
			req.BacklogMode = apb.GetLogsRequest_BACKLOG_ALL
		default:
			req.BacklogMode = apb.GetLogsRequest_BACKLOG_COUNT
			req.BacklogCount = int64(*logsTailN)
		}

		if *logsStream {
			req.StreamMode = apb.GetLogsRequest_STREAM_UNBUFFERED
		}

		if *logsRecursive {
			req.Filters = append(req.Filters, &apb.LogFilter{
				Filter: &apb.LogFilter_WithChildren_{WithChildren: &apb.LogFilter_WithChildren{}},
			})
		}

		stream, err := debugClient.GetLogs(ctx, req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get logs: %v\n", err)
			os.Exit(1)
		}
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					os.Exit(0)
				}
				fmt.Fprintf(os.Stderr, "Failed to stream logs: %v\n", err)
				os.Exit(1)
			}
			for _, entry := range res.BacklogEntries {
				entry, err := logtree.LogEntryFromProto(entry)
				if err != nil {
					fmt.Printf("error decoding entry: %v", err)
					continue
				}
				fmt.Println(entry.String())
			}
		}
	case "goldenticket":
		goldenticketCmd.Parse(os.Args[2:])
		ip := goldenticketCmd.Arg(0)
		res, err := debugClient.GetGoldenTicket(ctx, &apb.GetGoldenTicketRequest{ExternalIp: ip})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get golden ticket: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(res.Ticket)
	case "kubectl":
		// Always get a kubeconfig with cluster-admin (group system:masters), kubectl itself can impersonate
		kubeconfigFile, err := ioutil.TempFile("", "dbg_kubeconfig")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create kubeconfig temp file: %v\n", err)
			os.Exit(1)
		}
		defer kubeconfigFile.Close()
		defer os.Remove(kubeconfigFile.Name())

		res, err := debugClient.GetDebugKubeconfig(ctx, &apb.GetDebugKubeconfigRequest{Id: "debug-user", Groups: []string{"system:masters"}})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get kubeconfig: %v\n", err)
			os.Exit(1)
		}
		if _, err := kubeconfigFile.WriteString(res.DebugKubeconfig); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write kubeconfig: %v\n", err)
			os.Exit(1)
		}

		// This magic sets up everything as if this was just the kubectl binary. It sets the KUBECONFIG environment
		// variable so that it knows where the Kubeconfig is located and forcibly overwrites the arguments so that
		// the "wrapper" arguments are not visible to its flags parser. The base code is straight from
		// https://github.com/kubernetes/kubernetes/blob/master/cmd/kubectl/kubectl.go
		os.Setenv("KUBECONFIG", kubeconfigFile.Name())
		rand.Seed(time.Now().UnixNano())
		pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
		pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
		logs.InitLogs()
		defer logs.FlushLogs()
		command := cmd.NewDefaultKubectlCommandWithArgs(cmd.NewDefaultPluginHandler(plugin.ValidPluginFilenamePrefixes), os.Args[2:], os.Stdin, os.Stdout, os.Stderr)
		command.SetArgs(os.Args[2:])
		if err := command.Execute(); err != nil {
			os.Exit(1)
		}
	}
}
