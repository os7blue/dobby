package sender

var Senders = new(sender)

type sender struct {
	MailSender mailSender
}
