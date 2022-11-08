package api

import (
	"github.com/gin-gonic/gin"
	"dobby/common"
	"dobby/model"
	"dobby/service"
)

type pushApi struct {
}

func (p *pushApi) Push(c *gin.Context) {

	var pushInfoVm model.PushInfo
	err := c.ShouldBindJSON(&pushInfoVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = common.GlobalUtil.CheckArrStr(
		1,
		0,
		common.UUID,
		pushInfoVm.KeyStr,
	)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	ip := c.ClientIP()
	pushInfoVm.Ip = ip
	msg, err := service.Services.SendService.Send(pushInfoVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.SuccessWithMsg(c, msg)

}
