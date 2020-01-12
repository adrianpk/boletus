package seed

import (
	"log"
)

var (
	// TODO: Create a builder for users that reads values from somewere: file, csv, map, etc...
	users = []map[string]interface{}{

		map[string]interface{}{"ID": "", "Username": "superadmin", "Email": "superadmin@localhost", "GivenName": "", "MiddleNames": "", "FamilyName": ""},
	}
)

// CreateUsers seeding
func (s *step) CreateUsers() error {
	tx := s.GetTx()

	st := `INSERT INTO users (id, slug, username, password_digest, email, given_name, middle_names, family_name, last_ip,  confirmation_token, is_confirmed, geolocation, locale, base_tz, current_tz, starts_at, ends_at, is_active, is_deleted, created_by_id, updated_by_id, created_at, updated_at)
VALUES (:id, :slug, :username, :password_digest, :email, :given_name, :middle_names, :family_name, :last_ip, :confirmation_token, :is_confirmed, :geolocation, :locale, :base_tz, :current_tz, :starts_at, :ends_at, :is_active, :is_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at)`

	// NOTE: Continue processing following after error?
	for u := range users {
		_, err := tx.NamedExec(st, u)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
