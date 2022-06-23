package rpc

import (
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
