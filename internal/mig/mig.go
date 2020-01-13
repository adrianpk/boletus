package mig

import (
	fnd "github.com/adrianpk/foundation"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // package init.
)

const (
	devDb  = "foundation"
	testDb = "foundation_test"
	prodDb = "foundation_prod"
)

type (
	// Migrator is a migrator handler.
	Migrator struct {
		*fnd.Migrator
	}
)

// NewMigrator creates and returns a new migrator.
func NewMigrator(cfg *fnd.Config, log fnd.Logger, name string, db *sqlx.DB) (*Migrator, error) {
	log.Info("New migrator", "name", name)

	m := &Migrator{
		fnd.NewMigrator(cfg, log, name, db),
	}

	m.addSteps()

	return m, nil
}

// GetMigrator configured.
func (m *Migrator) addSteps() {
	// Migrations

	// Create users table
	s := &step{}
	s.Config(s.CreateUsersTable, s.DropUsersTable)
	m.AddMigration(s)

	// Create events table
	s = &step{}
	s.Config(s.CreateEventsTable, s.DropEventsTable)
	m.AddMigration(s)

	// Create tickets table
	s = &step{}
	s.Config(s.CreateTicketsTable, s.DropTicketsTable)
	m.AddMigration(s)
}
