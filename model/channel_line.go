package model

type ChannelLine struct {
	ID          uint `gorm:"primaryKey;autoIncrement" json:"id"`
	InfoId      uint `gorm:"not null" json:"infoId"`
	ChannelType int  `gorm:"not null;size:1;" json:"channelType"`
	ItemId      uint `gorm:"not null" json:"itemId"`
}