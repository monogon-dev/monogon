package main

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/proto/api"
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
	mgmt := api.NewManagementClient(cc)

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

	enc := newOutputEncoder()
	defer enc.close()

	for _, n := range nodes {
		// Filter the information we want client-side.
		if len(qids) != 0 {
			nid := identity.NodeID(n.Pubkey)
			if _, e := qids[nid]; !e {
				continue
			}
		}

		if err := enc.writeNodeID(n); err != nil {
			log.Fatalf("While listing nodes: %v", err)
		}
	}
}
