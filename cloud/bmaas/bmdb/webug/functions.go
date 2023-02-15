package webug

import (
	"strings"
)

var (
	// templateFuncs are helper functions accessible to the rendered templates.
	templateFuncs = map[string]any{
		// summarizeError attempts to make a Go-style "foo: bar: baz" error short by
		// using some ugly heuristics. This is currently used to show a shorter error
		// message in the backoff column of the machine list.
		//
		// TODO(q3k): fix backoff causes to be less verbose and nuke this.
		"summarizeError": func(in string) string {
			parts := strings.Split(in, ": ")
			for i, p := range parts {
				// Attempt to strip some common error prefixes.
				if strings.HasPrefix(p, "failed to ") {
					continue
				}
				if strings.HasPrefix(p, "when ") {
					continue
				}
				if strings.HasPrefix(p, "while ") {
					continue
				}
				// If we had some prefixes stripped but suddenly reached a part that is not
				// prefixed
				return "[...] " + strings.Join(parts[i:], ": ")
			}
			// If we stripped every single segment then just return the whole thing.
			return in
		},
	}
)
