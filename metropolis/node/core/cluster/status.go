package cluster

import (
	"errors"
	"fmt"

	"source.monogon.dev/metropolis/node/core/consensus/client"
	"source.monogon.dev/metropolis/node/core/curator"
	cpb "source.monogon.dev/metropolis/proto/common"
)

var (
	ErrNoLocalConsensus = errors.New("this node does not have direct access to etcd")
)

// Status is returned to Cluster clients (ie., node code) on Manager.Watch/.Get.
type Status struct {
	// State is the current state of the cluster, as seen by the node.
	State cpb.ClusterState

	// hasLocalConsensus is true if the local node is running a local consensus
	// (etcd) server.
	hasLocalConsensus bool
	// consensusClient is an etcd client to the local consensus server if the node
	// has such a server and the cluster state is HOME or SPLIT.
	consensusClient client.Namespaced

	// Credentials used for the node to authenticate to the Curator and other
	// cluster services.
	Credentials *curator.NodeCredentials
}

// ConsensusUser is the to-level user of an etcd client in Metropolis node
// code. These need to be defined ahead of time in an Go 'enum', and different
// ConsensusUsers should not be shared by different codepaths.
type ConsensusUser string

const (
	ConsensusUserKubernetesPKI ConsensusUser = "kubernetes-pki"
	ConsensusUserCurator       ConsensusUser = "curator"
)

// ConsensusClient returns an etcd/consensus client for a given ConsensusUser.
// The node must be running a local consensus/etcd server.
func (s *Status) ConsensusClient(user ConsensusUser) (client.Namespaced, error) {
	if !s.hasLocalConsensus {
		return nil, ErrNoLocalConsensus
	}

	// Ensure that we already are connected to etcd and are in a state in which we
	// should be handing out cluster connectivity.
	if s.consensusClient == nil {
		return nil, fmt.Errorf("not connected")
	}
	switch s.State {
	case cpb.ClusterState_CLUSTER_STATE_HOME:
	case cpb.ClusterState_CLUSTER_STATE_SPLIT:
		// The consensus client is resistant to being split off, and will serve
		// as soon as the split is resolved.
	default:
		return nil, fmt.Errorf("refusing connection with cluster state %v", s.State)
	}

	// Ensure only defined 'applications' are used to prevent programmer error and
	// casting to ConsensusUser from an arbitrary string.
	switch user {
	case ConsensusUserKubernetesPKI:
	case ConsensusUserCurator:
	default:
		return nil, fmt.Errorf("unknown ConsensusUser %q", user)
	}
	client, err := s.consensusClient.Sub(string(user))
	if err != nil {
		return nil, fmt.Errorf("retrieving subclient failed: %w", err)
	}
	return client, nil
}
