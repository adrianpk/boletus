package pg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserRepo struct {
		Cfg  *fnd.Config
		Log  fnd.Logger
		Name string
		DB   *sqlx.DB
	}
)

func NewUserRepo(cfg *fnd.Config, log fnd.Logger, name string, db *sqlx.DB) *UserRepo {
	return &UserRepo{
		Cfg:  cfg,
		Log:  log,
		Name: name,
		DB:   db,
	}
}

// Create a user
func (ur *UserRepo) Create(user *model.User, tx ...*sqlx.Tx) error {
	st := `INSERT INTO users (id, slug, role, username, password_digest, email, given_name, middle_names, family_name, last_ip,  confirmation_token, is_confirmed, locale, base_tz, current_tz, is_active, is_deleted, created_by_id, updated_by_id, created_at, updated_at)
VALUES (:id, :slug, :role, :username, :password_digest, :email, :given_name, :middle_names, :family_name, :last_ip, :confirmation_token, :is_confirmed, :locale, :base_tz, :current_tz, :is_active, :is_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at)`

	// Create a local transaction if it is not passed as argument.
	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	// Don't wait for repo to setup this values.
	// We want user ID to user as account owner ID.
	user.SetCreateValues()

	_, err = t.NamedExec(st, user)
	if err != nil {
		return err
	}

	// Commit on local transactions
	if local {
		return t.Commit()
	}

	return nil
}

// GetAll users from
func (ur *UserRepo) GetAll() (users []model.User, err error) {
	st := `SELECT * FROM users WHERE is_deleted IS NULL OR NOT is_deleted;`

	err = ur.DB.Select(&users, st)
	if err != nil {
		return users, err
	}

	return users, err
}

// Get user by ID.
func (ur *UserRepo) Get(id uuid.UUID) (user model.User, err error) {
	st := `SELECT * FROM USERS WHERE id = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, id.String())

	err = ur.DB.Get(&user, st)
	if err != nil {
		return user, err
	}

	return user, err
}

// GetBySlug user from repo by slug.
func (ur *UserRepo) GetBySlug(slug string) (user model.User, err error) {
	st := `SELECT * FROM USERS WHERE slug = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, slug)

	err = ur.DB.Get(&user, st)

	return user, err
}

// GetByUsername user from repo by username.
func (ur *UserRepo) GetByUsername(username string) (model.User, error) {
	var user model.User

	st := `SELECT * FROM USERS WHERE username = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, username)

	err := ur.DB.Get(&user, st)

	return user, err
}

// Update user data in
func (ur *UserRepo) Update(user *model.User, tx ...*sqlx.Tx) error {
	ref, err := ur.Get(user.ID)
	if err != nil {
		return fmt.Errorf("cannot retrieve reference user: %s", err.Error())
	}

	user.SetUpdateValues()

	var st strings.Builder
	pcu := false // previous column updated?

	st.WriteString("UPDATE users SET ")

	if user.Role != ref.Role {
		st.WriteString(fnd.SQLStrUpd("role", "role"))
		pcu = true
	}

	if user.Username != ref.Username {
		st.WriteString(fnd.SQLStrUpd("username", "username"))
		pcu = true
	}

	if user.PasswordDigest != ref.PasswordDigest {
		st.WriteString(fnd.SQLComma(pcu))
		st.WriteString(fnd.SQLStrUpd("password_digest", "password_digest"))
		pcu = true
	}

	if user.Email != ref.Email {
		st.WriteString(fnd.SQLComma(pcu))
		st.WriteString(fnd.SQLStrUpd("email", "email"))
		pcu = true
	}

	if user.GivenName != ref.GivenName {
		st.WriteString(fnd.SQLComma(pcu))
		st.WriteString(fnd.SQLStrUpd("given_name", "given_name"))
		pcu = true
	}

	if user.MiddleNames != ref.MiddleNames {
		st.WriteString(fnd.SQLComma(pcu))
		st.WriteString(fnd.SQLStrUpd("middle_names", "middle_names"))
		pcu = true
	}

	if user.FamilyName != ref.FamilyName {
		st.WriteString(fnd.SQLComma(pcu))
		st.WriteString(fnd.SQLStrUpd("family_name", "family_name"))
		pcu = true
	}

	if user.ConfirmationToken != ref.ConfirmationToken {
		st.WriteString(fnd.SQLComma(pcu))
		st.WriteString(fnd.SQLStrUpd("confirmation_token", "confirmation_token"))
		pcu = true
	}

	if user.IsConfirmed != ref.IsConfirmed {
		st.WriteString(fnd.SQLComma(pcu))
		st.WriteString(fnd.SQLStrUpd("is_confirmed", "is_confirmed"))
		pcu = true
	}

	if user.LastIP != ref.LastIP {
		st.WriteString(fnd.SQLComma(pcu))
		st.WriteString(fnd.SQLStrUpd("last_ip", "last_ip"))
		pcu = true
	}

	st.WriteString(" ")
	st.WriteString(fnd.SQLWhereID(ref.ID.String()))
	st.WriteString(";")

	//fmt.Println(st.String())

	if pcu == false {
		return errors.New("no fields to update")
	}

	// Create a local transaction if it is not passed as argument.
	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.NamedExec(st.String(), user)

	if local {
		ur.Log.Info("Transaction created by repo: committing")
		return t.Commit()
	}

	return nil
}

// Delete user from repo by ID.
func (ur *UserRepo) Delete(id uuid.UUID, tx ...*sqlx.Tx) error {
	st := `DELETE FROM USERS WHERE id = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, id)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// DeleteBySlug:w user from repo by slug.
func (ur *UserRepo) DeleteBySlug(slug string, tx ...*sqlx.Tx) error {
	st := `DELETE FROM USERS WHERE slug = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, slug)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// DeleteByusername user from repo by username.
func (ur *UserRepo) DeleteByUsername(username string, tx ...*sqlx.Tx) error {
	st := `DELETE FROM USERS WHERE username = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, username)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}
	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// GetBySlug user from repo by slug token.
func (ur *UserRepo) GetBySlugAndToken(slug, token string) (model.User, error) {
	var user model.User

	st := `SELECT * FROM USERS WHERE slug = '%s' AND confirmation_token = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, slug, token)

	err := ur.DB.Get(&user, st)

	return user, err
}

// Confirm user from repo by slug.
func (ur *UserRepo) ConfirmUser(slug, token string, tx ...*sqlx.Tx) (err error) {
	st := `UPDATE USERS SET is_confirmed = TRUE WHERE slug = '%s' AND confirmation_token = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, slug, token)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// SignIn user
func (ur *UserRepo) SignIn(username, password string) (model.User, error) {
	var u model.User

	st := `SELECT * FROM users WHERE username = '%s' OR email = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`

	st = fmt.Sprintf(st, username, username)

	err := ur.DB.Get(&u, st)

	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest.String), []byte(password))
	if err != nil {
		return u, err
	}

	return u, nil
}

func (ur *UserRepo) newTx() (tx *sqlx.Tx, err error) {
	tx, err = ur.DB.Beginx()
	if err != nil {
		return tx, err
	}

	return tx, err
}

func (ur *UserRepo) getTx(txs []*sqlx.Tx) (tx *sqlx.Tx, local bool, err error) {
	// Create a new transaction if its no passed as argument.
	if len(txs) > 0 {
		return txs[0], false, nil
	}

	tx, err = ur.DB.Beginx()
	if err != nil {
		return tx, true, err
	}

	return tx, true, err
}
