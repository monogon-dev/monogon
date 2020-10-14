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

package logtree

import (
	"sync/atomic"
)

// subscriber is an observer for new entries that are appended to the journal.
type subscriber struct {
	// filters that entries need to pass through in order to be sent to the subscriber.
	filters []filter
	// dataC is the channel to which entries that pass filters will be sent. The channel must be drained regularly in
	// order to prevent accumulation of goroutines and possible reordering of messages.
	dataC chan *LogEntry
	// doneC is a channel that is closed once the subscriber wishes to stop receiving notifications.
	doneC chan struct{}
	// missed is the amount of messages missed by the subscriber by not receiving from dataC fast enough
	missed uint64
}

// subscribe attaches a subscriber to the journal.
// mu must be taken in W mode
func (j *journal) subscribe(sub *subscriber) {
	j.subscribers = append(j.subscribers, sub)
}

// notify sends an entry to all subscribers that wish to receive it.
func (j *journal) notify(e *entry) {
	j.mu.Lock()
	defer j.mu.Unlock()

	newSub := make([]*subscriber, 0, len(j.subscribers))
	for _, sub := range j.subscribers {
		select {
		case <-sub.doneC:
			close(sub.dataC)
			continue
		default:
			newSub = append(newSub, sub)
		}

		for _, filter := range sub.filters {
			if !filter(e.origin, e.leveled.severity) {
				continue
			}
		}
		select {
		case sub.dataC <- &LogEntry{LeveledPayload: e.leveled, DN: e.origin}:
		default:
			atomic.AddUint64(&sub.missed, 1)
		}
	}
	j.subscribers = newSub
}
