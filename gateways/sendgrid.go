package gateways

import (
	"fmt"

	"github.com/ghmeier/bloodlines/config"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridI interface {
	SendEmail(string, string, string) error
}

type Sendgrid struct {
	config config.Sendgrid
}

func NewSendgrid(config config.Sendgrid) SendgridI {
	return &Sendgrid{config: config}
}

func (s *Sendgrid) SendEmail(target string, subject string, text string) error {
	m := s.getSGMail(target, subject, text)

	request := s.getSendRequest(m)
	response, err := sendgrid.API(request)
	if err != nil {
		return err
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
	to := mail.NewEmail("User", target)
	content := mail.NewContent("text/plain", text)

	return mail.NewV3MailInit(from, subject, to, content)
}
