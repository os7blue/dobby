package model

//item -> line -> info

type ChannelInfo struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"not null;size:20;unique" json:"name"`
	CreateTime    int64  `gorm:"not null;autoCreateTime:milli" json:"createTime"`
	ChannelType   int    `gorm:"not null" json:"channelType"`
	OptionJsonStr string `gorm:"not null" json:"optionJsonStr"`
}

type ChannelInfoView[T any] struct {
	ID            uint   ` json:"id"`
	Name          string ` json:"name"`
	CreateTime    int64  `json:"createTime"`
	ChannelType   int    `json:"channelType"`
	ChannelOption T      `json:"channelOption"`
}

type ChannelInfoCreateValidator struct {
	Name          string `json:"name" binding:"required,min=2,max=20"`
	ChannelType   int    `json:"channelType" bind:"required"`
	OptionJsonStr string `json:"optionJsonStr" binding:"required"`
}

type ChannelInfoUpdateValidator struct {
	ID            uint   `json:"id" binding:"required"`
	Name          string `json:"name" binding:"omitempty,min=2,max=20"`
	OptionJsonStr string `json:"optionJsonStr"`
}

type ChannelInfoSearchValidator struct {
	Name string `json:"name"`
}
