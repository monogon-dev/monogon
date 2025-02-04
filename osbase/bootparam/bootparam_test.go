// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

// If this is bootparam we have an import cycle
package bootparam_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"source.monogon.dev/osbase/bootparam"
	"source.monogon.dev/osbase/bootparam/ref"
)

// Fuzzers can be run with
// bazel test //osbase/bootparam:bootparam_test
//   --test_arg=-test.fuzz=FuzzMarshal
//   --test_arg=-test.fuzzcachedir=/tmp/fuzz
//   --test_arg=-test.fuzztime=60s

func FuzzUnmarshal(f *testing.F) {
	f.Add(`initrd="\test\some=value" root=yolo "definitely quoted" ro rootflags=`)
	f.Fuzz(func(t *testing.T, a string) {
		refOut, refRest := ref.Parse(a)
		out, rest, err := bootparam.Unmarshal(a)
		if err != nil {
			return
		}
		if diff := cmp.Diff(refOut, out); diff != "" {
			t.Errorf("Parse(%q): params mismatch (-want +got):\n%s", a, diff)
		}
		if refRest != rest {
			t.Errorf("Parse(%q): expected rest to be %q, got %q", a, refRest, rest)
		}
	})
}

func FuzzMarshal(f *testing.F) {
	// Choose delimiters which mean nothing to the parser
	f.Add("a:b;assd:9dsf;1234", "some fancy rest")
	f.Fuzz(func(t *testing.T, paramsRaw string, rest string) {
		paramsSeparated := strings.Split(paramsRaw, ";")
		var params bootparam.Params
		for _, p := range paramsSeparated {
			a, b, _ := strings.Cut(p, ":")
			params = append(params, bootparam.Param{Param: a, Value: b})
		}
		rest = bootparam.TrimLeftSpace(rest)
		encoded, err := bootparam.Marshal(params, rest)
		if err != nil {
			return // Invalid input
		}
		refOut, refRest := ref.Parse(encoded)
		if diff := cmp.Diff(refOut, params); diff != "" {
			t.Errorf("Marshal(%q): params mismatch (-want +got):\n%s", paramsRaw, diff)
		}
		if refRest != rest {
			t.Errorf("Parse(%q, %q): expected rest to be %q, got %q", paramsRaw, rest, refRest, rest)
		}
	})
}
