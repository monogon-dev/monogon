// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"context"
	"errors"
	"log"
	"sort"
	"testing"
	"time"

	"source.monogon.dev/osbase/logtree"
)

// TestHarness runs a supervisor in a harness designed for unit testing
// runnables and runnable trees.
//
// The given runnable will be run in a new supervisor, and the logs from this
// supervisor will be streamed to stderr. If the runnable returns a non-context
// error, the harness will throw a test error, but will not abort the test.
//
// The harness also returns a context cancel function that can be used to
// terminate the started supervisor early. Regardless of manual cancellation,
// the supervisor will always be terminated up at the end of the test/benchmark
// it's running in. The supervision tree will also be cleaned up and the test
// will block until all runnables have exited.
//
// The second returned value is the logtree used by this supervisor. It can be
// used to assert some log messages are emitted in tests that exercise some
// log-related functionality.
func TestHarness(t testing.TB, r func(ctx context.Context) error) (context.CancelFunc, *logtree.LogTree) {
	t.Helper()

	ctx, ctxC := context.WithCancel(context.Background())

	lt := logtree.New()

	// Only log to stderr when we're running in a test, not in a fuzz harness or a
	// benchmark - otherwise we just waste CPU cycles.
	verbose := false
	if _, ok := t.(*testing.T); ok {
		verbose = true
	}
	if verbose {
		logtree.PipeAllToTest(t, lt)
	}

	sup := New(ctx, func(ctx context.Context) error {
		Logger(ctx).Infof("Starting test %s...", t.Name())
		if err := r(ctx); err != nil && !errors.Is(err, ctx.Err()) {
			t.Errorf("Supervised runnable in harness returned error: %v", err)
			return err
		}
		return nil
	}, WithExistingLogtree(lt), WithPropagatePanic)

	t.Cleanup(func() {
		ctxC()
		if verbose {
			log.Printf("supervisor.TestHarness: Waiting for supervisor runnables to die...")
		}
		timeoutNag := time.Now().Add(5 * time.Second)

		for {
			live := sup.liveRunnables()
			if len(live) == 0 {
				if verbose {
					log.Printf("supervisor.TestHarness: All done.")
				}
				return
			}

			if time.Now().After(timeoutNag) {
				timeoutNag = time.Now().Add(5 * time.Second)
				sort.Strings(live)
				if verbose {
					log.Printf("supervisor.TestHarness: Still live:")
					for _, l := range live {
						log.Printf("supervisor.TestHarness: - %s", l)
					}
				}
			}

			time.Sleep(10 * time.Millisecond)
		}
	})
	return ctxC, lt
}
