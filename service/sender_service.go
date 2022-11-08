package service

import (
	"dobby/common"
	"dobby/model"
	"dobby/model/constant"
	"dobby/service/sender"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type sendService struct {
}

func (s *sendService) Send(info model.PushInfo) (string, error) {

	var channels []model.ChannelInfo
	err := common.DB.Model(&model.ChannelInfo{}).Where("key in (?)", strings.Split(info.KeyStr, ",")).Find(&channels)
	if err.Error != nil {
		return "", err.Error
	}
	if len(channels) != len(strings.Split(info.KeyStr, ",")) {
		return "", errors.New("请求key数量与实际通道数不匹配")
	}

	msgs := ""
	for _, channel := range channels {

		//check channel status
		if channel.Status != 10 {
			msgs = fmt.Sprintf("%s [通道:%s，key=%s，推送失败，原因：通道为未启用状态]", msgs, channel.Name, channel.Key)
			continue
		}

		//check request source ip
		if channel.WhiteListStr != "0.0.0.0" {

			ips := strings.Split(channel.WhiteListStr, ",")
			has := false
			for _, ip := range ips {

				if ip == info.Ip {
					has = true
					break
				}

			}

			if !has {
				msgs = fmt.Sprintf("%s [通道:%s，key=%s，推送失败，原因：来源非法]", msgs, channel.Name, channel.Key)
				break
			}

		}

		//dispatching different channel request
		switch channel.ChannelType {
		case constant.WEBHOOK:
			c := model.WebhookChannel{}
			err := json.Unmarshal([]byte(channel.OptionJsonStr), &c)
			if err != nil {
				msgs = fmt.Sprintf("%s [通道：%s，key=%s，参数加载失败]", msgs, channel.Name, channel.Key)
				break
			}
			err = sender.Senders.WebhookSender.Send(c.Url, info.Content, c.HookType)
			if err != nil {
				msgs = fmt.Sprintf("%s [通道：%s，key=%s，推送失败，失败原因：%s]", msgs, channel.Name, channel.Key, err.Error())

			}
			break
		case constant.EMAIL:

			c := model.EmailChannel{}
			err := json.Unmarshal([]byte(channel.OptionJsonStr), &c)
			if err != nil {
				msgs = fmt.Sprintf("%s [通道：%s，key=%s，参数加载失败]", msgs, channel.Name, channel.Key)
				break
			}
			err = sender.Senders.MailSender.Send(
				c.Host,
				c.Port,
				c.Username,
				c.Password,
				strings.Split(c.ToEmailListStr, ","),
				common.Option.Setting.PushTitle,
				info.Content,
			)
			if err != nil {
				msgs = fmt.Sprintf("%s [通道:%s，key=%s，推送失败，原因：%s]", msgs, channel.Name, channel.Key, err.Error())
			}
			break
		case constant.WXMP:
			c := model.WxMpChannel{}
			err := json.Unmarshal([]byte(channel.OptionJsonStr), &c)
			if err != nil {
				msgs = fmt.Sprintf("%s [通道：%s，key=%s，参数加载失败]", msgs, channel.Name, channel.Key)
				break
			}
			err = sender.Senders.WxMpSender.Send(
				c.AppId,
				c.AppSecret,
				c.TemplateId,
				c.ToUserListStr,
				info.Content,
			)
			if err != nil {
				msgs = fmt.Sprintf("%s [通道:%s，key=%s，推送失败，原因：%s]", msgs, channel.Name, channel.Key, err.Error())
			}
			break

		}

	}

	return msgs, nil
}
