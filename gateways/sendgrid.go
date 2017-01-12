package gateways

import (
	"fmt"

	"github.com/ghmeier/bloodlines/config"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

/*SendgridI is the interface for sendgrid gateways*/
type SendgridI interface {
	SendEmail(string, string, string) error
}

/*Sendgrid implements SendgridI for a given sendgrid config*/
type Sendgrid struct {
	config config.Sendgrid
}

/*NewSendgrid returns a pointer to a Sendgrid struct*/
func NewSendgrid(config config.Sendgrid) SendgridI {
	return &Sendgrid{config: config}
}

/*SendEmail sends an email through sendgrid with the given target email*/
func (s *Sendgrid) SendEmail(target string, subject string, text string) error {
	m := s.getSGMail(target, subject, text)

	request := s.getSendRequest(m)
	response, err := sendgrid.API(request)
	if err != nil {
		return err
	}
	if response.StatusCode > 299 {
		return fmt.Errorf("ERROR: invalid request, %s", response.Body)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	return nil
}

func (s *Sendgrid) getSendRequest(m *mail.SGMailV3) rest.Request {
	request := sendgrid.GetRequest(s.config.APIKey, "/v3/mail/send", s.config.Host)
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)

	return request
}

func (s *Sendgrid) getSGMail(target string, subject string, text string) *mail.SGMailV3 {
	from := mail.NewEmail(s.config.FromName, s.config.FromEmail)
	to := mail.NewEmail("Garret Meier", target)
	content := mail.NewContent("text/plain", text)

	return mail.NewV3MailInit(from, subject, to, content)
}
