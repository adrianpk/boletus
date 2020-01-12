package kabestan

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	MigratorIF interface {
		Migrate() error
		Rollback(n ...int) error
		RollbackAll() error
		Reset() error
	}
)

type (
	// Fx type alias
	MigFx = func() error

	// Migrator struct.
	Migrator struct {
		*Worker
		DB     *sqlx.DB
		schema string
		dbName string
		migs   []*Migration
	}

	// Exec interface.
	Exec interface {
		Config(up MigFx, down MigFx)
		GetName() (name string)
		GetUp() (up MigFx)
		GetDown() (down MigFx)
		SetTx(tx *sqlx.Tx)
		GetTx() (tx *sqlx.Tx)
	}

	// Migration struct.
	Migration struct {
		Order    int
		Executor Exec
	}

	migRecord struct {
		ID        uuid.UUID `db:"id" json:"id"`
		Name      string    `db:"name" json:"name"`
		UpFx      string    `db:"up_fx" json:"upFx"`
		DownFx    string    `db:"down_fx" json:"downFx"`
		IsApplied bool      `db:"is_applied" json:"isApplied"`
		CreatedAt time.Time `db:"created_at" json:"createdAt"`
	}
)

const (
	pgMigrationsTable = "migrations"

	pgCreateDbSt = `
		CREATE DATABASE %s;`

	pgDropDbSt = `
		DROP DATABASE %s;`

	pgCreateMigrationsSt = `CREATE TABLE %s.%s (
		id UUID PRIMARY KEY,
		name VARCHAR(64),
		up_fx VARCHAR(64),
		down_fx VARCHAR(64),
 		is_applied BOOLEAN,
		created_at TIMESTAMP
	);`

	pgDropMigrationsSt = `DROP TABLE %s.%s;`

	pgSelMigrationSt = `SELECT is_applied FROM %s.%s WHERE name = '%s' and is_applied = true`

	pgRecMigrationSt = `INSERT INTO %s.%s (id, name, up_fx, down_fx, is_applied, created_at)
		VALUES (:id, :name, :up_fx, :down_fx, :is_applied, :created_at);`

	pgDelMigrationSt = `DELETE FROM %s.%s WHERE name = '%s' and is_applied = true`
)

// NewMigrator.
func NewMigrator(cfg *Config, log Logger, name string, db *sqlx.DB) *Migrator {
	m := &Migrator{
		Worker: NewWorker(cfg, log, name),
		DB:     db,
		schema: cfg.ValOrDef("pg.schema", "public"),
		dbName: cfg.ValOrDef("pg.database", "granica_test"),
	}

	return m
}

// pgConnect to postgre database
// mainly user to create and drop app database.
func (m *Migrator) pgConnect() error {
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

// GetTx returns a new transaction from migrator connection.
func (m *Migrator) GetTx() *sqlx.Tx {
	return m.DB.MustBegin()
}

// PreSetup creates database
// and migrations table if needed.
func (m *Migrator) PreSetup() {
	if !m.dbExists() {
		m.CreateDb()
	}

	if !m.migTableExists() {
		m.createMigrationsTable()
	}
}

// dbExists returns true if migrator
// referenced database has been already created.
// Only for postgress at the moment.
func (m *Migrator) dbExists() bool {
	st := fmt.Sprintf(`select exists(
		SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('%s')
	);`, m.dbName)

	r, err := m.DB.Query(st)
	if err != nil {
		log.Printf("Error checking database: %s\n", err.Error())
		return false
	}

	for r.Next() {
		var exists sql.NullBool
		err = r.Scan(&exists)
		if err != nil {
			log.Printf("Cannot read query result: %s\n", err.Error())
			return false
		}
		return exists.Bool
	}
	return false
}

// migExists returns true if migrations table exists.
func (m *Migrator) migTableExists() bool {
	st := fmt.Sprintf(`SELECT EXISTS (
		SELECT 1
   	FROM   pg_catalog.pg_class c
   	JOIN   pg_catalog.pg_namespace n ON n.oid = c.relnamespace
   	WHERE  n.nspname = '%s'
   	AND    c.relname = '%s'
   	AND    c.relkind = 'r'
	);`, m.schema, m.dbName)

	r, err := m.DB.Query(st)
	if err != nil {
		log.Printf("Error checking database: %s\n", err.Error())
		return false
	}

	for r.Next() {
		var exists sql.NullBool
		err = r.Scan(&exists)
		if err != nil {
			log.Printf("Cannot read query result: %s\n", err.Error())
			return false
		}

		return exists.Bool
	}
	return false
}

// CreateDb migration.
func (m *Migrator) CreateDb() (string, error) {
	m.CloseAppConns()
	st := fmt.Sprintf(pgCreateDbSt, m.dbName)

	_, err := m.DB.Exec(st)
	if err != nil {
		return m.dbName, err
	}

	return m.dbName, nil
}

// DropDb migration.
func (m *Migrator) DropDb() (string, error) {
	m.CloseAppConns()
	st := fmt.Sprintf(pgDropDbSt, m.dbName)

	_, err := m.DB.Exec(st)
	if err != nil {
		return m.dbName, err
	}

	return m.dbName, nil
}

// CreateDb migration.
func (m *Migrator) CloseAppConns() (string, error) {
	dbName := m.Cfg.ValOrDef("pg.database", "")
	st := `SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '%s';`
	st = fmt.Sprintf(st, dbName)

	_, err := m.DB.Exec(st)
	if err != nil {
		return m.dbName, err
	}

	return m.dbName, nil
}

// DropDb migration.
func (m *Migrator) createMigrationsTable() (string, error) {
	tx := m.GetTx()

	st := fmt.Sprintf(pgCreateMigrationsSt, m.schema, pgMigrationsTable)

	_, err := tx.Exec(st)
	if err != nil {
		return pgMigrationsTable, err
	}

	return pgMigrationsTable, tx.Commit()
}

func (m *Migrator) AddMigration(e Exec) {
	m.migs = append(m.migs, &Migration{Executor: e})
}

func (m *Migrator) Migrate() error {
	m.PreSetup()

	for _, mg := range m.migs {
		exec := mg.Executor
		fn := getFxName(exec.GetUp())
		name := migName(fn)

		// Continue if already applied
		if !m.canApplyMigration(name) {
			log.Printf("Migration '%s' already applied.", name)
			continue
		}

		// Get a new Tx from migrator
		tx := m.GetTx()
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

		// Register migration
		err = m.recMigration(exec)

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

// Rollback migrations.
func (m *Migrator) Rollback(steps ...int) error {
	// Default to 1 step if no value is provided
	s := 1
	if len(steps) > 0 && steps[0] > 1 {
		s = steps[0]
	}

	// Default to max n° migration if steps is higher
	c := m.count()
	if s > c {
		s = c
	}

	m.rollback(s)
	return nil
}

// Rollback all migrations.
func (m *Migrator) RollbackAll() error {
	return m.rollback(m.count())
}

func (m *Migrator) rollback(steps int) error {
	count := m.count()
	stopAt := count - steps

	for i := count - 1; i >= stopAt; i-- {
		mg := m.migs[i]
		exec := mg.Executor
		fn := getFxName(exec.GetDown())
		// Migration name is associated to up migration
		name := migName(getFxName(exec.GetUp()))

		// Continue if already not rolledback
		if m.cancelRollback(name) {
			log.Printf("Rollback '%s' already executed.", name)
			continue
		}

		// Get a new Tx from migrator
		tx := m.GetTx()
		// Pass Tx to the executor
		exec.SetTx(tx)

		// Execute rollback
		values := reflect.ValueOf(exec).MethodByName(fn).Call([]reflect.Value{})

		// Read error
		err, ok := values[0].Interface().(error)
		if !ok && err != nil {
			log.Printf("Rollback not executed: %s\n", fn)
			log.Printf("Err '%+v' of type %T", err, err)
		}

		// Remove migration record.
		err = m.delMigration(exec)

		err = tx.Commit()
		if err != nil {
			msg := fmt.Sprintf("Cannot update migrations table: %s\n", err.Error())
			log.Printf("Commit error: %s", msg)
			tx.Rollback()
			return errors.New(msg)
		}

		log.Printf("Rollback executed: %s\n", fn)
	}

	return nil
}

func (m *Migrator) SoftReset() error {
	err := m.RollbackAll()
	if err != nil {
		log.Printf("Cannot rollback database: %s", err.Error())
		return err
	}

	err = m.Migrate()
	if err != nil {
		log.Printf("Cannot migrate database: %s", err.Error())
		return err
	}

	return nil
}

func (m *Migrator) Reset() error {
	_, err := m.DropDb()
	if err != nil {
		log.Printf("Drop database error: %s", err.Error())
		// Don't return maybe it was not created before.
	}

	_, err = m.CreateDb()
	if err != nil {
		log.Printf("Create database error: %s", err.Error())
		return err
	}

	err = m.Migrate()
	if err != nil {
		log.Printf("Drop database error: %s", err.Error())
		return err
	}

	return nil
}

func (m *Migrator) recMigration(e Exec) error {
	st := fmt.Sprintf(pgRecMigrationSt, m.schema, pgMigrationsTable)
	upFx := getFxName(e.GetUp())
	downFx := getFxName(e.GetDown())
	name := migName(upFx)
	log.Printf("%+s", upFx)

	_, err := e.GetTx().NamedExec(st, migRecord{
		ID:        uuid.NewV4(),
		Name:      name,
		UpFx:      upFx,
		DownFx:    downFx,
		IsApplied: true,
		CreatedAt: time.Now(),
	})

	if err != nil {
		msg := fmt.Sprintf("Cannot update migrations table: %s\n", err.Error())
		return errors.New(msg)
	}

	return nil
}

func (m *Migrator) cancelRollback(name string) bool {
	st := fmt.Sprintf(pgSelMigrationSt, m.schema, pgMigrationsTable, name)
	r, err := m.DB.Query(st)

	if err != nil {
		log.Printf("Cannot determine rollback status: %s\n", err.Error())
		return true
	}

	for r.Next() {
		var applied sql.NullBool
		err = r.Scan(&applied)
		if err != nil {
			log.Printf("Cannot determine migration status: %s\n", err.Error())
			return true
		}

		return !applied.Bool
	}

	return true
}

func (m *Migrator) canApplyMigration(name string) bool {
	st := fmt.Sprintf(pgSelMigrationSt, m.schema, pgMigrationsTable, name)
	r, err := m.DB.Query(st)

	if err != nil {
		log.Printf("Cannot determine migration status: %s\n", err.Error())
		return false
	}

	for r.Next() {
		var applied sql.NullBool
		err = r.Scan(&applied)
		if err != nil {
			log.Printf("Cannot determine migration status: %s\n", err.Error())
			return false
		}

		return !applied.Bool
	}

	return true
}

func (m *Migrator) delMigration(e Exec) error {
	name := migName(getFxName(e.GetUp()))
	st := fmt.Sprintf(pgDelMigrationSt, m.schema, pgMigrationsTable, name)
	_, err := e.GetTx().Exec(st)

	if err != nil {
		msg := fmt.Sprintf("Cannot update migrations table: %s\n", err.Error())
		return errors.New(msg)
	}

	return nil
}

func (m *Migrator) count() (last int) {
	return len(m.migs)
}

func (m *Migrator) last() (last int) {
	return m.count() - 1
}

func getFxName(i interface{}) string {
	n := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	t := strings.FieldsFunc(n, split)
	return t[len(t)-2]
}

func split(r rune) bool {
	return r == '.' || r == '-'
}

func migName(upFxName string) string {
	return toSnakeCase(upFxName)
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (m *Migrator) dbURL() string {
	host := m.Cfg.ValOrDef("pg.host", "localhost")
	port := m.Cfg.ValOrDef("pg.port", "5432")
	m.schema = m.Cfg.ValOrDef("pg.schema", "public")
	m.dbName = m.Cfg.ValOrDef("pg.database", "granica_test_d1x89s0l")
	user := m.Cfg.ValOrDef("pg.user", "granica")
	pass := m.Cfg.ValOrDef("pg.password", "granica")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbName=%s sslmode=disable search_path=%s", host, port, user, pass, m.dbName, m.schema)
}

func (m *Migrator) pgDbURL() string {
	host := m.Cfg.ValOrDef("pg.host", "localhost")
	port := m.Cfg.ValOrDef("pg.port", "5432")
	schema := "public"
	db := "postgres"
	user := m.Cfg.ValOrDef("pg.user", "granica")
	pass := m.Cfg.ValOrDef("pg.password", "granica")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbName=%s sslmode=disable search_path=%s", host, port, user, pass, db, schema)
}
