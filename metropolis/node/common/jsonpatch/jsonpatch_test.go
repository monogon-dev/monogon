// Copyright 2020 The Monogon Project Authors.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jsonpatch

import (
	"testing"
)

func TestEncodeJSONRefToken(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  string
	}{
		{"Passes through normal characters", "asdf123", "asdf123"},
		{"Encodes simple slashes", "a/b", "a~1b"},
		{"Encodes tildes", "m~n", "m~0n"},
		{"Encodes bot tildes and slashes", "a/m~n", "a~1m~0n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeJSONRefToken(tt.token); got != tt.want {
				t.Errorf("EncodeJSONRefToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointerFromParts(t *testing.T) {
	type args struct {
		pathParts []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty path", args{[]string{}}, ""},
		{"Single level path", args{[]string{"foo"}}, "/foo"},
		{"Multi-level path", args{[]string{"foo", "0"}}, "/foo/0"},
		{"Path starting with empty key", args{[]string{""}}, "/"},
		{"Path with part containing /", args{[]string{"a/b"}}, "/a~1b"},
		{"Path with part containing spaces", args{[]string{" "}}, "/ "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointerFromParts(tt.args.pathParts); got != tt.want {
				t.Errorf("PointerFromParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
