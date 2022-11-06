package service

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"message-push/common"
	"message-push/model"
	"message-push/model/constant"
	"message-push/service/sender"
	"strings"
)

type sendService struct {
}

func (s *sendService) Send(info model.PushInfo) (string, error) {

	var channels []model.ChannelInfo
	err := common.DB.Model(&model.ChannelInfo{}).Where("key in (?)", info.KeyStr).Find(&channels)
	if err.Error != nil {
		return "", err.Error
	}
	if len(channels) != len(strings.Split(info.KeyStr, ",")) {
		return "", errors.New("请求key数量与实际通道数不匹配")

	}

	msgs := ""
	for _, channel := range channels {
		if channel.Status != 10 {
			msgs = fmt.Sprintf("%s [通道:%s，key=%s，推送失败，原因：通道为未启用状态]", msgs, channel.Name, channel.Key)
			continue
		}
		switch channel.ChannelType {
		case constant.WEBHOOK:
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
				info.Content,
				info.Content,
			)
			if err != nil {
				msgs = fmt.Sprintf("%s [通道:%s，key=%s，推送失败，原因：%s]", msgs, channel.Name, channel.Key, err.Error())
			}
			break
		case constant.WXMP:
			break

		}

	}

	return msgs, nil
}
