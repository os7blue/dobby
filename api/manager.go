package api

var Apis = new(apiGroup)

type apiGroup struct {
	AuthApi        authApi
	ChannelInfoApi channelInfoApi
}
