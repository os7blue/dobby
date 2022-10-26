package model

type WebHookChannel struct {
	Url       string `json:"url" bind:"url"`
	ChannelId uint
}
