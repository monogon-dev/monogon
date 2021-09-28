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

package fietsje

// deps_gvisor.go contains all dependencies required by gVisor/runsc.

func depsGVisor(p *planner) {
	p.collect(
		"github.com/google/gvisor", "release-20201216.0",
		patches(
			"gvisor.patch",
			"gvisor-build-against-newer-runtime-specs.patch",
		),
	).use(
		"github.com/cenkalti/backoff",
		"github.com/gofrs/flock",
		"github.com/google/subcommands",
		"github.com/kr/pretty",
		"github.com/kr/pty",
		"github.com/mohae/deepcopy",
		"golang.org/x/time",
	)
	// gRPC is used by gvisor's bazel machinery, but not present in go.sum. Include it
	// manually.
	p.collect("github.com/grpc/grpc", "v1.29.1")
}
