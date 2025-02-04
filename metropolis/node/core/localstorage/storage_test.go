// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package localstorage

import (
	"testing"

	"source.monogon.dev/metropolis/node/core/localstorage/declarative"
)

func TestValidateAll(t *testing.T) {
	var r Root
	if err := declarative.Validate(&r); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestPlaceFS(t *testing.T) {
	var rr Root
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
		{rr.Data.Node.Credentials.Certificate, "/data/node/credentials/cert.pem"},
	} {
		if got, want := te.pl.FullPath(), te.want; got != want {
			t.Errorf("test %d: wanted path %q, got %q", i, want, got)
		}
	}
}
