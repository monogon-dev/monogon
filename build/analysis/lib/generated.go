// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package lib

import (
	"go/ast"
	"strings"
)

const (
	genPrefix = "// Code generated"
	genSuffix = "DO NOT EDIT."
)

// IsGeneratedFile returns true if the file is generated according to
// https://golang.org/s/generatedcode and other heuristics.
func IsGeneratedFile(file *ast.File) bool {
	for _, c := range file.Comments {
		for _, t := range c.List {
			if strings.HasPrefix(t.Text, genPrefix) && strings.HasSuffix(t.Text, genSuffix) {
				return true
			}
			// Generated testmain.go stubs from rules_go - for some reason, they don't
			// contain the expected markers.
			if strings.Contains(t.Text, "This package must be initialized before packages being tested.") {
				return true
			}
		}
	}
	return false
}
