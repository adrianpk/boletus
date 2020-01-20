package svc_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	svc "github.com/adrianpk/boletus/internal/app/svc"
	"github.com/adrianpk/boletus/internal/mig"
	"github.com/adrianpk/boletus/internal/repo/pg"
	"github.com/adrianpk/boletus/internal/seed"
	fnd "github.com/adrianpk/foundation"

	//"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
)

var (
	cfg *fnd.Config
	log fnd.Logger
	db  *sqlx.DB
)

var (
	mg *mig.Migrator
)

const (
	userSlug  = "lauriem-000000000004"
	eventSlug = "rockpartyinwrocław-000000000001"
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
func TestTicketSummary(t *testing.T) {
	// Create Service
	s := newService()

	// Invoke function to be tested
	tss, err := s.TicketSummary(eventSlug)
	if err != nil {
		t.Error(err)
		t.Error("Error calling TicketSummary")
	}

	//t.Log(spew.Sdump(tss))

	if len(tss) < 1 {
		t.Error("Empty result")
	}

	for _, ts := range tss {
		tp := ts.Type.String
		cy := ts.Currency.String

		ok0 := ts.EventSlug.String == "rockpartyinwrocław-000000000001"
		ok1 := tp == "standard" || tp == "golden-circle" || tp == "couples" || tp == "preemptive"
		ok2 := cy == "EUR" || cy == "USD" || cy == "PLN" || cy == "CZK" || cy == "RUB"

		if !(ok0 && ok1 && ok2) {
			t.Error("Does not match expected result")
			return
		}

		ok3 := len(ts.Prices) > 0

		if !ok3 {
			t.Error("Cannot do currency conversion")
			return
		}
	}
}

// TestPreBookGoldenCircleTickets test ticket reservation
// for standard tickets.
func TestPreBookStandardTickets(t *testing.T) {
	// Create Service
	s := newService()

	// Invoke function to be tested
	ts, err := s.PreBookTickets(eventSlug, "standard", 4, userSlug)
	if err != nil {
		t.Error(err)
		t.Error("Error calling PreBookTicket")
	}

	//t.Log(spew.Sdump(ts))

	if len(ts) != 4 {
		t.Error("Reserved tickets qty. does not match expected.")
	}

	zeroT := time.Time{}

	for _, tt := range ts {
		sl := tt.Slug.String          // rockpartyinwrocław-...
		nm := tt.Name.String          // "Rock Party in Wrocław"
		tp := tt.Type.String          // "standard"
		sr := tt.Serie.String         // "A"
		pr := tt.Price.Int32          // 30000
		cy := tt.Currency.String      // "EUR"
		ri := tt.ReservationID.String // != ""
		rb := tt.ReservedBy.String    // "00000000-0000-0000-0000-000000000004"
		ra := tt.ReservedAt.Time      // != time.Time{}
		st := tt.Status.String        // "reserved"

		ok0 := strings.Index(sl, "rockpartyinwrocław-") == 0
		ok1 := nm == "Rock Party in Wrocław"
		ok2 := tp == "standard"
		ok3 := sr == "A"
		ok4 := pr == 30000
		ok5 := cy == "EUR"
		ok6 := ri != ""
		ok7 := rb == "00000000-0000-0000-0000-000000000004"
		ok8 := ra != zeroT
		ok9 := st == "reserved"

		//fmt.Printf("ok0: %t, ok1: %t, ok2: %t, ok3: %t, ok4: %t, ok5: %t, ok6: %t, ok7: %t, ok8: %t, ok9: %t\n\n",
		//ok0, ok1, ok2, ok3, ok4, ok5, ok6, ok7, ok8, ok9)

		if !(ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8 && ok9) {
			t.Error("Does not match expected result")
			return
		}
	}
}

// TestPrBookGoldenCircleiTickets test ticket reservation for golden circle tickets.
func TestPreBookGoldenCircleTickets(t *testing.T) {
	// Create Service
	s := newService()

	// Invoke function to be tested
	ts, err := s.PreBookTickets(eventSlug, "golden-circle", 4, userSlug)
	if err != nil {
		t.Error(err)
		t.Error("Error calling PreBookTicket")
	}

	//t.Log(spew.Sdump(ts))
	fmt.Println(len(ts))
	if len(ts) != 20 {
		t.Error("Reserved tickets qty. does not match expected.")
	}

	zeroT := time.Time{}

	for _, tt := range ts {
		sl := tt.Slug.String          // rockpartyinwrocław-...
		nm := tt.Name.String          // "Rock Party in Wrocław"
		tp := tt.Type.String          // "golden-circle"
		sr := tt.Serie.String         // "A"
		pr := tt.Price.Int32          // 75000
		cy := tt.Currency.String      // "EUR"
		ri := tt.ReservationID.String // != ""
		rb := tt.ReservedBy.String    // "00000000-0000-0000-0000-000000000004"
		ra := tt.ReservedAt.Time      // != time.Time{}
		st := tt.Status.String        // "reserved"

		ok0 := strings.Index(sl, "rockpartyinwrocław-") == 0
		ok1 := nm == "Rock Party in Wrocław"
		ok2 := tp == "golden-circle"
		ok3 := sr == "A"
		ok4 := pr == 75000
		ok5 := cy == "EUR"
		ok6 := ri != ""
		ok7 := rb == "00000000-0000-0000-0000-000000000004"
		ok8 := ra != zeroT
		ok9 := st == "reserved"

		//fmt.Printf("ok0: %t, ok1: %t, ok2: %t, ok3: %t, ok4: %t, ok5: %t, ok6: %t, ok7: %t, ok8: %t, ok9: %t\n\n",
		//ok0, ok1, ok2, ok3, ok4, ok5, ok6, ok7, ok8, ok9)

		if !(ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8 && ok9) {
			t.Error("Does not match expected result")
			return
		}
	}
}

// TestPreBookCouplesTickets test ticket reservation
// for couples tickets.
func TestPreBookCouplesTickets(t *testing.T) {
	// Create Service
	s := newService()

	// Invoke function to be tested
	ts, err := s.PreBookTickets(eventSlug, "couples", 3, userSlug)
	if err != nil {
		t.Error(err)
		t.Error("Error calling PreBookTicket")
	}

	//t.Log(spew.Sdump(ts))

	if len(ts) != 4 {
		t.Error("Reserved tickets qty. does not match expected.")
	}

	zeroT := time.Time{}

	for _, tt := range ts {
		sl := tt.Slug.String          // rockpartyinwrocław-...
		nm := tt.Name.String          // "Rock Party in Wrocław"
		tp := tt.Type.String          // "couples"
		sr := tt.Serie.String         // "A"
		pr := tt.Price.Int32          // 30000
		cy := tt.Currency.String      // "EUR"
		ri := tt.ReservationID.String // != ""
		rb := tt.ReservedBy.String    // "00000000-0000-0000-0000-000000000004"
		ra := tt.ReservedAt.Time      // != time.Time{}
		st := tt.Status.String        // "reserved"

		ok0 := strings.Index(sl, "rockpartyinwrocław-") == 0
		ok1 := nm == "Rock Party in Wrocław"
		ok2 := tp == "couples"
		ok3 := sr == "A"
		ok4 := pr == 20000
		ok5 := cy == "EUR"
		ok6 := ri != ""
		ok7 := rb == "00000000-0000-0000-0000-000000000004"
		ok8 := ra != zeroT
		ok9 := st == "reserved"

		//fmt.Printf("ok0: %t, ok1: %t, ok2: %t, ok3: %t, ok4: %t, ok5: %t, ok6: %t, ok7: %t, ok8: %t, ok9: %t\n\n",
		//ok0, ok1, ok2, ok3, ok4, ok5, ok6, ok7, ok8, ok9)

		if !(ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8 && ok9) {
			t.Error("Does not match expected result")
			return
		}
	}
}

// Misc

func setup() (err error) {
	cfg = testConfig()
	log = testLogger()
	db, err = testDB(cfg)
	if err != nil {
		return err
	}

	// Prepare database
	mg, err = mig.NewMigrator(cfg, log, "migrator", db)
	if err != nil {
		return err
	}

	mg.Migrate()

	// Seed data
	sd, err := seed.NewSeeder(cfg, log, "seeder", db)
	if err != nil {
		return err
	}

	sd.Seed()

	return nil
}

func newService() *svc.Service {
	// Create ticket repo
	r := pg.NewTicketRepo(cfg, log, "ticket-repo", db)

	// Create Service instance
	s := svc.NewService(cfg, log, "service-instance", db)

	s.Init()

	s.TicketRepo = r

	return s
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

	db := conn

	return db, nil
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
