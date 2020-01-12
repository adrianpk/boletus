package svc

import (
	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
)

func (s *Service) IndexEvents() (users []model.Event, err error) {
	repo := s.EventRepo
	if repo == nil {
		return users, NoRepoErr
	}

	return repo.GetAll()
}

func (s *Service) CreateEvent(user *model.Event) (fnd.ValErrorSet, error) {
	// Validation
	v := NewEventValidator(user)

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
	userRepo := s.EventRepo
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

func (s *Service) GetEvent(slug string) (user model.Event, err error) {
	repo := s.EventRepo
	if err != nil {
		return user, err
	}

	user, err = repo.GetBySlug(slug)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) GetEventByName(username string) (user model.Event, err error) {
	repo := s.EventRepo
	if err != nil {
		return user, err
	}

	user, err = repo.GetByName(username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) UpdateEvent(slug string, user *model.Event) (fnd.ValErrorSet, error) {
	// Get a new transaction
	tx, err := s.getTx()
	if err != nil {
		return nil, err
	}

	userRepo := s.EventRepo
	if userRepo == nil {
		return nil, NoRepoErr
	}

	// Validation
	v := NewEventValidator(user)

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
