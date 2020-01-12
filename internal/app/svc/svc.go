package svc

import (
	"github.com/adrianpk/boletus/internal/repo"
	fnd "github.com/adrianpk/foundation"
	"github.com/jmoiron/sqlx"
	//repo "github.com/adrianpk/boletus/internal/repo/pg"
)

type (
	Service struct {
		*fnd.Service
		DB        *sqlx.DB
		UserRepo  repo.UserRepo
		EventRepo repo.EventRepo
	}
)

func NewService(cfg *fnd.Config, log fnd.Logger, name string, db *sqlx.DB) *Service {
	return &Service{
		Service: fnd.NewService(cfg, log, name),
		DB:      db,
	}
}

func (s *Service) getTx() (tx *sqlx.Tx, err error) {
	tx, err = s.DB.Beginx()
	if err != nil {
		return tx, err
	}

	return tx, nil
}
