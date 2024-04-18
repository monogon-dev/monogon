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

	"source.monogon.dev/metropolis/pkg/event"
)

var (
	// Type assert that *Value implements Value. We do this artificially, as
	// there currently is no code path that needs this to be strictly true. However,
	// users of this library might want to rely on the Value type instead of
	// particular Value implementations.
	_ event.Value[int] = &Value[int]{}
)

// Value is a 'memory value', which implements a event.Value stored in memory.
// It is safe to construct an empty object of this type. However, this must not
// be copied.
type Value[T any] struct {
	// mu guards the inner, innerSet and watchers fields.
	mu sync.RWMutex
	// inner is the latest data Set on the Value. It is used to provide the
	// newest version of the Set data to new watchers.
	inner T
	// innerSet is true when inner has been Set at least once. It is used to
	// differentiate between a nil and unset value.
	innerSet bool
	// watchers is the list of watchers that should be updated when new data is
	// Set. It will grow on every .Watch() and shrink any time a watcher is
	// determined to have been closed.
	watchers []*watcher[T]

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
func (m *Value[T]) Set(val T) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Update the data that is provided on first Get() to watchers.
	m.inner = val
	m.innerSet = true

	// Go through all watchers, updating them on the new value and filtering out
	// all closed watchers.
	newWatchers := m.watchers[:0]
	for _, w := range m.watchers {
		if w.closed() {
			continue
		}
		w.update(m.Sync, val)
		newWatchers = append(newWatchers, w)
	}
	if cap(newWatchers) > len(newWatchers)*3 {
		reallocated := make([]*watcher[T], 0, len(newWatchers)*2)
		newWatchers = append(reallocated, newWatchers...)
	}
	m.watchers = newWatchers
}

// watcher implements the event.Watcher interface for watchers returned by
// Value.
type watcher[T any] struct {
	// bufferedC is a buffered channel of size 1 for submitting values to the
	// watcher.
	bufferedC chan T
	// unbufferedC is an unbuffered channel, which is used when Sync is enabled.
	unbufferedC chan T

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
func (m *Value[T]) Watch() event.Watcher[T] {
	waiter := &watcher[T]{
		bufferedC:   make(chan T, 1),
		unbufferedC: make(chan T),
		close:       make(chan struct{}),
		getSem:      make(chan struct{}, 1),
	}

	m.mu.Lock()
	// If the watchers slice is at capacity, drop closed watchers, and
	// reallocate the slice at 2x length if it is not between 1.5x and 3x.
	if len(m.watchers) == cap(m.watchers) {
		newWatchers := m.watchers[:0]
		for _, w := range m.watchers {
			if !w.closed() {
				newWatchers = append(newWatchers, w)
			}
		}
		if cap(newWatchers)*2 < len(newWatchers)*3 || cap(newWatchers) > len(newWatchers)*3 {
			reallocated := make([]*watcher[T], 0, len(newWatchers)*2)
			newWatchers = append(reallocated, newWatchers...)
		}
		m.watchers = newWatchers
	}
	// Append this watcher to the Value.
	m.watchers = append(m.watchers, waiter)
	// If the Value already has some value set, put it in the buffered channel.
	if m.innerSet {
		waiter.bufferedC <- m.inner
	}
	m.mu.Unlock()

	return waiter
}

// closed returns whether this watcher has been closed.
func (m *watcher[T]) closed() bool {
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
func (m *watcher[T]) update(sync bool, val T) {
	// If synchronous delivery was requested, block until a watcher .Gets it,
	// or is closed.
	if sync {
		select {
		case m.unbufferedC <- val:
		case <-m.close:
		}
		return
	}

	// Otherwise, deliver asynchronously. If there is already a value in the
	// buffered channel that was not retrieved, drop it.
	select {
	case <-m.bufferedC:
	default:
	}
	// The channel is now empty, so sending to it cannot block.
	m.bufferedC <- val
}

func (m *watcher[T]) Close() error {
	close(m.close)
	return nil
}

// Get blocks until a Value's data is available. See event.Watcher.Get for
// guarantees and more information.
func (m *watcher[T]) Get(ctx context.Context, opts ...event.GetOption[T]) (T, error) {
	// Make sure we're the only active .Get call.
	var empty T
	select {
	case m.getSem <- struct{}{}:
	default:
		return empty, fmt.Errorf("cannot Get() concurrently on a single waiter")
	}
	defer func() {
		<-m.getSem
	}()

	var predicate func(t T) bool
	for _, opt := range opts {
		if opt.Predicate != nil {
			predicate = opt.Predicate
		}
		if opt.BacklogOnly {
			return empty, errors.New("BacklogOnly is not implemented for memory watchers")
		}
	}

	for {
		var val T
		// For Sync values, ensure the initial value in the buffered
		// channel is delivered first.
		select {
		case val = <-m.bufferedC:
		default:
			select {
			case <-ctx.Done():
				return empty, ctx.Err()
			case val = <-m.bufferedC:
			case val = <-m.unbufferedC:
			}
		}
		if predicate != nil && !predicate(val) {
			continue
		}
		return val, nil
	}
}
