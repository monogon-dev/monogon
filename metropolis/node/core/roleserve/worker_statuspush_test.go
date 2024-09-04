package roleserve

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/testing/protocmp"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/test/util"
	mversion "source.monogon.dev/metropolis/version"
	"source.monogon.dev/osbase/supervisor"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// statusRecodingCurator is a fake implementation of the Curator which updates
// UpdateNodeStatus requests and logs them.
type statusRecordingCurator struct {
	ipb.UnimplementedCuratorServer

	mu            sync.Mutex
	statusReports []*ipb.UpdateNodeStatusRequest
}

func (f *statusRecordingCurator) UpdateNodeStatus(ctx context.Context, req *ipb.UpdateNodeStatusRequest) (*ipb.UpdateNodeStatusResponse, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.statusReports = append(f.statusReports, req)
	return &ipb.UpdateNodeStatusResponse{}, nil
}

// expectReports waits until the given requests have been logged by the
// statusRecordingCurator.
func (f *statusRecordingCurator) expectReports(t *testing.T, want []*ipb.UpdateNodeStatusRequest) {
	t.Helper()

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Second * 10
	err := backoff.Retry(func() error {
		f.mu.Lock()
		defer f.mu.Unlock()

		if diff := cmp.Diff(want, f.statusReports, protocmp.Transform()); diff != "" {
			return fmt.Errorf("unexpected difference:\n%v", diff)
		}
		return nil
	}, bo)
	if err != nil {
		t.Fatal(err)
	}
}

// TestWorkerStatusPush ensures that the status push worker main loop behaves as
// expected. It does not exercise the 'map' runnables.
func TestWorkerStatusPush(t *testing.T) {
	chans := workerStatusPushChannels{
		address:           make(chan string),
		localControlPlane: make(chan *localControlPlane),
		curatorConnection: make(chan *CuratorConnection),
	}

	go supervisor.TestHarness(t, func(ctx context.Context) error {
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		return workerStatusPushLoop(ctx, &chans)
	})

	// Build a loopback gRPC server served by the statusRecordingCurator and connect
	// to it.
	cur := &statusRecordingCurator{}
	srv := grpc.NewServer()
	defer srv.Stop()
	ipb.RegisterCuratorServer(srv, cur)
	lis := bufconn.Listen(1024 * 1024)
	defer lis.Close()
	go func() {
		if err := srv.Serve(lis); err != nil {
			t.Errorf("GRPC serve failed: %v", err)
			return
		}
	}()
	withLocalDialer := grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return lis.Dial()
	})
	cl, err := grpc.Dial("local", withLocalDialer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
	defer cl.Close()

	eph := util.NewEphemeralClusterCredentials(t, 1)
	nodeID := eph.Nodes[0].ID()

	// Actual test code starts here.
	chans.curatorConnection <- &CuratorConnection{
		Credentials: eph.Nodes[0],
		conn:        cl,
	}
	cur.expectReports(t, nil)

	// Provide enough data for the first status report to be submitted.
	chans.address <- "192.0.2.10"
	cur.expectReports(t, []*ipb.UpdateNodeStatusRequest{
		{NodeId: nodeID, Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.10",
			Version:         mversion.Version,
		}},
	})

	// Spurious address update should be ignored.
	chans.address <- "192.0.2.10"
	chans.address <- "192.0.2.11"
	cur.expectReports(t, []*ipb.UpdateNodeStatusRequest{
		{NodeId: nodeID, Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.10",
			Version:         mversion.Version,
		}},
		{NodeId: nodeID, Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.11",
			Version:         mversion.Version,
		}},
	})

	// Enabling and disabling local curator should work.
	chans.localControlPlane <- &localControlPlane{
		curator:   &curator.Service{},
		consensus: &consensus.Service{},
	}
	chans.localControlPlane <- &localControlPlane{
		curator:   nil,
		consensus: nil,
	}
	cur.expectReports(t, []*ipb.UpdateNodeStatusRequest{
		{NodeId: nodeID, Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.10",
			Version:         mversion.Version,
		}},
		{NodeId: nodeID, Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.11",
			Version:         mversion.Version,
		}},
		{NodeId: nodeID, Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.11",
			RunningCurator: &cpb.NodeStatus_RunningCurator{
				Port: int32(common.CuratorServicePort),
			},
			Version: mversion.Version,
		}},
		{NodeId: nodeID, Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.11",
			Version:         mversion.Version,
		}},
	})
}
