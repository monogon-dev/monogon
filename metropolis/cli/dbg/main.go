// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/component-base/cli"
	"k8s.io/kubectl/pkg/cmd"
	"k8s.io/kubectl/pkg/cmd/plugin"
	"k8s.io/kubectl/pkg/cmd/util"

	"source.monogon.dev/osbase/logtree"

	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

func main() {
	ctx := context.Background()
	// Hardcode localhost since this should never be used to interface with a
	// production node because of missing encryption & authentication
	grpcClient, err := grpc.Dial("localhost:7837", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	conditionCmd := flag.NewFlagSet("condition", flag.ExitOnError)
	conditionCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s %s [options] component_path\n", os.Args[0], os.Args[1])
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "Example:\n  %s %s IPAssigned\n", os.Args[0], os.Args[1])
	}

	traceCmd := flag.NewFlagSet("trace", flag.ExitOnError)
	traceCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v %v [options] tracer\n", os.Args[0], os.Args[1])
		flag.PrintDefaults()
	}
	functionFilter := traceCmd.String("function_filter", "", "Only trace functions matched by this filter (comma-separated, supports wildcards via *)")
	functionGraphFilter := traceCmd.String("function_graph_filter", "", "Only trace functions matched by this filter and their children (syntax same as function_filter)")

	loadimageCmd := flag.NewFlagSet("loadimage", flag.ExitOnError)
	loadimageCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v %v [options] image\n", os.Args[0], os.Args[1])
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "Example: %v %v [options] helloworld_oci.tar.gz\n", os.Args[0], os.Args[1])
	}

	switch os.Args[1] {
	case "logs":
		logsCmd.Parse(os.Args[2:])
		dn := logsCmd.Arg(0)
		req := &apb.LogsRequest{
			Dn:          dn,
			BacklogMode: apb.LogsRequest_BACKLOG_MODE_DISABLE,
			StreamMode:  apb.LogsRequest_STREAM_MODE_DISABLE,
			Filters:     nil,
		}

		switch *logsTailN {
		case 0:
		case -1:
			req.BacklogMode = apb.LogsRequest_BACKLOG_MODE_ALL
		default:
			req.BacklogMode = apb.LogsRequest_BACKLOG_MODE_COUNT
			req.BacklogCount = int64(*logsTailN)
		}

		if *logsStream {
			req.StreamMode = apb.LogsRequest_STREAM_MODE_UNBUFFERED
		}

		if *logsRecursive {
			req.Filters = append(req.Filters, &cpb.LogFilter{
				Filter: &cpb.LogFilter_WithChildren_{WithChildren: &cpb.LogFilter_WithChildren{}},
			})
		}

		stream, err := debugClient.Logs(ctx, req)
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
			for _, entry := range res.StreamEntries {
				entry, err := logtree.LogEntryFromProto(entry)
				if err != nil {
					fmt.Printf("error decoding entry: %v", err)
					continue
				}
				fmt.Println(entry.String())
			}
		}
	case "kubectl":
		// Pop "kubectl" arg (the k8s cli library internally parses os.Args).
		os.Args = os.Args[1:]

		// Always get a kubeconfig with cluster-admin (group system:masters),
		// kubectl itself can impersonate
		kubeconfigFile, err := os.CreateTemp("", "dbg_kubeconfig")
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

		// This magic sets up everything as if this was just the kubectl
		// binary. It sets the KUBECONFIG environment variable so that it knows
		// where the Kubeconfig is located and forcibly overwrites the
		// arguments so that the "wrapper" arguments are not visible to its
		// flags parser.
		// The base code is straight from:
		//   https://github.com/kubernetes/kubernetes/blob/master/cmd/kubectl/kubectl.go
		os.Setenv("KUBECONFIG", kubeconfigFile.Name())
		command := cmd.NewDefaultKubectlCommandWithArgs(cmd.KubectlOptions{
			PluginHandler: cmd.NewDefaultPluginHandler(plugin.ValidPluginFilenamePrefixes),
			Arguments:     os.Args[2:],
			ConfigFlags:   genericclioptions.NewConfigFlags(true).WithDiscoveryBurst(300).WithDiscoveryQPS(50.0),
			IOStreams:     genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr},
		})
		if err := cli.RunNoErrOutput(command); err != nil {
			// Pretty-print the error and exit with an error.
			util.CheckErr(err)
		}
	case "trace":
		traceCmd.Parse(os.Args[2:])
		tracer := traceCmd.Arg(0)
		var fgf []string
		var ff []string
		if len(*functionGraphFilter) > 0 {
			fgf = strings.Split(*functionGraphFilter, ",")
		}
		if len(*functionFilter) > 0 {
			ff = strings.Split(*functionFilter, ",")
		}
		req := apb.TraceRequest{
			GraphFunctionFilter: fgf,
			FunctionFilter:      ff,
			Tracer:              tracer,
		}
		traceEvents, err := debugClient.Trace(ctx, &req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to trace: %v", err)
			os.Exit(1)
		}
		for {
			traceEvent, err := traceEvents.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Fprintf(os.Stderr, "stream aborted unexpectedly: %v", err)
				os.Exit(1)
			}
			fmt.Println(traceEvent.RawLine)
		}
	case "loadimage":
		loadimageCmd.Parse(os.Args[2:])
		imagePath := loadimageCmd.Arg(0)
		image, err := os.Open(imagePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open image file: %v\n", err)
			os.Exit(1)
		}
		defer image.Close()

		loadStream, err := debugClient.LoadImage(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to start loading image: %v\n", err)
			os.Exit(1)
		}
		buf := make([]byte, 64*1024)
		for {
			n, err := image.Read(buf)
			if err != io.EOF && err != nil {
				fmt.Fprintf(os.Stderr, "failed to read image: %v\n", err)
				os.Exit(1)
			}
			if err == io.EOF && n == 0 {
				break
			}
			if err := loadStream.Send(&apb.ImagePart{DataPart: buf[:n]}); err != nil {
				fmt.Fprintf(os.Stderr, "failed to send image part: %v\n", err)
				os.Exit(1)
			}
		}
		if _, err := loadStream.CloseAndRecv(); err != nil {
			fmt.Fprintf(os.Stderr, "image failed to load: %v\n", err)
		}
		fmt.Fprintf(os.Stderr, "Image loaded into Metropolis node\n")
	}
}
