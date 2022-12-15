package sender

import (
	"dobby/common"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"strings"
)

type wxMpSender struct {
}

func (w *wxMpSender) Send(AppId string, appSecret string, templateId string, to string, title string, content string) error {

	token := ""
	data, exist := common.LocalCache.Get(fmt.Sprintf("mpt-%s", AppId))

	if exist {
		j := gjson.Get(data.(string), "access_token")
		token = j.Str

	} else {
		getAccessTokenUrl := fmt.Sprintf(
			"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
			AppId,
			appSecret,
		)
		body, err := common.GlobalUtil.SendSimpleGet(getAccessTokenUrl)
		if err != nil {
			return err
		}
		t := gjson.Get(body, "access_token")
		o := gjson.Get(body, "expires_in")
		if !t.Exists() || !o.Exists() {
			return errors.New(fmt.Sprintf("推送失败，%s", body))
		}
		token = t.String()
		common.LocalCache.Set(
			fmt.Sprintf("mpt-%s", AppId),
			body,
			o.Int()-100,
		)

	}

	sendMagUrl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token)

	userTitle := common.Option.Setting.PushTitle
	if userTitle != "" {

		title = fmt.Sprintf("[%s]%s：", userTitle, title)

	} else {
		title += "："
	}

	tempMap := map[string]any{
		"template_id": templateId,
		"data": map[string]any{
			"content": map[string]any{
				"value": title + content,
			},
		},
	}
	tos := strings.Split(to, ",")
	for _, t := range tos {

		tempMap["touser"] = t

		body, err := common.GlobalUtil.SendSimplePost(sendMagUrl, tempMap)
		if err != nil {
			return err
		}
		ec := gjson.Get(body, "errcode")
		if ec.Int() != 0 {
			return errors.New(gjson.Get(body, "errmsg").String())
		}

	}
	return nil
}
