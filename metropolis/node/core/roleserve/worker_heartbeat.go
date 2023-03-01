package roleserve

import (
	"context"
	"fmt"
	"io"
	"time"

	"source.monogon.dev/metropolis/node/core/curator"
	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

// workerHeartbeat is a service that periodically updates node's heartbeat
// timestamps within the cluster.
type workerHeartbeat struct {
	network *network.Service

	// clusterMembership will be read.
	clusterMembership *memory.Value[*ClusterMembership]
}

func (s *workerHeartbeat) run(ctx context.Context) error {
	nw := s.network.Watch()
	defer nw.Close()

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

	stream, err := cur.Heartbeat(ctx)
	if err != nil {
		return err
	}

	for {
		if err := stream.Send(&ipb.HeartbeatUpdateRequest{}); err != nil {
			return fmt.Errorf("while sending a heartbeat: %v", err)
		}
		next := time.Now().Add(curator.HeartbeatTimeout)

		_, err := stream.Recv()
		if err == io.EOF {
			return fmt.Errorf("stream closed by the server. Restarting worker...")
		}
		if err != nil {
			return fmt.Errorf("while receiving a heartbeat reply: %v", err)
		}

		time.Sleep(time.Until(next))
	}
}
