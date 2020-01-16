package v1

import (
	"context"
	"time"

	"github.com/adrianpk/boletus/internal/app/svc"
	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	GRPCService struct {
		Cfg     *fnd.Config
		Log     fnd.Logger
		Name    string
		Service *svc.Service
	}
)

const (
	version = "v1"
)

// GRPCService is an implementation of Server proto interface
func NewGRPCService(cfg *fnd.Config, log fnd.Logger, name string) *GRPCService {
	return &GRPCService{
		Cfg:  cfg,
		Log:  log,
		Name: name,
	}
}

// checkAPI checks API version.
func (t *GRPCService) checkAPI(api string) error {
	if len(api) > 0 {
		if version != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: '%s' (current: '%s') '", api, version)
		}
	}
	return nil
}

// UserIndex returns all active Events
func (s *GRPCService) IndexEvents(ctx context.Context, req *EventIDReq) (*IndexEventsRes, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Get events list from registered service
	events, err := s.Service.IndexEvents()

	// Convert result list into a EventRes list
	list := toEventResList(events)

	return &IndexEventsRes{
		Api:    version,
		Events: list,
	}, err
}

// EventTicketsInfo returns info about ticket types, price and availability bv event.
func (s *GRPCService) EventTicketSummary(ctx context.Context, req *EventIDReq) (*TicketSummaryListRes, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// Get ticket info from service
	ts, err := s.Service.TicketSummary(req.Slug)

	// Convert result list into a EventRes list
	list := toTicketSummaryList(ts)

	return &TicketSummaryListRes{
		Api:  version,
		List: list,
	}, err
}

// Convertions
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
