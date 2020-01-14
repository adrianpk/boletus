package pg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	TicketRepo struct {
		Cfg  *fnd.Config
		Log  fnd.Logger
		Name string
		DB   *sqlx.DB
	}
)

func NewTicketRepo(cfg *fnd.Config, log fnd.Logger, name string, db *sqlx.DB) *TicketRepo {
	return &TicketRepo{
		Cfg:  cfg,
		Log:  log,
		Name: name,
		DB:   db,
	}
}

// Create a ticket
func (ur *TicketRepo) Create(ticket *model.Ticket, tx ...*sqlx.Tx) error {
	st := `INSERT INTO tickets (id, slug, name, event_id, serie, number, seat, price, currency, reserved_by, reserved_at, bought_by, bought_at, local_order_id, gateway_order_id, gateway_op_status, locale, base_tz, is_active, is_deleted, created_by_id, updated_by_id, created_at, updated_at)
VALUES (:id, :slug, :name, :event_id, :serie, :number, :seat, :price, :currency, :reserved_by, :reserved_at, :bought_by, :bought_at, :local_order_id, gateway_order_id, gateway_op_status, :locale, :base_tz, :current_tz, :is_active, :is_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at)`

	// Create a local transaction if it is not passed as argument.
	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	// Don't wait for repo to setup this values.
	// We want ticket ID to ticket as account owner ID.
	ticket.SetCreateValues()

	_, err = t.NamedExec(st, ticket)
	if err != nil {
		return err
	}

	// Commit on local transactions
	if local {
		return t.Commit()
	}

	return nil
}

// GetAll tickets from
func (ur *TicketRepo) GetAll() (tickets []model.Ticket, err error) {
	st := `SELECT * FROM tickets WHERE is_deleted IS NULL OR NOT is_deleted;`

	err = ur.DB.Select(&tickets, st)
	if err != nil {
		return tickets, err
	}

	return tickets, err
}

// Get ticket by ID.
func (ur *TicketRepo) Get(id uuid.UUID) (ticket model.Ticket, err error) {
	st := `SELECT * FROM tickets WHERE id = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, id.String())

	err = ur.DB.Get(&ticket, st)
	if err != nil {
		return ticket, err
	}

	return ticket, err
}

// GetBySlug ticket from repo by slug.
func (ur *TicketRepo) GetBySlug(slug string) (ticket model.Ticket, err error) {
	st := `SELECT * FROM tickets WHERE slug = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, slug)

	err = ur.DB.Get(&ticket, st)

	return ticket, err
}

// GetByEventID ticket from repo by eventID.
func (ur *TicketRepo) GetByEventID(eventID uuid.UUID) (tickets []model.Ticket, err error) {
	st := `SELECT * FROM tickets WHERE event_id = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, eventID)

	err = ur.DB.Select(&tickets, st)
	if err != nil {
		return tickets, err
	}

	return tickets, err
}

// Update ticket data in
func (ur *TicketRepo) Update(ticket *model.Ticket, tx ...*sqlx.Tx) error {
	ref, err := ur.Get(ticket.ID)
	if err != nil {
		return fmt.Errorf("cannot retrieve reference ticket: %s", err.Error())
	}

	ticket.SetUpdateValues()

	var st strings.Builder
	pcu := false // previous column updated?

	st.WriteString("UPDATE tickets SET ")

	if ticket.Name != ref.Name {
		st.WriteString(fnd.SQLStrUpd("name", "name"))
		pcu = true
	}

	if ticket.EventID != ref.EventID {
		st.WriteString(fnd.SQLStrUpd("event_id", "event_id"))
		pcu = true
	}

	if ticket.Serie != ref.Serie {
		st.WriteString(fnd.SQLStrUpd("serie", "serie"))
		pcu = true
	}

	if ticket.Number != ref.Number {
		st.WriteString(fnd.SQLStrUpd("number", "number"))
		pcu = true
	}

	if ticket.Seat != ref.Seat {
		st.WriteString(fnd.SQLStrUpd("seat", "seat"))
		pcu = true
	}

	if ticket.Price != ref.Price {
		st.WriteString(fnd.SQLStrUpd("price", "price"))
		pcu = true
	}

	if ticket.Currency != ref.Currency {
		st.WriteString(fnd.SQLStrUpd("currency", "currency"))
		pcu = true
	}

	if ticket.ReservedBy != ref.ReservedBy {
		st.WriteString(fnd.SQLStrUpd("reserved_by", "reserved_by"))
		pcu = true
	}

	if ticket.ReservedAt != ref.ReservedAt {
		st.WriteString(fnd.SQLStrUpd("reserved_at", "reserved_at"))
		pcu = true
	}

	if ticket.BoughtBy != ref.BoughtBy {
		st.WriteString(fnd.SQLStrUpd("bought_by", "bought_by"))
		pcu = true
	}

	if ticket.BoughtAt != ref.BoughtAt {
		st.WriteString(fnd.SQLStrUpd("bought_at", "bought_at"))
		pcu = true
	}

	if ticket.LocalOrderID != ref.LocalOrderID {
		st.WriteString(fnd.SQLStrUpd("local_order_id", "local_order_id"))
		pcu = true
	}

	if ticket.GatewayOrderID != ref.GatewayOrderID {
		st.WriteString(fnd.SQLStrUpd("gateway_order_id", "gateway_order_id"))
		pcu = true
	}

	if ticket.GatewayOpStatus != ref.GatewayOpStatus {
		st.WriteString(fnd.SQLStrUpd("gateway_op_status", "gateway_op_status"))
		pcu = true
	}

	st.WriteString(" ")
	st.WriteString(fnd.SQLWhereID(ref.ID.String()))
	st.WriteString(";")

	//fmt.Println(st.String())

	if pcu == false {
		return errors.New("no fields to update")
	}

	// Create a local transaction if it is not passed as argument.
	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.NamedExec(st.String(), ticket)

	if local {
		ur.Log.Info("Transaction created by repo: committing")
		return t.Commit()
	}

	return nil
}

// Delete ticket from repo by ID.
func (ur *TicketRepo) Delete(id uuid.UUID, tx ...*sqlx.Tx) error {
	st := `DELETE FROM tickets WHERE id = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, id)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// DeleteBySlug:w ticket from repo by slug.
func (ur *TicketRepo) DeleteBySlug(slug string, tx ...*sqlx.Tx) error {
	st := `DELETE FROM tickets WHERE slug = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, slug)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// DeleteByticketname ticket from repo by ticketname.
func (ur *TicketRepo) DeleteByTicketname(ticketname string, tx ...*sqlx.Tx) error {
	st := `DELETE FROM tickets WHERE ticketname = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, ticketname)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}
	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// GetBySlug ticket from repo by slug token.
func (ur *TicketRepo) GetBySlugAndToken(slug, token string) (model.Ticket, error) {
	var ticket model.Ticket

	st := `SELECT * FROM tickets WHERE slug = '%s' AND confirmation_token = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, slug, token)

	err := ur.DB.Get(&ticket, st)

	return ticket, err
}

func (ur *TicketRepo) newTx() (tx *sqlx.Tx, err error) {
	tx, err = ur.DB.Beginx()
	if err != nil {
		return tx, err
	}

	return tx, err
}

func (ur *TicketRepo) getTx(txs []*sqlx.Tx) (tx *sqlx.Tx, local bool, err error) {
	// Create a new transaction if its no passed as argument.
	if len(txs) > 0 {
		return txs[0], false, nil
	}

	tx, err = ur.DB.Beginx()
	if err != nil {
		return tx, true, err
	}

	return tx, true, err
}
