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

func (s *channelInfoService) LoadPage(page int, limit int, channelInfo model.ChannelInfoSearchValidator) ([]model.ChannelInfo, int64, error) {

	d := common.DB.Model(&model.ChannelInfo{})
	c := common.DB.Model(&model.ChannelInfo{})
	d.Select("*")

	if "" != channelInfo.Name {
		d.Where("name LIKE %?%", channelInfo.Name).Or("key like %?%", channelInfo.Name)
		c.Where("name LIKE %?%", channelInfo.Name).Or("key like %?%", channelInfo.Name)
	}

	var results []model.ChannelInfo
	err := d.Offset((page - 1) * limit).Limit(limit).Order("create_time DESC").Find(&results)
	if err.Error != nil {
		return nil, 0, errors.New("查找失败")
	}
	var count int64
	c.Count(&count)

	return results, count, nil
}

func (s *channelInfoService) CreateOne(name string) error {

	var count int64
	common.DB.Model(&model.ChannelInfo{}).Where("name=?", name).Count(&count)
	if count != 0 {
		return errors.New(fmt.Sprintf("通道名称《%s》已存在", name))
	}

	info := model.ChannelInfo{
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
