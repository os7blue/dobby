package api

import (
	"dobby/common"
	"dobby/model"
	"dobby/service"
	"github.com/gin-gonic/gin"
	"regexp"
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

	match, err := regexp.Match(common.UUID, []byte(pushInfoVm.Key))
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	if !match {
		common.R.FailWithMsg(c, "key格式不正确")
		return
	}

	ip := c.ClientIP()

	msg, err := service.SendService.Send(pushInfoVm.Key, pushInfoVm.Title, pushInfoVm.Content, ip)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.R.SuccessWithMsg(c, msg)

}
