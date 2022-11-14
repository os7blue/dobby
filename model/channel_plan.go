package model

type ChannelPlan struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null;size:20;unique" json:"name"`
	//status code , enable is 10 , disabled is 20
	Key            string `gorm:"not null;unique" json:"key"`
	CreateTime     int64  `gorm:"not null;autoCreateTime:milli" json:"createTime"`
	WhiteListStr   string `json:"whiteListStr"`
	ChannelListStr string `json:"channelListStr"`
}

type ChannelPlanCreateValidator struct {
	Name           string `json:"name"`
	WhiteListStr   string `json:"whiteListStr" binding:"required"`
	ChannelListStr string `json:"channelListStr" binding:"required"`
}

type ChannelPlanUpdateValidator struct {
	ID             uint   `json:"id" bind:"required"`
	Name           string `json:"name"`
	WhiteListStr   string `json:"whiteListStr" binding:"required"`
	ChannelListStr string `json:"channelListStr" binding:"required"`
}
