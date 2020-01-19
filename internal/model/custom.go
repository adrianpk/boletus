package model

import (
	"database/sql"
)

type (
	SellingOption string
)

var (
	// TicketType
	TicketTypes = ticketTypesReg()

	// SelingOptions
	NoneSO        SellingOption = "none"
	AllTogetherSO SellingOption = "all-together"
	EvenSO        SellingOption = "even"
	PreemptiveSO  SellingOption = "preemptive"
)

var (
	normalTT = TicketType{
		Name:          "normal",
		SellingOption: NoneSO,
	}

	goldenTT = TicketType{
		Name:          "golden-circle",
		SellingOption: AllTogetherSO,
	}

	couplesTT = TicketType{
		Name:          "couples",
		SellingOption: EvenSO,
	}

	preemptiveTT = TicketType{
		Name:          "preemptive",
		SellingOption: PreemptiveSO,
	}
)

// Ticket types
type (
	// Ticket
	TicketType struct {
		Name          string
		SellingOption SellingOption
	}

	ticketTypes struct {
		Normal     TicketType
		Golden     TicketType
		Couples    TicketType
		Preemptive TicketType
	}
)

// Ticket kinds
func ticketTypesReg() *ticketTypes {
	return &ticketTypes{
		Normal:     normalTT,
		Golden:     goldenTT,
		Couples:    couplesTT,
		Preemptive: preemptiveTT,
	}
}

func TicketTypeByName(name string) TicketType {
	switch name {

	case "normal":
		return TicketTypes.Normal

	case "golden-circle":
		return TicketTypes.Golden

	case "couples":
		return TicketTypes.Couples

	case "preemptive":
		return TicketTypes.Preemptive

	default:
		return TicketTypes.Normal
	}
}

// Ticket summary
type (
	// TicketSummary model
	TicketSummary struct {
		Qty       sql.NullInt32      `db:"qty"`
		Name      sql.NullString     `db:"name"`
		EventID   sql.NullString     `db:"event_id""`
		EventSlug sql.NullString     `db:"event_slug""`
		Type      sql.NullString     `db:"type"`
		Price     sql.NullFloat64    `db:"price"`
		Currency  sql.NullString     `db:"currency"`
		Prices    map[string]float32 `db:"-"`
	}
)

// PriceFloat32 converts float64 to float32.
// We lose precission but sql null types are float64
// but protobuffer interface generator uses float32.
func (ts *TicketSummary) PriceFloat32() float32 {
	return float32(ts.Price.Float64)
}

// PreBookOp
type (
	// PreBookOp model
	PreBookOp struct {
		Name      sql.NullString  `db:"name"`
		EventID   sql.NullString  `db:"event_id""`
		EventSlug sql.NullString  `db:"event_slug""`
		Type      sql.NullString  `db:"type"`
		Qty       sql.NullInt64   `db:"qty"`
		Price     sql.NullFloat64 `db:"price"`
		Total     sql.NullFloat64 `db:"total"`
		Currency  sql.NullString  `db:"currency"`
	}
)

func (pbo *PreBookOp) PriceFloat32() float32 {
	return float32(pbo.Price.Float64)
}

func (pbo *PreBookOp) TotalFloat32() float32 {
	return float32(pbo.Total.Float64)
}
