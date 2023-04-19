package bmdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
)

// StartSession creates a new BMDB session which will be maintained in a
// background goroutine as long as the given context is valid. Each Session is
// represented by an entry in a sessions table within the BMDB, and subsequent
// Transact calls emit SQL transactions which depend on that entry still being
// present and up to date. A garbage collection system (to be implemented) will
// remove expired sessions from the BMDB, but this mechanism is not necessary
// for the session expiry mechanism to work.
//
// When the session becomes invalid (for example due to network partition),
// subsequent attempts to call Transact will fail with ErrSessionExpired. This
// means that the caller within the component is responsible for recreating a
// new Session if a previously used one expires.
func (c *Connection) StartSession(ctx context.Context) (*Session, error) {
	intervalSeconds := 5

	res, err := model.New(c.db).NewSession(ctx, model.NewSessionParams{
		SessionComponentName:   c.bmdb.ComponentName,
		SessionRuntimeInfo:     c.bmdb.RuntimeInfo,
		SessionIntervalSeconds: int64(intervalSeconds),
	})
	if err != nil {
		return nil, fmt.Errorf("creating session failed: %w", err)
	}

	klog.Infof("Started session %s", res.SessionID)

	ctx2, ctxC := context.WithCancel(ctx)

	s := &Session{
		connection: c,
		interval:   time.Duration(intervalSeconds) * time.Second,

		UUID: res.SessionID,

		ctx:  ctx2,
		ctxC: ctxC,
	}
	go s.maintainHeartbeat(ctx2)
	return s, nil
}

// Session is a session (identified by UUID) that has been started in the BMDB.
// Its liveness is maintained by a background goroutine, and as long as that
// session is alive, it can perform transactions and work on the BMDB.
type Session struct {
	connection *Connection
	interval   time.Duration

	UUID uuid.UUID

	ctx  context.Context
	ctxC context.CancelFunc
}

// Expired returns true if this session is expired and will fail all subsequent
// transactions/work.
func (s *Session) Expired() bool {
	return s.ctx.Err() != nil
}

// expire is a helper which marks this session as expired and returns
// ErrSessionExpired.
func (s *Session) expire() error {
	s.ctxC()
	return ErrSessionExpired
}

var (
	// ErrSessionExpired is returned when attempting to Transact or Work on a
	// Session that has expired or been canceled. Once a Session starts returning
	// these errors, it must be re-created by another StartSession call, as no other
	// calls will succeed.
	ErrSessionExpired = errors.New("session expired")
	// ErrWorkConflict is returned when attempting to Work on a Session with a
	// process name that's already performing some work, concurrently, on the
	// requested machine.
	ErrWorkConflict = errors.New("conflicting work on machine")
)

// maintainHeartbeat will attempt to repeatedly poke the session at a frequency
// twice of that of the minimum frequency mandated by the configured 5-second
// interval. It will exit if it detects that the session cannot be maintained
// anymore, canceling the session's internal context and causing future
// Transact/Work calls to fail.
func (s *Session) maintainHeartbeat(ctx context.Context) {
	// Internal deadline, used to check whether we haven't dropped the ball on
	// performing the updates due to a lot of transient errors.
	deadline := time.Now().Add(s.interval)
	for {
		if ctx.Err() != nil {
			klog.Infof("Session %s: context over, exiting: %v", s.UUID, ctx.Err())
			return
		}

		err := s.Transact(ctx, func(q *model.Queries) error {
			sessions, err := q.SessionCheck(ctx, s.UUID)
			if err != nil {
				return fmt.Errorf("when retrieving session: %w", err)
			}
			if len(sessions) < 1 {
				return s.expire()
			}
			err = q.SessionPoke(ctx, s.UUID)
			if err != nil {
				return fmt.Errorf("when poking session: %w", err)
			}
			return nil
		})
		if err != nil {
			klog.Errorf("Session %s: update failed: %v", s.UUID, err)
			if errors.Is(err, ErrSessionExpired) || deadline.After(time.Now()) {
				// No way to recover.
				klog.Errorf("Session %s: exiting", s.UUID)
				s.ctxC()
				return
			}
			// Just retry in a bit. One second seems about right for a 5 second interval.
			//
			// TODO(q3k): calculate this based on the configured interval.
			time.Sleep(time.Second)
		}
		// Success. Keep going.
		deadline = time.Now().Add(s.interval)
		select {
		case <-ctx.Done():
			// Do nothing, next loop iteration will exit.
		case <-time.After(s.interval / 2):
			// Do nothing, next loop iteration will heartbeat.
		}
	}
}

// Transact runs a given function in the context of both a CockroachDB and BMDB
// transaction, retrying as necessary.
//
// Most pure (meaning without side effects outside the database itself) BMDB
// transactions should be run this way.
func (s *Session) Transact(ctx context.Context, fn func(q *model.Queries) error) error {
	return crdb.ExecuteTx(ctx, s.connection.db, nil, func(tx *sql.Tx) error {
		qtx := model.New(tx)
		sessions, err := qtx.SessionCheck(ctx, s.UUID)
		if err != nil {
			return fmt.Errorf("when retrieving session: %w", err)
		}
		if len(sessions) < 1 {
			return s.expire()
		}

		if err := fn(qtx); err != nil {
			return err
		}

		return nil
	})
}

var (
	ErrNothingToDo = errors.New("nothing to do")
	// PostgresUniqueViolation is returned by the lib/pq driver when a mutation
	// cannot be performed due to a UNIQUE constraint being violated as a result of
	// the query.
	postgresUniqueViolation = pq.ErrorCode("23505")
)

// Work starts work on a machine. Full work execution is performed in three
// phases:
//
//  1. Retrieval phase. This is performed by 'fn' given to this function.
//     The retrieval function must return zero or more machines that some work
//     should be performed on per the BMDB. The first returned machine will be
//     locked for work under the given process and made available in the Work
//     structure returned by this call. The function may be called multiple times,
//     as it's run within a CockroachDB transaction which may be retried an
//     arbitrary number of times. Thus, it should be side-effect free, ideally only
//     performing read queries to the database.
//  2. Work phase. This is performed by user code while holding on to the Work
//     structure instance.
//  3. Commit phase. This is performed by the function passed to Work.Finish. See
//     that method's documentation for more details.
//
// Important: after retrieving Work successfully, either Finish or Cancel must be
// called, otherwise the machine will be locked until the parent session expires
// or is closed! It's safe and recommended to `defer work.Close()` after calling
// Work().
//
// If no machine is eligible for work, ErrNothingToDo should be returned by the
// retrieval function, and the same error (wrapped) will be returned by Work. In
// case the retrieval function returns no machines and no error, that error will
// also be returned.
//
// The returned Work object is _not_ goroutine safe.
func (s *Session) Work(ctx context.Context, process model.Process, fn func(q *model.Queries) ([]uuid.UUID, error)) (*Work, error) {
	var mid *uuid.UUID
	err := s.Transact(ctx, func(q *model.Queries) error {
		mids, err := fn(q)
		if err != nil {
			return fmt.Errorf("could not retrieve machines for work: %w", err)
		}
		if len(mids) < 1 {
			return ErrNothingToDo
		}
		mid = &mids[0]
		err = q.StartWork(ctx, model.StartWorkParams{
			MachineID: mids[0],
			SessionID: s.UUID,
			Process:   process,
		})
		if err != nil {
			var perr *pq.Error
			if errors.As(err, &perr) && perr.Code == postgresUniqueViolation {
				return ErrWorkConflict
			}
			return fmt.Errorf("could not start work on %q: %w", mids[0], err)
		}
		err = q.WorkHistoryInsert(ctx, model.WorkHistoryInsertParams{
			MachineID: mids[0],
			Event:     model.WorkHistoryEventStarted,
			Process:   process,
		})
		if err != nil {
			return fmt.Errorf("could not insert history event: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	klog.Infof("Started work %q on machine %q (sess %q)", process, *mid, s.UUID)
	return &Work{
		Machine: *mid,
		s:       s,
		process: process,
	}, nil
}

// Work being performed on a machine.
type Work struct {
	// Machine that this work is being performed on, as retrieved by the retrieval
	// function passed to the Work method.
	Machine uuid.UUID
	// s is the parent session.
	s *Session
	// done marks that this work has already been canceled or finished.
	done bool
	// process that this work performs.
	process model.Process
}

// Cancel the Work started on a machine. If the work has already been finished
// or canceled, this is a no-op. In case of error, a log line will be emitted.
func (w *Work) Cancel(ctx context.Context) {
	if w.done {
		return
	}
	w.done = true

	klog.Infof("Canceling work %q on machine %q (sess %q)", w.process, w.Machine, w.s.UUID)
	// Eat error and log. There's nothing we can do if this fails, and if it does, it's
	// probably because our connectivity to the BMDB has failed. If so, our session
	// will be invalidated soon and so will the work being performed on this
	// machine.
	err := w.s.Transact(ctx, func(q *model.Queries) error {
		err := q.FinishWork(ctx, model.FinishWorkParams{
			MachineID: w.Machine,
			SessionID: w.s.UUID,
			Process:   w.process,
		})
		if err != nil {
			return err
		}
		return q.WorkHistoryInsert(ctx, model.WorkHistoryInsertParams{
			MachineID: w.Machine,
			Process:   w.process,
			Event:     model.WorkHistoryEventCanceled,
		})
	})
	if err != nil {
		klog.Errorf("Failed to cancel work %q on %q (sess %q): %v", w.process, w.Machine, w.s.UUID, err)
	}
}

// Finish work by executing a commit function 'fn' and releasing the machine
// from the work performed. The function given should apply tags to the
// processed machine in a way that causes it to not be eligible for retrieval
// again. As with the retriever function, the commit function might be called an
// arbitrary number of times as part of cockroachdb transaction retries.
//
// This may be called only once.
func (w *Work) Finish(ctx context.Context, fn func(q *model.Queries) error) error {
	if w.done {
		return fmt.Errorf("already finished")
	}
	w.done = true
	klog.Infof("Finishing work %q on machine %q (sess %q)", w.process, w.Machine, w.s.UUID)
	return w.s.Transact(ctx, func(q *model.Queries) error {
		err := q.FinishWork(ctx, model.FinishWorkParams{
			MachineID: w.Machine,
			SessionID: w.s.UUID,
			Process:   w.process,
		})
		if err != nil {
			return err
		}
		err = q.WorkHistoryInsert(ctx, model.WorkHistoryInsertParams{
			MachineID: w.Machine,
			Process:   w.process,
			Event:     model.WorkHistoryEventFinished,
		})
		if err != nil {
			return err
		}
		return fn(q)
	})
}

// Fail work and introduce backoff for a given duration (if given backoff is
// non-nil). As long as that backoff is active, no further work for this
// machine/process will be started. The given cause is an operator-readable
// string that will be persisted alongside the backoff and the work history/audit
// table.
func (w *Work) Fail(ctx context.Context, backoff *time.Duration, cause string) error {
	if w.done {
		return fmt.Errorf("already finished")
	}
	w.done = true

	return w.s.Transact(ctx, func(q *model.Queries) error {
		err := q.FinishWork(ctx, model.FinishWorkParams{
			MachineID: w.Machine,
			SessionID: w.s.UUID,
			Process:   w.process,
		})
		if err != nil {
			return err
		}
		err = q.WorkHistoryInsert(ctx, model.WorkHistoryInsertParams{
			MachineID: w.Machine,
			Process:   w.process,
			Event:     model.WorkHistoryEventFailed,
			FailedCause: sql.NullString{
				String: cause,
				Valid:  true,
			},
		})
		if err != nil {
			return err
		}
		if backoff != nil && backoff.Seconds() >= 1.0 {
			seconds := int64(backoff.Seconds())
			klog.Infof("Adding backoff for %q on machine %q (%d seconds)", w.process, w.Machine, seconds)
			return q.WorkBackoffInsert(ctx, model.WorkBackoffInsertParams{
				MachineID: w.Machine,
				Process:   w.process,
				Seconds:   seconds,
				Cause:     cause,
			})
		} else {
			return nil
		}
	})
}
