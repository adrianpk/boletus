package pg

import (
	"errors"
	"time"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/foundation/db"
	"github.com/adrianpk/foundation/db/pg"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type (
	EventRepo struct {
		Cfg  *fnd.Config
		Log  fnd.Logger
		Name string
	}
)

type (
	eventRow struct {
		mutable bool
		model   model.Event
	}
)

var (
	t1, _ = time.Parse(time.RFC3339, "2020-01-02T19:00:00Z00:00")
	t2, _ = time.Parse(time.RFC3339, "2020-01-08T20:00:00Z00:00")

	rockParty = model.Event{
		Identification: fnd.Identification{
			ID:   fnd.ToUUID("03bc0622-4612-4490-9818-c50b2fd2340"),
			Slug: db.ToNullString("rockparty-c0f432eb921d"),
		},
		Name:        db.ToNullString("Rock Party in Wrocław"),
		Description: db.ToNullString(""),
		Place:       db.ToNullString("Wrocław"),
		ScheduledAt: pg.ToNullTime(t1), // 1st February 19:00, 2020
	}

	jesseCook = model.Event{
		Identification: fnd.Identification{
			ID:   fnd.ToUUID("f2553fc3-3c63-4adb-bc92-be8c66aecd19"),
			Slug: db.ToNullString("jessecook-3f7080400ba1"),
		},
		Name:        db.ToNullString("Jesse Cook"),
		Description: db.ToNullString(""),
		Place:       db.ToNullString("National Forum of Music, Wrocław"),
		ScheduledAt: pg.ToNullTime(t2), // 8st February 20:00, 2020
	}

	eventsTable = map[uuid.UUID]eventRow{
		rockParty.ID: eventRow{mutable: false, model: rockParty},
		jesseCook.ID: eventRow{mutable: false, model: jesseCook},
	}
)

func (tr *EventRepo) Create(event *model.Event, tx ...*sqlx.Tx) error {
	_, ok := eventsTable[event.ID]
	if ok {
		errors.New("duplicate key violates unique constraint")
	}

	if event.ID == uuid.Nil {
		errors.New("Non valid primary key")
	}

	eventsTable[event.ID] = eventRow{
		mutable: true,
		model:   *event,
	}

	return nil
}

func (tr *EventRepo) GetAll() (events []model.Event, err error) {
	size := len(eventsTable)
	out := make([]model.Event, size)
	for _, row := range eventsTable {
		out = append(out, row.model)
	}
	return out, nil
}

func (tr *EventRepo) Get(id uuid.UUID) (event model.Event, err error) {
	for _, row := range eventsTable {
		if id == row.model.ID {
			return row.model, nil
		}
	}
	return model.Event{}, nil
}

func (tr *EventRepo) GetBySlug(slug string) (event model.Event, err error) {
	for _, row := range eventsTable {
		if slug == row.model.Slug.String {
			return row.model, nil
		}
	}
	return model.Event{}, nil
}

func (tr *EventRepo) GetByName(name string) (model.Event, error) {
	for _, row := range eventsTable {
		if name == row.model.Name.String {
			return row.model, nil
		}
	}
	return model.Event{}, nil
}

func (tr *EventRepo) Update(event *model.Event, tx ...*sqlx.Tx) error {
	for _, row := range eventsTable {
		if event.ID == row.model.ID {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			eventsTable[event.ID] = eventRow{
				mutable: true,
				model:   *event,
			}
			return nil
		}
	}
	return errors.New("no records updated")
}

func (tr *EventRepo) Delete(id uuid.UUID, tx ...*sqlx.Tx) error {
	for _, row := range eventsTable {
		if id == row.model.ID {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			delete(eventsTable, id)
			return nil
		}
	}
	return errors.New("no records deleted")
}

func (tr *EventRepo) DeleteySlug(slug string, tx ...*sqlx.Tx) error {
	for _, row := range eventsTable {
		if slug == row.model.Slug.String {
			if !row.mutable {
				return errors.New("non mutable row")
			}

			delete(eventsTable, row.model.ID)
			return nil
		}
	}
	return errors.New("no records deleted")
}
