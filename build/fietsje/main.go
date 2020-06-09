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
	"bytes"
	"flag"
	"io/ioutil"
	"log"
)

var (
	flagShelfPath          string
	flagRepositoresBzlPath string
)

func main() {
	flag.StringVar(&flagShelfPath, "shelf_path", "", "Path to shelf (cache/lockfile)")
	flag.StringVar(&flagRepositoresBzlPath, "repositories_bzl", "", "Path to output repositories.bzl file")
	flag.Parse()

	if flagShelfPath == "" {
		log.Fatalf("shelf_path must be set")
	}
	if flagRepositoresBzlPath == "" {
		log.Fatalf("repositories_bzl must be set")
	}

	shelf, err := shelfLoad(flagShelfPath)
	if err != nil {
		log.Fatalf("could not load shelf: %v", err)
	}

	p := &planner{
		available: make(map[string]*dependency),
		enabled:   make(map[string]bool),
		seen:      make(map[string]string),

		shelf: shelf,
	}

	// gRPC/proto deps (https://github.com/bazelbuild/rules_go/blob/master/go/workspace.rst#id8)
	// bump down from 1.28.1 to 1.26.0 because https://github.com/etcd-io/etcd/issues/11563
	p.collect(
		"google.golang.org/grpc", "v1.26.0",
	).use(
		"golang.org/x/net",
		"golang.org/x/text",
	)

	depsKubernetes(p)
	depsContainerd(p)
	depsGVisor(p)
	depsCilium(p)
	depsSQLBoiler(p)

	// our own deps, common
	p.collectOverride("go.uber.org/zap", "v1.15.0")
	p.collectOverride("golang.org/x/mod", "v0.3.0")
	p.collect("github.com/cenkalti/backoff/v4", "v4.0.2")

	p.collect("github.com/google/go-tpm", "ae6dd98980d4")
	p.collect("github.com/google/go-tpm-tools", "f8c04ff88181")
	p.collect("github.com/insomniacslk/dhcp", "5dd7202f19711228cb4a51aa8b3415421c2edefe")
	p.collect("github.com/mdlayher/ethernet", "0394541c37b7f86a10e0b49492f6d4f605c34163").use(
		"github.com/mdlayher/raw",
	)
	p.collect("github.com/rekby/gpt", "a930afbc6edcc89c83d39b79e52025698156178d")
	p.collect("github.com/yalue/native_endian", "51013b03be4fd97b0aabf29a6923e60359294186")

	// used by insomniacslk/dhcp for pkg/uio
	p.collect("github.com/u-root/u-root", "v6.0.0")

	// used by //core/cmd/mkimage
	p.collect("github.com/diskfs/go-diskfs", "v1.0.0").use(
		"gopkg.in/djherbis/times.v1",
	)

	// used by //build/bindata
	p.collect("github.com/kevinburke/go-bindata", "v3.16.0")

	// used by deltagen
	p.collectOverride("github.com/lyft/protoc-gen-star", "v0.4.14")

	// for interactive debugging during development (//:dlv alias)
	depsDelve(p)

	// First generate the repositories starlark rule into memory. This is because rendering will lock all unlocked
	// dependencies, which might take a while. If a use were to interrupt it now, they would end up with an incomplete
	// repositories.bzl and would have to restore from git.
	buf := bytes.NewBuffer(nil)
	err = p.render(buf)
	if err != nil {
		log.Fatalf("could not render deps: %v", err)
	}

	err = ioutil.WriteFile(flagRepositoresBzlPath, buf.Bytes(), 0666)
	if err != nil {
		log.Fatalf("could not write deps: %v", err)
	}
}