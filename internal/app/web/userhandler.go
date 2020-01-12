package web

import (
	"errors"
	"fmt"
	"net/http"

	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/boletus/internal/app/svc"
	"github.com/adrianpk/boletus/internal/model"
)

const (
	userRes = "user"
)

const (
	// Defined in 'assets/web/embed/i18n/xx.json'
	UserCreatedInfoMsg = "user_created_info_msg"
	UserUpdatedInfoMsg = "user_updated_info_msg"
	UserDeletedInfoMsg = "user_deleted_info_msg"
	SignedUpInfoMsg    = "signed_up_info_msg"
	ConfirmedInfoMsg   = "confirmed_info_msg"
	SignedInInfoMsg    = "signed_in_info_msg"
	SignedOutInfoMsg   = "signed_out_info_msg"
	// Error
	CreateUserErrMsg        = "create_user_err_msg"
	IndexUsersErrMsg        = "get_all_users_err_msg"
	GetUserErrMsg           = "get_user_err_msg"
	GetUsersErrMsg          = "get_users_err_msg"
	UpdateUserErrMsg        = "update_user_err_msg"
	DeleteUserErrMsg        = "delete_user_err_msg"
	CredentialsErrMsg       = "credentials_err_msg"
	SignUpUserErrMsg        = "signup_err_msg"
	SignInUserErrMsg        = "signin_err_msg"
	ConfirmUserErrMsg       = "confirm_user_err_msg"
	ConfirmationTokenErrMsg = "confirmation_token_err_msg"
)

// IndexUsers web endpoint.
func (ep *Endpoint) IndexUsers(w http.ResponseWriter, r *http.Request) {
	// Get users list from registered service
	users, err := ep.Service.IndexUsers()
	if err != nil {
		ep.ErrorRedirect(w, r, "/", IndexUsersErrMsg, err)
		return
	}

	// Convert result list into a form list
	// Models use sql null types but templates looks
	// clearer if we use plain Go type.
	// i.e.: $user.Username instead of $user.Username.String
	l := model.ToUserFormList(users)
	wr := ep.WrapRes(w, r, l, nil)

	// Get template to render from cache.
	ts, err := ep.TemplateFor(userRes, fnd.IndexTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, "/", CannotProcErrMsg, err)
		return
	}

	// Execute it and redirect if error.
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, "/", CannotProcErrMsg, err)
		return
	}
}

func (ep *Endpoint) NewUser(w http.ResponseWriter, r *http.Request) {
	userForm := model.UserForm{IsNew: true}

	// Wrap response
	wr := ep.WrapRes(w, r, &userForm, nil)
	wr.SetAction(userCreateAction())

	// Get template to render from cache.
	ts, err := ep.TemplateFor(userRes, fnd.NewTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	// Execute it and redirect if error.
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}
}

func (ep *Endpoint) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decode request data into a form.
	userForm := model.UserForm{}
	err := ep.FormToModel(r, &userForm)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Create a model using form values.
	user := userForm.ToModel()

	// Use registered service to do everything related
	// to user creation.
	ves, err := ep.Service.CreateUser(&user)

	// First take care of service validation errors.
	if !ves.IsEmpty() {
		ep.rerenderUserForm(w, r, user.ToForm(), ves, fnd.NewTmpl, userCreateAction())
		return
	}

	// Then take care of other kind of possible errors
	// that service can generate.
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Localize Ok info message, put it into a flash message
	// and redirect to index.
	m := ep.Localize(r, UserCreatedInfoMsg)
	ep.RedirectWithFlash(w, r, UserPath(), m, fnd.InfoMT)
}

// ShowUser web endpoint.
func (ep *Endpoint) ShowUser(w http.ResponseWriter, r *http.Request) {
	// Get slug from request context.
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Use registered service to do everything related
	// to user creation.
	user, err := ep.Service.GetUser(s)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), GetUserErrMsg, err)
		return
	}

	// Wrap response
	wr := ep.WrapRes(w, r, user.ToForm(), nil)

	// Template
	ts, err := ep.TemplateFor(userRes, fnd.ShowTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}
}

// EditUser web endpoint.
func (ep *Endpoint) EditUser(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Use registerd service to get the user from repo.
	user, err := ep.Service.GetUser(s)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), GetUserErrMsg, err)
		return
	}

	// Wrap response
	userForm := user.ToForm()
	wr := ep.WrapRes(w, r, &userForm, nil)
	wr.SetAction(userUpdateAction(&userForm))

	// Template
	ts, err := ep.TemplateFor(userRes, fnd.EditTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}
}

// UpdateUser web endpoint.
func (ep *Endpoint) UpdateUser(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), GetUserErrMsg, err)
		return
	}

	// Decode request data into a form.
	userForm := model.UserForm{}
	err = ep.FormToModel(r, &userForm)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Create a model using form values.
	user := userForm.ToModel()

	// Use registered service to do everything related
	// to user update.
	ves, err := ep.Service.UpdateUser(s, &user)

	// First take care of service validation errors.
	if !ves.IsEmpty() {
		ep.Log.Debug("Validation errors", "dump", fmt.Sprintf("%+v", ves.FieldErrors))
		ep.rerenderUserForm(w, r, user.ToForm(), ves, fnd.NewTmpl, userCreateAction())
		return
	}

	// Non validation errors
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), UpdateUserErrMsg, err)
		return
	}

	m := ep.Localize(r, UserUpdatedInfoMsg)
	ep.RedirectWithFlash(w, r, UserPath(), m, fnd.InfoMT)
}

// InitDeleteUser web endpoint.
func (ep *Endpoint) InitDeleteUser(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Use registerd service to get the user from repo.
	user, err := ep.Service.GetUser(s)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), GetUsersErrMsg, err)
		return
	}

	// Wrap response
	userForm := user.ToForm()
	wr := ep.WrapRes(w, r, &userForm, nil)
	wr.SetAction(userDeleteAction(&userForm))

	// Template
	ts, err := ep.TemplateFor(userRes, fnd.InitDelTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}
}

// DeleteUser web endpoint.
func (ep *Endpoint) DeleteUser(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), DeleteUserErrMsg, err)
		return
	}

	// Service
	err = ep.Service.DeleteUser(s)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), DeleteUserErrMsg, err)
		return
	}

	m := ep.Localize(r, UserDeletedInfoMsg)
	ep.RedirectWithFlash(w, r, UserPath(), m, fnd.InfoMT)
}

func (ep *Endpoint) InitSignUpUser(w http.ResponseWriter, r *http.Request) {
	userForm := model.UserForm{}

	// Wrap response
	wr := ep.WrapRes(w, r, &userForm, nil)
	wr.SetAction(userSignUpAction())

	// Get template to render from cache.
	ts, err := ep.TemplateFor(userRes, fnd.SignUpTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, AuthPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	// Execute it and redirect if error.
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, AuthPath(), CannotProcErrMsg, err)
		return
	}
}

// SignUpUser web endpoint.
func (ep *Endpoint) SignUpUser(w http.ResponseWriter, r *http.Request) {
	// Decode request data into a form.
	userForm := model.UserForm{}
	err := ep.FormToModel(r, &userForm)
	if err != nil {
		ep.ErrorRedirect(w, r, AuthPath(), CannotProcErrMsg, err)
		return
	}

	// Create a model using form values.
	user := userForm.ToModel()

	// Get IP from user request
	// user.LastIP = db.ToNullString("0.0.0.0/24")

	// Use registered service to do everything related
	// to user creation.
	ves, err := ep.Service.SignUpUser(&user)

	// First take care of service validation errors.
	if !ves.IsEmpty() {
		ep.rerenderUserForm(w, r, user.ToForm(), ves, fnd.NewTmpl, userSignUpAction())
		return
	}

	// Then take care of other kind of possible errors
	// that service can generate.
	if err != nil {
		ep.ErrorRedirect(w, r, AuthPath(), SignUpUserErrMsg, err)
		return
	}

	// Localize Ok info message, put it into a flash message
	// and redirect to index.
	m := ep.Localize(r, SignedUpInfoMsg)
	ep.RedirectWithFlash(w, r, "/", m, fnd.InfoMT)
}

func (ep *Endpoint) InitSignInUser(w http.ResponseWriter, r *http.Request) {
	userForm := model.UserForm{}

	// Wrap response
	wr := ep.WrapRes(w, r, &userForm, nil)
	wr.SetAction(userSignInAction())

	// Get template to render from cache.
	ts, err := ep.TemplateFor(userRes, fnd.SignInTmpl)
	if err != nil {
		ep.ErrorRedirect(w, r, AuthPath(), CannotProcErrMsg, err)
		return
	}

	// Write response
	// Execute it and redirect if error.
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, AuthPath(), CannotProcErrMsg, err)
		return
	}
}

// SignInUser web endpoint.
func (ep *Endpoint) SignInUser(w http.ResponseWriter, r *http.Request) {
	// Decode request data into a form.
	userForm := model.UserForm{}
	err := ep.FormToModel(r, &userForm)
	if err != nil {
		ep.ErrorRedirect(w, r, AuthPath(), CannotProcErrMsg, err)
		return
	}

	// Get IP from user request
	// ip := db.ToNullString("0.0.0.0/24")
	// TODO: Provide IP to the service in order to register last IP
	// Can be used to detect spurious logins.
	// user, err := ep.Service.SignInUser(userForm.Username, userForm.Password, ip)
	user, err := ep.Service.SignInUser(userForm.Username, userForm.Password)

	if err != nil {
		msgID := SignInUserErrMsg

		// Give a hint to user about kind of error.
		if err == svc.CredentialsErr {
			msgID = (err.(svc.Err)).MsgID()
			ep.rerenderUserForm(w, r, user.ToForm(), nil, fnd.SignInTmpl, userSignInAction())
			return
		}

		ep.ErrorRedirect(w, r, UserPath(), msgID, err)
		return
	}

	// Register user slug in session.
	ep.SignIn(w, r, user.Slug.String)
	ep.Log.Debug("User signed in", "user", user.Username.String)

	// Localize Ok info message, put it into a flash message
	// and redirect to index.
	m := ep.Localize(r, SignedInInfoMsg)
	ep.RedirectWithFlash(w, r, UserPath(), m, fnd.InfoMT)
}

// SignOutUser web endpoint.
func (ep *Endpoint) SignOutUser(w http.ResponseWriter, r *http.Request) {
	ep.Log.Info("Signing out user")
	ep.SignOut(w, r)

	// Localize Ok info message, put it into a flash message
	// and redirect to index.
	m := ep.Localize(r, SignedOutInfoMsg)
	ep.RedirectWithFlash(w, r, UserPath(), m, fnd.InfoMT)
}

// ConfirmUser web endpoint.
func (ep *Endpoint) ConfirmUser(w http.ResponseWriter, r *http.Request) {
	s, err := ep.getSlug(r)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	// Token
	t, err := ep.getToken(r)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), ConfirmationTokenErrMsg, err)
		return
	}

	// Service
	err = ep.Service.ConfirmUser(s, t)
	if err != nil {
		msgID := ConfirmUserErrMsg

		// Give a hint to user if it was already confirmed.
		if err == svc.AlreadyConfirmedErr {
			msgID = (err.(svc.Err)).MsgID()
		}

		ep.ErrorRedirect(w, r, UserPath(), msgID, err)
		return
	}

	m := ep.Localize(r, UserCreatedInfoMsg)
	ep.RedirectWithFlash(w, r, UserPath(), m, fnd.InfoMT)
}

func (ep *Endpoint) rerenderUserForm(w http.ResponseWriter, r *http.Request, data interface{}, valErrors fnd.ValErrorSet, template string, action fnd.FormAction) {
	wr := ep.WrapRes(w, r, data, valErrors)
	wr.AddErrorFlash(InputValuesErrMsg)
	wr.SetAction(action)

	ts, err := ep.TemplateFor(userRes, template)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), InputValuesErrMsg, err)
		return
	}

	// Write response
	err = ts.Execute(w, wr)
	if err != nil {
		ep.ErrorRedirect(w, r, UserPath(), CannotProcErrMsg, err)
		return
	}

	return
}

// Misc
func (ep *Endpoint) getToken(r *http.Request) (token string, err error) {
	ctx := r.Context()
	token, ok := ctx.Value(ConfCtxKey).(string)
	if !ok {
		err := errors.New("no token provided")
		return "", err
	}

	return token, nil
}

// userCreateAction
func userCreateAction() fnd.FormAction {
	return fnd.FormAction{Target: fmt.Sprintf("%s", UserPath()), Method: "POST"}
}

// userUpdateAction
func userUpdateAction(model fnd.Identifiable) fnd.FormAction {
	return fnd.FormAction{Target: UserPathSlug(model), Method: "PUT"}
}

// userDeleteAction
func userDeleteAction(model fnd.Identifiable) fnd.FormAction {
	return fnd.FormAction{Target: UserPathSlug(model), Method: "DELETE"}
}

// userSignUpAction
func userSignUpAction() fnd.FormAction {
	return fnd.FormAction{Target: AuthPathSignUp(), Method: "POST"}
}

// userSignInAction
func userSignInAction() fnd.FormAction {
	return fnd.FormAction{Target: AuthPathSignIn(), Method: "POST"}
}
