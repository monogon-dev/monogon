package roleserve

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"

	common "source.monogon.dev/metropolis/node"
	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// workerStatusPush is the Status Pusher, a service responsible for sending
// UpdateNodeStatus RPCs to a cluster whenever a Curator is available.
type workerStatusPush struct {
	network *network.Service

	// clusterMembership will be read.
	clusterMembership *memory.Value[*ClusterMembership]
}

// workerStatusPushChannels contain all the channels between the status pusher's
// 'map' runnables (waiting on Event Values) and the main loop.
type workerStatusPushChannels struct {
	// address of the node, or empty if none. Retrieved from network service.
	address chan string
	// nodeID of this node. Populated whenever available from ClusterMembership.
	nodeID chan string
	// curatorClient connecting to the cluster, populated whenever available from
	// ClusterMembership. Used to actually submit the update.
	curatorClient chan ipb.CuratorClient
	// localCurator describing whether this node has a locally running curator on
	// the default port. Retrieved from ClusterMembership.
	localCurator chan bool
}

// workerStatusPushLoop runs the main loop acting on data received from
// workerStatusPushChannels.
func workerStatusPushLoop(ctx context.Context, chans *workerStatusPushChannels) error {
	status := cpb.NodeStatus{}
	var cur ipb.CuratorClient
	var nodeID string

	for {
		changed := false

		select {
		case <-ctx.Done():
			return fmt.Errorf("while waiting for map updates: %w", ctx.Err())

		case address := <-chans.address:
			if address != status.ExternalAddress {
				supervisor.Logger(ctx).Infof("Got external address: %s", address)
				status.ExternalAddress = address
				changed = true
			}

		case newNodeID := <-chans.nodeID:
			if nodeID != newNodeID {
				supervisor.Logger(ctx).Infof("Got new NodeID: %s", newNodeID)
				nodeID = newNodeID
				changed = true
			}

		case cur = <-chans.curatorClient:
			supervisor.Logger(ctx).Infof("Got curator connection.")

		case localCurator := <-chans.localCurator:
			if status.RunningCurator == nil && localCurator {
				supervisor.Logger(ctx).Infof("Got new local curator state: running")
				status.RunningCurator = &cpb.NodeStatus_RunningCurator{
					Port: int32(common.CuratorServicePort),
				}
				changed = true
			}
			if status.RunningCurator != nil && !localCurator {
				supervisor.Logger(ctx).Infof("Got new local curator state: not running")
				status.RunningCurator = nil
				changed = true
			}
		}

		if cur != nil && nodeID != "" && changed && status.ExternalAddress != "" {
			txt, _ := prototext.Marshal(&status)
			supervisor.Logger(ctx).Infof("Submitting status: %q", txt)
			_, err := cur.UpdateNodeStatus(ctx, &ipb.UpdateNodeStatusRequest{
				NodeId: nodeID,
				Status: &status,
			})
			if err != nil {
				return fmt.Errorf("UpdateNodeStatus failed: %w", err)
			}
		}
	}
}

func (s *workerStatusPush) run(ctx context.Context) error {
	chans := workerStatusPushChannels{
		address:       make(chan string),
		nodeID:        make(chan string),
		curatorClient: make(chan ipb.CuratorClient),
		localCurator:  make(chan bool),
	}

	// All the channel sends in the map runnables are preemptible by a context
	// cancelation. This is because workerStatusPushLoop can crash while processing
	// the events, requiring a restart of this runnable. Without the preemption this
	// runnable could get stuck.

	supervisor.Run(ctx, "map-network", func(ctx context.Context) error {
		nw := s.network.Watch()
		defer nw.Close()

		supervisor.Signal(ctx, supervisor.SignalHealthy)
		for {
			st, err := nw.Get(ctx)
			if err != nil {
				return fmt.Errorf("getting network status failed: %w", err)
			}
			select {
			case chans.address <- st.ExternalAddress.String():
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})
	supervisor.Run(ctx, "map-cluster-membership", func(ctx context.Context) error {
		supervisor.Signal(ctx, supervisor.SignalHealthy)

		var conn *grpc.ClientConn
		defer func() {
			if conn != nil {
				conn.Close()
			}
		}()

		w := s.clusterMembership.Watch()
		defer w.Close()
		supervisor.Logger(ctx).Infof("Waiting for cluster membership...")
		for {
			cm, err := w.Get(ctx, FilterHome())
			if err != nil {
				return fmt.Errorf("getting cluster membership status failed: %w", err)
			}

			if conn == nil {
				conn, err = cm.DialCurator()
				if err != nil {
					return fmt.Errorf("when attempting to connect to curator: %w", err)
				}
				select {
				case chans.curatorClient <- ipb.NewCuratorClient(conn):
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			select {
			case chans.localCurator <- cm.localCurator != nil:
			case <-ctx.Done():
				return ctx.Err()
			}
			select {
			case chans.nodeID <- cm.NodeID():
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	return workerStatusPushLoop(ctx, &chans)
}
