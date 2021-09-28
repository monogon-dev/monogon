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

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

// getTransitiveDeps is a hairy ball of heuristic used to find all recursively
// transitive dependencies of a given dependency. It downloads a given dependency
// using `go get`, and performs analysis of standard (go.mod/go.sum) and project-
// specific dependency management configuration/lock files in order to build a full
// view of all known, versioned transitive dependencies.
func (d *dependency) getTransitiveDeps() (map[string]*dependency, error) {
	// First, lock the dependency. Downloading it later will also return a sum, and we
	// want to ensure both are the same.
	err := d.lock()
	if err != nil {
		return nil, fmt.Errorf("could not lock: %v", err)
	}

	_, path, sum, err := d.download()
	if err != nil {
		return nil, fmt.Errorf("could not download: %v", err)
	}

	if sum != d.locked.sum {
		return nil, fmt.Errorf("inconsistent sum: %q downloaded, %q in shelf/lock", sum, d.locked.sum)
	}

	exists := func(p string) bool {
		full := fmt.Sprintf("%s/%s", path, p)
		if _, err := os.Stat(full); err == nil {
			return true
		}
		if err != nil && !os.IsExist(err) {
			panic(fmt.Sprintf("checking file %q: %v", full, err))
		}
		return false
	}

	read := func(p string) []byte {
		full := fmt.Sprintf("%s/%s", path, p)
		data, err := ioutil.ReadFile(full)
		if err != nil {
			panic(fmt.Sprintf("reading file %q: %v", full, err))
		}
		return data
	}

	requirements := make(map[string]*dependency)

	// Read & parse go.mod if present.
	var mf *modfile.File
	if exists("go.mod") {
		log.Printf("%q: parsing go.mod\n", d.importpath)
		data := read("go.mod")
		mf, err = modfile.Parse("go.mod", data, nil)
		if err != nil {
			return nil, fmt.Errorf("parsing go.mod in %s: %v", d.importpath, err)
		}
	}

	// If a go.mod file was present, interpret it to populate dependencies.
	if mf != nil {
		for _, req := range mf.Require {
			requirements[req.Mod.Path] = d.child(req.Mod.Path, req.Mod.Version)
		}
		for _, rep := range mf.Replace {
			// skip filesystem rewrites
			if rep.New.Version == "" {
				continue
			}

			requirements[rep.New.Path] = d.child(rep.New.Path, rep.New.Version)
		}
	}

	// Read parse, and interpret. go.sum if present.
	// This should bring into view all recursively transitive dependencies.
	if exists("go.sum") {
		log.Printf("%q: parsing go.sum", d.importpath)
		data := read("go.sum")
		for _, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.Fields(line)
			if len(parts) != 3 {
				return nil, fmt.Errorf("parsing go.sum: unparseable line %q", line)
			}

			importpath, version := parts[0], parts[1]

			// Skip if already created from go.mod.
			// TODO(q3k): error if go.sum and go.mod disagree?
			if _, ok := requirements[importpath]; ok {
				continue
			}

			if strings.HasSuffix(version, "/go.mod") {
				version = strings.TrimSuffix(version, "/go.mod")
			}
			requirements[importpath] = d.child(importpath, version)
		}
	}

	// Special case: root Kubernetes repo - rewrite staging/ deps to k8s.io/ at correct
	// versions, quit early. Kubernetes vendors all dependencies into vendor/, and also
	// contains sub-projects (components) in staging/. This converts all staging
	// dependencies into appropriately versioned k8s.io/<dep> paths.
	if d.importpath == "k8s.io/kubernetes" {
		log.Printf("%q: special case for Kubernetes main repository", d.importpath)
		if mf == nil {
			return nil, fmt.Errorf("k8s.io/kubernetes needs a go.mod")
		}
		// extract the version, turn into component version
		version := d.version
		if !strings.HasPrefix(version, "v") {
			return nil, fmt.Errorf("invalid version format for k8s: %q", version)
		}
		version = version[1:]
		componentVersion := fmt.Sprintf("kubernetes-%s", version)

		// find all k8s.io 'components'
		components := make(map[string]bool)
		for _, rep := range mf.Replace {
			if !strings.HasPrefix(rep.Old.Path, "k8s.io/") || !strings.HasPrefix(rep.New.Path, "./staging/src/") {
				continue
			}
			components[rep.Old.Path] = true
		}

		// add them to planner at the 'kubernetes-$ver' tag
		for component, _ := range components {
			requirements[component] = d.child(component, componentVersion)
		}
		return requirements, nil
	}

	// Special case: github.com/containerd/containerd: read vendor.conf.
	if d.importpath == "github.com/containerd/containerd" {
		log.Printf("%q: special case for containerd", d.importpath)
		if !exists("vendor.conf") {
			panic("containerd needs vendor.conf")
		}
		data := read("vendor.conf")
		for _, line := range strings.Split(string(data), "\n") {
			// strip comments
			parts := strings.SplitN(line, "#", 2)
			line = parts[0]

			// skip empty contents
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			// read dep/version pairs
			parts = strings.Fields(line)
			if len(parts) < 2 {
				return nil, fmt.Errorf("unparseable line in containerd vendor.conf: %q", line)
			}
			importpath, version := parts[0], parts[1]
			requirements[importpath] = d.child(importpath, version)
		}
		return requirements, nil
	}

	return requirements, nil
}
