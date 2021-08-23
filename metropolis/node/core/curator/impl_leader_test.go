package curator

import (
	"bytes"
	"context"
	"testing"

	"go.etcd.io/etcd/integration"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	apb "source.monogon.dev/metropolis/proto/api"
)

// fakeLeader creates a curatorLeader without any underlying leader election, in
// its own etcd namespace.
//
// This is used to test functionality of the individual curatorLeader RPC
// implementations without the overhead of having to wait for a leader election.
func fakeLeader(t *testing.T) (*curatorLeader, context.CancelFunc) {
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

	// Return a normal curator leader object that directly implements the tested
	// RPC methods. This will be exercised by tests.
	return newCuratorLeader(leadership{
		lockKey: lockKey,
		lockRev: lockRev,
		etcd:    cl,
	}), ctxC
}

// TestManagementRegisterTicket exercises the Management.GetRegisterTicket RPC.
func TestManagementRegisterTicket(t *testing.T) {
	l, cancel := fakeLeader(t)
	defer cancel()

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Retrieve ticket twice.
	res1, err := l.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if err != nil {
		t.Fatalf("GetRegisterTicket failed: %v", err)
	}
	res2, err := l.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
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
