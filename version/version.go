// Package version is a companion library to the //version/spec proto.
package version

import (
	"fmt"
	"strings"

	"source.monogon.dev/version/spec"
)

// Release converts a spec.Version's Release field into a SemVer 2.0.0 compatible
// string in the X.Y.Z form.
func Release(v *spec.Version) string {
	if v == nil || v.Release == nil {
		return "0.0.0"
	}
	rel := v.Release
	return fmt.Sprintf("%d.%d.%d", rel.Major, rel.Minor, rel.Patch)
}

// Semver converts a spec.Version proto message into a SemVer 2.0.0 compatible
// string.
func Semver(v *spec.Version) string {
	ver := "v" + Release(v)
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
