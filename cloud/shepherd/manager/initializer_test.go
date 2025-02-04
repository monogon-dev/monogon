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
)

// TestInitializerSmokes makes sure the Initializer doesn't go up in flames on
// the happy path.
func TestInitializerSmokes(t *testing.T) {
	provider := newDummyProvider(100)

	ic := InitializerConfig{
		ControlLoopConfig: ControlLoopConfig{
			DBQueryLimiter: rate.NewLimiter(rate.Every(time.Second), 10),
		},
		Executable:        []byte("beep boop i'm a real program"),
		TargetPath:        "/fake/path",
		Endpoint:          "example.com:1234",
		SSHConnectTimeout: time.Second,
		SSHExecTimeout:    time.Second,
	}

	i, err := NewInitializer(provider, provider.sshClient(), ic)
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

	go RunControlLoop(ctx, conn, i)

	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Failed to create BMDB session for verifiaction: %v", err)
	}

	// Create 10 provided machines for testing.
	if _, err := provider.createDummyMachines(ctx, sess, 10); err != nil {
		t.Fatalf("Failed to create dummy machines: %v", err)
	}

	// Expect to find 0 machines needing start.
	for {
		time.Sleep(100 * time.Millisecond)

		var machines []model.MachineProvided
		err = sess.Transact(ctx, func(q *model.Queries) error {
			var err error
			machines, err = q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
				Limit:    100,
				Provider: provider.Type(),
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

	provider.muMachines.RLock()
	defer provider.muMachines.RUnlock()
	for _, m := range provider.machines {
		if !m.agentStarted {
			t.Fatalf("Initializer didn't start agent on machine %q", m.id)
		}
	}
}
