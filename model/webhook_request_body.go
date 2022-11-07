package model

type WxWorkBody struct {
	Msgtype string         `json:"msgtype"`
	Text    WxWorkBodyText `json:"text"`
}

type WxWorkBodyText struct {
	Content             string   `json:"content"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type DingTalkBody struct {
	At      DingTalkBodyAt   `json:"at"`
	Text    DingTalkBodyText `json:"text"`
	MsgType string           `json:"msgtype"`
}

type DingTalkBodyAt struct {
	IsAtAll bool `json:"isAtAll"`
}

type DingTalkBodyText struct {
	Content string `json:"content"`
}

type FeiShuBody struct {
	MsgType string            `json:"msg_type"`
	Content FeiShuBodyContent `json:"content"`
}

type FeiShuBodyContent struct {
	Text string `json:"text"`
}
