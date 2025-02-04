// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// Package noioutil contains a Go analysis pass designed to prevent use of the
// deprecated ioutil package for which a tree-wide migration was already done.
package noioutil

import (
	"strconv"

	"golang.org/x/tools/go/analysis"

	alib "source.monogon.dev/build/analysis/lib"
)

var Analyzer = &analysis.Analyzer{
	Name: "noioutil",
	Doc:  "noioutil checks for imports of the deprecated ioutil package",
	Run:  run,
}

func run(p *analysis.Pass) (interface{}, error) {
	for _, file := range p.Files {
		if alib.IsGeneratedFile(file) {
			continue
		}
		for _, i := range file.Imports {
			importPath, err := strconv.Unquote(i.Path.Value)
			if err != nil {
				continue
			}
			if importPath == "io/ioutil" {
				p.Report(analysis.Diagnostic{
					Pos:     i.Path.ValuePos,
					End:     i.Path.End(),
					Message: "File imports the deprecated io/ioutil package. See https://pkg.go.dev/io/ioutil for replacements.",
				})
			}
		}
	}

	return nil, nil
}
