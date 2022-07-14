package test

import (
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

func expectMetroctl(t *testing.T, ctx context.Context, args []string, expect string) error {
	t.Helper()

	path, err := datafile.ResolveRunfile("metropolis/cli/metroctl/metroctl_/metroctl")
	if err != nil {
		return fmt.Errorf("couldn't resolve metroctl binary: %v", err)
	}

	log.Printf("$ metroctl %s", strings.Join(args, " "))
	found, err := cmd.RunCommand(ctx, path, args, expect)
	if err != nil {
		return fmt.Errorf("while running metroctl: %v", err)
	}
	if !found {
		return fmt.Errorf("expected string wasn't found while running metroctl.")
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
			return expectMetroctl(t, ctx, args, "Successfully retrieved owner credentials")
		})
	})
	if !st {
		t.Fatalf("metroctl: Couldn't get cluster ownership.")
	}
}
