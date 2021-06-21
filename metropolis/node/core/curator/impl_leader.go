package curator

import (
	"context"
	"errors"
	"fmt"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/pkg/event/etcd"
)

// curatorLeader implements the curator acting as the elected leader of a
// cluster. It performs direct reads/writes from/to etcd as long as it remains
// leader.
//
// It effectively implements all the core management logic of a Metropolis
// cluster.
type curatorLeader struct {
	// lockKey is the etcd key which backs this leader-elected instance.
	lockKey string
	// lockRev is the revision at which lockKey was created. The leader will use it
	// in combination with lockKey to ensure all mutations/reads performed to etcd
	// succeed only if this leader election is still current.
	lockRev int64
	// etcd is the etcd client in which curator data and leader election state is
	// stored.
	etcd client.Namespaced
}

var (
	// lostLeadership is returned by txnAsLeader if the transaction got canceled
	// because leadership was lost.
	lostLeadership = errors.New("lost leadership")
)

// txnAsLeader performs an etcd transaction guarded by continued leadership.
// lostLeadership will be returned as an error in case the leadership is lost.
func (c *curatorLeader) txnAsLeader(ctx context.Context, ops ...clientv3.Op) (*clientv3.TxnResponse, error) {
	resp, err := c.etcd.Txn(ctx).If(
		clientv3.Compare(clientv3.CreateRevision(c.lockKey), "=", c.lockRev),
	).Then(ops...).Commit()
	if err != nil {
		return nil, fmt.Errorf("when running leader transaction: %w", err)
	}
	if !resp.Succeeded {
		return nil, lostLeadership
	}
	return resp, nil
}

// rpcError attempts to convert a given error to a high-level error that can be
// directly exposed to RPC clients. If false is returned, the error was not
// converted and is returned verbatim.
func rpcError(err error) (error, bool) {
	if errors.Is(err, lostLeadership) {
		return status.Error(codes.Unavailable, "lost leadership"), true
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return status.Errorf(codes.Unavailable, "%v", err), true
	}
	return err, false
}

func (l *curatorLeader) Watch(req *apb.WatchRequest, srv apb.Curator_WatchServer) error {
	nic, ok := req.Kind.(*apb.WatchRequest_NodeInCluster_)
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
	nodePath := nodeEtcdPath(nodeID)

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
		ev := &apb.WatchEvent{
			Nodes: []*apb.Node{
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
