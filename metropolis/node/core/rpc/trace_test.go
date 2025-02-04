// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package rpc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"source.monogon.dev/osbase/logtree"
)

// TestSpanRecording exercises the span->logtree forwarding functionality by
// adding an event to the span and expecting to find it as a log entry.
func TestSpanRecording(t *testing.T) {
	lt := logtree.New()
	span := newLogtreeSpan(lt.MustLeveledFor("test"))
	span.Printf("hello world")

	r, err := lt.Read("test", logtree.WithBacklog(logtree.BacklogAllAvailable))
	if err != nil {
		t.Fatalf("logtree read failed: %v", err)
	}
	defer r.Close()
	found := false
	needle := fmt.Sprintf("Span %x: hello world", span.uid)
	for _, e := range r.Backlog {
		if e.DN != "test" {
			continue
		}
		if e.Leveled == nil {
			continue
		}
		if e.Leveled.MessagesJoined() != needle {
			continue
		}
		if parts := strings.Split(e.Leveled.Location(), ":"); parts[0] != "trace_test.go" {
			t.Errorf("Trace/log location is %s, wanted something in trace_test.go", e.Leveled.Location())
		}
		found = true
		break
	}
	if !found {
		t.Fatalf("did not find expected logline")
	}
}

// TestSpanContext exercises a span context injection/extraction roundtrip.
func TestSpanContext(t *testing.T) {
	ctx := context.Background()

	lt := logtree.New()
	span := newLogtreeSpan(lt.MustLeveledFor("test"))
	ctx = contextWithSpan(ctx, span)
	span2 := Trace(ctx)
	if !span2.IsRecording() {
		t.Errorf("Expected span to be active")
	}

	v, ok := span2.(*logtreeSpan)
	if !ok {
		t.Fatalf("Retrieved span is not *logtreeSpan")
	}
	if v != span {
		t.Fatalf("Retrieved span differs from injected span")
	}
}

// TestSpanContextFallback exercises an empty span retrieved from a context with
// no span set.
func TestSpanContextFallback(t *testing.T) {
	ctx := context.Background()
	// We expect this to never panic, just to drop any event.
	Trace(ctx).Printf("plonk")
	if Trace(ctx).IsRecording() {
		t.Errorf("Expected span to be inactive")
	}
}
