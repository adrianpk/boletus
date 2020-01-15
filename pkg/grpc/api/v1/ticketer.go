package v1

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	version = "v1"
)

// Ticketer is implementation of TicketerServer proto interface
type Ticketer struct {
}

// NewTicketer creates a Ticketer gRPC server
func NewTicketer(db *sql.DB) TicketerServer {
	return &Ticketer{}
}

// checkAPI checks API version.
func (t *Ticketer) checkAPI(api string) error {
	if len(api) > 0 {
		if version != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: '%s' (current: '%s') '", api, version)
		}
	}
	return nil
}

// UserIndex returns all active Events
func (s *Ticketer) IndexEvents(ctx context.Context, req *EventIDReq) (*IndexEventsRes, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &IndexEventsRes{
		Api:    version,
		Events: list,
	}, nil
}
