package roleserve

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"net"

	"google.golang.org/grpc"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/pkg/event"
	"source.monogon.dev/metropolis/pkg/event/memory"
	cpb "source.monogon.dev/metropolis/proto/common"
)

// ClusterMembership is an Event Value structure used to keep state of the
// membership of this node in a cluster, the location of a working Curator API
// (local or remote) and the state of a locally running control plane.
//
// This amalgam of seemingly unrelated data is all required to have a single
// structure that can answer a seemingly trivial question: “Who am I and how do
// I contact a Curator?”.
//
// This structure is available to roleserver-internal workers (eg. the Kubernetes
// Worker Launcher and Updater) and to external code (eg. the Hostsfile
// service). It is also deeply intertwined with the Control Plane Worker which
// not only populates it when a Control Plane (and thus Curator) gets started,
// but also accesses it to pass over information about already known remote
// curators and to get the local node's identity.
type ClusterMembership struct {
	// localConsensus is set by the Control Plane Worker when this node runs control
	// plane services.
	localConsensus *consensus.Service
	// remoteCurators gets set by Cluster Enrolment code when Registering into a
	// cluster and gets propagated by the Control Plane Worker to maintain
	// connectivity to external Curators regardless of local curator health.
	//
	// TODO(q3k): also update this based on a live Cluster Directory from the
	// cluster.
	remoteCurators *cpb.ClusterDirectory
	// credentials is set whenever this node has full access to the Cluster and is
	// the a of credentials which can be used to perform authenticated (as the node)
	// access to the Curator.
	credentials *identity.NodeCredentials
	// pubkey is the public key of the local node, and is always set.
	pubkey ed25519.PublicKey
}

type ClusterMembershipValue struct {
	value memory.Value
}

func (c *ClusterMembershipValue) Watch() *ClusterMembershipWatcher {
	return &ClusterMembershipWatcher{
		w: c.value.Watch(),
	}
}

func (c *ClusterMembershipValue) set(v *ClusterMembership) {
	c.value.Set(v)
}

type ClusterMembershipWatcher struct {
	w event.Watcher
}

func (c *ClusterMembershipWatcher) Close() error {
	return c.w.Close()
}

func (c *ClusterMembershipWatcher) getAny(ctx context.Context) (*ClusterMembership, error) {
	v, err := c.w.Get(ctx)
	if err != nil {
		return nil, err
	}
	return v.(*ClusterMembership), nil
}

// GetNodeID returns the Node ID of the locally running node whenever available.
// NodeIDs are available early on in the node startup process and are guaranteed
// to never change at runtime. The watcher will then block all further Get calls
// until new information is available. This method should only be used if
// GetNodeID is the only method ran on the watcher.
func (c *ClusterMembershipWatcher) GetNodeID(ctx context.Context) (string, error) {
	for {
		cm, err := c.getAny(ctx)
		if err != nil {
			return "", err
		}
		if cm.pubkey != nil {
			return identity.NodeID(cm.pubkey), nil
		}
	}
}

// GetHome returns a ClusterMembership whenever the local node is HOME to a
// cluster (ie. whenever the node is fully a member of a cluster and can dial
// the cluster's Curator). See proto.common.ClusterState for more information
// about cluster states. The watcher will then block all futher Get calls until
// new information is available.
func (c *ClusterMembershipWatcher) GetHome(ctx context.Context) (*ClusterMembership, error) {
	for {
		cm, err := c.getAny(ctx)
		if err != nil {
			return nil, err
		}
		if cm.credentials == nil {
			continue
		}
		if cm.remoteCurators == nil {
			continue
		}
		return cm, nil
	}
}

// DialCurator returns an authenticated gRPC client connection to the Curator,
// either local or remote. No load balancing will be performed across local and
// remote curators, so if the local node starts running a local curator but old
// connections are still used, they will continue to target only remote
// curators. Same goes for local consensus being turned down - however, in this
// case, calls will error out and the client can be redialed on errors.
//
// It is thus recommended to only use DialCurator in short-lived contexts, and
// perform a GetHome/DialCurator process on any gRPC error. A smarter
// load-balancing/re-dialing client will be implemented in the future.
func (m *ClusterMembership) DialCurator() (*grpc.ClientConn, error) {
	// Dial first curator.
	// TODO(q3k): load balance
	if m.remoteCurators == nil || len(m.remoteCurators.Nodes) < 1 {
		return nil, fmt.Errorf("no curators available")
	}
	host := m.remoteCurators.Nodes[0].Addresses[0].Host
	addr := net.JoinHostPort(host, common.CuratorServicePort.PortString())
	creds := rpc.NewAuthenticatedCredentials(m.credentials.TLSCredentials(), m.credentials.ClusterCA())
	return grpc.Dial(addr, grpc.WithTransportCredentials(creds))
}

func (m *ClusterMembership) NodePubkey() ed25519.PublicKey {
	if m.pubkey == nil {
		// This shouldn't happen - it means a user got this structure too early or
		// constructed it from scratch.
		panic("node pubkey not available")
	}
	return m.pubkey
}

func (m *ClusterMembership) NodeID() string {
	return identity.NodeID(m.NodePubkey())
}
