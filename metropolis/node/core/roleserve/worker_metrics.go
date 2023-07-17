package roleserve

import (
	"context"

	cpb "source.monogon.dev/metropolis/proto/common"

	"source.monogon.dev/metropolis/node/core/metrics"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

// workerMetrics runs the Metrics Service, which runs local Prometheus collectors
// (themselves usually instances of existing Prometheus Exporters running as
// sub-processes), and a forwarding service that lets external users access them
// over HTTPS using the Cluster CA.
type workerMetrics struct {
	curatorConnection *memory.Value[*curatorConnection]
	localRoles        *memory.Value[*cpb.NodeRoles]
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
		Curator:     ipb.NewCuratorClient(cc.conn),
		LocalRoles:  s.localRoles,
	}
	return svc.Run(ctx)
}
