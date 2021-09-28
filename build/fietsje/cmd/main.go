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

// fietsje is the standalone command line tool which calls into the Fietsje
// library (//build/fietsje) to perform actual work. The split between cli and
// main library is used so that Fietsje can be called from other Go tooling
// without having to shell out to a binary.
package main

import (
	"log"
	"os"
	"path"

	"source.monogon.dev/build/fietsje"
	"source.monogon.dev/build/toolbase"
	"source.monogon.dev/build/toolbase/gotoolchain"
)

func main() {
	// Get absolute path of Monogon workspace directory currently operating on
	// (either via bazel run or by running it directly in the root of a checkout),
	// use it to build paths to shelf.pb.txt and repositories.bzl.
	wd, err := toolbase.WorkspaceDirectory()
	if err != nil {
		log.Fatalf("%v", err)
	}
	shelfPath := path.Join(wd, "third_party/go/shelf.pb.text")
	repositoriesBzlPath := path.Join(wd, "third_party/go/repositories.bzl")
	// Set GOROOT as required by fietsje/go-the-tool.
	os.Setenv("GOROOT", gotoolchain.Root)

	if err := fietsje.Monogon(shelfPath, repositoriesBzlPath); err != nil {
		log.Fatal(err)
	}
}
