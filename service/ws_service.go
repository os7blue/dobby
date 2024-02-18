package service

import (
	"dobby/common"
	"dobby/service/sender"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

type wsService struct {
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wss wsService) Conn(key string, token string, c *gin.Context) {

	//ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	common.R.FailWithMsg(c, err.Error())
	//	return
	//}

}

func (wss wsService) SendCode(email string, key string) error {

	emails := common.Option.Email.Admin
	emails = fmt.Sprintf("%s,%s", emails, common.Option.Email.Username)

	if !strings.Contains(emails, email) {
		return errors.New("email地址不存在")
	}

	code := uuid.New().String()
	err := sender.Senders.MailSender.Send(
		common.Option.Email.Host,
		common.Option.Email.Port,
		common.Option.Email.Username,
		common.Option.Email.Password,
		[]string{email}, fmt.Sprintf("ws通道%s授权码", key),
		code,
	)
	if err != nil {
		return err
	}

	m := map[string]string{
		"key": key,
	}
	common.LocalCache.SetWithTTL(
		fmt.Sprintf("wscode-%s", code),
		m,
		1,
		time.Hour*24,
	)

	return nil

}
