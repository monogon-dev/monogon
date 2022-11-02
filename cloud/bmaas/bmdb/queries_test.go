package bmdb

import (
	"context"
	"fmt"
	"testing"
	"time"

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
	expectNoCandidates := func() {
		if err := session.Transact(ctx, func(q *model.Queries) error {
			candidates, err := q.GetMachinesForAgentStart(ctx, 1)
			if err != nil {
				t.Fatalf("Could not retrieve machines for agent start: %v", err)
			}
			if want, got := 0, len(candidates); want != got {
				t.Fatalf("Wanted %d machines for agent start, got %+v", want, candidates)
			}
			return nil
		}); err != nil {
			t.Fatal(err)
		}
	}
	expectNoCandidates()

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
	if err := session.Transact(ctx, func(q *model.Queries) error {
		candidates, err := q.GetMachinesForAgentStart(ctx, 1)
		if err != nil {
			t.Fatalf("Could not retrieve machines for agent start: %v", err)
		}
		if want, got := 1, len(candidates); want != got {
			t.Fatalf("Wanted %d machines for agent start, got %+v", want, candidates)
		}
		if want, got := machine.MachineID, candidates[0].MachineID; want != got {
			t.Fatalf("Wanted %s for agent start, got %s", want, got)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

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
	expectNoCandidates()
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
	}{
		// No start, no heartbeat -> no recovery expected.
		{false, time.Time{}, time.Time{}},
		// Started recently, no heartbeat -> no recovery expected.
		{false, time.Now(), time.Time{}},
		// Started a while ago, heartbeat active -> no recovery expected.
		{false, time.Now().Add(-40 * time.Minute), time.Now()},

		// Started a while ago, no heartbeat -> recovery expected.
		{true, time.Now().Add(-40 * time.Minute), time.Time{}},
		// Started a while ago, no recent heartbeat -> recovery expected.
		{true, time.Now().Add(-40 * time.Minute), time.Now().Add(-20 * time.Minute)},
	} {
		if err := session.Transact(ctx, func(q *model.Queries) error {
			machine, err := q.NewMachine(ctx)
			if err != nil {
				return fmt.Errorf("NewMachine: %w", err)
			}
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
			return nil
		}); err != nil {
			t.Errorf("%d: setup failed: %v", i, err)
			continue
		}

		if err := session.Transact(ctx, func(q *model.Queries) error {
			candidates, err := q.GetMachineForAgentRecovery(ctx, 1)
			if err != nil {
				return fmt.Errorf("GetMachinesForAgentRecovery: %w", err)
			}
			if scenario.wantRun && len(candidates) == 0 {
				return fmt.Errorf("machine unscheduled for recovery")
			}
			if !scenario.wantRun && len(candidates) != 0 {
				return fmt.Errorf("machine scheduled for recovery")
			}
			return nil
		}); err != nil {
			t.Errorf("%d: test failed: %v", i, err)
		}
	}
}
