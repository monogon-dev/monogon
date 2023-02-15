package bmdb

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/bmaas/bmdb/reflection"
)

// TestReflection exercises the BMDB reflection schema reflection and data
// retrieval code. Ideally this code would live in //cloud/bmaas/bmdb/reflection,
// but due to namespacing issues it lives here.
func TestReflection(t *testing.T) {
	b := dut()
	conn, err := b.Open(true)
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	sess, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("StartSession: %v", err)
	}

	// Create 10 test machines.
	var mids []uuid.UUID
	sess.Transact(ctx, func(q *model.Queries) error {
		for i := 0; i < 10; i += 1 {
			mach, err := q.NewMachine(ctx)
			if err != nil {
				return err
			}
			err = q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
				MachineID:  mach.MachineID,
				Provider:   model.ProviderEquinix,
				ProviderID: fmt.Sprintf("test-%d", i),
			})
			if err != nil {
				return err
			}
			mids = append(mids, mach.MachineID)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	// Start and fail work on one of the machines with an hour long backoff.
	w, err := sess.Work(ctx, model.ProcessUnitTest1, func(q *model.Queries) ([]uuid.UUID, error) {
		return mids[0:1], nil
	})
	if err != nil {
		t.Fatal(err)
	}
	to := time.Hour
	w.Fail(ctx, &to, "failure test")

	// On another machine, create a failure with a 1 second backoff.
	w, err = sess.Work(ctx, model.ProcessUnitTest1, func(q *model.Queries) ([]uuid.UUID, error) {
		return mids[1:2], nil
	})
	if err != nil {
		t.Fatal(err)
	}
	to = time.Second
	w.Fail(ctx, &to, "failure test")
	// Later on in the test we must wait for this backoff to actually elapse. Start
	// counting now.
	elapsed := time.NewTicker(to * 1)
	defer elapsed.Stop()

	// On another machine, create work and don't finish it yet.
	_, err = sess.Work(ctx, model.ProcessUnitTest1, func(q *model.Queries) ([]uuid.UUID, error) {
		return mids[2:3], nil
	})
	if err != nil {
		t.Fatal(err)
	}

	schema, err := conn.Reflect(ctx)
	if err != nil {
		t.Fatalf("ReflectTagTypes: %v", err)
	}

	// Dump all in strict mode.
	opts := &reflection.GetMachinesOpts{
		Strict: true,
	}
	res, err := schema.GetMachines(ctx, opts)
	if err != nil {
		t.Fatalf("Dump failed: %v", err)
	}
	if res.Query == "" {
		t.Errorf("Query not set on result")
	}
	machines := res.Data
	if want, got := 10, len(machines); want != got {
		t.Fatalf("Expected %d machines in dump, got %d", want, got)
	}

	// Expect Provided tag on all machines. Do a detailed check on fields, too.
	for _, machine := range machines {
		tag, ok := machine.Tags["Provided"]
		if !ok {
			t.Errorf("No Provided tag on machine.")
			continue
		}
		if want, got := "Provided", tag.Type.Name(); want != got {
			t.Errorf("Provided tag should have type %q, got %q", want, got)
		}
		if provider := tag.Field("provider"); provider != nil {
			if want, got := provider.HumanValue(), "Equinix"; want != got {
				t.Errorf("Wanted Provided.provider value %q, got %q", want, got)
			}
		} else {
			t.Errorf("Provider tag has no provider field")
		}
		if providerId := tag.Field("provider_id"); providerId != nil {
			if !strings.HasPrefix(providerId.HumanValue(), "test-") {
				t.Errorf("Unexpected provider_id value %q", providerId.HumanValue())
			}
		} else {
			t.Errorf("Provider tag has no provider_id field")
		}
	}

	// Now just dump one machine.
	opts.FilterMachine = &mids[0]
	res, err = schema.GetMachines(ctx, opts)
	if err != nil {
		t.Fatalf("Dump failed: %v", err)
	}
	machines = res.Data
	if want, got := 1, len(machines); want != got {
		t.Fatalf("Expected %d machines in dump, got %d", want, got)
	}
	if want, got := mids[0].String(), machines[0].ID.String(); want != got {
		t.Fatalf("Expected machine %s, got %s", want, got)
	}

	// Now dump a machine that doesn't exist. That should just return an empty list.
	fakeMid := uuid.New()
	opts.FilterMachine = &fakeMid
	res, err = schema.GetMachines(ctx, opts)
	if err != nil {
		t.Fatalf("Dump failed: %v", err)
	}
	machines = res.Data
	if want, got := 0, len(machines); want != got {
		t.Fatalf("Expected %d machines in dump, got %d", want, got)
	}

	// Finally, check the special case machines. The first one should have an active
	// backoff.
	opts.FilterMachine = &mids[0]
	res, err = schema.GetMachines(ctx, opts)
	if err != nil {
		t.Errorf("Dump failed: %v", err)
	} else {
		machine := res.Data[0]
		if _, ok := machine.Backoffs["UnitTest1"]; !ok {
			t.Errorf("Expected UnitTest1 backoff on machine")
		}
	}
	// The second one should have an expired backoff that shouldn't be reported in a
	// normal call..
	<-elapsed.C
	opts.FilterMachine = &mids[1]
	res, err = schema.GetMachines(ctx, opts)
	if err != nil {
		t.Errorf("Dump failed: %v", err)
	} else {
		machine := res.Data[0]
		if _, ok := machine.Backoffs["UnitTest1"]; ok {
			t.Errorf("Expected no UnitTest1 backoff on machine")
		}
	}
	// But if we ask for expired backoffs, we should get it.
	opts.ExpiredBackoffs = true
	res, err = schema.GetMachines(ctx, opts)
	if err != nil {
		t.Errorf("Dump failed: %v", err)
	} else {
		machine := res.Data[0]
		if _, ok := machine.Backoffs["UnitTest1"]; !ok {
			t.Errorf("Expected UnitTest1 backoff on machine")
		}
	}
	// Finally, the third machine should have an active Work item.
	opts.FilterMachine = &mids[2]
	res, err = schema.GetMachines(ctx, opts)
	if err != nil {
		t.Errorf("Dump failed: %v", err)
	} else {
		machine := res.Data[0]
		if _, ok := machine.Work["UnitTest1"]; !ok {
			t.Errorf("Expected UnitTest1 work item on machine")
		}
	}
}
