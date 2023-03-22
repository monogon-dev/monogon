package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
)

var nodeCmd = &cobra.Command{
	Short: "Updates and queries node information.",
	Use:   "node",
}

var nodeDescribeCmd = &cobra.Command{
	Short:   "Describes cluster nodes.",
	Use:     "describe [node-id] [--filter] [--output] [--format]",
	Example: "metroctl node describe metropolis-c556e31c3fa2bf0a36e9ccb9fd5d6056",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := clicontext.WithInterrupt(context.Background())
		cc := dialAuthenticated(ctx)
		mgmt := apb.NewManagementClient(cc)

		nodes, err := core.GetNodes(ctx, mgmt, flags.filter)
		if err != nil {
			log.Fatalf("While calling Management.GetNodes: %v", err)
		}

		printNodes(nodes, args, nil)
	},
	Args: cobra.ArbitraryArgs,
}

var nodeListCmd = &cobra.Command{
	Short:   "Lists cluster nodes.",
	Use:     "list [node-id] [--filter] [--output] [--format]",
	Example: "metroctl node list --filter node.status.external_address==\"10.8.0.2\"",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := clicontext.WithInterrupt(context.Background())
		cc := dialAuthenticated(ctx)
		mgmt := apb.NewManagementClient(cc)

		nodes, err := core.GetNodes(ctx, mgmt, flags.filter)
		if err != nil {
			log.Fatalf("While calling Management.GetNodes: %v", err)
		}

		printNodes(nodes, args, map[string]bool{"node id": true})
	},
	Args: cobra.ArbitraryArgs,
}

func init() {
	nodeCmd.AddCommand(nodeDescribeCmd)
	nodeCmd.AddCommand(nodeListCmd)
	rootCmd.AddCommand(nodeCmd)
}

func printNodes(nodes []*apb.Node, args []string, onlyColumns map[string]bool) {
	o := io.WriteCloser(os.Stdout)
	if flags.output != "" {
		of, err := os.Create(flags.output)
		if err != nil {
			log.Fatalf("Couldn't create the output file at %s: %v", flags.output, err)
		}
		o = of
	}

	// Narrow down the output set to supplied node IDs, if any.
	qids := make(map[string]bool)
	if len(args) != 0 && args[0] != "all" {
		for _, a := range args {
			qids[a] = true
		}
	}

	var t table
	for _, n := range nodes {
		// Filter the information we want client-side.
		if len(qids) != 0 {
			nid := identity.NodeID(n.Pubkey)
			if _, e := qids[nid]; !e {
				continue
			}
		}
		t.add(nodeEntry(n))
	}

	t.print(o, onlyColumns)
}
