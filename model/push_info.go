package model

type PushInfo struct {
	KeyStr  string `json:"keyStr" binding:"required"`
	Content string `json:"content" binding:"required"`
	Ip      string `json:"ip"`
}
