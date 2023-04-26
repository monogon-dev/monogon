package scruffy

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
)

// bmdbStatsRunner collects metrics from the BMDB and exposes them as Prometheus
// metrics via a registry passed to newBMDBStatsRunner.
type bmdbStatsRunner struct {
	s          *Server
	collectors []*statsCollector
}

// A statsCollectorDefinition describes how to gather a given metric via a BMDB
// SQL query.
type statsCollectorDefinition struct {
	// name of the metric. Used in actual metric name, prefixed with 'bmdb_stats_'.
	name string
	// help string emitted in prometheus endpoint.
	help string
	// labels is the label 'type definition', containing information about the
	// dimensions of this metric.
	labels labelDefinitions
	// query used to retrieve the metric data.
	query func(*model.Queries, context.Context) ([]model.MetricValue, error)
}

// labelProcess is the type definition of the 'process' label 'type', which is a
// fixed-cardinality representation of the database Process enum.
var labelProcess = labelDefinition{
	name: "process",
	initialValues: []string{
		string(model.ProcessShepherdAccess),
		string(model.ProcessShepherdAgentStart),
		string(model.ProcessShepherdRecovery),
	},
}

var collectorDefs = []statsCollectorDefinition{
	{
		name:   "active_backoffs",
		help:   "Number of active backoffs, partitioned by process. There may be more than one active backoff per machine.",
		query:  model.WrapLabeledMetric((*model.Queries).CountActiveBackoffs),
		labels: []labelDefinition{labelProcess},
	},
	{
		name:   "active_work",
		help:   "Number of active work, partitioned by process. There may be more than one active work item per machine.",
		query:  model.WrapLabeledMetric((*model.Queries).CountActiveWork),
		labels: []labelDefinition{labelProcess},
	},
	{
		name:  "machines",
		help:  "Number of machines in the BMDB.",
		query: model.WrapSimpleMetric((*model.Queries).CountMachines),
	},
	{
		name:  "machines_provided",
		help:  "Number of provided machines in the BMDB.",
		query: model.WrapSimpleMetric((*model.Queries).CountMachinesProvided),
	},
	{
		name:  "machines_heartbeating",
		help:  "Number of machines with a currently heartbeating agent.",
		query: model.WrapSimpleMetric((*model.Queries).CountMachinesAgentHeartbeating),
	},
	{
		name:  "machines_pending_installation",
		help:  "Number of machines pending installation.",
		query: model.WrapSimpleMetric((*model.Queries).CountMachinesInstallationPending),
	},
	{
		name:  "machines_installed",
		help:  "Number of machines succesfully installed.",
		query: model.WrapSimpleMetric((*model.Queries).CountMachinesInstallationComplete),
	},
	{
		name:  "machines_pending_agent_start",
		help:  "Number of machines pending the agent start workflow.",
		query: model.WrapSimpleMetric((*model.Queries).CountMachinesForAgentStart),
	},
	{
		name:  "machines_pending_agent_recovery",
		help:  "Number of machines pending the agent recovery workflow.",
		query: model.WrapSimpleMetric((*model.Queries).CountMachinesForAgentRecovery),
	},
}

// A statsCollector is an instantiated statsCollectorDefinition which carries the
// actual prometheus gauge backing the metric.
type statsCollector struct {
	gauge *prometheus.GaugeVec
	def   *statsCollectorDefinition
}

// setDefaults emits gauges with zero values for all metrics of the runner, using
// the initialLabel data gathered from each metric definition.
func (b *bmdbStatsRunner) setDefaults() {
	for _, collector := range b.collectors {
		info := collector.def
		initial := info.labels.initialLabels()
		if len(initial) == 0 {
			collector.gauge.With(nil).Set(0.0)
		} else {
			for _, labels := range initial {
				collector.gauge.With(labels).Set(0.0)
			}
		}
	}
}

// newBMDBStatsRunner builds a bmdbStatsRunner from the collectorDefs above. The
// bmdbStatsRunner then has the given's Server BMDB connection bound to it and
// can perform actual database statistic gathering.
func newBMDBStatsRunner(s *Server, reg *prometheus.Registry) *bmdbStatsRunner {
	var collectors []*statsCollector

	for _, info := range collectorDefs {
		info := info
		gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "bmdb_stats_" + info.name,
			Help: info.help,
		}, info.labels.names())
		reg.MustRegister(gauge)

		collectors = append(collectors, &statsCollector{
			gauge: gauge,
			def:   &info,
		})
	}

	res := &bmdbStatsRunner{
		s:          s,
		collectors: collectors,
	}
	res.setDefaults()
	return res
}

func (b *bmdbStatsRunner) run(ctx context.Context) {
	klog.Infof("Starting stats runner...")

	ti := time.NewTicker(b.s.Config.StatsRunnerRate)

	for {
		err := b.runOnce(ctx)
		if err != nil {
			if errors.Is(err, ctx.Err()) {
				return
			}
			klog.Errorf("Stats run failed: %v", err)
		}
		select {
		case <-ti.C:
		case <-ctx.Done():
			klog.Infof("Exiting stats runner (%v)...", ctx.Err())
			return
		}
	}
}

func (b *bmdbStatsRunner) runOnce(ctx context.Context) error {
	sess, err := b.s.session(ctx)
	if err != nil {
		return err
	}

	results := make(map[string][]model.MetricValue)
	// TODO(q3k): don't fail entire run if we can't collect just one metric.
	err = sess.Transact(ctx, func(q *model.Queries) error {
		for _, c := range b.collectors {
			res, err := c.def.query(q, ctx)
			if err != nil {
				return fmt.Errorf("collecting %s failed: %v", c.def.name, err)
			} else {
				results[c.def.name] = res
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	b.setDefaults()
	for _, c := range b.collectors {
		for _, m := range results[c.def.name] {
			klog.Infof("Setting %s (%v) to %d", c.def.name, m.Labels, m.Count)
			c.gauge.With(m.Labels).Set(float64(m.Count))
		}
	}

	return nil
}
