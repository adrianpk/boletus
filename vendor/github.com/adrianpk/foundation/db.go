package kabestan

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	// "github.com/dewski/spatial"
	"errors"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // package init.
)

type (
	pgManager struct {
		*Worker
	}
)

func NewPostgresConn(cfg *Config, log Logger) (*sqlx.DB, error) {
	pm := pgManager{
		Worker: NewWorker(cfg, log, "pg-db-manager"),
	}

	res := <-pm.retryConnection()
	if res.err != nil {
		return nil, res.err
	}
	return res.conn, nil
}

// retryConnection implments a backoff mechanism for establishing a connection
// to Postgres; this is especially useful in containerized environments where
// components can be started out of order.
func (pm *pgManager) retryConnection() chan retryResult {
	res := make(chan retryResult)

	cbmax := uint64(pm.Cfg.ValAsInt("pg.backoff.maxtries", 5))
	bo := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), cbmax)

	go func() {
		defer close(res)

		url := pm.dbURL()

		for i := 0; i <= int(cbmax); i++ {

			pm.Log.Info("Dialing to Postgres", "host", url)

			conn, err := sqlx.Open("postgres", url)
			if err != nil {
				pm.Log.Error(err, "Postgres connection error")
			}

			err = conn.Ping()
			if err == nil {
				pm.Log.Info("Postgres connection established")
				res <- retryResult{conn, nil}
				return
			}

			pm.Log.Error(err, "Postgres connection error")

			// Backoff
			nb := bo.NextBackOff()
			if nb == backoff.Stop {
				pm.Log.Info("Postgres connection failed", "reason", "max number of attempts reached")
				err := errors.New("Postgres max number of connection attempts reached")
				res <- retryResult{nil, err}
				bo.Reset()
				return
			}

			pm.Log.Info("Postgres connection failed", "retrying-in", nb.String(), "unit", "seconds")
			time.Sleep(nb)
		}
	}()

	return res
}

func (pm *pgManager) dbURL() string {
	host := pm.Cfg.ValOrDef("pg.host", "")
	port := pm.Cfg.ValAsInt("pg.port", 5432)
	db := pm.Cfg.ValOrDef("pg.database", "")
	user := pm.Cfg.ValOrDef("pg.user", "")
	pass := pm.Cfg.ValOrDef("pg.password", "")
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
}

// Misc
// Point struct
type Point struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (p *Point) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p.Lng, p.Lat)
}

// Scan implements the Scanner interface.
func (p *Point) Scan(val interface{}) error {
	b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %d", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}

	return nil
}

// Value implements the driver Value interface.
func (p Point) Value() (driver.Value, error) {
	return p.String(), nil
}

// Postgres
type retryResult struct {
	conn *sqlx.DB
	err  error
}
