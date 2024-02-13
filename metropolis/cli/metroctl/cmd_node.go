package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"source.monogon.dev/go/clitable"
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

var nodeUpdateCmd = &cobra.Command{
	Short:   "Updates the operating system of a cluster node.",
	Use:     "update [NodeID] --bundle-url bundleURL [--activation-mode <none|reboot|kexec>]",
	Example: "metroctl node update --bundle-url https://example.com/bundle.zip --activation-mode reboot metropolis-25fa5f5e9349381d4a5e9e59de0215e3",
	RunE: func(cmd *cobra.Command, args []string) error {
		bundleUrl, err := cmd.Flags().GetString("bundle-url")
		if err != nil {
			return err
		}

		if len(bundleUrl) == 0 {
			return fmt.Errorf("flag bundle-url is required")
		}

		activationMode, err := cmd.Flags().GetString("activation-mode")
		if err != nil {
			return err
		}

		var am apb.ActivationMode
		switch strings.ToLower(activationMode) {
		case "none":
			am = apb.ActivationMode_ACTIVATION_NONE
		case "reboot":
			am = apb.ActivationMode_ACTIVATION_REBOOT
		case "kexec":
			am = apb.ActivationMode_ACTIVATION_KEXEC
		default:
			return fmt.Errorf("invalid value for flag activation-mode")
		}

		ctx := clicontext.WithInterrupt(context.Background())

		cacert, err := core.GetClusterCAWithTOFU(ctx, connectOptions())
		if err != nil {
			return fmt.Errorf("could not get CA certificate: %w", err)
		}

		mgmt := apb.NewManagementClient(dialAuthenticated(ctx))

		nodes, err := core.GetNodes(ctx, mgmt, "")
		if err != nil {
			return fmt.Errorf("while calling Management.GetNodes: %v", err)
		}
		// Narrow down the output set to supplied node IDs, if any.
		qids := make(map[string]bool)
		if len(args) != 0 && args[0] != "all" {
			for _, a := range args {
				qids[a] = true
			}
		}

		updateReq := &apb.UpdateNodeRequest{
			BundleUrl:      bundleUrl,
			ActivationMode: am,
		}

		for _, n := range nodes {
			// Filter the information we want client-side.
			if len(qids) != 0 {
				nid := identity.NodeID(n.Pubkey)
				if _, e := qids[nid]; !e {
					continue
				}
			}

			cc := dialAuthenticatedNode(ctx, n.Id, n.Status.ExternalAddress, cacert)
			nodeMgmt := apb.NewNodeManagementClient(cc)
			log.Printf("sending update request to: %s (%s)", n.Id, n.Status.ExternalAddress)
			_, err := nodeMgmt.UpdateNode(ctx, updateReq)
			if err != nil {
				return err
			}
		}

		return nil
	},
	Args: cobra.ExactArgs(1),
}

var nodeDeleteCmd = &cobra.Command{
	Short:   "Deletes a node from the cluster.",
	Use:     "delete [NodeID] [--bypass-has-roles] [--bypass-not-decommissioned]",
	Example: "metroctl node delete metropolis-25fa5f5e9349381d4a5e9e59de0215e3",
	RunE: func(cmd *cobra.Command, args []string) error {
		bypassHasRoles, err := cmd.Flags().GetBool("bypass-has-roles")
		if err != nil {
			return err
		}

		bypassNotDecommissioned, err := cmd.Flags().GetBool("bypass-not-decommissioned")
		if err != nil {
			return err
		}

		ctx := clicontext.WithInterrupt(context.Background())
		mgmt := apb.NewManagementClient(dialAuthenticated(ctx))

		nodes, err := core.GetNodes(ctx, mgmt, fmt.Sprintf("node.id==%q", args[0]))
		if err != nil {
			return fmt.Errorf("while calling Management.GetNodes: %v", err)
		}

		if len(nodes) == 0 {
			return fmt.Errorf("could not find node with id: %s", args[0])
		}

		if len(nodes) != 1 {
			return fmt.Errorf("expected one node, got %d", len(nodes))
		}

		n := nodes[0]
		log.Printf("deleting node: %s (%s)", n.Id, n.Status.ExternalAddress)

		req := &apb.DeleteNodeRequest{
			Node: &apb.DeleteNodeRequest_Id{
				Id: n.Id,
			},
		}

		if bypassHasRoles {
			req.SafetyBypassHasRoles = &apb.DeleteNodeRequest_SafetyBypassHasRoles{}
		}

		if bypassNotDecommissioned {
			req.SafetyBypassNotDecommissioned = &apb.DeleteNodeRequest_SafetyBypassNotDecommissioned{}
		}

		_, err = mgmt.DeleteNode(ctx, req)
		return err
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	nodeUpdateCmd.Flags().String("bundle-url", "", "The URL to the new version")
	nodeUpdateCmd.Flags().String("activation-mode", "reboot", "How the update should be activated (kexec, reboot, none)")

	nodeDeleteCmd.Flags().Bool("bypass-has-roles", false, "Allows to bypass the HasRoles check")
	nodeDeleteCmd.Flags().Bool("bypass-not-decommissioned", false, "Allows to bypass the NotDecommissioned check")

	nodeCmd.AddCommand(nodeDescribeCmd)
	nodeCmd.AddCommand(nodeListCmd)
	nodeCmd.AddCommand(nodeUpdateCmd)
	nodeCmd.AddCommand(nodeDeleteCmd)
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

	var t clitable.Table
	for _, n := range nodes {
		// Filter the information we want client-side.
		if len(qids) != 0 {
			nid := identity.NodeID(n.Pubkey)
			if _, e := qids[nid]; !e {
				continue
			}
		}
		t.Add(nodeEntry(n))
	}

	t.Print(o, onlyColumns)
}
