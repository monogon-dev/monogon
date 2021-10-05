package curator

import (
	"context"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/pkg/event/etcd"
)

// leaderCurator implements the Curator gRPC API (cpb.Curator) as a curator
// leader.
type leaderCurator struct {
	*leadership
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
			// TODO(q3k): log err
			return status.Error(codes.Unavailable, "internal error")
		}
		node := v.(*Node)
		ev := &cpb.WatchEvent{
			Nodes: []*cpb.Node{
				{
					Id:     node.ID(),
					Roles:  node.proto().Roles,
					Status: node.status,
				},
			},
		}
		if err := srv.Send(ev); err != nil {
			return err
		}
	}
}

// UpdateNodeStatus is called by nodes in the cluster to report their own
// status. This status is recorded by the curator and can be retrieed via
// Watch.
func (l *leaderCurator) UpdateNodeStatus(ctx context.Context, req *cpb.UpdateNodeStatusRequest) (*cpb.UpdateNodeStatusResponse, error) {
	// Ensure that the given node_id matches the calling node. We currently
	// only allow for direct self-reporting of status by nodes.
	pi := rpc.GetPeerInfo(ctx)
	if pi == nil || pi.Node == nil {
		return nil, status.Error(codes.PermissionDenied, "only nodes can update node status")
	}
	id := identity.NodeID(pi.Node.PublicKey)
	if id != req.NodeId {
		return nil, status.Errorf(codes.PermissionDenied, "node %q cannot update the status of node %q", id, req.NodeId)
	}

	// Verify sent status. Currently we assume the entire status must be set at
	// once, and cannot be unset.
	if req.Status == nil || req.Status.ExternalAddress == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Status and Status.ExternalAddress must be set")
	}

	// As we're performing a node update with two etcd transactions below (one
	// to retrieve, one to save and upate node), take a local lock to ensure
	// that we don't have a race between either two UpdateNodeStatus calls or
	// an UpdateNodeStatus call and some other mutation to the node store.
	l.muNodes.Lock()
	defer l.muNodes.Unlock()

	// Retrieve node ...
	key, err := nodeEtcdPrefix.Key(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid node id")
	}
	res, err := l.txnAsLeader(ctx, clientv3.OpGet(key))
	if err != nil {
		if rpcErr, ok := rpcError(err); ok {
			return nil, rpcErr
		}
		return nil, status.Errorf(codes.Unavailable, "could not retrieve node: %v", err)
	}
	kvs := res.Responses[0].GetResponseRange().Kvs
	if len(kvs) < 1 {
		return nil, status.Error(codes.NotFound, "no such node")
	}
	node, err := nodeUnmarshal(kvs[0].Value)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to unmarshal node: %v", err)
	}
	// ... update its' status ...
	node.status = req.Status
	// ... and save it to etcd.
	bytes, err := proto.Marshal(node.proto())
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "failed to marshal node: %v", err)
	}
	_, err = l.txnAsLeader(ctx, clientv3.OpPut(key, string(bytes)))
	if err != nil {
		if rpcErr, ok := rpcError(err); ok {
			return nil, rpcErr
		}
		return nil, status.Errorf(codes.Unavailable, "could not update node: %v", err)
	}

	return &cpb.UpdateNodeStatusResponse{}, nil
}
