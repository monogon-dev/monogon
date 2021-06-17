package curator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

type curatorLeader struct {
	lockKey string
	lockRev int64
	etcd    client.Namespaced
}

func (l *curatorLeader) Watch(req *apb.WatchRequest, srv apb.Curator_WatchServer) error {
	return status.Error(codes.Unimplemented, "curator leader not implemented")
}
