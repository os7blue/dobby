package model

//item -> line -> info

type PushChannelInfo struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"not null;size:20;unique" json:"name"`
	Status     int    `gorm:"not null;size:2" json:"status"`
	Key        string `gorm:"not null;unique" json:"key"`
	CreateTime int64  `gorm:"not null;autoCreateTime:milli" json:"createTime"`
}

type PushChannelInfoCreateValidator struct {
	Name string `json:"name" binding:"required,min=2,max=20"`
}

type PushChannelInfoUpdateValidator struct {
	ID     uint   `json:"id" binding:"required"`
	Name   string `json:"name" binding:"min=2,max=20"`
	Status int    `json:"status"`
}
