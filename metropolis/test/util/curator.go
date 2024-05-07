package util

import (
	"context"
	"net"
	"testing"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"

	"source.monogon.dev/osbase/event/memory"
)

// TestCurator is a shim Curator implementation that serves pending Watch
// requests based on data submitted to a channel.
type TestCurator struct {
	apb.UnimplementedCuratorServer

	watchC    chan *apb.WatchEvent
	updateReq memory.Value[*apb.UpdateNodeClusterNetworkingRequest]
}

// Watch implements a minimum Watch which just returns all nodes at once.
func (t *TestCurator) Watch(_ *apb.WatchRequest, srv apb.Curator_WatchServer) error {
	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ev := <-t.watchC:
			if err := srv.Send(ev); err != nil {
				return err
			}
		}
	}
}

func (t *TestCurator) UpdateNodeClusterNetworking(ctx context.Context, req *apb.UpdateNodeClusterNetworkingRequest) (*apb.UpdateNodeClusterNetworkingResponse, error) {
	t.updateReq.Set(req)
	return &apb.UpdateNodeClusterNetworkingResponse{}, nil
}

// NodeWithPrefixes submits a given node/key/address with prefixes to the Watch
// event channel.
func (t *TestCurator) NodeWithPrefixes(key wgtypes.Key, id, address string, prefixes ...string) {
	var p []*cpb.NodeClusterNetworking_Prefix
	for _, prefix := range prefixes {
		p = append(p, &cpb.NodeClusterNetworking_Prefix{Cidr: prefix})
	}
	n := &apb.Node{
		Id: id,
		Status: &cpb.NodeStatus{
			ExternalAddress: address,
		},
		Clusternet: &cpb.NodeClusterNetworking{
			WireguardPubkey: key.PublicKey().String(),
			Prefixes:        p,
		},
		Roles: &cpb.NodeRoles{
			ConsensusMember: &cpb.NodeRoles_ConsensusMember{},
		},
	}
	t.watchC <- &apb.WatchEvent{
		Nodes: []*apb.Node{
			n,
		},
	}
}

// DeleteNode submits a given node for deletion to the Watch event channel.
func (t *TestCurator) DeleteNode(id string) {
	t.watchC <- &apb.WatchEvent{
		NodeTombstones: []*apb.WatchEvent_NodeTombstone{
			{
				NodeId: id,
			},
		},
	}
}

// MakeTestCurator returns a working TestCurator alongside a grpc connection to
// it.
func MakeTestCurator(t *testing.T) (*TestCurator, *grpc.ClientConn) {
	cur := &TestCurator{
		watchC: make(chan *apb.WatchEvent),
	}

	srv := grpc.NewServer()
	apb.RegisterCuratorServer(srv, cur)
	externalLis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := srv.Serve(externalLis); err != nil {
			t.Errorf("GRPC serve failed: %v", err)
			return
		}
	}()
	withLocalDialer := grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return externalLis.Dial()
	})
	cl, err := grpc.Dial("local", withLocalDialer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Dialing GRPC failed: %v", err)
	}

	return cur, cl
}
