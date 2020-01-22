package app

import (
	"fmt"
	"net/http"
	"sync"

	svc "github.com/adrianpk/boletus/internal/svc"
	v1 "github.com/adrianpk/boletus/pkg/grpc/api/v1"
	"github.com/adrianpk/boletus/pkg/web"
	fnd "github.com/adrianpk/foundation"

	"net"

	"google.golang.org/grpc"
)

type (
	App struct {
		*fnd.App
		WebEP     *web.Endpoint
		GRPCAPIV1 *v1.GRPCService
		Scheduler *scheduler
	}
)

// NewApp creates a new app  worker instance.func NewApp(cfg *fnd.Config, log fnd.Logger, name string, core *fnd.Worker) (*App, error) {
func NewApp(cfg *fnd.Config, log fnd.Logger, name string) (*App, error) {
	app := App{
		App: fnd.NewApp(cfg, log, name),
	}

	// Web Endpoint
	wep, err := web.NewEndpoint(cfg, log, "web-endpoint")
	if err != nil {
		return nil, err
	}
	app.WebEP = wep

	// Router
	app.WebRouter = app.NewWebRouter()

	// gRPC Server
	gsv1 := v1.NewGRPCService(cfg, log, "grpc-service-v1")
	app.GRPCAPIV1 = gsv1

	// Scheduler
	app.Scheduler = NewScheduler(cfg, log, "scheduler")

	return &app, nil
}

// Init runs pre Start process.
func (app *App) Init() error {
	//return app.Migrator.RollbackAll()
	err := app.Migrator.Migrate()
	if err != nil {
		return err
	}

	err = app.Seeder.Seed()
	if err != nil {
		return err
	}

	app.GRPCAPIV1.Service.Init()

	return nil
}

func (app *App) Start() error {
	var wg sync.WaitGroup

	err := app.Scheduler.Start()
	if err != nil {
		return err
	}
	app.Log.Info("Scheduler started")

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

	wg.Add(1)
	go func() {
		app.StartGRPC()

		wg.Done()
	}()

	wg.Wait()
	return nil
}

func (app *App) Stop() {
	// TODO: Gracefully stop all workers
}

func (app *App) SetService(svc *svc.Service) {
	if app.WebEP != nil {
		app.WebEP.Service = svc
	}

	if app.GRPCAPIV1 != nil {
		app.GRPCAPIV1.Service = svc
	}

	if app.Scheduler != nil {
		app.Scheduler.Service = svc
	}
}

// StartGRPC starts a web server to publish ticketer service.
func (app *App) StartWeb() error {
	p := app.Cfg.ValOrDef("web.server.port", "8080")
	p = fmt.Sprintf(":%s", p)

	app.Log.Info("Web server initializing", "port", p)

	err := http.ListenAndServe(p, app.WebRouter)
	app.Log.Error(err)

	return err
}

// StartGRPC starts a gRPC server to publish ticketer service.
func (app *App) StartGRPC() error {
	p := app.Cfg.ValOrDef("grpc.server.port", "8082")
	p = fmt.Sprintf(":%s", p)

	app.Log.Info("gRPC server initializing", "port", p)

	listen, err := net.Listen("tcp", p)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	v1.RegisterTicketerServer(s, app.GRPCAPIV1)

	return s.Serve(listen)
}
