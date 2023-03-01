package roleserve

import (
	"context"
	"fmt"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// workerRoleFetch is the Role Fetcher, an internal bookkeeping service
// responsible for populating localRoles based on a clusterMembership whenever
// the node is HOME and cluster credentials / curator access is available.
type workerRoleFetch struct {
	clusterMembership *memory.Value[*ClusterMembership]

	// localRoles will be written.
	localRoles *memory.Value[*cpb.NodeRoles]
}

func (s *workerRoleFetch) run(ctx context.Context) error {
	w := s.clusterMembership.Watch()
	defer w.Close()
	supervisor.Logger(ctx).Infof("Waiting for cluster membership...")
	cm, err := w.Get(ctx, FilterHome())
	if err != nil {
		return err
	}
	supervisor.Logger(ctx).Infof("Got cluster membership, starting...")

	nodeID := cm.NodeID()
	conn, err := cm.DialCurator()
	if err != nil {
		return err
	}
	defer conn.Close()
	cur := ipb.NewCuratorClient(conn)

	// Start watch for current node, update localRoles whenever we get something
	// new.
	srv, err := cur.Watch(ctx, &ipb.WatchRequest{Kind: &ipb.WatchRequest_NodeInCluster_{
		NodeInCluster: &ipb.WatchRequest_NodeInCluster{
			NodeId: nodeID,
		},
	}})
	if err != nil {
		return fmt.Errorf("watch failed: %w", err)
	}
	defer srv.CloseSend()

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	for {
		ev, err := srv.Recv()
		if err != nil {
			return fmt.Errorf("watch event receive failed: %w", err)
		}
		for _, n := range ev.Nodes {
			n := n
			// Skip spuriously sent other nodes.
			if n.Id != nodeID {
				continue
			}
			supervisor.Logger(ctx).Infof("Got new node data. Roles:")
			if n.Roles.ConsensusMember != nil {
				supervisor.Logger(ctx).Infof(" - control plane member, existing peers: %+v", n.Roles.ConsensusMember.Peers)
			}
			if n.Roles.KubernetesController != nil {
				supervisor.Logger(ctx).Infof(" - kubernetes controller")
			}
			if n.Roles.KubernetesWorker != nil {
				supervisor.Logger(ctx).Infof(" - kubernetes worker")
			}
			s.localRoles.Set(n.Roles)
			break
		}
	}
}
