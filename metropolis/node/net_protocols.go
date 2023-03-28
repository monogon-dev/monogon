package node

import "github.com/vishvananda/netlink"

// These are netlink protocol numbers used internally for various netlink
// resource (e.g. route) owners/manager.
const (
	// ProtocolClusternet is used by //metropolis/node/core/clusternet when
	// creating/removing routes pointing to the clusternet interface.
	ProtocolClusternet netlink.RouteProtocol = 129
)
