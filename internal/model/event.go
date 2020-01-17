package model

import (
	"database/sql"
	"fmt"
	"time"

	fnd "github.com/adrianpk/foundation"
	"github.com/adrianpk/foundation/db"
	"github.com/adrianpk/foundation/db/pg"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type (
	// Event model
	Event struct {
		fnd.Identification
		Name        sql.NullString `db:"name" json:"name" schema:"name"`
		Description sql.NullString `db:"description" json:"description" schema:"description"`
		Place       sql.NullString `db:"place" json:"place" schema:"place"`
		ScheduledAt pq.NullTime    `db:"scheduled_at" json:"scheduled_at" schema:"scheduled-at"`
		Locale      sql.NullString `db:"locale" json:"-" schema:"-"`
		BaseTZ      sql.NullString `db:"base_tz" json:"-" schema:"-"`
		IsActive    sql.NullBool   `db:"is_active" json:"-" schema:"-"`
		IsDeleted   sql.NullBool   `db:"is_deleted" json:"-" schema:"-"`
		fnd.Audit
	}
)

type (
	EventForm struct {
		Slug        string `json:"slug" schema:"slug"`
		Name        string `json:"name" schema:"name"`
		Description string `json:"description" schema:"description"`
		Place       string `json:"place" schema:"place"`
		ScheduledAt string `json:"scheduledAt" schema:"scheduled-at"`
		Year        string `json:"year" schema:"year"`
		Month       string `json:"mont" schema:"month"`
		Day         string `json:"day" schema:"day"`
		Hour        string `json:"hours" schema:"hour"`
		Minute      string `json:"minute" schema:"minute"`
		BaseTZ      string `json:"baseTZ" schema:"base-tz"`
		IsNew       bool   `json:"-" schema:"-"`
	}
)

func ToEventFormList(events []Event) (fs []EventForm) {
	for _, m := range events {
		fs = append(fs, m.ToForm())
	}
	return fs
}

// SetCreateValues sets de ID and slug.
func (event *Event) SetCreateValues() error {
	// Set create values only only if they were not previously
	if event.Identification.ID == uuid.Nil ||
		event.Identification.Slug.String == "" {
		pfx := event.Name.String
		event.Identification.SetCreateValues(pfx)
		event.Audit.SetCreateValues()
	}
	return nil
}

// SetUpdateValues
func (event *Event) SetUpdateValues() error {
	event.Audit.SetUpdateValues()
	return nil
}

// Match condition for
func (event *Event) Match(tc *Event) bool {
	r := event.Identification.Match(tc.Identification) &&
		event.Name == tc.Name &&
		event.Description == tc.Description &&
		event.Place == tc.Place &&
		event.ScheduledAt == tc.ScheduledAt &&
		event.BaseTZ == tc.BaseTZ
	return r
}

// ToForm lets convert a model to its associated form type.
// This convertion step could be avoided since gorilla schema allows
// to register custom decoders and in this case we need one because
// struct properties are not using Go standard types but their sql
// null conterpart types. As long as is relatively simple, ergonomic
// and could be easily implemented with generators I prefer to avoid
// the use of reflection.
// TODO: Move type conversion to web packate.
func (event *Event) ToForm() EventForm {
	date := destructureDate(event.ScheduledAt.Time)
	return EventForm{
		Slug:        event.Slug.String,
		Name:        event.Name.String,
		Description: event.Description.String,
		Place:       event.Place.String,
		ScheduledAt: event.ScheduledAt.Time.Format(time.RFC3339),
		Year:        date["year"],
		Month:       date["month"],
		Day:         date["day"],
		Hour:        date["hour"],
		Minute:      date["minute"],
		BaseTZ:      date["timezone"],
		IsNew:       event.IsNew(),
	}
}

// ToModel lets covert a form type to its associated model.
func (eventForm *EventForm) ToModel() Event {
	return Event{
		Identification: fnd.Identification{
			Slug: db.ToNullString(eventForm.Slug),
		},
		Name:        db.ToNullString(eventForm.Name),
		Description: db.ToNullString(eventForm.Description),
		Place:       db.ToNullString(eventForm.Place),
		ScheduledAt: pg.ToNullTime(eventForm.ScheduledAtParsed()),
	}
}

func (eventForm *EventForm) GetSlug() string {
	return eventForm.Slug
}

func (eventForm *EventForm) ScheduledAtDateMap() map[string]string {
	return map[string]string{
		"year":     eventForm.Year,
		"month":    eventForm.Month,
		"day":      eventForm.Day,
		"minute":   eventForm.Minute,
		"timezone": "UTC",
	}
}

func (eventForm *EventForm) ScheduledAtParsed() time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	dateStr := normalizeDate(eventForm.ScheduledAtDateMap())

	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}
	}

	return t
}

func destructureDate(date time.Time) map[string]string {
	return map[string]string{
		"year":     toStr(date.Year()),
		"month":    toStr(int(date.Month())),
		"day":      toStr(date.Day()),
		"hour":     toStr(date.Hour()),
		"minute":   toStr(date.Minute()),
		"timezone": "UTC",
	}
}

func normalizeDate(dm map[string]string) string {
	y := dm["year"]
	m := dm["month"]
	d := dm["day"]
	h := dm["hour"]
	mn := dm["minute"]
	return fmt.Sprintf("%s-%s-%sT%s:%s:00.000Z", y, m, d, h, mn)
}

func toStr(n int) string {
	return fmt.Sprintf("%d", n)
}
