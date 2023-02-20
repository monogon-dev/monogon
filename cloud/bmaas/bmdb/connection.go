package bmdb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/bmaas/bmdb/reflection"
)

// Open creates a new Connection to the BMDB for the calling component. Multiple
// connections can be opened (although there is no advantage to doing so, as
// Connections manage an underlying CockroachDB connection pool, which performs
// required reconnects and connection pooling automatically).
func (b *BMDB) Open(migrate bool) (*Connection, error) {
	if b.Config.Database.Migrations == nil {
		klog.Infof("Using default migrations source.")
		m, err := model.MigrationsSource()
		if err != nil {
			klog.Exitf("failed to prepare migrations source: %w", err)
		}
		b.Config.Database.Migrations = m
	}
	if migrate {
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

		DatabaseName: b.Config.Database.DatabaseName,
		Address:      b.Config.Database.EndpointHost,
		InMemory:     b.Config.Database.InMemory,
	}, nil
}

// Connection to the BMDB. Internally, this contains a sql.DB connection pool,
// so components can (and likely should) reuse Connections as much as possible
// internally.
type Connection struct {
	bmdb *BMDB
	db   *sql.DB

	// The database name that we're connected to.
	DatabaseName string
	// The address of the CockroachDB endpoint we've connected to.
	Address string
	// Whether this connection is to an in-memory database. Note: this only works if
	// this Connection came directly from calling Open on a BMDB that was defined to
	// be in-memory. If you just connect to an in-memory CRDB manually, this will
	// still be false.
	InMemory bool
}

// Reflect returns a reflection.Schema as detected by inspecting the table
// information of this connection to the BMDB. The Schema can then be used to
// retrieve arbitrary tag/machine information without going through the
// concurrency/ordering mechanism of the BMDB.
//
// This should only be used to implement debugging tooling and should absolutely
// not be in the path of any user requests.
//
// This Connection will be used not only to query the Schema information, but
// also for all subsequent data retrieval operations on it. Please ensure that
// the Schema is rebuilt in the event of a database connection failure. Ideally,
// you should be rebuilding the schema often, to follow what is currently
// available on the production database - but not for every request. Use a cache
// or something.
func (c *Connection) Reflect(ctx context.Context) (*reflection.Schema, error) {
	return reflection.Reflect(ctx, c.db)
}

// ListHistoryOf retrieves a full audit history of a machine, sorted
// chronologically. It can be read without a session / transaction for debugging
// purposes.
func (c *Connection) ListHistoryOf(ctx context.Context, machine uuid.UUID) ([]model.WorkHistory, error) {
	return model.New(c.db).ListHistoryOf(ctx, machine)
}

// GetSession retrieves all information about a session. It can be read without a
// session/transaction for debugging purposes.
func (c *Connection) GetSession(ctx context.Context, session uuid.UUID) ([]model.Session, error) {
	return model.New(c.db).GetSession(ctx, session)
}
