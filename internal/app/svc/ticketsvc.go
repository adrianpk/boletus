package svc

import (
	"errors"
	"fmt"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
)

func (s *Service) IndexTickets() (users []model.Ticket, err error) {
	repo := s.TicketRepo
	if repo == nil {
		return users, NoRepoErr
	}

	return repo.GetAll()
}

func (s *Service) CreateTicket(user *model.Ticket) (fnd.ValErrorSet, error) {
	// Validation
	v := NewTicketValidator(user)

	err := v.ValidateForCreate()
	if err != nil {
		return v.Errors, err
	}

	user.SetCreateValues()

	// Get a new transaction
	tx, err := s.getTx()
	if err != nil {
		return nil, err
	}

	// Repo
	userRepo := s.TicketRepo
	if userRepo == nil {
		return nil, NoRepoErr
	}

	err = userRepo.Create(user, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Output
	return nil, nil
}

func (s *Service) GetTicket(slug string) (user model.Ticket, err error) {
	repo := s.TicketRepo
	if err != nil {
		return user, err
	}

	user, err = repo.GetBySlug(slug)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) UpdateTicket(slug string, user *model.Ticket) (fnd.ValErrorSet, error) {
	// Get a new transaction
	tx, err := s.getTx()
	if err != nil {
		return nil, err
	}

	userRepo := s.TicketRepo
	if userRepo == nil {
		return nil, NoRepoErr
	}

	// Validation
	v := NewTicketValidator(user)

	err = v.ValidateForUpdate()
	if err != nil {
		return v.Errors, err
	}

	// Update user
	err = userRepo.Update(user, tx)
	if err != nil {
		return v.Errors, err
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Output
	return v.Errors, nil
}

func (s *Service) DeleteTicket(slug string) error {
	repo := s.TicketRepo
	if repo == nil {
		return NoRepoErr
	}

	err := repo.DeleteBySlug(slug)
	if err != nil {
		return err
	}

	// Output
	return nil
}

// Custom queries and process
// TicketSummary returns a list including availability and price for each ticket type od an event.
func (s *Service) TicketSummary(eventSlug string) (users []model.TicketSummary, err error) {
	repo := s.TicketRepo
	if repo == nil {
		return users, NoRepoErr
	}

	return repo.TicketSummary(eventSlug)
}

// PreBookTickets
func (s *Service) PreBookTickets(eventSlug, ticketType string, qty int, userSlug string) (tickets []model.Ticket, err error) {
	repo := s.TicketRepo
	if repo == nil {
		return tickets, NoRepoErr
	}

	// Get a new transaction
	tx, err := s.getTx()
	if err != nil {
		return tickets, err
	}

	// Non resource intensive availabilty check
	ts, err := repo.Available(eventSlug, ticketType)
	if err != nil {
		tx.Commit()
		return tickets, err
	}

	avail := int(ts.Qty.Int32)
	if qty > avail {
		msg := fmt.Sprintf("requested quantity exceeds availability (avail.: %d)", avail)
		return tickets, errors.New(msg)
	}

	// Pre book tickets
	reservationID := fnd.GenShortID()
	tickets, err = repo.PreBook(eventSlug, ticketType, qty, reservationID, userSlug, tx)
	if err != nil {
		tx.Rollback()
		return tickets, err
	}

	// Commit on local transactions
	err = tx.Commit()
	if err != nil {
		return tickets, err
	}

	return tickets, nil
}

// ExpireTicketReservations
func (s *Service) ExpireTicketReservations() {
	s.Log.Info("Expire tickets process init.")
	repo := s.TicketRepo
	if repo == nil {
		s.Log.Error(NoRepoErr)
	}

	mins := int(s.Cfg.ValAsInt("reservation.expire.minutes", 15))

	err := repo.ExpireReservations(mins)
	if err != nil {
		s.Log.Error(err)
	}
}
