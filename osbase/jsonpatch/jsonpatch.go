// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package jsonpatch contains data structures and encoders for JSON Patch (RFC
// 6902) and JSON Pointers (RFC 6901)
package jsonpatch

import "strings"

// JsonPatchOp describes a JSON Patch operation (RFC 6902 Section 4)
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
