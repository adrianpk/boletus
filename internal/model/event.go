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
		Role        sql.NullString `db:"role" json:"role" schema:"role"`
		Name        sql.NullString `db:"name" json:"name" schema:"name"`
		Description sql.NullString `db:"description" json:"description" schema:"description"`
		Place       sql.NullString `db:"place" json:"place" schema:"place"`
		ScheduledAt pq.NullTime    `db:"schedlued_at" json:"scheduled_at" schema:"scheduled-at"`
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
		Name        string `json:"username" schema:"username"`
		Description string `json:"password" schema:"password"`
		Place       string `json:"place" schema:"place"`
		Year        string `json:"year" schema:"year"`
		Month       string `json:"mont" schema:"month"`
		Day         string `json:"day" schema:"day"`
		Hour        string `json:"hours" schema:"hour"`
		Minute      string `json:"minute" schema:"minute"`
		Timezone    string `json:"timezone" schema:"timezone"`
		IsNew       bool   `json:"-" schema:"-"`
	}
)

func ToEventFormList(users []Event) (fs []EventForm) {
	for _, m := range users {
		fs = append(fs, m.ToForm())
	}
	return fs
}

// SetCreateValues sets de ID and slug.
func (user *Event) SetCreateValues() error {
	// Set create values only only if they were not previously
	if user.Identification.ID == uuid.Nil ||
		user.Identification.Slug.String == "" {
		pfx := user.Name.String
		user.Identification.SetCreateValues(pfx)
		user.Audit.SetCreateValues()
	}
	return nil
}

// SetUpdateValues
func (user *Event) SetUpdateValues() error {
	user.Audit.SetUpdateValues()
	return nil
}

// Match condition for
func (user *Event) Match(tc *Event) bool {
	r := user.Identification.Match(tc.Identification) &&
		user.Name == tc.Name &&
		user.Description == tc.Description &&
		user.Place == tc.Place &&
		user.ScheduledAt == tc.ScheduledAt &&
		user.BaseTZ == tc.BaseTZ
	return r
}

// ToForm lets convert a model to its associated form type.
// This convertion step could be avoided since gorilla schema allows
// to register custom decoders and in this case we need one because
// struct properties are not using Go standard types but their sql
// null conterpart types. As long as is relatively simple, ergonomic
// and could be easily implemented with generators I prefer to avoid
// the use of reflection.
func (user *Event) ToForm() EventForm {
	date := destructureDate(user.ScheduledAt.Time)
	return EventForm{
		Slug:        user.Slug.String,
		Name:        user.Name.String,
		Description: user.Description.String,
		Place:       user.Place.String,
		Year:        date["year"],
		Month:       date["month"],
		Day:         date["day"],
		Hour:        date["hour"],
		Minute:      date["minute"],
		Timezone:    date["timezone"],
		IsNew:       user.IsNew(),
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
