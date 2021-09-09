package curator

import (
	"bytes"
	"context"
	"testing"

	"go.etcd.io/etcd/integration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/rpc"
	apb "source.monogon.dev/metropolis/proto/api"
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
	leader := newCuratorLeader(leadership{
		lockKey: lockKey,
		lockRev: lockRev,
		etcd:    cl,
	})

	// Build a test cluster PKI and node/manager certificates, and create the
	// listener security parameters which will authenticate incoming requests.
	ephemeral := rpc.NewEphemeralClusterCredentials(t, 1)
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
