package kabestan

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/markbates/pkger"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type (
	Router struct {
		Cfg  *Config
		Log  Logger
		Name string
		chi.Router
		i18nBundle *i18n.Bundle
	}
)

func NewRouter(cfg *Config, log Logger, name string) *Router {
	name = genName(name, "router")

	rt := Router{
		Cfg:    cfg,
		Log:    log,
		Name:   name,
		Router: chi.NewRouter(),
	}

	rt.Use(middleware.RequestID)
	rt.Use(middleware.RealIP)
	rt.Use(middleware.Recoverer)
	rt.Use(middleware.Timeout(60 * time.Second))
	rt.Use(rt.MethodOverride)
	rt.Use(rt.CSRFProtection)
	rt.Use(rt.I18N)

	return &rt
}

// Middlewares
// MethodOverride to emulate PUT and PATCH HTTP methodh.
func (rt *Router) MethodOverride(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			method := r.PostFormValue("_method")
			if method == "" {
				method = r.Header.Get("X-HTTP-Method-Override")
			}
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				r.Method = method
			}
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// CSRFProtection add cross-site request forgery protecction to the handler.
func (rt *Router) CSRFProtection(next http.Handler) http.Handler {
	return csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(next)
}

// I18N

func (rt *Router) I18N(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		bundle := rt.I18NBundle()
		l := i18n.NewLocalizer(bundle, lang, accept)
		ctx := context.WithValue(r.Context(), I18NorCtxKey, l)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func (rt *Router) I18NBundle() *i18n.Bundle {
	if rt.i18nBundle != nil {
		return rt.i18nBundle
	}

	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc("json", json.Unmarshal)

	locales := []string{"en", "pl", "de", "es"}

	for _, loc := range locales {
		path := fmt.Sprintf("/assets/web/embed/i18n/%s.json", loc)

		// Open pkger filer
		f, err := pkger.Open(path)
		if err != nil {
			rt.Log.Info("Opening embedded resource", "path", path)
			rt.Log.Error(err)
		}
		defer f.Close()

		// Read file content
		fs, err := f.Stat()
		if err != nil {
			rt.Log.Info("Stating embedded resource", "path", path)
			rt.Log.Error(err)
		}
		fd := make([]byte, fs.Size())
		f.Read(fd)

		// Load into bundle
		b.ParseMessageFileBytes(fd, path)
		//b.LoadEmbeddedMessageFile(fd, path)
	}

	// Cache it
	rt.i18nBundle = b

	return b
}
