package clusternet

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/supervisor"

	apb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// testCurator is a shim Curator implementation that serves pending Watch
// requests based on data submitted to a channel.
type testCurator struct {
	apb.UnimplementedCuratorServer

	watchC    chan *apb.WatchEvent
	updateReq memory.Value[*apb.UpdateNodeClusterNetworkingRequest]
}

// Watch implements a minimum Watch which just returns all nodes at once.
func (t *testCurator) Watch(_ *apb.WatchRequest, srv apb.Curator_WatchServer) error {
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

func (t *testCurator) UpdateNodeClusterNetworking(ctx context.Context, req *apb.UpdateNodeClusterNetworkingRequest) (*apb.UpdateNodeClusterNetworkingResponse, error) {
	t.updateReq.Set(req)
	return &apb.UpdateNodeClusterNetworkingResponse{}, nil
}

// nodeWithPrefix submits a given node/key/address with prefixes to the Watch
// event channel.
func (t *testCurator) nodeWithPrefixes(key wgtypes.Key, id, address string, prefixes ...string) {
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
	}
	t.watchC <- &apb.WatchEvent{
		Nodes: []*apb.Node{
			n,
		},
	}
}

// deleteNode submits a given node for deletion to the Watch event channel.
func (t *testCurator) deleteNode(id string) {
	t.watchC <- &apb.WatchEvent{
		NodeTombstones: []*apb.WatchEvent_NodeTombstone{
			{
				NodeId: id,
			},
		},
	}
}

// makeTestCurator returns a working testCurator alongside a grpc connection to
// it.
func makeTestCurator(t *testing.T) (*testCurator, *grpc.ClientConn) {
	cur := &testCurator{
		watchC: make(chan *apb.WatchEvent),
	}

	srv := grpc.NewServer()
	apb.RegisterCuratorServer(srv, cur)
	externalLis := bufconn.Listen(1024 * 1024)
	go func() {
		if err := srv.Serve(externalLis); err != nil {
			t.Fatalf("GRPC serve failed: %v", err)
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

// fakeWireguard implements wireguard while keeping peer information internally.
type fakeWireguard struct {
	k wgtypes.Key

	muNodes        sync.Mutex
	nodes          map[string]*node
	failNextUpdate bool
}

func (f *fakeWireguard) ensureOnDiskKey(_ *localstorage.DataKubernetesClusterNetworkingDirectory) error {
	f.k, _ = wgtypes.GeneratePrivateKey()
	return nil
}

func (f *fakeWireguard) setup(clusterNet *net.IPNet) error {
	f.muNodes.Lock()
	defer f.muNodes.Unlock()
	f.nodes = make(map[string]*node)
	return nil
}

func (f *fakeWireguard) configurePeers(nodes []*node) error {
	f.muNodes.Lock()
	defer f.muNodes.Unlock()
	if f.failNextUpdate {
		f.failNextUpdate = false
		return fmt.Errorf("synthetic test failure")
	}
	for _, n := range nodes {
		f.nodes[n.id] = n
	}
	return nil
}

func (f *fakeWireguard) unconfigurePeer(n *node) error {
	f.muNodes.Lock()
	defer f.muNodes.Unlock()
	delete(f.nodes, n.id)
	return nil
}

func (f *fakeWireguard) key() wgtypes.Key {
	return f.k
}

func (f *fakeWireguard) close() {
}

// TestClusternetBasic exercises clusternet with a fake curator and fake
// wireguard, trying to exercise as many edge cases as possible.
func TestClusternetBasic(t *testing.T) {
	key1, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	key2, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	cur, cl := makeTestCurator(t)
	defer cl.Close()
	curator := apb.NewCuratorClient(cl)

	var nval memory.Value[*network.Status]

	var podNetwork memory.Value[*Prefixes]
	wg := &fakeWireguard{}
	svc := Service{
		Curator: curator,
		ClusterNet: net.IPNet{
			IP:   net.IP([]byte{10, 10, 0, 0}),
			Mask: net.IPv4Mask(255, 255, 0, 0),
		},
		DataDirectory:             nil,
		LocalKubernetesPodNetwork: &podNetwork,
		Network:                   &nval,

		wg: wg,
	}
	supervisor.TestHarness(t, svc.Run)

	checkState := func(nodes map[string]*node) error {
		t.Helper()
		wg.muNodes.Lock()
		defer wg.muNodes.Unlock()
		for nid, n := range nodes {
			n2, ok := wg.nodes[nid]
			if !ok {
				return fmt.Errorf("node %q missing in programmed peers", nid)
			}
			if n2.pubkey != n.pubkey {
				return fmt.Errorf("node %q pubkey mismatch: %q in programmed peers, %q wanted", nid, n2.pubkey, n.pubkey)
			}
			if n2.address != n.address {
				return fmt.Errorf("node %q address mismatch: %q in programmed peers, %q wanted", nid, n2.address, n.address)
			}
			p := strings.Join(n.prefixes, ",")
			p2 := strings.Join(n2.prefixes, ",")
			if p != p2 {
				return fmt.Errorf("node %q prefixes mismatch: %v in programmed peers, %v wanted", nid, n2.prefixes, n.prefixes)
			}
		}
		for nid, _ := range wg.nodes {
			if _, ok := nodes[nid]; !ok {
				return fmt.Errorf("node %q present in programmed peers", nid)
			}
		}
		return nil
	}

	assertStateEventual := func(nodes map[string]*node) {
		t.Helper()
		deadline := time.Now().Add(5 * time.Second)
		for {
			err := checkState(nodes)
			if err == nil {
				break
			}
			if time.Now().After(deadline) {
				t.Error(err)
				return
			}
		}

	}

	// Start with a single node.
	cur.nodeWithPrefixes(key1, "metropolis-fake-1", "1.2.3.4")
	assertStateEventual(map[string]*node{
		"metropolis-fake-1": {
			pubkey:   key1.PublicKey().String(),
			address:  "1.2.3.4",
			prefixes: nil,
		},
	})
	// Change the node's peer address.
	cur.nodeWithPrefixes(key1, "metropolis-fake-1", "1.2.3.5")
	assertStateEventual(map[string]*node{
		"metropolis-fake-1": {
			pubkey:   key1.PublicKey().String(),
			address:  "1.2.3.5",
			prefixes: nil,
		},
	})
	// Add another node.
	cur.nodeWithPrefixes(key2, "metropolis-fake-2", "1.2.3.6")
	assertStateEventual(map[string]*node{
		"metropolis-fake-1": {
			pubkey:   key1.PublicKey().String(),
			address:  "1.2.3.5",
			prefixes: nil,
		},
		"metropolis-fake-2": {
			pubkey:   key2.PublicKey().String(),
			address:  "1.2.3.6",
			prefixes: nil,
		},
	})
	// Add some prefixes to both nodes, but fail the next configurePeers call.
	wg.muNodes.Lock()
	wg.failNextUpdate = true
	wg.muNodes.Unlock()
	cur.nodeWithPrefixes(key1, "metropolis-fake-1", "1.2.3.5", "10.100.10.0/24", "10.100.20.0/24")
	cur.nodeWithPrefixes(key2, "metropolis-fake-2", "1.2.3.6", "10.100.30.0/24", "10.100.40.0/24")
	assertStateEventual(map[string]*node{
		"metropolis-fake-1": {
			pubkey:  key1.PublicKey().String(),
			address: "1.2.3.5",
			prefixes: []string{
				"10.100.10.0/24", "10.100.20.0/24",
			},
		},
		"metropolis-fake-2": {
			pubkey:  key2.PublicKey().String(),
			address: "1.2.3.6",
			prefixes: []string{
				"10.100.30.0/24", "10.100.40.0/24",
			},
		},
	})
	// Delete one of the nodes.
	cur.deleteNode("metropolis-fake-1")
	assertStateEventual(map[string]*node{
		"metropolis-fake-2": {
			pubkey:  key2.PublicKey().String(),
			address: "1.2.3.6",
			prefixes: []string{
				"10.100.30.0/24", "10.100.40.0/24",
			},
		},
	})
}

// TestWireguardImplementation makes sure localWireguard behaves as expected.
func TestWireguardIntegration(t *testing.T) {
	if os.Getenv("IN_KTEST") != "true" {
		t.Skip("Not in ktest")
	}

	root := &localstorage.Root{}
	tmp, err := os.MkdirTemp("", "clusternet")
	if err != nil {
		t.Fatal(err)
	}
	err = declarative.PlaceFS(root, tmp)
	if err != nil {
		t.Fatal(err)
	}
	os.MkdirAll(root.Data.Kubernetes.ClusterNetworking.FullPath(), 0700)
	wg := &localWireguard{}

	// Ensure key once and make note of it.
	if err := wg.ensureOnDiskKey(&root.Data.Kubernetes.ClusterNetworking); err != nil {
		t.Fatalf("Could not ensure wireguard key: %v", err)
	}
	key := wg.key().String()
	// Do it again, and make sure the key hasn't changed.
	wg = &localWireguard{}
	if err := wg.ensureOnDiskKey(&root.Data.Kubernetes.ClusterNetworking); err != nil {
		t.Fatalf("Could not ensure wireguard key second time: %v", err)
	}
	if want, got := key, wg.key().String(); want != got {
		t.Fatalf("Key changed, was %q, became %q", want, got)
	}

	// Setup the interface.
	cnet := net.IPNet{
		IP:   net.IP([]byte{10, 10, 0, 0}),
		Mask: net.IPv4Mask(255, 255, 0, 0),
	}
	if err := wg.setup(&cnet); err != nil {
		t.Fatalf("Failed to setup interface: %v", err)
	}
	// Do it again.
	wg.close()
	if err := wg.setup(&cnet); err != nil {
		t.Fatalf("Failed to setup interface second time: %v", err)
	}

	// Check that the key and listen port are configured correctly.
	wgClient, err := wgctrl.New()
	if err != nil {
		t.Fatalf("Failed to create wireguard client: %v", err)
	}
	wgDev, err := wgClient.Device(clusterNetDeviceName)
	if err != nil {
		t.Fatalf("Failed to connect to netlink's WireGuard config endpoint: %v", err)
	}
	if want, got := key, wgDev.PrivateKey.String(); want != got {
		t.Errorf("Wireguard key mismatch, wanted %q, got %q", want, got)
	}
	if want, got := int(common.WireGuardPort), wgDev.ListenPort; want != got {
		t.Errorf("Wireguard port mismatch, wanted %d, got %d", want, got)
	}

	// Add some peers and check that we got them.
	pkeys := make([]wgtypes.Key, 2)
	pkeys[0], err = wgtypes.GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	pkeys[1], err = wgtypes.GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	err = wg.configurePeers([]*node{
		{
			pubkey:  pkeys[0].PublicKey().String(),
			address: "10.100.0.1",
			prefixes: []string{
				"10.0.0.0/24",
				"10.0.1.0/24",
			},
		},
		{
			pubkey:  pkeys[1].PublicKey().String(),
			address: "10.100.1.1",
			prefixes: []string{
				"10.1.0.0/24",
				"10.1.1.0/24",
			},
		},
	})
	if err != nil {
		t.Fatalf("Configuring peers failed: %v", err)
	}

	wgDev, err = wgClient.Device(clusterNetDeviceName)
	if err != nil {
		t.Fatalf("Failed to connect to netlink's WireGuard config endpoint: %v", err)
	}
	if want, got := 2, len(wgDev.Peers); want != got {
		t.Errorf("Wanted %d peers, got %d", want, got)
	} else {
		for i := 0; i < 2; i++ {
			if want, got := pkeys[i].PublicKey().String(), wgDev.Peers[i].PublicKey.String(); want != got {
				t.Errorf("Peer %d should have key %q, got %q", i, want, got)
			}
			if want, got := fmt.Sprintf("10.100.%d.1:%s", i, common.WireGuardPort.PortString()), wgDev.Peers[i].Endpoint.String(); want != got {
				t.Errorf("Peer %d should have endpoint %q, got %q", i, want, got)
			}
			if want, got := 2, len(wgDev.Peers[i].AllowedIPs); want != got {
				t.Errorf("Peer %d should have %d peers, got %d", i, want, got)
			} else {
				for j := 0; j < 2; j++ {
					if want, got := fmt.Sprintf("10.%d.%d.0/24", i, j), wgDev.Peers[i].AllowedIPs[j].String(); want != got {
						t.Errorf("Peer %d should have allowed ip %d %q, got %q", i, j, want, got)
					}
				}
			}
		}
	}

	// Update one of the peers and check that things got applied.
	err = wg.configurePeers([]*node{
		{
			pubkey:  pkeys[0].PublicKey().String(),
			address: "10.100.0.3",
			prefixes: []string{
				"10.0.0.0/24",
			},
		},
	})
	if err != nil {
		t.Fatalf("Failed to connect to netlink's WireGuard config endpoint: %v", err)
	}
	wgDev, err = wgClient.Device(clusterNetDeviceName)
	if err != nil {
		t.Fatalf("Failed to connect to netlink's WireGuard config endpoint: %v", err)
	}
	if want, got := 2, len(wgDev.Peers); want != got {
		t.Errorf("Wanted %d peers, got %d", want, got)
	} else {
		if want, got := pkeys[0].PublicKey().String(), wgDev.Peers[0].PublicKey.String(); want != got {
			t.Errorf("Peer 0 should have key %q, got %q", want, got)
		}
		if want, got := fmt.Sprintf("10.100.0.3:%s", common.WireGuardPort.PortString()), wgDev.Peers[0].Endpoint.String(); want != got {
			t.Errorf("Peer 0 should have endpoint %q, got %q", want, got)
		}
		if want, got := 1, len(wgDev.Peers[0].AllowedIPs); want != got {
			t.Errorf("Peer 0 should have %d peers, got %d", want, got)
		} else {
			if want, got := "10.0.0.0/24", wgDev.Peers[0].AllowedIPs[0].String(); want != got {
				t.Errorf("Peer 0 should have allowed ip 0 %q, got %q", want, got)
			}
		}
	}

	// Remove one of the peers and make sure it's gone.
	err = wg.unconfigurePeer(&node{
		pubkey: pkeys[0].PublicKey().String(),
	})
	if err != nil {
		t.Fatalf("Failed to unconfigure peer: %v", err)
	}
	err = wg.unconfigurePeer(&node{
		pubkey: pkeys[0].PublicKey().String(),
	})
	if err != nil {
		t.Fatalf("Failed to unconfigure peer a second time: %v", err)
	}
	wgDev, err = wgClient.Device(clusterNetDeviceName)
	if err != nil {
		t.Fatalf("Failed to connect to netlink's WireGuard config endpoint: %v", err)
	}
	if want, got := 1, len(wgDev.Peers); want != got {
		t.Errorf("Wanted %d peer, got %d", want, got)
	}
}
