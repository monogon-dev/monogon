// hostsfile implements a service which owns and writes all node-local
// files/interfaces used by the system to resolve the local node's name and the
// names of other nodes in the cluster:
//
// 1. All cluster node names are written into /etc/hosts for DNS resolution.
// 2. The local node's name is written into /etc/machine-id.
// 3. The local node's name is set as the UNIX hostname of the machine (via the
//    sethostname call).
//
// The hostsfile Service can start up in two modes: with cluster connectivity
// and without cluster connectivity. Without cluster connectivity, only
// information about the current node (as retrieved from the network service)
// will be used to populate local data. In cluster mode, information about other
// nodes is also used.
package hostsfile

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"fmt"
	"net"
	"sort"

	"golang.org/x/sys/unix"
	"google.golang.org/grpc"

	ipb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/pkg/supervisor"
)

type Config struct {
	// NodePublicKey is the local node's public key, used for calculating the local
	// node's ID/name. This is used instead of passing identity.Node as the
	// hostsfile service starts very early, earlier than a node might have its' full
	// identity (which includes cluster certificates), but will have its public key.
	NodePublicKey ed25519.PublicKey
	// Network is a handle to the Network service, used to update the hostsfile
	// service with information about the local node's external IP address.
	Network *network.Service
	// Ephemeral is the root of the ephemeral storage of the node, into which the
	// service will write its managed files.
	Ephemeral *localstorage.EphemeralDirectory
	// ClusterDialer is an optional function that the service will call to establish
	// connectivity to cluster services. If given, the local files will be augmented
	// with information about all the cluster's nodes (ie. the local node will be
	// able to resolve all nodes in the cluster by DNS name), otherwise only the
	// local node's addresse/name will be populated.
	//
	// MVP: this should instead be an event value (maybe a roleserver bit?) which
	// allows the service to access the cluster data gracefully, instead of it
	// having to be restarted when cluster connectivity becomes available.
	ClusterDialer ClusterDialer
}

// Service is the hostsfile service instance. See package-level documentation
// for more information.
type Service struct {
	Config

	// localC is a channel populated by the local sub-runnable with the newest
	// available information about the local node's address. It is automatically
	// created and closed by Run.
	localC chan string
	// clusterC is a channel populated by the cluster sub-runnable with the newest
	// available information about the cluster nodes. It is automatically created and
	// closed by Run.
	clusterC chan nodeMap
}

type ClusterDialer func(ctx context.Context) (*grpc.ClientConn, error)

// nodeMap is a map from node ID (effectively DNS name) to node IP address.
type nodeMap map[string]string

// hosts generates a complete /etc/hosts file based on the contents of the
// nodeMap. Apart from the addresses in the nodeMap, entries for localhost
// pointing to 127.0.0.1 and ::1 will also be generated.
func (m nodeMap) hosts(ctx context.Context) []byte {
	var nodeIdsSorted []string
	for k, _ := range m {
		nodeIdsSorted = append(nodeIdsSorted, k)
	}
	sort.Slice(nodeIdsSorted, func(i, j int) bool {
		return nodeIdsSorted[i] < nodeIdsSorted[j]
	})

	lines := [][]byte{
		[]byte("127.0.0.1 localhost"),
		[]byte("::1 localhost"),
	}
	for _, nid := range nodeIdsSorted {
		addr := m[nid]
		line := fmt.Sprintf("%s %s", addr, nid)
		supervisor.Logger(ctx).Infof("Hosts entry: %s", line)
		lines = append(lines, []byte(line))
	}
	lines = append(lines, []byte(""))

	return bytes.Join(lines, []byte("\n"))
}

func (s *Service) Run(ctx context.Context) error {
	s.localC = make(chan string)
	defer close(s.localC)
	s.clusterC = make(chan nodeMap)
	defer close(s.clusterC)

	nodeID := identity.NodeID(s.NodePublicKey)

	if err := supervisor.Run(ctx, "local", s.runLocal); err != nil {
		return err
	}

	if s.ClusterDialer != nil {
		supervisor.Logger(ctx).Infof("Running with cluster support.")
		if err := supervisor.Run(ctx, "cluster", s.runCluster); err != nil {
			return err
		}
	} else {
		supervisor.Logger(ctx).Infof("Running without cluster support, only local node will be present.")
	}

	// Immediately update machine-id and hostname, we don't need network addresses
	// for that.
	if err := s.Ephemeral.MachineID.Write([]byte(nodeID), 0644); err != nil {
		return fmt.Errorf("failed to write /ephemeral/machine-id: %w", err)
	}
	if err := unix.Sethostname([]byte(nodeID)); err != nil {
		return fmt.Errorf("failed to set runtime hostname: %w", err)
	}
	// Immediately write an /etc/hosts just containing localhost, even if we don't
	// yet have a network address.
	nodes := make(nodeMap)
	if err := s.Ephemeral.Hosts.Write(nodes.hosts(ctx), 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", s.Ephemeral.Hosts.FullPath(), err)
	}

	supervisor.Signal(ctx, supervisor.SignalHealthy)
	// Update nodeMap in a loop, issuing writes/updates when any change occurred.
	for {
		changed := false
		select {
		case <-ctx.Done():
			return ctx.Err()
		case u := <-s.localC:
			// Ignore spurious updates.
			if nodes[nodeID] == u {
				break
			}
			supervisor.Logger(ctx).Infof("Got new local address: %s", u)
			nodes[nodeID] = u
			changed = true
		case u := <-s.clusterC:
			// Loop through the nodeMap from the cluster subrunnable, making note of what
			// changed. By design we don't care about any nodes disappearing from the
			// nodeMap: we'd rather keep stale data about nodes that don't exist any more,
			// as these might either be spurious or have a long tail of effectively still
			// being used by the local node for communications while the node gets fully
			// drained/disowned.
			//
			// MVP: we should at least log removed nodes.
			for id, addr := range u {
				// We're not interested in what the cluster thinks about our local node, as that
				// might be outdated (eg. when we haven't yet reported a new local address to
				// the cluster).
				if id == nodeID {
					continue
				}
				if nodes[id] == addr {
					continue
				}
				supervisor.Logger(ctx).Infof("Got new cluster address: %s is %s", id, addr)
				nodes[id] = addr
				changed = true
			}
		}

		if !changed {
			continue
		}

		supervisor.Logger(ctx).Infof("Updating hosts file: %d nodes", len(nodes))
		if err := s.Ephemeral.Hosts.Write(nodes.hosts(ctx), 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", s.Ephemeral.Hosts.FullPath(), err)
		}

		// Check that we are self-resolvable.
		if _, err := net.ResolveIPAddr("ip", nodeID); err != nil {
			supervisor.Logger(ctx).Errorf("Failed to self-resolve %q: %v", nodeID, err)
		}

	}
}

// runLocal updates s.localC with the IP address of the local node, as retrieved
// from the network service.
func (s *Service) runLocal(ctx context.Context) error {
	nw := s.Network.Watch()
	for {
		ns, err := nw.Get(ctx)
		if err != nil {
			return err
		}
		addr := ns.ExternalAddress.String()
		if addr != "" {
			s.localC <- addr
		}
	}
}

// runCluster updates s.clusterC with the IP addresses of cluster nodes, as
// retrieved from a Curator client from the ClusterDialer. The returned map
// reflects the up-to-date view of the cluster returned from the Curator Watch
// call, including any node deletions.
func (s *Service) runCluster(ctx context.Context) error {
	cl, err := s.Config.ClusterDialer(ctx)
	if err != nil {
		return fmt.Errorf("cluster dial failed: %w", err)
	}
	defer cl.Close()
	curator := ipb.NewCuratorClient(cl)

	w, err := curator.Watch(ctx, &ipb.WatchRequest{
		Kind: &ipb.WatchRequest_NodesInCluster_{
			NodesInCluster: &ipb.WatchRequest_NodesInCluster{},
		},
	})
	if err != nil {
		return fmt.Errorf("curator watch failed: %w", err)
	}

	nodes := make(nodeMap)
	for {
		ev, err := w.Recv()
		if err != nil {
			return fmt.Errorf("receive failed: %w", err)
		}
		for _, n := range ev.Nodes {
			if n.Status == nil || n.Status.ExternalAddress == "" {
				continue
			}
			nodes[n.Id] = n.Status.ExternalAddress
		}
		for _, t := range ev.NodeTombstones {
			delete(nodes, t.NodeId)
		}
		s.clusterC <- nodes
	}
}