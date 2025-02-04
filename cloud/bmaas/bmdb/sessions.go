// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

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

	"source.monogon.dev/cloud/bmaas/bmdb/metrics"
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
func (c *Connection) StartSession(ctx context.Context, opts ...SessionOption) (*Session, error) {
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

	var processor metrics.Processor
	for _, opt := range opts {
		if opt.Processor != "" {
			processor = opt.Processor
		}
	}

	s := &Session{
		connection: c,
		interval:   time.Duration(intervalSeconds) * time.Second,

		UUID: res.SessionID,

		ctx:  ctx2,
		ctxC: ctxC,
		m:    c.bmdb.metrics.Recorder(processor),
	}
	s.m.OnSessionStarted()
	go s.maintainHeartbeat(ctx2)
	return s, nil
}

type SessionOption struct {
	Processor metrics.Processor
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

	m *metrics.ProcessorRecorder
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
	var attempts int64

	err := crdb.ExecuteTx(ctx, s.connection.db, nil, func(tx *sql.Tx) error {
		attempts += 1
		s.m.OnTransactionStarted(attempts)

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
	if err != nil {
		s.m.OnTransactionFailed()
	}
	return err
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
	var exisingingBackoff *existingBackoff
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
		backoffs, err := q.WorkBackoffOf(ctx, model.WorkBackoffOfParams{
			MachineID: mids[0],
			Process:   process,
		})
		if err != nil {
			return fmt.Errorf("could not get backoffs: %w", err)
		}
		if len(backoffs) > 0 {
			// If the backoff exists but the last interval is null (e.g. is from a previous
			// version of the schema when backoffs had no interval data) pretend it doesn't
			// exist. Then the backoff mechanism can restart from a clean slate and populate
			// a new, full backoff row.
			if backoff := backoffs[0]; backoff.LastIntervalSeconds.Valid {
				klog.Infof("Existing backoff: %d seconds", backoff.LastIntervalSeconds.Int64)
				exisingingBackoff = &existingBackoff{
					lastInterval: time.Second * time.Duration(backoff.LastIntervalSeconds.Int64),
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	w := &Work{
		Machine: *mid,
		s:       s,
		process: process,
		backoff: exisingingBackoff,
		m:       s.m.WithProcess(process),
	}
	w.m.OnWorkStarted()
	klog.Infof("Started work %q on machine %q (sess %q)", process, *mid, s.UUID)
	return w, nil
}

// existingBackoff contains backoff information retrieved from a work item that
// has previously failed with a backoff.
type existingBackoff struct {
	// lastInterval is the last interval as stored in the backoff table.
	lastInterval time.Duration
}

// Backoff describes the configuration of backoff for a failed work item. It can
// be passed to Work.Fail to cause an item to not be processed again (to be 'in
// backoff') for a given period of time. Exponential backoff can be configured so
// that subsequent failures of a process will have exponentially increasing
// backoff periods, up to some maximum length.
//
// The underlying unit of backoff period length in the database is one second.
// What that means is that all effective calculated backoff periods must be an
// integer number of seconds. This is performed by always rounding up this period
// to the nearest second. A side effect of this is that with exponential backoff,
// non-integer exponents will be less precisely applied for small backoff values,
// e.g. an exponent of 1.1 with initial backoff of 1s will generate the following
// sequence of backoff periods:
//
// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 15, 17
//
// Which corresponds to the following approximate multipliers in between periods:
//
// 2.00, 1.50, 1.33, 1.25, 1.20, 1.17, 1.14, 1.12, 1.11, 1.10, 1.18, 1.15, 1.13
//
// Thus, the exponent value should be treated more as a limit that the sequence
// of periods will approach than a hard rule for calculating the periods.
// However, if the exponent is larger than 1 (i.e. any time exponential backoff
// is requested), this guarantees that the backoff won't get 'stuck' on a
// repeated period value due to a rounding error.
//
// A zero backoff structure is valid and represents a non-exponential backoff of
// one second.
//
// A partially filled structure is also valid. See the field comments for more
// information about how fields are capped if not set. The described behaviour
// allows for two useful shorthands:
//
//  1. If only Initial is set, then the backoff is non-exponential and will always
//     be of value Initial (or whatever the previous period already persisted the
//     database).
//  2. If only Maximum and Exponent are set, the backoff will be exponential,
//     starting at one second, and exponentially increasing to Maximum.
//
// It is recommended to construct Backoff structures as const values and treat
// them as read-only 'descriptors', one per work kind / process.
//
// One feature currently missing from the Backoff implementation is jitter. This
// might be introduced in the future if deemed necessary.
type Backoff struct {
	// Initial backoff period, used for the backoff if this item failed for the first
	// time (i.e. has not had a Finish call in between two Fail calls).
	//
	// Subsequent calls will ignore this field if the backoff is exponential. If
	// non-exponential, the initial time will always override whatever was previously
	// persisted in the database, i.e. the backoff will always be of value 'Initial'.
	//
	// Cannot be lower than one second. If it is, it will be capped to it.
	Initial time.Duration `u:"initial"`

	// Maximum time for backoff. If the calculation of the next back off period
	// (based on the Exponent and last backoff value) exceeds this maximum, it will
	// be capped to it.
	//
	// Maximum is not persisted in the database. Instead, it is always read from this
	// structure.
	//
	// Cannot be lower than Initial. If it is, it will be capped to it.
	Maximum time.Duration `u:"maximum"`

	// Exponent used for next backoff calculation. Any time a work item fails
	// directly after another failure, the previous backoff period will be multiplied
	// by the exponent to yield the new backoff period. The new period will then be
	// capped to Maximum.
	//
	// Exponent is not persisted in the database. Instead, it is always read from
	// this structure.
	//
	// Cannot be lower than 1.0. If it is, it will be capped to it.
	Exponent float64 `u:"exponent"`
}

// normalized copies the given backoff and returns a 'normalized' version of it,
// with the 'when zero/unset' rules described in the Backoff documentation
// strings.
func (b *Backoff) normalized() *Backoff {
	c := *b

	if c.Exponent < 1.0 {
		c.Exponent = 1.0
	}
	if c.Initial < time.Second {
		c.Initial = time.Second
	}
	if c.Maximum < c.Initial {
		c.Maximum = c.Initial
	}
	return &c
}

func (b *Backoff) simple() bool {
	// Non-normalized simple backoffs will have a zero exponent.
	if b.Exponent == 0.0 {
		return true
	}
	// Normalized simple backoffs will have a 1.0 exponent.
	if b.Exponent == 1.0 {
		return true
	}
	return false
}

// next calculates the backoff period based on a backoff descriptor and previous
// existing backoff information. Both or either can be nil.
func (b *Backoff) next(e *existingBackoff) int64 {
	second := time.Second.Nanoseconds()

	// Minimum interval is one second. Start with that.
	last := second
	// Then, if we have a previous interval, and it's greater than a second, use that
	// as the last interval.
	if e != nil {
		if previous := e.lastInterval.Nanoseconds(); previous > second {
			last = previous
		}
	}

	// If no backoff is configured, go with either the minimum of one second, or
	// whatever the last previous interval was.
	if b == nil {
		return last / second
	}

	// Make a copy of the backoff descriptor, normalizing as necessary.
	c := b.normalized()

	// Simple backoffs always return Initial.
	if b.simple() {
		return c.Initial.Nanoseconds() / second
	}

	// If there is no existing backoff, return the initial backoff value directly.
	if e == nil {
		return c.Initial.Nanoseconds() / second
	}

	// Start out with the persisted interval.
	next := last
	// If by any chance we persisted an interval less than one second, clamp it.
	if next < second {
		next = second
	}

	// Multiply by exponent from descriptor.
	next = int64(float64(next) * c.Exponent)

	// Handle overflows. If multiplying by a positive number resulted in a lower
	// value than what we started with, it means we overflowed and wrapped around. If
	// so, clamp to maximum.
	if next < last {
		next = c.Maximum.Nanoseconds()
	}

	// Clamp to maximum.
	if next > c.Maximum.Nanoseconds() {
		next = c.Maximum.Nanoseconds()
	}
	// Round up to the nearest second.
	if next%second == 0 {
		return next / second
	} else {
		return next/second + 1
	}
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

	backoff *existingBackoff

	m *metrics.ProcessRecorder
}

// Cancel the Work started on a machine. If the work has already been finished
// or canceled, this is a no-op. In case of error, a log line will be emitted.
func (w *Work) Cancel(ctx context.Context) {
	if w.done {
		return
	}
	w.done = true
	w.m.OnWorkFinished(metrics.WorkResultCanceled)

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
	w.m.OnWorkFinished(metrics.WorkResultFinished)

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
		err = q.WorkBackoffDelete(ctx, model.WorkBackoffDeleteParams{
			MachineID: w.Machine,
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

// Fail work and introduce backoff. The given cause is an operator-readable
// string that will be persisted alongside the backoff and the work history/audit
// table.
//
// The backoff describes a period during which the same process will not be
// retried on this machine until its expiration.
//
// The given backoff is a structure which describes both the initial backoff
// period if the work failed for the first time, and a mechanism to exponentially
// increase the backoff period if that work failed repeatedly. The work is
// defined to have failed repeatedly if it only resulted in Cancel/Fail calls
// without any Finish calls in the meantime.
//
// Only the last backoff period is persisted in the database. The exponential
// backoff behaviour (including its maximum time) is always calculated based on
// the given backoff structure.
//
// If nil, the backoff defaults to a non-exponential, one second backoff. This is
// the minimum designed to keep the system chugging along without repeatedly
// trying a failed job in a loop. However, the backoff should generally be set to
// some well engineered value to prevent spurious retries.
func (w *Work) Fail(ctx context.Context, backoff *Backoff, cause string) error {
	if w.done {
		return fmt.Errorf("already finished")
	}
	w.done = true
	w.m.OnWorkFinished(metrics.WorkResultFailed)

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
		if backoff == nil {
			klog.Warningf("Nil backoff for %q on machine %q: defaulting to one second non-exponential.", w.process, w.Machine)
		}
		seconds := backoff.next(w.backoff)
		klog.Infof("Adding backoff for %q on machine %q: %d seconds", w.process, w.Machine, seconds)
		return q.WorkBackoffInsert(ctx, model.WorkBackoffInsertParams{
			MachineID: w.Machine,
			Process:   w.process,
			Cause:     cause,
			Seconds:   seconds,
		})
	})
}
