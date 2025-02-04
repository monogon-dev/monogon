// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build tools
// +build tools

package node

import (
	_ "github.com/containerd/containerd/v2/cmd/containerd"
	_ "github.com/containerd/containerd/v2/cmd/containerd-shim-runc-v2"
	_ "github.com/containernetworking/plugins/plugins/ipam/host-local"
	_ "github.com/containernetworking/plugins/plugins/main/loopback"
	_ "github.com/containernetworking/plugins/plugins/main/ptp"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/opencontainers/runc"
	_ "github.com/prometheus/node_exporter"
	_ "gvisor.dev/gvisor/runsc"
)
