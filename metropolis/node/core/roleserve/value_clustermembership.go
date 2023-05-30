package roleserve

import (
	"crypto/ed25519"

	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
	"source.monogon.dev/metropolis/pkg/event"
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
	// credentials is set whenever this node has full access to the Cluster and is
	// the a of credentials which can be used to perform authenticated (as the node)
	// access to the Curator.
	credentials *identity.NodeCredentials
	// pubkey is the public key of the local node, and is always set.
	pubkey ed25519.PublicKey
	// resolver will be used to dial the cluster via DialCurator().
	resolver *resolver.Resolver
}

// FilterHome returns a ClusterMembership whenever the local node is HOME to a
// cluster (ie. whenever the node is fully a member of a cluster and can dial
// the cluster's Curator). See proto.common.ClusterState for more information
// about cluster states. The watcher will then block all futher Get calls until
// new information is available.
func FilterHome() event.GetOption[*ClusterMembership] {
	return event.Filter(func(cm *ClusterMembership) bool {
		if cm.credentials == nil {
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
