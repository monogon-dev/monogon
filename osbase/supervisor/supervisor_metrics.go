// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"sync"
	"time"
)

// Metrics is an interface from the supervisor to any kind of metrics-collecting
// component.
type Metrics interface {
	// NotifyNodeState is called whenever a given runnable at a given DN changes
	// state. Called synchronously from the supervisor's processor loop, so must not
	// block, but is also guaranteed to only be called from a single goroutine.
	NotifyNodeState(dn string, state NodeState)
}

// metricsFanout is used internally to fan out a single Metrics interface (which
// it implements) onto multiple subordinate Metrics interfaces (as provided by
// the user via WithMetrics).
type metricsFanout struct {
	sub []Metrics
}

func (m *metricsFanout) NotifyNodeState(dn string, state NodeState) {
	for _, sub := range m.sub {
		sub.NotifyNodeState(dn, state)
	}
}

// InMemoryMetrics is a simple Metrics implementation that keeps an in-memory
// mirror of the state of all DNs in the supervisor. The zero value for
// InMemoryMetrics is ready to use.
type InMemoryMetrics struct {
	mu  sync.RWMutex
	dns map[string]DNState
}

// DNState is the state of a supervisor runnable, recorded alongside a timestamp
// of when the State changed.
type DNState struct {
	// State is the current state of the runnable.
	State NodeState
	// Transition is the time at which the runnable reached its State.
	Transition time.Time
}

func (m *InMemoryMetrics) NotifyNodeState(dn string, state NodeState) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.dns == nil {
		m.dns = make(map[string]DNState)
	}
	m.dns[dn] = DNState{
		State:      state,
		Transition: time.Now(),
	}
}

// DNs returns a copy (snapshot in time) of the recorded DN states, in a map from
// DN to DNState. The returned value can be mutated.
func (m *InMemoryMetrics) DNs() map[string]DNState {
	m.mu.RLock()
	defer m.mu.RUnlock()

	res := make(map[string]DNState)
	for k, v := range m.dns {
		res[k] = v
	}
	return res
}
