package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/packethost/packngo"
	"golang.org/x/time/rate"

	"source.monogon.dev/cloud/bmaas/bmdb"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
	"source.monogon.dev/cloud/shepherd/manager"
)

type initializerDut struct {
	f        *fakequinix
	i        *manager.Initializer
	bmdb     *bmdb.Connection
	ctx      context.Context
	provider *equinixProvider
}

func newInitializerDut(t *testing.T) *initializerDut {
	t.Helper()

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

	ic := manager.InitializerConfig{
		ControlLoopConfig: manager.ControlLoopConfig{
			DBQueryLimiter: rate.NewLimiter(rate.Every(time.Second), 10),
		},
		Executable:        []byte("beep boop i'm a real program"),
		TargetPath:        "/fake/path",
		Endpoint:          "example.com:1234",
		SSHConnectTimeout: time.Second,
		SSHExecTimeout:    time.Second,
	}

	i, err := manager.NewInitializer(provider, &manager.FakeSSHClient{}, ic)
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

	if err := provider.SSHEquinixEnsure(ctx); err != nil {
		t.Fatalf("Failed to ensure SSH key: %v", err)
	}
	go manager.RunControlLoop(ctx, conn, i)

	return &initializerDut{
		f:        f,
		i:        i,
		bmdb:     conn,
		ctx:      ctx,
		provider: provider,
	}
}

// TestInitializerSmokes makes sure the Initializer doesn't go up in flames on
// the happy path.
func TestInitializerSmokes(t *testing.T) {
	dut := newInitializerDut(t)
	f := dut.f
	ctx := dut.ctx
	conn := dut.bmdb

	reservations, _ := f.ListReservations(ctx, f.pid)
	kid, err := dut.provider.sshEquinixId(ctx)
	if err != nil {
		t.Fatalf("Failed to retrieve equinix key ID: %v", err)
	}
	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Failed to create BMDB session for verifiaction: %v", err)
	}

	// Create 10 provided machines for testing.
	for i := 0; i < 10; i++ {
		res := reservations[i]
		dev, _ := f.CreateDevice(ctx, &packngo.DeviceCreateRequest{
			Hostname:              fmt.Sprintf("test-%d", i),
			OS:                    "fake",
			ProjectID:             f.pid,
			HardwareReservationID: res.ID,
			ProjectSSHKeys:        []string{kid},
		})
		f.devices[dev.ID].Network = []*packngo.IPAddressAssignment{
			{
				IpAddressCommon: packngo.IpAddressCommon{
					ID:            "fake",
					Address:       "1.2.3.4",
					Management:    true,
					AddressFamily: 4,
					Public:        true,
				},
			},
		}
		err = sess.Transact(ctx, func(q *model.Queries) error {
			machine, err := q.NewMachine(ctx)
			if err != nil {
				return err
			}
			return q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
				MachineID:  machine.MachineID,
				Provider:   model.ProviderEquinix,
				ProviderID: dev.ID,
			})
		})
		if err != nil {
			t.Fatalf("Failed to create BMDB machine: %v", err)
		}
	}

	// Expect to find 0 machines needing start.
	for {
		time.Sleep(100 * time.Millisecond)

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
		if len(machines) == 0 {
			break
		}
	}
}
