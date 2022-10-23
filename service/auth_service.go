package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"message-push/common"
	"message-push/service/sender"
)

type authService struct {
}

func (a *authService) SendCode(email string) error {

	code := common.RandCodeString(6)

	emails := common.Option.Email.Admin
	emails = append(emails, common.Option.Email.Username)

	for _, e := range emails {

		if e == email {

			err := sender.MailSender.Send([]string{email}, "推送平台登录验证码", code)
			if err != nil {
				return errors.New("发送验证码失败")
			}
			common.LocalCache.Set(fmt.Sprintf("code-%s", email), code, 18000)

			return nil

		}

	}

	return errors.New("email地址不存在")
}

func (a *authService) Login(email string, code string) (string, error) {

	emails := common.Option.Email.Admin
	emails = append(emails, common.Option.Email.Username)

	for _, e := range emails {

		if e == email {

			key := fmt.Sprintf("code-%s", email)
			c, exist := common.LocalCache.Get(key)
			if !exist {
				return "", errors.New("验证码有误")
			}

			if c.(string) != code {
				return "", errors.New("验证码有误")
			}

			tokenCode := uuid.New()
			common.LocalCache.Set(fmt.Sprintf("auth-%s", tokenCode.String()), email, 7200)

			common.LocalCache.Del(key)
			return tokenCode.String(), nil

		}

	}

	return "", errors.New("email地址不存在")

}
