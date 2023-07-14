package sender

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"sync"
)

type webSocketSender struct {
}

// WebSocketGroup 一级 map 为通道列表 二级 map 为 socket 实例保存列表
var webSocketGroup = map[string]*map[string]*websocket.Conn{}

func (w *webSocketSender) NewConn(channelId string, conn *websocket.Conn) string {
	channel := webSocketGroup[channelId]
	clientID := uuid.New().String()

	if channel == nil {
		clientMap := map[string]*websocket.Conn{
			clientID: conn,
		}
		webSocketGroup[channelId] = &clientMap
	} else {
		clientMap := *channel
		clientMap[clientID] = conn
	}

	return clientID
}

func (w *webSocketSender) SendClient(channelId string, ClientId string) {

}

func (w *webSocketSender) Send(channelId string) error {

	channel := *webSocketGroup[channelId]
	if channel == nil || len(channel) == 0 {
		return errors.New(fmt.Sprintf("通道[%s]没有接收端在线", channelId))

	}

	var wait sync.WaitGroup
	wait.Add(len(channel))

	for id, conn := range channel {

		go func(id string, conn *websocket.Conn) {

			if conn == nil {

			}

			wait.Done()

		}(id, conn)

	}

	wait.Wait()

	return nil
}
