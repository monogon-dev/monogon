package kmod

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"unicode"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	kmodpb "source.monogon.dev/osbase/kmod/spec"
)

func TestParsePattern(t *testing.T) {
	cases := []struct {
		name          string
		pattern       string
		expectedNodes []*kmodpb.RadixNode
	}{
		{"Empty", "", nil},
		{"SingleLiteral", "asdf", []*kmodpb.RadixNode{{Type: kmodpb.RadixNode_TYPE_LITERAL, Literal: "asdf"}}},
		{"SingleWildcard", "as*df", []*kmodpb.RadixNode{
			{Type: kmodpb.RadixNode_TYPE_LITERAL, Literal: "as"},
			{Type: kmodpb.RadixNode_TYPE_WILDCARD},
			{Type: kmodpb.RadixNode_TYPE_LITERAL, Literal: "df"},
		}},
		{"EscapedWildcard", "a\\*", []*kmodpb.RadixNode{{Type: kmodpb.RadixNode_TYPE_LITERAL, Literal: "a*"}}},
		{"SingleRange", "[y-z]", []*kmodpb.RadixNode{{Type: kmodpb.RadixNode_TYPE_BYTE_RANGE, StartByte: 121, EndByte: 122}}},
		{"SingleWildcardChar", "a?c", []*kmodpb.RadixNode{
			{Type: kmodpb.RadixNode_TYPE_LITERAL, Literal: "a"},
			{Type: kmodpb.RadixNode_TYPE_SINGLE_WILDCARD},
			{Type: kmodpb.RadixNode_TYPE_LITERAL, Literal: "c"},
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out, err := parsePattern(c.pattern)
			if err != nil {
				t.Fatal(err)
			}
			diff := cmp.Diff(c.expectedNodes, out, protocmp.Transform())
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestLookupComplex(t *testing.T) {
	root := &kmodpb.RadixNode{
		Type: kmodpb.RadixNode_TYPE_LITERAL,
	}
	if err := AddPattern(root, "usb:v0B95p1790d*dc*dsc*dp*icFFiscFFip00in*", 2); err != nil {
		t.Error(err)
	}
	if err := AddPattern(root, "usb:v0B95p178Ad*dc*dsc*dp*icFFiscFFip00in*", 3); err != nil {
		t.Error(err)
	}
	if err := AddPattern(root, "acpi*:PNP0C14:*", 10); err != nil {
		t.Error(err)
	}
	matches := make(map[uint32]bool)
	lookupModulesRec(root, "acpi:PNP0C14:asdf", matches)
	if !matches[10] {
		t.Error("value should match pattern 10")
	}
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func FuzzRadixImpl(f *testing.F) {
	f.Add("acpi*:PNP0C14:*\x00usb:v0B95p1790d*dc*dsc*dp*icFFiscFFip00in*", "acpi:PNP0C14:asdf\x00usb:v0B95p1790d0dc0dsc0dp0icFFiscFFip00in")
	f.Fuzz(func(t *testing.T, a string, b string) {
		patternsRaw := strings.Split(a, "\x00")
		values := strings.Split(b, "\x00")
		var patternsRegexp []regexp.Regexp
		root := &kmodpb.RadixNode{
			Type: kmodpb.RadixNode_TYPE_LITERAL,
		}
		for i, p := range patternsRaw {
			if !isASCII(p) {
				// Ignore non-ASCII patterns, there are tons of edge cases with them
				return
			}
			pp, err := parsePattern(p)
			if err != nil {
				// Bad pattern
				return
			}
			if err := AddPattern(root, p, uint32(i)); err != nil {
				t.Fatal(err)
			}
			var regexb strings.Builder
			regexb.WriteString("(?s)^")
			for _, part := range pp {
				switch part.Type {
				case kmodpb.RadixNode_TYPE_LITERAL:
					regexb.WriteString(regexp.QuoteMeta(part.Literal))
				case kmodpb.RadixNode_TYPE_SINGLE_WILDCARD:
					regexb.WriteString(".")
				case kmodpb.RadixNode_TYPE_WILDCARD:
					regexb.WriteString(".*")
				case kmodpb.RadixNode_TYPE_BYTE_RANGE:
					regexb.WriteString(fmt.Sprintf("[%s-%s]", regexp.QuoteMeta(string([]rune{rune(part.StartByte)})), regexp.QuoteMeta(string([]rune{rune(part.EndByte)}))))
				default:
					t.Errorf("Unknown node type %v", part.Type)
				}
			}
			regexb.WriteString("$")
			patternsRegexp = append(patternsRegexp, *regexp.MustCompile(regexb.String()))
		}
		for _, v := range values {
			if !isASCII(v) {
				// Ignore non-ASCII values
				return
			}
			if len(v) > 64 {
				// Ignore big values as they are not realistic and cause the
				// wildcard matches to be very expensive.
				return
			}
			radixMatchesSet := make(map[uint32]bool)
			lookupModulesRec(root, v, radixMatchesSet)
			for i, re := range patternsRegexp {
				if re.MatchString(v) {
					if !radixMatchesSet[uint32(i)] {
						t.Errorf("Pattern %q is expected to match %q but didn't", patternsRaw[i], v)
					}
				} else {
					if radixMatchesSet[uint32(i)] {
						t.Errorf("Pattern %q is not expected to match %q but did", patternsRaw[i], v)
					}
				}
			}
		}
	})
}
