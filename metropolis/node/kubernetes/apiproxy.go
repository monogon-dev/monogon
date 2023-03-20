package kubernetes

import (
	"context"
	"fmt"
	"net"

	"source.monogon.dev/go/net/tinylb"
	"source.monogon.dev/metropolis/node"
	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/pkg/event/memory"
)

// updateLoadBalancerAPIServers provides a tinylb BackendSet memory value with
// the currently known nodes running a Kubernetes apiserver as retrieved from the
// given curator client.
func updateLoadbalancerAPIServers(ctx context.Context, val *memory.Value[tinylb.BackendSet], cur ipb.CuratorClient) error {
	w, err := cur.Watch(ctx, &ipb.WatchRequest{
		Kind: &ipb.WatchRequest_NodesInCluster_{
			NodesInCluster: &ipb.WatchRequest_NodesInCluster{},
		},
	})
	if err != nil {
		return fmt.Errorf("watch failed: %w", err)
	}
	defer w.CloseSend()

	set := &tinylb.BackendSet{}
	val.Set(set.Clone())
	for {
		ev, err := w.Recv()
		if err != nil {
			return fmt.Errorf("receive failed: %w", err)
		}

		for _, n := range ev.Nodes {
			if n.Status == nil || n.Status.ExternalAddress == "" {
				set.Delete(n.Id)
				continue
			}
			if n.Roles.KubernetesController == nil {
				set.Delete(n.Id)
				continue
			}
			set.Insert(n.Id, &tinylb.SimpleTCPBackend{Remote: net.JoinHostPort(n.Status.ExternalAddress, node.KubernetesAPIPort.PortString())})
		}
		for _, t := range ev.NodeTombstones {
			set.Delete(t.NodeId)
		}
		val.Set(set.Clone())
	}
}
