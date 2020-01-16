package model

import (
	"database/sql"
)

var (
	TicketTypes = ticketTypesReg()
)

var (
	normalTT = TicketType{
		SellingOptions: []string{"none"},
	}

	goldenTT = TicketType{
		SellingOptions: []string{"all-together"},
	}

	silverTT = TicketType{
		SellingOptions: []string{"even", "all-together"},
	}

	bronzeTT = TicketType{
		SellingOptions: []string{"none", "even", "all-together"},
	}

	couplesTT = TicketType{
		SellingOptions: []string{"even"},
	}
)

// Ticket kind
type (
	// Ticket
	TicketType struct {
		SellingOptions []string
	}

	ticketTypes struct {
		Normal  TicketType
		Golden  TicketType
		Silver  TicketType
		Bronze  TicketType
		Couples TicketType
	}
)

// Ticket kinds
func ticketTypesReg() *ticketTypes {
	return &ticketTypes{
		Normal:  normalTT,
		Golden:  goldenTT,
		Silver:  silverTT,
		Bronze:  bronzeTT,
		Couples: couplesTT,
	}
}

// Ticket summary
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
