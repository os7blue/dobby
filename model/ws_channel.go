package model

type WebSocketChannel struct {
	Key string `json:"key" validate:"required,min=2"`
}
