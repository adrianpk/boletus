package svc

import (
	fnd "github.com/adrianpk/foundation"
)

type (
	Err struct {
		fnd.Err
	}

	noRepoError struct {
		fnd.Err
	}
)

var (
	NoRepoErr           = NewErr("no_repo_err", nil)
	AlreadyConfirmedErr = NewErr("already_confirmed_err_msg", nil)
	CredentialsErr      = NewErr("creadentials_err_msg", nil)
)

func NewErr(msgID string, err error) Err {
	return Err{
		fnd.NewErr(msgID, err),
	}
}
