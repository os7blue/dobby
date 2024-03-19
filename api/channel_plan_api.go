package api

import (
	"dobby/common"
	"dobby/model"
	"dobby/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type channelPlanApi struct {
}

func (p *channelPlanApi) Create(c *gin.Context) {

	var createVm model.ChannelPlanCreateValidator
	err := c.ShouldBindJSON(&createVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = common.GlobalUtil.CheckArrStr(
		1,
		0,
		common.IpV4Regx,
		createVm.WhiteListStr,
	)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = common.GlobalUtil.CheckArrStr(
		1,
		0,
		"",
		createVm.ChannelIdListStr,
	)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	var plan model.ChannelPlan
	err = common.StructConvert[model.ChannelPlanCreateValidator, model.ChannelPlan](createVm, &plan)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	err = service.ChannelPlanService.Create(plan)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)
}

func (p *channelPlanApi) Load(c *gin.Context) {

	var loadVm model.PageBody[model.ChannelPlanSearchValidator]
	err := c.ShouldBindJSON(&loadVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	result, count, err := service.ChannelPlanService.LoadPage(loadVm.Page, loadVm.Limit, loadVm.Param.Name)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.SuccessWithDataCount(c, result, count)

}

func (p *channelPlanApi) Update(c *gin.Context) {

	var updateVm model.ChannelPlanUpdateValidator
	err := c.ShouldBindJSON(&updateVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = common.GlobalUtil.CheckArrStr(
		1,
		0,
		common.IpV4Regx,
		updateVm.WhiteListStr,
	)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = common.GlobalUtil.CheckArrStr(
		1,
		0,
		"",
		updateVm.ChannelIdListStr,
	)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	var plan model.ChannelPlan
	err = common.StructConvert[model.ChannelPlanUpdateValidator, model.ChannelPlan](updateVm, &plan)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	err = service.ChannelPlanService.Update(plan)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)

}

func (p *channelPlanApi) ChangeStatus(c *gin.Context) {

	var updateVm model.ChannelPlanUpdateValidator

	err := c.ShouldBindJSON(&updateVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	_, err = service.ChannelPlanService.GetOne(updateVm.ID)
	if err != nil {
		common.R.FailWithMsg(c, fmt.Sprintf("id为 %d 的通道方案可能不存在", updateVm.ID))
		return
	}

	err = service.ChannelPlanService.Update(model.ChannelPlan{
		ID:     updateVm.ID,
		Status: updateVm.Status,
	})

	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	common.R.Success(c)

}

func (p *channelPlanApi) Delete(c *gin.Context) {

	var updateVm model.ChannelPlanUpdateValidator
	err := c.ShouldBindJSON(&updateVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = service.ChannelPlanService.Delete(updateVm.ID)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)

}

func (p *channelPlanApi) RefreshKey(c *gin.Context) {

	var updateVm model.ChannelPlanUpdateValidator
	err := c.ShouldBindJSON(&updateVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	key, err := service.ChannelPlanService.RefreshKey(updateVm.ID)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.SuccessWithData(c, key)

}
