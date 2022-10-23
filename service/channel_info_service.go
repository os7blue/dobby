package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"message-push/common"
	"message-push/model"
	"time"
)

type channelInfoService struct {
}

func (s *channelInfoService) CreateOne(name string) error {

	var count int64
	common.DB.Model(&model.PushChannelInfo{}).Where("name=?", name).Count(&count)
	if count != 0 {
		return errors.New(fmt.Sprintf("通道名称《%s》已存在", name))
	}

	info := model.PushChannelInfo{
		Name:       name,
		Status:     0,
		Key:        uuid.New().String(),
		CreateTime: time.Now().UnixMilli(),
	}
	err := common.DB.Create(&info)
	if err.Error != nil {
		return errors.New("创建失败")
	}

	return nil
}
