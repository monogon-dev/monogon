// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package bmdb

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	apb "source.monogon.dev/cloud/agent/api"
	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/bmaas/bmdb/reflection"
	"source.monogon.dev/cloud/bmaas/server/api"
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
	err = sess.Transact(ctx, func(q *model.Queries) error {
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
	backoff := Backoff{
		Initial: time.Hour,
	}
	w.Fail(ctx, &backoff, "failure test")

	// On another machine, create a failure with a 1 second backoff.
	w, err = sess.Work(ctx, model.ProcessUnitTest1, func(q *model.Queries) ([]uuid.UUID, error) {
		return mids[1:2], nil
	})
	if err != nil {
		t.Fatal(err)
	}
	backoff = Backoff{
		Initial: time.Second,
	}
	w.Fail(ctx, &backoff, "failure test")
	// Later on in the test we must wait for this backoff to actually elapse. Start
	// counting now.
	elapsed := time.NewTicker(time.Second * 1)
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

// TestReflectionProtoFields ensures that the basic proto field introspection
// functionality works.
func TestReflectionProtoFields(t *testing.T) {
	s := dut()
	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	bmdb, err := s.Open(true)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	sess, err := bmdb.StartSession(ctx)
	if err != nil {
		t.Fatalf("StartSession: %v", err)
	}
	var machine model.Machine
	err = sess.Transact(ctx, func(q *model.Queries) error {
		machine, err = q.NewMachine(ctx)
		if err != nil {
			return err
		}

		report := &api.AgentHardwareReport{
			Report: &apb.Node{
				Manufacturer:         "Charles Babbage",
				Product:              "Analytical Engine",
				SerialNumber:         "183701",
				MemoryInstalledBytes: 14375,
				MemoryUsableRatio:    1.0,
				Cpu: []*apb.CPU{
					{
						Architecture:    nil,
						HardwareThreads: 1,
						Cores:           1,
					},
				},
			},
			Warning: []string{"something went wrong"},
		}
		b, _ := proto.Marshal(report)
		return q.MachineSetHardwareReport(ctx, model.MachineSetHardwareReportParams{
			MachineID:         machine.MachineID,
			HardwareReportRaw: b,
		})
	})
	if err != nil {
		t.Fatalf("Failed to submit hardware report: %v", err)
	}

	schem, err := bmdb.Reflect(ctx)
	if err != nil {
		t.Fatalf("Failed to reflect on database: %v", err)
	}

	machines, err := schem.GetMachines(ctx, &reflection.GetMachinesOpts{FilterMachine: &machine.MachineID, Strict: true})
	if err != nil {
		t.Fatalf("Failed to get machine: %v", err)
	}
	if len(machines.Data) != 1 {
		t.Errorf("Expected one machine, got %d", len(machines.Data))
	} else {
		machine := machines.Data[0]
		ty := machine.Tags["HardwareReport"].Field("hardware_report_raw").Type.HumanType()
		if want, got := "cloud.bmaas.server.api.AgentHardwareReport", ty; want != got {
			t.Errorf("Mismatch in type: wanted %q, got %q", want, got)
		}
		v := machine.Tags["HardwareReport"].Field("hardware_report_raw").HumanValue()
		if !strings.Contains(v, "manufacturer:") {
			t.Errorf("Invalid serialized prototext: %s", v)
		}
		fv, err := machine.Tags["HardwareReport"].Field("hardware_report_raw").Index("report.cpu[0].cores")
		if err != nil {
			t.Errorf("Could not get report.cpu[0].cores from hardware_report_raw: %v", err)
		} else {
			if want, got := "1", fv; want != got {
				t.Errorf("report.cpu[0].cores should be %q, got %q", want, got)
			}
		}
	}
}
