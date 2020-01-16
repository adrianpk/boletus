package svc

import (
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
func (s *Service) TicketSummary(eventSlug string) (users []model.TicketSummary, err error) {
	repo := s.TicketRepo
	if repo == nil {
		return users, NoRepoErr
	}

	return repo.TicketSummary(eventSlug)
}
