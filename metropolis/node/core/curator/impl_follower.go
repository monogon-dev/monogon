package curator

import (
	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	apb "source.monogon.dev/metropolis/proto/api"
)

type curatorFollower struct {
	apb.UnimplementedAAAServer
	apb.UnimplementedManagementServer
	cpb.UnimplementedCuratorServer
}
