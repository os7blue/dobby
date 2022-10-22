package model

//item -> line -> info

type PushChannelInfo struct {
	Id         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"not null;size:20;unique" json:"name"`
	Status     int    `gorm:"not null;size:2" json:"status"`
	Key        string `gorm:"not null;unique" json:"key"`
	CreateTime uint   `gorm:"not null;autoCreateTime:milli" json:"createTime"`
}
