// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package haslicense

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/analysis"

	alib "source.monogon.dev/build/analysis/lib"
)

const header = `// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

`

var Analyzer = &analysis.Analyzer{
	Name: "haslicense",
	Doc:  "haslicense checks for an existing license header in monogon source code.",
	Run: func(p *analysis.Pass) (any, error) {
		for _, file := range p.Files {
			if alib.IsGeneratedFile(file) {
				continue
			}

			if len(file.Comments) > 0 {
				var sb strings.Builder
				for _, c := range file.Comments[0].List {
					sb.WriteString(c.Text)
					sb.WriteRune('\n')
				}
				sb.WriteRune('\n')

				if strings.HasPrefix(sb.String(), header) {
					continue
				}
			}

			p.Report(analysis.Diagnostic{
				Pos:     file.FileStart,
				End:     file.FileStart,
				Message: "File is missing license header. Please add it.",
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: fmt.Sprintf("should prepend file with `%s`", header),
						TextEdits: []analysis.TextEdit{
							{
								Pos:     file.FileStart,
								End:     file.FileStart,
								NewText: []byte(header),
							},
						},
					},
				},
			})
		}

		return nil, nil
	},
}
