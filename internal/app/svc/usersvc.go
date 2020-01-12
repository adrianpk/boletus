package svc

import (
	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
)

const (
	userAccountType = "user"
	orgAccountType  = "org"
)

const (
	UserConfirmationErrMsg = "user_confirmation_err_msg"
	UserSignInErrMsg       = "user_signin_err_msg"
)

func (s *Service) IndexUsers() (users []model.User, err error) {
	repo := s.UserRepo
	if repo == nil {
		return users, NoRepoErr
	}

	return repo.GetAll()
}

func (s *Service) CreateUser(user *model.User) (fnd.ValErrorSet, error) {
	// Validation
	v := NewUserValidator(user)

	err := v.ValidateForCreate()
	if err != nil {
		return v.Errors, err
	}

	// Confirmation
	user.GenAutoConfirmationToken()
	user.SetCreateValues()

	// Get a new transaction
	tx, err := s.getTx()
	if err != nil {
		return nil, err
	}

	// Repo
	userRepo := s.UserRepo
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

func (s *Service) GetUser(slug string) (user model.User, err error) {
	repo := s.UserRepo
	if err != nil {
		return user, err
	}

	user, err = repo.GetBySlug(slug)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) GetUserByUsername(username string) (user model.User, err error) {
	repo := s.UserRepo
	if err != nil {
		return user, err
	}

	user, err = repo.GetByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *Service) UpdateUser(slug string, user *model.User) (fnd.ValErrorSet, error) {
	// Get a new transaction
	tx, err := s.getTx()
	if err != nil {
		return nil, err
	}

	userRepo := s.UserRepo
	if userRepo == nil {
		return nil, NoRepoErr
	}

	// Get user
	current, err := userRepo.GetBySlug(slug)
	if err != nil {
		return nil, err
	}

	// Create a model
	// ID shouldn't change.
	user.ID = current.ID
	// Username can change if system enabled.
	// Set envar BLT_APP_USERNAME_UPDATABLE=true
	// to let username be updatable.
	if !(s.Cfg.ValAsBool("fnd.username.updatable", false)) {
		user.Username = current.Username
	}

	// Validation
	v := NewUserValidator(user)

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

func (s *Service) DeleteUser(slug string) error {
	repo := s.UserRepo
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

func (s *Service) SignUpUser(user *model.User) (fnd.ValErrorSet, error) {
	// Validation
	v := NewUserValidator(user)

	err := v.ValidateForSignUp()
	if err != nil {
		return v.Errors, err
	}

	// Generate confirmation token
	user.GenConfirmationToken()

	// Repo
	repo := s.UserRepo
	if repo == nil {
		return nil, NoRepoErr
	}

	err = repo.Create(user)
	if err != nil {
		return nil, err
	}

	// Mail confirmation
	s.sendConfirmationEmail(user)

	// Output
	return nil, nil
}

func (s *Service) ConfirmUser(slug, token string) error {
	repo := s.UserRepo
	if repo == nil {
		return NoRepoErr
	}

	user, err := repo.GetBySlugAndToken(slug, token)
	if err != nil {
		return NewErr(UserConfirmationErrMsg, err)
	}

	if user.IsConfirmed.Bool {
		return AlreadyConfirmedErr
	}

	err = repo.ConfirmUser(user.Slug.String, user.ConfirmationToken.String)
	if err != nil {
		return NewErr(UserConfirmationErrMsg, err)
	}

	// Output
	return nil
}

func (s *Service) SignInUser(username, password string) (user model.User, err error) {
	repo := s.UserRepo
	if repo == nil {
		return user, NoRepoErr
	}

	user, err = repo.SignIn(username, password)
	if err != nil {
		return user, CredentialsErr
	}

	// Output
	return user, nil
}
