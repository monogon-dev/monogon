package roleserve

import (
	"context"
	"fmt"
	"net"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/kubernetes"
	"source.monogon.dev/metropolis/node/kubernetes/containerd"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// runKubernetesWorkerLauncher runs a launcher responsible for maintaining the
// main Kubernetes Worker service from //metropolis/node/kubernetes.
//
// TODO(q3k): make this generic as we have more workloads/launchers (and maybe
// Go type parameters).
func (s *Service) runKubernetesWorkerLauncher(ctx context.Context) error {
	var n *cpb.Node
	select {
	case <-ctx.Done():
		return ctx.Err()
	case n = <-s.kwC:
	}

	kw := n.Roles.KubernetesWorker
	if kw != nil {
		supervisor.Logger(ctx).Info("Node is a Kubernetes Worker, starting...")
		containerdSvc := &containerd.Service{
			EphemeralVolume: &s.StorageRoot.Ephemeral.Containerd,
		}
		if err := supervisor.Run(ctx, "containerd", containerdSvc.Run); err != nil {
			return fmt.Errorf("failed to start containerd service: %w", err)
		}

		kubeSvc := kubernetes.New(kubernetes.Config{
			// TODO(q3k): make this configurable.
			ServiceIPRange: net.IPNet{
				IP: net.IP{10, 0, 255, 1},
				// That's a /24.
				Mask: net.IPMask{0xff, 0xff, 0xff, 0x00},
			},
			ClusterNet: net.IPNet{
				IP: net.IP{10, 0, 0, 0},
				// That's a /16.
				Mask: net.IPMask{0xff, 0xff, 0x00, 0x00},
			},
			KPKI:    s.KPKI,
			Root:    s.StorageRoot,
			Network: s.Network,
			Node:    s.Node,
		})
		if err := supervisor.Run(ctx, "run", kubeSvc.Run); err != nil {
			return fmt.Errorf("failed to start kubernetes service: %w", err)
		}
		s.kwSvcC <- kubeSvc
	}
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case n := <-s.kwC:
			if kw == nil && n.Roles.KubernetesWorker != nil {
				return fmt.Errorf("node is now a kubernetes worker, restarting...")
			}
			if kw != nil && n.Roles.KubernetesWorker == nil {
				return fmt.Errorf("node is not a kubernetes worker anymore, restarting...")
			}
		}
	}
}
