// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"source.monogon.dev/osbase/event"
)

// TestAsync exercises the high-level behaviour of a Value, in which a
// watcher is able to catch up to the newest Set value.
func TestAsync(t *testing.T) {
	p := Value[int]{}
	p.Set(0)

	ctx := context.Background()

	// The 0 from Set() should be available via .Get().
	watcher := p.Watch()
	val, err := watcher.Get(ctx)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if want, got := 0, val; want != got {
		t.Fatalf("Value: got %d, wanted %d", got, want)
	}

	// Send a large amount of updates that the watcher does not actively .Get().
	for i := 1; i <= 100; i++ {
		p.Set(i)
	}

	// The watcher should still end up with the newest .Set() value on the next
	// .Get() call.
	val, err = watcher.Get(ctx)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if want, got := 100, val; want != got {
		t.Fatalf("Value: got %d, wanted %d", got, want)
	}
}

// TestMultipleGets verifies that calling .Get() on a single watcher from two
// goroutines is prevented by returning an error in exactly one of them.
func TestMultipleGets(t *testing.T) {
	p := Value[int]{}
	ctx := context.Background()

	w := p.Watch()

	tryError := func(errs chan error) {
		_, err := w.Get(ctx)
		errs <- err
	}
	errs := make(chan error, 2)
	go tryError(errs)
	go tryError(errs)

	for err := range errs {
		if err == nil {
			t.Fatalf("A Get call succeeded, while it should have blocked or returned an error")
		} else {
			// Found the error, test succeeded.
			break
		}
	}
}

// TestConcurrency attempts to stress the Value/Watcher
// implementation to design limits (a hundred simultaneous watchers), ensuring
// that the watchers all settle to the final set value.
func TestConcurrency(t *testing.T) {
	ctx := context.Background()

	p := Value[int]{}
	p.Set(0)

	// Number of watchers to create.
	watcherN := 100
	// Expected final value to be Set().
	final := 100
	// Result channel per watcher.
	resC := make([]chan error, watcherN)

	// Spawn watcherN watchers.
	for i := 0; i < watcherN; i++ {
		resC[i] = make(chan error, 1)
		go func(id int) {
			// done is a helper function that will put an error on the
			// respective watcher's resC.
			done := func(err error) {
				resC[id] <- err
				close(resC[id])
			}

			watcher := p.Watch()
			// prev is used to ensure the values received are monotonic.
			prev := -1
			for {
				val, err := watcher.Get(ctx)
				if err != nil {
					done(err)
					return
				}

				// Ensure monotonicity of received data.
				if val <= prev {
					done(fmt.Errorf("received out of order data: %d after %d", val, prev))
				}
				prev = val

				// Quit when the final value is received.
				if val == final {
					done(nil)
					return
				}

				// Sleep a bit, depending on the watcher. This makes each
				// watcher behave slightly differently, and attempts to
				// exercise races dependent on sleep time between subsequent
				// Get calls.
				time.Sleep(time.Millisecond * time.Duration(id))
			}
		}(i)
	}

	// Set 1..final on the value.
	for i := 1; i <= final; i++ {
		p.Set(i)
	}

	// Ensure all watchers exit with no error.
	for i, c := range resC {
		err := <-c
		if err != nil {
			t.Errorf("Watcher %d returned %v", i, err)
		}
	}
}

// TestCanceling exercises whether a context canceling in a .Get() gracefully
// aborts that particular Get call, but also allows subsequent use of the same
// watcher.
func TestCanceling(t *testing.T) {
	p := Value[int]{}

	ctx, ctxC := context.WithCancel(context.Background())

	watcher := p.Watch()

	// errs will contain the error returned by Get.
	errs := make(chan error, 1)
	go func() {
		// This Get will block, as no initial data has been Set on the value.
		_, err := watcher.Get(ctx)
		errs <- err
	}()

	// Cancel the context, and expect that context error to propagate to the .Get().
	ctxC()
	if want, got := ctx.Err(), <-errs; !errors.Is(got, want) {
		t.Fatalf("Get should've returned %v, got %v", want, got)
	}

	// Do another .Get() on the same watcher with a new context. Even though the
	// call was aborted via a context cancel, the watcher should continue working.
	ctx = context.Background()
	go func() {
		_, err := watcher.Get(ctx)
		errs <- err
	}()

	// Unblock the .Get now.
	p.Set(1)
	if want, got := error(nil), <-errs; !errors.Is(got, want) {
		t.Fatalf("Get should've returned %v, got %v", want, got)
	}
}

// TestSetAfterWatch ensures that if a value is updated between a Watch and the
// initial Get, only the newest Set value is returns.
func TestSetAfterWatch(t *testing.T) {
	ctx := context.Background()

	p := Value[int]{}
	p.Set(0)

	watcher := p.Watch()
	p.Set(1)

	data, err := watcher.Get(ctx)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if want, got := 1, data; want != got {
		t.Errorf("Get should've returned %v, got %v", want, got)
	}
}

// TestWatchersList ensures that the list of watchers is managed correctly,
// i.e. there is no memory leak and closed watchers are removed while
// keeping all non-closed watchers.
func TestWatchersList(t *testing.T) {
	ctx := context.Background()
	p := Value[int]{}

	var watchers []event.Watcher[int]
	for i := 0; i < 100; i++ {
		watchers = append(watchers, p.Watch())
	}
	for i := 0; i < 10000; i++ {
		watchers[10].Close()
		watchers[10] = p.Watch()
	}

	if want, got := 1000, cap(p.watchers); want <= got {
		t.Fatalf("Got capacity %d, wanted less than %d", got, want)
	}

	p.Set(1)
	if want, got := 100, len(p.watchers); want != got {
		t.Fatalf("Got %d watchers, wanted %d", got, want)
	}

	for _, watcher := range watchers {
		data, err := watcher.Get(ctx)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}
		if want, got := 1, data; want != got {
			t.Errorf("Get should've returned %v, got %v", want, got)
		}
	}
}
