package curator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

type curatorFollower struct {
}

func (f *curatorFollower) Watch(req *apb.WatchRequest, srv apb.Curator_WatchServer) error {
	return status.Error(codes.Unimplemented, "curator follower not implemented")
}
