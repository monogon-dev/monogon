package roleserve

import (
	"google.golang.org/grpc"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/rpc"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
	"source.monogon.dev/metropolis/node/kubernetes"
)

// localControlPlane is an internal EventValue structure which carries
// information about whether the node has a locally running consensus and curator
// service. When it does, the structure pointer inside the EventValue will be
// non-nil and its consensus and curator members will also be non-nil. If it
// doesn't, either the pointer inside the EventValue will be nil, or will carry
// nil pointers. Because of this, it is recommended to use the exists() method to
// check for consensus/curator presence.
type localControlPlane struct {
	consensus *consensus.Service
	curator   *curator.Service
}

func (l *localControlPlane) exists() bool {
	if l == nil {
		return false
	}
	if l.consensus == nil || l.curator == nil {
		return false
	}
	return true
}

// CuratorConnection carries information about the node having successfully
// established connectivity to its cluster's control plane.
//
// It carries inside it a single gRPC client connection which is built using the
// main roleserver resolver. This connection will automatically use any available
// curator, whether running locally or remotely.
//
// This structure should also be used by roleserver runnables that simply wish to
// access the node's credentials.
type CuratorConnection struct {
	Credentials *identity.NodeCredentials
	resolver    *resolver.Resolver
	conn        *grpc.ClientConn
}

func newCuratorConnection(creds *identity.NodeCredentials, res *resolver.Resolver) *CuratorConnection {
	c := rpc.NewAuthenticatedCredentials(creds.TLSCredentials(), rpc.WantRemoteCluster(creds.ClusterCA()))
	conn, err := grpc.Dial(resolver.MetropolisControlAddress, grpc.WithTransportCredentials(c), grpc.WithResolvers(res))
	if err != nil {
		// TOOD(q3k): triple check that Dial will not fail
		panic(err)
	}
	return &CuratorConnection{
		Credentials: creds,
		resolver:    res,
		conn:        conn,
	}
}

func (c *CuratorConnection) nodeID() string {
	return identity.NodeID(c.Credentials.PublicKey())
}

// KubernetesStatus is an Event Value structure populated by a running
// Kubernetes instance. It allows external services to access the Kubernetes
// Service whenever available (ie. enabled and started by the Role Server).
type KubernetesStatus struct {
	Controller *kubernetes.Controller
}
