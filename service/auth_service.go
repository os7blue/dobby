package service

import (
	"fmt"
	"message-push/bootstrap"
	"message-push/common"
	"message-push/service/sender"
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

func (a *authService) Login(email string, code string) bool {

	emails := bootstrap.Option.Email.Admin
	emails = append(emails, bootstrap.Option.Email.Username)

	for _, e := range emails {

		if e == email {

			key := fmt.Sprintf("code-%s", email)
			c, exist := bootstrap.LocalCache.Get(key)
			if !exist {
				return exist
			}

			if c.(string) == code {
				bootstrap.LocalCache.Del(key)
				return true
			}

		}

	}

	return false

}
