package main

import (
	"context"
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
	Args:    cobra.ArbitraryArgs,
	Run:     doAdd,
}

var removeRoleCmd = &cobra.Command{
	Short:   "Updates node roles.",
	Use:     "role <KubernetesController|KubernetesWorker|ConsensusMember> [NodeID, ...]",
	Example: "metroctl node remove role KubernetesWorker metropolis-25fa5f5e9349381d4a5e9e59de0215e3",
	Args:    cobra.ArbitraryArgs,
	Run:     doRemove,
}

func init() {
	addCmd.AddCommand(addRoleCmd)
	nodeCmd.AddCommand(addCmd)

	removeCmd.AddCommand(removeRoleCmd)
	nodeCmd.AddCommand(removeCmd)
}

func doAdd(cmd *cobra.Command, args []string) {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	cc := dialAuthenticated(ctx)
	mgmt := api.NewManagementClient(cc)

	if len(args) < 2 {
		log.Fatal("Provide the role parameter together with at least one node ID.")
	}

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
			log.Fatalf("Unknown role: %s", role)
		}

		_, err := mgmt.UpdateNodeRoles(ctx, req)
		if err != nil {
			log.Printf("Couldn't update node \"%s\": %v", node, err)
		}
		log.Printf("Updated node %s.", node)
	}
}

func doRemove(cmd *cobra.Command, args []string) {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	cc := dialAuthenticated(ctx)
	mgmt := api.NewManagementClient(cc)

	if len(args) < 2 {
		log.Fatal("Provide the role parameter together with at least one node ID.")
	}

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
			log.Fatalf("Unknown role: %s. Must be one of: KubernetesController, KubernetesWorker, ConsensusMember.", role)
		}

		_, err := mgmt.UpdateNodeRoles(ctx, req)
		if err != nil {
			log.Printf("Couldn't update node \"%s\": %v", node, err)
		}
		log.Printf("Updated node %s.", node)
	}
}
