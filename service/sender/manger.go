package sender

var Senders = new(sender)

type sender struct {
	WebhookSender webhookSender
	MailSender    mailSender
	WxMpSender    wxMpSender
	WebSocket     webSocketSender
}
