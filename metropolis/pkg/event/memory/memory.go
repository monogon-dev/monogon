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
	"fmt"
	"sync"

	"source.monogon.dev/metropolis/pkg/event"
)

var (
	// Type assert that *Value implements Value. We do this artificially, as
	// there currently is no code path that needs this to be strictly true. However,
	// users of this library might want to rely on the Value type instead of
	// particular Value implementations.
	_ event.Value = &Value{}
)

// Value is a 'memory value', which implements a event.Value stored in memory.
// It is safe to construct an empty object of this type. However, this must not
// be copied.
type Value struct {
	// mu guards the inner, innerSet and watchers fields.
	mu sync.RWMutex
	// inner is the latest data Set on the Value. It is used to provide the
	// newest version of the Set data to new watchers.
	inner interface{}
	// innerSet is true when inner has been Set at least once. It is used to
	// differentiate between a nil and unset value.
	innerSet bool
	// watchers is the list of watchers that should be updated when new data is
	// Set. It will grow on every .Watch() and shrink any time a watcher is
	// determined to have been closed.
	watchers []*watcher

	// Sync, if set to true, blocks all .Set() calls on the Value until all
	// Watchers derived from it actively .Get() the new value. This can be used
	// to ensure Watchers always receive a full log of all Set() calls.
	//
	// This must not be changed after the first .Set/.Watch call.
	//
	// This is an experimental API and subject to change. It might be migrated
	// to per-Watcher settings defined within the main event.Value/Watcher
	// interfaces.
	Sync bool
}

// Set updates the Value to the given data. It is safe to call this from
// multiple goroutines, including concurrently.
//
// For more information about guarantees, see event.Value.Set.
func (m *Value) Set(val interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Update the data that is provided on first Get() to watchers.
	m.inner = val
	m.innerSet = true

	// Go through all watchers, updating them on the new value and filtering out
	// all closed watchers.
	newWatchers := make([]*watcher, 0, len(m.watchers))
	for _, w := range m.watchers {
		w := w
		if w.closed() {
			continue
		}
		w.update(m.Sync, val)
		newWatchers = append(newWatchers, w)
	}
	m.watchers = newWatchers
}

// watcher implements the event.Watcher interface for watchers returned by
// Value.
type watcher struct {
	// activeReqC is a channel used to request an active submission channel
	// from a pending Get function, if any.
	activeReqC chan chan interface{}
	// deadletterSubmitC is a channel used to communicate a value that
	// attempted to be submitted via activeReqC. This will be received by the
	// deadletter worker of this watcher and passed on to the next .Get call
	// that occurs.
	deadletterSubmitC chan interface{}

	// getSem is a channel-based semaphore (which is of size 1, and thus in
	// fact a mutex) that is used to ensure that only a single .Get() call is
	// active. It is implemented as a channel to permit concurrent .Get() calls
	// to error out instead of blocking.
	getSem chan struct{}
	// close is a channel that is closed when this watcher is itself Closed.
	close chan struct{}
}

// Watch retrieves a Watcher that keeps track on the version of the data
// contained within the Value that was last seen by a consumer.
//
// For more information about guarantees, see event.Value.Watch.
func (m *Value) Watch() event.Watcher {
	waiter := &watcher{
		activeReqC:        make(chan chan interface{}),
		deadletterSubmitC: make(chan interface{}),
		close:             make(chan struct{}),
		getSem:            make(chan struct{}, 1),
	}
	// Start the deadletter worker as a goroutine. It will be stopped when the
	// watcher is Closed() (as signaled by the close channel).
	go waiter.deadletterWorker()

	// Append this watcher to the Value.
	m.mu.Lock()
	m.watchers = append(m.watchers, waiter)
	// If the Value already has some value set, communicate that to the
	// first Get call by going through the deadletter worker.
	if m.innerSet {
		waiter.deadletterSubmitC <- m.inner
	}
	m.mu.Unlock()

	return waiter
}

// deadletterWorker runs the 'deadletter worker', as goroutine that contains
// any data that has been Set on the Value that is being watched that was
// unable to be delivered directly to a pending .Get call.
//
// It watches the deadletterSubmitC channel for updated data, and overrides
// previously received data. Then, when a .Get() begins to pend (and respond to
// activeReqC receives), the deadletter worker will deliver that value.
func (m *watcher) deadletterWorker() {
	// Current value, and flag to mark it as set (vs. nil).
	var cur interface{}
	var set bool

	for {
		if !set {
			// If no value is yet available, only attempt to receive one from the
			// submit channel, as there's nothing to submit to pending .Get() calls
			// yet.
			val, ok := <-m.deadletterSubmitC
			if !ok {
				// If the channel has been closed (by Close()), exit.
				return
			}
			cur = val
			set = true
		} else {
			// If a value is available, update the inner state. Otherwise, if a
			// .Get() is pending, submit our current state and unset it.
			select {
			case val, ok := <-m.deadletterSubmitC:
				if !ok {
					// If the channel has been closed (by Close()), exit.
					return
				}
				cur = val
			case c := <-m.activeReqC:
				// Potential race: a .Get() might've been active, but might've
				// quit by the time we're here (and will not receive on the
				// responded channel). Handle this gracefully by just returning
				// to the main loop if that's the case.
				select {
				case c <- cur:
					set = false
				default:
				}
			}
		}
	}
}

// closed returns whether this watcher has been closed.
func (m *watcher) closed() bool {
	select {
	case _, ok := <-m.close:
		if !ok {
			return true
		}
	default:
	}
	return false
}

// update is the high level update-this-watcher function called by Value.
func (m *watcher) update(sync bool, val interface{}) {
	// If synchronous delivery was requested, block until a watcher .Gets it.
	if sync {
		c := <-m.activeReqC
		c <- val
		return
	}

	// Otherwise, deliver asynchronously. This means either delivering directly
	// to a pending .Get if one exists, or submitting to the deadletter worker
	// otherwise.
	select {
	case c := <-m.activeReqC:
		// Potential race: a .Get() might've been active, but might've  quit by
		// the time we're here (and will not receive on the responded channel).
		// Handle this gracefully by falling back to the deadletter worker.
		select {
		case c <- val:
		default:
			m.deadletterSubmitC <- val
		}
	default:
		m.deadletterSubmitC <- val
	}
}

func (m *watcher) Close() error {
	close(m.deadletterSubmitC)
	close(m.close)
	return nil
}

// Get blocks until a Value's data is available. See event.Watcher.Get for
// guarantees and more information.
func (m *watcher) Get(ctx context.Context) (interface{}, error) {
	// Make sure we're the only active .Get call.
	select {
	case m.getSem <- struct{}{}:
	default:
		return nil, fmt.Errorf("cannot Get() concurrently on a single waiter")
	}
	defer func() {
		<-m.getSem
	}()

	c := make(chan interface{})

	// Start responding on activeReqC. This signals to .update and to the
	// deadletter worker that we're ready to accept data updates.

	// There is a potential for a race condition here that hasn't been observed
	// in tests but might happen:
	//   1) Value.Watch returns a Watcher 'w'.
	//   2) w.Set(0) is called, no .Get() is pending, so 0 is submitted to the
	//      deadletter worker.
	//   3) w.Get() is called, and activeReqC begins to be served.
	//   4) Simultaneously:
	//     a) w.Set(1) is called, attempting to submit via activeReqC
	//     b) the deadletter worker attempts to submit via activeReqC
	//
	// This could theoretically cause .Get() to first return 1, and then 0, if
	// the Set activeReqC read and subsequent channel write is served before
	// the deadletter workers' read/write is.
	// As noted, however, this has not been observed in practice, even though
	// TestConcurrency explicitly attempts to trigger this condition. More
	// research needs to be done to attempt to trigger this (or to lawyer the
	// Go channel spec to see if this has some guarantees that resolve this
	// either way), or a preemptive fix can be attempted by adding monotonic
	// counters associated with each .Set() value, ensuring an older value does
	// not replace a newer value.
	//
	// TODO(q3k): investigate this.
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case m.activeReqC <- c:
		case val := <-c:
			return val, nil
		}
	}
}
