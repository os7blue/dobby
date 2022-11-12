package model

type ChannelPlan struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null;size:20;unique" json:"name"`
	//status code , enable is 10 , disabled is 20
	Key          string `gorm:"not null;unique" json:"key"`
	CreateTime   int64  `gorm:"not null;autoCreateTime:milli" json:"createTime"`
	WhiteListStr string `json:"whiteListStr"`
}
