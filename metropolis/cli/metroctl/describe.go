package main

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	"source.monogon.dev/metropolis/node/core/identity"
	apb "source.monogon.dev/metropolis/proto/api"
)

var describeCmd = &cobra.Command{
	Short:   "Describes cluster nodes.",
	Use:     "describe [node-id] [--filter] [--output] [--format]",
	Example: "metroctl node describe metropolis-c556e31c3fa2bf0a36e9ccb9fd5d6056",
	Run:     doDescribe,
	Args:    cobra.ArbitraryArgs,
}

func init() {
	nodeCmd.AddCommand(describeCmd)
}

func printNodes(of func(*encoder, *apb.Node) error, nodes []*apb.Node, args []string) {
	// Narrow down the output set to supplied node IDs, if any.
	qids := make(map[string]bool)
	if len(args) != 0 && args[0] != "all" {
		for _, a := range args {
			qids[a] = true
		}
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

		if err := of(enc, n); err != nil {
			log.Fatalf("While listing nodes: %v", err)
		}
	}
}

func doDescribe(cmd *cobra.Command, args []string) {
	ctx := clicontext.WithInterrupt(context.Background())
	cc := dialAuthenticated(ctx)
	mgmt := apb.NewManagementClient(cc)

	nodes, err := core.GetNodes(ctx, mgmt, flags.filter)
	if err != nil {
		log.Fatalf("While calling Management.GetNodes: %v", err)
	}

	of := func(enc *encoder, n *apb.Node) error {
		return enc.writeNode(n)
	}
	printNodes(of, nodes, args)
}
