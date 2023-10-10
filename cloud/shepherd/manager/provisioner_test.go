package manager

import (
	"context"
	"testing"
	"time"

	"golang.org/x/time/rate"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
	"source.monogon.dev/cloud/shepherd"
)

// TestProvisionerSmokes makes sure the Provisioner doesn't go up in flames on
// the happy path.
func TestProvisionerSmokes(t *testing.T) {
	pc := ProvisionerConfig{
		MaxCount: 10,
		// We need 3 iterations to provide 10 machines with a chunk size of 4.
		ReconcileLoopLimiter:  rate.NewLimiter(rate.Every(10*time.Second), 3),
		DeviceCreationLimiter: rate.NewLimiter(rate.Every(time.Second), 10),
		ChunkSize:             4,
	}

	provider := newDummyProvider(100)

	p, err := NewProvisioner(provider, pc)
	if err != nil {
		t.Fatalf("Could not create Provisioner: %v", err)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	b := bmdb.BMDB{
		Config: bmdb.Config{
			Database: component.CockroachConfig{
				InMemory: true,
			},
			ComponentName: "test",
			RuntimeInfo:   "test",
		},
	}
	conn, err := b.Open(true)
	if err != nil {
		t.Fatalf("Could not create in-memory BMDB: %v", err)
	}

	go p.Run(ctx, conn)

	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Failed to create BMDB session for verification: %v", err)
	}
	for {
		time.Sleep(100 * time.Millisecond)

		var provided []model.MachineProvided
		err = sess.Transact(ctx, func(q *model.Queries) error {
			var err error
			provided, err = q.GetProvidedMachines(ctx, provider.Type())
			return err
		})
		if err != nil {
			t.Errorf("Transact failed: %v", err)
		}
		if len(provided) < 10 {
			continue
		}
		if len(provided) > 10 {
			t.Errorf("%d machines provided (limit: 10)", len(provided))
		}

		for _, mp := range provided {
			if provider.machines[shepherd.ProviderID(mp.ProviderID)] == nil {
				t.Errorf("BMDB machine %q has unknown provider ID %q", mp.MachineID, mp.ProviderID)
			}
		}

		return
	}
}

// TestProvisioner_resolvePossiblyUsed makes sure the PossiblyUsed state is
// resolved correctly.
func TestProvisioner_resolvePossiblyUsed(t *testing.T) {
	const providedMachineID = "provided-machine"

	providedMachines := map[shepherd.ProviderID]bool{
		providedMachineID: true,
	}

	tests := []struct {
		name         string
		machineID    shepherd.ProviderID
		machineState shepherd.State
		wantedState  shepherd.State
	}{
		{
			name:         "skip KnownUsed",
			machineState: shepherd.StateKnownUsed,
			wantedState:  shepherd.StateKnownUsed,
		},
		{
			name:         "skip KnownUnused",
			machineState: shepherd.StateKnownUnused,
			wantedState:  shepherd.StateKnownUnused,
		},
		{
			name:         "invalid ID",
			machineID:    shepherd.InvalidProviderID,
			machineState: shepherd.StatePossiblyUsed,
			wantedState:  shepherd.StateKnownUnused,
		},
		{
			name:         "valid ID, not in providedMachines",
			machineID:    "unused-machine",
			machineState: shepherd.StatePossiblyUsed,
			wantedState:  shepherd.StateKnownUnused,
		},
		{
			name:         "valid ID, in providedMachines",
			machineID:    providedMachineID,
			machineState: shepherd.StatePossiblyUsed,
			wantedState:  shepherd.StateKnownUsed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provisioner{}
			if got := p.resolvePossiblyUsed(&dummyMachine{id: tt.machineID, state: tt.machineState}, providedMachines); got != tt.wantedState {
				t.Errorf("resolvePossiblyUsed() = %v, want %v", got, tt.wantedState)
			}
		})
	}
}
