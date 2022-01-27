package fsspec

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
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
		if err := proto.UnmarshalText(string(specRaw), &spec); err != nil {
			return nil, fmt.Errorf("failed to parse spec %q: %w", p, err)
		}
		for _, f := range spec.File {
			mergedSpec.File = append(mergedSpec.File, f)
		}
		for _, d := range spec.Directory {
			mergedSpec.Directory = append(mergedSpec.Directory, d)
		}
		for _, s := range spec.SymbolicLink {
			mergedSpec.SymbolicLink = append(mergedSpec.SymbolicLink, s)
		}
		for _, s := range spec.SpecialFile {
			mergedSpec.SpecialFile = append(mergedSpec.SpecialFile, s)
		}
	}
	return &mergedSpec, nil
}
