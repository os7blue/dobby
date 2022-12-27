package service

import (
	"dobby/common"
	"dobby/service/sender"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type authService struct {
}

func (a *authService) SendCode(email string) error {

	emails := common.Option.Email.Admin
	emails = fmt.Sprintf("%s,%s", emails, common.Option.Email.Username)

	if !strings.Contains(emails, email) {
		return errors.New("email地址不存在")
	}

	code := common.GlobalUtil.RandCodeString(6)
	err := sender.Senders.MailSender.Send(
		common.Option.Email.Host,
		common.Option.Email.Port,
		common.Option.Email.Username,
		common.Option.Email.Password,
		[]string{email}, "推送平台登录验证码",
		code,
	)
	if err != nil {
		return err
	}
	common.LocalCache.SetWithTTL(
		fmt.Sprintf("code-%s", email),
		code,
		1,
		time.Second*360,
	)

	return nil

}

func (a *authService) Login(email string, code string) (string, error) {

	emails := common.Option.Email.Admin
	emails = fmt.Sprintf("%s,%s", emails, common.Option.Email.Username)

	if !strings.Contains(emails, email) {
		return "", errors.New("email地址不存在")
	}

	key := fmt.Sprintf("code-%s", email)
	c, exist := common.LocalCache.Get(key)
	if !exist {
		return "", errors.New("验证码有误")
	}

	if c.(string) != code {
		return "", errors.New("验证码有误")
	}

	tokenCode := uuid.New()
	common.LocalCache.SetWithTTL(
		fmt.Sprintf("auth-%s", tokenCode.String()),
		email,
		1,
		time.Second*7200,
	)

	common.LocalCache.Del(key)
	return tokenCode.String(), nil

}
