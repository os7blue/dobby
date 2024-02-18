package sender

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type wsSender struct {
}

// WebSocketGroup 一级 map 为通道列表 二级 map 为 socket 实例保存列表
var webSocketGroup = map[string]*map[string]*websocket.Conn{}

func (w *wsSender) NewConn(channelId string, conn *websocket.Conn) string {
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
