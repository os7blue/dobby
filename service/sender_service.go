package service

import (
	"dobby/model"
	"dobby/model/constant"
	"dobby/service/sender"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
)

// ================== Sender 接口 ==================
type Sender interface {
	Send(channel model.ChannelInfo, title, content string) string
}

// ================== Sender 自动注册 ==================
var senderRegistry = make(map[int]Sender)

func RegisterSender(channelType int, sender Sender) {
	senderRegistry[channelType] = sender
}

// ================== 各通道实现 ==================
type WebhookSender struct{}

func (s WebhookSender) Send(channel model.ChannelInfo, title, content string) string {
	c := model.WebhookChannel{}
	if err := json.Unmarshal([]byte(channel.OptionJsonStr), &c); err != nil {
		return fmt.Sprintf("[通道：%s，参数加载失败]", channel.Name)
	}
	if err := sender.WebhookSender.Send(c.Url, title, content, c.HookType); err != nil {
		return fmt.Sprintf("[通道：%s，推送失败，原因：%s]", channel.Name, err.Error())
	}
	return ""
}

type EmailSender struct{}

func (s EmailSender) Send(channel model.ChannelInfo, title, content string) string {
	c := model.EmailChannel{}
	if err := json.Unmarshal([]byte(channel.OptionJsonStr), &c); err != nil {
		return fmt.Sprintf("[通道：%s，参数加载失败]", channel.Name)
	}
	if err := sender.MailSender.Send(
		c.Host, c.Port, c.Username, c.Password,
		strings.Split(c.ToEmailListStr, ","),
		title, content,
	); err != nil {
		return fmt.Sprintf("[通道：%s，推送失败，原因：%s]", channel.Name, err.Error())
	}
	return ""
}

type WxMpSender struct{}

func (s WxMpSender) Send(channel model.ChannelInfo, title, content string) string {
	c := model.WxMpChannel{}
	if err := json.Unmarshal([]byte(channel.OptionJsonStr), &c); err != nil {
		return fmt.Sprintf("[通道：%s，参数加载失败]", channel.Name)
	}
	if err := sender.WxMpSender.Send(
		c.AppId, c.AppSecret, c.TemplateId, c.ToUserListStr,
		title, content,
	); err != nil {
		return fmt.Sprintf("[通道：%s，推送失败，原因：%s]", channel.Name, err.Error())
	}
	return ""
}

type WsSender struct{}

func (s WsSender) Send(channel model.ChannelInfo, title, content string) string {
	if err := sender.WsSender.Send(channel.ID, title, content); err != nil {
		return fmt.Sprintf("[通道：%s，推送失败，原因：%s]", channel.Name, err.Error())
	}
	return ""
}

// ================== 自动注册通道 ==================
func init() {
	RegisterSender(constant.WEBHOOK, WebhookSender{})
	RegisterSender(constant.EMAIL, EmailSender{})
	RegisterSender(constant.WXMP, WxMpSender{})
	RegisterSender(constant.WS, WsSender{})
}

// ================== sendService ==================
type sendService struct{}

const (
	ModeConcurrent int = 1 // 并发
)

func (s *sendService) Send(key, title, content, ip string) (string, error) {
	plan, err := ChannelPlanService.getOneByKey(key)
	if err != nil {
		return "", err
	}

	if plan.Status != 10 {
		return "", errors.Errorf("方案[%s]推送失败，原因：方案未启用", plan.Name)
	}

	if plan.WhiteListStr != "0.0.0.0" {
		ips := strings.Split(plan.WhiteListStr, ",")
		if !ipAllowed(ips, ip) {
			return "", errors.Errorf("方案[%s]推送失败，来源非法", plan.Name)
		}
	}

	msgChan := make(chan string, len(plan.ChannelInfoList))

	if plan.ChannelModel == ModeConcurrent {
		// 并发执行
		for _, ch := range plan.ChannelInfoList {
			go func(channel model.ChannelInfo) {
				msgChan <- runSender(channel, title, content)
			}(ch)
		}
	} else {
		// 顺序执行
		for _, ch := range plan.ChannelInfoList {
			msgChan <- runSender(ch, title, content)
		}
	}

	var msgList []string
	for i := 0; i < len(plan.ChannelInfoList); i++ {
		msg := <-msgChan
		if msg != "" {
			msgList = append(msgList, msg)
		}
	}
	close(msgChan)

	if len(msgList) == 0 {
		return "全部发送成功", nil
	}
	return strings.Join(msgList, "; "), nil
}

func ipAllowed(whitelist []string, ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false // 无效 IP
	}

	for _, w := range whitelist {
		w = strings.TrimSpace(w)
		if strings.Contains(w, "/") { // CIDR 段
			_, ipnet, err := net.ParseCIDR(w)
			if err != nil {
				continue
			}
			if ipnet.Contains(ip) {
				return true
			}
		} else { // 单个 IP
			if ip.Equal(net.ParseIP(w)) {
				return true
			}
		}
	}
	return false
}

// runSender 根据通道类型调用对应 Sender
func runSender(channel model.ChannelInfo, title, content string) string {
	senderObj, ok := senderRegistry[channel.ChannelType]
	if !ok {
		return fmt.Sprintf("[通道：%s，未知通道类型]", channel.Name)
	}
	return senderObj.Send(channel, title, content)
}

// ================== 工具函数 ==================
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
