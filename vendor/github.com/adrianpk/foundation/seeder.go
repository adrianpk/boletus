package foundation

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/jmoiron/sqlx"
)

type (
	SeederIF interface {
		Seed() error
	}
)

type (
	// Fx type alias
	SeedFx = func() error

	// Seeder struct.
	Seeder struct {
		*Worker
		DB     *sqlx.DB
		schema string
		dbName string
		seeds  []*Seed
	}

	// Exec interface.
	SeedExec interface {
		Config(seed SeedFx)
		GetSeed() (up SeedFx)
		SetTx(tx *sqlx.Tx)
		GetTx() (tx *sqlx.Tx)
	}

	// Seed struct.
	Seed struct {
		Executor SeedExec
	}
)

// NewSeeder.
func NewSeeder(cfg *Config, log Logger, name string, db *sqlx.DB) *Seeder {
	m := &Seeder{
		Worker: NewWorker(cfg, log, name),
		DB:     db,
		schema: cfg.ValOrDef("pg.schema", "public"),
		dbName: cfg.ValOrDef("pg.database", "granica_test"),
	}

	return m
}

// pgConnect to postgre database
// mainly user to create and drop app database.
func (m *Seeder) pgConnect() error {
	db, err := sqlx.Open("postgres", m.pgDbURL())
	if err != nil {
		log.Printf("Connection error: %s\n", err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Connection error: %s", err.Error())
		return err
	}

	m.DB = db
	return nil
}

// GetTx returns a new transaction from seeder connection.
func (s *Seeder) GetTx() *sqlx.Tx {
	return s.DB.MustBegin()
}

func (s *Seeder) AddSeed(e SeedExec) {
	s.seeds = append(s.seeds, &Seed{Executor: e})
}

func (s *Seeder) Seed() error {
	for _, mg := range s.seeds {
		exec := mg.Executor
		fn := getFxName(exec.GetSeed())

		// Get a new Tx from seeder
		tx := s.GetTx()
		// Pass Tx to the executor
		exec.SetTx(tx)

		// Execute migration
		values := reflect.ValueOf(exec).MethodByName(fn).Call([]reflect.Value{})

		// Read error
		err, ok := values[0].Interface().(error)
		if !ok && err != nil {
			log.Printf("Migration not executed: %s\n", fn) // TODO: Remove log
			log.Printf("Err  %+v' of type %T\n", err, err) // TODO: Remove log.
			msg := fmt.Sprintf("cannot run migration '%s': %s", fn, err.Error())
			tx.Rollback()
			return errors.New(msg)
		}

		err = tx.Commit()
		if err != nil {
			msg := fmt.Sprintf("Cannot update migrations table: %s\n", err.Error())
			log.Printf("Commit error: %s", msg)
			tx.Rollback()
			return errors.New(msg)
		}

		log.Printf("Migration executed: %s\n", fn)
	}

	return nil
}

func (m *Seeder) dbURL() string {
	host := m.Cfg.ValOrDef("pg.host", "localhost")
	port := m.Cfg.ValOrDef("pg.port", "5432")
	m.schema = m.Cfg.ValOrDef("pg.schema", "public")
	m.dbName = m.Cfg.ValOrDef("pg.database", "granica_test_d1x89s0l")
	user := m.Cfg.ValOrDef("pg.user", "granica")
	pass := m.Cfg.ValOrDef("pg.password", "granica")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbName=%s sslmode=disable search_path=%s", host, port, user, pass, m.dbName, m.schema)
}

func (m *Seeder) pgDbURL() string {
	host := m.Cfg.ValOrDef("pg.host", "localhost")
	port := m.Cfg.ValOrDef("pg.port", "5432")
	schema := "public"
	db := "postgres"
	user := m.Cfg.ValOrDef("pg.user", "granica")
	pass := m.Cfg.ValOrDef("pg.password", "granica")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbName=%s sslmode=disable search_path=%s", host, port, user, pass, db, schema)
}
