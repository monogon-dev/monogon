package bmdb

import (
	"database/sql"
	"fmt"

	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
)

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
