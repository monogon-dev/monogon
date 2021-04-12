// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"fmt"

	"source.monogon.dev/metropolis/pkg/pki"
)

// ClusterState is the state of the cluster from the point of view of the
// current node. Clients within the node code can watch this state to change
// their behaviour as needed.
type ClusterState int

const (
	// ClusterStateUnknown means the node has not yet determined the existence
	// of a cluster it should join or start. This is a transient, initial state
	// that should only manifest during boot.
	ClusterUnknown ClusterState = iota
	// ClusterForeign means the node is attempting to register into an already
	// existing cluster with which it managed to make preliminary contact, but
	// which the cluster has not yet fully productionized (eg. the node is
	// still being hardware attested, or the operator needs to confirm the
	// registration of this node).
	ClusterForeign
	// ClusterTrusted means the node is attempting to register into an already
	// registered cluster, and has been trusted by it. The node is now
	// attempting to finally commit into registering the cluster.
	ClusterTrusted
	// ClusterHome means the node is part of a cluster. This is the bulk of
	// time in which this node will spend its time.
	ClusterHome
	// ClusterDisowning means the node has been disowned (ie., removed) by the
	// cluster, and that it will not be ever part of any cluster again, and
	// that it will be decommissioned by the operator.
	ClusterDisowning
	// ClusterSplit means that the node would usually be Home in a cluster, but
	// has been split from the consensus of the cluster. This can happen for
	// nodes running consensus when consensus is lost (eg. when there is no
	// quorum or this node has been netsplit), and for other nodes if they have
	// lost network connectivity to the consensus nodes. Clients should make
	// their own decision what action to perform in this state, depending on
	// the level of consistency required and whether it makes sense for the
	// node to fence its services off.
	ClusterSplit
)

func (s ClusterState) String() string {
	switch s {
	case ClusterForeign:
		return "ClusterForeign"
	case ClusterTrusted:
		return "ClusterTrusted"
	case ClusterHome:
		return "ClusterHome"
	case ClusterDisowning:
		return "ClusterDisowning"
	case ClusterSplit:
		return "ClusterSplit"
	}
	return fmt.Sprintf("Invalid(%d)", s)
}

var (
	PKINamespace = pki.Namespaced("/cluster-pki/")
	PKICA        = PKINamespace.New(pki.SelfSigned, "cluster-ca", pki.CA("Metropolis Cluster CA"))
)
