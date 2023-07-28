package roleserve

import (
	"context"
	"fmt"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"

	"source.monogon.dev/metropolis/node/core/metrics"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// workerMetrics runs the Metrics Service, which runs local Prometheus collectors
// (themselves usually instances of existing Prometheus Exporters running as
// sub-processes), and a forwarding service that lets external users access them
// over HTTPS using the Cluster CA.
type workerMetrics struct {
	curatorConnection *memory.Value[*curatorConnection]
	localRoles        *memory.Value[*cpb.NodeRoles]
	localControlplane *memory.Value[*localControlPlane]
}

func (s *workerMetrics) run(ctx context.Context) error {
	w := s.curatorConnection.Watch()
	defer w.Close()

	supervisor.Logger(ctx).Infof("Waiting for curator connection")
	cc, err := w.Get(ctx)
	if err != nil {
		return err
	}
	supervisor.Logger(ctx).Infof("Got curator connection, starting...")

	svc := metrics.Service{
		Credentials: cc.credentials,
		Discovery: metrics.Discovery{
			Curator: ipb.NewCuratorClient(cc.conn),
		},
	}

	err = supervisor.Run(ctx, "watch-consensus", func(ctx context.Context) error {
		isConsensusMember := func(roles *cpb.NodeRoles) bool {
			return roles.ConsensusMember != nil
		}

		w := s.localRoles.Watch()
		defer w.Close()

		r, err := w.Get(ctx)
		if err != nil {
			return err
		}

		if isConsensusMember(r) {
			if err := supervisor.Run(ctx, "discovery", svc.Discovery.Run); err != nil {
				return err
			}
		}

		for {
			nr, err := w.Get(ctx)
			if err != nil {
				return err
			}

			changed := isConsensusMember(r) != isConsensusMember(nr)
			if changed {
				return fmt.Errorf("restarting")
			}
		}
	})
	if err != nil {
		return err
	}

	return svc.Run(ctx)
}
