package app

import (
	"fmt"
	"net/http"
	"sync"

	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/boletus/internal/app/web"
)

type (
	App struct {
		*fnd.App
		WebEP *web.Endpoint
	}
)

// NewApp creates a new app  worker instance.func NewApp(cfg *fnd.Config, log fnd.Logger, name string, core *fnd.Worker) (*App, error) {
func NewApp(cfg *fnd.Config, log fnd.Logger, name string) (*App, error) {
	app := App{
		App: fnd.NewApp(cfg, log, name),
	}

	// Endpoint
	wep, err := web.NewEndpoint(cfg, log, "web-endpoint")
	if err != nil {
		return nil, err
	}
	app.WebEP = wep

	// Router
	app.WebRouter = app.NewWebRouter()

	return &app, nil
}

// Init runs pre Start process.
func (app *App) Init() error {
	//return app.Migrator.RollbackAll()
	return app.Migrator.Migrate()
}

func (app *App) Start() error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		app.StartWeb()

		wg.Done()
	}()

	//wg.Add(1)
	//go func() {
	//a.StartJSONREST()
	//wg.Done()
	//}()

	wg.Wait()
	return nil
}

func (app *App) Stop() {
	// TODO: Gracefully stop all workers
}

func (app *App) StartWeb() error {
	p := app.Cfg.ValOrDef("web.server.port", "8080")
	p = fmt.Sprintf(":%s", p)

	app.Log.Info("Web server initializing", "port", p)

	err := http.ListenAndServe(p, app.WebRouter)
	app.Log.Error(err)

	return err
}
