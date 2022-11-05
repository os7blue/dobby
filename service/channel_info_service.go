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

func (s *channelInfoService) CreateOne(info model.ChannelInfo) error {

	var count int64
	common.DB.Model(&model.ChannelInfo{}).Where("name=?", info.Name).Count(&count)
	if count != 0 {
		return errors.New(fmt.Sprintf("通道名称《%s》已存在", info.Name))
	}

	info.CreateTime = time.Now().UnixMilli()
	info.Status = 20
	info.Key = uuid.New().String()

	err := common.DB.Create(&info)
	if err.Error != nil {
		return errors.New("创建失败")
	}

	return nil
}

func (s *channelInfoService) Update(info model.ChannelInfo) error {

	err := common.DB.Model(&model.ChannelInfo{}).Where("id=?", info.ID).Updates(info)
	if err.Error != nil {
		return errors.New("修改失败")
	}

	return nil
}

func (s *channelInfoService) Delete(id uint) error {

	err := common.DB.Delete(&model.ChannelInfo{}, id)
	if err.Error != nil {
		return errors.New("删除失败")
	}
	return nil
}

func (s *channelInfoService) GetOne(id uint) (model.ChannelInfo, error) {

	var result model.ChannelInfo
	err := common.DB.First(&result, id)
	if err.Error != nil {
		return result, err.Error
	}

	return result, nil
}
