// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// MetricsPrometheus is a Metrics implementation which exports the supervisor
// metrics over some prometheus registry.
//
// This structure must be constructed with NewMetricsPrometheus.
//
// The metrics exported are:
//   - monogon_supervisor_dn_state_total
//   - monogon_superfisor_dn_state_transition_count
type MetricsPrometheus struct {
	exportedState *prometheus.GaugeVec
	exportedEdge  *prometheus.CounterVec
	cachedState   map[string]*NodeState
}

// NewMetricsPrometheus initializes Supervisor metrics in a prometheus registry
// and return a Metrics instance to be used with WithMetrics.
//
// This should only be called once for a given registry.
func NewMetricsPrometheus(registry *prometheus.Registry) *MetricsPrometheus {
	factory := promauto.With(registry)
	res := &MetricsPrometheus{
		exportedState: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "monogon",
			Subsystem: "supervisor",
			Name:      "dn_state_total",
			Help:      "Total count of supervisor runnables, broken up by DN and state",
		}, []string{"dn", "state"}),
		exportedEdge: factory.NewCounterVec(prometheus.CounterOpts{
			Namespace:   "monogon",
			Subsystem:   "supervisor",
			Name:        "dn_state_transition_count",
			Help:        "Total count of supervisor runnable state transitions, broken up by DN and (old_state, new_state) tuple",
			ConstLabels: nil,
		}, []string{"dn", "old_state", "new_state"}),
		cachedState: make(map[string]*NodeState),
	}
	return res
}

func (m *MetricsPrometheus) exportState(dn string, state NodeState, value float64) {
	m.exportedState.With(map[string]string{
		"state": state.String(),
		"dn":    dn,
	}).Set(value)
}

func (m *MetricsPrometheus) exportEdge(dn string, oldState, newState NodeState) {
	m.exportedEdge.With(map[string]string{
		"old_state": oldState.String(),
		"new_state": newState.String(),
		"dn":        dn,
	}).Inc()
}

func (m *MetricsPrometheus) NotifyNodeState(dn string, state NodeState) {
	// Set all other exported states to zero, so that a given DN is only in a single
	// state.
	for _, st := range NodeStates {
		if st == state {
			continue
		}
		m.exportState(dn, st, 0.0)
	}
	// Export new state.
	m.exportState(dn, state, 1.0)

	// Export edge transition (assume previous state was Dead if this is the first
	// time we see this DN).
	previous := NodeStateDead
	if m.cachedState[dn] != nil {
		previous = *m.cachedState[dn]
	}
	m.exportEdge(dn, previous, state)
	m.cachedState[dn] = &state
}
