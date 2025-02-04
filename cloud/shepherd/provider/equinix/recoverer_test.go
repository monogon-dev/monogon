// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"testing"
	"time"

	"github.com/packethost/packngo"
	"golang.org/x/time/rate"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
	"source.monogon.dev/cloud/shepherd/manager"
)

type recovererDut struct {
	f    *fakequinix
	r    *manager.Recoverer
	bmdb *bmdb.Connection
	ctx  context.Context
}

func newRecovererDut(t *testing.T) *recovererDut {
	t.Helper()

	rc := manager.RecovererConfig{
		ControlLoopConfig: manager.ControlLoopConfig{
			DBQueryLimiter: rate.NewLimiter(rate.Every(time.Second), 10),
		},
	}

	sc := providerConfig{
		ProjectId:    "noproject",
		KeyLabel:     "somekey",
		DevicePrefix: "test-",
	}

	_, key, _ := ed25519.GenerateKey(rand.Reader)
	k := manager.SSHKey{
		Key: key,
	}

	f := newFakequinix(sc.ProjectId, 100)
	provider, err := sc.New(&k, f)
	if err != nil {
		t.Fatalf("Could not create Provider: %v", err)
	}

	r, err := manager.NewRecoverer(provider, rc)
	if err != nil {
		t.Fatalf("Could not create Initializer: %v", err)
	}

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

	ctx, ctxC := context.WithCancel(context.Background())
	t.Cleanup(ctxC)

	go manager.RunControlLoop(ctx, conn, r)

	return &recovererDut{
		f:    f,
		r:    r,
		bmdb: conn,
		ctx:  ctx,
	}
}

// TestRecoverySmokes makes sure that the Initializer in recovery mode doesn't go
// up in flames on the happy path.
func TestRecoverySmokes(t *testing.T) {
	dut := newRecovererDut(t)
	f := dut.f
	ctx := dut.ctx
	conn := dut.bmdb

	reservations, _ := f.ListReservations(ctx, "fake")

	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Failed to create BMDB session: %v", err)
	}

	// Create test machine that should be selected for recovery.
	// First in Fakequinix...
	dev, _ := f.CreateDevice(ctx, &packngo.DeviceCreateRequest{
		Hostname:              "test-devices",
		OS:                    "fake",
		ProjectID:             "fake",
		HardwareReservationID: reservations[0].ID,
		ProjectSSHKeys:        []string{},
	})
	// ... and in BMDB.
	err = sess.Transact(ctx, func(q *model.Queries) error {
		machine, err := q.NewMachine(ctx)
		if err != nil {
			return err
		}
		err = q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
			MachineID:  machine.MachineID,
			Provider:   model.ProviderEquinix,
			ProviderID: dev.ID,
		})
		if err != nil {
			return err
		}
		return q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
			MachineID:      machine.MachineID,
			AgentStartedAt: time.Now().Add(time.Hour * -10),
			AgentPublicKey: []byte("fakefakefakefake"),
		})
	})
	if err != nil {
		t.Fatalf("Failed to create test machine: %v", err)
	}

	// Expect to find 0 machines needing recovery.
	deadline := time.Now().Add(10 * time.Second)
	for {
		if time.Now().After(deadline) {
			t.Fatalf("Machines did not get processed in time")
		}
		time.Sleep(100 * time.Millisecond)

		var machines []model.MachineProvided
		err = sess.Transact(ctx, func(q *model.Queries) error {
			var err error
			machines, err = q.GetMachineForAgentRecovery(ctx, model.GetMachineForAgentRecoveryParams{
				Limit:    100,
				Provider: model.ProviderEquinix,
			})
			return err
		})
		if err != nil {
			t.Fatalf("Failed to run Transaction: %v", err)
		}
		if len(machines) == 0 {
			break
		}
	}

	// Expect the target machine to have been rebooted.
	dut.f.mu.Lock()
	reboots := dut.f.reboots[dev.ID]
	dut.f.mu.Unlock()
	if want, got := 1, reboots; want != got {
		t.Fatalf("Wanted %d reboot, got %d", want, got)
	}

	// Expect machine to now be available again for agent start.
	var machines []model.MachineProvided
	err = sess.Transact(ctx, func(q *model.Queries) error {
		var err error
		machines, err = q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
			Limit:    100,
			Provider: model.ProviderEquinix,
		})
		return err
	})
	if err != nil {
		t.Fatalf("Failed to run Transaction: %v", err)
	}
	if want, got := 1, len(machines); want != got {
		t.Fatalf("Wanted %d machine ready for agent start, got %d", want, got)
	}
}
