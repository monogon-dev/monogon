package roleserve

import (
	"context"
	"fmt"
	"net"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/kubernetes"
	"source.monogon.dev/metropolis/node/kubernetes/containerd"
	kpki "source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// workerKubernetes is the Kubernetes Worker, responsible for launching
// (currently a control plane / data plane converged) Kubernetes payload on
// Metropolis.
//
// This currently requires locally available Consensus, until it is split up
// into Control/Data plane parts (where then the API server must still be
// colocated with Consensus, but all the other services don't have to).
type workerKubernetes struct {
	network     *network.Service
	storageRoot *localstorage.Root

	localRoles        *localRolesValue
	clusterMembership *ClusterMembershipValue
	kubernetesStatus  *KubernetesStatusValue
}

// kubernetesStartup is used internally to provide a reduced (as in MapReduce
// reduced) datum for the main Kubernetes launcher responsible for starting the
// service, if at all.
type kubernetesStartup struct {
	roles      *cpb.NodeRoles
	membership *ClusterMembership
}

// changed informs the Kubernetes launcher whether two different
// kubernetesStartups differ to the point where a restart of Kubernetes should
// happen.
func (k *kubernetesStartup) changed(o *kubernetesStartup) bool {
	hasKubernetesA := k.roles.KubernetesWorker != nil
	hasKubernetesB := o.roles.KubernetesWorker != nil
	if hasKubernetesA != hasKubernetesB {
		return true
	}

	return false
}

func (s *workerKubernetes) run(ctx context.Context) error {
	// TODO(q3k): stop depending on local consensus, split up k8s into control plane
	// and workers.

	// Map/Reduce a *kubernetesStartup from different data sources. This will then
	// populate an Event Value that the actual launcher will use to start
	// Kubernetes.
	//
	// ClusterMambership -M-> clusterMembershipC --R---> startupV
	//                                             |
	//         NodeRoles -M-> rolesC --------------'
	//
	var startupV memory.Value

	clusterMembershipC := make(chan *ClusterMembership)
	rolesC := make(chan *cpb.NodeRoles)

	supervisor.RunGroup(ctx, map[string]supervisor.Runnable{
		// Plain conversion from Event Value to channel.
		"map-cluster-membership": func(ctx context.Context) error {
			supervisor.Signal(ctx, supervisor.SignalHealthy)
			w := s.clusterMembership.Watch()
			defer w.Close()
			for {
				v, err := w.GetHome(ctx)
				if err != nil {
					return err
				}
				clusterMembershipC <- v
			}
		},
		// Plain conversion from Event Value to channel.
		"map-roles": func(ctx context.Context) error {
			supervisor.Signal(ctx, supervisor.SignalHealthy)
			w := s.localRoles.Watch()
			defer w.Close()
			for {
				v, err := w.Get(ctx)
				if err != nil {
					return err
				}
				rolesC <- v
			}
		},
		// Provide config from clusterMembership and roles.
		"reduce-config": func(ctx context.Context) error {
			supervisor.Signal(ctx, supervisor.SignalHealthy)
			var lr *cpb.NodeRoles
			var cm *ClusterMembership
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case lr = <-rolesC:
				case cm = <-clusterMembershipC:
				}
				if lr != nil && cm != nil {
					startupV.Set(&kubernetesStartup{
						roles:      lr,
						membership: cm,
					})
				}
			}
		},
	})

	supervisor.Run(ctx, "run", func(ctx context.Context) error {
		w := startupV.Watch()
		defer w.Close()
		supervisor.Logger(ctx).Infof("Waiting for startup data...")

		// Acquire kubernetesStartup, waiting for it to contain local consensus and a
		// KubernetesWorker local role.
		var d *kubernetesStartup
		for {
			dV, err := w.Get(ctx)
			if err != nil {
				return err
			}
			d = dV.(*kubernetesStartup)
			supervisor.Logger(ctx).Infof("Got new startup data.")
			if d.roles.KubernetesWorker == nil {
				supervisor.Logger(ctx).Infof("No Kubernetes role, not starting.")
				continue
			}
			if d.membership.localConsensus == nil {
				supervisor.Logger(ctx).Warningf("No local consensus, cannot start.")
				continue
			}

			break
		}
		supervisor.Logger(ctx).Infof("Waiting for local consensus...")
		cstW := d.membership.localConsensus.Watch()
		defer cstW.Close()
		cst, err := cstW.GetRunning(ctx)
		if err != nil {
			return fmt.Errorf("waiting for local consensus: %w", err)
		}

		supervisor.Logger(ctx).Infof("Got data, starting Kubernetes...")
		kkv, err := cst.KubernetesClient()
		if err != nil {
			return fmt.Errorf("retrieving kubernetes client: %w", err)
		}

		// Start containerd.
		containerdSvc := &containerd.Service{
			EphemeralVolume: &s.storageRoot.Ephemeral.Containerd,
		}
		if err := supervisor.Run(ctx, "containerd", containerdSvc.Run); err != nil {
			return fmt.Errorf("failed to start containerd service: %w", err)
		}

		// Start building Kubernetes service...
		pki := kpki.New(supervisor.Logger(ctx), kkv)

		kubeSvc := kubernetes.New(kubernetes.Config{
			Node: &d.membership.credentials.Node,
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
			KPKI:    pki,
			Root:    s.storageRoot,
			Network: s.network,
		})
		// Start Kubernetes.
		if err := supervisor.Run(ctx, "kubernetes", kubeSvc.Run); err != nil {
			return fmt.Errorf("failed to start kubernetes service: %w", err)
		}

		// Let downstream know that Kubernetes is running.
		s.kubernetesStatus.set(&KubernetesStatus{
			Svc: kubeSvc,
		})

		supervisor.Signal(ctx, supervisor.SignalHealthy)

		// Restart everything if we get a significantly different config (ie., a config
		// whose change would/should either turn up or tear down Kubernetes).
		for {
			ncI, err := w.Get(ctx)
			if err != nil {
				return err
			}
			nc := ncI.(*kubernetesStartup)
			if nc.changed(d) {
				supervisor.Logger(ctx).Errorf("watcher got new config, restarting")
				return fmt.Errorf("restarting")
			}
		}
	})

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	<-ctx.Done()
	return ctx.Err()
}
