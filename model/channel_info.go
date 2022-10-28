package model

//item -> line -> info

type ChannelInfo struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null;size:20;unique" json:"name"`
	//status code , enable is 10 , disabled is 20
	Status        int    `gorm:"not null;size:2" json:"status"`
	Key           string `gorm:"not null;unique" json:"key"`
	CreateTime    int64  `gorm:"not null;autoCreateTime:milli" json:"createTime"`
	WhiteList     string `json:"whiteList"`
	ChannelType   int    `gorm:"not null" json:"channelType"`
	OptionJsonStr string `gorm:"not null" json:"optionJsonStr"`
}

type ChannelInfoCreateValidator struct {
	Name        string `json:"name" binding:"required,min=2,max=20"`
	WhiteList   string `json:"whiteList"`
	ChannelType int    `gorm:"not null" json:"channelType"`
}

type ChannelInfoUpdateValidator struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"omitempty,min=2,max=20"`
	WhiteList   string `json:"whiteList"`
	ChannelType int    `json:"channelType"`
	Status      int    `json:"status"`
}

type ChannelInfoSearchValidator struct {
	Name string `json:"name" binding:"omitempty"`
}
