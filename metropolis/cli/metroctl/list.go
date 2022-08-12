package main

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	apb "source.monogon.dev/metropolis/proto/api"
)

var listCmd = &cobra.Command{
	Short:   "Lists cluster nodes.",
	Use:     "list [node-id] [--filter] [--output] [--format]",
	Example: "metroctl node list --filter node.status.external_address==\"10.8.0.2\"",
	Run:     doList,
	Args:    cobra.ArbitraryArgs,
}

func init() {
	nodeCmd.AddCommand(listCmd)
}

func doList(cmd *cobra.Command, args []string) {
	ctx := clicontext.WithInterrupt(context.Background())
	cc := dialAuthenticated(ctx)
	mgmt := apb.NewManagementClient(cc)

	// Narrow down the output set to supplied node IDs, if any.
	qids := make(map[string]bool)
	if len(args) != 0 && args[0] != "all" {
		for _, a := range args {
			qids[a] = true
		}
	}

	nodes, err := core.GetNodes(ctx, mgmt, flags.filter)
	if err != nil {
		log.Fatalf("While calling Management.GetNodes: %v", err)
	}

	of := func(enc *encoder, n *apb.Node) error {
		return enc.writeNodeID(n)
	}
	printNodes(of, nodes, args)
}
