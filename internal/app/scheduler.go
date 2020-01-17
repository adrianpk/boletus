package app

import (
	"github.com/adrianpk/boletus/internal/app/svc"
	fnd "github.com/adrianpk/foundation"
	"github.com/robfig/cron/v3"
)

type (
	scheduler struct {
		Cfg     *fnd.Config
		Log     fnd.Logger
		Name    string
		cron    *cron.Cron
		Service *svc.Service
	}
)

func NewScheduler(cfg *fnd.Config, log fnd.Logger, name string) *scheduler {
	return &scheduler{
		Cfg:  cfg,
		Log:  log,
		Name: name,
	}
}

func (sc *scheduler) Start() error {
	mins := int(sc.Cfg.ValAsInt("scheduler.one.minutes", 1))

	// Expire reservations job
	c, err := newCron(mins, sc.Service.ExpireTicketReservations)
	if err != nil {
		return err
	}

	sc.cron = c
	sc.cron.Start()

	return nil
}

func newCron(mins int, f func()) (*cron.Cron, error) {
	c := cron.New()

	// FIX: Runing every minute.
	// Lib robfig/cron/ changed behaviour in v3.
	//cstr := fmt.Sprintf("0 %d * * *", mins)
	cstr := "* * * * *"

	_, err := c.AddFunc(cstr, f)
	if err != nil {
		return c, err
	}

	return c, nil
}
