package curator

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"testing"

	"go.etcd.io/etcd/integration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/test/bufconn"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/pkg/pki"
	apb "source.monogon.dev/metropolis/proto/api"
)

// fakeLeader creates a curatorLeader without any underlying leader election, in
// its own etcd namespace. It starts a gRPC listener of its public services
// implementation and returns a client to it.
//
// The entire gRPC layer is encrypted, authenticated and authorized in the same
// way as by the full Curator codebase running in Metropolis. An ephemeral
// cluster CA and node/manager credentials are created, and are used to
// establish a secure channel when creating the gRPC listener and client.
//
// This is used to test functionality of the individual curatorLeader RPC
// implementations without the overhead of having to wait for a leader election.
func fakeLeader(t *testing.T) (grpc.ClientConnInterface, context.CancelFunc) {
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
	node, manager, ca := pki.EphemeralClusterCredentials(t)
	sec := &listenerSecurity{
		nodeCredentials:      node,
		clusterCACertificate: ca,
	}

	// Create a curator gRPC server which performs authentication as per the created
	// listenerSecurity and is backed by the created leader.
	srv := sec.setupPublicGRPC(leader)
	// The gRPC server will listen on an internal 'loopback' buffer.
	lis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := srv.Serve(lis); err != nil {
			t.Fatalf("GRPC serve failed: %v", err)
		}
	}()
	// Stop the gRPC server on context cancel.
	go func() {
		<-ctx.Done()
		srv.Stop()
	}()

	// Create an authenticated manager gRPC client.
	// TODO(q3k): factor this out to its own library, alongside the code in //metropolis/test/e2e/client.go.
	pool := x509.NewCertPool()
	pool.AddCert(ca)
	gclCreds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{manager},
		RootCAs:      pool,
	})
	gcl, err := grpc.Dial("test-server", grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(gclCreds))
	if err != nil {
		t.Fatalf("Dialing local GRPC failed: %v", err)
	}
	// Close the client on context cancel.
	go func() {
		<-ctx.Done()
		gcl.Close()
	}()

	return gcl, ctxC
}

// TestManagementRegisterTicket exercises the Management.GetRegisterTicket RPC.
func TestManagementRegisterTicket(t *testing.T) {
	cl, cancel := fakeLeader(t)
	defer cancel()

	mgmt := apb.NewManagementClient(cl)

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
