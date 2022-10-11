package bmdb

import (
	"context"
	"errors"
	"testing"
	"time"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/lib/component"
)

func dut() *BMDB {
	return &BMDB{
		Config: Config{
			Database: component.CockroachConfig{
				InMemory: true,
			},
		},
	}
}

// TestSessionExpiry exercises the session heartbeat logic, making sure that if
// a session stops being maintained subsequent Transact calls will fail.
func TestSessionExpiry(t *testing.T) {
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

	// A transaction in a brand-new session should work.
	var machine model.Machine
	err = session.Transact(ctx, func(q *model.Queries) error {
		machine, err = q.NewMachine(ctx)
		return err
	})
	if err != nil {
		t.Fatalf("First transaction failed: %v", err)
	}

	time.Sleep(6 * time.Second)

	// A transaction after the 5-second session interval should continue to work.
	err = session.Transact(ctx, func(q *model.Queries) error {
		_, err = q.NewMachine(ctx)
		return err
	})
	if err != nil {
		t.Fatalf("Second transaction failed: %v", err)
	}

	// A transaction after the 5-second session interval should fail if we don't
	// maintain its heartbeat.
	session.ctxC()
	time.Sleep(6 * time.Second)

	err = session.Transact(ctx, func(q *model.Queries) error {
		return q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
			MachineID:  machine.MachineID,
			Provider:   "foo",
			ProviderID: "bar",
		})
	})
	if !errors.Is(err, ErrSessionExpired) {
		t.Fatalf("Second transaction should've failed due to expired session, got %v", err)
	}

}

// TestWork exercises the per-{process,machine} mutual exclusion mechanism of
// Work items.
func TestWork(t *testing.T) {
	b := dut()
	conn, err := b.Open(true)
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	// Start two session for testing.
	session1, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Starting session failed: %v", err)
	}
	session2, err := conn.StartSession(ctx)
	if err != nil {
		t.Fatalf("Starting session failed: %v", err)
	}

	var machine model.Machine
	err = session1.Transact(ctx, func(q *model.Queries) error {
		machine, err = q.NewMachine(ctx)
		return err
	})
	if err != nil {
		t.Fatalf("Creating machine failed: %v", err)
	}

	// Create a subcontext for a long-term work request. We'll cancel it later as
	// part of the test.
	ctxB, ctxBC := context.WithCancel(ctx)
	defer ctxBC()
	// Start work which will block forever. We have to go rendezvous through a
	// channel to make sure the work actually starts.
	started := make(chan error)
	done := make(chan error, 1)
	go func() {
		err := session1.Work(ctxB, machine.MachineID, model.ProcessUnitTest1, func() error {
			started <- nil
			<-ctxB.Done()
			return ctxB.Err()
		})
		done <- err
		if err != nil {
			started <- err
		}
	}()
	err = <-started
	if err != nil {
		t.Fatalf("Starting first work failed: %v", err)
	}

	// Starting more work on the same machine but a different process should still
	// be allowed.
	for _, session := range []*Session{session1, session2} {
		err = session.Work(ctxB, machine.MachineID, model.ProcessUnitTest2, func() error {
			return nil
		})
		if err != nil {
			t.Errorf("Could not run concurrent process on machine: %v", err)
		}
	}

	// However, starting work with the same process on the same machine should
	// fail.
	for _, session := range []*Session{session1, session2} {
		err = session.Work(ctxB, machine.MachineID, model.ProcessUnitTest1, func() error {
			return nil
		})
		if !errors.Is(err, ErrWorkConflict) {
			t.Errorf("Concurrent work with same process should've been forbidden, got %v", err)
		}
	}

	// Now, cancel the first long-running request and wait for it to return.
	ctxBC()
	err = <-done
	if !errors.Is(err, ctxB.Err()) {
		t.Fatalf("First work item should've failed with %v, got %v", ctxB.Err(), err)
	}

	// We should now be able to perform 'test1' work again against this machine.
	for _, session := range []*Session{session1, session2} {
		err = session.Work(ctx, machine.MachineID, model.ProcessUnitTest1, func() error {
			return nil
		})
		if err != nil {
			t.Errorf("Could not run work against machine: %v", err)
		}
	}
}
