package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianpk/boletus/internal/app"
	"github.com/adrianpk/boletus/internal/app/svc"
	"github.com/adrianpk/boletus/internal/mig"
	repo "github.com/adrianpk/boletus/internal/repo/pg"

	//vrepo "github.com/adrianpk/boletus/internal/repo/vol"
	"github.com/adrianpk/boletus/internal/seed"
	fnd "github.com/adrianpk/foundation"
)

type contextKey string

const (
	appName = "boletus"
)

var (
	a *app.App
)

func main() {
	cfg := fnd.LoadConfig("blt")
	log := fnd.NewLogger(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	checkStopEvents(ctx, cancel)

	// App
	a, err := app.NewApp(cfg, log, appName)
	if err != nil {
		exit(log, err)
	}

	// Database connection
	db, err := fnd.NewPostgresConn(cfg, log)
	if err != nil {
		exit(log, err)
	}

	// Migrator
	mg, err := mig.NewMigrator(cfg, log, "migrator", db)
	if err != nil {
		log.Error(err)
	}

	// Seeder
	sd, err := seed.NewSeeder(cfg, log, "seeder", db)
	if err != nil {
		log.Error(err)
	}

	// Mailer
	ml, err := fnd.NewSESMailer(cfg, log, "mailer")
	if err != nil {
		exit(log, err)
	}

	// Repos
	userRepo := repo.NewUserRepo(cfg, log, "user-repo", db)
	eventRepo := repo.NewEventRepo(cfg, log, "event-repo", db)
	ticketRepo := repo.NewTicketRepo(cfg, log, "ticket-repo", db)

	// Core service
	svc := svc.NewService(cfg, log, "core-service", db)

	// Service dependencies
	svc.Mailer = ml
	svc.UserRepo = userRepo
	svc.EventRepo = eventRepo
	svc.TicketRepo = ticketRepo

	// App dependencies
	a.Migrator = mg
	a.Seeder = sd
	a.SetService(svc)

	// Init service
	err = a.Init()
	if err != nil {
		exit(log, err)
	}

	// Start service
	err = a.Start()
	if err != nil {
		exit(log, err)
	}

	log.Error(err, fmt.Sprintf("%s service stoped", appName))
}

func exit(log fnd.Logger, err error) {
	log.Error(err)
	os.Exit(1)
}

func checkStopEvents(ctx context.Context, cancel context.CancelFunc) {
	go checkSigterm(cancel)
	go checkCancel(ctx)
}

func checkSigterm(cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	cancel()
}

func checkCancel(ctx context.Context) {
	<-ctx.Done()
	a.Stop()
	os.Exit(1)
}
