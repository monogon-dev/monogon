// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package toolbase

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBazelLabelParse(t *testing.T) {
	for i, te := range []struct {
		p string
		t *BazelLabel
	}{
		{"//foo/bar", &BazelLabel{"@", []string{"foo", "bar"}, "bar"}},
		{"//foo/bar:baz", &BazelLabel{"@", []string{"foo", "bar"}, "baz"}},
		{"//:foo", &BazelLabel{"@", nil, "foo"}},

		{"@test//foo/bar", &BazelLabel{"test", []string{"foo", "bar"}, "bar"}},
		{"@test//foo/bar:baz", &BazelLabel{"test", []string{"foo", "bar"}, "baz"}},
		{"@test//:foo", &BazelLabel{"test", nil, "foo"}},

		{"", nil},
		{"//", nil},
		{"//foo:bar/foo", nil},
		{"//foo//bar/foo", nil},
		{"/foo/bar/foo", nil},
		{"foo/bar/foo", nil},
		{"@//foo/bar/foo", nil},
		{"@foo/bar//foo/bar/foo", nil},
		{"@foo:bar//foo/bar/foo", nil},
		{"foo//foo/bar/foo", nil},
	} {
		want := te.t
		got := ParseBazelLabel(te.p)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("case %d (%q): %s", i, te.p, diff)
		}
	}
}

func TestBazelLabelString(t *testing.T) {
	for i, te := range []struct {
		in   string
		want string
	}{
		{"//foo/bar", "@@//foo/bar:bar"},
		{"//foo:bar", "@@//foo:bar"},
		{"@com_github_example//:run", "@com_github_example//:run"},
	} {
		l := ParseBazelLabel(te.in)
		if l == nil {
			t.Errorf("case %d: wanted %q, got nil", i, te.want)
			continue
		}
		if want, got := te.want, l.String(); want != got {
			t.Errorf("case %d: wanted %q, got %q", i, want, got)
		}
	}
}
