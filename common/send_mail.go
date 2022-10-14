package common

type emailConfig struct {
	Form    string
	To      []string
	Auth    string
	Title   string
	Content string
}

var MailSender = new(mailSender)

type mailSender struct {
}

func (m *mailSender) CreateConfig() *emailConfig {
	return new(emailConfig)
}
