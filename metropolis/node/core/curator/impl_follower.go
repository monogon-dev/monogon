package curator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/consensus/client"
	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	ppb "source.monogon.dev/metropolis/node/core/curator/proto/private"
	"source.monogon.dev/metropolis/node/core/rpc"
)

type curatorFollower struct {
	lock       *ppb.LeaderElectionValue
	etcd       client.Namespaced
	followerID string
}

func (f *curatorFollower) GetCurrentLeader(_ *cpb.GetCurrentLeaderRequest, srv cpb.CuratorLocal_GetCurrentLeaderServer) error {
	ctx := srv.Context()
	if f.lock == nil {
		return status.Errorf(codes.Unavailable, "could not determine current leader")
	}
	nodeId := f.lock.NodeId

	// Manually load node status data from etcd, even though we are not a leader.
	// This is fine, as if we ever end up serving stale data, the client will
	// realize and call us again.
	key, err := nodeEtcdPrefix.Key(nodeId)
	if err != nil {
		rpc.Trace(ctx).Printf("invalid leader node id %q: %v", nodeId, err)
		return status.Errorf(codes.Internal, "current leader has invalid node id")
	}
	res, err := f.etcd.Get(ctx, key)
	if err != nil {
		rpc.Trace(ctx).Printf("Get(%q) failed: %v", key, err)
		return status.Errorf(codes.Unavailable, "could not retrieve leader node from etcd")
	}
	if len(res.Kvs) != 1 {
		rpc.Trace(ctx).Printf("Get(%q) returned %d nodes", key, len(res.Kvs))
		return status.Errorf(codes.Internal, "current leader not found in etcd")
	}
	node, err := nodeUnmarshal(res.Kvs[0].Value)
	if err != nil {
		rpc.Trace(ctx).Printf("could not unmarshal leader node %s: %v", nodeId, err)
		return status.Errorf(codes.Unavailable, "could not unmarshal leader node")
	}
	if node.status == nil || node.status.ExternalAddress == "" {
		rpc.Trace(ctx).Printf("leader node %s has no address in status", nodeId)
		return status.Errorf(codes.Unavailable, "current leader has no reported address")
	}

	err = srv.Send(&cpb.GetCurrentLeaderResponse{
		LeaderNodeId: nodeId,
		LeaderHost:   node.status.ExternalAddress,
		LeaderPort:   int32(common.CuratorServicePort),
		ThisNodeId:   f.followerID,
	})
	if err != nil {
		return err
	}

	<-ctx.Done()
	rpc.Trace(ctx).Printf("Interrupting due to context cancellation")
	return nil
}
