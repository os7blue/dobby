package bootstrap

import (
	"dobby/api"
	"github.com/gin-gonic/gin"
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
	channelInfoGroup.POST("/refresh_key", api.Apis.ChannelInfoApi.RefreshKey)

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

}
