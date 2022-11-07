package sender

import (
	"fmt"
	"message-push/common"
	"message-push/model"
	"message-push/model/constant"
)

type webhookSender struct {
}

func (w *webhookSender) Send(url string, content string, hookType int) error {

	switch hookType {
	case constant.DING_TALK_HOOK:
		body := model.DingTalkBody{
			MsgType: "text",
		}
		body.At.IsAtAll = true
		body.Text.Content = fmt.Sprintf("%s%s", common.Option.Setting.PushTitle, content)
		_, err := common.GlobalUtil.SendSimplePost(url, body)
		if err != nil {
			return err
		}

		break
	case constant.FEI_SHU_HOOK:
		body := model.FeiShuBody{
			MsgType: "text",
		}

		body.Content.Text = fmt.Sprintf("<at user_id=\"all\"></at>%s%s", common.Option.Setting.PushTitle, content)
		_, err := common.GlobalUtil.SendSimplePost(url, body)
		if err != nil {
			return err
		}

		break
	case constant.WECHAT_WORK_HOOK:
		body := model.WxWorkBody{
			Msgtype: "text",
		}
		body.Text.Content = common.Option.Setting.PushTitle + content
		body.Text.MentionedMobileList = []string{"@all"}

		_, err := common.GlobalUtil.SendSimplePost(url, body)
		if err != nil {
			return err
		}

		break

	}

	return nil
}
