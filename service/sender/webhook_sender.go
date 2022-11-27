package sender

import (
	"dobby/common"
	"dobby/model"
	"dobby/model/constant"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type webhookSender struct {
}

func (w *webhookSender) Send(url string, title string, content string, hookType int) error {

	userTitle := common.Option.Setting.PushTitle
	if title != "" {

		userTitle = fmt.Sprintf("%s[%s]：", userTitle, title)

	}

	switch hookType {
	case constant.DING_TALK_HOOK:
		body := model.DingTalkBody{
			MsgType: "text",
		}
		body.At.IsAtAll = true
		body.Text.Content = fmt.Sprintf("%s%s", userTitle, content)
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

		body.Content.Text = fmt.Sprintf("<at user_id=\"all\"></at>%s%s", userTitle, content)
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
		body.Text.Content = userTitle + content
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
