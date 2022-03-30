//go:build tools
// +build tools

package node

import (
	_ "github.com/containerd/containerd"
	_ "github.com/containernetworking/plugins/plugins/ipam/host-local"
	_ "github.com/containernetworking/plugins/plugins/main/loopback"
	_ "github.com/containernetworking/plugins/plugins/main/ptp"
	_ "github.com/coredns/coredns"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/opencontainers/runc"
	_ "gvisor.dev/gvisor/runsc"
)
