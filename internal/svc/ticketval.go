package svc

import (
	"errors"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
)

type (
	TicketValidator struct {
		Model *model.Ticket
		fnd.Validator
	}
)

func NewTicketValidator(u *model.Ticket) TicketValidator {
	return TicketValidator{
		Model:     u,
		Validator: fnd.NewValidator(),
	}
}

func (uv TicketValidator) ValidateForCreate() error {
	// Name
	ok0 := uv.ValidateRequiredName()
	ok1 := uv.ValidateMinLengthName(4)
	ok2 := uv.ValidateMaxLengthName(16)

	if ok0 && ok1 && ok2 {
		return nil
	}

	return errors.New("ticket has errors")
}

// NOTE: Update validations shoud be different
// than the ones executed on creation.
func (uv TicketValidator) ValidateForUpdate() error {
	// Name
	ok0 := uv.ValidateRequiredName()
	ok1 := uv.ValidateMinLengthName(4)
	ok2 := uv.ValidateMaxLengthName(16)

	if ok0 && ok1 && ok2 {
		return nil
	}

	return errors.New("ticket has errors")
}

func (uv TicketValidator) ValidateRequiredName(errMsg ...string) (ok bool) {
	u := uv.Model

	ok = uv.ValidateRequired(u.Name.String)
	if ok {
		return true
	}

	msg := fnd.ValMsg.RequiredErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	uv.Errors["Name"] = append(uv.Errors["Name"], msg)
	return false
}

func (uv TicketValidator) ValidateMinLengthName(min int, errMsg ...string) (ok bool) {
	u := uv.Model

	ok = uv.ValidateMinLength(u.Name.String, min)
	if ok {
		return true
	}

	msg := fnd.ValMsg.MinLengthErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	uv.Errors["Name"] = append(uv.Errors["Name"], msg)
	return false
}

func (uv TicketValidator) ValidateMaxLengthName(max int, errMsg ...string) (ok bool) {
	u := uv.Model

	ok = uv.ValidateMaxLength(u.Name.String, max)
	if ok {
		return true
	}

	msg := fnd.ValMsg.MaxLengthErrMsg
	if len(errMsg) > 0 {
		msg = errMsg[0]
	}

	uv.Errors["Name"] = append(uv.Errors["Name"], msg)
	return false
}
