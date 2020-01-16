package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/foundation/db"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type (
	// Ticket model
	Ticket struct {
		fnd.Identification
		Name            sql.NullString `db:"name_id" json:"name_id" schema:"name_id"`
		EventID         sql.NullString `db:"event_id" json:"event_id" schema:"event_id"`
		Type            sql.NullString `db:"type" json:"type" schema:"type"`
		Serie           sql.NullString `db:"serie" json:"serie" schema:"serie"`
		Number          sql.NullString `db:"number" json:"number" schema:"number"`
		Seat            sql.NullString `db:"seat" json:"seat" schema:"seat"`
		Price           sql.NullInt32  `db:"price" json:"price" schema:"price"`
		Currency        sql.NullString `db:"currency" json:"currency" schema:"currency"`
		ReservedBy      sql.NullString `db:"reserved_by" json:"reservedBy" schema:"reserved-by"`
		ReservedAt      pq.NullTime    `db:"reserved_at" json:"reserved_at" schema:"reserved-at"`
		BoughtBy        sql.NullString `db:"bought_by" json:"boughtBy" schema:"bought-by"`
		BoughtAt        pq.NullTime    `db:"bought_at" json:"boughtAt" schema:"bought-at"`
		LocalOrderID    sql.NullString `db:"local_order_id" json:"localOrderID" schema:"local-order-at"`
		GatewayOrderID  sql.NullString `db:"gateway_order_id" json:"gatewayOrderID" schema:"local-order-at"`
		GatewayOpStatus sql.NullString `db:"gateway_op_status" json:"gatewayOpStatus" schema:"gateway-op-status"`
		Locale          sql.NullString `db:"locale" json:"-" schema:"-"`
		BaseTZ          sql.NullString `db:"base_tz" json:"-" schema:"-"`
		IsActive        sql.NullBool   `db:"is_active" json:"-" schema:"-"`
		IsDeleted       sql.NullBool   `db:"is_deleted" json:"-" schema:"-"`
		fnd.Audit
	}
)

type (
	TicketForm struct {
		Slug            string `json:"slug" schema:"slug"`
		Name            string `json:"name" schema:"name"`
		EventID         string `json:"eventID" schema:"event-id"`
		Type            string `json:"type" schema:"type"`
		Serie           string `json:"serie" schema:"serie"`
		Number          string `json:"number" schema:"number"`
		Seat            string `json:"seat" schema:"seat"`
		Price           string `json:"price" schema:"price"`
		Currency        string `json:"currency" schema:"currency"`
		ReservedBy      string `json:"reservedBy" schema:"reserved-by"`
		ReservedAt      string `json:"reservedAt" schema:"reserved-at"`
		BoughtBy        string `json:"boughtBy" schema:"bought-by"`
		BoughtAt        string `json:"boughtAt" schema:"bought-at"`
		LocalOrderID    string `json:"localOrderID" schema:"local-order-id"`
		GatewayOrderID  string `json:"GatewayOrderID" schema:"gateway-order-id"`
		GatewayOpStatus string `json:"GatewayOpStatus" schema:"gateway-op-status"`
		IsNew           bool   `json:"-" schema:"-"`
	}
)

func ToTicketFormList(tickets []Ticket) (fs []TicketForm) {
	for _, m := range tickets {
		fs = append(fs, m.ToForm())
	}
	return fs
}

// SetCreateValues sets de ID and slug.
func (ticket *Ticket) SetCreateValues() error {
	// Set create values only only if they were not previously
	if ticket.Identification.ID == uuid.Nil ||
		ticket.Identification.Slug.String == "" {
		pfx := ticket.Name.String
		ticket.Identification.SetCreateValues(pfx)
		ticket.Audit.SetCreateValues()
	}
	return nil
}

// SetUpdateValues
func (ticket *Ticket) SetUpdateValues() error {
	ticket.Audit.SetUpdateValues()
	return nil
}

// Match condition for
func (ticket *Ticket) Match(tc *Ticket) bool {
	r := ticket.Identification.Match(tc.Identification) &&
		ticket.Name.String == tc.Name.String &&
		ticket.EventID.String == tc.EventID.String &&
		ticket.Type.String == tc.Type.String &&
		ticket.Serie.String == tc.Serie.String &&
		ticket.Number.String == tc.Number.String &&
		ticket.Seat.String == tc.Seat.String &&
		ticket.Price.Int32 == tc.Price.Int32 &&
		ticket.Currency.String == tc.Currency.String &&
		ticket.ReservedBy.String == tc.ReservedBy.String &&
		ticket.ReservedAt.Time == tc.ReservedAt.Time &&
		ticket.BoughtBy.String == tc.BoughtBy.String &&
		ticket.BoughtAt.Time == tc.BoughtAt.Time &&
		ticket.LocalOrderID.String == tc.LocalOrderID.String &&
		ticket.GatewayOrderID.String == tc.GatewayOrderID.String &&
		ticket.GatewayOpStatus.String == tc.GatewayOpStatus.String
	return r
}

// ToForm lets convert a model to its associated form type.
// This convertion step could be avoided since gorilla schema allows
// to register custom decoders and in this case we need one because
// struct properties are not using Go standard types but their sql
// null conterpart types. As long as is relatively simple, ergonomic
// and could be easily implemented with generators I prefer to avoid
// the use of reflection.
// TODO: Move type conversion to web packate.
func (ticket *Ticket) ToForm() TicketForm {
	return TicketForm{
		Slug:            ticket.Slug.String,
		Name:            ticket.Name.String,
		EventID:         ticket.EventID.String,
		Type:            ticket.Type.String,
		Serie:           ticket.Serie.String,
		Number:          ticket.Number.String,
		Seat:            ticket.Seat.String,
		Price:           formatCurrency(ticket.Price.Int32),
		Currency:        ticket.Currency.String,
		ReservedBy:      ticket.ReservedBy.String,
		ReservedAt:      formatTime(ticket.ReservedAt.Time),
		BoughtBy:        ticket.BoughtBy.String,
		BoughtAt:        formatTime(ticket.BoughtAt.Time),
		LocalOrderID:    ticket.LocalOrderID.String,
		GatewayOrderID:  ticket.GatewayOrderID.String,
		GatewayOpStatus: ticket.GatewayOpStatus.String,
		IsNew:           ticket.IsNew(),
	}
}

// ToModel lets covert a form type to its associated model.
func (ticketForm *TicketForm) ToModel() Ticket {
	return Ticket{
		Identification: fnd.Identification{
			Slug: db.ToNullString(ticketForm.Slug),
		},
		Name:            db.ToNullString(ticketForm.Name),
		EventID:         db.ToNullString(ticketForm.EventID),
		Type:            db.ToNullString(ticketForm.Type),
		Serie:           db.ToNullString(ticketForm.Serie),
		Number:          db.ToNullString(ticketForm.Number),
		Seat:            db.ToNullString(ticketForm.Seat),
		Price:           db.ToNullInt32(toCurrency(ticketForm.Price)),
		Currency:        db.ToNullString(ticketForm.Currency),
		ReservedBy:      db.ToNullString(ticketForm.ReservedBy),
		BoughtBy:        db.ToNullString(ticketForm.BoughtBy),
		LocalOrderID:    db.ToNullString(ticketForm.LocalOrderID),
		GatewayOrderID:  db.ToNullString(ticketForm.GatewayOrderID),
		GatewayOpStatus: db.ToNullString(ticketForm.GatewayOpStatus),
	}
}

func (ticketForm *TicketForm) GetSlug() string {
	return ticketForm.Slug
}

func formatTime(time time.Time) string {
	layout := "2006-01-02T15:04:05.000Z"
	return time.Format(layout)
}

func toCurrencyStr(millis string) string {
	val := toCurrency(millis)
	return fmt.Sprintf("%d", val)
}

func toCurrency(millis string) int32 {
	val, err := strconv.ParseInt(millis, 10, 32)
	if err != nil {
		return int32(0)
	}
	return int32(val)
}

func formatCurrency(millis int32) string {
	return fmt.Sprintf("%.2f", millis/1000)
}
