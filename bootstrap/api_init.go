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
	channelGroup.POST("/channel_info/create", api.Apis.ChannelInfoApi.Create)

}
