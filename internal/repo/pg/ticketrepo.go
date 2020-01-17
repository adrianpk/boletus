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
	st := `INSERT INTO tickets (id, slug, name, event_id, serie, number, seat, price, currency, reservation_id, reserved_by_id, reserved_at, bought_by_id, bought_at, status, local_order_id, gateway_op_id,gateway_order_id, gateway_op_status, locale, base_tz, is_active, is_deleted, created_by_id, updated_by_id, created_at, updated_at)
VALUES (:id, :slug, :name, :event_id, :serie, :number, :seat, :price, :currency, :reservation_id, :reserved_by_id, :reserved_at, :bought_by_id, :bought_at, :status, :local_order_id, :gateway_op_id, gateway_order_id, gateway_op_status, :locale, :base_tz, :current_tz, :is_active, :is_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at)`

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

	if ticket.ReservationID != ref.ReservationID {
		st.WriteString(fnd.SQLStrUpd("reservation_id", "reservation_id"))
		pcu = true
	}

	if ticket.ReservedBy != ref.ReservedBy {
		st.WriteString(fnd.SQLStrUpd("reserved_by_id", "reserved_by_id"))
		pcu = true
	}

	if ticket.ReservedAt != ref.ReservedAt {
		st.WriteString(fnd.SQLStrUpd("reserved_at", "reserved_at"))
		pcu = true
	}

	if ticket.BoughtBy != ref.BoughtBy {
		st.WriteString(fnd.SQLStrUpd("bought_by_id", "bought_by_id"))
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

// Custom queries and process

// TicketSummary returns an availability of tickets report for all ticket types in an event.
func (ur *TicketRepo) TicketSummary(eventSlug string) (ts []model.TicketSummary, err error) {
	st := `SELECT count(tickets.id) as qty, event_id, events.slug as event_slug, type, (tickets.price/1000) as price, currency FROM tickets INNER JOIN events ON tickets.event_id = events.id WHERE events.slug = '%s' AND (tickets.is_deleted IS NULL OR NOT tickets.is_deleted) AND (events.is_deleted IS NULL OR NOT events.is_deleted) AND (reserved_by_id IS NULL OR reserved_by_id::text='00000000-0000-0000-0000-000000000000') AND (bought_by_id IS NULL OR NOT bought_by_id::text='00000000-0000-0000-0000-00000000') AND (status IS NULL OR status='') AND (gateway_op_id IS NULL or gateway_op_id='') GROUP BY event_id, event_slug, type, price, currency ORDER BY tickets.price ASC;`

	st = fmt.Sprintf(st, eventSlug)

	err = ur.DB.Select(&ts, st)

	return ts, err
}

// Available returns a report of number of available tickets for a specific ticket type in an event.
func (ur *TicketRepo) Available(eventSlug, ticketType string) (ts model.TicketSummary, err error) {
	st := `SELECT count(tickets.id) as qty, event_id, events.slug as event_slug, type, (tickets.price/1000) as price, currency FROM tickets INNER JOIN events ON tickets.event_id = events.id WHERE events.slug = '%s' AND tickets.type = '%s' AND (tickets.is_deleted IS NULL OR NOT tickets.is_deleted) AND (events.is_deleted IS NULL OR NOT events.is_deleted) AND (reserved_by_id IS NULL OR reserved_by_id::text='00000000-0000-0000-0000-000000000000') AND (bought_by_id IS NULL OR NOT bought_by_id::text='00000000-0000-0000-0000-00000000') AND (status IS NULL OR status='') AND (gateway_op_id IS NULL or gateway_op_id='') GROUP BY event_id, event_slug, type, price, currency ORDER BY tickets.price ASC LIMIT 1;`

	st = fmt.Sprintf(st, eventSlug, ticketType)

	err = ur.DB.Get(&ts, st)

	return ts, err
}

// Available returns a set of available tickets for a specific ticket type in an event.
func (ur *TicketRepo) GetAvailable(eventSlug, ticketType string, qty int) (tickets []model.Ticket, err error) {
	st := `SELECT tickets.* FROM tickets INNER JOIN events ON tickets.event_id = events.id WHERE events.slug = '%s' AND tickets.type = '%s' AND (tickets.is_deleted IS NULL OR NOT tickets.is_deleted) AND (events.is_deleted IS NULL OR NOT events.is_deleted) AND (reserved_by_id IS NULL OR reserved_by_id::text='00000000-0000-0000-0000-000000000000') AND (bought_by_id IS NULL OR NOT bought_by_id::text='00000000-0000-0000-0000-00000000') AND (status IS NULL OR status='') AND (gateway_op_id IS NULL or gateway_op_id='') GROUP BY event_id, event_slug, type, price, currency ORDER BY tickets.price ASC LIMIT %s;`

	st = fmt.Sprintf(st, eventSlug, ticketType, qty)

	err = ur.DB.Select(&tickets, st)

	return tickets, err
}

// PreBook mark as reserved a specific number of tickets of a certain type associated to an event.
func (ur *TicketRepo) PreBook(eventSlug, ticketType string, qty int, reservationID, userSlug string, tx ...*sqlx.Tx) (tickets []model.Ticket, err error) {
	// Create a local transaction if it is not passed as argument.
	t, local, err := ur.getTx(tx)
	if err != nil {
		return tickets, err
	}

	// TODO: Replace this extra query making another join
	// in update statement.
	st := `select ID from users where slug = '%s';`
	st = fmt.Sprintf(st, userSlug)

	var userID string
	row := t.QueryRow(st)
	err = row.Scan(&userID)
	if err != nil {
		ur.Log.Error(err)
		return tickets, err
	}

	st = `UPDATE tickets
					SET
						reservation_id = '%s',
						reserved_by_id = '%s',
						reserved_at = now(),
						status = 'reserved',
						updated_by_id = '%s',
						updated_at = now()
					WHERE tickets.id IN (
						SELECT tickets.id FROM tickets INNER JOIN events ON tickets.event_id = events.id WHERE events.slug = '%s' AND tickets.type = '%s' AND (tickets.is_deleted IS NULL OR NOT tickets.is_deleted) AND (events.is_deleted IS NULL OR NOT events.is_deleted) AND (reserved_by_id IS NULL OR reserved_by_id::text='00000000-0000-0000-0000-000000000000') AND (bought_by_id IS NULL OR NOT bought_by_id::text='00000000-0000-0000-0000-00000000') AND (status IS NULL OR status='') AND (gateway_op_id IS NULL or gateway_op_id='') ORDER BY tickets.price ASC LIMIT %d);`

	st = fmt.Sprintf(st, reservationID, userID, userID, eventSlug, ticketType, qty)

	// Update
	_, err = t.Query(st)
	if err != nil {
		return tickets, err
	}

	// Select all updated
	st = `SELECT tickets.* FROM tickets INNER JOIN events ON tickets.event_id = events.id WHERE events.slug = '%s' AND tickets.type = '%s' AND (tickets.is_deleted IS NULL OR NOT tickets.is_deleted) AND (events.is_deleted IS NULL OR NOT events.is_deleted) AND reservation_id = '%s' AND reserved_by_id::text='%s' AND (bought_by_id IS NULL OR NOT bought_by_id::text='00000000-0000-0000-0000-00000000') AND status='reserved' AND (gateway_op_id IS NULL or gateway_op_id='') ORDER BY tickets.updated_at ASC;`

	st = fmt.Sprintf(st, eventSlug, ticketType, reservationID, userID)

	err = ur.DB.Select(&tickets, st)
	if err != nil {
		return tickets, err
	}

	// Commit on local transactions
	if local {
		return tickets, t.Commit()
	}

	return tickets, nil
}

// ExpireReservations removes reservation marks from tickets if they are not confirmed after some expTimeMins minutes.
func (ur *TicketRepo) ExpireReservations(expMins int) (err error) {
	st := `UPDATE tickets
					SET
						reservation_id = '',
						reserved_by_id = '00000000-0000-0000-0000-000000000000',
						reserved_at = NULL,
						status = '',
						updated_by_id = '00000000-0000-0000-0000-000000000001',
						updated_at = now()
					WHERE tickets.id IN (
						SELECT tickets.id FROM tickets WHERE (tickets.is_deleted IS NULL OR NOT tickets.is_deleted) AND (reserved_by_id IS NOT NULL AND NOT reserved_by_id::text='00000000-0000-0000-0000-000000000000') AND (bought_by_id IS NULL OR bought_by_id::text='00000000-0000-0000-0000-000000000000') AND (status='reserved') AND (gateway_op_id IS NULL or gateway_op_id='') AND ((DATE_PART('hour', NOW() - tickets.reserved_at) * 60 + DATE_PART('minute', NOW() - reserved_at) > %d)));`

	st = fmt.Sprintf(st, expMins)

	// Update
	_, err = ur.DB.Query(st)
	if err != nil {
		return err
	}

	return nil
}

// Tx

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
