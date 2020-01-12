package app

import (
	"net/http"
	"os"

	fnd "github.com/adrianpk/foundation"
	"github.com/markbates/pkger"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type (
	textResponse string
)

var (
	langMatcher = language.NewMatcher(message.DefaultCatalog.Languages())
)

func (app *App) NewWebRouter() *fnd.Router {
	rt := app.makeWebHomeRouter(app.Cfg, app.Log)
	// Auth
	app.addWebAuthRouter(rt)
	// User
	app.addWebUserRouter(rt)
	// Event
	app.addWebEventRouter(rt)
	app.addWebEventAdminRouter(rt)
	return rt
}

func (app *App) makeWebHomeRouter(cfg *fnd.Config, log fnd.Logger) *fnd.Router {
	rt := fnd.NewRouter(cfg, log, "web-home-router")
	app.addWebHomeRoutes(rt)
	return rt
}

func (app *App) addWebHomeRoutes(rt *fnd.Router) {
	dir := "/assets/web/embed/public"
	fs := http.FileServer(fnd.FileSystem{pkger.Dir(dir)})

	rt.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := pkger.Stat(dir + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)

		} else {
			fs.ServeHTTP(w, r)
		}
	})
}

func (t textResponse) write(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(t))
}
