package bmdb

import (
	"testing"
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
