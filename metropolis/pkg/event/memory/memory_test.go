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

package memory

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"source.monogon.dev/metropolis/pkg/event"
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

// TestSyncBlocks exercises the Value's 'Sync' field, which makes all
// Set() calls block until all respective watchers .Get() the updated data.
// This particular test ensures that .Set() calls to a Watcher result in a
// prefect log of updates being transmitted to a watcher.
func TestSync(t *testing.T) {
	p := Value[int]{
		Sync: true,
	}
	values := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ctx := context.Background()
		watcher := p.Watch()
		wg.Done()
		for {
			value, err := watcher.Get(ctx)
			if err != nil {
				panic(err)
			}
			values <- value
		}
	}()

	p.Set(0)
	wg.Wait()

	want := []int{1, 2, 3, 4}
	for _, w := range want {
		p.Set(w)
	}

	timeout := time.After(time.Second)
	for i, w := range append([]int{0}, want...) {
		select {
		case <-timeout:
			t.Fatalf("timed out on value %d (%d)", i, w)
		case val := <-values:
			if w != val {
				t.Errorf("value %d was %d, wanted %d", i, val, w)
			}
		}
	}
}

// TestSyncBlocks exercises the Value's 'Sync' field, which makes all
// Set() calls block until all respective watchers .Get() the updated data.
// This particular test ensures that .Set() calls actually block when a watcher
// is unattended.
func TestSyncBlocks(t *testing.T) {
	p := Value[int]{
		Sync: true,
	}
	ctx := context.Background()

	// Shouldn't block, as there's no declared watchers.
	p.Set(0)

	watcher := p.Watch()

	// Should retrieve the zero, more requests will pend.
	value, err := watcher.Get(ctx)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if want, got := 0, value; want != got {
		t.Fatalf("Got initial value %d, wanted %d", got, want)
	}

	// .Set() Should block, as watcher is unattended.
	//
	// Whether something blocks in Go is untestable in a robust way (see: halting
	// problem). We work around this this by introducing a 'stage' int64, which is
	// put on the 'c' channel after the needs-to-block function returns. We then
	// perform an action that should unblock this function right after updating
	// 'stage' to a different value.
	// Then, we observe what was put on the channel: If it's the initial value, it
	// means the function didn't block when expected. Otherwise, it means the
	// function unblocked when expected.
	stage := int64(0)
	c := make(chan int64, 1)
	go func() {
		p.Set(1)
		c <- atomic.LoadInt64(&stage)
	}()

	// Getting should unblock the provider. Mark via 'stage' variable that
	// unblocking now is expected.
	atomic.StoreInt64(&stage, int64(1))
	// Potential race: .Set() unblocks here due to some bug, before .Get() is
	// called, and we record a false positive.
	value, err = watcher.Get(ctx)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	res := <-c
	if res != int64(1) {
		t.Fatalf("Set() returned before Get()")
	}

	if want, got := 1, value; want != got {
		t.Fatalf("Wanted value %d, got %d", want, got)
	}

	// Closing the watcher and setting should not block anymore.
	if err := watcher.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	// Last step, if this blocks we will get a deadlock error and the test will panic.
	p.Set(2)
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
	p := Value[int]{
		Sync: true,
	}

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
