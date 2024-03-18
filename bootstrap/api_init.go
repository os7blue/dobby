package bootstrap

import (
	"dobby/api"
	"github.com/gin-gonic/gin"
)

func apiInit(g *gin.Engine) {

	rootGroup := g.Group("/api")

	//login , push
	loginGroup := rootGroup.Group("/login")
	loginGroup.POST("/send_code", api.AuthApi.SendCode)
	loginGroup.POST("/login", api.AuthApi.Login)
	loginGroup.POST("/logOut", api.AuthApi.LogOut)

	//channel
	channelGroup := rootGroup.Group("/admin/channel")

	//channel info
	channelInfoGroup := channelGroup.Group("/info")
	channelInfoGroup.POST("/create", api.ChannelInfoApi.Create)
	channelInfoGroup.POST("/load", api.ChannelInfoApi.Load)
	channelInfoGroup.POST("/update", api.ChannelInfoApi.Update)
	channelInfoGroup.POST("/del", api.ChannelInfoApi.Delete)
	channelInfoGroup.POST("/push_ws", api.ChannelInfoApi.PushWS)

	//channel plan
	channelPlanGroup := channelGroup.Group("/plan")
	channelPlanGroup.POST("/create", api.ChannelPlanApi.Create)
	channelPlanGroup.POST("/load", api.ChannelPlanApi.Load)
	channelPlanGroup.POST("/update", api.ChannelPlanApi.Update)
	channelPlanGroup.POST("/change_status", api.ChannelPlanApi.ChangeStatus)
	channelPlanGroup.POST("/del", api.ChannelPlanApi.Delete)
	channelPlanGroup.POST("/refresh_key", api.ChannelPlanApi.RefreshKey)

	//push
	g.POST("/push", api.PushApi.Push)

	//ws
	g.GET("/ws/:ID/:Key", api.WsApi.Conn)

}
