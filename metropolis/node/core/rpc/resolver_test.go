package rpc

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

type testImplementationClusterAware struct {
	ipb.UnimplementedCuratorServer
	apb.UnimplementedAAAServer
	apb.UnimplementedManagementServer

	addresses map[string]string
}

func (t *testImplementationClusterAware) GetClusterInfo(ctx context.Context, req *apb.GetClusterInfoRequest) (*apb.GetClusterInfoResponse, error) {
	return &apb.GetClusterInfoResponse{}, nil
}

func (t *testImplementationClusterAware) Watch(_ *ipb.WatchRequest, srv ipb.Curator_WatchServer) error {
	var nodes []*ipb.Node
	for name, addr := range t.addresses {
		nodes = append(nodes, &ipb.Node{
			Id:    name,
			Roles: &cpb.NodeRoles{ConsensusMember: &cpb.NodeRoles_ConsensusMember{}},
			Status: &cpb.NodeStatus{
				ExternalAddress: addr,
			},
		})
	}
	err := srv.Send(&ipb.WatchEvent{
		Nodes: nodes,
	})
	if err != nil {
		return err
	}
	<-srv.Context().Done()
	return srv.Context().Err()
}

// TestResolverSimple exercises the happy path of the gRPC ResolverBuilder,
// checking that a single node can be used to bootstrap multiple nodes from, and
// ensuring that nodes are being dialed in a round-robin fashion.
//
// TODO(q3k): exercise node removal and re-dial of updater to another node.
func TestResolverSimple(t *testing.T) {
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Make three nodes for testing, each with its own bufconn listener.
	numCurators := 3
	eph := NewEphemeralClusterCredentials(t, numCurators)

	listeners := make([]net.Listener, numCurators)
	for i := 0; i < numCurators; i++ {
		lis, err := net.Listen("tcp", "")
		if err != nil {
			t.Fatalf("Listen failed: %v", err)
		}
		listeners[i] = lis
	}

	addresses := make(map[string]string)
	for i, lis := range listeners {
		name := eph.Nodes[i].ID()
		addresses[name] = lis.Addr().String()
	}
	impls := make([]*testImplementationClusterAware, numCurators)
	for i := 0; i < numCurators; i++ {
		impls[i] = &testImplementationClusterAware{
			addresses: addresses,
		}
	}

	servers := make([]*grpc.Server, numCurators)
	for i := 0; i < numCurators; i++ {
		i := i
		ss := ServerSecurity{
			NodeCredentials: eph.Nodes[i],
		}
		servers[i] = grpc.NewServer(ss.GRPCOptions(nil)...)
		ipb.RegisterCuratorServer(servers[i], impls[i])
		apb.RegisterAAAServer(servers[i], impls[i])
		apb.RegisterManagementServer(servers[i], impls[i])
		go func() {
			if err := servers[i].Serve(listeners[i]); err != nil {
				t.Fatalf("GRPC serve failed: %v", err)
			}
		}()

		defer listeners[i].Close()
		defer servers[i].Stop()
	}

	r := NewClusterResolver()
	r.logger = func(f string, args ...interface{}) {
		log.Printf(f, args...)
	}
	defer r.Close()

	creds := credentials.NewTLS(&tls.Config{
		Certificates:          []tls.Certificate{eph.Manager},
		InsecureSkipVerify:    true,
		VerifyPeerCertificate: verifyClusterCertificate(eph.CA),
	})
	cl, err := grpc.Dial("metropolis:///control", grpc.WithTransportCredentials(creds), grpc.WithResolvers(r))
	if err != nil {
		t.Fatalf("Could not dial: %v", err)
	}

	// Add first node to bootstrap node information from.
	r.AddNode(eph.Nodes[0].ID(), listeners[0].Addr().String())

	mgmt := apb.NewManagementClient(cl)
	_, err = mgmt.GetClusterInfo(ctx, &apb.GetClusterInfoRequest{})
	if err != nil {
		t.Fatalf("Running initial GetClusterInfo failed: %v", err)
	}

	// Wait until client finds all three nodes.
	r.condCurators.L.Lock()
	for len(r.curators) < 3 {
		r.condCurators.Wait()
	}
	curators := r.curators
	r.condCurators.L.Unlock()

	// Ensure the three nodes as are expected.
	for i, node := range eph.Nodes {
		if got, want := curators[node.ID()], listeners[i].Addr().String(); want != got {
			t.Errorf("Node %s: wanted address %q, got %q", node.ID(), want, got)
		}
	}

	// Stop first node, make sure the call now reaches the other servers. This will
	// happen due to the resolver's round-robin behaviour, not because this node is
	// dropped from the active set of nodes.
	servers[0].Stop()
	listeners[0].Close()

	for i := 0; i < 10; i++ {
		_, err = mgmt.GetClusterInfo(ctx, &apb.GetClusterInfoRequest{})
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		t.Errorf("Running GetClusterInfo after stopping first node failed: %v", err)
	}

	// Close the builder, new dials should fail.
	r.Close()
	_, err = grpc.Dial(MetropolisControlAddress, grpc.WithTransportCredentials(creds), grpc.WithResolvers(r), grpc.WithBlock())
	// String comparison required because the underlying gRPC code does not wrap the
	// error.
	if want, got := ResolverClosed, err; !strings.Contains(got.Error(), want.Error()) {
		t.Errorf("Unexpected dial error after closing builder, wanted %q, got %q", want, got)
	}
}
