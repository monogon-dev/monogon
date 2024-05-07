package ha_cold

import (
	"context"
	"fmt"
	"testing"
	"time"

	mlaunch "source.monogon.dev/metropolis/test/launch"
	"source.monogon.dev/metropolis/test/util"
	"source.monogon.dev/osbase/test/launch"

	cpb "source.monogon.dev/metropolis/proto/common"
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

// TestE2EColdStartHA exercises an HA cluster being fully shut down then
// restarted again.
//
// Metropolis currently doesn't support cold startups from TPM/Secure clusters,
// so we test a non-TPM/Insecure cluster.
func TestE2EColdStartHA(t *testing.T) {
	// Set a global timeout to make sure this terminates
	ctx, cancel := context.WithTimeout(context.Background(), globalTestTimeout)
	defer cancel()

	// Launch cluster.
	clusterOptions := mlaunch.ClusterOptions{
		NumNodes:        3,
		NodeLogsToFiles: true,
		InitialClusterConfiguration: &cpb.ClusterConfiguration{
			TpmMode:               cpb.ClusterConfiguration_TPM_MODE_DISABLED,
			StorageSecurityPolicy: cpb.ClusterConfiguration_STORAGE_SECURITY_POLICY_NEEDS_INSECURE,
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

	launch.Log("E2E: Cluster running, starting tests...")

	util.MustTestEventual(t, "Add ConsensusMember roles", ctx, smallTestTimeout, func(ctx context.Context) error {
		// Make everything but the first node into ConsensusMember.
		for i := 1; i < clusterOptions.NumNodes; i++ {
			err := cluster.MakeConsensusMember(ctx, cluster.NodeIDs[i])
			if err != nil {
				return util.Permanent(fmt.Errorf("MakeConsensusMember(%d/%s): %w", i, cluster.NodeIDs[i], err))
			}
		}
		return nil
	})
	util.TestEventual(t, "Heartbeat test successful", ctx, 20*time.Second, cluster.AllNodesHealthy)

	// Shut every node down.
	for i := 0; i < clusterOptions.NumNodes; i++ {
		if err := cluster.ShutdownNode(i); err != nil {
			t.Fatalf("Could not shutdown node %d", i)
		}
	}
	// Start every node back up.
	for i := 0; i < clusterOptions.NumNodes; i++ {
		if err := cluster.StartNode(i); err != nil {
			t.Fatalf("Could not shutdown node %d", i)
		}
	}
	// Check if the cluster comes back up.
	util.TestEventual(t, "Heartbeat test successful", ctx, 60*time.Second, cluster.AllNodesHealthy)
}
