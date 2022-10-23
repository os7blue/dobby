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

	var createVm model.PushChannelInfoCreateValidator

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
