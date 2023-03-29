package roleserve

import (
	"context"
	"crypto/ed25519"

	"google.golang.org/grpc"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
	"source.monogon.dev/metropolis/pkg/event"
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
	// localConsensus and localCurator are set by the Control Plane Worker when this
	// node runs control plane services.
	localConsensus *consensus.Service
	localCurator   *curator.Service
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
	// resolver will be used to dial the cluster via DialCurator().
	resolver *resolver.Resolver
}

// GetNodeID returns the Node ID of the locally running node whenever available.
// NodeIDs are available early on in the node startup process and are guaranteed
// to never change at runtime. The watcher will then block all further Get calls
// until new information is available. This method should only be used if
// GetNodeID is the only method ran on the watcher.
func GetNodeID(ctx context.Context, watcher event.Watcher[*ClusterMembership]) (string, error) {
	for {
		cm, err := watcher.Get(ctx)
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
func FilterHome() event.GetOption[*ClusterMembership] {
	return event.Filter(func(cm *ClusterMembership) bool {
		if cm.credentials == nil {
			return false
		}
		if cm.remoteCurators == nil {
			return false
		}
		return true
	})
}

// DialCurator returns an authenticated gRPC client connection to the Curator
// using the long-lived roleserver cluster resolver. RPCs will automatically be
// forwarded to the current control plane leader, and this gRPC client
// connection can be used long-term by callers.
func (m *ClusterMembership) DialCurator() (*grpc.ClientConn, error) {
	// Always make sure the resolver has fresh data about curators, both local and
	// remote. This would be better done only when ClusterMembership is set() with
	// new data, but that would require a bit of a refactor.
	//
	// TODO(q3k): take care of the above, possibly when the roleserver is made more generic.
	if m.localConsensus != nil {
		m.resolver.AddEndpoint(resolver.NodeByHostPort("127.0.0.1", uint16(common.CuratorServicePort)))
	}
	for _, n := range m.remoteCurators.Nodes {
		for _, addr := range n.Addresses {
			m.resolver.AddEndpoint(resolver.NodeByHostPort(addr.Host, uint16(common.CuratorServicePort)))
		}
	}
	creds := rpc.NewAuthenticatedCredentials(m.credentials.TLSCredentials(), rpc.WantRemoteCluster(m.credentials.ClusterCA()))
	return grpc.Dial(resolver.MetropolisControlAddress, grpc.WithTransportCredentials(creds), grpc.WithResolvers(m.resolver))
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
