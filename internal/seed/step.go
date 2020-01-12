package seed

import (
	"github.com/jmoiron/sqlx"
	fnd "github.com/adrianpk/foundation"
)

type (
	step struct {
		name string
		seed fnd.SeedFx
		tx   *sqlx.Tx
	}
)

func (s *step) Config(seed fnd.SeedFx) {
	s.seed = seed
}

func (s *step) GetSeed() (up fnd.MigFx) {
	return s.seed
}

func (s *step) SetTx(tx *sqlx.Tx) {
	s.tx = tx
}

func (s *step) GetTx() (tx *sqlx.Tx) {
	return s.tx
}
