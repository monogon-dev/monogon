// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package node

// These are netlink protocol numbers used internally for various netlink
// resource (e.g. route) owners/manager.
const (
	// ProtocolClusternet is used by //metropolis/node/core/clusternet when
	// creating/removing routes pointing to the clusternet interface.
	ProtocolClusternet int = 129
)

// Netlink link groups used for interface classification and traffic matching.
const (
	// LinkGroupK8sPod is set on all host side PtP interfaces going to K8s
	// pods.
	LinkGroupK8sPod uint32 = 8
	// LinkGroupClusternet is set on all interfaces not needing SNAT from the
	// K8s internal IPs.
	LinkGroupClusternet uint32 = 9
)
