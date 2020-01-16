package mig

// CreateTicketsable migration
func (s *step) CreateTicketsTable() error {
	tx := s.GetTx()

	st := `CREATE TABLE tickets
	(
		id UUID PRIMARY KEY,
		slug VARCHAR(36) UNIQUE,
		name VARCHAR(32),
		event_id UUID,
		type VARCHAR(32),
		serie VARCHAR(8),
		number INTEGER,
		seat VARCHAR(32),
		price INTEGER,
		currency VARCHAR(8),
		reserved_by UUID,
		reserved_at TIMESTAMP,
		bought_by UUID,
		bought_at TIMESTAMP,
		local_order_id VARCHAR(128),
		gateway_order_id VARCHAR(128),
		gateway_op_status VARCHAR(128)
	);`

	_, err := tx.Exec(st)
	if err != nil {
		return err
	}

	st = `
		ALTER TABLE tickets
		ADD COLUMN locale VARCHAR(32),
		ADD COLUMN base_tz VARCHAR(64),
		ADD COLUMN is_active BOOLEAN,
		ADD COLUMN is_deleted BOOLEAN,
		ADD COLUMN created_by_id UUID,
		ADD COLUMN updated_by_id UUID,
		ADD COLUMN created_at TIMESTAMP,
		ADD COLUMN updated_at TIMESTAMP;`

	_, err = tx.Exec(st)
	if err != nil {
		return err
	}

	return nil
}

// DropTicketsTable rollback
func (s *step) DropTicketsTable() error {
	tx := s.GetTx()

	st := `DROP TABLE tickets;`

	_, err := tx.Exec(st)
	if err != nil {
		return err
	}

	return nil
}
