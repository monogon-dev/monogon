package roleserve

import (
	"context"
	"net"

	"source.monogon.dev/metropolis/node/core/clusternet"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

type workerClusternet struct {
	storageRoot *localstorage.Root

	// clusterMembership will be read.
	clusterMembership *memory.Value[*ClusterMembership]
	// podNetwork will be read.
	podNetwork *memory.Value[*clusternet.Prefixes]
	network    *network.Service
}

func (s *workerClusternet) run(ctx context.Context) error {
	w := s.clusterMembership.Watch()
	defer w.Close()
	supervisor.Logger(ctx).Infof("Waiting for cluster membership...")
	cm, err := w.Get(ctx, FilterHome())
	if err != nil {
		return err
	}
	supervisor.Logger(ctx).Infof("Got cluster membership, starting...")

	conn, err := cm.DialCurator()
	if err != nil {
		return err
	}
	defer conn.Close()
	cur := ipb.NewCuratorClient(conn)

	svc := clusternet.Service{
		Curator: cur,
		ClusterNet: net.IPNet{
			IP:   []byte{10, 0, 0, 0},
			Mask: net.IPMask{255, 255, 0, 0},
		},
		DataDirectory:             &s.storageRoot.Data.Kubernetes.ClusterNetworking,
		LocalKubernetesPodNetwork: s.podNetwork,
		Network:                   s.network.Value(),
	}
	return svc.Run(ctx)
}
