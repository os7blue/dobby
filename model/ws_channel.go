package model

type WsChannel struct {
	Key string `json:"key" validate:"required,min=2"`
}

type WsConnView struct {
	ID  uint   `validate:"required"`
	Key string `validate:"required,min=2"`
}

type PushWs struct {
	ID      uint   `json:"id" validate:"required"`
	Content string `json:"content" validate:"required"`
}
