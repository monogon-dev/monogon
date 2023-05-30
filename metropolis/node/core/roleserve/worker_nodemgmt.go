package roleserve

import (
	"context"

	"source.monogon.dev/metropolis/node/core/mgmt"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

type workerNodeMgmt struct {
	curatorConnection *memory.Value[*curatorConnection]
	logTree           *logtree.LogTree
}

func (s *workerNodeMgmt) run(ctx context.Context) error {
	w := s.curatorConnection.Watch()
	defer w.Close()
	supervisor.Logger(ctx).Infof("Waiting for cluster membership...")
	cc, err := w.Get(ctx)
	if err != nil {
		return err
	}

	supervisor.Logger(ctx).Infof("Got cluster membership, starting...")
	srv := mgmt.Service{
		NodeCredentials: cc.credentials,
		LogTree:         s.logTree,
	}
	return srv.Run(ctx)
}
