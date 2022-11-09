package bootstrap

import (
	"dobby/api"
	"github.com/gin-gonic/gin"
)

func apiInit(g *gin.Engine) {

	//login , push
	g.POST("/login/send_code", api.Apis.AuthApi.SendCode)
	g.POST("/login/login", api.Apis.AuthApi.Login)
	g.POST("/login/logOut", api.Apis.AuthApi.LogOut)
	g.POST("/push", api.Apis.PushApi.Push)

	//channel
	channelGroup := g.Group("/admin/channel")
	//channel info
	channelInfoGroup := channelGroup.Group("/channel_info")
	channelInfoGroup.POST("/create", api.Apis.ChannelInfoApi.Create)
	channelInfoGroup.POST("/load", api.Apis.ChannelInfoApi.Load)
	channelInfoGroup.POST("/update", api.Apis.ChannelInfoApi.Update)
	channelInfoGroup.POST("/change_status", api.Apis.ChannelInfoApi.ChangeStatus)
	channelInfoGroup.POST("/del", api.Apis.ChannelInfoApi.Delete)
	channelInfoGroup.POST("/refresh_key", api.Apis.ChannelInfoApi.RefreshKey)

}
