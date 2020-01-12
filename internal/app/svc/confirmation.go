package svc

import (
	"fmt"

	"github.com/adrianpk/boletus/internal/model"
	fnd "github.com/adrianpk/foundation"
)

func (s *Service) MakeConfirmationEmail(u *model.User) fnd.Email {
	cfg := s.Cfg

	name := cfg.ValOrDef("mailer.agent.name", "mailer")
	from := cfg.ValOrDef("mailer.agent.mail", "dontreply@localhost")
	to := u.Email.String
	subject := fmt.Sprintf("%s, please confirm your account!", u.Username.String)

	site := cfg.ValOrDef("site.url", "localhost")
	path := cfg.ValOrDef("user.confirmation.path", "auth/%s/verify/%s")
	confPath := fmt.Sprintf(path, u.Slug.String, u.ConfirmationToken.String)
	link := fmt.Sprintf("https://%s/%s", site, confPath)

	body := "<p>Hi %s, follow this link to confirm your account: <br/><br/>"
	body = body + "<a href=\"%s\">%s</a><br/<br/>"
	body = body + "Thanks!"
	body = fmt.Sprintf(body, u.Username.String, link, link)

	m := fnd.MakeEmail(name, from, to, "", "", subject, body)

	s.Log.Info("User account confirmation", "mail-body", fmt.Sprintf("%+v", m))

	return m
}

func (s *Service) makeConfirmationEmail(u *model.User) fnd.Email {
	cfg := s.Cfg

	name := cfg.ValOrDef("mailer.agent.name", "mailer")
	from := cfg.ValOrDef("mailer.agent.mail", "dontreply@localhost")
	to := u.Email.String
	subject := fmt.Sprintf("%s, please confirm your account!", u.Username.String)

	site := cfg.ValOrDef("site.url", "localhost")
	path := cfg.ValOrDef("user.confirmation.path", "auth/%s/verify/%s")
	confPath := fmt.Sprintf(path, u.Slug.String, u.ConfirmationToken.String)
	link := fmt.Sprintf("https://%s/%s", site, confPath)

	body := "<p>Hi %s, follow this link to confirm your account: <br/><br/>"
	body = body + "<a href=\"%s\">%s</a><br/<br/>"
	body = body + "Thanks!"
	body = fmt.Sprintf(body, u.Username.String, link, link)

	m := fnd.MakeEmail(name, from, to, "", "", subject, body)

	// s.Log.Info("User account confirmation", "mail-body", fmt.Sprintf("%+v", m))

	return m
}

// NOTE: This is just to get an out of the box solution to send emails.
// Resend management implementation is still incomplete.
// Anyway, an option to configure an external dispatching mechanism will be implemented.
// and also something like "Your account has not been confirmed, resend email to xxxx@yyyyyy.com?"
// will also be developed.
func (s *Service) sendConfirmationEmail(u *model.User) {
	cfg := s.Cfg

	debug := cfg.ValAsBool("user.confirmation.debug", false)
	send := cfg.ValAsBool("user.confirmation.send", false)

	var m fnd.Email

	if debug {
		m = s.makeConfirmationEmail(u)
		s.Log.Debug("Confirmation email", "subject", m.Subject, "body", m.Body)
	}

	if !send {
		s.Log.Info("User signup email confirmation send is disabled")
		return
	}

	// Build mail
	m = s.makeConfirmationEmail(u)

	// Send it
	go func() {
		_, err := s.Mailer.Send(m)
		if err != nil {
			s.Log.Error(err)
		}
	}()
}
