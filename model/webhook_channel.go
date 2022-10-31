package model

type WebhookChannel struct {
	Url string `json:"url" validate:"url"`
}
