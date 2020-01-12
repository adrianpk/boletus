package mig

import "log"

// CreateUsersTable migration
func (s *step) CreateEventsTable() error {
	tx := s.GetTx()

	st := `CREATE TABLE events
	(
		id UUID PRIMARY KEY,
		slug VARCHAR(36) UNIQUE,
		name VARCHAR(32) UNIQUE,
		description VARCHAR(512) UNIQUE,
		place VARCHAR(255) UNIQUE,
		scheduled_at TIMESTAMP WITH TIME ZONE
	);`

	_, err := tx.Exec(st)
	if err != nil {
		return err
	}

	st = `
		ALTER TABLE events
		ADD COLUMN locale VARCHAR(32),
		ADD COLUMN base_tz VARCHAR(2),
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

// DropEventsTable rollback
func (s *step) DropEventsTable() error {
	tx := s.GetTx()

	st := `DROP TABLE events;`

	_, err := tx.Exec(st)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
