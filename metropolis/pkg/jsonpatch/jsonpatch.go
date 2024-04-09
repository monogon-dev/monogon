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

// Package jsonpatch contains data structures and encoders for JSON Patch (RFC
// 6902) and JSON Pointers (RFC 6901)
package jsonpatch

import "strings"

// JSON Patch operation (RFC 6902 Section 4)
type JsonPatchOp struct {
	Operation string      `json:"op"`
	Path      string      `json:"path"` // Technically a JSON Pointer, but called Path in the RFC
	From      string      `json:"from,omitempty"`
	Value     interface{} `json:"value,omitempty"`
}

// EncodeJSONRefToken encodes a JSON reference token as part of a JSON Pointer
// (RFC 6901 Section 2)
func EncodeJSONRefToken(token string) string {
	x := strings.ReplaceAll(token, "~", "~0")
	return strings.ReplaceAll(x, "/", "~1")
}

// PointerFromParts returns an encoded JSON Pointer from parts
func PointerFromParts(pathParts []string) string {
	var encodedParts []string
	encodedParts = append(encodedParts, "")
	for _, part := range pathParts {
		encodedParts = append(encodedParts, EncodeJSONRefToken(part))
	}
	return strings.Join(encodedParts, "/")
}
