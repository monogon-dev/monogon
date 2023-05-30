package roleserve

import (
	"context"

	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/core/network/hostsfile"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
)

// workerHostsfile run the //metropolis/node/core/network/hostsfile service,
// which in turn populates /etc/hosts, /etc/machine-id and updates the ESP-stored
// ClusterDirectory (used to Join the cluster after a machine reboots).
type workerHostsfile struct {
	storageRoot *localstorage.Root

	// network will be read. It provides data about the local node's address.
	network           *network.Service
	curatorConnection *memory.Value[*curatorConnection]

	// clusterDirectorySaved will be written. A value of true indicates that the
	// cluster directory has been successfully written at least once to the ESP.
	clusterDirectorySaved *memory.Value[bool]
}

func (s *workerHostsfile) run(ctx context.Context) error {
	w := s.curatorConnection.Watch()
	defer w.Close()
	supervisor.Logger(ctx).Infof("Waiting for curator connection...")
	cc, err := w.Get(ctx)
	if err != nil {
		return err
	}
	supervisor.Logger(ctx).Infof("Got cluster membership, starting...")

	cur := ipb.NewCuratorClient(cc.conn)

	svc := hostsfile.Service{
		Config: hostsfile.Config{
			Network:               s.network,
			Ephemeral:             &s.storageRoot.Ephemeral,
			ESP:                   &s.storageRoot.ESP,
			NodeID:                cc.nodeID(),
			Curator:               cur,
			ClusterDirectorySaved: s.clusterDirectorySaved,
		},
	}

	return svc.Run(ctx)
}
