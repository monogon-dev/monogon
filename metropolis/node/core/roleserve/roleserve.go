// package roleserve implements the roleserver/“Role Server”.
//
// The Role Server runs on every node and is responsible for running all of the
// node's role dependant services, like the control plane (Consensus/etcd and
// Curator) and Kubernetes. It watches the node roles as assigned by the
// cluster's curator, updates the status of the node within the curator, and
// spawns on-demand services.
//
//
//  .-----------.          .--------.  Watches  .------------.
//  | Cluster   |--------->| Role   |<----------| Node Roles |
//  | Enrolment | Provides | Server |  Updates  '------------'
//  '-----------'   Data   |        |----.      .-------------.
//                         '--------'    '----->| Node Status |
//                    Spawns |    | Spawns      '-------------'
//                     .-----'    '-----.
//                     V                V
//                 .-----------. .------------.
//                 | Consensus | | Kubernetes |
//                 | & Curator | |            |
//                 '-----------' '------------'
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
// instance. The state of of resolution of this bootstrap problem is maintained
// within ClusterMembership, which contains critical information about the
// control plane, like the information required to connect to a Curator (local
// or remote). It is updated both by external processes (ie. data from the
// Cluster Enrolment) as well as logic responsible for spawning the control
// plane.
//
package roleserve

import (
	"context"
	"crypto/ed25519"

	"source.monogon.dev/metropolis/node/core/identity"
	"source.monogon.dev/metropolis/node/core/localstorage"
	"source.monogon.dev/metropolis/node/core/network"
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
}

// Service is the roleserver/“Role Server” service. See the package-level
// documentation for more details.
type Service struct {
	Config

	ClusterMembership ClusterMembershipValue
	KubernetesStatus  KubernetesStatusValue
	bootstrapData     bootstrapDataValue
	localRoles        localRolesValue

	controlPlane *workerControlPlane
	statusPush   *workerStatusPush
	kubernetes   *workerKubernetes
	rolefetch    *workerRoleFetch
}

// New creates a Role Server services from a Config.
func New(c Config) *Service {
	s := &Service{
		Config: c,
	}

	s.controlPlane = &workerControlPlane{
		storageRoot: s.StorageRoot,

		bootstrapData:     &s.bootstrapData,
		clusterMembership: &s.ClusterMembership,
		localRoles:        &s.localRoles,
	}

	s.statusPush = &workerStatusPush{
		network: s.Network,

		clusterMembership: &s.ClusterMembership,
	}

	s.kubernetes = &workerKubernetes{
		network:     s.Network,
		storageRoot: s.StorageRoot,

		localRoles:        &s.localRoles,
		clusterMembership: &s.ClusterMembership,

		kubernetesStatus: &s.KubernetesStatus,
	}

	s.rolefetch = &workerRoleFetch{
		clusterMembership: &s.ClusterMembership,

		localRoles: &s.localRoles,
	}

	return s
}

func (s *Service) ProvideBootstrapData(privkey ed25519.PrivateKey, iok, cuk []byte) {
	s.ClusterMembership.set(&ClusterMembership{
		pubkey: privkey.Public().(ed25519.PublicKey),
	})
	s.bootstrapData.set(&bootstrapData{
		nodePrivateKey:   privkey,
		initialOwnerKey:  iok,
		clusterUnlockKey: cuk,
	})
}

func (s *Service) ProvideRegisterData(credentials identity.NodeCredentials, directory *cpb.ClusterDirectory) {
	s.ClusterMembership.set(&ClusterMembership{
		remoteCurators: directory,
		credentials:    &credentials,
		pubkey:         credentials.PublicKey(),
	})
}

// Run the Role Server service, which uses intermediary workload launchers to
// start/stop subordinate services as the Node's roles change.
func (s *Service) Run(ctx context.Context) error {
	supervisor.Run(ctx, "controlplane", s.controlPlane.run)
	supervisor.Run(ctx, "kubernetes", s.kubernetes.run)
	supervisor.Run(ctx, "statuspush", s.statusPush.run)
	supervisor.Run(ctx, "rolefetch", s.rolefetch.run)
	supervisor.Signal(ctx, supervisor.SignalHealthy)

	<-ctx.Done()
	return ctx.Err()
}
