package main

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/adrianpk/boletus/pkg/grpc/api/v1"
	fnd "github.com/adrianpk/foundation"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
)

type (
	client struct {
		Cfg      *fnd.Config
		Log      fnd.Logger
		Conn     *grpc.ClientConn
		Ticketer v1.TicketerClient
	}
)

const (
	apiVer = "v1"

	// To test replace by valid user and event slugs

	// select slug from users;
	userSlug = "lauriem-000000000004"

	// select slug from events;
	eventSlug = "rockpartyinwrocław-000000000001"

	// PreBook ticket type [normal, golden-circle, silver-circle, bronce-circle, couple]
	ticketType = "normal"

	// replace by a a valid reservation ID
	reservationID = "cab86283242c"
)

// This is simple client that can be used to
// manually test ticketer gRPC server exposed functions
// TODO: Accept commands and IDs as flags to avoid harcoding arguments.
func main() {
	// Replace by custom envar prefix
	cfg := fnd.LoadConfig("blt")
	log := fnd.NewLogger(cfg)

	// Create a client
	clt, err := NewClient(cfg, log)
	if err != nil {
		log.Error(err, "Connection failed")
	}
	defer clt.Conn.Close()

	// IndexEvents
	//log.Info("IndexEvents begin")
	//clt.IndexEvents()
	//log.Info("IndexEvents end\n")

	// Ticket summary
	//log.Info("TicketSummary begin")
	//clt.EventTicketSummary()
	//log.Info("TicketSummary end\n")

	// PreBook
	//log.Info("PreBook begin")
	//clt.PreBook()
	//log.Info("PreBook end\n")

	// ConfirmBooking
	log.Info("ConfirmBooking begin")
	clt.ConfirmBooking()
	log.Info("ConfirmBooking end\n")
}

// NewClient for Ticketer gRPC server
func NewClient(cfg *fnd.Config, log fnd.Logger) (*client, error) {
	host := cfg.ValOrDef("grpc.client.host", "localhost")
	port := cfg.ValOrDef("grpc.client.port", "8082")
	addr := fmt.Sprintf("%s:%s", host, port)

	log.Info("Dialing gRPC server", "address", addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	t := v1.NewTicketerClient(conn)

	return &client{
		Cfg:      cfg,
		Log:      log,
		Conn:     conn,
		Ticketer: t,
	}, nil
}

// IndexEvents
func (c *client) IndexEvents() error {
	// EventID request
	req := v1.EventIDReq{
		Api: apiVer,
	}

	// Context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Ticketer.IndexEvents(ctx, &req)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	// Dump result
	c.Log.Info("Result:")
	c.Log.Info(spew.Sdump(res))
	return nil
}

// EventTicketSummary
func (c *client) EventTicketSummary() error {
	// EventID request
	req := v1.EventIDReq{
		Api:  apiVer,
		Slug: eventSlug,
	}

	// Context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Ticketer.EventTicketSummary(ctx, &req)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	// Dump result
	c.Log.Info("Result:")
	c.Log.Info(spew.Sdump(res))
	return nil
}

// PreBook
func (c *client) PreBook() error {
	// EventID request
	req := v1.PreBookReq{
		Api:        apiVer,
		UserSlug:   userSlug,
		EventSlug:  eventSlug,
		TicketType: ticketType,
		Qty:        4,
	}

	// Context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Ticketer.PreBook(ctx, &req)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	// Dump result
	c.Log.Info("Result:")
	c.Log.Info(spew.Sdump(res))
	return nil
}

// ConfirmBooking
func (c *client) ConfirmBooking() error {
	// EventID request
	req := v1.ConfirmBookingReq{
		Api:           apiVer,
		EventSlug:     eventSlug,
		UserSlug:      userSlug,
		ReservationID: reservationID,
	}

	// Context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := c.Ticketer.ConfirmBooking(ctx, &req)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	// Dump result
	c.Log.Info("Result:")
	c.Log.Info(spew.Sdump(res))
	return nil
}
