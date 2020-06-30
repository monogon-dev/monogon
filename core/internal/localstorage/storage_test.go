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

package localstorage

import (
	"testing"

	"git.monogon.dev/source/nexantic.git/core/internal/localstorage/declarative"
)

func TestValidateAll(t *testing.T) {
	r := Root{}
	if err := declarative.Validate(&r); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestPlaceFS(t *testing.T) {
	rr := Root{}
	err := declarative.PlaceFS(&rr, "")
	if err != nil {
		t.Errorf("Placement failed: %v", err)
	}

	// Re-placing should fail.
	err = declarative.PlaceFS(&rr, "/foo")
	if err == nil {
		t.Errorf("Re-placement didn't fail")
	}

	// Check some absolute paths.
	for i, te := range []struct {
		pl   declarative.Placement
		want string
	}{
		{rr.ESP, "/esp"},
		{rr.Data.Etcd, "/data/etcd"},
		{rr.Data.Node.Certificate, "/data/node_pki/cert.pem"},
	} {
		if got, want := te.pl.FullPath(), te.want; got != want {
			t.Errorf("test %d: wanted path %q, got %q", i, want, got)
		}
	}
}
