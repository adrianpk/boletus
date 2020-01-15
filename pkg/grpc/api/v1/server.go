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
	Server struct {
		Cfg     *fnd.Config
		Log     fnd.Logger
		Name    string
		Service *svc.Service
	}
)

const (
	version = "v1"
)

// Server is an mplementation of Server proto interface
func NewServer(cfg *fnd.Config, log fnd.Logger, name string) TicketerServer {
	return &Server{
		Cfg:  cfg,
		Log:  log,
		Name: name,
	}
}

// checkAPI checks API version.
func (t *Server) checkAPI(api string) error {
	if len(api) > 0 {
		if version != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: '%s' (current: '%s') '", api, version)
		}
	}
	return nil
}

// UserIndex returns all active Events
func (s *Server) IndexEvents(ctx context.Context, req *EventIDReq) (*IndexEventsRes, error) {
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

// Misc
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
