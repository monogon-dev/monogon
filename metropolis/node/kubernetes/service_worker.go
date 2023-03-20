package kubernetes

import (
	"context"
	"fmt"
	"net"

	"source.monogon.dev/go/net/tinylb"
	"source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

type ConfigWorker struct {
	ServiceIPRange net.IPNet
	ClusterNet     net.IPNet
	ClusterDomain  string

	Root          *localstorage.Root
	Network       *network.Service
	NodeID        string
	CuratorClient ipb.CuratorClient
}

type Worker struct {
	c ConfigWorker
}

func NewWorker(c ConfigWorker) *Worker {
	s := &Worker{
		c: c,
	}
	return s
}

func (s *Worker) Run(ctx context.Context) error {
	// Run apiproxy, which load-balances connections from worker components to this
	// cluster's api servers. This is necessary as we want to round-robin across all
	// available apiservers, and Kubernetes components do not implement client-side
	// load-balancing.
	err := supervisor.Run(ctx, "apiproxy", func(ctx context.Context) error {
		lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", node.KubernetesWorkerLocalAPIPort))
		if err != nil {
			return fmt.Errorf("failed to listen: %w", err)
		}
		defer lis.Close()

		v := memory.Value[tinylb.BackendSet]{}
		srv := tinylb.Server{
			Provider: &v,
			Listener: lis,
		}
		err = supervisor.Run(ctx, "updater", func(ctx context.Context) error {
			return updateLoadbalancerAPIServers(ctx, &v, s.c.CuratorClient)
		})
		if err != nil {
			return err
		}

		supervisor.Logger(ctx).Infof("Starting proxy...")
		return srv.Run(ctx)
	})
	if err != nil {
		return err
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	return nil
}
