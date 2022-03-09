package rpc

import (
	cpb "source.monogon.dev/metropolis/node/core/curator/proto/api"
	apb "source.monogon.dev/metropolis/proto/api"
	epb "source.monogon.dev/metropolis/proto/ext"
)

var (
	// nodePermissions are the set of metropolis.common.ext.authorization
	// permissions automatically given to nodes when connecting to curator gRPC
	// services, either locally or remotely.
	nodePermissions = Permissions{
		epb.Permission_PERMISSION_READ_CLUSTER_STATUS: true,
		epb.Permission_PERMISSION_UPDATE_NODE_SELF:    true,
	}
)

// ClusterServices is the interface containing all gRPC services that a
// Metropolis Cluster implements. With the current implementation of Metropolis,
// this is all implemented by the Curator.
type ClusterServices interface {
	cpb.CuratorServer
	apb.AAAServer
	apb.ManagementServer
}
