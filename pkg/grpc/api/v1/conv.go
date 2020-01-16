package v1

import (
	"time"

	"github.com/adrianpk/boletus/internal/model"
)

func toEventResList(events []model.Event) (ers []*EventRes) {
	for _, m := range events {
		ers = append(ers, toEventRes(m))
	}
	return ers
}

func toEventRes(event model.Event) *EventRes {
	return &EventRes{
		Slug:        event.Slug.String,
		Name:        event.Name.String,
		Description: event.Description.String,
		Place:       event.Place.String,
		ScheduledAt: event.ScheduledAt.Time.Format(time.RFC3339),
		IsNew:       event.IsNew(),
	}
}

func toTicketSummaryList(tss []model.TicketSummary) (ers []*TicketSummaryRes) {
	for _, m := range tss {
		ers = append(ers, toTicketSummaryRes(m))
	}
	return ers
}

func toTicketSummaryRes(ts model.TicketSummary) *TicketSummaryRes {
	return &TicketSummaryRes{
		Qty:       ts.Qty.Int32,
		Name:      ts.Name.String,
		EventSlug: ts.EventSlug.String,
		Type:      ts.Type.String,
		Price:     ts.PriceFloat32(),
		Currency:  ts.Currency.String,
	}
}
