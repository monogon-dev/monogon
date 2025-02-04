// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"testing"
	"time"

	"github.com/packethost/packngo"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
)

type updaterDut struct {
	f    *fakequinix
	u    *Updater
	bmdb *bmdb.Connection
	ctx  context.Context
}

func newUpdaterDut(t *testing.T) *updaterDut {
	t.Helper()

	uc := UpdaterConfig{
		Enable:        true,
		IterationRate: time.Second,
	}

	f := newFakequinix("fake", 100)
	u, err := uc.New(f)
	if err != nil {
		t.Fatalf("Could not create Updater: %v", err)
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

	go u.Run(ctx, conn)

	return &updaterDut{
		f:    f,
		u:    u,
		bmdb: conn,
		ctx:  ctx,
	}
}

func TestUpdater(t *testing.T) {
	dut := newUpdaterDut(t)
	f := dut.f
	ctx := dut.ctx
	conn := dut.bmdb

	reservations, _ := f.ListReservations(ctx, "fake")

	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Failed to create BMDB session: %v", err)
	}

	// Create test machine that should be selected for updating.
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
		t.Fatalf("failed to execute bmdb transaction: %v", err)
	}

	deadline := time.Now().Add(time.Second * 10)
	for {
		time.Sleep(100 * time.Millisecond)
		if time.Now().After(deadline) {
			t.Fatalf("Deadline exceeded")
		}

		var provided []model.MachineProvided
		err = sess.Transact(ctx, func(q *model.Queries) error {
			var err error
			provided, err = q.GetProvidedMachines(ctx, model.ProviderEquinix)
			return err
		})
		if err != nil {
			t.Fatalf("Transact: %v", err)
		}
		if len(provided) < 1 {
			continue
		}
		p := provided[0]
		if p.ProviderStatus.ProviderStatus != model.ProviderStatusRunning {
			continue
		}
		if p.ProviderLocation.String != "wad" {
			continue
		}
		if p.ProviderIpAddress.String != "1.2.3.4" {
			continue
		}
		if p.ProviderReservationID.String != reservations[0].ID {
			continue
		}
		break
	}
}
