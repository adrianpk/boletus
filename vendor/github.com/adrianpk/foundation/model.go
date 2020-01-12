package kabestan

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/lib/pq"
	"github.com/adrianpk/foundation/db"
	"github.com/adrianpk/foundation/db/pg"

	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	// Identification model
	Identification struct {
		ID       uuid.UUID      `db:"id" json:-" schema:"-"`
		TenantID sql.NullString `db:"tenant_id" json:"-" schema:"-"`
		Slug     sql.NullString `db:"slug" json:"-" schema:"-"`
	}
)

// GetID representation.
func (i *Identification) GetID() uuid.UUID {
	return i.ID
}

// SetID for user.
func (i *Identification) SetID(id uuid.UUID) {
	i.ID = id
}

// GenID for user.
func (i *Identification) GenID() {
	if i.ID == uuid.Nil {
		i.ID = uuid.NewV4()
	}
}

// UpdateSlug if it was not set.
func (i *Identification) UpateSlug(prefix string) (slug string, err error) {
	if strings.Trim(i.Slug.String, " ") == "" {
		s, err := i.genSlug(prefix)
		if err != nil {
			return "", err
		}

		i.Slug = db.ToNullString(s)
	}

	return i.Slug.String, nil
}

// genSlug if it was not set.
func (i *Identification) genSlug(prefix string) (slug string, err error) {
	if strings.TrimSpace(prefix) == "" {
		return "", errors.New("no slug prefix defined")
	}

	prefix = strings.Replace(prefix, "-", "", -1)
	prefix = strings.Replace(prefix, "_", "", -1)

	if !utf8.ValidString(prefix) {
		v := make([]rune, 0, len(prefix))
		for i, r := range prefix {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(prefix[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		prefix = string(v)
	}

	prefix = strings.ToLower(prefix)

	s := strings.Split(uuid.NewV4().String(), "-")
	l := s[len(s)-1]

	return strings.ToLower(fmt.Sprintf("%s-%s", prefix, l)), nil
}

func (i *Identification) IsNew() bool {
	return i.ID == uuid.UUID{}
}

// SetCreateValues sets de ID and slug.
func (i *Identification) SetCreateValues(slugPrefix string) error {
	i.GenID()
	_, err := i.UpateSlug(slugPrefix)
	if err != nil {
		return err
	}
	return nil
}

// Match for Identification.
func (identification *Identification) Match(tc Identification) bool {
	r := identification.ID == tc.ID &&
		identification.TenantID == tc.TenantID &&
		identification.Slug == tc.Slug
	return r
}

type Audit struct {
	CreatedByID sql.NullString `db:"created_by_id" json:"-" schema:"-"`
	UpdatedByID sql.NullString `db:"updated_by_id" json:"-" schema:"-"`
	CreatedAt   pq.NullTime    `db:"created_at" json:"-" schema:"-"`
	UpdatedAt   pq.NullTime    `db:"updated_at" json:"-" schema:"-"`
}

// SetCreateValues sets de ID and slug.
func (a *Audit) SetCreateValues() error {
	now := time.Now()
	a.CreatedAt = pg.ToNullTime(now)
	a.UpdatedAt = pg.NullTime()
	return nil
}

// SetUpdateValues
func (a *Audit) SetUpdateValues() error {
	now := time.Now()
	a.UpdatedAt = pg.ToNullTime(now)
	return nil
}

// ToUUID parses string and return a UUID.
// If error returns aspecial form of UUID that
// is specified to have all 128 bits set to zero.
func ToUUID(uuidStr string) uuid.UUID {
	u, err := uuid.FromString(uuidStr)
	if err != nil {
		return uuid.Nil
	}

	return u
}

// ParseUUID string.
func ParseUUID(uuidStr string) (uuid.UUID, error) {
	return uuid.FromString(uuidStr)
}
