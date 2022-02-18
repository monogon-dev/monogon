package rpc

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	apb "source.monogon.dev/metropolis/proto/api"
	epb "source.monogon.dev/metropolis/proto/ext"
)

// testImplementations implements ClusterServices by returning 'unimplementd'
// for every RPC call.
type testImplementation struct {
	cpb.UnimplementedCuratorServer
	apb.UnimplementedAAAServer
	apb.UnimplementedManagementServer
}

// TestExternalServerSecurity ensures that the unary interceptor of the
// ServerSecurity structure works, and authenticates/authorizes incoming RPCs as
// expected.
func TestExternalServerSecurity(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	eph := NewEphemeralClusterCredentials(t, 1)
	permissions := make(Permissions)
	for k, v := range nodePermissions {
		permissions[k] = v
	}
	ss := ExternalServerSecurity{
		NodeCredentials: eph.Nodes[0],
		nodePermissions: permissions,
	}

	impl := &testImplementation{}
	srv := ss.SetupExternalGRPC(nil, impl)
	lis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := srv.Serve(lis); err != nil {
			t.Fatalf("GRPC serve failed: %v", err)
		}
	}()
	defer lis.Close()
	defer srv.Stop()

	// Authenticate as manager externally, ensure that GetRegisterTicket runs.
	cl, err := NewAuthenticatedClientTest(lis, eph.Manager, eph.CA)
	if err != nil {
		t.Fatalf("NewAuthenticatedClient: %v", err)
	}
	defer cl.Close()
	mgmt := apb.NewManagementClient(cl)
	_, err = mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if s, ok := status.FromError(err); !ok || s.Code() != codes.Unimplemented {
		t.Errorf("GetRegisterTicket returned %v, wanted codes.Unimplemented", err)
	}

	// Authenticate as node externally, ensure that GetRegisterTicket is refused
	// (this is because nodes miss the GET_REGISTER_TICKET permissions).
	cl, err = NewAuthenticatedClientTest(lis, eph.Nodes[0].TLSCredentials(), eph.CA)
	if err != nil {
		t.Fatalf("NewAuthenticatedClient: %v", err)
	}
	defer cl.Close()
	mgmt = apb.NewManagementClient(cl)
	_, err = mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if s, ok := status.FromError(err); !ok || s.Code() != codes.PermissionDenied {
		t.Errorf("GetRegisterTicket (by external node) returned %v, wanted codes.PermissionDenied", err)
	}

	// Give the node GET_REGISTER_TICKET permissions and try again. This should pass.
	permissions[epb.Permission_PERMISSION_GET_REGISTER_TICKET] = true
	_, err = mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if s, ok := status.FromError(err); !ok || s.Code() != codes.Unimplemented {
		t.Errorf("GetRegisterTicket returned %v, wanted codes.Unimplemented", err)
	}
	permissions[epb.Permission_PERMISSION_GET_REGISTER_TICKET] = false

	// Authenticate with an ephemeral/self-signed certificate, ensure that
	// GetRegisterTicket is refused (this is because GetRegisterTicket requires an
	// authenticated connection).
	_, sk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	cl, err = NewEphemeralClientTest(lis, sk, eph.CA)
	if err != nil {
		t.Fatalf("NewEphemeralClient: %v", err)
	}
	defer cl.Close()
	mgmt = apb.NewManagementClient(cl)
	_, err = mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if s, ok := status.FromError(err); !ok || s.Code() != codes.Unauthenticated {
		t.Errorf("GetRegisterTicket (by ephemeral cert) returned %v, wanted codes.Unauthenticated", err)
	}
}

// TestLocalServerSecurity ensures that the unary interceptor of the
// LocalServerSecurity structure works, and authenticates/authorizes incoming
// RPCs as expected.
func TestLocalServerSecurity(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	eph := NewEphemeralClusterCredentials(t, 1)

	permissions := make(Permissions)
	for k, v := range nodePermissions {
		permissions[k] = v
	}

	ls := LocalServerSecurity{
		Node:            &eph.Nodes[0].Node,
		nodePermissions: permissions,
	}

	impl := &testImplementation{}
	srv := ls.SetupLocalGRPC(nil, impl)
	lis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := srv.Serve(lis); err != nil {
			t.Fatalf("GRPC serve failed: %v", err)
		}
	}()
	defer lis.Close()
	defer srv.Stop()

	// Nodes should have access to Curator.Watch.
	cl, err := NewNodeClientTest(lis)
	if err != nil {
		t.Fatalf("NewAuthenticatedClient: %v", err)
	}
	defer cl.Close()

	curator := cpb.NewCuratorClient(cl)
	req := &cpb.WatchRequest{
		Kind: &cpb.WatchRequest_NodeInCluster_{
			NodeInCluster: &cpb.WatchRequest_NodeInCluster{
				NodeId: eph.Nodes[0].ID(),
			},
		},
	}
	w, err := curator.Watch(ctx, req)
	if err != nil {
		t.Fatalf("Watch: %v", err)
	}
	_, err = w.Recv()
	if s, ok := status.FromError(err); !ok || s.Code() != codes.Unimplemented {
		t.Errorf("Watch (by local node) returned %v, wanted codes.Unimplemented", err)
	}

	// Take away the node's PERMISSION_READ_CLUSTER_STATUS permissions and try
	// again. This should fail.
	permissions[epb.Permission_PERMISSION_READ_CLUSTER_STATUS] = false
	w, err = curator.Watch(ctx, req)
	if err != nil {
		t.Fatalf("Watch: %v", err)
	}
	_, err = w.Recv()
	if s, ok := status.FromError(err); !ok || s.Code() != codes.PermissionDenied {
		t.Errorf("Watch (by local node after removing permission) returned %v, wanted codes.PermissionDenied", err)
	}
}
