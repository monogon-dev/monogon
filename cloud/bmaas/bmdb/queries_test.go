package bmdb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
)

// TestAgentStart exercises GetMachinesForAgentStart.
func TestAgentStart(t *testing.T) {
	b := dut()
	conn, err := b.Open(true)
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	session, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Starting session failed: %v", err)
	}

	// Create a test machine.
	var machine model.Machine
	err = session.Transact(ctx, func(q *model.Queries) error {
		machine, err = q.NewMachine(ctx)
		return err
	})
	if err != nil {
		t.Fatalf("Creating machine failed: %v", err)
	}

	// It should be, by default, not a candidate for agent start as it's not yet
	// provided by any provider.
	expectCandidates := func(want int) {
		t.Helper()
		if err := session.Transact(ctx, func(q *model.Queries) error {
			candidates, err := q.GetMachinesForAgentStart(ctx, 1)
			if err != nil {
				t.Fatalf("Could not retrieve machines for agent start: %v", err)
			}
			if got := len(candidates); want != got {
				t.Fatalf("Wanted %d machines for agent start, got %+v", want, candidates)
			}
			return nil
		}); err != nil {
			t.Fatal(err)
		}
	}

	// Provide machine, and check it is now a candidate.
	if err := session.Transact(ctx, func(q *model.Queries) error {
		return q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
			MachineID:  machine.MachineID,
			Provider:   model.ProviderEquinix,
			ProviderID: "123",
		})
	}); err != nil {
		t.Fatalf("could not add provided tag to machine: %v", err)
	}
	expectCandidates(1)

	// Add a start tag. Machine shouldn't be a candidate anymore.
	if err := session.Transact(ctx, func(q *model.Queries) error {
		return q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
			MachineID:      machine.MachineID,
			AgentStartedAt: time.Now(),
			AgentPublicKey: []byte("fakefakefakefake"),
		})
	}); err != nil {
		t.Fatalf("could not add provided tag to machine: %v", err)
	}
	expectCandidates(0)

	// Add a new machine which has an unfulfilled installation request. It should be
	// a candidate.
	if err = session.Transact(ctx, func(q *model.Queries) error {
		machine, err = q.NewMachine(ctx)
		if err != nil {
			return err
		}
		if err := q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
			MachineID:  machine.MachineID,
			Provider:   model.ProviderEquinix,
			ProviderID: "234",
		}); err != nil {
			return err
		}
		if err := q.MachineSetOSInstallationRequest(ctx, model.MachineSetOSInstallationRequestParams{
			MachineID:  machine.MachineID,
			Generation: 10,
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		t.Fatalf("could not add new machine with installation request: %v", err)
	}
	expectCandidates(1)

	// Fulfill installation request on machine with an older generation. it should
	// remain a candidate.
	if err = session.Transact(ctx, func(q *model.Queries) error {
		if err := q.MachineSetOSInstallationReport(ctx, model.MachineSetOSInstallationReportParams{
			MachineID:  machine.MachineID,
			Generation: 9,
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		t.Fatalf("could not fulfill installation request with older generation: %v", err)
	}
	expectCandidates(1)

	// Fulfill installation request with correct generation. The machine should not
	// be a candidate anymore.
	if err = session.Transact(ctx, func(q *model.Queries) error {
		if err := q.MachineSetOSInstallationReport(ctx, model.MachineSetOSInstallationReportParams{
			MachineID:  machine.MachineID,
			Generation: 10,
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		t.Fatalf("could not fulfill installation request with current generation: %v", err)
	}
	expectCandidates(0)
}

// TestAgentRecovery exercises GetMachinesForAgentRecovery though a few
// different scenarios in which a test machine is present with different tags
// set.
func TestAgentRecovery(t *testing.T) {
	b := dut()
	conn, err := b.Open(true)
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	session, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Starting session failed: %v", err)
	}

	for i, scenario := range []struct {
		// Whether recovery is expected to run.
		wantRun bool
		// started will add a AgentStarted tag for a given time, if set.
		started time.Time
		// heartbeat will add a AgentHeartbeat tag for a given time, if set.
		heartbeat time.Time
		// requestGeneration will populate a OSInstallationRequest if not zero.
		requestGeneration int64
		// requestGeneration will populate a OSInstallationResponse if not zero.
		reportGeneration int64
	}{
		// No start, no heartbeat -> no recovery expected.
		{false, time.Time{}, time.Time{}, 0, 0},
		// Started recently, no heartbeat -> no recovery expected.
		{false, time.Now(), time.Time{}, 0, 0},
		// Started a while ago, heartbeat active -> no recovery expected.
		{false, time.Now().Add(-40 * time.Minute), time.Now(), 0, 0},

		// Started a while ago, no heartbeat -> recovery expected.
		{true, time.Now().Add(-40 * time.Minute), time.Time{}, 0, 0},
		// Started a while ago, no recent heartbeat -> recovery expected.
		{true, time.Now().Add(-40 * time.Minute), time.Now().Add(-20 * time.Minute), 0, 0},

		// Installation request without report -> recovery expected.
		{true, time.Now().Add(-40 * time.Minute), time.Time{}, 10, 0},
		{true, time.Now().Add(-40 * time.Minute), time.Now().Add(-20 * time.Minute), 10, 0},
		// Installation request mismatching report -> recovery expected.
		{true, time.Now().Add(-40 * time.Minute), time.Time{}, 10, 9},
		{true, time.Now().Add(-40 * time.Minute), time.Now().Add(-20 * time.Minute), 10, 9},
		// Installation request matching report -> no recovery expected.
		{false, time.Now().Add(-40 * time.Minute), time.Time{}, 10, 10},
		{false, time.Now().Add(-40 * time.Minute), time.Now().Add(-20 * time.Minute), 10, 10},
	} {
		var machineID uuid.UUID
		if err := session.Transact(ctx, func(q *model.Queries) error {
			machine, err := q.NewMachine(ctx)
			if err != nil {
				return fmt.Errorf("NewMachine: %w", err)
			}
			machineID = machine.MachineID
			if err := q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
				MachineID:  machine.MachineID,
				Provider:   model.ProviderEquinix,
				ProviderID: fmt.Sprintf("test-%d", i),
			}); err != nil {
				return fmt.Errorf("MachineAddProvided: %w", err)
			}
			if !scenario.started.IsZero() {
				if err := q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
					MachineID:      machine.MachineID,
					AgentStartedAt: scenario.started,
					AgentPublicKey: []byte("fake"),
				}); err != nil {
					return fmt.Errorf("MachineSetAgentStarted: %w", err)
				}
			}
			if !scenario.heartbeat.IsZero() {
				if err := q.MachineSetAgentHeartbeat(ctx, model.MachineSetAgentHeartbeatParams{
					MachineID:        machine.MachineID,
					AgentHeartbeatAt: scenario.heartbeat,
				}); err != nil {
					return fmt.Errorf("MachineSetAgentHeartbeat: %w", err)
				}
			}
			if scenario.requestGeneration != 0 {
				if err := q.MachineSetOSInstallationRequest(ctx, model.MachineSetOSInstallationRequestParams{
					MachineID:  machine.MachineID,
					Generation: scenario.requestGeneration,
				}); err != nil {
					return fmt.Errorf("MachineSetOSInstallationRequest: %w", err)
				}
			}
			if scenario.reportGeneration != 0 {
				if err := q.MachineSetOSInstallationReport(ctx, model.MachineSetOSInstallationReportParams{
					MachineID:  machine.MachineID,
					Generation: scenario.reportGeneration,
				}); err != nil {
					return fmt.Errorf("MachineSetOSInstallationReport: %w", err)
				}
			}
			return nil
		}); err != nil {
			t.Errorf("%d: setup failed: %v", i, err)
			continue
		}

		found := false
		if err := session.Transact(ctx, func(q *model.Queries) error {
			candidates, err := q.GetMachineForAgentRecovery(ctx, 100)
			if err != nil {
				return fmt.Errorf("GetMachinesForAgentRecovery: %w", err)
			}
			for _, c := range candidates {
				if c.MachineID == machineID {
					found = true
					break
				}
			}
			return nil
		}); err != nil {
			t.Errorf("%d: failed to retrieve candidates: %v", i, err)
		}
		if scenario.wantRun && !found {
			t.Errorf("%d: expected recovery but not scheduled", i)
		}
		if !scenario.wantRun && found {
			t.Errorf("%d: expected no recovery but is scheduled", i)
		}
	}
}
