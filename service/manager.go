package service

var Services = new(serviceGroup)

type serviceGroup struct {
	AuthService        authService
	ChannelInfoService channelInfoService
	ChannelPlanService channelPlanService
	SendService        sendService
}
