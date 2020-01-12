package web

import (
	//"encoding/gob"

	"errors"
	"net/http"

	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/boletus/internal/app/svc"
)

type (
	Endpoint struct {
		*fnd.WebEndpoint
		Service *svc.Service
	}
)

const (
	SlugCtxKey fnd.ContextKey = "slug"
	ConfCtxKey fnd.ContextKey = "conf"
)

const (
	// Generic
	CannotProcErrMsg  = "cannot_proc_err_msg"
	InputValuesErrMsg = "input_values_err_msg"
	// Fields
	RequiredErrMsg   = "required_err_msg"
	MinLengthErrMsg  = "min_length_err_msg"
	MaxLengthErrMsg  = "max_length_err_msg"
	NotAllowedErrMsg = "not_allowed_err_msg"
	NotEmailErrMsg   = "not_email_err_msg"
	ConfMatchErrMsg  = "conf_match_err_msg"
)

func NewEndpoint(cfg *fnd.Config, log fnd.Logger, name string) (*Endpoint, error) {
	//registerGobTypes()

	wep, err := fnd.MakeWebEndpoint(cfg, log, pathFxs)
	if err != nil {
		return nil, err
	}

	return &Endpoint{
		WebEndpoint: wep,
	}, nil
}

// registerGobTypes
func registerGobTypes() {
	// gob.Register(CustomType1{})
	// gob.Register(CustomType2{})
	// gob.Register(CustomType3{})
}

// Middlewares
// ReqAuth require user authentication middleware.
func (ep *Endpoint) ReqAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		s, ok := ep.IsAuthenticated(r)
		if !ok {
			ep.Log.Debug("User not authenticated")
			http.Redirect(w, r, AuthPathSignIn(), 302)
			return
		}

		w.Header().Add("Cache-Control", "no-store")

		ep.Log.Debug("User authenticated", "slug", s)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (ep *Endpoint) getSlug(r *http.Request) (slug string, err error) {
	ctx := r.Context()
	slug, ok := ctx.Value(SlugCtxKey).(string)
	if !ok {
		err := errors.New("no slug provided")
		return "", err
	}

	return slug, nil
}
