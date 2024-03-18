package service

import (
	"dobby/model"
	"dobby/model/constant"
	"dobby/service/sender"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

type sendService struct {
}

func (s *sendService) Send(key string, title string, content string, ip string) (string, error) {

	plan, err := ChannelPlanService.getOneByKey(key)
	if err != nil {
		return "", err
	}

	//check channel status
	if plan.Status != 10 {
		return "", errors.New(fmt.Sprintf("方案[%s]推送失败，原因：方案为未启用状态]", plan.Name))
	}

	//check request source ip
	if plan.WhiteListStr != "0.0.0.0" {

		ips := strings.Split(plan.WhiteListStr, ",")
		has := false
		for _, p := range ips {

			if p == ip {
				has = true
				break
			}

		}

		if !has {
			return "", errors.New(fmt.Sprintf("[方案[%s]推送失败，原因：来源非法]", plan.Name))
		}

	}

	msgs := ""
	var wait sync.WaitGroup
	wait.Add(len(plan.ChannelInfoList))

	for _, channel := range plan.ChannelInfoList {

		//dispatching different channel request
		switch channel.ChannelType {
		case constant.WEBHOOK:
			go func(channel model.ChannelInfo, title string, content string) {
				c := model.WebhookChannel{}
				err := json.Unmarshal([]byte(channel.OptionJsonStr), &c)
				if err != nil {
					msgs = fmt.Sprintf("%s [通道：%s，参数加载失败]", msgs, channel.Name)
					wait.Done()
					return
				}
				err = sender.WebhookSender.Send(c.Url, title, content, c.HookType)
				if err != nil {
					msgs = fmt.Sprintf("%s [通道：%s，推送失败，失败原因：%s]", msgs, channel.Name, err.Error())

				}
				wait.Done()

			}(channel, title, content)
			break
		case constant.EMAIL:

			go func(channel model.ChannelInfo, title string, content string) {
				c := model.EmailChannel{}
				err := json.Unmarshal([]byte(channel.OptionJsonStr), &c)
				if err != nil {
					msgs = fmt.Sprintf("%s [通道：%s，参数加载失败]", msgs, channel.Name)
					wait.Done()

					return
				}
				err = sender.MailSender.Send(
					c.Host,
					c.Port,
					c.Username,
					c.Password,
					strings.Split(c.ToEmailListStr, ","),
					title,
					content,
				)
				if err != nil {
					msgs = fmt.Sprintf("%s [通道:%s，推送失败，原因：%s]", msgs, channel.Name, err.Error())

				}
				wait.Done()
			}(channel, title, content)

			break
		case constant.WXMP:

			go func(channel model.ChannelInfo, title string, content string) {
				c := model.WxMpChannel{}
				err := json.Unmarshal([]byte(channel.OptionJsonStr), &c)
				if err != nil {
					msgs = fmt.Sprintf("%s [通道：%s，参数加载失败]", msgs, channel.Name)
					wait.Done()

					return
				}
				err = sender.WxMpSender.Send(
					c.AppId,
					c.AppSecret,
					c.TemplateId,
					c.ToUserListStr,
					title,
					content,
				)
				if err != nil {
					msgs = fmt.Sprintf("%s [通道:%s，推送失败，原因：%s]", msgs, channel.Name, err.Error())

				}
				wait.Done()
			}(channel, title, content)

			break

		case constant.WS:

			go func(channel model.ChannelInfo, title string, content string) {

				err := sender.WsSender.Send(channel.ID, title, content)
				if err != nil {
					msgs = fmt.Sprintf("%s [通道:%s，推送失败，原因：%s]", msgs, channel.Name, err.Error())

				}
				wait.Done()
			}(channel, title, content)

		}

	}

	wait.Wait()

	return msgs, nil
}
