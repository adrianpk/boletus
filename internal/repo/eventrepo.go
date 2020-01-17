package repo

import (
	"github.com/adrianpk/boletus/internal/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	EventRepo interface {
		Create(u *model.Event, tx ...*sqlx.Tx) error
		GetAll() (users []model.Event, err error)
		Get(id uuid.UUID) (user model.Event, err error)
		GetBySlug(slug string) (user model.Event, err error)
		GetByName(username string) (model.Event, error)
		Update(user *model.Event, tx ...*sqlx.Tx) error
		Delete(id uuid.UUID, tx ...*sqlx.Tx) error
		DeleteBySlug(slug string, tx ...*sqlx.Tx) error
	}
)
