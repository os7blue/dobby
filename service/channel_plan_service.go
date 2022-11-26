package service

import (
	"dobby/common"
	"dobby/model"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type channelPlanService struct {
}

func (cps *channelPlanService) getOne(id uint) model.ChannelPlanView {
	data, exist := common.LocalCache.Get(fmt.Sprintf("plan-%d", id))
	if exist {
		return data.(model.ChannelPlanView)
	}

	var plan model.ChannelPlanView
	common.DB.Model(&model.ChannelPlan{}).Select("*").Where("id = ?", id).Find(&plan)
}

func (cps *channelPlanService) Create(plan model.ChannelPlan) error {

	plan.CreateTime = time.Now().UnixMilli()
	plan.Key = uuid.New().String()

	err := common.DB.Create(&plan)

	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (cps *channelPlanService) LoadPage(page int, limit int, name string) (interface{}, int64, error) {

	d := common.DB.Model(&model.ChannelPlan{}).Select("*,(SELECT '[' || GROUP_CONCAT(json_object ('id',id,'name',name,'channelType',channel_type,'optionJsonStr',option_json_str)) || ']' FROM channel_info i WHERE channel_id_list_str LIKE '%' || i.id || ',%' OR channel_id_list_str LIKE '%,' || i.id || ',' OR channel_id_list_str LIKE '%' || i.id) AS channel_info_list_json_str")
	c := common.DB.Model(&model.ChannelPlan{})

	if "" != name {
		like := fmt.Sprintf("%%%s%%", name)

		d.Where("name LIKE ?", like).Or("key like ?", like).Or("(SELECT count(1) FROM channel_info WHERE id in (channel_plan.id) and name like %?%) > 0", name)
		c.Where("name LIKE ?", like).Or("key like ?", like).Or("(SELECT count(1) FROM channel_info WHERE id in (channel_plan.id) and name like %?%) > 0", name)
	}

	var results []model.ChannelPlanView
	err := d.Offset((page - 1) * limit).Limit(limit).Order("create_time DESC").Find(&results)
	if err.Error != nil {
		return nil, 0, errors.New("查找失败")
	}
	var count int64
	c.Count(&count)

	return results, count, nil

}

func (cps *channelPlanService) Update(plan model.ChannelPlan) error {

	err := common.DB.Updates(plan)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (cps *channelPlanService) Delete(id uint) error {

	err := common.DB.Delete(&model.ChannelPlan{}, id)
	if err.Error != nil {
		return err.Error
	}
	return nil

}

func (cps *channelPlanService) GetOne(id uint) (model.ChannelPlan, error) {

	var result model.ChannelPlan
	err := common.DB.First(&result, id)
	if err.Error != nil {
		return result, err.Error
	}
	return result, nil
}

func (cps *channelPlanService) RefreshKey(id uint) (string, error) {

	key := uuid.New().String()
	err := common.DB.Model(&model.ChannelPlan{}).Where("id=?", id).Update("key", key)
	if err.Error != nil {
		return "", err.Error
	}

	return key, nil
}
