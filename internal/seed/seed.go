package seed

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // package init.
	fnd "github.com/adrianpk/foundation"
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
	// Create users
	st := &step{}
	st.Config(st.CreateUsers)
	s.AddSeed(st)
}
