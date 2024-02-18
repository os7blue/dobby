package api

var Apis = new(apiGroup)

type apiGroup struct {
	AuthApi        authApi
	ChannelInfoApi channelInfoApi
	ChannelPlanApi channelPlanApi
	PushApi        pushApi
	WsApi          wsApi
}
