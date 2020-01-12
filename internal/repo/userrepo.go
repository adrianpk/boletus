package repo

import (
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/adrianpk/boletus/internal/model"
)

type (
	UserRepo interface {
		Create(u *model.User, tx ...*sqlx.Tx) error
		GetAll() (users []model.User, err error)
		Get(id uuid.UUID) (user model.User, err error)
		GetBySlug(slug string) (user model.User, err error)
		GetByUsername(username string) (model.User, error)
		Update(user *model.User, tx ...*sqlx.Tx) error
		Delete(id uuid.UUID, tx ...*sqlx.Tx) error
		DeleteBySlug(slug string, tx ...*sqlx.Tx) error
		DeleteByUsername(username string, tx ...*sqlx.Tx) error
		GetBySlugAndToken(slug, token string) (model.User, error)
		ConfirmUser(slug, token string, tx ...*sqlx.Tx) (err error)
		SignIn(username, password string) (model.User, error)
	}
)
