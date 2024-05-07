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
	"source.monogon.dev/osbase/test/launch"
)

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

	rPath, err := runfiles.Rlocation("_main/metropolis/test/e2e/testimages_manifest.prototxt")
	if err != nil {
		t.Fatalf("Resolving registry manifest failed: %v", err)
	}
	df, err := os.ReadFile(rPath)
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

	launch.Log("E2E: Cluster running, starting tests...")

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
