package sender

import (
	"github.com/jordan-wright/email"
	"message-push/bootstrap"
	"net/smtp"
)

type mailSender struct {
}

func (m *mailSender) Send(to []string, title string, content string) error {

	ec := bootstrap.Option.Email
	e := email.NewEmail()
	e.From = ec.Username
	e.To = to
	e.Subject = title
	e.Text = []byte(content)

	err := e.Send(ec.Host+":"+ec.Port, smtp.PlainAuth("", ec.Username, ec.Password, ec.Host))
	if err != nil {
		return err
	}

	return nil
}
