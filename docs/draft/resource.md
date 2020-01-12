# Notes

Steps needed to create a new resource

# Account

## Create a resource generator input.
  * File: `assets/gen/account.yaml`
  * As a reference, not needed right now, generator not finished.
  * After the generator is finished this guide will be not needed anymore.

## Create a model

  * File: `internal/model/account.go`
  * This file also includes de `accountForm` struct.

## Create a migration

  * File: `internal/mig/00003createaccountstable.go`

## Add migration

  * File: `internal/mig/mig.go`

```go
// GetMigrator configured.
func (m *Migrator) addSteps() {
	// Migrations
	// Enable Postgis
	s := &step{}
	s.Config(s.EnablePostgis, s.DropPostgis)
	m.AddMigration(s)

	// CreateUsersTable
	s = &step{}
	s.Config(s.CreateUsersTable, s.DropUsersTable)
	m.AddMigration(s)

	// CreateAccountsTable <-- Something like this block
	s = &step{}
	s.Config(s.CreateAccountsTable, s.DropAccountsTable)
	m.AddMigration(s)
}
```

## Create resource repo interface

  * File: `internal/repo/accountrepo.go`

## Create resource repo implementation

  * PostgreSQL
  * File: `internal/repo/pg/accountrepo.go`
  * Implement al interface methods
  * Tests: `internal/repo/pg/accountrepo_test.go`

  * Volatile
  * Location: `internal/repo/mem/accountrepo.go`
  * Implement al interface methods
  * Only if needed

## Create resource router

  * File: `internal/app/accountrouter.go`

## Add router to parent router

  * Edit: `internal/app/router.go`

```go
func (app *App) NewWebRouter() *fnd.Router {
	rt := app.makeWebHomeRouter(app.Cfg, app.Log)
	app.addWebAuthRouter(rt)
	app.addWebUserRouter(rt)
	app.addWebAccountRouter(rt) <- Something like this
	return rt
}
```

## Create resource web endpoint

  * File: `internal/web/accounthandler.go`

## Create resource path routes

  * File: `internal/web/accountpath.go`

## Create resource service

  * File: `internal/svc/accountsvc.go`

## Create resource validator

  * File: `internal/svc/accountval.go`

## Add resource repo dependency to service

  * Edit: `internal/svc/svc.go`

```go
type (
	Service struct {
		*fnd.Service
		UserRepo    repo.UserRepo
		AccountRepo repo.AccountRepo <- Something like this
	}
)
```
## Create templates

  * Files: `assets/web/embed/template/account`

## Edit content and translations

  * Files: `assets/web/embed/i18n`


