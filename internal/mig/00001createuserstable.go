package mig

import "log"

// CreateUsersTable migration
func (s *step) CreateUsersTable() error {
	tx := s.GetTx()

	st := `CREATE TABLE users
	(
		id UUID PRIMARY KEY,
		slug VARCHAR(36) UNIQUE,
		role VARCHAR(16),
		username VARCHAR(32) UNIQUE,
		password_digest CHAR(128),
		email VARCHAR(255) UNIQUE,
		given_name VARCHAR(32),
		middle_names VARCHAR(32) NULL,
		family_name VARCHAR(64),
		last_ip INET
	);`

	_, err := tx.Exec(st)
	if err != nil {
		return err
	}

	st = `
		ALTER TABLE users
		ADD COLUMN confirmation_token VARCHAR(36),
		ADD COLUMN is_confirmed BOOLEAN,
		ADD COLUMN locale VARCHAR(32),
		ADD COLUMN base_tz VARCHAR(2),
		ADD COLUMN current_tz VARCHAR(2),
		ADD COLUMN is_active BOOLEAN,
		ADD COLUMN is_deleted BOOLEAN,
		ADD COLUMN created_by_id UUID,
		ADD COLUMN updated_by_id UUID,
		ADD COLUMN created_at TIMESTAMP WITH TIME ZONE,
		ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE;`

	_, err = tx.Exec(st)
	if err != nil {
		return err
	}

	return nil
}

// DropUsersTable rollback
func (s *step) DropUsersTable() error {
	tx := s.GetTx()

	st := `DROP TABLE users;`

	_, err := tx.Exec(st)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
