package seed

import (
	"fmt"
	"log"
	"strings"
	"time"
	"unicode/utf8"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	// TODO: Create a builder for users that reads values from somewere: json, csv, xml, etc...
	users = []map[string]interface{}{

		newUserMap("superadmin", "superadmin", "greenhurley", "superadmin@boletus.club", "", "", ""),

		newUserMap("admin", "admin", "tallkristine", "admin@bloetus.club", "", "", ""),

		newUserMap("user", "lauriem", "openmontana", "lauriem@username.club", "Laurie", "Anne", "Miles"),
	}
)

// Users seeding
func (s *step) Users() error {
	tx := s.GetTx()

	st := `INSERT INTO users (id, slug, role, username, password_digest, email, given_name, middle_names, family_name, last_ip,  confirmation_token, is_confirmed, locale, base_tz, current_tz, is_active, is_deleted, created_by_id, updated_by_id, created_at, updated_at)
VALUES (:id, :slug, :role, :username, :password_digest, :email, :given_name, :middle_names, :family_name, :last_ip, :confirmation_token, :is_confirmed, :locale, :base_tz, :current_tz, :is_active, :is_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at)`

	// NOTE: Continue processing following after error?
	for _, u := range users {
		_, err := tx.NamedExec(st, u)
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
	}

	return nil
}

func newUserMap(role, username, password, email, givenName, middleNames, familyName string) map[string]interface{} {

	return map[string]interface{}{
		"id":                 genUUID(),
		"slug":               genSlug(username),
		"role":               role,
		"username":           username,
		"password_digest":    genPassDigest(password),
		"email":              email,
		"given_name":         givenName,
		"middle_names":       middleNames,
		"family_name":        familyName,
		"last_ip":            "198.24.10.0/24",
		"confirmation_token": genUUIDStr(),
		"is_confirmed":       true,
		"locale":             "en-US",
		"base_tz":            "GMT",
		"current_tz":         "CET",
		"is_active":          true,
		"is_deleted":         false,
		"created_by_id":      uuid.Nil,
		"updated_by_id":      uuid.Nil,
		"created_at":         time.Now(),
		"updated_at":         time.Time{},
	}
}

func genPassDigest(password string) string {
	pd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pd)
}

func genUUIDStr() string {
	return genUUID().String()
}

func genUUID() uuid.UUID {
	return uuid.NewV4()
}

func genSlug(prefix string) (slug string) {
	if strings.TrimSpace(prefix) == "" {
		prefix = "slug"
	}

	prefix = strings.Replace(prefix, "-", "", -1)
	prefix = strings.Replace(prefix, " ", "", -1)

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

	return strings.ToLower(fmt.Sprintf("%s-%s", prefix, l))
}
