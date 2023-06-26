package bootstrap

import (
	"dobby/api"
	"dobby/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

func apiInit(g *gin.Engine) {

	rootGroup := g.Group("/api")

	//login , push
	loginGroup := rootGroup.Group("/login")
	loginGroup.POST("/send_code", api.Apis.AuthApi.SendCode)
	loginGroup.POST("/login", api.Apis.AuthApi.Login)
	loginGroup.POST("/logOut", api.Apis.AuthApi.LogOut)

	//channel
	channelGroup := rootGroup.Group("/admin/channel")

	//channel info
	channelInfoGroup := channelGroup.Group("/info")
	channelInfoGroup.POST("/create", api.Apis.ChannelInfoApi.Create)
	channelInfoGroup.POST("/load", api.Apis.ChannelInfoApi.Load)
	channelInfoGroup.POST("/update", api.Apis.ChannelInfoApi.Update)
	channelInfoGroup.POST("/del", api.Apis.ChannelInfoApi.Delete)

	//channel plan
	channelPlanGroup := channelGroup.Group("/plan")
	channelPlanGroup.POST("/create", api.Apis.ChannelPlanApi.Create)
	channelPlanGroup.POST("/load", api.Apis.ChannelPlanApi.Load)
	channelPlanGroup.POST("/update", api.Apis.ChannelPlanApi.Update)
	channelPlanGroup.POST("/change_status", api.Apis.ChannelPlanApi.ChangeStatus)
	channelPlanGroup.POST("/del", api.Apis.ChannelPlanApi.Delete)
	channelPlanGroup.POST("/refresh_key", api.Apis.ChannelPlanApi.RefreshKey)

	//push
	g.POST("/push", api.Apis.PushApi.Push)

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	var wss = map[string]*map[string]*websocket.Conn{}

	var index = 1

	g.GET("/ws/send/:key/:id/:msg", func(c *gin.Context) {
		m := map[string]string{}
		c.BindUri(&m)
		first := wss[m["key"]]
		if first == nil {
			common.R.FailWithMsg(c, "不存在 key")
			return
		}

		second := *first
		if m["id"] == "0" {

			for _, value := range second {
				_ = value.WriteMessage(websocket.TextMessage, []byte(m["msg"]))

			}
		} else {

			if second[m["id"]] != nil {
				second[m["id"]].WriteMessage(websocket.TextMessage, []byte(m["msg"]))
			}

		}

		common.R.Success(c)

	})

	G.GET("/ws/list/load", func(c *gin.Context) {
		common.R.SuccessWithData(c, wss)
	})

	//ws
	g.GET("/ws/:key", func(c *gin.Context) {

		m := map[string]string{}
		err := c.BindUri(&m)
		if err != nil {
			return
		}

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			common.R.FailWithMsg(c, err.Error())
			return
		}

		if m["key"] != "laowang" {
			ws.WriteMessage(websocket.TextMessage, []byte("404"))

			err := ws.Close()
			if err != nil {
				return
			}
			return
		}

		if wss["laowang"] == nil {
			wss["laowang"] = &map[string]*websocket.Conn{}
		}
		m1 := *wss["laowang"]
		m1[strconv.Itoa(index)] = ws

		ws.WriteMessage(websocket.TextMessage, []byte("id:"+strconv.Itoa(index)))

		index++

		for true {
			message, i, err := ws.ReadMessage()
			fmt.Println(message)
			fmt.Println(i)
			if err != nil {
				return
			}

			err = ws.WriteMessage(websocket.TextMessage, []byte("hello"))
			if err != nil {
				return
			}
		}

	})

}
