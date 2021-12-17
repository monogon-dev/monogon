// importsort implements import grouping style checks as per
// CODING_STANDARDS.md.
package importsort

import (
	"fmt"
	"go/ast"
	"go/token"
	"sort"
	"strconv"

	"golang.org/x/tools/go/analysis"

	alib "source.monogon.dev/build/analysis/lib"
)

var Analyzer = &analysis.Analyzer{
	Name: "importsort",
	Doc:  "importsort ensures imports are properly sorted and grouped",
	Run:  run,
}

func run(p *analysis.Pass) (interface{}, error) {
	for _, file := range p.Files {
		if alib.IsGeneratedFile(file) {
			continue
		}
		imp := getImportBlock(p, file)
		if imp == nil {
			continue
		}
		ensureSorted(p, imp, file, p.Fset)
	}
	return nil, nil
}

// getImportBlock returns the 'main' import block from a file. If more than one
// import block is present, the first one is returned and an diagnostic is
// added. If no import blocks are present, nil is returned.
func getImportBlock(p *analysis.Pass, f *ast.File) *ast.GenDecl {
	var first *ast.GenDecl
	for _, decl := range f.Decls {
		// Only interested in import declarations that aren't CGO import blocks.
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.IMPORT {
			continue
		}
		for _, spec := range gen.Specs {
			path, _ := strconv.Unquote(spec.(*ast.ImportSpec).Path.Value)
			if path == "C" {
				continue
			}
		}

		// Got our first import block.
		if first == nil {
			first = gen
			continue
		}

		// Second import block. Shouldn't happen.
		p.Report(analysis.Diagnostic{
			Pos:     gen.Pos(),
			End:     gen.End(),
			Message: "more than one import block",
		})
	}
	return first
}

// ensureSorted performs a style pass on a given import block/decl, reporting
// any issues found.
func ensureSorted(p *analysis.Pass, gen *ast.GenDecl, f *ast.File, fset *token.FileSet) {
	// Not a block but a single import - nothing to do here.
	if !gen.Lparen.IsValid() {
		return
	}

	// Find comment lines. These are the only entries allowed in an import block
	// apart from actual imports and newlines, so we need to know them to figure out
	// where newlines are, which in turn will be used to split imports into groups.
	commentLines := make(map[int]bool)
	for _, comment := range f.Comments {
		line := fset.Position(comment.Pos()).Line
		commentLines[line] = true
	}

	// Split imports into groups, where each group contains a list of indices into
	// gen.Specs.
	var groups [][]int
	var curGroup []int
	for i, spec := range gen.Specs {
		line := fset.Position(spec.Pos()).Line

		// First group.
		if len(curGroup) == 0 {
			curGroup = []int{i}
			continue
		}
		prevInGroup := curGroup[len(curGroup)-1]

		// Check for difference between the line number of this import and the expected
		// next line per the last recorded import.
		expectedNext := fset.Position(gen.Specs[prevInGroup].Pos()).Line + 1

		// No extra lines between this decl and expected decl per previous decl. Still
		// part of the same group.
		if line == expectedNext {
			curGroup = append(curGroup, i)
			continue
		}

		// Some lines between previous spec and this spec. If they're not all comments,
		// this makes a new group.
		allComments := true
		for j := expectedNext; j < line; j++ {
			if !commentLines[j] {
				allComments = false
				break
			}
		}
		if !allComments {
			groups = append(groups, curGroup)
			curGroup = []int{i}
			continue
		}

		// All lines in between were comments. Still part of the same group.
		curGroup = append(curGroup, i)
	}
	// Close last group.
	if len(curGroup) > 0 {
		groups = append(groups, curGroup)
	}

	// This shouldn't happened, but let's just make sure.
	if len(groups) == 0 {
		return
	}

	// Helper function to report a diagnoses on a given group.
	reportGroup := func(i int, msg string) {
		group := groups[i]
		groupStart := gen.Specs[group[0]].Pos()
		groupEnd := gen.Specs[group[len(group)-1]].End()
		p.Report(analysis.Diagnostic{
			Pos:     groupStart,
			End:     groupEnd,
			Message: msg,
		})

	}

	// Imports are now grouped. Ensure each group individually is sorted. Also use
	// this pass to classify all groups into kinds (stdlib, global, local).
	groupClasses := make([]groupClass, len(groups))
	mixed := false
	for i, group := range groups {
		importNames := make([]string, len(group))
		for i, j := range group {
			spec := gen.Specs[j]
			path := spec.(*ast.ImportSpec).Path.Value
			path, err := strconv.Unquote(path)
			if err != nil {
				p.Report(analysis.Diagnostic{
					Pos:     spec.Pos(),
					End:     spec.End(),
					Message: fmt.Sprintf("could not unquote import: %v", err),
				})
			}
			importNames[i] = path
		}
		groupClasses[i] = classifyImportGroup(importNames)

		if !sort.StringsAreSorted(importNames) {
			reportGroup(i, "imports within group are not sorted")
		}
		if groupClasses[i] == groupClassMixed {
			reportGroup(i, "import classes within group are mixed")
			mixed = true
		}
	}

	// If we had any mixed up group, abort here and let the user figure that out
	// first.
	if mixed {
		return
	}

	// Ensure group classes are in the right order.
	seenGlobal := false
	seenLocal := false
	for i, class := range groupClasses {
		switch class {
		case groupClassStdlib:
			if seenGlobal || seenLocal {
				reportGroup(i, "stdlib import group after non-stdlib import group")
			}
		case groupClassGlobal:
			if seenLocal {
				reportGroup(i, "global import group after local import group")
			}
			seenGlobal = true
		case groupClassLocal:
			seenLocal = true
		}
	}
}
