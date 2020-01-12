package model

import (
	"database/sql"
	"fmt"
	"strings"

	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/foundation/db"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	// User model
	User struct {
		fnd.Identification
		Role              sql.NullString `db:"role" json:"role" schema:"role"`
		Username          sql.NullString `db:"username" json:"username" schema:"username"`
		Password          string         `db:"-" json:"password" schema:"password"`
		PasswordDigest    sql.NullString `db:"password_digest" json:"-" schema:"-"`
		Email             sql.NullString `db:"email" json:"email" schema:"email"`
		EmailConfirmation sql.NullString `db:"-" json:"emailConfirmation" schema:"email-confirmation"`
		GivenName         sql.NullString `db:"given_name" json:"givenName" schema:"given-name"`
		MiddleNames       sql.NullString `db:"middle_names" json:"middleNames" schema:"middle-names"`
		FamilyName        sql.NullString `db:"family_name" json:"familyName" schema:"family-name"`
		LastIP            sql.NullString `db:"last_ip" json:"-" schema:"-"`
		ConfirmationToken sql.NullString `db:"confirmation_token" json:"-" schema:"-"`
		IsConfirmed       sql.NullBool   `db:"is_confirmed" json:"-" schema:"-"`
		Locale            sql.NullString `db:"locale" json:"-" schema:"-"`
		BaseTZ            sql.NullString `db:"base_tz" json:"-" schema:"-"`
		CurrentTZ         sql.NullString `db:"current_tz" json:"-" schema:"-"`
		IsActive          sql.NullBool   `db:"is_active" json:"-" schema:"-"`
		IsDeleted         sql.NullBool   `db:"is_deleted" json:"-" schema:"-"`
		fnd.Audit
	}
)

type (
	UserForm struct {
		Slug              string `json:"slug" schema:"slug"`
		Username          string `json:"username" schema:"username"`
		Password          string `json:"password" schema:"password"`
		Email             string `json:"email" schema:"email"`
		EmailConfirmation string `json:"emailConfirmation" schema:"email-confirmation"`
		GivenName         string `json:"givenName" schema:"given-name"`
		MiddleNames       string `json:"middleNames" schema:"middle-names"`
		FamilyName        string `json:"familyName" schema:"family-name"`
		IsNew             bool   `json:"-" schema:"-"`
	}
)

func ToUserFormList(users []User) (fs []UserForm) {
	for _, m := range users {
		fs = append(fs, m.ToForm())
	}
	return fs
}

func (user *User) FullName() string {
	return strings.Trim(fmt.Sprintf("%s %s", user.GivenName.String, user.FamilyName.String), " ")
}

// UpdatePasswordDigest if password changed.
func (user *User) UpdatePasswordDigest() (digest string, err error) {
	if user.Password == "" {
		return user.PasswordDigest.String, nil
	}

	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.PasswordDigest.String, err
	}
	user.PasswordDigest = db.ToNullString(string(hpass))
	return user.PasswordDigest.String, nil
}

// SetCreateValues sets de ID and slug.
func (user *User) SetCreateValues() error {
	// Set create values only only if they were not previously
	if user.Identification.ID == uuid.Nil ||
		user.Identification.Slug.String == "" {
		pfx := user.Username.String
		user.Identification.SetCreateValues(pfx)
		user.Audit.SetCreateValues()
		user.UpdatePasswordDigest()
	}
	return nil
}

// SetUpdateValues
func (user *User) SetUpdateValues() error {
	user.Audit.SetUpdateValues()
	user.UpdatePasswordDigest()
	return nil
}

// GenConfirmationToken
func (user *User) GenConfirmationToken() {
	user.ConfirmationToken = db.ToNullString(uuid.NewV4().String())
	user.IsConfirmed = db.ToNullBool(false)
}

// GenAutoConfirmationToken
func (user *User) GenAutoConfirmationToken() {
	user.ConfirmationToken = db.ToNullString(uuid.NewV4().String())
	user.IsConfirmed = db.ToNullBool(true)
}

// Match condition for
func (user *User) Match(tc *User) bool {
	r := user.Identification.Match(tc.Identification) &&
		user.Username == tc.Username &&
		user.PasswordDigest == tc.PasswordDigest &&
		user.Email == tc.Email &&
		user.GivenName == tc.GivenName &&
		user.MiddleNames == tc.MiddleNames &&
		user.FamilyName == tc.FamilyName &&
		user.ConfirmationToken == tc.ConfirmationToken &&
		user.IsConfirmed == tc.IsConfirmed &&
		user.BaseTZ == tc.BaseTZ &&
		user.CurrentTZ == tc.CurrentTZ
	return r
}

// ToForm lets convert a model to its associated form type.
// This convertion step could be avoided since gorilla schema allows
// to register custom decoders and in this case we need one because
// struct properties are not using Go standard types but their sql
// null conterpart types. As long as is relatively simple, ergonomic
// and could be easily implemented with generators I prefer to avoid
// the use of reflection.
func (user *User) ToForm() UserForm {
	return UserForm{
		Slug:              user.Slug.String,
		Username:          user.Username.String,
		Email:             user.Email.String,
		EmailConfirmation: user.Email.String,
		GivenName:         user.GivenName.String,
		MiddleNames:       user.MiddleNames.String,
		FamilyName:        user.FamilyName.String,
		IsNew:             user.IsNew(),
	}
}

// ToModel lets covert a form type to its associated model.
func (userForm *UserForm) ToModel() User {
	return User{
		Identification: fnd.Identification{
			Slug: db.ToNullString(userForm.Slug),
		},
		Username:          db.ToNullString(userForm.Username),
		Password:          userForm.Password,
		Email:             db.ToNullString(userForm.Email),
		EmailConfirmation: db.ToNullString(userForm.EmailConfirmation),
		GivenName:         db.ToNullString(userForm.GivenName),
		MiddleNames:       db.ToNullString(userForm.MiddleNames),
		FamilyName:        db.ToNullString(userForm.FamilyName),
	}
}

func (userForm *UserForm) GetSlug() string {
	return userForm.Slug
}
