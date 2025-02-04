// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package fsspec

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/prototext"
)

// ReadMergeSpecs reads FSSpecs from all files in paths and merges them into
// a single FSSpec.
func ReadMergeSpecs(paths []string) (*FSSpec, error) {
	var mergedSpec FSSpec
	for _, p := range paths {
		specRaw, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("failed to open spec: %w", err)
		}

		var spec FSSpec
		if err := prototext.Unmarshal(specRaw, &spec); err != nil {
			return nil, fmt.Errorf("failed to parse spec %q: %w", p, err)
		}
		mergedSpec.File = append(mergedSpec.File, spec.File...)
		mergedSpec.Directory = append(mergedSpec.Directory, spec.Directory...)
		mergedSpec.SymbolicLink = append(mergedSpec.SymbolicLink, spec.SymbolicLink...)
		mergedSpec.SpecialFile = append(mergedSpec.SpecialFile, spec.SpecialFile...)
	}
	return &mergedSpec, nil
}
