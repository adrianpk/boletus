package seed

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
	// Seeder is a seeder handler.
	Seeder struct {
		*fnd.Seeder
	}
)

// NewSeeder creates and returns a new seeder.
func NewSeeder(cfg *fnd.Config, log fnd.Logger, name string, db *sqlx.DB) (*Seeder, error) {
	log.Info("New seeder", "name", name)

	m := &Seeder{
		fnd.NewSeeder(cfg, log, name, db),
	}

	m.addSteps()

	return m, nil
}

// GetSeeder configured.
func (s *Seeder) addSteps() {
	// Seeds
	// Users
	st := &step{}
	st.Config(st.Users)
	s.AddSeed(st)

	// Events and tickets
	st = &step{}
	st.Config(st.EventAndTickets)
	s.AddSeed(st)
}
