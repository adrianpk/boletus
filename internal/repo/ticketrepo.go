package repo

import (
	"github.com/adrianpk/boletus/internal/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	TicketRepo interface {
		Create(u *model.Ticket, tx ...*sqlx.Tx) error
		GetAll() (users []model.Ticket, err error)
		Get(id uuid.UUID) (user model.Ticket, err error)
		GetBySlug(slug string) (user model.Ticket, err error)
		GetByEventID(eventID uuid.UUID) (user []model.Ticket, err error)
		Update(user *model.Ticket, tx ...*sqlx.Tx) error
		Delete(id uuid.UUID, tx ...*sqlx.Tx) error
		DeleteBySlug(slug string, tx ...*sqlx.Tx) error
		// WIP: Later this methods can be moved to a custom repo
		TicketSummary(eventSlug string) (ticketSummary []model.TicketSummary, err error)
		Available(eventSlug, ticketType string) (ts model.TicketSummary, err error)
		GetAvailable(eventSlug, ticketType string, qty int) (tickets []model.Ticket, err error)
		PreBook(eventSlug, ticketType string, qty int, reservationID, userSlug string, tx ...*sqlx.Tx) (ts []model.Ticket, err error)
		PreBookType(eventSlug, ticketType string, reservationID, userSlug string, tx ...*sqlx.Tx) (ts []model.Ticket, err error)
		ExpireReservations(expMins int) (err error)
		ConfirmReservation(eventSlug, reservationID, userSlug string, tx ...*sqlx.Tx) (ts []model.Ticket, err error)
	}
)
