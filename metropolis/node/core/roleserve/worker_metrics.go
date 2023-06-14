package roleserve

import (
	"context"

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
	}
	return svc.Run(ctx)
}
