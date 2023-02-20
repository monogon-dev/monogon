package bmdb

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"

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

	constantRetriever := func(_ *model.Queries) ([]uuid.UUID, error) {
		return []uuid.UUID{machine.MachineID}, nil
	}

	// Start work on machine which we're not gonna finish for a while.
	work1, err := session1.Work(ctxB, model.ProcessUnitTest1, constantRetriever)
	if err != nil {
		t.Fatalf("Starting first work failed: %v", err)
	}

	// Starting more work on the same machine but a different process should still
	// be allowed.
	for _, session := range []*Session{session1, session2} {
		work2, err := session.Work(ctxB, model.ProcessUnitTest2, constantRetriever)
		if err != nil {
			t.Errorf("Could not run concurrent process on machine: %v", err)
		} else {
			work2.Cancel(ctxB)
		}
	}

	// However, starting work with the same process on the same machine should
	// fail.
	for _, session := range []*Session{session1, session2} {
		work2, err := session.Work(ctxB, model.ProcessUnitTest1, constantRetriever)
		if !errors.Is(err, ErrWorkConflict) {
			t.Errorf("Concurrent work with same process should've been forbidden, got %v", err)
			work2.Cancel(ctxB)
		}
	}

	// Now, finish the long-running work.
	work1.Cancel(ctx)

	// We should now be able to perform 'test1' work again against this machine.
	for _, session := range []*Session{session1, session2} {
		work1, err := session.Work(ctxB, model.ProcessUnitTest1, constantRetriever)
		if err != nil {
			t.Errorf("Could not run work against machine: %v", err)
		} else {
			work1.Cancel(ctxB)
		}
	}
}

func TestWorkBackoff(t *testing.T) {
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

	var machine model.Machine
	// Create machine.
	err = session.Transact(ctx, func(q *model.Queries) error {
		machine, err = q.NewMachine(ctx)
		if err != nil {
			return err
		}
		return q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
			MachineID:  machine.MachineID,
			Provider:   model.ProviderEquinix,
			ProviderID: "123",
		})
	})
	if err != nil {
		t.Fatalf("Creating machine failed: %v", err)
	}

	// Work on machine, but fail it with a backoff.
	work, err := session.Work(ctx, model.ProcessShepherdAccess, func(q *model.Queries) ([]uuid.UUID, error) {
		machines, err := q.GetMachinesForAgentStart(ctx, 1)
		if err != nil {
			return nil, err
		}
		if len(machines) < 1 {
			return nil, ErrNothingToDo
		}
		return []uuid.UUID{machines[0].MachineID}, nil
	})
	if err != nil {
		t.Fatalf("Starting work failed: %v", err)
	}
	backoff := time.Hour
	if err := work.Fail(ctx, &backoff, "test"); err != nil {
		t.Fatalf("Failing work failed: %v", err)
	}

	// The machine shouldn't be returned now.
	err = session.Transact(ctx, func(q *model.Queries) error {
		machines, err := q.GetMachinesForAgentStart(ctx, 1)
		if err != nil {
			return err
		}
		if len(machines) > 0 {
			t.Errorf("Expected no machines ready for agent start.")
		}
		return nil
	})
	if err != nil {
		t.Errorf("Failed to retrieve machines for agent start: %v", err)
	}

	// Instead of waiting for the backoff to expire, set it again, but this time
	// make it immediate. This works because the backoff query acts as an upsert.
	err = session.Transact(ctx, func(q *model.Queries) error {
		return q.WorkBackoffInsert(ctx, model.WorkBackoffInsertParams{
			MachineID: machine.MachineID,
			Process:   model.ProcessShepherdAccess,
			Seconds:   0,
		})
	})
	if err != nil {
		t.Errorf("Failed to update backoff: %v", err)
	}

	// Just in case.
	time.Sleep(100 * time.Millisecond)

	// The machine should now be returned again.
	err = session.Transact(ctx, func(q *model.Queries) error {
		machines, err := q.GetMachinesForAgentStart(ctx, 1)
		if err != nil {
			return err
		}
		if len(machines) != 1 {
			t.Errorf("Expected exactly one machine ready for agent start.")
		}
		return nil
	})
	if err != nil {
		t.Errorf("Failed to retrieve machines for agent start: %v", err)
	}
}

// TestAgentStartWorkflow exercises the agent start workflow within the BMDB.
func TestAgentStartWorkflow(t *testing.T) {
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

	// Create machine. Drop its ID.
	err = session.Transact(ctx, func(q *model.Queries) error {
		machine, err := q.NewMachine(ctx)
		if err != nil {
			return err
		}
		return q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
			MachineID:  machine.MachineID,
			Provider:   model.ProviderEquinix,
			ProviderID: "123",
		})
	})
	if err != nil {
		t.Fatalf("Creating machine failed: %v", err)
	}

	// Start working on a machine.
	var machine uuid.UUID
	startedC := make(chan struct{})
	doneC := make(chan struct{})
	errC := make(chan error)
	go func() {
		work, err := session.Work(ctx, model.ProcessShepherdAccess, func(q *model.Queries) ([]uuid.UUID, error) {
			machines, err := q.GetMachinesForAgentStart(ctx, 1)
			if err != nil {
				return nil, err
			}
			if len(machines) < 1 {
				return nil, ErrNothingToDo
			}
			machine = machines[0].MachineID
			return []uuid.UUID{machines[0].MachineID}, nil
		})
		defer work.Cancel(ctx)

		if err != nil {
			close(startedC)
			errC <- err
			return
		}

		// Simulate work by blocking on a channel.
		close(startedC)

		<-doneC

		err = work.Finish(ctx, func(q *model.Queries) error {
			return q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
				MachineID:      work.Machine,
				AgentStartedAt: time.Now(),
				AgentPublicKey: []byte("fakefakefake"),
			})
		})
		errC <- err
	}()
	<-startedC
	// Work on the machine has started. Attempting to get more machines now should
	// return no machines.
	err = session.Transact(ctx, func(q *model.Queries) error {
		machines, err := q.GetMachinesForAgentStart(ctx, 1)
		if err != nil {
			return err
		}
		if len(machines) > 0 {
			t.Errorf("Expected no machines ready for agent start.")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Failed to retrieve machines for start in parallel: %v", err)
	}
	// Finish working on machine.
	close(doneC)
	err = <-errC
	if err != nil {
		t.Fatalf("Failed to finish work on machine: %v", err)
	}
	// That machine has its agent started, so we still expect no work to have to be
	// done.
	err = session.Transact(ctx, func(q *model.Queries) error {
		machines, err := q.GetMachinesForAgentStart(ctx, 1)
		if err != nil {
			return err
		}
		if len(machines) > 0 {
			t.Errorf("Expected still no machines ready for agent start.")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Failed to retrieve machines for agent start after work finished: %v", err)
	}

	// Check history has been recorded.
	var history []model.WorkHistory
	err = session.Transact(ctx, func(q *model.Queries) error {
		history, err = q.ListHistoryOf(ctx, machine)
		return err
	})
	if err != nil {
		t.Fatalf("Failed to retrieve machine history: %v", err)
	}
	// Expect two history items: started and finished.
	if want, got := 2, len(history); want != got {
		t.Errorf("Wanted %d history items, got %d", want, got)
	} else {
		if want, got := model.WorkHistoryEventStarted, history[0].Event; want != got {
			t.Errorf("Wanted first history event to be %s, got %s", want, got)
		}
		if want, got := model.WorkHistoryEventFinished, history[1].Event; want != got {
			t.Errorf("Wanted second history event to be %s, got %s", want, got)
		}
	}
	// Check all other history event data.
	for i, el := range history {
		if want, got := machine, el.MachineID; want.String() != got.String() {
			t.Errorf("Wanted %d history event machine ID to be %s, got %s", i, want, got)
		}
		if want, got := model.ProcessShepherdAccess, el.Process; want != got {
			t.Errorf("Wanted %d history event process to be %s, got %s", i, want, got)
		}
	}
}

// TestAgentStartWorkflowParallel starts work on three machines by six workers
// and makes sure that there are no scheduling conflicts between them.
func TestAgentStartWorkflowParallel(t *testing.T) {
	b := dut()
	conn, err := b.Open(true)
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}

	ctx, ctxC := context.WithCancel(context.Background())
	defer ctxC()

	makeMachine := func(providerID string) {
		ctxS, ctxC := context.WithCancel(ctx)
		defer ctxC()
		session, err := conn.StartSession(ctxS)
		if err != nil {
			t.Fatalf("Starting session failed: %v", err)
		}
		err = session.Transact(ctx, func(q *model.Queries) error {
			machine, err := q.NewMachine(ctx)
			if err != nil {
				return err
			}
			return q.MachineAddProvided(ctx, model.MachineAddProvidedParams{
				MachineID:  machine.MachineID,
				Provider:   model.ProviderEquinix,
				ProviderID: providerID,
			})
		})
		if err != nil {
			t.Fatalf("Creating machine failed: %v", err)
		}
	}
	// Make six machines for testing.
	for i := 0; i < 6; i++ {
		makeMachine(fmt.Sprintf("test%d", i))
	}

	workStarted := make(chan struct{})
	workDone := make(chan struct {
		machine  uuid.UUID
		workerID int
	})

	workOnce := func(ctx context.Context, workerID int, session *Session) error {
		work, err := session.Work(ctx, model.ProcessShepherdAccess, func(q *model.Queries) ([]uuid.UUID, error) {
			machines, err := q.GetMachinesForAgentStart(ctx, 1)
			if err != nil {
				return nil, err
			}
			if len(machines) < 1 {
				return nil, ErrNothingToDo
			}
			return []uuid.UUID{machines[0].MachineID}, nil
		})

		if err != nil {
			return err
		}
		defer work.Cancel(ctx)

		select {
		case <-workStarted:
		case <-ctx.Done():
			return ctx.Err()
		}

		select {
		case workDone <- struct {
			machine  uuid.UUID
			workerID int
		}{
			machine:  work.Machine,
			workerID: workerID,
		}:
		case <-ctx.Done():
			return ctx.Err()
		}

		return work.Finish(ctx, func(q *model.Queries) error {
			return q.MachineSetAgentStarted(ctx, model.MachineSetAgentStartedParams{
				MachineID:      work.Machine,
				AgentStartedAt: time.Now(),
				AgentPublicKey: []byte("fakefakefake"),
			})
		})
	}

	worker := func(workerID int) {
		ctxS, ctxC := context.WithCancel(ctx)
		defer ctxC()
		session, err := conn.StartSession(ctxS)
		if err != nil {
			t.Fatalf("Starting session failed: %v", err)
		}
		for {
			err := workOnce(ctxS, workerID, session)
			if err != nil {
				if errors.Is(err, ctxS.Err()) {
					return
				}
				t.Fatalf("worker failed: %v", err)
			}
		}
	}
	// Start three workers.
	for i := 0; i < 3; i++ {
		go worker(i)
	}

	// Wait for at least three workers to be alive.
	for i := 0; i < 3; i++ {
		workStarted <- struct{}{}
	}

	// Allow all workers to continue running from now on.
	close(workStarted)

	// Expect six machines to have been handled in parallel by three workers.
	seenWorkers := make(map[int]bool)
	seenMachines := make(map[string]bool)
	for i := 0; i < 6; i++ {
		res := <-workDone
		seenWorkers[res.workerID] = true
		seenMachines[res.machine.String()] = true
	}

	if want, got := 3, len(seenWorkers); want != got {
		t.Errorf("Expected %d workers, got %d", want, got)
	}
	if want, got := 6, len(seenMachines); want != got {
		t.Errorf("Expected %d machines, got %d", want, got)
	}
}
