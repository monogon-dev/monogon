// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package roleserve

import (
	"context"

	"source.monogon.dev/metropolis/node/core/mgmt"
	"source.monogon.dev/metropolis/node/core/update"
	"source.monogon.dev/osbase/event/memory"
	"source.monogon.dev/osbase/logtree"
	"source.monogon.dev/osbase/supervisor"
)

type workerNodeMgmt struct {
	curatorConnection *memory.Value[*CuratorConnection]
	logTree           *logtree.LogTree
	updateService     *update.Service
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
		NodeCredentials: cc.Credentials,
		LogTree:         s.logTree,
		UpdateService:   s.updateService,
	}
	return srv.Run(ctx)
}
