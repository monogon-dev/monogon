// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package scruffy

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"

	aapi "source.monogon.dev/cloud/agent/api"
	"source.monogon.dev/cloud/bmaas/server/api"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
)

type filler func(ctx context.Context, q *model.Queries) error

func fill() filler {
	return func(ctx context.Context, q *model.Queries) error {
		return nil
	}
}

func (f filler) chain(n func(ctx context.Context, q *model.Queries) error) filler {
	return func(ctx context.Context, q *model.Queries) error {
		if err := f(ctx, q); err != nil {
			return err
		}
		return n(ctx, q)
	}
}

type fillerMachine struct {
	f filler

	provider   *model.Provider
	providerID *string

	location *string

	threads *int32
	ramgb   *int64

	agentStartedAt *time.Time

	agentHeartbeatAt *time.Time

	installationRequestGeneration *int64

	installationReportGeneration *int64
}

func (f filler) machine() *fillerMachine {
	return &fillerMachine{
		f: f,
	}
}

func (m *fillerMachine) provided(p model.Provider, pid string) *fillerMachine {
	m.provider = &p
	m.providerID = &pid
	return m
}

func (m *fillerMachine) providedE(pid string) *fillerMachine {
	return m.provided(model.ProviderEquinix, pid)
}

func (m *fillerMachine) located(location string) *fillerMachine {
	m.location = &location
	return m
}

func (m *fillerMachine) hardware(threads int32, ramgb int64) *fillerMachine {
	m.threads = &threads
	m.ramgb = &ramgb
	return m
}

func (m *fillerMachine) agentStarted(t time.Time) *fillerMachine {
	m.agentStartedAt = &t
	return m
}

func (m *fillerMachine) agentHeartbeat(t time.Time) *fillerMachine {
	m.agentHeartbeatAt = &t
	return m
}

func (m *fillerMachine) agentHealthy() *fillerMachine {
	now := time.Now()
	return m.agentStarted(now.Add(-30 * time.Minute)).agentHeartbeat(now.Add(-1 * time.Minute))
}

func (m *fillerMachine) agentStoppedHeartbeating() *fillerMachine {
	now := time.Now()
	return m.agentStarted(now.Add(-30 * time.Minute)).agentHeartbeat(now.Add(-20 * time.Minute))
}

func (m *fillerMachine) agentNeverHeartbeat() *fillerMachine {
	now := time.Now()
	return m.agentStarted(now.Add(-30 * time.Minute))
}

func (m *fillerMachine) installRequested(gen int64) *fillerMachine {
	m.installationRequestGeneration = &gen
	return m
}

func (m *fillerMachine) installReported(gen int64) *fillerMachine {
	m.installationReportGeneration = &gen
	return m
}

func (m *fillerMachine) build() filler {
	return m.f.chain(func(ctx context.Context, q *model.Queries) error {
		mach, err := q.NewMachine(ctx)
		if err != nil {
			return err
		}
		if m.providerID != nil {
			err = q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
				MachineID:  mach.MachineID,
				Provider:   *m.provider,
				ProviderID: *m.providerID,
			})
			if err != nil {
				return err
			}
			if m.location != nil {
				err = q.MachineUpdateProviderStatus(ctx, model.MachineUpdateProviderStatusParams{
					ProviderID:       *m.providerID,
					Provider:         *m.provider,
					ProviderLocation: sql.NullString{Valid: true, String: *m.location},
				})
				if err != nil {
					return err
				}
			}
		}
		if m.threads != nil {
			report := api.AgentHardwareReport{
				Report: &aapi.Node{
					MemoryInstalledBytes: *m.ramgb << 30,
					MemoryUsableRatio:    1.0,
					Cpu: []*aapi.CPU{
						{
							HardwareThreads: *m.threads,
							Cores:           *m.threads,
						},
					},
				},
				Warning: nil,
			}
			raw, err := proto.Marshal(&report)
			if err != nil {
				return err
			}
			err = q.MachineSetHardwareReport(ctx, model.MachineSetHardwareReportParams{
				MachineID:         mach.MachineID,
				HardwareReportRaw: raw,
			})
			if err != nil {
				return err
			}
		}
		if m.agentStartedAt != nil {
			err = q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
				MachineID:      mach.MachineID,
				AgentStartedAt: *m.agentStartedAt,
				AgentPublicKey: []byte("fakefakefake"),
			})
			if err != nil {
				return err
			}
		}
		if m.agentHeartbeatAt != nil {
			err = q.MachineSetAgentHeartbeat(ctx, model.MachineSetAgentHeartbeatParams{
				MachineID:        mach.MachineID,
				AgentHeartbeatAt: *m.agentHeartbeatAt,
			})
			if err != nil {
				return err
			}
		}
		if m.installationRequestGeneration != nil {
			err = q.MachineSetOSInstallationRequest(ctx, model.MachineSetOSInstallationRequestParams{
				MachineID:  mach.MachineID,
				Generation: *m.installationRequestGeneration,
			})
			if err != nil {
				return err
			}
		}
		if m.installationReportGeneration != nil {
			err = q.MachineSetOSInstallationReport(ctx, model.MachineSetOSInstallationReportParams{
				MachineID:            mach.MachineID,
				Generation:           *m.installationReportGeneration,
				OsInstallationResult: model.MachineOsInstallationResultSuccess,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func TestHWStats(t *testing.T) {
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
	runner := newHWStatsRunner(&s, registry)

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	res, err := registry.Gather()
	if err != nil {
		t.Fatalf("Gather: %v", err)
	}
	if want, got := 0, len(res); want != got {
		t.Fatalf("Expected no metrics with empty database, got %d", got)
	}

	conn, err := s.Config.BMDB.Open(true)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("StartSession: %v", err)
	}
	// Populate database with some test data.
	err = sess.Transact(ctx, func(q *model.Queries) error {
		f := fill().
			machine().provided(model.ProviderEquinix, "1").hardware(32, 256).located("dark-bramble").build().
			machine().provided(model.ProviderEquinix, "2").hardware(32, 256).located("dark-bramble").build().
			machine().provided(model.ProviderEquinix, "3").hardware(32, 256).located("dark-bramble").build().
			machine().provided(model.ProviderEquinix, "4").hardware(32, 256).located("brittle-hollow").build().
			machine().provided(model.ProviderEquinix, "5").hardware(32, 256).located("timber-hearth").build().
			machine().provided(model.ProviderEquinix, "6").hardware(32, 256).located("timber-hearth").build()
		return f(ctx, q)
	})
	if err != nil {
		t.Fatalf("Transact: %v", err)
	}

	s.bmdb = conn
	s.sessionC = make(chan *bmdb.Session)
	go s.sessionWorker(ctx)

	// Do a statistics run and check results.
	if err := runner.runOnce(ctx); err != nil {
		t.Fatalf("runOnce: %v", err)
	}

	mfs, err := registry.Gather()
	if err != nil {
		t.Fatalf("Gatcher: %v", err)
	}

	// metric name -> provider -> location -> value
	values := make(map[string]map[string]map[string]float64)
	for _, mf := range mfs {
		values[*mf.Name] = make(map[string]map[string]float64)
		for _, m := range mf.Metric {
			var provider, location string
			for _, pair := range m.Label {
				switch *pair.Name {
				case "location":
					location = *pair.Value
				case "provider":
					provider = *pair.Value
				}
			}
			if _, ok := values[*mf.Name][provider]; !ok {
				values[*mf.Name][provider] = make(map[string]float64)
			}
			switch {
			case m.Gauge != nil && m.Gauge.Value != nil:
				values[*mf.Name][provider][location] = *m.Gauge.Value
			}
		}
	}

	for _, te := range []struct {
		provider model.Provider
		location string
		threads  int32
		ramgb    int64
	}{
		{model.ProviderEquinix, "dark-bramble", 96, 768},
		{model.ProviderEquinix, "brittle-hollow", 32, 256},
		{model.ProviderEquinix, "timber-hearth", 64, 512},
	} {
		threads := values["bmdb_hwstats_region_cpu_threads"][string(te.provider)][te.location]
		bytes := values["bmdb_hwstats_region_ram_bytes"][string(te.provider)][te.location]

		if want, got := te.threads, int32(threads); want != got {
			t.Errorf("Wanted %d threads in %s/%s, got %d", want, te.provider, te.location, got)
		}
		if want, got := te.ramgb, int64(bytes)>>30; want != got {
			t.Errorf("Wanted %d GB RAM in %s/%s, got %d", want, te.provider, te.location, got)
		}
	}
}
