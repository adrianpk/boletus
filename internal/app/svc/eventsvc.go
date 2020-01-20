package svc

import (
	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
)

func (s *Service) IndexEvents() (events []model.Event, err error) {
	repo := s.EventRepo
	if repo == nil {
		return events, NoRepoErr
	}

	return repo.GetAll()
}

func (s *Service) CreateEvent(event *model.Event) (fnd.ValErrorSet, error) {
	// Validation
	v := NewEventValidator(event)

	err := v.ValidateForCreate()
	if err != nil {
		return v.Errors, err
	}

	event.SetCreateValues()

	// Get a new transaction
	tx, err := s.getTx()
	if err != nil {
		return nil, err
	}

	// Repo
	userRepo := s.EventRepo
	if userRepo == nil {
		return nil, NoRepoErr
	}

	err = userRepo.Create(event, tx)
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

func (s *Service) GetEvent(slug string) (event model.Event, err error) {
	repo := s.EventRepo
	if err != nil {
		return event, err
	}

	event, err = repo.GetBySlug(slug)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (s *Service) GetEventByName(name string) (event model.Event, err error) {
	repo := s.EventRepo
	if err != nil {
		return event, err
	}

	event, err = repo.GetByName(name)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (s *Service) UpdateEvent(slug string, event *model.Event) (fnd.ValErrorSet, error) {
	// Get a new transaction
	tx, err := s.getTx()
	if err != nil {
		return nil, err
	}

	eventRepo := s.EventRepo
	if eventRepo == nil {
		return nil, NoRepoErr
	}

	// Validation
	v := NewEventValidator(event)

	err = v.ValidateForUpdate()
	if err != nil {
		return v.Errors, err
	}

	// Update user
	err = eventRepo.Update(event, tx)
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

func (s *Service) DeleteEvent(slug string) error {
	repo := s.EventRepo
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
