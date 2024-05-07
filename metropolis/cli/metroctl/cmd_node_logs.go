package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"source.monogon.dev/metropolis/cli/metroctl/core"
	"source.monogon.dev/metropolis/pkg/logtree"
	lpb "source.monogon.dev/metropolis/pkg/logtree/proto"
	"source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

type metroctlLogFlags struct {
	// follow (ie. stream) logs live.
	follow bool
	// dn to query.
	dn string
	// exact dn query, i.e. without children/recursion.
	exact bool
	// concise logging output format.
	concise bool
	// backlog: >0 for a concrete limit, -1 for all, 0 for none
	backlog int
}

var logFlags metroctlLogFlags

var nodeLogsCmd = &cobra.Command{
	Short: "Get/stream logs from node",
	Long: `Get or stream logs from node.

Node logs are structured in a 'log tree' structure, in which different subsystems
log to DNs (distinguished names). For example, service 'foo' might log to
root.role.foo, while service 'bar' might log to root.role.bar.

To set the DN you want to request logs from, use --dn. The default is to return
all logs. The default output is also also a good starting point to figure out
what DNs are active in the system.

When requesting logs for a DN by default all sub-DNs will also be returned (ie.
with the above example, when requesting DN 'root.role' logs at root.role.foo and
root.role.bar would also be returned). This behaviour can be disabled by setting
--exact.

To stream logs, use --follow.

By default, all available logs are returned. To limit the number of historical
log lines (a.k.a. 'backlog') to return, set --backlog. This similar to requesting
all lines and then piping the result through 'tail' - but more efficient, as no
unnecessary lines are fetched.
`,
	Use:  "logs [node-id]",
	Args: cobra.MinimumNArgs(1),
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

		cacert, err := core.GetClusterCAWithTOFU(ctx, connectOptions())
		if err != nil {
			return fmt.Errorf("could not get CA certificate: %w", err)
		}

		fmt.Printf("=== Logs from %s (%s):\n", n.Id, n.Status.ExternalAddress)
		// Dial the actual node at its management port.
		cl := dialAuthenticatedNode(ctx, n.Id, n.Status.ExternalAddress, cacert)
		nmgmt := api.NewNodeManagementClient(cl)

		streamMode := api.GetLogsRequest_STREAM_DISABLE
		if logFlags.follow {
			streamMode = api.GetLogsRequest_STREAM_UNBUFFERED
		}
		var filters []*cpb.LogFilter
		if !logFlags.exact {
			filters = append(filters, &cpb.LogFilter{
				Filter: &cpb.LogFilter_WithChildren_{
					WithChildren: &cpb.LogFilter_WithChildren{},
				},
			})
		}
		backlogMode := api.GetLogsRequest_BACKLOG_ALL
		var backlogCount int64
		switch {
		case logFlags.backlog > 0:
			backlogMode = api.GetLogsRequest_BACKLOG_COUNT
			backlogCount = int64(logFlags.backlog)
		case logFlags.backlog == 0:
			backlogMode = api.GetLogsRequest_BACKLOG_DISABLE
		}

		srv, err := nmgmt.Logs(ctx, &api.GetLogsRequest{
			Dn:           logFlags.dn,
			BacklogMode:  backlogMode,
			BacklogCount: backlogCount,
			StreamMode:   streamMode,
			Filters:      filters,
		})
		if err != nil {
			return fmt.Errorf("failed to get logs: %w", err)
		}
		for {
			res, err := srv.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("=== Done.")
				break
			}
			if err != nil {
				return fmt.Errorf("log stream failed: %w", err)
			}
			for _, entry := range res.BacklogEntries {
				printEntry(entry)
			}
			for _, entry := range res.StreamEntries {
				printEntry(entry)
			}
		}

		return nil
	},
}

func printEntry(e *lpb.LogEntry) {
	entry, err := logtree.LogEntryFromProto(e)
	if err != nil {
		fmt.Printf("invalid stream entry: %v\n", err)
		return
	}
	if logFlags.concise {
		fmt.Println(entry.ConciseString(logtree.MetropolisShortenDict, 0))
	} else {
		fmt.Println(entry.String())
	}
}

func init() {
	nodeLogsCmd.Flags().BoolVarP(&logFlags.follow, "follow", "f", false, "Continue streaming logs after fetching backlog.")
	nodeLogsCmd.Flags().StringVar(&logFlags.dn, "dn", "", "Distinguished Name to get logs from (and children, if --exact is not set). If not set, defaults to '', which is the top-level DN.")
	nodeLogsCmd.Flags().BoolVarP(&logFlags.exact, "exact", "e", false, "Only show logs for exactly the DN, do not recurse down the tree.")
	nodeLogsCmd.Flags().BoolVarP(&logFlags.concise, "concise", "c", false, "Output concise logs.")
	nodeLogsCmd.Flags().IntVar(&logFlags.backlog, "backlog", -1, "How many lines of historical log data to return. The default (-1) returns all available lines. Zero value means no backlog is returned (useful when using --follow).")
	nodeCmd.AddCommand(nodeLogsCmd)
}
