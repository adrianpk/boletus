package svc

import (
	"fmt"

	"github.com/robfig/cron"
)

type (
	scheduler struct {
		cron *cron.Cron
	}
)

func (s *Service) InitScheduler() error {
	s.Scheduler = &scheduler{}
	return s.initCron()
}

func (s *Service) initCron() error {
	mins := int(s.Cfg.ValAsInt("scheduler.one.minutes", 15))
	c, err := newCron(mins, s.expireReservations)
	if err != nil {
		return err
	}

	s.Scheduler.cron = c
	return nil
}

func (sc *scheduler) Start() {
	sc.cron.Start()
}

func newCron(mins int, f func()) (*cron.Cron, error) {
	c := cron.New()

	cstr := fmt.Sprintf("0 %d * * * *", mins)

	err := c.AddFunc(cstr, f)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (s *Service) expireReservations() {
	// TODO: Expire reservations
	panic("not implemented")
}
