package svc_test

import (
	"fmt"
	"os"
	"testing"

	svc "github.com/adrianpk/boletus/internal/app/svc"
	"github.com/adrianpk/boletus/internal/mig"
	"github.com/adrianpk/boletus/internal/seed"
	fnd "github.com/adrianpk/foundation"
	"github.com/jmoiron/sqlx"
)

var (
	cfg *fnd.Config
	log fnd.Logger
	db  *sqlx.DB
	mg  *mig.Migrator
)

const (
	eventSlug = "rockpartyinwrocław-000000000001"
)

var (
	user = map[string]string{
		"username":          "username",
		"password":          "password",
		"email":             "username@mail.com",
		"emailConfirmation": "username@mail.com",
		"givenName":         "name",
		"middleNames":       "middles",
		"familyName":        "family",
	}
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}
	code := m.Run()
	teardown()
	os.Exit(code)
}

// TestTicketSummary verifies Test.
// Signature: TicketSummary(eventSlug string) (tss []model.TicketSummary, err error)
func TestTicketSummary(t *testing.T) {
	// Create Service instance
	s := svc.NewService(cfg, log, "service-test-instance", db)
	s.Init()

	// Invoke function to be tested
	s.IndexTickets()

	if true {
		t.Error("TestTicketSummary not implemented")
	}
}

func setup() (err error) {
	cfg = testConfig()
	log = testLogger()
	db, err := testDB(cfg)
	if err != nil {
		return err
	}

	// Prepare database
	mg, err = mig.NewMigrator(cfg, log, "test-migrator", db)
	if err != nil {
		return err
	}

	mg.Migrate()

	// Seed data
	sd, err := seed.NewSeeder(cfg, log, "test-seeder", db)
	if err != nil {
		return err
	}

	sd.Seed()
	return nil
}

func teardown() {
	mg.RollbackAll()
}

func testConfig() *fnd.Config {
	cfg := &fnd.Config{}
	values := map[string]string{
		"pg.host":               "localhost",
		"pg.port":               "5432",
		"pg.schema":             "public",
		"pg.database":           "boletus_test",
		"pg.user":               "boletus",
		"pg.password":           "boletus",
		"pg.backoff.maxentries": "3",
	}

	cfg.SetNamespace("grc")
	cfg.SetValues(values)
	return cfg
}

func testLogger() *fnd.Log {
	return fnd.NewDevLogger(0, "boletus", "n/a")
}

// getTestDB returns a connection to test DB.
func testDB(cfg *fnd.Config) (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", dbURL(cfg))
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// dbURL returns a Postgres connection string.
func dbURL(cfg *fnd.Config) string {
	host := cfg.ValOrDef("pg.host", "localhost")
	port := cfg.ValOrDef("pg.port", "5432")
	schema := cfg.ValOrDef("pg.schema", "public")
	db := cfg.ValOrDef("pg.database", "boletus_test")
	user := cfg.ValOrDef("pg.user", "boletus")
	pass := cfg.ValOrDef("pg.password", "boletus")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, pass, db, schema)
}
