package main

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	clicontext "source.monogon.dev/metropolis/cli/pkg/context"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/proto/api"
)

var approveCmd = &cobra.Command{
	Short: "Approves a candidate node, if specified; lists nodes pending approval otherwise.",
	Use:   "approve [node-id]",
	Args:  cobra.MaximumNArgs(1), // One positional argument: node ID
	Run:   doApprove,
}

func init() {
	rootCmd.AddCommand(approveCmd)
}

// nodeById returns the node matching id, if it exists within nodes.
func nodeById(nodes []*api.Node, id string) *api.Node {
	for _, n := range nodes {
		if identity.NodeID(n.Pubkey) == id {
			return n
		}
	}
	return nil
}

func doApprove(cmd *cobra.Command, args []string) {
	cc := dialAuthenticated()
	mgmt := api.NewManagementClient(cc)

	// Get a list of all nodes pending approval by calling Management.GetNodes.
	// We need this list regardless of whether we're actually approving nodes, or
	// just listing them.
	ctx := clicontext.WithInterrupt(context.Background())
	nodes, err := core.GetNodes(ctx, mgmt, "node.state == NODE_STATE_NEW")
	if err != nil {
		log.Fatalf("While fetching a list of nodes pending approval: %v", err)
	}

	if len(args) == 0 {
		// If no id was given, just list the nodes pending approval.
		if len(nodes) != 0 {
			for _, n := range nodes {
				fmt.Print(identity.NodeID(n.Pubkey))
			}
		} else {
			log.Print("There are no nodes pending approval at this time.")
		}
	} else {
		// Otherwise, try to approve the node matching the id.
		tgtNodeId := args[0]

		n := nodeById(nodes, tgtNodeId)
		if n == nil {
			log.Fatalf("Couldn't find a new node matching id %s", tgtNodeId)
		}
		_, err := mgmt.ApproveNode(ctx, &api.ApproveNodeRequest{
			Pubkey: n.Pubkey,
		})
		if err != nil {
			log.Fatalf("While approving node %s: %v", tgtNodeId, err)
		}
		log.Printf("Approved node %s.", tgtNodeId)
	}
}
