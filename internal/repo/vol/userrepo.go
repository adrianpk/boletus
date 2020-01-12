package pg

import (
	"errors"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/foundation/db"
	"golang.org/x/crypto/bcrypt"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	UserRepo struct {
		Cfg  *fnd.Config
		Log  fnd.Logger
		Name string
	}
)

type (
	userRow struct {
		mutable bool
		model   model.User
	}
)

var (
	superadmin = model.User{
		Identification: fnd.Identification{
			ID:   fnd.ToUUID("e8b43223-17fe-4e36-bd0f-a7d96e867d95"),
			Slug: db.ToNullString("superadmin-aa0298fe796f"),
		},
		Username:       db.ToNullString("superadmin"),
		PasswordDigest: db.ToNullString(digestPass("superadmin")),
		Email:          db.ToNullString("superadmin@boletus.club"),
		IsConfirmed:    db.ToNullBool(true),
	}

	admin = model.User{
		Identification: fnd.Identification{
			ID:   fnd.ToUUID("fc86c00c-2d4f-400b-ae57-d9d5c87d13c8"),
			Slug: db.ToNullString("admin-7463efd308b4"),
		},
		Username:       db.ToNullString("admin"),
		PasswordDigest: db.ToNullString(digestPass("admin")),
		Email:          db.ToNullString("admin@boletus.club"),
		IsConfirmed:    db.ToNullBool(true),
	}

	usersTable = map[uuid.UUID]userRow{
		superadmin.ID: userRow{mutable: false, model: superadmin},
		admin.ID:      userRow{mutable: false, model: admin},
	}
)

func (ur *UserRepo) Create(user *model.User, tx ...*sqlx.Tx) error {
	_, ok := usersTable[user.ID]
	if ok {
		errors.New("duplicate key violates unique constraint")
	}

	if user.ID == uuid.Nil {
		errors.New("Non valid primary key")
	}

	usersTable[user.ID] = userRow{
		mutable: true,
		model:   *user,
	}

	return nil
}

func (ur *UserRepo) GetAll() (users []model.User, err error) {
	size := len(usersTable)
	out := make([]model.User, size)
	for _, row := range usersTable {
		out = append(out, row.model)
	}
	return out, nil
}

func (ur *UserRepo) Get(id uuid.UUID) (user model.User, err error) {
	for _, row := range usersTable {
		if id == row.model.ID {
			return row.model, nil
		}
	}
	return model.User{}, nil
}

func (ur *UserRepo) GetBySlug(slug string) (user model.User, err error) {
	for _, row := range usersTable {
		if slug == row.model.Slug.String {
			return row.model, nil
		}
	}
	return model.User{}, nil
}

func (ur *UserRepo) GetByUsername(name string) (model.User, error) {
	for _, row := range usersTable {
		if name == row.model.Username.String {
			return row.model, nil
		}
	}
	return model.User{}, nil
}

func (ur *UserRepo) Update(user *model.User, tx ...*sqlx.Tx) error {
	for _, row := range usersTable {
		if user.ID == row.model.ID {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			usersTable[user.ID] = userRow{
				mutable: true,
				model:   *user,
			}
			return nil
		}
	}
	return errors.New("no records updated")
}

func (ur *UserRepo) Delete(id uuid.UUID, tx ...*sqlx.Tx) error {
	for _, row := range usersTable {
		if id == row.model.ID {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			delete(usersTable, id)
			return nil
		}
	}
	return errors.New("no records deleted")
}

func (ur *UserRepo) DeleteySlug(slug string, tx ...*sqlx.Tx) error {
	for _, row := range usersTable {
		if slug == row.model.Slug.String {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			delete(usersTable, row.model.ID)
			return nil
		}
	}
	return errors.New("no records deleted")
}

func (ur *UserRepo) ConfirmUser(slug, token string, tx ...*sqlx.Tx) (err error) {
	for _, row := range usersTable {
		if slug == row.model.Slug.String {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			user := row.model
			user.IsConfirmed = db.ToNullBool(true)

			usersTable[user.ID] = userRow{
				mutable: true,
				model:   user,
			}
			return nil
		}
	}
	return errors.New("no records updated")
}

// SignIn user
func (ur *UserRepo) SignIn(username, password string) (model.User, error) {
	for _, row := range usersTable {
		if username == row.model.Username.String {
			u := row.model

			// Validate password
			err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest.String), []byte(password))
			if err != nil {
				return u, err
			}
		}
	}
	return model.User{}, nil
}

// Misc.

func digestPass(pass string) string {
	hpass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hpass)
}
