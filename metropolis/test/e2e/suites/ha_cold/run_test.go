package ha_cold

import (
	"context"
	"fmt"
	"testing"
	"time"

	mlaunch "source.monogon.dev/metropolis/test/launch"
	"source.monogon.dev/metropolis/test/util"
	"source.monogon.dev/osbase/test/launch"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	apb "source.monogon.dev/metropolis/proto/api"
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
			ClusterDomain:         "cluster.test",
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

	// Test node role removal.
	curC, err := cluster.CuratorClient()
	if err != nil {
		t.Fatalf("Could not get CuratorClient: %v", err)
	}
	mgmt := apb.NewManagementClient(curC)
	cur := ipb.NewCuratorClient(curC)

	util.MustTestEventual(t, "Remove KubernetesController role", ctx, 10*time.Second, func(ctx context.Context) error {
		fa := false
		_, err := mgmt.UpdateNodeRoles(ctx, &apb.UpdateNodeRolesRequest{
			Node: &apb.UpdateNodeRolesRequest_Id{
				Id: cluster.NodeIDs[0],
			},
			KubernetesController: &fa,
		})
		return err
	})
	util.MustTestEventual(t, "Remove ConsensusMember role", ctx, time.Minute, func(ctx context.Context) error {
		fa := false
		_, err := mgmt.UpdateNodeRoles(ctx, &apb.UpdateNodeRolesRequest{
			Node: &apb.UpdateNodeRolesRequest_Id{
				Id: cluster.NodeIDs[0],
			},
			ConsensusMember: &fa,
		})
		return err
	})

	// Test that removing the ConsensusMember role from a node removed the
	// corresponding etcd member from the cluster.
	var st *ipb.GetConsensusStatusResponse
	util.MustTestEventual(t, "Get ConsensusStatus", ctx, time.Minute, func(ctx context.Context) error {
		st, err = cur.GetConsensusStatus(ctx, &ipb.GetConsensusStatusRequest{})
		return err
	})

	for _, member := range st.EtcdMember {
		if member.Id == cluster.NodeIDs[0] {
			t.Errorf("member still present in etcd")
		}
	}

	// Test that that the cluster still works after deleting the first node and
	// restarting the remaining nodes.
	util.MustTestEventual(t, "Delete first node", ctx, 10*time.Second, func(ctx context.Context) error {
		_, err := mgmt.DeleteNode(ctx, &apb.DeleteNodeRequest{
			Node: &apb.DeleteNodeRequest_Id{
				Id: cluster.NodeIDs[0],
			},
			SafetyBypassNotDecommissioned: &apb.DeleteNodeRequest_SafetyBypassNotDecommissioned{},
		})
		return err
	})

	// Shut every remaining node down.
	for i := 1; i < clusterOptions.NumNodes; i++ {
		if err := cluster.ShutdownNode(i); err != nil {
			t.Fatalf("Could not shutdown node %d", i)
		}
	}
	// Start every remaining node back up.
	for i := 1; i < clusterOptions.NumNodes; i++ {
		if err := cluster.StartNode(i); err != nil {
			t.Fatalf("Could not shutdown node %d", i)
		}
	}
	// Check if the cluster comes back up.
	util.TestEventual(t, "Heartbeat test successful", ctx, 60*time.Second, cluster.AllNodesHealthy)
}
