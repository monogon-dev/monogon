package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

type configurableClusterKey struct {
	key         string
	description string
	set         func(value []string) (*apb.ConfigureClusterRequest, error)
	get         func(c *cpb.ClusterConfiguration) (string, error)
}

var configurableClusterKeys = []configurableClusterKey{
	{
		key:         "kubernetes.node_labels_to_synchronize",
		description: "list of label regexes to sync from Metropolis to Kubernetes nodes",
		set: func(value []string) (*apb.ConfigureClusterRequest, error) {
			res := &apb.ConfigureClusterRequest{
				NewConfig: &cpb.ClusterConfiguration{
					Kubernetes: &cpb.ClusterConfiguration_Kubernetes{},
				},
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{"kubernetes.node_labels_to_synchronize"},
				},
			}
			for _, v := range value {
				_, err := regexp.Compile(v)
				if err != nil {
					return nil, fmt.Errorf("%q is not a valid regexp: %w", v, err)
				}
				res.NewConfig.Kubernetes.NodeLabelsToSynchronize = append(res.NewConfig.Kubernetes.NodeLabelsToSynchronize, &cpb.ClusterConfiguration_Kubernetes_NodeLabelsToSynchronize{
					Regexp: v,
				})
			}
			return res, nil
		},
		get: func(c *cpb.ClusterConfiguration) (string, error) {
			var res []string
			if kc := c.Kubernetes; kc != nil {
				for _, r := range kc.NodeLabelsToSynchronize {
					res = append(res, fmt.Sprintf("%q", r.Regexp))
				}
			}
			return strings.Join(res, ", "), nil
		},
	},
}

var clusterConfigureCommand = &cobra.Command{
	Use:   "configure <set/get> <field>",
	Short: "Gets/sets values in the cluster configuration structure",
	Long: `Gets/sets values in the cluster configuration structure.

The cluster is configured through a ClusterConfiguration structure, of which
a subset of fields can be modified. To set a field's value, use:

    cluster configure set <field> <value>

To get a field's current value, use the:

    cluster configure get <field>

Available configuration fields:

`,
	Args: PrintUsageOnWrongArgs(cobra.MinimumNArgs(2)),
	RunE: func(cmd *cobra.Command, args []string) error {
		mode := strings.ToLower(args[0])
		isSet := false
		switch mode {
		case "set":
			isSet = true
		case "get":
		default:
			return fmt.Errorf("invalid mode %q: must be set or get", mode)
		}

		var key *configurableClusterKey
		for _, k := range configurableClusterKeys {
			if k.key == args[1] {
				key = &k
				break
			}
		}
		if key == nil {
			return fmt.Errorf("unknown field %q, see help for list of supported fields", args[1])
		}

		ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
		cc, err := dialAuthenticated(ctx)
		if err != nil {
			return fmt.Errorf("while dialing node: %w", err)
		}
		mgmt := apb.NewManagementClient(cc)

		if isSet {
			req, err := key.set(args[2:])
			if err != nil {
				return err
			}
			res, err := mgmt.ConfigureCluster(ctx, req)
			if err != nil {
				return fmt.Errorf("could not mutate config: %w", err)
			}
			newValue, err := key.get(res.ResultingConfig)
			if err != nil {
				return fmt.Errorf("could not extract value from new config: %w", err)
			}
			log.Printf("New value: %s", newValue)
		} else {
			if len(args[2:]) > 0 {
				return fmt.Errorf("get <field> takes no extra arguments")
			}
			ci, err := mgmt.GetClusterInfo(ctx, &apb.GetClusterInfoRequest{})
			if err != nil {
				return fmt.Errorf("could not get cluster information: %w", err)
			}
			newValue, err := key.get(ci.ClusterConfiguration)
			if err != nil {
				return fmt.Errorf("could not extract value from new config: %w", err)
			}
			log.Printf("Value: %s", newValue)
		}
		return nil
	},
}

func init() {
	for _, key := range configurableClusterKeys {
		clusterConfigureCommand.Long += fmt.Sprintf("  - %s: %s", key.key, key.description)
	}
	clusterCmd.AddCommand(clusterConfigureCommand)
}
