package importsort

import (
	"strings"

	alib "source.monogon.dev/build/analysis/lib"
)

// groupClass is the 'class' of a given import path or import group.
type groupClass string

const (
	// groupClassMixed are import group that contain multiple different classes of
	// import paths.
	groupClassMixed = "mixed"
	// groupClassStdlib is an import path or group that contains only Go standard
	// library imports.
	groupClassStdlib = "stdlib"
	// groupClassGlobal is an import path or group that contains only third-party
	// packages, ie. all packages that aren't part of stdlib and aren't local to the
	// Monogon codebase.
	groupClassGlobal = "global"
	// groupClassLocal is an import path or group that contains only package that
	// are local to the Monogon codebase.
	groupClassLocal = "local"
)

// classifyImport returns a groupClass for a given import path.
func classifyImport(path string) groupClass {
	if alib.StdlibPackages[path] {
		return groupClassStdlib
	}
	if strings.HasPrefix(path, "source.monogon.dev/") {
		return groupClassLocal
	}
	return groupClassGlobal
}

// classifyImportGroup returns a groupClass for a list of import paths that are
// part of a single import group.
func classifyImportGroup(paths []string) groupClass {
	res := groupClass("")
	for _, p := range paths {
		if res == "" {
			res = classifyImport(p)
			continue
		}
		class := classifyImport(p)
		if res != class {
			return groupClassMixed
		}
	}
	return res
}
