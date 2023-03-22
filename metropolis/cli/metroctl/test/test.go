package test

import (
	"bufio"
	"context"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"source.monogon.dev/metropolis/cli/pkg/datafile"
	"source.monogon.dev/metropolis/pkg/cmd"
	"source.monogon.dev/metropolis/test/launch/cluster"
	"source.monogon.dev/metropolis/test/util"
)

// resolveMetroctl resolves metroctl filesystem path. It will return a correct
// path, or terminate test execution.
func resolveMetroctl() string {
	path, err := datafile.ResolveRunfile("metropolis/cli/metroctl/metroctl_/metroctl")
	if err != nil {
		log.Fatalf("Couldn't resolve metroctl binary: %v", err)
	}
	return path
}

// mctlRun starts metroctl, and waits till it exits. It returns nil, if the run
// was successful.
func mctlRun(t *testing.T, ctx context.Context, args []string) error {
	t.Helper()

	path := resolveMetroctl()
	log.Printf("$ metroctl %s", strings.Join(args, " "))
	logf := func(line string) {
		log.Printf("metroctl: %s", line)
	}
	_, err := cmd.RunCommand(ctx, path, args, cmd.WaitUntilCompletion(logf))
	return err
}

// mctlExpectOutput returns true in the event the expected string is found in
// metroctl output, and false otherwise.
func mctlExpectOutput(t *testing.T, ctx context.Context, args []string, expect string) (bool, error) {
	t.Helper()

	path := resolveMetroctl()
	log.Printf("$ metroctl %s", strings.Join(args, " "))
	// Terminate metroctl as soon as the expected output is found.
	logf := func(line string) {
		log.Printf("metroctl: %s", line)
	}
	found, err := cmd.RunCommand(ctx, path, args, cmd.TerminateIfFound(expect, logf))
	if err != nil {
		return false, fmt.Errorf("while running metroctl: %v", err)
	}
	return found, nil
}

// mctlFailIfMissing will return a non-nil error value either if the expected
// output string s is missing in metroctl output, or in case metroctl can't be
// launched.
func mctlFailIfMissing(t *testing.T, ctx context.Context, args []string, s string) error {
	found, err := mctlExpectOutput(t, ctx, args, s)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("expected output is missing: \"%s\"", s)
	}
	return nil
}

// mctlFailIfFound will return a non-nil error value either if the expected
// output string s is found in metroctl output, or in case metroctl can't be
// launched.
func mctlFailIfFound(t *testing.T, ctx context.Context, args []string, s string) error {
	found, err := mctlExpectOutput(t, ctx, args, s)
	if err != nil {
		return err
	}
	if found {
		return fmt.Errorf("unexpected output was found: \"%s\"", s)
	}
	return nil
}

func TestMetroctl(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	co := cluster.ClusterOptions{
		NumNodes: 2,
	}
	cl, err := cluster.LaunchCluster(context.Background(), co)
	if err != nil {
		t.Fatalf("LaunchCluster failed: %v", err)
	}
	defer func() {
		err := cl.Close()
		if err != nil {
			t.Fatalf("cluster Close failed: %v", err)
		}
	}()

	socksRemote := fmt.Sprintf("localhost:%d", cl.Ports[cluster.SOCKSPort])
	var clusterEndpoints []string
	for _, ep := range cl.Nodes {
		clusterEndpoints = append(clusterEndpoints, ep.ManagementAddress)
	}

	ownerPem := pem.EncodeToMemory(&pem.Block{
		Type:  "METROPOLIS INITIAL OWNER PRIVATE KEY",
		Bytes: cluster.InsecurePrivateKey,
	})
	if err := os.WriteFile("owner-key.pem", ownerPem, 0644); err != nil {
		log.Fatal("Couldn't write owner-key.pem")
	}

	commonOpts := []string{
		"--proxy=" + socksRemote,
		"--config=.",
	}

	var endpointOpts []string
	for _, ep := range clusterEndpoints {
		endpointOpts = append(endpointOpts, "--endpoints="+ep)
	}

	log.Printf("metroctl: Cluster's running, starting tests...")
	st := t.Run("Init", func(t *testing.T) {
		util.TestEventual(t, "metroctl takeownership", ctx, 30*time.Second, func(ctx context.Context) error {
			// takeownership needs just a single endpoint pointing at the initial node.
			var args []string
			args = append(args, commonOpts...)
			args = append(args, endpointOpts[0])
			args = append(args, "takeownership")
			return mctlFailIfMissing(t, ctx, args, "Successfully retrieved owner credentials")
		})
	})
	if !st {
		t.Fatalf("metroctl: Couldn't get cluster ownership.")
	}
	t.Run("list", func(t *testing.T) {
		util.TestEventual(t, "metroctl list", ctx, 10*time.Second, func(ctx context.Context) error {
			var args []string
			args = append(args, commonOpts...)
			args = append(args, endpointOpts...)
			args = append(args, "node", "list")
			// Expect both node IDs to show up in the results.
			if err := mctlFailIfMissing(t, ctx, args, cl.NodeIDs[0]); err != nil {
				return err
			}
			return mctlFailIfMissing(t, ctx, args, cl.NodeIDs[1])
		})
	})
	t.Run("list [nodeID]", func(t *testing.T) {
		util.TestEventual(t, "metroctl list [nodeID]", ctx, 10*time.Second, func(ctx context.Context) error {
			var args []string
			args = append(args, commonOpts...)
			args = append(args, endpointOpts...)
			args = append(args, "node", "list", cl.NodeIDs[1])
			// Expect just the supplied node IDs to show up in the results.
			if err := mctlFailIfFound(t, ctx, args, cl.NodeIDs[0]); err != nil {
				return err
			}
			return mctlFailIfMissing(t, ctx, args, cl.NodeIDs[1])
		})
	})
	t.Run("list --output", func(t *testing.T) {
		util.TestEventual(t, "metroctl list --output", ctx, 10*time.Second, func(ctx context.Context) error {
			var args []string
			args = append(args, commonOpts...)
			args = append(args, endpointOpts...)
			args = append(args, "node", "list", "--output", "list.txt")
			// In this case metroctl should write its output to a file rather than stdout.
			if err := mctlFailIfFound(t, ctx, args, cl.NodeIDs[0]); err != nil {
				return err
			}
			od, err := os.ReadFile("list.txt")
			if err != nil {
				return fmt.Errorf("while reading metroctl output file: %v", err)
			}
			if !strings.Contains(string(od), cl.NodeIDs[0]) {
				return fmt.Errorf("expected node ID hasn't been found in metroctl output")
			}
			return nil
		})
	})
	t.Run("list --filter", func(t *testing.T) {
		util.TestEventual(t, "metroctl list --filter", ctx, 10*time.Second, func(ctx context.Context) error {
			nid := cl.NodeIDs[1]
			naddr := cl.Nodes[nid].ManagementAddress

			var args []string
			args = append(args, commonOpts...)
			args = append(args, endpointOpts...)
			// Filter list results based on nodes' external addresses.
			args = append(args, "node", "list", "--filter", fmt.Sprintf("node.status.external_address==\"%s\"", naddr))
			// Expect the second node's ID to show up in the results.
			if err := mctlFailIfMissing(t, ctx, args, cl.NodeIDs[1]); err != nil {
				return err
			}
			// The first node should've been filtered away.
			return mctlFailIfFound(t, ctx, args, cl.NodeIDs[0])
		})
	})
	t.Run("describe --filter", func(t *testing.T) {
		util.TestEventual(t, "metroctl list --filter", ctx, 10*time.Second, func(ctx context.Context) error {
			nid := cl.NodeIDs[0]
			naddr := cl.Nodes[nid].ManagementAddress

			var args []string
			args = append(args, commonOpts...)
			args = append(args, endpointOpts...)

			// Filter out the first node. Afterwards, only one node should be left.
			args = append(args, "node", "describe", "--output", "describe.txt", "--filter", fmt.Sprintf("node.status.external_address==\"%s\"", naddr))
			if err := mctlRun(t, ctx, args); err != nil {
				return err
			}

			// Try matching metroctl output against the advertised format.
			f, err := os.Open("describe.txt")
			if err != nil {
				return fmt.Errorf("while opening metroctl output: %v", err)
			}
			scanner := bufio.NewScanner(f)
			if !scanner.Scan() {
				return fmt.Errorf("expected header line")
			}
			if !scanner.Scan() {
				return fmt.Errorf("expected result line")
			}
			line := scanner.Text()
			t.Logf("Line: %q", line)

			var onid, ostate, onaddr, onstatus, onroles string
			var ontimeout int

			_, err = fmt.Sscanf(line, "%s%s%s%s%s%ds", &onid, &ostate, &onaddr, &onstatus, &onroles, &ontimeout)
			if err != nil {
				return fmt.Errorf("while parsing metroctl output: %v", err)
			}
			if onid != nid {
				return fmt.Errorf("node id mismatch")
			}
			if ostate != "UP" {
				return fmt.Errorf("node state mismatch")
			}
			if onaddr != naddr {
				return fmt.Errorf("node address mismatch")
			}
			if onstatus != "HEALTHY" {
				return fmt.Errorf("node status mismatch")
			}
			if want, got := "ConsensusMember,KubernetesController", onroles; want != got {
				return fmt.Errorf("node role mismatch: wanted %q, got %q", want, got)
			}
			if ontimeout < 0 || ontimeout > 30 {
				return fmt.Errorf("node timeout mismatch")
			}
			return nil
		})
	})
	t.Run("set/unset role", func(t *testing.T) {
		util.TestEventual(t, "metroctl set/unset role KubernetesController", ctx, 10*time.Second, func(ctx context.Context) error {
			nid := cl.NodeIDs[1]
			naddr := cl.Nodes[nid].ManagementAddress

			// In this test we'll unset a node role, make sure that it's been in fact
			// unset, then set it again, and check again. This exercises commands of
			// the form "metroctl set/unset role KubernetesController [NodeID, ...]".

			// Check that KubernetesController role is set initially.
			var describeArgs []string
			describeArgs = append(describeArgs, commonOpts...)
			describeArgs = append(describeArgs, endpointOpts...)
			describeArgs = append(describeArgs, "node", "describe", "--filter", fmt.Sprintf("node.status.external_address==\"%s\"", naddr))
			if err := mctlFailIfMissing(t, ctx, describeArgs, "KubernetesController"); err != nil {
				return err
			}
			// Remove the role.
			var unsetArgs []string
			unsetArgs = append(unsetArgs, commonOpts...)
			unsetArgs = append(unsetArgs, endpointOpts...)
			unsetArgs = append(unsetArgs, "node", "remove", "role", "KubernetesController", nid)
			if err := mctlRun(t, ctx, unsetArgs); err != nil {
				return err
			}
			// Check that the role is unset.
			if err := mctlFailIfFound(t, ctx, describeArgs, "KubernetesController"); err != nil {
				return err
			}
			// Set the role back to the initial value.
			var setArgs []string
			setArgs = append(setArgs, commonOpts...)
			setArgs = append(setArgs, endpointOpts...)
			setArgs = append(setArgs, "node", "add", "role", "KubernetesController", nid)
			if err := mctlRun(t, ctx, setArgs); err != nil {
				return err
			}
			// Chack that the role is set.
			return mctlFailIfMissing(t, ctx, describeArgs, "KubernetesController")
		})
	})
}
