package model

type ChannelPlan struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null;size:20;unique" json:"name"`
	//status code , enable is 10 , disabled is 20
	Status           int    `gorm:"not null" json:"status"`
	Key              string `gorm:"not null;unique" json:"key"`
	CreateTime       int64  `gorm:"not null;autoCreateTime:milli" json:"createTime"`
	WhiteListStr     string `json:"whiteListStr"`
	ChannelIdListStr string `json:"ChannelIdListStr"`
}

type ChannelPlanView struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	//status code , enable is 10 , disabled is 20
	Status                 int    `json:"status"`
	Key                    string `json:"key"`
	CreateTime             int64  `json:"createTime"`
	WhiteListStr           string `json:"whiteListStr"`
	ChannelIdListStr       string `json:"ChannelIdListStr"`
	ChannelInfoListJsonStr string `json:"channelInfoListJsonStr"`
}

type ChannelPlanCreateValidator struct {
	Name             string `json:"name"`
	WhiteListStr     string `json:"whiteListStr" binding:"required"`
	ChannelIdListStr string `json:"channelIdListStr" binding:"required"`
}

type ChannelPlanUpdateValidator struct {
	ID               uint   `json:"id" bind:"required"`
	Name             string `json:"name"`
	Status           int    `json:"status"`
	WhiteListStr     string `json:"whiteListStr"`
	ChannelIdListStr string `json:"ChannelIdListStr"`
}

type ChannelPlanSearchValidator struct {
	Name string `json:"name" binding:"omitempty"`
}
