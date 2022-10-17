package service

import (
	"fmt"
	"message-push/bootstrap"
	"message-push/common"
	"message-push/common/sender"
)

type authService struct {
}

func (a *authService) SendCode(email string) error {

	code := common.RandCodeString(6)

	err := sender.MailSender.Send([]string{email}, "推送平台登录验证码", code)
	if err != nil {
		return err
	}

	bootstrap.LocalCache.Set(fmt.Sprintf("code-%s", email), code, 18000)

	return nil
}
