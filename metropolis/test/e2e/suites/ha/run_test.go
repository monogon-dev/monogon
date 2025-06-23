// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package ha

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"

	mlaunch "source.monogon.dev/metropolis/test/launch"
	"source.monogon.dev/metropolis/test/localregistry"
	"source.monogon.dev/metropolis/test/util"
)

var (
	// These are filled by bazel at linking time with the canonical path of
	// their corresponding file. Inside the init function we resolve it
	// with the rules_go runfiles package to the real path.
	xTestImagesManifestPath string
)

func init() {
	var err error
	for _, path := range []*string{
		&xTestImagesManifestPath,
	} {
		*path, err = runfiles.Rlocation(*path)
		if err != nil {
			panic(err)
		}
	}
}

const (
	// Timeout for the global test context.
	//
	// Bazel would eventually time out the test after 900s ("large") if, for
	// some reason, the context cancellation fails to abort it.
	globalTestTimeout = 600 * time.Second

	// Timeouts for individual end-to-end tests of different sizes.
	smallTestTimeout = 60 * time.Second
	largeTestTimeout = 120 * time.Second
)

// TestE2ECoreHA exercises the basics of a high-availability control plane by
// starting up a 3-node cluster, turning all nodes into ConsensusMembers, then
// performing a rolling restart.
func TestE2ECoreHA(t *testing.T) {
	// Set a global timeout to make sure this terminates
	ctx, cancel := context.WithTimeout(context.Background(), globalTestTimeout)
	defer cancel()

	df, err := os.ReadFile(xTestImagesManifestPath)
	if err != nil {
		t.Fatalf("Reading registry manifest failed: %v", err)
	}
	lr, err := localregistry.FromBazelManifest(df)
	if err != nil {
		t.Fatalf("Creating test image registry failed: %v", err)
	}
	// Launch cluster.
	clusterOptions := mlaunch.ClusterOptions{
		NumNodes:        3,
		LocalRegistry:   lr,
		NodeLogsToFiles: true,
		Node: mlaunch.NodeOptions{
			// ESP, 2 system partitions, and data partition.
			DiskBytes: (128 + 2*1024 + 512) * 1024 * 1024,
		},
	}
	cluster, err := mlaunch.LaunchCluster(ctx, clusterOptions)
	if err != nil {
		t.Fatalf("LaunchCluster failed: %v", err)
	}
	defer func() {
		err := cluster.Close()
		if err != nil {
			t.Fatalf("cluster Close failed: %v", err)
		}
	}()

	util.MustTestEventual(t, "Add ConsensusMember roles", ctx, smallTestTimeout, func(ctx context.Context) error {
		// Make everything but the first node into ConsensusMember.
		for i := 1; i < clusterOptions.NumNodes; i++ {
			err := cluster.MakeConsensusMember(ctx, cluster.NodeIDs[i])
			if err != nil {
				return fmt.Errorf("MakeConsensusMember(%d/%s): %w", i, cluster.NodeIDs[i], err)
			}
		}
		return nil
	})
	util.TestEventual(t, "Heartbeat test successful", ctx, 20*time.Second, cluster.AllNodesHealthy)

	// Perform rolling restart of all nodes. When a node rejoins it must be able to
	// contact the cluster, so this also exercises that the cluster is serving even
	// with the node having rebooted.
	for i := 0; i < clusterOptions.NumNodes; i++ {
		util.MustTestEventual(t, fmt.Sprintf("Node %d rejoin successful", i), ctx, 60*time.Second, func(ctx context.Context) error {
			// Ensure nodes rejoin the cluster after a reboot by reboting the 1st node.
			if err := cluster.RebootNode(ctx, i); err != nil {
				return fmt.Errorf("while rebooting a node: %w", err)
			}
			return nil
		})
	}
}
