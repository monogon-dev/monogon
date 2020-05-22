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

package main

func depsContainerd(p *planner) {
	p.collectOverride(
		"github.com/containerd/containerd", "8e685f78cf66e2901b2fbed2fdddd64449a74ab9",
		buildTags("no_zfs", "no_aufs", "no_devicemapper", "no_btrfs"),
		disabledProtoBuild,
	).use(
		"github.com/BurntSushi/toml",
		"github.com/Microsoft/go-winio",
		"github.com/beorn7/perks",
		"github.com/cespare/xxhash/v2",
		"github.com/cilium/ebpf",
		"github.com/containerd/btrfs",
		"github.com/containerd/console",
		"github.com/containerd/continuity",
		"github.com/containerd/fifo",
		"github.com/containerd/go-cni",
		"github.com/containerd/go-runc",
		"github.com/containerd/ttrpc",
		"github.com/containerd/typeurl",
		"github.com/containernetworking/cni",
		"github.com/coreos/go-systemd/v22",
		"github.com/cpuguy83/go-md2man",
		"github.com/davecgh/go-spew",
		"github.com/docker/distribution",
		"github.com/docker/docker",
		"github.com/docker/go-events",
		"github.com/docker/go-metrics",
		"github.com/docker/go-units",
		"github.com/docker/spdystream",
		"github.com/emicklei/go-restful",
		"github.com/godbus/dbus/v5",
		"github.com/gogo/protobuf",
		"github.com/google/gofuzz",
		"github.com/google/uuid",
		"github.com/hashicorp/errwrap",
		"github.com/hashicorp/go-multierror",
		"github.com/hashicorp/golang-lru",
		"github.com/imdario/mergo",
		"github.com/json-iterator/go",
		"github.com/konsorten/go-windows-terminal-sequences",
		"github.com/matttproud/golang_protobuf_extensions",
		"github.com/modern-go/concurrent",
		"github.com/modern-go/reflect2",
		"github.com/opencontainers/go-digest",
		"github.com/opencontainers/image-spec",
		"github.com/opencontainers/runc",
		"github.com/opencontainers/runtime-spec",
		"github.com/pkg/errors",
		"github.com/prometheus/client_golang",
		"github.com/prometheus/client_model",
		"github.com/prometheus/common",
		"github.com/prometheus/procfs",
		"github.com/russross/blackfriday",
		"github.com/seccomp/libseccomp-golang",
		"github.com/sirupsen/logrus",
		"github.com/stretchr/testify",
		"github.com/syndtr/gocapability",
		"github.com/tchap/go-patricia",
		"github.com/urfave/cli",
		"go.etcd.io/bbolt",
		"go.opencensus.io",
		"golang.org/x/crypto",
		"golang.org/x/oauth2",
		"golang.org/x/sync",
		"golang.org/x/sys",
		"google.golang.org/genproto",
		"gopkg.in/inf.v0",
		"gopkg.in/yaml.v2",
		"gotest.tools",
		"k8s.io/klog",
		"sigs.k8s.io/yaml",
	).with(disabledProtoBuild).use(
		"github.com/Microsoft/hcsshim",
		"github.com/containerd/cgroups",
		"github.com/containerd/cri",
		"github.com/gogo/googleapis",
	).with(buildTags("selinux")).use(
		"github.com/opencontainers/selinux",
	)

	// containernetworking/plugins
	p.collectOverride(
		"github.com/containernetworking/plugins", "v0.8.2",
		patches("cni-plugins-build.patch"),
	).use(
		"github.com/alexflint/go-filemutex",
		"github.com/coreos/go-iptables",
		"github.com/j-keck/arping",
		"github.com/safchain/ethtool",
	)
}
