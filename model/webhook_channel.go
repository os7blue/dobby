package model

type WebhookChannel struct {
	Url      string `json:"url" validate:"url"`
	HookType int    `json:"hookType" validate:"required"`
}
