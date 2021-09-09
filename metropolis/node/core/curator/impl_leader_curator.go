package curator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/pkg/event/etcd"
)

// leaderCurator implements the Curator gRPC API (cpb.Curator) as a curator
// leader.
type leaderCurator struct {
	leadership
}

// Watch returns a stream of updates concerning some part of the cluster
// managed by the curator.
//
// See metropolis.node.core.curator.proto.api.Curator for more information.
func (l *leaderCurator) Watch(req *cpb.WatchRequest, srv cpb.Curator_WatchServer) error {
	nic, ok := req.Kind.(*cpb.WatchRequest_NodeInCluster_)
	if !ok {
		return status.Error(codes.Unimplemented, "unsupported watch kind")
	}
	nodeID := nic.NodeInCluster.NodeId
	// Constructing arbitrary etcd path: this is okay, as we only have node objects
	// underneath the NodeEtcdPrefix. Worst case an attacker can do is request a node
	// that doesn't exist, and that will just hang . All access is privileged, so
	// there's also no need to filter anything.
	// TODO(q3k): formalize and strongly type etcd paths for cluster state?
	// Probably worth waiting for type parameters before attempting to do that.
	nodePath, err := nodeEtcdPrefix.Key(nodeID)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid node name: %v", err)
	}

	value := etcd.NewValue(l.etcd, nodePath, func(data []byte) (interface{}, error) {
		return nodeUnmarshal(data)
	})
	w := value.Watch()
	defer w.Close()

	for {
		v, err := w.Get(srv.Context())
		if err != nil {
			if rpcErr, ok := rpcError(err); ok {
				return rpcErr
			}
		}
		node := v.(*Node)
		ev := &cpb.WatchEvent{
			Nodes: []*cpb.Node{
				{
					Id:    node.ID(),
					Roles: node.proto().Roles,
				},
			},
		}
		if err := srv.Send(ev); err != nil {
			return err
		}
	}
}
