package model

type PushInfo struct {
	Key     string `json:"key" binding:"required"`
	Title   string `json:"title"`
	Content string `json:"content" binding:"required"`
}
