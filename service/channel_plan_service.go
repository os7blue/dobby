package service

import (
	"dobby/common"
	"dobby/model"
)

type channelPlanService struct {
}

func (c *channelPlanService) Create(plan model.ChannelPlan) error {

	finalStr := common.GlobalUtil.ArrStrItemFormat(
		plan.ChannelListStr,
		"{%s}",
		",",
		",",
	)
	plan.ChannelListStr = finalStr

	err := common.DB.Create(&plan)

	if err.Error != nil {
		return err.Error
	}
	return nil
}
