// Package metrics implements a Prometheus metrics submission interface for BMDB
// client components. A Metrics object can be attached to a BMDB object, which
// will make all BMDB sessions/transactions/work statistics be submitted to that
// Metrics object.
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
)

// Processor describes some cloud component and possibly sub-component which acts
// upon the BMDB. When starting a BMDB session, this Processor can be provided to
// contextualize the metrics emitted by this session. Because the selected
// Processor ends up directly as a Prometheus metric label, it must be
// low-cardinality - thus all possible values are defined as an enum here. If a
// Session is not configured with a Processor, the default (ProcessorUnknown)
// will be used.
type Processor string

const (
	ProcessorUnknown             Processor = ""
	ProcessorShepherdInitializer Processor = "shepherd-initializer"
	ProcessorShepherdProvisioner Processor = "shepherd-provisioner"
	ProcessorShepherdRecoverer   Processor = "shepherd-recoverer"
	ProcessorShepherdUpdater     Processor = "shepherd-updater"
	ProcessorBMSRV               Processor = "bmsrv"
	ProcessorScruffyStats        Processor = "scruffy-stats"
)

// String returns the Prometheus label value for use with the 'processor' label
// key.
func (p Processor) String() string {
	switch p {
	case ProcessorUnknown:
		return "unknown"
	default:
		return string(p)
	}
}

// MetricsSet contains all the Prometheus metrics objects related to a BMDB
// client.
//
// The MetricsSet object is goroutine-safe.
//
// An empty MetricsSet object is not valid, and should be instead constructed
// using New.
//
// A nil MetricsSet object is valid and represents a no-op metrics recorder
// that's never collected.
type MetricsSet struct {
	sessionStarted      *prometheus.CounterVec
	transactionExecuted *prometheus.CounterVec
	transactionRetried  *prometheus.CounterVec
	transactionFailed   *prometheus.CounterVec
	workStarted         *prometheus.CounterVec
	workFinished        *prometheus.CounterVec
}

func processorCounter(name, help string, labels ...string) *prometheus.CounterVec {
	labels = append([]string{"processor"}, labels...)
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
		labels,
	)
}

// New creates a new BMDB MetricsSet object which can be then attached to a BMDB
// object by calling BMDB.EnableMetrics on the MetricsSet object.
//
// The given registry must be a valid Prometheus registry, and all metrics
// contained in this MetricsSet object will be registered into it.
//
// The MetricsSet object can be shared between multiple BMDB object.
//
// The MetricsSet object is goroutine-safe.
func New(registry *prometheus.Registry) *MetricsSet {
	m := &MetricsSet{
		sessionStarted:      processorCounter("bmdb_session_started", "How many sessions this worker started"),
		transactionExecuted: processorCounter("bmdb_transaction_executed", "How many transactions were performed by this worker"),
		transactionRetried:  processorCounter("bmdb_transaction_retried", "How many transaction retries were performed by this worker"),
		transactionFailed:   processorCounter("bmdb_transaction_failed", "How many transactions failed permanently on this worker"),
		workStarted:         processorCounter("bmdb_work_started", "How many work items were performed by this worker, partitioned by process", "process"),
		workFinished:        processorCounter("bmdb_work_finished", "How many work items were finished by this worker, partitioned by process and result", "process", "result"),
	}
	registry.MustRegister(
		m.sessionStarted,
		m.transactionExecuted,
		m.transactionRetried,
		m.transactionFailed,
		m.workStarted,
		m.workFinished,
	)
	return m
}

// ProcessorRecorder wraps a MetricsSet object with the context of some
// Processor. It exposes methods that record specific events into the managed
// Metrics.
//
// The ProcessorRecorder object is goroutine safe.
//
// An empty ProcessorRecorder object is not valid, and should be instead
// constructed using Metrics.Recorder.
//
// A nil ProcessorRecorder object is valid and represents a no-op metrics
// recorder.
type ProcessorRecorder struct {
	m      *MetricsSet
	labels prometheus.Labels
}

// Recorder builds a ProcessorRecorder for the given Metrics and a given
// Processor.
func (m *MetricsSet) Recorder(p Processor) *ProcessorRecorder {
	if m == nil {
		return nil
	}
	return &ProcessorRecorder{
		m: m,
		labels: prometheus.Labels{
			"processor": p.String(),
		},
	}
}

// OnTransactionStarted should be called any time a BMDB client starts or
// re-starts a BMDB Transaction. The attempt should either be '1' (for the first
// attempt) or a number larger than 1 for any subsequent attempt (i.e. retry) of
// a transaction.
func (r *ProcessorRecorder) OnTransactionStarted(attempt int64) {
	if r == nil {
		return
	}
	if attempt == 1 {
		r.m.transactionExecuted.With(r.labels).Inc()
	} else {
		r.m.transactionRetried.With(r.labels).Inc()
	}
}

// OnTransactionFailed should be called any time a BMDB client fails a
// BMDB Transaction permanently.
func (r *ProcessorRecorder) OnTransactionFailed() {
	if r == nil {
		return
	}
	r.m.transactionFailed.With(r.labels).Inc()
}

// OnSessionStarted should be called any time a BMDB client opens a new BMDB
// Session.
func (r *ProcessorRecorder) OnSessionStarted() {
	if r == nil {
		return
	}
	r.m.sessionStarted.With(r.labels).Inc()
}

// ProcessRecorder wraps a ProcessorRecorder with an additional model.Process.
// The resulting object can then record work-specific events.
//
// The PusherWithProcess object is goroutine-safe.
type ProcessRecorder struct {
	*ProcessorRecorder
	labels prometheus.Labels
}

// WithProcess wraps a given Pusher with a Process.
//
// The resulting PusherWithProcess object is goroutine-safe.
func (r *ProcessorRecorder) WithProcess(process model.Process) *ProcessRecorder {
	if r == nil {
		return nil
	}
	return &ProcessRecorder{
		ProcessorRecorder: r,
		labels: prometheus.Labels{
			"processor": r.labels["processor"],
			"process":   string(process),
		},
	}
}

// OnWorkStarted should be called any time a BMDB client starts a new Work item.
func (r *ProcessRecorder) OnWorkStarted() {
	if r == nil {
		return
	}
	r.m.workStarted.With(r.labels).Inc()
}

type WorkResult string

const (
	WorkResultFinished WorkResult = "finished"
	WorkResultCanceled WorkResult = "canceled"
	WorkResultFailed   WorkResult = "failed"
)

// OnWorkFinished should be called any time a BMDB client finishes, cancels or
// fails a Work item.
func (r *ProcessRecorder) OnWorkFinished(result WorkResult) {
	if r == nil {
		return
	}
	r.m.workFinished.MustCurryWith(r.labels).With(prometheus.Labels{"result": string(result)}).Inc()
}
