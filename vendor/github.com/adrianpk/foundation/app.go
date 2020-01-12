package kabestan

import (
	"context"
)

type (
	App struct {
		Cfg      *Config
		Log      Logger
		cancel   context.CancelFunc
		Name     string
		Revision string
		// Health
		ready bool
		alive bool
		// Migrator
		Migrator MigratorIF
		// Seeder
		Seeder SeederIF
		// Routers
		WebRouter      *Router
		JSONRESTRouter *Router
	}
)

func NewApp(cfg *Config, log Logger, name string) *App {
	name = genName(name, "app")

	return &App{
		Cfg:  cfg,
		Log:  log,
		Name: name,
	}
}
