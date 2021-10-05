// package roleserve implements the roleserver/“Role Server”.
//
// The Role Server is responsible for watching the cluster/node state and
// starting/stopping any 'application' code (also called workload code) based on
// the Node's roles.
//
// Each workload code (which would usually be a supervisor runnable) is started
// by a dedicated 'launcher'. These launchers wait for node roles to be
// available from the curator, and either start the related workload sub-runners
// or do nothing ; then they declare themselves as healthy to the supervisor. If
// at any point the role of the node changes (meaning that the node should now
// start or stop the workloads) the launcher just exits and the supervisor
// will restart it.
//
// Currently, this is used to start up the Kubernetes worker code.
package roleserve

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/kubernetes"
	"source.monogon.dev/metropolis/node/kubernetes/pki"
	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// Config is the configuration of the role server.
type Config struct {
	// CuratorDial is a function that the roleserver will use to dial the curator.
	// As both the curator listener and roleserver might restart, this dial function
	// is needed to possibly re-establish connectivity after a full restart of
	// either.
	CuratorDial func(ctx context.Context) (*grpc.ClientConn, error)

	// StorageRoot is a handle to access all of the Node's storage. This is needed
	// as the roleserver spawns complex workloads like Kubernetes which need access
	// to a broad range of storage.
	StorageRoot *localstorage.Root

	// Network is a handle to the network service, used by workloads.
	Network *network.Service

	// KPKI is a handle to initialized Kubernetes PKI stored on etcd. In the future
	// this will probably be provisioned by the Kubernetes workload itself.
	KPKI *pki.PKI

	// NodeID is the node ID on which the roleserver is running.
	NodeID string
}

// Service is the roleserver/“Role Server” service. See the package-level
// documentation for more details.
type Service struct {
	Config

	value memory.Value

	// kwC is a channel populated with updates to the local Node object from the
	// curator, passed over to the Kubernetes Worker launcher.
	kwC chan *cpb.Node
	// kwSvcC is a channel populated by the Kubernetes Worker launcher when the
	// service is started (which then contains the value of spawned Kubernetes
	// workload service) or stopped (which is then nil).
	kwSvcC chan *kubernetes.Service

	// gRPC channel to curator, acquired via Config.CuratorDial, active for the
	// lifecycle of the Service runnable. It's used by the updater
	// sub-runnable.
	curator cpb.CuratorClient
}

// Status is updated by the role service any time one of the subordinate
// workload services is started or stopped.
type Status struct {
	// Kubernetes is set to the Kubernetes workload Service if started/restarted or
	// nil if stopped.
	Kubernetes *kubernetes.Service
}

// New creates a Role Server services from a Config.
func New(c Config) *Service {
	return &Service{
		Config: c,
		kwC:    make(chan *cpb.Node),
		kwSvcC: make(chan *kubernetes.Service),
	}
}

type Watcher struct {
	event.Watcher
}

func (s *Service) Watch() Watcher {
	return Watcher{
		Watcher: s.value.Watch(),
	}
}

func (w *Watcher) Get(ctx context.Context) (*Status, error) {
	v, err := w.Watcher.Get(ctx)
	if err != nil {
		return nil, err
	}
	st := v.(Status)
	return &st, nil
}

// Run the Role Server service, which uses intermediary workload launchers to
// start/stop subordinate services as the Node's roles change.
func (s *Service) Run(ctx context.Context) error {
	supervisor.Logger(ctx).Info("Dialing curator...")
	conn, err := s.CuratorDial(ctx)
	if err != nil {
		return fmt.Errorf("could not dial cluster curator: %w", err)
	}
	defer conn.Close()
	s.curator = cpb.NewCuratorClient(conn)

	if err := supervisor.Run(ctx, "updater", s.runUpdater); err != nil {
		return fmt.Errorf("failed to launch updater: %w", err)
	}

	if err := supervisor.Run(ctx, "cluster-agent", s.runClusterAgent); err != nil {
		return fmt.Errorf("failed to launch cluster agent: %w", err)
	}

	if err := supervisor.Run(ctx, "kubernetes-worker", s.runKubernetesWorkerLauncher); err != nil {
		return fmt.Errorf("failed to start kubernetes-worker launcher: %w", err)
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)

	status := Status{}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case svc := <-s.kwSvcC:
			status.Kubernetes = svc
			s.value.Set(status)
		}
	}
}

// runUpdater runs the updater, a runnable which watchers the cluster via
// curator for any pertinent node changes and distributes them to respective
// workload launchers.
//
// TODO(q3k): this should probably be implemented somewhere as a curator client
// library, maybe one that implements the Event Value interface.
func (s *Service) runUpdater(ctx context.Context) error {
	srv, err := s.curator.Watch(ctx, &cpb.WatchRequest{Kind: &cpb.WatchRequest_NodeInCluster_{
		NodeInCluster: &cpb.WatchRequest_NodeInCluster{
			NodeId: s.NodeID,
		},
	}})
	if err != nil {
		return fmt.Errorf("watch failed: %w", err)
	}
	defer srv.CloseSend()

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	for {
		ev, err := srv.Recv()
		if err != nil {
			return fmt.Errorf("watch event receive failed: %w", err)
		}
		supervisor.Logger(ctx).Infof("Received node event: %+v", ev)
		for _, node := range ev.Nodes {
			if node.Id != s.NodeID {
				continue
			}
			s.kwC <- node
		}
	}

}
