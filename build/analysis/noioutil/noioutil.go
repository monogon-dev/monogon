// Package noioutil contains a Go analysis pass designed to prevent use of the
// deprecated ioutil package for which a tree-wide migration was already done.
package noioutil

import (
	"go/ast"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "noioutil",
	Doc:  "noioutil checks for imports of the deprecated ioutil package",
	Run:  run,
}

func run(p *analysis.Pass) (interface{}, error) {
	for _, file := range p.Files {
		if isGeneratedFile(file) {
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

const (
	genPrefix = "// Code generated"
	genSuffix = "DO NOT EDIT."
)

// isGeneratedFile returns true if the file is generated
// according to https://golang.org/s/generatedcode.
func isGeneratedFile(file *ast.File) bool {
	for _, c := range file.Comments {
		for _, t := range c.List {
			if strings.HasPrefix(t.Text, genPrefix) && strings.HasSuffix(t.Text, genSuffix) {
				return true
			}
		}
	}
	return false
}
