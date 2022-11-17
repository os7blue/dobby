package api

import (
	"dobby/common"
	"dobby/model"
	"dobby/service"
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
		createVm.ChannelListStr,
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
	err = service.Services.ChannelPlanService.Create(plan)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)
}

func (p *channelPlanApi) Load(context *gin.Context) {

}

func (p *channelPlanApi) Update(context *gin.Context) {

}

func (p *channelPlanApi) ChangeStatus(context *gin.Context) {

}

func (p *channelPlanApi) Delete(context *gin.Context) {

}

func (p *channelPlanApi) RefreshKey(context *gin.Context) {

}
