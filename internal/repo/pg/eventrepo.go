package pg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	EventRepo struct {
		Cfg  *fnd.Config
		Log  fnd.Logger
		Name string
		DB   *sqlx.DB
	}
)

func NewEventRepo(cfg *fnd.Config, log fnd.Logger, name string, db *sqlx.DB) *EventRepo {
	return &EventRepo{
		Cfg:  cfg,
		Log:  log,
		Name: name,
		DB:   db,
	}
}

// Create a event
func (ur *EventRepo) Create(event *model.Event, tx ...*sqlx.Tx) error {
	st := `INSERT INTO events (id, slug, name, description, place, scheduled_at, locale, base_tz, is_active, is_deleted, created_by_id, updated_by_id, created_at, updated_at)
VALUES (:id, :slug, :name, :description, :place, :scheduled_at, :locale, :base_tz, :current_tz, :is_active, :is_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at)`

	// Create a local transaction if it is not passed as argument.
	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	// Don't wait for repo to setup this values.
	// We want event ID to event as account owner ID.
	event.SetCreateValues()

	_, err = t.NamedExec(st, event)
	if err != nil {
		return err
	}

	// Commit on local transactions
	if local {
		return t.Commit()
	}

	return nil
}

// GetAll events from
func (ur *EventRepo) GetAll() (events []model.Event, err error) {
	st := `SELECT * FROM events WHERE is_deleted IS NULL OR NOT is_deleted;`

	err = ur.DB.Select(&events, st)
	if err != nil {
		return events, err
	}

	return events, err
}

// Get event by ID.
func (ur *EventRepo) Get(id uuid.UUID) (event model.Event, err error) {
	st := `SELECT * FROM events WHERE id = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, id.String())

	err = ur.DB.Get(&event, st)
	if err != nil {
		return event, err
	}

	return event, err
}

// GetBySlug event from repo by slug.
func (ur *EventRepo) GetBySlug(slug string) (event model.Event, err error) {
	st := `SELECT * FROM events WHERE slug = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, slug)

	err = ur.DB.Get(&event, st)

	return event, err
}

// GetByName event from repo by eventname.
func (ur *EventRepo) GetByName(name string) (model.Event, error) {
	var event model.Event

	st := `SELECT * FROM events WHERE name = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, name)

	err := ur.DB.Get(&event, st)

	return event, err
}

// Update event data in
func (ur *EventRepo) Update(event *model.Event, tx ...*sqlx.Tx) error {
	ref, err := ur.Get(event.ID)
	if err != nil {
		return fmt.Errorf("cannot retrieve reference event: %s", err.Error())
	}

	event.SetUpdateValues()

	var st strings.Builder
	pcu := false // previous column updated?

	st.WriteString("UPDATE events SET ")

	if event.Name != ref.Name {
		st.WriteString(fnd.SQLStrUpd("name", "name"))
		pcu = true
	}

	if event.Description != ref.Description {
		st.WriteString(fnd.SQLStrUpd("description", "description"))
		pcu = true
	}

	if event.Place != ref.Place {
		st.WriteString(fnd.SQLStrUpd("place", "place"))
		pcu = true
	}

	if event.ScheduledAt != ref.ScheduledAt {
		st.WriteString(fnd.SQLStrUpd("scheduled_at", "scheduled_at"))
		pcu = true
	}

	st.WriteString(" ")
	st.WriteString(fnd.SQLWhereID(ref.ID.String()))
	st.WriteString(";")

	//fmt.Println(st.String())

	if pcu == false {
		return errors.New("no fields to update")
	}

	// Create a local transaction if it is not passed as argument.
	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.NamedExec(st.String(), event)

	if local {
		ur.Log.Info("Transaction created by repo: committing")
		return t.Commit()
	}

	return nil
}

// Delete event from repo by ID.
func (ur *EventRepo) Delete(id uuid.UUID, tx ...*sqlx.Tx) error {
	st := `DELETE FROM events WHERE id = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, id)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// DeleteBySlug:w event from repo by slug.
func (ur *EventRepo) DeleteBySlug(slug string, tx ...*sqlx.Tx) error {
	st := `DELETE FROM events WHERE slug = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, slug)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}

	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// DeleteByeventname event from repo by eventname.
func (ur *EventRepo) DeleteByEventname(eventname string, tx ...*sqlx.Tx) error {
	st := `DELETE FROM events WHERE eventname = '%s' AND (is_deleted IS NULL OR NOT is_deleted);`
	st = fmt.Sprintf(st, eventname)

	t, local, err := ur.getTx(tx)
	if err != nil {
		return err
	}
	_, err = t.Exec(st)

	if local {
		return t.Commit()
	}

	return err
}

// GetBySlug event from repo by slug token.
func (ur *EventRepo) GetBySlugAndToken(slug, token string) (model.Event, error) {
	var event model.Event

	st := `SELECT * FROM events WHERE slug = '%s' AND confirmation_token = '%s' AND (is_deleted IS NULL OR NOT is_deleted) LIMIT 1;`
	st = fmt.Sprintf(st, slug, token)

	err := ur.DB.Get(&event, st)

	return event, err
}

func (ur *EventRepo) newTx() (tx *sqlx.Tx, err error) {
	tx, err = ur.DB.Beginx()
	if err != nil {
		return tx, err
	}

	return tx, err
}

func (ur *EventRepo) getTx(txs []*sqlx.Tx) (tx *sqlx.Tx, local bool, err error) {
	// Create a new transaction if its no passed as argument.
	if len(txs) > 0 {
		return txs[0], false, nil
	}

	tx, err = ur.DB.Beginx()
	if err != nil {
		return tx, true, err
	}

	return tx, true, err
}
