package v1

import (
	"context"

	"github.com/adrianpk/boletus/internal/app/svc"
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
	if err := s.checkAPI(req.GetApi()); err != nil {
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
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	// Get ticket info from service
	ts, err := s.Service.TicketSummary(req.GetSlug())

	// Convert result list into a EventRes list
	list := toTicketSummaryList(ts)

	return &TicketSummaryListRes{
		Api:  version,
		List: list,
	}, err
}

// PreBook makes a ticket reservation for an event.
// NOTE: Isn't it risky to let pre-book ticket that as 'all-together' as selling option?
// It will make all tickets appear unavailable for 15 minutes.
// and nothing prevents from pre-booking all them again later.
// Maybe charging a fee will discourage massive reservations.
func (s *GRPCService) PreBook(ctx context.Context, req *PreBookReq) (*TicketOpRes, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	// Get ticket info from service
	qty := int(req.GetQty())
	ts, err := s.Service.PreBookTickets(req.GetEventSlug(), req.GetTicketType(), qty, req.GetUserSlug())

	// Convert result list into a TicketRes list
	list, total, currency, resID := toTicketResList(ts)

	return &TicketOpRes{
		Api:           version,
		List:          list,
		Total:         total,
		Currency:      currency,
		ReservationID: resID,
		Status:        "reserved",
	}, err
}

// PreBooking confirms reservations
// NOTE: This is a work in progress concept.
// Probably, depending on payment system implementation,
// Confirmation should first put the tickets in an intermediate state at the beginning of
// transaction ('processing-payment'), call the payment gateway using required values
// (card data, blik code, etc) and a callback address to close the loop: mark  tickets
// as 'payed' and send notifications to user.
func (s *GRPCService) ConfirmBooking(ctx context.Context, req *ConfirmBookingReq) (*TicketOpRes, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.GetApi()); err != nil {
		return nil, err
	}

	// Confirm booking
	ts, err := s.Service.ConfirmTicketsReservation(req.GetEventSlug(), req.GetReservationID(), req.GetUserSlug())

	// Convert result list into a TicketRes list
	list, total, currency, resID := toTicketResList(ts)

	return &TicketOpRes{
		Api:           version,
		List:          list,
		Total:         total,
		Currency:      currency,
		ReservationID: resID,
		Status:        "paid",
	}, err
}
