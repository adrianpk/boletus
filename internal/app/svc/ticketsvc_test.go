package svc

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"gitlab.com/adrianpk/boletus/internal/mig"
	"gitlab.com/adrianpk/boletus/internal/model"
	"gitlab.com/adrianpk/boletus/internal/repo"
)

var (
	userDataValid = map[string]string{
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
	mgr := setup()
	code := m.Run()
	teardown(mgr)
	os.Exit(code)
}

// TestTicketSummary verifies Test.
// Signature: TicketSummary(eventSlug string) (tss []model.TicketSummary, err error)
func TestTicketSummary(t *testing.T) {
	cfg := testConfig()
	log := testLogger()
	db := testDb()

	// Create Service instance
	svc := NewService(cfg, log, "service-test-instance", db)
	svc.Init()

	// Invoke function to be tested
	//svc.IndexTicket(...)

	notImplementd = true
	if notImplemented {
		t.Error("TestTicketSummary not implemented")
	}
}

func setup() *fnd.Migrator {
	// Config for tests
	cfg := testConfig()

	// Prepare database
	m := mig.GetMigrator(testConfig())
	m.Migrate()

	// Seed data
	//s := seed.GetSeeder(testConfig())
	//m.Seed()
	return m
}

func teardown(m *fnd.Migrator) {
	m.RollbackAll()
}

func testConfig() *Config {
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

func testLogger() *Logger {
	return fnd.NewDevLogger(0, "boletus", "n/a")
}

// getTestDB returns a connection to test DB.
func testDB() (*sqlx.DB, error) {
	cfg := testConfig()
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
func dbURL(cfg *Config) string {
	host := cfg.ValOrDef("pg.host", "localhost")
	port := cfg.ValOrDef("pg.port", "5432")
	schema := cfg.ValOrDef("pg.schema", "public")
	db := cfg.ValOrDef("pg.database", "boletus_test")
	user := cfg.ValOrDef("pg.user", "boletus")
	pass := cfg.ValOrDef("pg.password", "boletus")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, pass, db, schema)
}
