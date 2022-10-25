package bootstrap

import (
	"github.com/gin-gonic/gin"
	"message-push/api"
)

func apiInit(g *gin.Engine) {

	//login & index
	g.POST("/login/send_code", api.Apis.AuthApi.SendCode)
	g.POST("/login/login", api.Apis.AuthApi.Login)

	//channel
	channelGroup := g.Group("/admin/channel")
	//channel info
	channelInfoGroup := channelGroup.Group("/channel_info")
	channelInfoGroup.POST("/create", api.Apis.ChannelInfoApi.Create)
	channelInfoGroup.POST("/load", api.Apis.ChannelInfoApi.Load)
	channelInfoGroup.POST("/update", api.Apis.ChannelInfoApi.Update)

}
