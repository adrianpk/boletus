package pg

import (
	"errors"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/foundation/db"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	TicketRepo struct {
		Cfg  *fnd.Config
		Log  fnd.Logger
		Name string
	}
)

type (
	ticketRow struct {
		mutable bool
		model   model.Ticket
	}
)

var (
	rockParty1 = model.Ticket{
		Identification: fnd.Identification{
			ID:   fnd.ToUUID("03bc0622-4612-4490-9818-c50b2fd2340"),
			Slug: db.ToNullString("rockparty-d11167cf0a82"),
		},
		Name:     db.ToNullString("Rock Party in Wrocław"),
		EventID:  db.ToNullString("03bc0622-4612-4490-9818-c50b2fd2340"),
		Serie:    db.ToNullString("A"),
		Number:   db.ToNullString("125"),
		Seat:     db.ToNullString("25/A"),
		Price:    db.ToNullInt32(50000),
		Currency: db.ToNullString("PLN"),
	}

	jesseCook1 = model.Ticket{
		Identification: fnd.Identification{
			ID:   fnd.ToUUID("363f901e-5aa8-468b-8805-4e4fa560f6e0"),
			Slug: db.ToNullString("jessecook-7dc35c1216b"),
		},
		Name:     db.ToNullString("Jesse Cook"),
		EventID:  db.ToNullString("f2553fc3-3c63-4adb-bc92-be8c66aecd19"),
		Serie:    db.ToNullString("C"),
		Number:   db.ToNullString("12"),
		Seat:     db.ToNullString("n/a"),
		Price:    db.ToNullInt32(38000),
		Currency: db.ToNullString("PLN"),
	}

	ticketsTable = map[uuid.UUID]ticketRow{
		rockParty1.ID: ticketRow{mutable: false, model: rockParty1},
		jesseCook1.ID: ticketRow{mutable: false, model: jesseCook1},
	}
)

func (ur *TicketRepo) Create(ticket *model.Ticket, tx ...*sqlx.Tx) error {
	_, ok := ticketsTable[ticket.ID]
	if ok {
		errors.New("duplicate key violates unique constraint")
	}

	if ticket.ID == uuid.Nil {
		errors.New("Non valid primary key")
	}

	ticketsTable[ticket.ID] = ticketRow{
		mutable: true,
		model:   *ticket,
	}

	return nil
}

func (ur *TicketRepo) GetAll() (tickets []model.Ticket, err error) {
	size := len(ticketsTable)
	out := make([]model.Ticket, size)
	for _, row := range ticketsTable {
		out = append(out, row.model)
	}
	return out, nil
}

func (ur *TicketRepo) Get(id uuid.UUID) (ticket model.Ticket, err error) {
	for _, row := range ticketsTable {
		if id == row.model.ID {
			return row.model, nil
		}
	}
	return model.Ticket{}, nil
}

func (ur *TicketRepo) GetBySlug(slug string) (ticket model.Ticket, err error) {
	for _, row := range ticketsTable {
		if slug == row.model.Slug.String {
			return row.model, nil
		}
	}
	return model.Ticket{}, nil
}

func (ur *TicketRepo) GetByEventID(eventID uuid.UUID) (ticket model.Ticket, err error) {
	out := make([]model.Ticket, 1000)
	for _, row := range ticketsTable {
		if eventID == row.model.ID {
			out = append(out, row.model)
		}
	}

	return model.Ticket{}, nil
}

func (ur *TicketRepo) Update(ticket *model.Ticket, tx ...*sqlx.Tx) error {
	for _, row := range ticketsTable {
		if ticket.ID == row.model.ID {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			ticketsTable[ticket.ID] = ticketRow{
				mutable: true,
				model:   *ticket,
			}
			return nil
		}
	}
	return errors.New("no records updated")
}

func (ur *TicketRepo) Delete(id uuid.UUID, tx ...*sqlx.Tx) error {
	for _, row := range ticketsTable {
		if id == row.model.ID {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			delete(ticketsTable, id)
			return nil
		}
	}
	return errors.New("no records deleted")
}

func (ur *TicketRepo) DeleteySlug(slug string, tx ...*sqlx.Tx) error {
	for _, row := range ticketsTable {
		if slug == row.model.Slug.String {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			delete(ticketsTable, row.model.ID)
			return nil
		}
	}
	return errors.New("no records deleted")
}

func (ur *TicketRepo) TicketSummary(eventSlug string) (ts []model.TicketSummary, err error) {
	panic("not implemented")
}

func (ur *TicketRepo) Available(eventSlug, ticketType string) (ts model.TicketSummary, err error) {
	panic("not implemented")
}

func (ur *TicketRepo) PreBook(eventSlug, ticketType string, qty int, userSlug string, tx ...*sqlx.Tx) (ts []model.Ticket, err error) {
	panic("not implemented")
}
