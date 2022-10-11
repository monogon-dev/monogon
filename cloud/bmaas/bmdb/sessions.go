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
		time.Sleep(s.interval / 2)
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

// Work runs a given function as a work item with a given process name against
// some identified machine. Not more than one process of a given name can run
// against a machine concurrently.
//
// Most impure (meaning with side effects outside the database itself) BMDB
// transactions should be run this way.
func (s *Session) Work(ctx context.Context, machine uuid.UUID, process model.Process, fn func() error) error {
	err := model.New(s.connection.db).StartWork(ctx, model.StartWorkParams{
		MachineID: machine,
		SessionID: s.UUID,
		Process:   process,
	})
	if err != nil {
		var perr *pq.Error
		if errors.As(err, &perr) && perr.Code == "23505" {
			return ErrWorkConflict
		}
		return fmt.Errorf("could not start work: %w", err)
	}
	klog.Infof("Started work: %q on machine %s, session %s", process, machine, s.UUID)

	defer func() {
		err := model.New(s.connection.db).FinishWork(s.ctx, model.FinishWorkParams{
			MachineID: machine,
			SessionID: s.UUID,
			Process:   process,
		})
		klog.Errorf("Finished work: %q on machine %s, session %s", process, machine, s.UUID)
		if err != nil && !errors.Is(err, s.ctx.Err()) {
			klog.Errorf("Failed to finish work: %v", err)
			klog.Errorf("Closing session out of an abundance of caution")
			s.ctxC()
		}
	}()

	return fn()
}
