package resolver

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/rpc"
	apb "source.monogon.dev/metropolis/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
	"source.monogon.dev/metropolis/test/util"
)

// fakeCuratorClusterAware is a fake curator implementation that has a vague
// concept of a cluster and the current leader within a cluster. Every instance
// of a fakeCuratorClusterAware is backed by a net.Listener, and is aware of
// the other implementations' net.Listeners.
type fakeCuratorClusterAware struct {
	ipb.UnimplementedCuratorServer
	apb.UnimplementedAAAServer
	apb.UnimplementedManagementServer

	// mu guards all the other fields of this struct.
	mu sync.Mutex
	// listeners is the collection of all listeners that make up this cluster, keyed
	// by node ID.
	listeners map[string]net.Listener
	// thisNode is the node ID of this fake.
	thisNode string
	// leader is the node ID of the leader of the cluster.
	leader string
}

// Watch implements a minimum Watch which just returns all nodes at once.
func (t *fakeCuratorClusterAware) Watch(_ *ipb.WatchRequest, srv ipb.Curator_WatchServer) error {
	var nodes []*ipb.Node
	for name, listener := range t.listeners {
		addr := listener.Addr().String()
		host, port, _ := net.SplitHostPort(addr)
		portNum, _ := strconv.ParseUint(port, 10, 16)

		nodes = append(nodes, &ipb.Node{
			Id:    name,
			Roles: &cpb.NodeRoles{ConsensusMember: &cpb.NodeRoles_ConsensusMember{}},
			Status: &cpb.NodeStatus{
				ExternalAddress: host,
				RunningCurator: &cpb.NodeStatus_RunningCurator{
					Port: int32(portNum),
				},
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

// GetCurrentLeader returns this fake cluster's current leader, based on
// thisNode and leader from the struct.
func (t *fakeCuratorClusterAware) GetCurrentLeader(req *ipb.GetCurrentLeaderRequest, srv ipb.CuratorLocal_GetCurrentLeaderServer) error {
	ctx := srv.Context()

	t.mu.Lock()
	leaderName := t.leader
	thisNode := t.thisNode
	t.mu.Unlock()
	leader := t.listeners[leaderName]

	host, port, _ := net.SplitHostPort(leader.Addr().String())
	portNum, _ := strconv.ParseUint(port, 10, 16)
	srv.Send(&ipb.GetCurrentLeaderResponse{
		LeaderNodeId: leaderName,
		LeaderHost:   host,
		LeaderPort:   int32(portNum),
		ThisNodeId:   thisNode,
	})
	<-ctx.Done()
	return ctx.Err()
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
	eph := util.NewEphemeralClusterCredentials(t, numCurators)

	listeners := make([]net.Listener, numCurators)
	for i := 0; i < numCurators; i++ {
		lis, err := net.Listen("tcp", "")
		if err != nil {
			t.Fatalf("Listen failed: %v", err)
		}
		listeners[i] = lis
	}

	// Make fakeCuratorClusterAware implementations for every node.
	listenerMap := make(map[string]net.Listener)
	for i, lis := range listeners {
		log.Printf("Test listener: %s", lis.Addr().String())
		name := eph.Nodes[i].ID()
		listenerMap[name] = lis
	}
	impls := make([]*fakeCuratorClusterAware, numCurators)

	for i := 0; i < numCurators; i++ {
		impls[i] = &fakeCuratorClusterAware{
			listeners: listenerMap,
			leader:    eph.Nodes[0].ID(),
			thisNode:  eph.Nodes[i].ID(),
		}
	}

	// Make gRPC servers for every node.
	servers := make([]*grpc.Server, numCurators)
	for i := 0; i < numCurators; i++ {
		i := i
		ss := rpc.ServerSecurity{
			NodeCredentials: eph.Nodes[i],
		}
		servers[i] = grpc.NewServer(ss.GRPCOptions(nil)...)
		ipb.RegisterCuratorServer(servers[i], impls[i])
		apb.RegisterAAAServer(servers[i], impls[i])
		apb.RegisterManagementServer(servers[i], impls[i])
		ipb.RegisterCuratorLocalServer(servers[i], impls[i])
		go func() {
			if err := servers[i].Serve(listeners[i]); err != nil {
				t.Fatalf("GRPC serve failed: %v", err)
			}
		}()

		defer listeners[i].Close()
		defer servers[i].Stop()
	}

	// Create our DUT resolver.
	r := New(ctx)
	r.logger = func(f string, args ...interface{}) {
		log.Printf(f, args...)
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{eph.Manager},
		InsecureSkipVerify: true,
	})
	cl, err := grpc.Dial("metropolis:///control", grpc.WithTransportCredentials(creds), grpc.WithResolvers(r))
	if err != nil {
		t.Fatalf("Could not dial: %v", err)
	}

	// Test logic follows.

	// Add first node to bootstrap node information from.
	r.AddEndpoint(nodeAtListener(listeners[0]))

	// The first node should be answering.
	cpl := ipb.NewCuratorLocalClient(cl)
	srv, err := cpl.GetCurrentLeader(ctx, &ipb.GetCurrentLeaderRequest{})
	if err != nil {
		t.Fatalf("GetCurrentLeader: %v", err)
	}
	leader, err := srv.Recv()
	if err != nil {
		t.Fatalf("GetCurrentLeader.Recv: %v", err)
	}
	if want, got := eph.Nodes[0].ID(), leader.ThisNodeId; want != got {
		t.Fatalf("Expected node %q to answer (current leader), got answer from %q", want, got)
	}

	// Wait for all curators to be picked up by the resolver.
	bo := backoff.NewExponentialBackOff()
	bo.MaxInterval = time.Second
	bo.MaxElapsedTime = 10 * time.Second
	err = backoff.Retry(func() error {
		req := &request{
			cmg: &requestCuratorMapGet{
				resC: make(chan *curatorMap),
			},
		}
		r.reqC <- req
		cm := <-req.cmg.resC
		if len(cm.curators) == 3 {
			return nil
		}
		return fmt.Errorf("have %d leaders, wanted 3", len(cm.curators))
	}, bo)
	if err != nil {
		t.Fatal(err)
	}

	// Move leadership to second node, resolver should follow.
	servers[0].Stop()
	for _, impl := range impls {
		impl.mu.Lock()
	}
	for _, impl := range impls {
		impl.leader = eph.Nodes[1].ID()
	}
	for _, impl := range impls {
		impl.mu.Unlock()
	}

	// Give it a few attempts. This isn't time bound, we _expect_ the resolver to
	// move over quickly to the correct leader.
	for i := 0; i < 3; i += 1 {
		srv, err = cpl.GetCurrentLeader(ctx, &ipb.GetCurrentLeaderRequest{})
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		leader, err = srv.Recv()
		if err != nil {
			continue
		}
		if want, got := eph.Nodes[1].ID(), leader.ThisNodeId; want != got {
			t.Fatalf("Expected node %q to answer (current leader), got answer from %q", want, got)
		}
		break
	}
	if err != nil {
		t.Fatalf("GetCurrentLeader after leadership change: %v", err)
	}
}
