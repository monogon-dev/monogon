package main

import (
	"crypto/x509"
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/proto/api"
)

var nodeLogsCmd = &cobra.Command{
	Short: "Get/stream logs from node",
	Use:   "logs [node-id]",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		// First connect to the main management service and figure out the node's IP
		// address.
		cc := dialAuthenticated(ctx)
		mgmt := api.NewManagementClient(cc)
		nodes, err := core.GetNodes(ctx, mgmt, fmt.Sprintf("node.id == %q", args[0]))
		if err != nil {
			return fmt.Errorf("when getting node info: %w", err)
		}

		if len(nodes) == 0 {
			return fmt.Errorf("no such node")
		}
		if len(nodes) > 1 {
			return fmt.Errorf("expression matched more than one node")
		}
		n := nodes[0]
		if n.Status == nil || n.Status.ExternalAddress == "" {
			return fmt.Errorf("node has no external address")
		}

		// TODO(q3k): save CA certificate on takeover
		info, err := mgmt.GetClusterInfo(ctx, &api.GetClusterInfoRequest{})
		if err != nil {
			return fmt.Errorf("couldn't get cluster info: %w", err)
		}
		cacert, err := x509.ParseCertificate(info.CaCertificate)
		if err != nil {
			return fmt.Errorf("remote CA certificate invalid: %w", err)
		}

		fmt.Printf("Getting logs from %s (%s)...\n", n.Id, n.Status.ExternalAddress)
		// Dial the actual node at its management port.
		cl := dialAuthenticatedNode(ctx, n.Id, n.Status.ExternalAddress, cacert)
		nmgmt := api.NewNodeManagementClient(cl)

		srv, err := nmgmt.Logs(ctx, &api.GetLogsRequest{
			Dn:          "",
			BacklogMode: api.GetLogsRequest_BACKLOG_ALL,
			StreamMode:  api.GetLogsRequest_STREAM_DISABLE,
			Filters: []*api.LogFilter{
				{
					Filter: &api.LogFilter_WithChildren_{
						WithChildren: &api.LogFilter_WithChildren{},
					},
				},
			},
		})
		if err != nil {
			return fmt.Errorf("failed to get logs: %w", err)
		}
		for {
			res, err := srv.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("Done.")
				break
			}
			if err != nil {
				return fmt.Errorf("log stream failed: %w", err)
			}
			for _, entry := range res.BacklogEntries {
				entry, err := logtree.LogEntryFromProto(entry)
				if err != nil {
					fmt.Printf("invalid entry: %v\n", err)
					continue
				}
				fmt.Println(entry.String())
			}
		}

		return nil
	},
}
