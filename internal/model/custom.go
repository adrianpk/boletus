package model

import (
	"database/sql"
)

type (
	// TicketSummary model
	TicketSummary struct {
		Qty       sql.NullInt32   `db:"qty"`
		Name      sql.NullString  `db:"name"`
		EventID   sql.NullString  `db:"event_id""`
		EventSlug sql.NullString  `db:"event_slug""`
		Type      sql.NullString  `db:"type"`
		Price     sql.NullFloat64 `db:"price"`
		Currency  sql.NullString  `db:"currency"`
	}
)

// PriceFloat32 converts float64 to float32.
// We lose precission but sql null types are float64
// but protobuffer interface generator uses float32.
func (ts *TicketSummary) PriceFloat32() float32 {
	return float32(ts.Price.Float64)
}
