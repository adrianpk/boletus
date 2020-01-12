package kabestan

import (
	"fmt"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	uuid "github.com/satori/go.uuid"
)

type (
	Mailer interface {
		Send(em Email) (resend bool, err error)
	}
)

type SESMailer struct {
	*Worker
	client *ses.SES
}

// Email struct
type Email struct {
	ID      uuid.UUID
	Name    string
	From    string
	To      string
	CC      string
	BCC     string
	Subject string
	Body    string
	Charset string
}

// MakeEmail
func MakeEmail(name, from, to, cc, bcc, subject, body string) Email {
	return Email{
		ID:      uuid.NewV4(),
		Name:    name,
		From:    from,
		To:      to,
		CC:      cc,
		BCC:     bcc,
		Subject: subject,
		Body:    body,
		Charset: "UTF-8",
	}
}

const (
	// TODO: This value should be configurable
	region = "eu-west-1"
)

// NewSESMailer creates and returns a new SES Mailer.
func NewSESMailer(cfg *Config, log Logger, name string) (*SESMailer, error) {
	if name == "" {
		name = fmt.Sprintf("ses-mailer-%s", nameSufix())
	}

	h, err := newSESMailer(cfg, log, name)
	if err != nil {
		return nil, err
	}

	log.Info("New handler", "name", name)

	return h, nil
}

// Send an email.
func (h *SESMailer) Send(em Email) (resend bool, err error) {
	email := newSESEmail(em.From, em.To, em.CC, em.BCC, em.Subject, em.Body, em.Charset)
	result, err := h.client.SendEmail(email)

	// Actually, all error cases are solved in the same way.
	// In case that, eventually, it is not required to modify
	// this behavior for some particular case, the following block
	// could be replaced by a single line of code:
	// return true, err
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {

			case ses.ErrCodeMessageRejected:
				// SES mail sending not succeed
				// It probably does not exist but we can try again
				return true, fmt.Errorf("cannot send the email: %s", err.Error())

			case ses.ErrCodeMailFromDomainNotVerifiedException:
				// SES cannot read MX record.
				// It probably does not exist but we can try again
				// just in cae it was a temporary failure
				return true, fmt.Errorf("target domain not verified: %s", err.Error())

			case ses.ErrCodeConfigurationSetDoesNotExistException:
				// Configuration error, try a resend.
				return true, fmt.Errorf("configuration error: %s", err.Error())

			default:
				// Default condition for SES related errors.
				return true, fmt.Errorf("cannot send the email: %s", err.Error())
			}
		}
		// Default condition for SES non codified errors.
		return true, fmt.Errorf("cannot send the email: %s", err.Error())
	}

	h.Log.Info("SES mailer mail sending", "result", result.GoString())

	return false, nil
}

func newSESEmail(from, to, cc, bcc, subject, body, charset string) *ses.SendEmailInput {
	// Assemble the email.
	email := &ses.SendEmailInput{
		Destination: &ses.Destination{
			BccAddresses: []*string{
				aws.String(bcc),
			},
			CcAddresses: []*string{
				aws.String(cc),
			},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charset),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(from),
	}

	return email
}

func newSESMailer(cfg *Config, log Logger, name string) (*SESMailer, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		return nil, err
	}

	// Create a SES session.
	clt := ses.New(sess)

	return &SESMailer{
		Worker: NewWorker(cfg, log, name),
		client: clt,
	}, nil
}

// Client return the provider client.
func (p *SESMailer) Client() interface{} {
	return p.client
}
