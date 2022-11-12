package model

//item -> line -> info

type ChannelInfo struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null;size:20;unique" json:"name"`
	//status code , enable is 10 , disabled is 20
	Key           string `gorm:"not null;unique" json:"key"`
	CreateTime    int64  `gorm:"not null;autoCreateTime:milli" json:"createTime"`
	WhiteListStr  string `json:"whiteListStr"`
	ChannelType   int    `gorm:"not null" json:"channelType"`
	OptionJsonStr string `gorm:"not null" json:"optionJsonStr"`
}

type ChannelInfoView[T any] struct {
	ID   uint   ` json:"id"`
	Name string ` json:"name"`
	//status code , enable is 10 , disabled is 20
	Status        int    `json:"status"`
	Key           string `json:"key"`
	CreateTime    int64  `json:"createTime"`
	WhiteListStr  string `json:"whiteListStr"`
	ChannelType   int    `json:"channelType"`
	ChannelOption T      `json:"channelOption"`
}

type ChannelInfoCreateValidator struct {
	Name          string `json:"name" binding:"required,min=2,max=20"`
	WhiteListStr  string `json:"whiteListStr" binding:"required"`
	ChannelType   int    `json:"channelType" bind:"required"`
	OptionJsonStr string `json:"optionJsonStr" binding:"required"`
}

type ChannelInfoUpdateValidator struct {
	ID            uint   `json:"id" binding:"required"`
	Name          string `json:"name" binding:"omitempty,min=2,max=20"`
	WhiteListStr  string `json:"whiteListStr"`
	Status        int    `json:"status"`
	OptionJsonStr string `json:"optionJsonStr"`
}

type ChannelInfoSearchValidator struct {
	Name string `json:"name" binding:"omitempty"`
}
