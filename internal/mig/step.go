package mig

import (
	fnd "github.com/adrianpk/foundation"
	"github.com/jmoiron/sqlx"
)

type (
	step struct {
		name string
		up   fnd.MigFx
		down fnd.MigFx
		tx   *sqlx.Tx
	}
)

func (s *step) Config(up fnd.MigFx, down fnd.MigFx) {
	s.up = up
	s.down = down
}

func (s *step) GetName() (name string) {
	return s.name
}

func (s *step) GetUp() (up fnd.MigFx) {
	return s.up
}

func (s *step) GetDown() (down fnd.MigFx) {
	return s.down
}

func (s *step) SetTx(tx *sqlx.Tx) {
	s.tx = tx
}

func (s *step) GetTx() (tx *sqlx.Tx) {
	return s.tx
}
