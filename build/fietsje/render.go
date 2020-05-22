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

import (
	"fmt"
	"io"
	"sort"
)

// render writes a gazelle-compatible starlark file based on the enabled dependencies in this planner.
func (p *planner) render(w io.Writer) error {
	fmt.Fprintln(w, `load("@bazel_gazelle//:deps.bzl", "go_repository")`)
	fmt.Fprintln(w, ``)
	fmt.Fprintln(w, `def go_repositories():`)

	// Get and sort all enabled importpaths.
	var enabled []string
	for importpath, _ := range p.enabled {
		enabled = append(enabled, importpath)
	}
	sort.Slice(enabled, func(i, j int) bool { return enabled[i] < enabled[j] })

	// Render all importpaths.
	for _, importpath := range enabled {
		d := p.available[importpath]
		if err := d.lock(); err != nil {
			return fmt.Errorf("could not lock %q: %v", importpath, err)
		}

		fmt.Fprintf(w, "    go_repository(\n")
		fmt.Fprintf(w, "        name = %q,\n", d.locked.bazelName)
		fmt.Fprintf(w, "        importpath = %q,\n", d.importpath)
		fmt.Fprintf(w, "        version = %q,\n", d.locked.semver)
		fmt.Fprintf(w, "        sum = %q,\n", d.locked.sum)
		if d.disableProtoBuild {
			fmt.Fprintf(w, "        build_file_proto_mode = %q,\n", "disable")
		}
		if d.buildTags != nil {
			fmt.Fprintf(w, "        build_tags = [\n")
			for _, tag := range d.buildTags {
				fmt.Fprintf(w, "            %q,\n", tag)
			}
			fmt.Fprintf(w, "        ],\n")
		}
		if d.patches != nil {
			fmt.Fprintf(w, "        patches = [\n")
			for _, patch := range d.patches {
				fmt.Fprintf(w, "            %q,\n", "//third_party/go/patches:"+patch)
			}
			fmt.Fprintf(w, "        ],\n")
			fmt.Fprintf(w, "        patch_args = [%q],\n", "-p1")
		}

		fmt.Fprintf(w, "    )\n")
	}
	return nil
}
