package model

type WxMpChannel struct {
	AppId         string `json:"appId" validate:"not null"`
	AppSecret     string `json:"appSecret" validate:"not null"`
	TemplateId    string `json:"templateId" validate:"not null"`
	ToUserListStr string `json:"toUserListStr" validate:"not null"`
}
