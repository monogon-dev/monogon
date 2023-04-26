package scruffy

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
)

func TestBMDBStats(t *testing.T) {
	s := Server{
		Config: Config{
			BMDB: bmdb.BMDB{
				Config: bmdb.Config{
					Database: component.CockroachConfig{
						InMemory: true,
					},
				},
			},
		},
	}

	registry := prometheus.NewRegistry()
	runner := newBMDBStatsRunner(&s, registry)

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	expect := func(wantValues map[string]int64) {
		t.Helper()
		res, err := registry.Gather()
		if err != nil {
			t.Fatalf("Gather: %v", err)
		}
		gotValues := make(map[string]bool)
		for _, mf := range res {
			if len(mf.Metric) != 1 {
				for _, m := range mf.Metric {
					var lvs []string
					for _, lp := range m.Label {
						lvs = append(lvs, fmt.Sprintf("%s=%s", *lp.Name, *lp.Value))
					}
					sort.Strings(lvs)
					name := fmt.Sprintf("%s[%s]", *mf.Name, strings.Join(lvs, ","))
					gotValues[name] = true
					if _, ok := wantValues[name]; !ok {
						t.Errorf("MetricFamily %s: unexpected", name)
					}
					if want, got := wantValues[name], int64(*m.Gauge.Value); want != got {
						t.Errorf("MetricFamily %s: wanted %d, got %d", *mf.Name, want, got)
					}
				}
			} else {
				m := mf.Metric[0]
				gotValues[*mf.Name] = true
				if want, got := wantValues[*mf.Name], int64(*m.Gauge.Value); want != got {
					t.Errorf("MetricFamily %s: wanted %d, got %d", *mf.Name, want, got)
				}
				if _, ok := wantValues[*mf.Name]; !ok {
					t.Errorf("MetricFamily %s: unexpected", *mf.Name)
				}
			}
		}
		for mf, _ := range wantValues {
			if !gotValues[mf] {
				t.Errorf("MetricFamily %s: missing", mf)
			}
		}
	}

	expect(map[string]int64{
		"bmdb_stats_machines":                                    0,
		"bmdb_stats_machines_provided":                           0,
		"bmdb_stats_machines_heartbeating":                       0,
		"bmdb_stats_machines_pending_installation":               0,
		"bmdb_stats_machines_installed":                          0,
		"bmdb_stats_machines_pending_agent_start":                0,
		"bmdb_stats_machines_pending_agent_recovery":             0,
		"bmdb_stats_active_backoffs[process=ShepherdAccess]":     0,
		"bmdb_stats_active_backoffs[process=ShepherdAgentStart]": 0,
		"bmdb_stats_active_backoffs[process=ShepherdRecovery]":   0,
		"bmdb_stats_active_work[process=ShepherdAccess]":         0,
		"bmdb_stats_active_work[process=ShepherdAgentStart]":     0,
		"bmdb_stats_active_work[process=ShepherdRecovery]":       0,
	})

	conn, err := s.Config.BMDB.Open(true)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("StartSession: %v", err)
	}

	s.bmdb = conn
	s.sessionC = make(chan *bmdb.Session)
	go s.sessionWorker(ctx)
	if err := runner.runOnce(ctx); err != nil {
		t.Fatal(err)
	}

	expect(map[string]int64{
		"bmdb_stats_machines":                                    0,
		"bmdb_stats_machines_provided":                           0,
		"bmdb_stats_machines_heartbeating":                       0,
		"bmdb_stats_machines_pending_installation":               0,
		"bmdb_stats_machines_installed":                          0,
		"bmdb_stats_machines_pending_agent_start":                0,
		"bmdb_stats_machines_pending_agent_recovery":             0,
		"bmdb_stats_active_backoffs[process=ShepherdAccess]":     0,
		"bmdb_stats_active_backoffs[process=ShepherdAgentStart]": 0,
		"bmdb_stats_active_backoffs[process=ShepherdRecovery]":   0,
		"bmdb_stats_active_work[process=ShepherdAccess]":         0,
		"bmdb_stats_active_work[process=ShepherdAgentStart]":     0,
		"bmdb_stats_active_work[process=ShepherdRecovery]":       0,
	})

	f := fill().
		// Provided, needs installation.
		machine().providedE("1").build().
		// Three machines needing recovery.
		machine().providedE("2").agentNeverHeartbeat().build().
		machine().providedE("3").agentNeverHeartbeat().build().
		machine().providedE("4").agentNeverHeartbeat().build().
		// One machine correctly heartbeating.
		machine().providedE("5").agentHealthy().build().
		// Two machines heartbeating and pending installation.
		machine().providedE("6").agentHealthy().installRequested(10).build().
		machine().providedE("7").agentHealthy().installRequested(10).installReported(9).build().
		// Machine which is pending installation _and_ recovery.
		machine().providedE("8").agentNeverHeartbeat().installRequested(10).build().
		// Machine which has been successfully installed.
		machine().providedE("9").agentStoppedHeartbeating().installRequested(10).installReported(10).build()

	err = sess.Transact(ctx, func(q *model.Queries) error {
		return f(ctx, q)
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := runner.runOnce(ctx); err != nil {
		t.Fatal(err)
	}

	expect(map[string]int64{
		"bmdb_stats_machines":                                    9,
		"bmdb_stats_machines_provided":                           9,
		"bmdb_stats_machines_heartbeating":                       3,
		"bmdb_stats_machines_pending_installation":               3,
		"bmdb_stats_machines_installed":                          1,
		"bmdb_stats_machines_pending_agent_start":                1,
		"bmdb_stats_machines_pending_agent_recovery":             4,
		"bmdb_stats_active_backoffs[process=ShepherdAccess]":     0,
		"bmdb_stats_active_backoffs[process=ShepherdAgentStart]": 0,
		"bmdb_stats_active_backoffs[process=ShepherdRecovery]":   0,
		"bmdb_stats_active_work[process=ShepherdAccess]":         0,
		"bmdb_stats_active_work[process=ShepherdAgentStart]":     0,
		"bmdb_stats_active_work[process=ShepherdRecovery]":       0,
	})
}
