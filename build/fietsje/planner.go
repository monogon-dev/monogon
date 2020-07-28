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
)

// The Planner provides the main DSL and high-level control logic for resolving dependencies. It is the main API that
// fietsje users should consume.

// planner is a builder for a single world of Go package dependencies, and what is then emitted into a Starlark file
// containing gazelle go_repository rules.
// The planner's builder system covers three increasingly specific contextx:
//  - planner (this structure, allows for 'collecting' in high-level dependencies. ie. collections)
//  - collection (represents what has been pulled in by a high-level dependency, and allows for 'using' transitive
//    dependencies from a collection)
//  - optionized (represents a collection with extra build flags, eg. disabled proto builds)
type planner struct {
	// available is a map of importpaths to dependencies that the planner knows. This is a flat structure that is the
	// main source of truth of actual dependency data, like a registry of everything that the planner knows about.
	// The available dependency for a given importpath, as the planner progresses, might change, ie. when there is a
	// version conflict. As such, code should use importpaths as atoms describing dependencies, instead of holding
	// dependency pointers.
	available map[string]*dependency
	// enabled is a map of dependencies that will be emitted by the planner into the build via Gazelle.
	enabled map[string]bool
	// seen is a map of 'dependency' -> 'parent' importpaths, ie. returns what higher-level dependency (ie. one enabled
	// with .collect()) pulled in a given dependency. This is only used for error messages to help the user find what
	// a transitive  dependency has been pulled in by.
	seen map[string]string

	shelf *shelf
}

func (p *planner) collect(importpath, version string, opts ...buildOpt) *collection {
	return p.collectInternal(importpath, version, false, opts...)
}

func (p *planner) collectOverride(importpath, version string, opts ...buildOpt) *collection {
	return p.collectInternal(importpath, version, true, opts...)
}

// collectInternal pulls in a high-level dependency into the planner and
// enables it. It also parses all of its transitive // dependencies (not just
// directly transitive, but recursively transitive) and makes the planner aware
// of them. It does not enable these transitive dependencies, but returns a
// collection builder, which can be used to do se by calling .use().
func (p *planner) collectInternal(importpath, version string, override bool, opts ...buildOpt) *collection {
	// Ensure overrides are explicit and minimal.
	by, ok := p.seen[importpath]
	if ok && !override {
		panic(fmt.Errorf("%s is being collected, but has already been declared by %s; replace it by a use(%q) call on %s or use collectOverride", importpath, by, importpath, by))
	}
	if !ok && override {
		panic(fmt.Errorf("%s is being collected with override, but has not been seen as a dependency previously - use .collect(%q, %q) instead", importpath, importpath, version))
	}

	d := &dependency{
		shelf:      p.shelf,
		importpath: importpath,
		version:    version,
	}
	for _, o := range opts {
		o(d)
	}

	// automatically enable direct import
	p.enabled[d.importpath] = true
	p.available[d.importpath] = d

	td, err := d.getTransitiveDeps()
	if err != nil {
		panic(fmt.Errorf("could not get transitive deps for %q: %v", d.importpath, err))
	}
	// add transitive deps to 'available' map
	for k, v := range td {
		// skip dependencies that have already been enabled, dependencies are 'first enabled version wins'.
		if _, ok := p.available[k]; ok && p.enabled[k] {
			continue
		}

		p.available[k] = v

		// make note of the high-level dependency that pulled in the dependency.
		p.seen[v.importpath] = d.importpath
	}

	return &collection{
		p:          p,
		highlevel:  d,
		transitive: td,
	}
}

// collection represents the context of the planner after pulling/collecting in a high-level dependency. In this state,
// the planner can be used to enable transitive dependencies of the high-level dependency.
type collection struct {
	p *planner

	highlevel  *dependency
	transitive map[string]*dependency
}

// use enables given dependencies defined in the collection by a high-level dependency.
func (c *collection) use(paths ...string) *collection {
	return c.with().use(paths...)
}

// replace injects a new dependency with a replacement importpath. This is used to reflect 'replace' stanzas in go.mod
// files of third-party dependencies. This is not done automatically by Fietsje, as a replacement is global to the
// entire build tree, and should be done knowingly and explicitly by configuration. The 'oldpath' importpath will be
// visible to the build system, but will be backed at 'newpath' locked at 'version'.
func (c *collection) replace(oldpath, newpath, version string) *collection {
	// Ensure oldpath is in use. We want as little replacements as possible, and if it's not being used by anything,
	// it means that we likely don't need it.
	c.use(oldpath)

	d := c.highlevel.child(oldpath, version)
	d.replace = newpath
	c.transitive[oldpath] = d
	c.p.available[oldpath] = d
	c.p.enabled[oldpath] = true

	return c
}

// inject adds a dependency to a collection as if requested by the high-level dependency of the collection. This should
// be used sparingly, for instance when high-level dependencies contain bazel code that uses some external workspaces
// from Go modules, and those workspaces are not defined in parsed transitive dependency definitions like go.mod/sum.
func (c *collection) inject(importpath, version string) *collection {
	d := c.highlevel.child(importpath, version)
	c.transitive[importpath] = d
	c.p.available[importpath] = d
	c.p.enabled[importpath] = true

	return c
}

// with transforms a collection into an optionized, by setting some build options.
func (c *collection) with(o ...buildOpt) *optionized {
	return &optionized{
		c:    c,
		opts: o,
	}
}

// optionized is a collection that has some build options set, that will be applied to all dependencies 'used' in this
// context
type optionized struct {
	c    *collection
	opts []buildOpt
}

// buildOpt is a build option passed to Gazelle.
type buildOpt func(d *dependency)

// buildTags sets the given buildTags in affected dependencies.
func buildTags(tags ...string) buildOpt {
	return func(d *dependency) {
		d.buildTags = tags
	}
}

// disabledProtoBuild disables protobuf builds in affected dependencies.
func disabledProtoBuild(d *dependency) {
	d.disableProtoBuild = true
}

// patches applies patches in affected dependencies after BUILD file generation.
func patches(patches ...string) buildOpt {
	return func(d *dependency) {
		d.patches = patches
	}
}

// prePatches applies patches in affected dependencies before BUILD file generation.
func prePatches(patches ...string) buildOpt {
	return func(d *dependency) {
		d.prePatches = patches
	}
}

func forceBazelGeneration(d *dependency) {
	d.forceBazelGeneration = true
}

func buildExtraArgs(args ...string) buildOpt {
	return func(d *dependency) {
		d.buildExtraArgs = args
	}
}

// use enables given dependencies defined in the collection by a high-level dependency, with any set build options.
// After returning, the builder degrades to a collection - ie, all build options are reset.
func (o *optionized) use(paths ...string) *collection {
	for _, path := range paths {
		el, ok := o.c.transitive[path]
		if !ok {
			msg := fmt.Sprintf("dependency %q not found in %q", path, o.c.highlevel.importpath)
			if alternative, ok := o.c.p.seen[path]; ok {
				msg += fmt.Sprintf(" (but found in %q)", alternative)
			} else {
				msg += " or any other collected library"
			}
			panic(msg)
		}
		for _, o := range o.opts {
			o(el)
		}
		o.c.p.enabled[path] = true
	}

	return o.c
}
