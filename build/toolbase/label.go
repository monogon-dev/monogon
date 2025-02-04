// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package toolbase

import (
	"fmt"
	"regexp"
	"strings"
)

// BazelLabel is a label, as defined by Bazel's documentation:
//
// https://docs.bazel.build/versions/main/skylark/lib/Label.html
type BazelLabel struct {
	WorkspaceName string
	PackagePath   []string
	Name          string
}

func (b BazelLabel) Package() string {
	return strings.Join(b.PackagePath, "/")
}

func (b BazelLabel) String() string {
	return fmt.Sprintf("@%s//%s:%s", b.WorkspaceName, b.Package(), b.Name)
}

var (
	// reLabel splits a Bazel label into a workspace name (if set) and a
	// workspace root relative package/name.
	reLabel = regexp.MustCompile(`^(@[^:/]+)?//(.+)$`)
	// rePathPart matches valid label path parts.
	rePathPart = regexp.MustCompile(`^[^:/]+$`)
)

// ParseBazelLabel converts parses a string representation of a Bazel Label. If
// the given representation is invalid or for some other reason unparseable,
// nil is returned.
func ParseBazelLabel(s string) *BazelLabel {
	res := BazelLabel{
		WorkspaceName: "@",
	}

	// Split label into workspace name (if set) and a workspace root relative
	// package/name.
	m := reLabel.FindStringSubmatch(s)
	if m == nil {
		return nil
	}
	packageRel := m[2]
	if m[1] != "" {
		res.WorkspaceName = m[1][1:]
	}

	// Split path by ':', which is the target name delimiter. If it appears
	// exactly once, interpret everything to its right as the target name.
	targetSplit := strings.Split(packageRel, ":")
	switch len(targetSplit) {
	case 1:
	case 2:
		packageRel = targetSplit[0]
		res.Name = targetSplit[1]
		if !rePathPart.MatchString(res.Name) {
			return nil
		}
	default:
		return nil
	}

	// Split the package path by /, and if the name was not explicitly given,
	// use the last element of the package path.
	if packageRel == "" {
		res.PackagePath = nil
	} else {
		res.PackagePath = strings.Split(packageRel, "/")
	}
	if res.Name == "" {
		res.Name = res.PackagePath[len(res.PackagePath)-1]
	}

	// Ensure all parts of the package path are valid.
	for _, p := range res.PackagePath {
		if !rePathPart.MatchString(p) {
			return nil
		}
	}

	return &res
}
