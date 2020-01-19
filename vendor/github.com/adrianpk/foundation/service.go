package foundation

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

type (
	ServiceIF interface {
		// NOTE: Something should be here for sure.
	}
)

type (
	Service struct {
		Cfg    *Config
		Log    Logger
		Name   string
		Mailer Mailer
	}
)

func NewService(cfg *Config, log Logger, name string) *Service {
	name = genName(name, "service")

	return &Service{
		Cfg:  cfg,
		Log:  log,
		Name: name,
	}
}

func GenShortID() string {
	s := strings.Split(uuid.NewV4().String(), "-")
	return s[len(s)-1]
}
