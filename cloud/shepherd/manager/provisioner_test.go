// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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
			t.Fatalf("Transact failed: %v", err)
		}
		if len(provided) < 10 {
			continue
		}
		if len(provided) > 10 {
			t.Fatalf("%d machines provided (limit: 10)", len(provided))
		}

		for _, mp := range provided {
			provider.muMachines.RLock()
			if provider.machines[shepherd.ProviderID(mp.ProviderID)] == nil {
				t.Fatalf("BMDB machine %q has unknown provider ID %q", mp.MachineID, mp.ProviderID)
			}
			provider.muMachines.RUnlock()
		}

		return
	}
}

// TestProvisioner_resolvePossiblyUsed makes sure the PossiblyUsed availability is
// resolved correctly.
func TestProvisioner_resolvePossiblyUsed(t *testing.T) {
	const providedMachineID = "provided-machine"

	providedMachines := map[shepherd.ProviderID]shepherd.Machine{
		providedMachineID: nil,
	}

	tests := []struct {
		name                string
		machineID           shepherd.ProviderID
		machineAvailability shepherd.Availability
		wantedAvailability  shepherd.Availability
	}{
		{
			name:                "skip KnownUsed",
			machineAvailability: shepherd.AvailabilityKnownUsed,
			wantedAvailability:  shepherd.AvailabilityKnownUsed,
		},
		{
			name:                "skip KnownUnused",
			machineAvailability: shepherd.AvailabilityKnownUnused,
			wantedAvailability:  shepherd.AvailabilityKnownUnused,
		},
		{
			name:                "invalid ID",
			machineID:           shepherd.InvalidProviderID,
			machineAvailability: shepherd.AvailabilityPossiblyUsed,
			wantedAvailability:  shepherd.AvailabilityKnownUnused,
		},
		{
			name:                "valid ID, not in providedMachines",
			machineID:           "unused-machine",
			machineAvailability: shepherd.AvailabilityPossiblyUsed,
			wantedAvailability:  shepherd.AvailabilityKnownUnused,
		},
		{
			name:                "valid ID, in providedMachines",
			machineID:           providedMachineID,
			machineAvailability: shepherd.AvailabilityPossiblyUsed,
			wantedAvailability:  shepherd.AvailabilityKnownUsed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provisioner{}
			if got := p.resolvePossiblyUsed(&dummyMachine{id: tt.machineID, availability: tt.machineAvailability}, providedMachines); got != tt.wantedAvailability {
				t.Fatalf("resolvePossiblyUsed() = %v, want %v", got, tt.wantedAvailability)
			}
		})
	}
}
