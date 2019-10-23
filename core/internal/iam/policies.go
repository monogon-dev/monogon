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
	"context"
	"fmt"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/util"
)

type dataSetProfile struct {
	numTokens int
	numPaths  int
}

func main() {
	ctx := context.Background()
	compiler := ast.NewCompiler()
	module := ast.MustParseModule(policy)

	compiler.Compile(map[string]*ast.Module{"": module})
	if compiler.Failed() {
	}

	r := rego.New(
		rego.Compiler(compiler),
		rego.Input(util.MustUnmarshalJSON([]byte(`{
			"token_id": "deadbeef",
			"path": "mna",
			"method": "GET"
		}`))),
		rego.Query("data.restauthz"),
	)

	rs, err := r.Eval(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", rs)
}

const policy = `package restauthz

default allow = false

allow {
	input.method == "GET"
}

allow {
	not input.method == "GET"
}
`
