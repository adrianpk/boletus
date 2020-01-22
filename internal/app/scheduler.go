package app

import (
	"github.com/adrianpk/boletus/internal/svc"
	fnd "github.com/adrianpk/foundation"
	"github.com/robfig/cron/v3"
)

type (
	scheduler struct {
		Cfg     *fnd.Config
		Log     fnd.Logger
		Name    string
		cron1   *cron.Cron
		cron2   *cron.Cron
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
	err := sc.startCron1()
	if err != nil {
		return err
	}

	return sc.startCron2()
}

func (sc *scheduler) startCron1() error {
	cronStr1 := sc.Cfg.ValOrDef("scheduler.cron.one.str", "* * * * *")

	// Expire reservations job
	c1, err := newCron(cronStr1, sc.Service.ExpireTicketReservations)
	if err != nil {
		return err
	}

	sc.cron1 = c1
	sc.cron1.Start()
	return nil
}

func (sc *scheduler) startCron2() error {
	// Update currency rates cache
	cronStr1 := sc.Cfg.ValOrDef("scheduler.cron.two.str", "0 * * * *")

	c2, err := newCron(cronStr1, sc.Service.UpdateRates)
	if err != nil {
		return err
	}

	sc.cron2 = c2
	sc.cron2.Start()

	return nil
}

func newCron(cronStr string, f func()) (*cron.Cron, error) {
	c := cron.New()

	_, err := c.AddFunc(cronStr, f)
	if err != nil {
		return c, err
	}

	return c, nil
}
