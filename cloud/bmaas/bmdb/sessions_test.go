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

// TestWorkBackoff exercises the backoff functionality within the BMDB.
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

	waitMachine := func(nsec int64) *Work {
		t.Helper()

		deadline := time.Now().Add(time.Duration(nsec) * 2 * time.Second)
		for {
			if time.Now().After(deadline) {
				t.Fatalf("Deadline expired")
			}
			work, err := session.Work(ctx, model.ProcessShepherdAgentStart, func(q *model.Queries) ([]uuid.UUID, error) {
				machines, err := q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
					Limit:    1,
					Provider: model.ProviderEquinix,
				})
				if err != nil {
					return nil, err
				}
				if len(machines) < 1 {
					return nil, ErrNothingToDo
				}
				return []uuid.UUID{machines[0].MachineID}, nil
			})
			if err == nil {
				return work
			}
			if !errors.Is(err, ErrNothingToDo) {
				t.Fatalf("Unexpected work error: %v", err)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Work on machine, but fail it with a backoff.
	work := waitMachine(1)
	backoff := Backoff{
		Initial:  time.Second,
		Maximum:  5 * time.Second,
		Exponent: 2,
	}
	if err := work.Fail(ctx, &backoff, "test"); err != nil {
		t.Fatalf("Failing work failed: %v", err)
	}

	expect := func(count int) {
		t.Helper()

		var machines []model.MachineProvided
		var err error
		err = session.Transact(ctx, func(q *model.Queries) error {
			machines, err = q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
				Limit:    1,
				Provider: model.ProviderEquinix,
			})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			t.Errorf("Failed to retrieve machines for agent start: %v", err)
		}
		if want, got := count, len(machines); want != got {
			t.Errorf("Expected %d machines, got %d", want, got)
		}
	}

	// The machine shouldn't be returned now.
	expect(0)

	// Wait for the backoff to expire.
	time.Sleep(1100 * time.Millisecond)

	// The machine should now be returned again.
	expect(1)

	// Prepare helper for checking exponential backoffs.
	failAndCheck := func(nsec int64) {
		t.Helper()
		work := waitMachine(nsec)
		if err := work.Fail(ctx, &backoff, "test"); err != nil {
			t.Fatalf("Failing work failed: %v", err)
		}

		var backoffs []model.WorkBackoff
		err = session.Transact(ctx, func(q *model.Queries) error {
			var err error
			backoffs, err = q.WorkBackoffOf(ctx, model.WorkBackoffOfParams{
				MachineID: machine.MachineID,
				Process:   model.ProcessShepherdAgentStart,
			})
			return err
		})
		if err != nil {
			t.Errorf("Failed to retrieve machines for agent start: %v", err)
		}
		if len(backoffs) < 1 {
			t.Errorf("No backoff")
		} else {
			backoff := backoffs[0]
			if want, got := nsec, backoff.LastIntervalSeconds.Int64; want != got {
				t.Fatalf("Wanted backoff of %d seconds, got %d", want, got)
			}
		}
	}

	// Exercise exponential backoff functionality.
	failAndCheck(2)
	failAndCheck(4)
	failAndCheck(5)
	failAndCheck(5)

	// If the job now succeeds, subsequent failures should start from 1 again.
	work = waitMachine(5)
	err = work.Finish(ctx, func(q *model.Queries) error {
		// Not setting any tags that would cause subsequent queries to not return the
		// machine anymore.
		return nil
	})
	if err != nil {
		t.Fatalf("Could not finish work: %v", err)
	}

	failAndCheck(1)
	failAndCheck(2)
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
		work, err := session.Work(ctx, model.ProcessShepherdAgentStart, func(q *model.Queries) ([]uuid.UUID, error) {
			machines, err := q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
				Limit:    1,
				Provider: model.ProviderEquinix,
			})
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

	// Mutual exclusion with AgentStart:
	err = session.Transact(ctx, func(q *model.Queries) error {
		machines, err := q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
			Limit:    1,
			Provider: model.ProviderEquinix,
		})
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

	// Mutual exclusion with Recovery:
	err = session.Transact(ctx, func(q *model.Queries) error {
		machines, err := q.GetMachineForAgentRecovery(ctx, model.GetMachineForAgentRecoveryParams{
			Limit:    1,
			Provider: model.ProviderEquinix,
		})
		if err != nil {
			return err
		}
		if len(machines) > 0 {
			t.Errorf("Expected no machines ready for agent recovery.")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Failed to retrieve machines for recovery in parallel: %v", err)
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
		machines, err := q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
			Limit:    1,
			Provider: model.ProviderEquinix,
		})
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
		if want, got := model.ProcessShepherdAgentStart, el.Process; want != got {
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
		work, err := session.Work(ctx, model.ProcessShepherdAgentStart, func(q *model.Queries) ([]uuid.UUID, error) {
			machines, err := q.GetMachinesForAgentStart(ctx, model.GetMachinesForAgentStartParams{
				Limit:    1,
				Provider: model.ProviderEquinix,
			})
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
		ctxS, ctxSC := context.WithCancel(ctx)
		defer ctxSC()
		session, err := conn.StartSession(ctxS)
		if err != nil {
			t.Errorf("Starting session failed: %v", err)
			ctxC()
			return
		}
		for {
			err := workOnce(ctxS, workerID, session)
			if err != nil {
				if errors.Is(err, ErrNothingToDo) {
					continue
				}
				if errors.Is(err, ctxS.Err()) {
					return
				}
				t.Errorf("worker failed: %v", err)
				ctxC()
				return
			}
		}
	}
	// Start three workers.
	for i := 0; i < 3; i++ {
		go worker(i)
	}

	// Wait for at least three workers to be alive.
	for i := 0; i < 3; i++ {
		select {
		case workStarted <- struct{}{}:
		case <-ctx.Done():
			t.FailNow()
		}
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
