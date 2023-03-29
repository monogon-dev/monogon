package roleserve

import (
	"context"

	"source.monogon.dev/metropolis/node/core/mgmt"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

type workerNodeMgmt struct {
	clusterMembership *memory.Value[*ClusterMembership]
}

func (s *workerNodeMgmt) run(ctx context.Context) error {
	w := s.clusterMembership.Watch()
	defer w.Close()
	supervisor.Logger(ctx).Infof("Waiting for cluster membership...")
	cm, err := w.Get(ctx, FilterHome())
	if err != nil {
		return err
	}

	supervisor.Logger(ctx).Infof("Got cluster membership, starting...")
	srv := mgmt.Service{
		NodeCredentials: cm.credentials,
	}
	return srv.Run(ctx)
}
