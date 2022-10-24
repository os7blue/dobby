package api

import (
	"github.com/gin-gonic/gin"
	"message-push/common"
	"message-push/model"
	"message-push/service"
)

type channelInfoApi struct {
}

func (i *channelInfoApi) Create(c *gin.Context) {

	var createVm model.ChannelInfoCreateValidator

	err := c.ShouldBindJSON(&createVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = service.Services.ChannelInfoService.CreateOne(createVm.Name)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.Success(c)

}

func (i *channelInfoApi) Load(c *gin.Context) {

	var loadVm model.PageBody[model.ChannelInfoSearchValidator]
	err := c.ShouldBindJSON(&loadVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	result, count, err := service.Services.ChannelInfoService.LoadPage(loadVm.Page, loadVm.Limit, loadVm.Param)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.SuccessWithDataCount(c, result, count)

}
