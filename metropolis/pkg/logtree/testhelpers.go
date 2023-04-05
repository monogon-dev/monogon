package logtree

import (
	"context"
	"testing"
)

// PipeAllToTest starts a goroutine that will forward all logtree entries
// t.Logf(), in the canonical logtree payload representation.
//
// It's designed to be used in tests, and will automatically stop when the
// test/benchmark it's running in exits.
func PipeAllToTest(t testing.TB, lt *LogTree) {
	t.Helper()

	reader, err := lt.Read("", WithChildren(), WithStream())
	if err != nil {
		t.Fatalf("Failed to set up logtree reader: %v", err)
	}

	// Internal context used to cancel the goroutine. This could also be a
	// implemented via a channel.
	ctx, ctxC := context.WithCancel(context.Background())
	t.Cleanup(ctxC)

	go func() {
		t.Helper()
		for {
			select {
			case <-ctx.Done():
				return
			case p := <-reader.Stream:
				t.Logf("%s", p.String())
			}
		}
	}()
}
