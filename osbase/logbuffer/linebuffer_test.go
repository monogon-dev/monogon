// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package logbuffer

import (
	"fmt"
	"testing"
)

func TestLineBuffer(t *testing.T) {
	var lines []*Line
	lb := NewLineBuffer(1024, func(l *Line) {
		lines = append(lines, l)
	})

	compare := func(a []*Line, b ...string) string {
		msg := fmt.Sprintf("want %v, got %v", a, b)
		if len(a) != len(b) {
			return msg
		}
		for i := range a {
			if a[i].String() != b[i] {
				return msg
			}
		}
		return ""
	}

	// Write some data.
	fmt.Fprintf(lb, "foo ")
	if diff := compare(lines); diff != "" {
		t.Fatal(diff)
	}
	fmt.Fprintf(lb, "bar\n")
	if diff := compare(lines, "foo bar"); diff != "" {
		t.Fatal(diff)
	}
	fmt.Fprintf(lb, "baz")
	if diff := compare(lines, "foo bar"); diff != "" {
		t.Fatal(diff)
	}
	fmt.Fprintf(lb, " baz")
	if diff := compare(lines, "foo bar"); diff != "" {
		t.Fatal(diff)
	}
	// Close and expect flush.
	if err := lb.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if diff := compare(lines, "foo bar", "baz baz"); diff != "" {
		t.Fatal(diff)
	}

	// Check behaviour after close
	if _, err := fmt.Fprintf(lb, "nope"); err == nil {
		t.Fatalf("Write after Close: wanted  error, got nil")
	}
	if err := lb.Close(); err == nil {
		t.Fatalf("second Close: wanted error, got nil")
	}
}
