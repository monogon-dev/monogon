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
	"source.monogon.dev/cloud/lib/component"
)

// BMDB is the Bare Metal Database, a common schema to store information about
// bare metal machines in CockroachDB. This struct is supposed to be
// embedded/contained by different components that interact with the BMDB, and
// provides a common interface to BMDB operations to these components.
//
// The BMDB provides two mechanisms facilitating a 'reactive work system' being
// implemented on the bare metal machine data:
//
//   - Sessions, which are maintained by heartbeats by components and signal the
//     liveness of said components to other components operating on the BMDB. These
//     effectively extend CockroachDB's transactions to be visible as row data. Any
//     session that is not actively being updated by a component can be expired by a
//     component responsible for lease garbage collection.
//   - Work locking, which bases on Sessions and allows long-standing
//     multi-transaction work to be performed on given machines, preventing
//     conflicting work from being performed by other components. As both Work
//     locking and Sessions are plain row data, other components can use SQL queries
//     to exclude machines to act on by constraining SELECT queries to not return
//     machines with some active work being performed on them.
type BMDB struct {
	Config
}

// Config is the configuration of the BMDB connector.
type Config struct {
	Database component.CockroachConfig

	// ComponentName is a human-readable name of the component connecting to the
	// BMDB, and is stored in any Sessions managed by this component's connector.
	ComponentName string
	// RuntimeInfo is a human-readable 'runtime information' (eg. software version,
	// host machine/job information, IP address, etc.) stored alongside the
	// ComponentName in active Sessions.
	RuntimeInfo string
}

// Open creates a new Connection to the BMDB for the calling component. Multiple
// connections can be opened (although there is no advantage to doing so, as
// Connections manage an underlying CockroachDB connection pool, which performs
// required reconnects and connection pooling automatically).
func (b *BMDB) Open(migrate bool) (*Connection, error) {
	if migrate {
		if b.Config.Database.Migrations == nil {
			klog.Infof("Using default migrations source.")
			m, err := model.MigrationsSource()
			if err != nil {
				klog.Exitf("failed to prepare migrations source: %w", err)
			}
			b.Config.Database.Migrations = m
		}
		if err := b.Database.MigrateUp(); err != nil {
			return nil, fmt.Errorf("migration failed: %w", err)
		}
	}
	db, err := b.Database.Connect()
	if err != nil {
		return nil, err
	}
	return &Connection{
		bmdb: b,
		db:   db,
	}, nil
}

// Connection to the BMDB. Internally, this contains a sql.DB connection pool,
// so components can (and likely should) reuse Connections as much as possible
// internally.
type Connection struct {
	bmdb *BMDB
	db   *sql.DB
}

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
				return ErrSessionExpired
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
			return ErrSessionExpired
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
		return q.FinishWork(ctx, model.FinishWorkParams{
			MachineID: w.Machine,
			SessionID: w.s.UUID,
			Process:   w.process,
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
		return fn(q)
	})
}
