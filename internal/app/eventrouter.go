package app

import (
	"context"
	"net/http"

	"github.com/adrianpk/boletus/internal/app/web"
	"github.com/go-chi/chi"
)

// These routers require authorization
func (app *App) addWebEventRouter(parent chi.Router) chi.Router {
	return parent.Route("/events", func(uar chi.Router) {
		uar.Use(app.WebEP.ReqAuth)
		uar.Get("/", app.WebEP.IndexEvents)
		uar.Get("/new", app.WebEP.NewEvent)
		uar.Post("/", app.WebEP.CreateEvent)
		uar.Route("/{slug}", func(uarid chi.Router) {
			uarid.Use(eventCtx)
			uarid.Get("/", app.WebEP.ShowEvent)
			uarid.Get("/edit", app.WebEP.EditEvent)
			uarid.Patch("/", app.WebEP.UpdateEvent)
			uarid.Put("/", app.WebEP.UpdateEvent)
			uarid.Post("/init-delete", app.WebEP.InitDeleteEvent)
			uarid.Delete("/", app.WebEP.DeleteEvent)
		})
	})
}

func eventCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		ctx := context.WithValue(r.Context(), web.SlugCtxKey, slug)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
