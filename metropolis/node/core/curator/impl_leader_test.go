package curator

import (
	"bytes"
	"context"
	"testing"

	"go.etcd.io/etcd/integration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/rpc"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// fakeLeader creates a curatorLeader without any underlying leader election, in
// its own etcd namespace. It starts public and local gRPC listeners and returns
// clients to them.
//
// The gRPC listeners are replicated to behave as when running the Curator
// within Metropolis, so all calls performed will be authenticated and encrypted
// the same way.
//
// This is used to test functionality of the individual curatorLeader RPC
// implementations without the overhead of having to wait for a leader election.
func fakeLeader(t *testing.T) fakeLeaderData {
	t.Helper()
	// Set up context whose cancel function will be returned to the user for
	// terminating all harnesses started by this function.
	ctx, ctxC := context.WithCancel(context.Background())

	// Start a single-node etcd cluster.
	cluster := integration.NewClusterV3(nil, &integration.ClusterConfig{
		Size: 1,
	})
	// Terminate the etcd cluster on context cancel.
	go func() {
		<-ctx.Done()
		cluster.Terminate(nil)
	}()

	// Create etcd client to test cluster.
	cl := client.NewLocal(cluster.Client(0))

	// Create a fake lock key/value and retrieve its revision. This replaces the
	// leader election functionality in the curator to enable faster and more
	// focused tests.
	lockKey := "/test-lock"
	res, err := cl.Put(ctx, lockKey, "fake key")
	if err != nil {
		t.Fatalf("setting fake leader key failed: %v", err)
	}
	lockRev := res.Header.Revision

	// Build a curator leader object. This implements methods that will be
	// exercised by tests.
	leader := newCuratorLeader(&leadership{
		lockKey: lockKey,
		lockRev: lockRev,
		etcd:    cl,
	})

	// Build a test cluster PKI and node/manager certificates, and create the
	// listener security parameters which will authenticate incoming requests.
	ephemeral := rpc.NewEphemeralClusterCredentials(t, 2)
	nodeCredentials := ephemeral.Nodes[0]

	cNode := NewNodeForBootstrap(nil, nodeCredentials.PublicKey())
	// Inject new node into leader, using curator bootstrap functionality.
	if err := BootstrapFinish(ctx, cl, &cNode, nodeCredentials.PublicKey()); err != nil {
		t.Fatalf("could not finish node bootstrap: %v", err)
	}

	// Create security interceptors for both gRPC listeners.
	externalSec := &rpc.ExternalServerSecurity{
		NodeCredentials: nodeCredentials,
	}
	localSec := &rpc.LocalServerSecurity{
		Node: &nodeCredentials.Node,
	}
	// Create a curator gRPC server which performs authentication as per the created
	// listenerSecurity and is backed by the created leader.
	externalSrv := externalSec.SetupExternalGRPC(leader)
	localSrv := localSec.SetupLocalGRPC(leader)
	// The gRPC server will listen on an internal 'loopback' buffer.
	externalLis := bufconn.Listen(1024 * 1024)
	localLis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := externalSrv.Serve(externalLis); err != nil {
			t.Fatalf("GRPC serve failed: %v", err)
		}
	}()
	go func() {
		if err := localSrv.Serve(localLis); err != nil {
			t.Fatalf("GRPC serve failed: %v", err)
		}
	}()

	// Stop the gRPC server on context cancel.
	go func() {
		<-ctx.Done()
		externalSrv.Stop()
		localSrv.Stop()
	}()

	// Create an authenticated manager gRPC client.
	mcl, err := rpc.NewAuthenticatedClientTest(externalLis, ephemeral.Manager, ephemeral.CA)
	if err != nil {
		t.Fatalf("Dialing external GRPC failed: %v", err)
	}

	// Create a locally authenticated node gRPC client.
	lcl, err := rpc.NewNodeClientTest(localLis)
	if err != nil {
		t.Fatalf("Dialing local GRPC failed: %v", err)
	}

	// Close the clients on context cancel.
	go func() {
		<-ctx.Done()
		mcl.Close()
		lcl.Close()
	}()

	return fakeLeaderData{
		mgmtConn:      mcl,
		localNodeConn: lcl,
		localNodeID:   nodeCredentials.ID(),
		otherNodeID:   ephemeral.Nodes[1].ID(),
		cancel:        ctxC,
	}
}

// fakeLeaderData is returned by fakeLeader and contains information about the
// newly created leader and connections to its gRPC listeners.
type fakeLeaderData struct {
	// mgmtConn is a gRPC connection to the leader's public gRPC interface,
	// authenticated as a cluster manager.
	mgmtConn grpc.ClientConnInterface
	// localNodeConn is a gRPC connection to the leader's internal/local node gRPC
	// interface, which usually runs on a domain socket and is only available to
	// other Metropolis node code.
	localNodeConn grpc.ClientConnInterface
	// localNodeID is the NodeID of the fake node that the leader is running on.
	localNodeID string
	// otherNodeID is the NodeID of some other node present in the curator
	// state.
	otherNodeID string
	// cancel shuts down the fake leader and all client connections.
	cancel context.CancelFunc
}

// TestManagementRegisterTicket exercises the Management.GetRegisterTicket RPC.
func TestManagementRegisterTicket(t *testing.T) {
	cl := fakeLeader(t)
	defer cl.cancel()

	mgmt := apb.NewManagementClient(cl.mgmtConn)

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Retrieve ticket twice.
	res1, err := mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if err != nil {
		t.Fatalf("GetRegisterTicket failed: %v", err)
	}
	res2, err := mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if err != nil {
		t.Fatalf("GetRegisterTicket failed: %v", err)
	}

	// Ensure tickets are set and the same.
	if len(res1.Ticket) == 0 {
		t.Errorf("Ticket is empty")
	}
	if !bytes.Equal(res1.Ticket, res2.Ticket) {
		t.Errorf("Unexpected ticket change between calls")
	}
}

// TestClusterUpdateNodeStatus exercises the Curator.UpdateNodeStatus RPC by
// sending node updates and making sure they are reflected in subsequent Watch
// events.
func TestClusterUpdateNodeStatus(t *testing.T) {
	cl := fakeLeader(t)
	defer cl.cancel()

	curator := ipb.NewCuratorClient(cl.localNodeConn)

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Retrieve initial node data, it should have no status set.
	value, err := curator.Watch(ctx, &ipb.WatchRequest{
		Kind: &ipb.WatchRequest_NodeInCluster_{
			NodeInCluster: &ipb.WatchRequest_NodeInCluster{
				NodeId: cl.localNodeID,
			},
		},
	})
	if err != nil {
		t.Fatalf("Could not request node watch: %v", err)
	}
	ev, err := value.Recv()
	if err != nil {
		t.Fatalf("Could not receive initial node value: %v", err)
	}
	if status := ev.Nodes[0].Status; status != nil {
		t.Errorf("Initial node value contains status, should be nil: %+v", status)
	}

	// Update status...
	_, err = curator.UpdateNodeStatus(ctx, &ipb.UpdateNodeStatusRequest{
		NodeId: cl.localNodeID,
		Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.10",
		},
	})
	if err != nil {
		t.Fatalf("UpdateNodeStatus: %v", err)
	}

	// ... and expect it to be reflected in the new node value.
	for {
		ev, err = value.Recv()
		if err != nil {
			t.Fatalf("Could not receive second node value: %v", err)
		}
		// Keep waiting until we get a status.
		status := ev.Nodes[0].Status
		if status == nil {
			continue
		}
		if want, got := "192.0.2.10", status.ExternalAddress; want != got {
			t.Errorf("Wanted external address %q, got %q", want, got)
		}
		break
	}

	// Expect updating some other node's ID to fail.
	_, err = curator.UpdateNodeStatus(ctx, &ipb.UpdateNodeStatusRequest{
		NodeId: cl.otherNodeID,
		Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.10",
		},
	})
	if err == nil {
		t.Errorf("UpdateNodeStatus for other node (%q vs local %q) succeeded, should have failed", cl.localNodeID, cl.otherNodeID)
	}
}

// TestManagementClusterInfo exercises GetClusterInfo after setting a status.
func TestMangementClusterInfo(t *testing.T) {
	cl := fakeLeader(t)
	defer cl.cancel()

	mgmt := apb.NewManagementClient(cl.mgmtConn)
	curator := ipb.NewCuratorClient(cl.localNodeConn)

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Update status to set an external address.
	_, err := curator.UpdateNodeStatus(ctx, &ipb.UpdateNodeStatusRequest{
		NodeId: cl.localNodeID,
		Status: &cpb.NodeStatus{
			ExternalAddress: "192.0.2.10",
		},
	})
	if err != nil {
		t.Fatalf("UpdateNodeStatus: %v", err)
	}

	// Retrieve cluster info and make sure it's as expected.
	res, err := mgmt.GetClusterInfo(ctx, &apb.GetClusterInfoRequest{})
	if err != nil {
		t.Fatalf("GetClusterInfo failed: %v", err)
	}

	nodes := res.ClusterDirectory.Nodes
	if want, got := 1, len(nodes); want != got {
		t.Fatalf("ClusterDirectory.Nodes contains %d elements, wanted %d", want, got)
	}
	node := nodes[0]

	// Address should match address set from status.
	if want, got := 1, len(node.Addresses); want != got {
		t.Fatalf("ClusterDirectory.Nodes[0].Addresses has %d elements, wanted %d", want, got)
	}
	if want, got := "192.0.2.10", node.Addresses[0].Host; want != got {
		t.Fatalf("Nodes[0].Addresses[0].Host is %q, wanted %q", want, got)
	}
}
