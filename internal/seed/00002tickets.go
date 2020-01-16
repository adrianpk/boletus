package seed

import (
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

var (
	t1, _ = time.Parse(time.RFC3339, "2020-01-02T19:00:00Z00:00")

	event = newEventMap("Rock Party in Wrocław", "Make some noise!", "Wrocław Stadion Miejski", t1)
)

// EventsAndTickets seeding
func (s *step) EventAndTickets() error {
	err := s.Events()
	if err != nil {
		return err
	}

	return s.Tickets()
}

// Events seeding
func (s *step) Events() error {
	tx := s.GetTx()

	st := `INSERT INTO events (id, slug, name, description, place, scheduled_at, base_tz, is_active, is_deleted, created_by_id, updated_by_id, created_at, updated_at)
VALUES (:id, :slug, :name, :description, :place, :scheduled_at, :base_tz, :is_active, :is_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at)`

	_, err := tx.NamedExec(st, event)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	return nil
}

// Tickets seeding
func (s *step) Tickets() error {
	tx := s.GetTx()

	st := `INSERT INTO tickets (id, slug, name, event_id, type, serie, number, seat, price, currency, reserved_by, reserved_at, bought_by, bought_at, local_order_id, gateway_order_id, gateway_op_status, base_tz, is_active, is_deleted, created_by_id, updated_by_id, created_at, updated_at)
VALUES (:id, :slug, :name, :event_id, :type, :serie, :number, :seat, :price, :currency, :reserved_by, :reserved_at, :bought_by, :bought_at, :local_order_id, :gateway_order_id, :gateway_op_status, :base_tz, :is_active, :is_deleted, :created_by_id, :updated_by_id, :created_at, :updated_at);`

	// eventMap, serie, qty, priceInMillis
	tickets := newTicketSerie(event, "normal", "A", 500, 30000)

	for _, t := range tickets {
		_, err := tx.NamedExec(st, t)
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
	}

	tickets = newTicketSerie(event, "golden-circle", "A", 20, 75000)

	for _, t := range tickets {
		_, err := tx.NamedExec(st, t)
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
	}

	tickets = newTicketSerie(event, "silver-circle", "A", 40, 50000)

	for _, t := range tickets {
		_, err := tx.NamedExec(st, t)
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
	}

	tickets = newTicketSerie(event, "bronze-circle", "A", 40, 40000)

	for _, t := range tickets {
		_, err := tx.NamedExec(st, t)
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
	}

	return nil
}

func newEventMap(name, description, place string, scheduledAt time.Time) map[string]interface{} {

	return map[string]interface{}{
		"id":            genUUID(),
		"slug":          genSlug(name),
		"name":          name,
		"description":   description,
		"place":         place,
		"scheduled_at":  scheduledAt,
		"timezone":      "CES",
		"base_tz":       "GMT",
		"is_active":     true,
		"is_deleted":    false,
		"created_by_id": uuid.Nil,
		"updated_by_id": uuid.Nil,
		"created_at":    time.Now(),
		"updated_at":    time.Time{},
	}
}

func newTicketSerie(eventMap map[string]interface{}, ticketType, serie string, qty, priceInMillis int) (ticketMap []map[string]interface{}) {
	//ts := make([]map[string]interface{}, qty)
	ts := []map[string]interface{}{}

	eid, ok := eventMap["id"]
	if !ok {
		log.Println("Invalid event ID")
		return ticketMap
	}

	n, ok := eventMap["name"]
	if !ok {
		log.Println("Invalid ticket name")
		return ticketMap
	}

	s, ok := eventMap["scheduled_at"]
	if !ok {
		log.Println("Invalid schedule date")
		return ticketMap
	}

	eventID, ok := eid.(uuid.UUID)
	if !ok {
		log.Println("Invalid event ID")
		return ticketMap
	}

	name, ok := n.(string)
	if !ok {
		log.Println("Invalid event name")
		return ticketMap
	}

	schedule, ok := s.(time.Time)
	if !ok {
		log.Println("Invalid schedule date")
		return ticketMap
	}

	for i := 0; i < qty; i++ {
		t := newTicketMap(eventID, name, ticketType, schedule, priceInMillis, serie, i)
		//fmt.Printf("%+v", t)
		ts = append(ts, t)
	}

	//spew.Sdump(ts)
	return ts
}

func newTicketMap(eventID uuid.UUID, name string, ticketType string, scheduledAt time.Time, priceInMillis int, serie string, number int) (ticketMap map[string]interface{}) {
	return map[string]interface{}{
		"id":                genUUID(),
		"slug":              genSlug(name),
		"name":              name,
		"event_id":          eventID,
		"type":              ticketType,
		"serie":             serie,
		"number":            number,
		"seat":              scheduledAt,
		"price":             priceInMillis,
		"currency":          "PLN",
		"reserved_by":       uuid.Nil,
		"reserved_at":       time.Time{},
		"bought_by":         uuid.Nil,
		"bought_at":         time.Time{},
		"local_order_id":    "",
		"gateway_order_id":  "",
		"gateway_op_status": "",
		"base_tz":           "CES",
		"is_active":         true,
		"is_deleted":        false,
		"created_by_id":     uuid.Nil,
		"updated_by_id":     uuid.Nil,
		"created_at":        time.Now(),
		"updated_at":        time.Time{},
	}
}
