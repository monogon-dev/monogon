package manager

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"testing"
	"time"

	"golang.org/x/time/rate"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
)

// TestProvisionerSmokes makes sure the Provisioner doesn't go up in flames on
// the happy path.
func TestProvisionerSmokes(t *testing.T) {
	pc := ProvisionerConfig{
		OS:       "fake",
		MaxCount: 10,
		// We need 3 iterations to provide 10 machines with a chunk size of 4.
		ReconcileLoopLimiter:  rate.NewLimiter(rate.Every(10*time.Second), 3),
		DeviceCreationLimiter: rate.NewLimiter(rate.Every(time.Second), 10),
		ReservationChunkSize:  4,
	}
	_, key, _ := ed25519.GenerateKey(rand.Reader)
	sc := SharedConfig{
		ProjectId:    "noproject",
		KeyLabel:     "somekey",
		Key:          key,
		DevicePrefix: "test-",
	}
	f := newFakequinix(sc.ProjectId, 100)
	p, err := pc.New(f, &sc)
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

	if err := sc.SSHEquinixEnsure(ctx, f); err != nil {
		t.Fatalf("Failed to ensure SSH key: %v", err)
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
			provided, err = q.GetProvidedMachines(ctx, model.ProviderEquinix)
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
			if f.devices[mp.ProviderID] == nil {
				t.Errorf("BMDB machine %q has unknown provider ID %q", mp.MachineID, mp.ProviderID)
			}
		}

		return
	}
}
