package model

import (
	"embed"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrationData embed.FS

func MigrationsSource() (source.Driver, error) {
	return iofs.New(migrationData, "migrations")
}
