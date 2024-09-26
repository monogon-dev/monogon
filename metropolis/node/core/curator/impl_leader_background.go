package curator

import (
	"context"
	"fmt"
	"time"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/osbase/supervisor"
)

// leaderBackground holds runnables which perform background processing on the
// curator leader.
type leaderBackground struct {
	*leadership
}

func (l *leaderBackground) background(ctx context.Context) error {
	if err := supervisor.Run(ctx, "sync-etcd", l.backgroundSyncEtcd); err != nil {
		return err
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	return ctx.Err()
}

// backgroundSyncEtcd ensures consistency between the set of nodes with
// ConsensusMember role, and the set of etcd members.
//
// When updating the ConsensusMember role, etcd membership is always added or
// removed first. Only after etcd accepts the change, the role is updated. If
// the role update fails, then roles and membership are inconsistent. To resolve
// the inconsistency, we take etcd membership as the source of truth, and update
// roles to match membership. That way, partially applied membership changes are
// made complete.
//
// Another way to end up in an inconsistent state is by deleting a node from the
// cluster which still has etcd membership, either by setting
// SafetyBypassHasRoles, or if the role has not been synced to match membership
// yet. In this case, we remove etcd membership.
func (l *leaderBackground) backgroundSyncEtcd(ctx context.Context) error {
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	for {
		// Process every 5 seconds.
		select {
		case <-time.After(5 * time.Second):
		case <-ctx.Done():
			return ctx.Err()
		}
		err := l.doSyncEtcd(ctx)
		if err != nil {
			return err
		}
	}
}

func (l *leaderBackground) doSyncEtcd(ctx context.Context) error {
	// Take muNodes to prevent concurrent role updates.
	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Get etcd members.
	members, err := l.consensusStatus.ClusterClient().MemberList(ctx)
	if err != nil {
		return fmt.Errorf("could not get etcd members: %w", err)
	}
	isEtcdMember := make(map[string]bool)
	// etcdMemberIDs is a map from etcd member ID (uint64) to node ID (string).
	etcdMemberIDs := make(map[uint64]string)
	for _, member := range members.Members {
		nodeID := consensus.GetEtcdMemberNodeId(member)
		isEtcdMember[nodeID] = true
		etcdMemberIDs[member.ID] = nodeID
	}

	// Get cluster nodes.
	res, err := l.txnAsLeader(ctx, NodeEtcdPrefix.Range())
	if err != nil {
		return fmt.Errorf("could not get nodes: %w", err)
	}
	isClusterNode := make(map[string]bool)
	var clusterNodes []*Node
	for _, kv := range res.Responses[0].GetResponseRange().Kvs {
		node, err := nodeUnmarshal(kv)
		if err != nil {
			return fmt.Errorf("could not unmarshal node %q: %w", kv.Key, err)
		}
		isClusterNode[node.ID()] = true
		clusterNodes = append(clusterNodes, node)
	}

	// Removing ConsensusMember roles is a potentially dangerous operation. If
	// something goes wrong and we fail to match an etcd member to the
	// corresponding cluster node, we could end up removing both the etcd
	// membership and the role. As a safety measure, refuse to perform any changes
	// if we do not find the local node ID in both etcd and cluster members.
	// Additionally, we first remove etcd members before removing roles, which
	// should prevent breaking consensus even if we fail to match cluster nodes to
	// corresponding etcd members.
	if !isEtcdMember[l.leaderID] {
		return fmt.Errorf("did not find local node ID in etcd members; refusing to do anything")
	}
	if !isClusterNode[l.leaderID] {
		return fmt.Errorf("did not find local node ID in cluster nodes; refusing to do anything")
	}

	// Remove etcd members that don't exist as nodes.
	//
	// Note that etcd membership operations are not guarded by curator leadership,
	// so we cannot assume that we are still leader. If we do not find a node,
	// this could be either because it has been deleted, or because it has been
	// created after we retrieved the list of nodes (assuming node IDs are never
	// reused). The second case would be bad, but it cannot occur because we get
	// the list of etcd members before the list of nodes.
	for memberID, nodeID := range etcdMemberIDs {
		if !isClusterNode[nodeID] {
			supervisor.Logger(ctx).Infof("Removing etcd member of non-existent node: %x / %s...", memberID, nodeID)
			_, err := l.consensusStatus.ClusterClient().MemberRemove(ctx, memberID)
			if err != nil {
				return fmt.Errorf("failed to remove etcd member: %w", err)
			}
			// Perform at most one change per call to doSyncEtcd.
			return nil
		}
	}

	// Update consensus roles to match etcd membership.
	for _, node := range clusterNodes {
		nodeID := node.ID()
		switch {
		case isEtcdMember[nodeID] && node.consensusMember == nil:
			supervisor.Logger(ctx).Infof("Adding ConsensusMember role to node which is etcd member: %s...", nodeID)
			// The node is already etcd member. We only call AddNode to obtain the
			// join parameters.
			join, err := l.consensusStatus.AddNode(ctx, nodeID, node.pubkey)
			if err != nil {
				return fmt.Errorf("failed to obtain consensus join parameters: %w", err)
			}
			node.EnableConsensusMember(join)
		case !isEtcdMember[nodeID] && node.consensusMember != nil:
			if node.kubernetesController != nil {
				// The KubernetesController role requires the ConsensusMember role,
				// so we need to remove it first.
				supervisor.Logger(ctx).Infof("Removing KubernetesController role from node which is not etcd member: %s...", nodeID)
				node.DisableKubernetesController()
			}
			supervisor.Logger(ctx).Infof("Removing ConsensusMember role from node which is not etcd member: %s...", nodeID)
			node.DisableConsensusMember()
		default:
			continue
		}
		if err := nodeSave(ctx, l.leadership, node); err != nil {
			return fmt.Errorf("could not save node with updated roles: %w", err)
		}
		return nil
	}

	return nil
}
