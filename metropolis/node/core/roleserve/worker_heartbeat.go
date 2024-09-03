package roleserve

import (
	"context"
	"fmt"
	"io"
	"time"

	"source.monogon.dev/metropolis/node/core/curator"
	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/osbase/event/memory"
	"source.monogon.dev/osbase/supervisor"
)

// workerHeartbeat is a service that periodically updates node's heartbeat
// timestamps within the cluster.
type workerHeartbeat struct {
	network *network.Service

	// curatorConnection will be read.
	curatorConnection *memory.Value[*CuratorConnection]
}

func (s *workerHeartbeat) run(ctx context.Context) error {
	nw := s.network.Status.Watch()
	defer nw.Close()

	w := s.curatorConnection.Watch()
	defer w.Close()
	supervisor.Logger(ctx).Infof("Waiting for curator connection...")
	cc, err := w.Get(ctx)
	if err != nil {
		return err
	}
	supervisor.Logger(ctx).Infof("Got curator connection, starting...")
	cur := ipb.NewCuratorClient(cc.conn)

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
			return fmt.Errorf("stream closed by the server, restarting worker... ")
		}
		if err != nil {
			return fmt.Errorf("while receiving a heartbeat reply: %v", err)
		}

		time.Sleep(time.Until(next))
	}
}
