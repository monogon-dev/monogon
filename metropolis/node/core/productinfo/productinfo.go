// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package productinfo

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/coreos/go-semver/semver"

	vpb "source.monogon.dev/version/spec"

	"source.monogon.dev/osbase/oci/osimage"
	"source.monogon.dev/version"
)

// ProductInfo is a wrapper of [osimage.ProductInfo] with some additional
// representations of the same data.
type ProductInfo struct {
	// Info is the product info parsed from JSON.
	Info *osimage.ProductInfo
	// Version parsed as [vpb.Version]
	Version *vpb.Version
	// VersionString is Info.Version with short commit and dirty indicator.
	VersionString string
	// HumanCommitDate is Info.CommitDate formatted to be more readable, in UTC.
	// This is empty if stamping is disabled.
	HumanCommitDate string
}

// path to the product info file. This may be replaced by x_defs in tests, and
// will then be resolved with runfiles.
var path = "/etc/product-info.json"

// Get returns the product info of the running system.
var Get = sync.OnceValue(func() *ProductInfo {
	resolvedPath := path
	if resolvedPath[0] != '/' {
		var err error
		resolvedPath, err = runfiles.Rlocation(resolvedPath)
		if err != nil {
			panic(err)
		}
	}
	productInfo, err := read(resolvedPath)
	if err != nil {
		panic(err)
	}
	return productInfo
})

// read parses the JSON file at the given path as [osimage.ProductInfo].
func read(path string) (*ProductInfo, error) {
	rawProductInfo, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var productInfo *osimage.ProductInfo
	err = json.Unmarshal(rawProductInfo, &productInfo)
	if err != nil {
		return nil, err
	}

	specVersion, err := versionFromProductInfo(productInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to extract version: %w", err)
	}

	var humanCommitDate string
	if productInfo.CommitDate != "" {
		commitDate, err := time.Parse(time.RFC3339, productInfo.CommitDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse commit date: %w", err)
		}
		humanCommitDate = commitDate.UTC().Format(time.DateTime)
	}

	info := &ProductInfo{
		Info:            productInfo,
		Version:         specVersion,
		VersionString:   version.Semver(specVersion),
		HumanCommitDate: humanCommitDate,
	}
	return info, nil
}

var rePrereleaseCommitOffset = regexp.MustCompile(`^dev([0-9]+)$`)

func versionFromProductInfo(productInfo *osimage.ProductInfo) (*vpb.Version, error) {
	version := &vpb.Version{}

	if productInfo.CommitHash != "" {
		if len(productInfo.CommitHash) < 8 {
			return nil, fmt.Errorf("git commit hash too short")
		}
		buildTreeState := vpb.Version_GitInformation_BUILD_TREE_STATE_CLEAN
		if productInfo.BuildTreeDirty {
			buildTreeState = vpb.Version_GitInformation_BUILD_TREE_STATE_DIRTY
		}
		version.GitInformation = &vpb.Version_GitInformation{
			CommitHash:     productInfo.CommitHash[:8],
			BuildTreeState: buildTreeState,
		}
	}

	if productInfo.Version != "" {
		v, err := semver.NewVersion(productInfo.Version)
		if err != nil {
			return nil, fmt.Errorf("invalid %s version %q: %w", productInfo.ID, productInfo.Version, err)
		}
		// Parse prerelease strings (v1.2.3-foo-bar -> [foo, bar])
		for _, el := range v.PreRelease.Slice() {
			preCommitOffset := rePrereleaseCommitOffset.FindStringSubmatch(el)
			switch {
			case el == "":
				// Skip empty slices which happens when there's a semver string with no
				// prerelease data.
			case el == "nostamp":
				// Ignore field, we have it from CommitHash.
			case preCommitOffset != nil:
				offset, err := strconv.ParseUint(preCommitOffset[1], 10, 64)
				if err != nil {
					return nil, fmt.Errorf("invalid commit offset value: %w", err)
				}
				if version.GitInformation == nil {
					return nil, fmt.Errorf("have git offset but no git commit")
				}
				version.GitInformation.CommitsSinceRelease = offset
			default:
				return nil, fmt.Errorf("invalid prerelease string %q (in %q)", el, productInfo.Version)
			}
		}
		version.Release = &vpb.Version_Release{
			Major: v.Major,
			Minor: v.Minor,
			Patch: v.Patch,
		}
	}
	return version, nil
}
