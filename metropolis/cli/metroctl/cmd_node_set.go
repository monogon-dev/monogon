package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/proto/api"
)

var addCmd = &cobra.Command{
	Short: "Updates node configuration.",
	Use:   "add",
}

var removeCmd = &cobra.Command{
	Short: "Updates node configuration.",
	Use:   "remove",
}

var addRoleCmd = &cobra.Command{
	Short:   "Updates node roles.",
	Use:     "role <KubernetesController|KubernetesWorker|ConsensusMember> [NodeID, ...]",
	Example: "metroctl node add role KubernetesWorker metropolis-25fa5f5e9349381d4a5e9e59de0215e3",
	Args:    PrintUsageOnWrongArgs(cobra.MinimumNArgs(2)),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
		cc, err := dialAuthenticated(ctx)
		if err != nil {
			return fmt.Errorf("while dialing node: %w", err)
		}
		mgmt := api.NewManagementClient(cc)

		role := strings.ToLower(args[0])
		nodes := args[1:]

		opt := func(v bool) *bool { return &v }
		for _, node := range nodes {
			req := &api.UpdateNodeRolesRequest{
				Node: &api.UpdateNodeRolesRequest_Id{
					Id: node,
				},
			}
			switch role {
			case "kubernetescontroller", "kc":
				req.KubernetesController = opt(true)
			case "kubernetesworker", "kw":
				req.KubernetesWorker = opt(true)
			case "consensusmember", "cm":
				req.ConsensusMember = opt(true)
			default:
				return fmt.Errorf("unknown role: %s", role)
			}

			_, err := mgmt.UpdateNodeRoles(ctx, req)
			if err != nil {
				log.Printf("Couldn't update node \"%s\": %v", node, err)
			} else {
				log.Printf("Updated node %s.", node)
			}
		}
		return nil
	},
}

var removeRoleCmd = &cobra.Command{
	Short:   "Updates node roles.",
	Use:     "role <KubernetesController|KubernetesWorker|ConsensusMember> [NodeID, ...]",
	Example: "metroctl node remove role KubernetesWorker metropolis-25fa5f5e9349381d4a5e9e59de0215e3",
	Args:    PrintUsageOnWrongArgs(cobra.ArbitraryArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
		cc, err := dialAuthenticated(ctx)
		if err != nil {
			return fmt.Errorf("while dialing node: %w", err)
		}
		mgmt := api.NewManagementClient(cc)

		role := strings.ToLower(args[0])
		nodes := args[1:]

		opt := func(v bool) *bool { return &v }
		for _, node := range nodes {
			req := &api.UpdateNodeRolesRequest{
				Node: &api.UpdateNodeRolesRequest_Id{
					Id: node,
				},
			}

			switch role {
			case "kubernetescontroller", "kc":
				req.KubernetesController = opt(false)
			case "kubernetesworker", "kw":
				req.KubernetesWorker = opt(false)
			case "consensusmember", "cm":
				req.ConsensusMember = opt(false)
			default:
				return fmt.Errorf("unknown role: %s. Must be one of: KubernetesController, KubernetesWorker, ConsensusMember", role)
			}

			_, err := mgmt.UpdateNodeRoles(ctx, req)
			if err != nil {
				log.Printf("Couldn't update node \"%s\": %v", node, err)
			} else {
				log.Printf("Updated node %s.", node)
			}
		}
		return nil
	},
}

func init() {
	addCmd.AddCommand(addRoleCmd)
	nodeCmd.AddCommand(addCmd)

	removeCmd.AddCommand(removeRoleCmd)
	nodeCmd.AddCommand(removeCmd)
}
