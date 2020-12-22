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

func depsCilium(p *planner) {
	p.collect(
		"github.com/cilium/cilium", "v1.8.0-rc1",
		disabledProtoBuild,
	).replace(
		// Cilium forks this library to introduce an API that they need, but
		// which the upstream rejected. Since this new API does not cause any
		// backwards incompatibility, we pull in their fork.
		// See https://github.com/miekg/dns/pull/949#issuecomment-487832458 for
		// more information about the fork.
		"github.com/miekg/dns", "github.com/cilium/dns", "8e25ec9a0ff3",
	).replace(
		// Cilium forks this library (a Go Kafka client) to apply the following
		// patches on top:
		//   - 01ce283: Fix de/serialization of null arrays
		//   - c411825: Correcly check msgSize in ReadResp before discarding.
		//   - 947cc36: Kafka : Cilium panics with error *index out of range
		//              error* on receiving messages of very large size.
		// serge@ has not found any trace of the Cilium project trying to
		// upstream this, but the patches seem to be only bugfixes, not
		// breaking functionality.
		// However, the fork-off point of the upstream project is fairly old
		// (commit b5a758db, dated Dec 7, 2017 - pre v1.5.0 of upstream). This
		// might cause issues in the future when we start to have other
		// consumers of this library.
		"github.com/optiopay/kafka", "github.com/cilium/kafka", "01ce283b732b",
	).use(
		"github.com/hashicorp/go-immutable-radix",
		"github.com/sasha-s/go-deadlock",
		"github.com/google/gopacket",
		"github.com/hashicorp/consul/api",
		"github.com/pborman/uuid",
		"github.com/petermattis/goid",
		"github.com/kr/text",
		"github.com/hashicorp/go-cleanhttp",
		"github.com/hashicorp/serf",
		"github.com/joho/godotenv",
		"github.com/envoyproxy/protoc-gen-validate",
		"github.com/hashicorp/go-rootcerts",
		"github.com/armon/go-metrics",
		"github.com/shirou/gopsutil",
		"github.com/cncf/udpa/go",
		"github.com/cpuguy83/go-md2man/v2",
		"github.com/russross/blackfriday/v2",
		"github.com/shurcooL/sanitized_anchor_name",
		"github.com/google/gops",
		"github.com/mattn/go-shellwords",
		"github.com/c9s/goprocinfo",
		"github.com/cilium/ipam",
		"github.com/kardianos/osext",
		"github.com/servak/go-fastping",
		"github.com/golang/snappy",
		"github.com/cilium/arping",
	).with(disabledProtoBuild, forceBazelGeneration).use(
		"github.com/cilium/proxy",
	).with(disabledProtoBuild, buildExtraArgs("-exclude=src")).use(
		// -exclude=src fixes a build issue with Gazelle. See:
		// https://github.com/census-instrumentation/opencensus-proto/issues/200
		"github.com/census-instrumentation/opencensus-proto",
	)
}
