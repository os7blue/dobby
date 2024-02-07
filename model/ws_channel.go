package model

// 只需要一个key
type WebSocketChannel struct {
	Key string `json:"key" validate:"required,min=2"`
}
