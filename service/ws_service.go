package service

import (
	"dobby/common"
	"dobby/model"
	"dobby/service/sender"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"net/http"
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

func (wss wsService) Conn(ID uint, key string, c *gin.Context) (string, error) {

	channel := model.ChannelInfo{}
	db := common.DB.First(&channel, ID)
	if db.Error != nil {
		return "", db.Error
	}

	r := gjson.Get(channel.OptionJsonStr, "key").Str
	if r != key {
		return "", errors.New("key错误")
	}

	upgrade, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return "", err
	}

	clientId := sender.WsSender.Conn(ID, upgrade)

	err = upgrade.WriteJSON(model.Result{
		Code:  1,
		Msg:   "链接成功",
		Data:  clientId,
		Count: 0,
	})

	if err != nil {
		return "", err
	}

	return clientId, nil
}

func (wss wsService) Send(id uint, content string) error {

	err := sender.WsSender.Send(id, "", content)
	if err != nil {
		return err
	}

	return nil
}
