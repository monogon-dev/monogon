package roleserve

import (
	"context"
	"fmt"
	"net"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/pkg/supervisor"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// runClusterAgent runs the ClusterAgent, a runnable responsible for reporting
// the status of the local node to the cluster.
//
// This currently only reports the node's external address to the cluster
// whenever it changes.
func (s *Service) runClusterAgent(ctx context.Context) error {
	w := s.Network.Watch()
	defer w.Close()

	var external net.IP

	for {
		st, err := w.Get(ctx)
		if err != nil {
			return fmt.Errorf("getting network status failed: %w", err)
		}

		if external.Equal(st.ExternalAddress) {
			continue
		}

		external = st.ExternalAddress
		supervisor.Logger(ctx).Infof("New external address (%s), submitting update to cluster...", external.String())

		_, err = s.curator.UpdateNodeStatus(ctx, &ipb.UpdateNodeStatusRequest{
			NodeId: s.NodeID,
			Status: &cpb.NodeStatus{
				ExternalAddress: external.String(),
			},
		})
		if err != nil {
			return fmt.Errorf("UpdateNodeStatus failed: %w", err)
		}
		supervisor.Logger(ctx).Infof("Updated.")
	}
}
