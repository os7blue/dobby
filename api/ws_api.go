package api

import (
	"dobby/common"
	"dobby/service"
	"github.com/gin-gonic/gin"
)

type wsApi struct {
}

func (a wsApi) conn(c *gin.Context) {

}

func (a wsApi) Conn(c *gin.Context) {

	m := map[string]string{}
	err := c.BindUri(&m)
	if err != nil {
		return
	}
	if m["key"] == "" {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	if m["token"] == "" {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	service.Services.WsService.Conn(m["key"], m["token"], c)

}
