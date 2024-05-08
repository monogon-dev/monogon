// Package version is a companion library to the //version/spec proto.
package version

import (
	"fmt"
	"strings"

	"source.monogon.dev/version/spec"
)

// Release converts a spec.Version's Release field into a SemVer 2.0.0 compatible
// string in the X.Y.Z form.
func Release(rel *spec.Version_Release) string {
	if rel == nil {
		return "0.0.0"
	}
	return fmt.Sprintf("%d.%d.%d", rel.Major, rel.Minor, rel.Patch)
}

// Semver converts a spec.Version proto message into a SemVer 2.0.0 compatible
// string.
func Semver(v *spec.Version) string {
	ver := "v" + Release(v.Release)
	var prerelease []string
	if git := v.GitInformation; git != nil {
		if n := git.CommitsSinceRelease; n != 0 {
			prerelease = append(prerelease, fmt.Sprintf("dev%d", n))
		}
		prerelease = append(prerelease, fmt.Sprintf("g%s", git.CommitHash[:8]))
		if git.BuildTreeState != spec.Version_GitInformation_BUILD_TREE_STATE_CLEAN {
			prerelease = append(prerelease, "dirty")
		}
	}

	if len(prerelease) > 0 {
		ver += "-" + strings.Join(prerelease, ".")
	}
	return ver
}

// ReleaseLessThan returns true if Release a is lexicographically smaller than b.
func ReleaseLessThan(a, b *spec.Version_Release) bool {
	if a.Major != b.Major {
		return a.Major < b.Major
	}
	if a.Minor != b.Minor {
		return a.Minor < b.Minor
	}
	if a.Patch != b.Patch {
		return a.Patch < b.Patch
	}
	return false
}
