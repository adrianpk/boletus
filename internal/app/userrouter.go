package app

import (
	"context"
	"net/http"

	"github.com/adrianpk/boletus/internal/app/web"
	"github.com/go-chi/chi"
)

func (app *App) addWebAuthRouter(parent chi.Router) chi.Router {
	return parent.Route("/auth", func(uar chi.Router) {
		uar.Get("/signup", app.WebEP.InitSignUpUser)
		uar.Post("/signup", app.WebEP.SignUpUser)
		uar.Get("/signin", app.WebEP.InitSignInUser)
		uar.Post("/signin", app.WebEP.SignInUser)
		uar.Get("/signout", app.WebEP.SignOutUser)
		uar.Route("/{slug}", func(uarid chi.Router) {
			uarid.Use(userCtx)
			uarid.Route("/{token}", func(uartkn chi.Router) {
				uartkn.Use(confCtx)
				uartkn.Get("/confirm", app.WebEP.ConfirmUser)
			})
		})
	})
}

// Thes handlers require authorization
func (app *App) addWebUserRouter(parent chi.Router) chi.Router {
	return parent.Route("/users", func(uar chi.Router) {
		uar.Use(app.WebEP.ReqAuth)
		uar.Get("/", app.WebEP.IndexUsers)
		uar.Get("/new", app.WebEP.NewUser)
		uar.Post("/", app.WebEP.CreateUser)
		uar.Route("/{slug}", func(uarid chi.Router) {
			uarid.Use(userCtx)
			uarid.Get("/", app.WebEP.ShowUser)
			uarid.Get("/edit", app.WebEP.EditUser)
			uarid.Patch("/", app.WebEP.UpdateUser)
			uarid.Put("/", app.WebEP.UpdateUser)
			uarid.Post("/init-delete", app.WebEP.InitDeleteUser)
			uarid.Delete("/", app.WebEP.DeleteUser)
			uarid.Route("/{token}", func(uartkn chi.Router) {
				uartkn.Use(confCtx)
				uartkn.Get("/confirm", app.WebEP.ConfirmUser)
			})
		})
	})
}

func userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		ctx := context.WithValue(r.Context(), web.SlugCtxKey, slug)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func confCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "token")
		ctx := context.WithValue(r.Context(), web.ConfCtxKey, slug)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
