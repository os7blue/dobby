package sender

import (
	"dobby/common"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type mailSender struct {
}

func (m *mailSender) Send(host string, port string, username string, password string, to []string, title string, content string) error {

	userTitle := common.Option.Setting.PushTitle
	if userTitle != "" {

		title = fmt.Sprintf("[%s]%s", userTitle, title)

	}

	e := email.NewEmail()
	e.From = username
	e.To = to
	e.Subject = title
	e.Text = []byte(content)

	err := e.Send(host+":"+port, smtp.PlainAuth("", username, password, host))
	if err != nil {
		return err
	}

	return nil
}
