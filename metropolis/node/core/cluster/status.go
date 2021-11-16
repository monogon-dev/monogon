package cluster

import (
	"errors"

	"source.monogon.dev/metropolis/node/core/consensus"
	"source.monogon.dev/metropolis/node/core/identity"
	cpb "source.monogon.dev/metropolis/proto/common"
)

var (
	ErrNoLocalConsensus = errors.New("this node does not have direct access to etcd")
)

// Status is returned to Cluster clients (ie., node code) on Manager.Watch/.Get.
type Status struct {
	// State is the current state of the cluster, as seen by the node.
	State cpb.ClusterState

	// Consensus is a handle to a running Consensus service, or nil if this node
	// does not run a Consensus instance.
	Consensus consensus.ServiceHandle

	// Credentials used for the node to authenticate to the Curator and other
	// cluster services.
	Credentials *identity.NodeCredentials
}
