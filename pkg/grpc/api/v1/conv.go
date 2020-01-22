package v1

import (
	"github.com/adrianpk/boletus/internal/model"
	"time"
)

func toEventResList(events []model.Event) (ers []*EventRes) {
	for _, m := range events {
		ers = append(ers, toEventRes(m))
	}
	return ers
}

func toEventRes(event model.Event) *EventRes {
	return &EventRes{
		Api:         "v1",
		Slug:        event.Slug.String,
		Name:        event.Name.String,
		Description: event.Description.String,
		Place:       event.Place.String,
		ScheduledAt: event.ScheduledAt.Time.Format(time.RFC3339),
	}
}

func toTicketSummaryList(tss []model.TicketSummary) (tsr []*TicketSummaryRes) {
	for _, t := range tss {
		tsr = append(tsr, toTicketSummaryRes(t))
	}
	return tsr
}

func toTicketSummaryRes(ts model.TicketSummary) *TicketSummaryRes {
	return &TicketSummaryRes{
		Api:       "v1",
		Qty:       ts.Qty.Int32,
		Name:      ts.Name.String,
		EventSlug: ts.EventSlug.String,
		Type:      ts.Type.String,
		Price:     ts.PriceFloat32(),
		Currency:  ts.Currency.String,
		Prices:    ts.Prices,
	}
}

func toTicketResList(tickets []model.Ticket) (trs []*TicketRes, total float32, currency, reservationID string) {
	for _, t := range tickets {
		trs = append(trs, toTicketRes(&t))
		total = total + float32(t.Price.Int32/1000)
	}

	if len(tickets) > 0 {
		currency = tickets[0].Currency.String
		reservationID = tickets[0].ReservationID.String
	}

	return trs, total, currency, reservationID
}

func toTicketRes(ticket *model.Ticket) *TicketRes {
	return &TicketRes{
		Api:             "v1",
		Name:            ticket.Name.String,
		EventSlug:       ticket.EventID.String,
		Type:            ticket.Type.String,
		Serie:           ticket.Serie.String,
		Number:          int32(ticket.Number.Int32),
		Seat:            ticket.Seat.String,
		Price:           ticket.Price.Int32,
		Currency:        ticket.Currency.String,
		Status:          ticket.Status.String,
		ReservationID:   ticket.ReservationID.String,
		LocalOrderID:    ticket.LocalOrderID.String,
		GatewayOpID:     ticket.GatewayOpID.String,
		GatewayOpStatus: ticket.GatewayOpStatus.String,
	}
}
