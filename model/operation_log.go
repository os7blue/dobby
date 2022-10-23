package model

type OperationLog struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Content    string `gorm:"not null" json:"content"`
	CreateTime int64  `gorm:"not null;autoCreateTime=milli" json:"createTime"`
}
