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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bazelbuild/bazel-gazelle/label"
)

// dependency is an external Go package/module, requested by the user of Fietsje directly or indirectly.
type dependency struct {
	// importpath is the Go import path that was used to import this dependency.
	importpath string
	// version at which this dependency has been requested. This can be in any form that `go get` or the go module
	// system understands.
	version string

	// locked is the 'resolved' version of a dependency, containing information about the dependency's hash, etc.
	locked *locked

	// parent is the dependency that pulled in this one, or nil if pulled in by the user.
	parent *dependency

	shelf *shelf

	// Build specific settings passed to gazelle.
	disableProtoBuild bool
	buildTags         []string
	patches           []string
}

// locked is information about a dependency resolved from the go module system. It is expensive to get, and as such
// it is cached both in memory (as .locked in a dependency) and in the shelf.
type locked struct {
	// bazelName is the external workspace name that Bazel should use for this dependency, eg. com_github_google_glog.
	bazelName string
	// sum is the gomod compatible checksum of the depdendency, egh1:4A07+ZFc2wgJwo8YNlQpr1rVlgUDlxXHhPJciaPY5gs=.
	sum string
	// semver is the gomod-compatible version of this dependency. If the dependency was requested by git hash that does
	// not resolve to a particular release, this will be in the form of v0.0.0-20200520133742-deadbeefcafe.
	semver string
}

// child creates a new child dependence for this dependency, ie. one where the 'parent' pointer points to the dependency
// on which this method is called.
func (d *dependency) child(importpath, version string) *dependency {
	return &dependency{
		importpath: importpath,
		version:    version,
		shelf:      d.shelf,
		parent:     d,
	}
}

func (d *dependency) String() string {
	return fmt.Sprintf("%s@%s", d.importpath, d.version)
}

// lock ensures that this dependency is locked, which means that it has been resolved to a particular, stable version
// and VCS details. We lock a dependency by either asking the go module subsystem (via a go module proxy or a download),
// or by consulting the shelf as a cache.
func (d *dependency) lock() error {
	// If already locked in-memory, use that.
	if d.locked != nil {
		return nil
	}

	// If already locked in the shelf, use that.
	if shelved := d.shelf.get(d.importpath, d.version); shelved != nil {
		d.locked = shelved
		return nil
	}

	// Otherwise, download module.
	semver, _, sum, err := d.download()
	if err != nil {
		return fmt.Errorf("could not download: %v", err)
	}

	// And resolve its bazelName.
	name := label.ImportPathToBazelRepoName(d.importpath)

	// Hack for github.com/google/gvisor: it requests @com_github_opencontainers_runtime-spec.
	// We fix the generated name for this repo so it conforms to what gvisor expects.
	// TODO(q3k): instead of this, patch gvisor?
	if name == "com_github_opencontainers_runtime_spec" {
		name = "com_github_opencontainers_runtime-spec"
	}

	d.locked = &locked{
		bazelName: name,
		sum:       sum,
		semver:    semver,
	}
	log.Printf("%s: locked to %s", d, d.locked)

	// Save locked version to shelf.
	d.shelf.put(d.importpath, d.version, d.locked)
	return d.shelf.save()
}

func (l *locked) String() string {
	return fmt.Sprintf("%s@%s", l.bazelName, l.sum)
}

// download ensures that this dependency is download locally, and returns the download location and the dependency's
// gomod-compatible sum.
func (d *dependency) download() (version, dir, sum string, err error) {
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		err = fmt.Errorf("GOROOT must be set")
		return
	}
	goTool := filepath.Join(goroot, "bin", "go")

	query := fmt.Sprintf("%s@%s", d.importpath, d.version)
	cmd := exec.Command(goTool, "mod", "download", "-json", "--", query)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("go mod returned: %q", out)
		err = fmt.Errorf("go mod failed: %v", err)
		return
	}

	var res struct{ Version, Sum, Dir string }
	err = json.Unmarshal(out, &res)
	if err != nil {
		return
	}

	version = res.Version
	dir = res.Dir
	sum = res.Sum
	return
}
