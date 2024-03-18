package sender

import (
	"dobby/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type wsSender struct {
}

// WebSocketGroup 一级 map 为通道列表 二级 map 为 socket 实例保存列表
var (
	webSocketGroup = map[uint]*map[string]*websocket.Conn{}
	mu             sync.Mutex
	checkPeriod    = 60 * time.Second // 检查连接状态的时间间隔
)

func init() {
	// 启动定时任务来检查连接状态
	go checkConnections()
}

func checkConnections() {

	for {

		for channelId, group := range webSocketGroup {
			for clientId, conn := range *group {
				// 尝试从连接中读取数据，如果失败说明连接已断开
				_, _, err := conn.ReadMessage()
				if err != nil {
					// 从group中删除已断开的连接
					delete(*group, clientId)
					fmt.Printf("连接已断开：通道%d 客户端%s\n", channelId, clientId)
				}
			}
		}
		time.Sleep(60 * time.Second)

	}
}
func (w wsSender) Conn(channelId uint, conn *websocket.Conn) string {

	group, ext := webSocketGroup[channelId]
	if !ext {

		group = &map[string]*websocket.Conn{}
		webSocketGroup[channelId] = group

	}

	clientId := uuid.New().String()

	(*group)[clientId] = conn

	return clientId
}

func (w wsSender) Send(channelId uint, title string, content string) error {

	group, exist := webSocketGroup[channelId]
	if !exist {
		return errors.New(fmt.Sprintf("通道%d没有客户端在线", channelId))
	}

	wait := sync.WaitGroup{}
	wait.Add(len(*group))

	for _, v := range *group {

		go func(title string, content string, v *websocket.Conn) {
			defer wait.Done()
			_ = v.WriteJSON(
				model.Result{
					Code:  2,
					Msg:   "",
					Data:  fmt.Sprintf("%s%s", title, content),
					Count: 0,
				},
			)

		}(title, content, v)

	}

	wait.Wait()

	return nil
}

func (w wsSender) ClientNum(id uint) int {

	num := 0
	group, exist := webSocketGroup[id]
	if exist {
		num = len(*group)
	}

	return num
}
