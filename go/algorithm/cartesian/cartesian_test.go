// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package cartesian

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProduct(t *testing.T) {
	for i, te := range []struct {
		data [][]string
		want [][]string
	}{
		{
			data: [][]string{
				{"a", "b"},
				{"c", "d"},
			},
			want: [][]string{
				{"a", "c"},
				{"a", "d"},
				{"b", "c"},
				{"b", "d"},
			},
		},
		{
			data: [][]string{
				{"a", "b"},
			},
			want: [][]string{
				{"a"},
				{"b"},
			},
		},
		{
			data: [][]string{},
			want: nil,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := Product(te.data...)
			if diff := cmp.Diff(te.want, got); diff != "" {
				t.Fatalf("Diff:\n%s", diff)
			}
		})
	}
}
