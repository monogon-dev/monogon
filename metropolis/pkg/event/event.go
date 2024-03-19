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

// Package event defines and implements Event Values, a mechanism in which
// multiple consumers can watch a value for updates in a reliable way.
//
// Values currently are kept in memory (see: MemoryValue), but a future
// implementation might exist for other storage backends, eg. etcd.
//
// # Background and intended use
//
// The Event Value library is intended to be used within Metropolis'
// supervisor-based runnables to communicate state changes to other runnables,
// while permitting both sides to restart if needed. It grew out of multiple
// codebases reimplementing an ad-hoc observer pattern, and from the
// realization that implementing all possible edge cases of such patterns is
// non-trivial and subject to programming errors. As such, it was turned into a
// self-standing library.
//
// Why not just channels?
//
// Plain channels have multiple deficiencies for this usecase:
//   - Strict FIFO behaviour: all values sent to a channel must be received, and
//     historic and newest data must be treated in the same way. This means that
//     a consumer of state changes must process all updates to the value as if
//     they are the newest, and unable to skip rapid updates when a system is
//     slowly settling due to a cascading state change.
//   - Implementation overhead: implementing an observer
//     registration/unregistration pattern is prone to programming bugs,
//     especially for features like always first sending the current state to a
//     new observer.
//   - Strict buffer size: due to their FIFO nature and the possibility of
//     consumers not receiving actively, channels would have to buffer all
//     existing updates, requiring some arbitrary best-guess channel buffer
//     sizing that would still not prevent blocking writes or data loss in a
//     worst case scenario.
//
// Or, in other words: Go channels are a synchronization primitive, not a
// ready-made solution to this problem. The Event Value implementation in fact
// extensively uses Go channels within its implementation as a building block.
//
// Why not just condition variables (sync.Cond)?
//
// Go's condition variable implementation doesn't fully address our needs
// either:
//   - No context/canceling support: once a condition is being Wait()ed on,
//     this cannot be interrupted. This is especially painful and unwieldy when
//     dealing with context-heavy code, such as Metropolis.
//   - Spartan API: expecting users to plainly use sync.Cond is risky, as the API
//     is fairly low-level.
//   - No solution for late consumers: late consumers (ones that missed the value
//     being set by a producer) would still have to implement logic in order to
//     find out such a value, as sync.Cond only supports what amounts to
//     edge-level triggers as part of its Broadcast/Signal system.
//
// It would be possible to implement MemoryValue using a sync.Cond internally,
// but such an implementation would likely be more complex than the current
// implementation based on channels and mutexes, as it would have to work
// around issues like lack of canceling, etc.
package event

import (
	"context"
	"errors"

	"source.monogon.dev/metropolis/pkg/supervisor"
)

// A Value is an 'Event Value', some piece of data that can be updated ('Set')
// by Producers and retrieved by Consumers.
type Value[T any] interface {
	// Set updates the Value to the given data. It is safe to call this from
	// multiple goroutines, including concurrently.
	//
	// Any time Set is called, any consumers performing a Watch on this Value
	// will be notified with the new data - even if the Set data is the same as
	// the one that was already stored.
	//
	// A Value will initially have no data set. This 'no data' state is seen by
	// consumers by the first .Get() call on the Watcher blocking until data is Set.
	//
	// All updates will be serialized in an arbitrary order - if multiple
	// producers wish to perform concurrent actions to update the Value partially,
	// this should be negotiated and serialized externally by the producers.
	Set(val T)

	// ValueWatch implements the Watch method. It is split out into another
	// interface to allow some 'Event Values' to implement only the watch/read
	// part, with the write side being implicit or defined by a more complex
	// interface than a simple Set().
	ValueWatch[T]
}

// ValueWatch is the read side of an 'Event Value', witch can by retrieved by
// Consumers by performing a Watch operation on it.
type ValueWatch[T any] interface {
	// Watch retrieves a Watcher that keeps track on the version of the data
	// contained within the Value that was last seen by a consumer. Once a
	// Watcher is retrieved, it can be used to then get the actual data stored
	// within the Value, and to reliably retrieve updates to it without having
	// to poll for changes.
	Watch() Watcher[T]
}

// A Watcher keeps track of the last version of data seen by a consumer for a
// given Value. Each consumer should use an own Watcher instance, and it is not
// safe to use this type concurrently. However, it is safe to move/copy it
// across different goroutines, as long as no two goroutines access it
// simultaneously.
type Watcher[T any] interface {
	// Get blocks until a Value's data is available:
	//  - On first use of a Watcher, Get will return the data contained in the
	//    value at the time of calling .Watch(), or block if no data has been
	//    .Set() on it yet. If a value has been Set() since the initial
	//    creation of the Watch() but before Get() is called for the first
	//    time, the first Get() call will immediately return the new value.
	//  - On subsequent uses of a Watcher, Get will block until the given Value
	//    has been Set with new data. This does not necessarily mean that the
	//    new data is different - consumers should always perform their own
	//    checks on whether the update is relevant to them (ie., the data has
	//    changed in a significant way), unless specified otherwise by a Value
	//    publisher.
	//
	// Get() will always return the current newest data that has been Set() on
	// the Value, and not a full log of historical events. This is geared
	// towards event values where consumers only care about changes to data
	// since last retrieval, not every value that has been Set along the way.
	// Thus, consumers need not make sure that they actively .Get() on a
	// watcher all the times.
	//
	// If the context is canceled before data is available to be returned, the
	// context's error will be returned. However, the Watcher will still need to be
	// Closed, as it is still fully functional after the context has been canceled.
	//
	// Concurrent requests to Get result in an error. The reasoning to return
	// an error instead of attempting to serialize the requests is that any
	// concurrent access from multiple goroutines would cause a desync in the
	// next usage of the Watcher. For example:
	//   1) w.Get() (in G0) and w.Get(G1) start. They both block waiting for an
	//      initial value.
	//   2) v.Set(0)
	//   3) w.Get() in G0 returns 0,
	//   4) v.Set(1)
	//   4) w.Get() in G1 returns 1,
	// This would cause G0 and G1 to become desynchronized between eachother
	// (both have different value data) and subsequent updates will also
	// continue skipping some updates.
	// If multiple goroutines need to access the Value, they should each use
	// their own Watcher.
	Get(context.Context, ...GetOption[T]) (T, error)

	// Close must be called if the Watcher is not going to be used anymore -
	// otherwise, it will not be garbage collected.
	Close() error
}

type GetOption[T any] struct {
	Predicate   func(t T) bool
	BacklogOnly bool
}

func Filter[T any](pred func(T) bool) GetOption[T] {
	return GetOption[T]{
		Predicate: pred,
	}
}

// BacklogOnly will prevent Get from blocking on waiting for more updates from
// etcd, by instead returning BacklogDone whenever no more data is currently
// locally available. This is different however, from establishing that there
// are no more pending updates from the etcd cluster - the only way to ensure
// the local client is up to date is by performing Get calls without this option
// set.
//
// This mode of retrieval should only be used for the retrieval of the existing
// data in the etcd cluster on the initial creation of the Watcher (by
// repeatedly calling Get until BacklogDone is returned), and shouldn't be set
// for any subsequent call. Any use of this option after that initial fetch is
// undefined behaviour that exposes the internals of the Get implementation, and
// must not be relied on. However, in the future, this behaviour might be
// formalized.
//
// This mode is particularly useful for ranged watchers. Non-ranged watchers can
// still use this option to distinguish between blocking because of the
// nonexistence of an object vs. blocking because of networking issues. However,
// non-ranged retrieval semantics generally will rarely need to make this
// distinction.
func BacklogOnly[T any]() GetOption[T] {
	return GetOption[T]{BacklogOnly: true}
}

var (
	// BacklogDone is returned by Get when BacklogOnly is set and there is no more
	// event data stored in the Watcher client, ie. when the initial cluster state
	// of the requested key has been retrieved.
	BacklogDone = errors.New("no more backlogged data")
)

// Pipe a Value's initial state and subsequent updates to an already existing
// channel in a supervisor.Runnable. This is mostly useful when wanting to select
// {} on many Values.
//
// The given channel will NOT be closed when the runnable exits. The process
// receiving from the channel should be running in a group with the pipe
// runnable, so that both restart if either does. This ensures that there is always
// at least one value in the channel when the receiver starts.
func Pipe[T any](value Value[T], c chan<- T, opts ...GetOption[T]) supervisor.Runnable {
	return func(ctx context.Context) error {
		supervisor.Signal(ctx, supervisor.SignalHealthy)
		w := value.Watch()
		defer w.Close()
		for {
			v, err := w.Get(ctx, opts...)
			if err != nil {
				return err
			}
			c <- v
		}
	}
}
