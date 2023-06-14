// Package roleserve implements the roleserver/“Role Server”.
//
// The Role Server runs on every node and is responsible for running all of the
// node's role dependant services, like the control plane (Consensus/etcd and
// Curator) and Kubernetes. It watches the node roles as assigned by the
// cluster's curator, updates the status of the node within the curator, and
// spawns on-demand services.
//
//	.-----------.          .--------.  Watches  .------------.
//	| Cluster   |--------->| Role   |<----------| Node Roles |
//	| Enrolment | Provides | Server |  Updates  '------------'
//	'-----------'   Data   |        |----.      .-------------.
//	                       '--------'    '----->| Node Status |
//	                  Spawns |    | Spawns      '-------------'
//	                   .-----'    '-----.
//	                   V                V
//	               .-----------. .------------.
//	               | Consensus | | Kubernetes |
//	               | & Curator | |            |
//	               '-----------' '------------'
//
// The internal state of the Role Server (eg. status of services, input from
// Cluster Enrolment, current node roles as retrieved from the cluster) is
// stored as in-memory Event Value variables, with some of them being exposed
// externally for other services to consume (ie. ones that wish to depend on
// some information managed by the Role Server but which do not need to be
// spawned on demand by the Role Server). These Event Values and code which acts
// upon them form a reactive/dataflow-driven model which drives the Role Server
// logic forward.
//
// The Role Server also has to handle the complex bootstrap problem involved in
// simultaneously accessing the control plane (for node roles and other cluster
// data) while maintaining (possibly the only one in the cluster) control plane
// instance. This problem is resolved by using the RPC resolver package which
// allows dynamic reconfiguration of endpoints as the cluster is running.
package roleserve

import (
	"context"
	"crypto/ed25519"

	common "source.monogon.dev/metropolis/node"
	"source.monogon.dev/metropolis/node/core/clusternet"
	"source.monogon.dev/metropolis/node/core/curator"
	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
	"source.monogon.dev/metropolis/node/core/rpc/resolver"
	"source.monogon.dev/metropolis/pkg/event/memory"
	"source.monogon.dev/metropolis/pkg/logtree"
	"source.monogon.dev/metropolis/pkg/supervisor"

	cpb "source.monogon.dev/metropolis/proto/common"
)

// Config is the configuration of the role server.
type Config struct {
	// StorageRoot is a handle to access all of the Node's storage. This is needed
	// as the roleserver spawns complex workloads like Kubernetes which need access
	// to a broad range of storage.
	StorageRoot *localstorage.Root

	// Network is a handle to the network service, used by workloads.
	Network *network.Service

	// resolver is the main, long-lived, authenticated cluster resolver that is used
	// for all subsequent gRPC calls by the subordinates of the roleserver. It is
	// created early in the roleserver lifecycle, and is seeded with node
	// information from the ProvideXXX methods.
	Resolver *resolver.Resolver

	LogTree *logtree.LogTree
}

// Service is the roleserver/“Role Server” service. See the package-level
// documentation for more details.
type Service struct {
	Config

	KubernetesStatus      memory.Value[*KubernetesStatus]
	bootstrapData         memory.Value[*bootstrapData]
	localRoles            memory.Value[*cpb.NodeRoles]
	podNetwork            memory.Value[*clusternet.Prefixes]
	clusterDirectorySaved memory.Value[bool]
	localControlPlane     memory.Value[*localControlPlane]
	CuratorConnection     memory.Value[*curatorConnection]

	controlPlane *workerControlPlane
	statusPush   *workerStatusPush
	heartbeat    *workerHeartbeat
	kubernetes   *workerKubernetes
	rolefetch    *workerRoleFetch
	nodeMgmt     *workerNodeMgmt
	clusternet   *workerClusternet
	hostsfile    *workerHostsfile
	metrics      *workerMetrics
}

// New creates a Role Server services from a Config.
func New(c Config) *Service {
	s := &Service{
		Config: c,
	}
	s.controlPlane = &workerControlPlane{
		storageRoot: s.StorageRoot,

		bootstrapData: &s.bootstrapData,
		localRoles:    &s.localRoles,
		resolver:      s.Resolver,

		localControlPlane: &s.localControlPlane,
		curatorConnection: &s.CuratorConnection,
	}

	s.statusPush = &workerStatusPush{
		network: s.Network,

		curatorConnection:     &s.CuratorConnection,
		localControlPlane:     &s.localControlPlane,
		clusterDirectorySaved: &s.clusterDirectorySaved,
	}

	s.heartbeat = &workerHeartbeat{
		network: s.Network,

		curatorConnection: &s.CuratorConnection,
	}

	s.kubernetes = &workerKubernetes{
		network:     s.Network,
		storageRoot: s.StorageRoot,

		localRoles:        &s.localRoles,
		localControlPlane: &s.localControlPlane,
		curatorConnection: &s.CuratorConnection,

		kubernetesStatus: &s.KubernetesStatus,
		podNetwork:       &s.podNetwork,
	}

	s.rolefetch = &workerRoleFetch{
		curatorConnection: &s.CuratorConnection,

		localRoles: &s.localRoles,
	}

	s.nodeMgmt = &workerNodeMgmt{
		curatorConnection: &s.CuratorConnection,
		logTree:           s.LogTree,
	}

	s.clusternet = &workerClusternet{
		storageRoot: s.StorageRoot,

		curatorConnection: &s.CuratorConnection,
		podNetwork:        &s.podNetwork,
		network:           s.Network,
	}

	s.hostsfile = &workerHostsfile{
		storageRoot:           s.StorageRoot,
		network:               s.Network,
		curatorConnection:     &s.CuratorConnection,
		clusterDirectorySaved: &s.clusterDirectorySaved,
	}

	s.metrics = &workerMetrics{
		curatorConnection: &s.CuratorConnection,
	}

	return s
}

func (s *Service) ProvideBootstrapData(privkey ed25519.PrivateKey, iok, cuk, nuk, jkey []byte, icc *curator.Cluster, tpmUsage cpb.NodeTPMUsage) {
	pubkey := privkey.Public().(ed25519.PublicKey)
	nid := identity.NodeID(pubkey)

	// This is the first time we have the node ID, tell the resolver that it's
	// available on the loopback interface.
	s.Resolver.AddOverride(nid, resolver.NodeByHostPort("127.0.0.1", uint16(common.CuratorServicePort)))
	s.Resolver.AddEndpoint(resolver.NodeByHostPort("127.0.0.1", uint16(common.CuratorServicePort)))

	s.bootstrapData.Set(&bootstrapData{
		nodePrivateKey:              privkey,
		initialOwnerKey:             iok,
		clusterUnlockKey:            cuk,
		nodeUnlockKey:               nuk,
		nodePrivateJoinKey:          jkey,
		initialClusterConfiguration: icc,
		nodeTPMUsage:                tpmUsage,
	})
}

func (s *Service) ProvideRegisterData(credentials identity.NodeCredentials, directory *cpb.ClusterDirectory) {
	// This is the first time we have the node ID, tell the resolver that it's
	// available on the loopback interface.
	s.Resolver.AddOverride(credentials.ID(), resolver.NodeByHostPort("127.0.0.1", uint16(common.CuratorServicePort)))
	// Also tell the resolver about all the existing nodes in the cluster we just
	// registered into. The directory passed here was used to issue the initial
	// Register call, which means at least one of the nodes was running the control
	// plane and thus can be used to seed the rest of the resolver.
	for _, n := range directory.Nodes {
		for _, addr := range n.Addresses {
			s.Resolver.AddEndpoint(resolver.NodeAtAddressWithDefaultPort(addr.Host))
		}
	}

	s.CuratorConnection.Set(newCuratorConnection(&credentials, s.Resolver))
}

func (s *Service) ProvideJoinData(credentials identity.NodeCredentials, directory *cpb.ClusterDirectory) {
	// This is the first time we have the node ID, tell the resolver that it's
	// available on the loopback interface.
	s.Resolver.AddOverride(credentials.ID(), resolver.NodeByHostPort("127.0.0.1", uint16(common.CuratorServicePort)))
	// Also tell the resolver about all the existing nodes in the cluster we just
	// joined into. The directory passed here was used to issue the initial
	// Join call, which means at least one of the nodes was running the control
	// plane and thus can be used to seed the rest of the resolver.
	for _, n := range directory.Nodes {
		for _, addr := range n.Addresses {
			s.Resolver.AddEndpoint(resolver.NodeAtAddressWithDefaultPort(addr.Host))
		}
	}

	s.CuratorConnection.Set(newCuratorConnection(&credentials, s.Resolver))
	s.clusterDirectorySaved.Set(true)
}

// Run the Role Server service, which uses intermediary workload launchers to
// start/stop subordinate services as the Node's roles change.
func (s *Service) Run(ctx context.Context) error {
	supervisor.Run(ctx, "controlplane", s.controlPlane.run)
	supervisor.Run(ctx, "kubernetes", s.kubernetes.run)
	supervisor.Run(ctx, "statuspush", s.statusPush.run)
	supervisor.Run(ctx, "heartbeat", s.heartbeat.run)
	supervisor.Run(ctx, "rolefetch", s.rolefetch.run)
	supervisor.Run(ctx, "nodemgmt", s.nodeMgmt.run)
	supervisor.Run(ctx, "clusternet", s.clusternet.run)
	supervisor.Run(ctx, "hostsfile", s.hostsfile.run)
	supervisor.Run(ctx, "metrics", s.metrics.run)
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	<-ctx.Done()
	return ctx.Err()
}
