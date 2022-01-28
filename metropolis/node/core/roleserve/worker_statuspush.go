package roleserve

import (
	"context"
	"fmt"
	"net"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/supervisor"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// workerStatusPush is the Status Pusher, a service responsible for sending
// UpdateNodeStatus RPCs to a cluster whenever a Curator is available.
//
// TODO(q3k): factor this out of the roleserver, there's no need for this to be
// internal, as it only depends on ClusterMembership. This could maybe even live
// in the Network service?
type workerStatusPush struct {
	network *network.Service

	// clusterMembership will be read.
	clusterMembership *ClusterMembershipValue
}

func (s *workerStatusPush) run(ctx context.Context) error {
	nw := s.network.Watch()
	defer nw.Close()

	w := s.clusterMembership.Watch()
	defer w.Close()
	supervisor.Logger(ctx).Infof("Waiting for cluster membership...")
	cm, err := w.GetHome(ctx)
	if err != nil {
		return err
	}
	supervisor.Logger(ctx).Infof("Got cluster membership, starting...")

	nodeID := cm.NodeID()
	conn, err := cm.DialCurator()
	if err != nil {
		return err
	}
	defer conn.Close()
	cur := ipb.NewCuratorClient(conn)

	// Start watch on Network service, update IP address whenever new one is set.
	supervisor.Signal(ctx, supervisor.SignalHealthy)
	var external net.IP
	for {
		st, err := nw.Get(ctx)
		if err != nil {
			return fmt.Errorf("getting network status failed: %w", err)
		}

		if external.Equal(st.ExternalAddress) {
			continue
		}
		supervisor.Logger(ctx).Infof("New external address (%s), submitting update to cluster...", st.ExternalAddress.String())
		_, err = cur.UpdateNodeStatus(ctx, &ipb.UpdateNodeStatusRequest{
			NodeId: nodeID,
			Status: &cpb.NodeStatus{
				ExternalAddress: st.ExternalAddress.String(),
			},
		})
		if err != nil {
			return fmt.Errorf("UpdateNodeStatus failed: %w", err)
		}
		external = st.ExternalAddress
		supervisor.Logger(ctx).Infof("Updated.")
	}
}
