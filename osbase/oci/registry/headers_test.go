// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package registry

import "testing"

func TestParseAuthenticateHeader(t *testing.T) {
	testCases := []struct {
		desc   string
		header []string
		parsed []authenticateChallenge
	}{
		{"absent", nil, nil},
		{"no params",
			[]string{"Basic, !#$%&'*+-.^_`|~019abzABZ"},
			[]authenticateChallenge{{scheme: "Basic"}, {scheme: "!#$%&'*+-.^_`|~019abzABZ"}}},
		{"token68",
			[]string{"0 a", "1 abzABZ019-._~+/, 2 abc=, 3   a==="},
			[]authenticateChallenge{
				{scheme: "0", info: "a"},
				{scheme: "1", info: "abzABZ019-._~+/"},
				{scheme: "2", info: "abc="},
				{scheme: "3", info: "a==="},
			}},
		{"params",
			[]string{`0 a="=,", empty =  "", escape="\a\\\"", ` + "1 token!#$%&'*+-.^_`|~019abzABZ=!#$%&'*+-.^_`|~019abzABZ"},
			[]authenticateChallenge{
				{scheme: "0", params: map[string]string{"a": "=,", "empty": "", "escape": `a\"`}},
				{scheme: "1", params: map[string]string{"token!#$%&'*+-.^_`|~019abzabz": "!#$%&'*+-.^_`|~019abzABZ"}},
			}},
		{"duplicate param", []string{`Basic realm="apps", REALM=other`}, nil},
		{"empty", []string{"", " ", "\t", ",", " , ,,\t ,", "Basic"}, []authenticateChallenge{{scheme: "Basic"}}},
		{"RFC example",
			[]string{`Basic realm="simple", Newauth realm="apps", type=1, title="Login to \"apps\""`},
			[]authenticateChallenge{
				{scheme: "Basic", params: map[string]string{"realm": "simple"}},
				{scheme: "Newauth", params: map[string]string{"realm": "apps", "type": "1", "title": `Login to "apps"`}},
			}},
		{"extra commas",
			[]string{` , , Basic , , realm="simple" , , Newauth ,realm="apps",type=1` + "\t" + `, ,title="Login to \"apps\"" , , `},
			[]authenticateChallenge{
				{scheme: "Basic", params: map[string]string{"realm": "simple"}},
				{scheme: "Newauth", params: map[string]string{"realm": "apps", "type": "1", "title": `Login to "apps"`}},
			}},
		{"missing comma between challenges", []string{"Basic\tBearer"}, nil},
		{"missing comma between challenges 2", []string{"Basic !"}, nil},
		{"missing comma after token68", []string{"Basic a Bearer"}, nil},
		{"missing comma between params", []string{`Basic realm="simple" type=1`}, nil},
		{"missing quote", []string{`Basic realm="simple`}, nil},
		{"missing value", []string{`Basic !=`}, nil},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := parseAuthenticateHeader(tC.header)
			if want, got := len(tC.parsed), len(actual); want != got {
				t.Fatalf("Expected %d challenges, got %d", want, got)
			}
			for i, actualC := range actual {
				wantC := tC.parsed[i]
				if want, got := wantC.scheme, actualC.scheme; want != got {
					t.Errorf("Expected scheme %q, got %q", want, got)
				}
				if want, got := wantC.info, actualC.info; want != got {
					t.Errorf("Expected info %q, got %q", want, got)
				}
				for param, want := range wantC.params {
					got, ok := actualC.params[param]
					if !ok {
						t.Errorf("Scheme %s: Missing param %q", wantC.scheme, param)
					} else if want != got {
						t.Errorf("Scheme %s: Expected %s=%q, got %q", wantC.scheme, param, want, got)
					}
				}
				for param := range actualC.params {
					if _, ok := wantC.params[param]; !ok {
						t.Errorf("Scheme %s: Extra param %q", wantC.scheme, param)
					}
				}
			}
		})
	}
}
