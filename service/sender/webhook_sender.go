package sender

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"dobby/common"
	"dobby/model"
	"dobby/model/constant"
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
		rpBody, err := common.GlobalUtil.SendSimplePost(url, body)
		if err != nil {
			return err
		}
		g := gjson.Parse(rpBody)

		if !g.Get("errcode").Exists() || g.Get("errcode").Int() != 0 {
			return errors.New(rpBody)
		}

		break
	case constant.FEI_SHU_HOOK:
		body := model.FeiShuBody{
			MsgType: "text",
		}

		body.Content.Text = fmt.Sprintf("<at user_id=\"all\"></at>%s%s", common.Option.Setting.PushTitle, content)
		rpBody, err := common.GlobalUtil.SendSimplePost(url, body)
		if err != nil {
			return err
		}
		g := gjson.Parse(rpBody)

		if !g.Get("StatusCode").Exists() || g.Get("StatusCode").Int() != 0 {
			return errors.New(rpBody)
		}

		break
	case constant.WECHAT_WORK_HOOK:
		body := model.WxWorkBody{
			Msgtype: "text",
		}
		body.Text.Content = common.Option.Setting.PushTitle + content
		body.Text.MentionedMobileList = []string{"@all"}

		rpBody, err := common.GlobalUtil.SendSimplePost(url, body)
		if err != nil {
			return err
		}
		g := gjson.Parse(rpBody)

		if !g.Get("errcode").Exists() || g.Get("errcode").Int() != 0 {
			return errors.New(rpBody)
		}

		break

	}

	return nil
}
