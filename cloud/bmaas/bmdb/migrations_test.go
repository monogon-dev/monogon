// Copyright The Monogon Project Authors.
// SPDX-License-Identifier: Apache-2.0

package bmdb

import (
	"context"
	"testing"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
)

// TestMigrateUpDown performs a full-up and full-down migration test on an
// in-memory database twice.
//
// Doing this the first time allows us to check the up migrations are valid and
// that the down migrations clean up enough after themselves for earlier down
// migrations to success.
//
// Doing this the second time allows us to make sure the down migrations cleaned
// up enough after themselves that they have left no table/type behind.
func TestMigrateUpDown(t *testing.T) {
	// Start with an empty database.
	b := dut()
	_, err := b.Open(false)
	if err != nil {
		t.Fatalf("Starting empty database failed: %v", err)
	}

	// Migrations go up.
	if err := b.Database.MigrateUp(); err != nil {
		t.Fatalf("Initial up migration failed: %v", err)
	}
	// Migrations go down.
	if err := b.Database.MigrateDownDangerDanger(); err != nil {
		t.Fatalf("Initial down migration failed: %v", err)
	}
	// Migrations go up.
	if err := b.Database.MigrateUp(); err != nil {
		t.Fatalf("Second up migration failed: %v", err)
	}
	// Migrations go down.
	if err := b.Database.MigrateDownDangerDanger(); err != nil {
		t.Fatalf("Second down migration failed: %v", err)
	}
}

// TestMigrateTwice makes sure we don't hit https://review.monogon.dev/1502 again.
func TestMigrateTwice(t *testing.T) {
	// Start with an empty database.
	b := dut()
	_, err := b.Open(false)
	if err != nil {
		t.Fatalf("Starting empty database failed: %v", err)
	}

	// Migrations go up.
	if err := b.Database.MigrateUp(); err != nil {
		t.Fatalf("Initial up migration failed: %v", err)
	}
	// Migrations go up again.
	if err := b.Database.MigrateUp(); err != nil {
		t.Fatalf("Initial up migration failed: %v", err)
	}
}

func TestMigration1681826233(t *testing.T) {
	// This migration adds a new nullable field to backoffs.

	// This guarantees that versions [prev, ver] can run concurrently in a cluster.
	minVer := uint(1672749980)
	maxVer := uint(1681826233)

	ctx, ctxC := context.WithCancel(context.Background())
	defer t.Cleanup(ctxC)

	b := dut()
	conn, err := b.Open(false)
	if err != nil {
		t.Fatalf("Starting empty database failed: %v", err)
	}

	// First, make sure the change can actually progress if we have some backoffs
	// already.
	if err := b.Database.MigrateUpToIncluding(minVer); err != nil {
		t.Fatalf("Migration to minimum version failed: %v", err)
	}

	// Create machine and old-style backoff.
	q := model.New(conn.db)
	machine, err := q.NewMachine(ctx)
	if err != nil {
		t.Fatalf("Could not create machine: %v", err)
	}
	_, err = conn.db.Exec(`
		INSERT INTO work_backoff
		    (machine_id, process, until, cause)
		VALUES
		    ($1, 'UnitTest1', now(), 'test');
	`, machine.MachineID)
	if err != nil {
		t.Fatalf("Could not create old-style backoff on old version: %v", err)
	}

	// Migrate to newer version.
	if err := b.Database.MigrateUpToIncluding(maxVer); err != nil {
		t.Fatalf("Migration to maximum version failed: %v", err)
	}

	// The migration should be read succesfully.
	boffs, err := q.WorkBackoffOf(ctx, model.WorkBackoffOfParams{
		MachineID: machine.MachineID,
		Process:   "UnitTest1",
	})
	if err != nil {
		t.Fatalf("Reading backoff failed: %v", err)
	}
	if len(boffs) != 1 {
		t.Errorf("No backoff found")
	} else {
		boff := boffs[0]
		if boff.LastIntervalSeconds.Valid {
			t.Errorf("Expected interval to be NULL")
		}
	}

	// Simultaneously, any concurrently running bmdb user on an older version should
	// still be able to insert and read backoffs old style.
	_, err = conn.db.Exec(`
		INSERT INTO work_backoff
		    (machine_id, process, until, cause)
		VALUES
		    ($1, 'UnitTest2', now(), 'test');
	`, machine.MachineID)
	if err != nil {
		t.Fatalf("Could not create old-style backoff on new version: %v", err)
	}
	rows, err := conn.db.Query(`
		SELECT machine_id, process, until, cause FROM work_backoff
	`)
	if err != nil {
		t.Fatalf("Could not fetch old-style backoff data: %v", err)
	}
	for rows.Next() {
		var mid, process, until, cause string
		if err := rows.Scan(&mid, &process, &until, &cause); err != nil {
			t.Errorf("Scan failed: %v", err)
		}
	}
}
