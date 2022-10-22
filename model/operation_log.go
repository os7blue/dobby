package model

type OperationLog struct {
	Id         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Content    string `gorm:"not null" json:"content"`
	CreateTime uint   `gorm:"not null;autoCreateTime=milli" json:"createTime"`
}
