// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package rpc

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	apb "source.monogon.dev/metropolis/proto/api"
	epb "source.monogon.dev/metropolis/proto/ext"
	"source.monogon.dev/metropolis/test/util"
)

// testImplementations implements a subset of test cluster services by returning
// 'unimplemented' for every RPC call.
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

	eph := util.NewEphemeralClusterCredentials(t, 1)
	permissions := make(Permissions)
	for k, v := range nodePermissions {
		permissions[k] = v
	}
	ss := ServerSecurity{
		NodeCredentials: eph.Nodes[0],
		nodePermissions: permissions,
	}

	impl := &testImplementation{}
	srv := grpc.NewServer(ss.GRPCOptions(nil)...)
	cpb.RegisterCuratorServer(srv, impl)
	apb.RegisterManagementServer(srv, impl)
	apb.RegisterAAAServer(srv, impl)
	lis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := srv.Serve(lis); err != nil {
			t.Errorf("GRPC serve failed: %v", err)
			return
		}
	}()
	defer lis.Close()
	defer srv.Stop()

	withLocalDialer := grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return lis.Dial()
	})

	// Authenticate as manager externally, ensure that GetRegisterTicket runs.
	cl, err := grpc.NewClient("passthrough:///local",
		grpc.WithTransportCredentials(NewAuthenticatedCredentials(eph.Manager, WantRemoteCluster(eph.CA))),
		withLocalDialer)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	defer cl.Close()
	mgmt := apb.NewManagementClient(cl)
	_, err = mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if s, ok := status.FromError(err); !ok || s.Code() != codes.Unimplemented {
		t.Errorf("GetRegisterTicket returned %v, wanted codes.Unimplemented", err)
	}

	// Authenticate as node externally, ensure that GetRegisterTicket is refused
	// (this is because nodes miss the GET_REGISTER_TICKET permissions).
	cl, err = grpc.NewClient("passthrough:///local",
		grpc.WithTransportCredentials(NewAuthenticatedCredentials(eph.Nodes[0].TLSCredentials(), WantRemoteCluster(eph.CA))),
		withLocalDialer)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
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

	ephCreds, err := NewEphemeralCredentials(sk, WantRemoteCluster(eph.CA))
	if err != nil {
		t.Fatalf("NewEphemeralCredentials: %v", err)
	}
	cl, err = grpc.NewClient("passthrough:///local", grpc.WithTransportCredentials(ephCreds), withLocalDialer)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	defer cl.Close()
	mgmt = apb.NewManagementClient(cl)
	_, err = mgmt.GetRegisterTicket(ctx, &apb.GetRegisterTicketRequest{})
	if s, ok := status.FromError(err); !ok || s.Code() != codes.Unauthenticated {
		t.Errorf("GetRegisterTicket (by ephemeral cert) returned %v, wanted codes.Unauthenticated", err)
	}
}
