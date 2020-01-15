package main

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/adrianpk/boletus/pkg/grpc/api/v1"
	fnd "github.com/adrianpk/foundation"
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
	version = "v1"
	// Replace by a valid event slug
	eventSlug = "wroclawparty-5806da1bbad3"
)

// This is simple client that can be used to
// manually test ticketer gRPC server exposed functions
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

	// Call IndexEvents
	clt.IndexEvents()
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

// IndexEvents list
func (c *client) IndexEvents() error {
	// EventID request
	req := v1.EventIDReq{
		Api: version,
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
	c.Log.Info("IndexEvents result:")
	c.Log.Info(fmt.Sprintf("%+v", res))
	return nil
}
