package api

import (
	"dobby/common"
	"dobby/model"
	"dobby/service"
	"github.com/gin-gonic/gin"
)

type wsApi struct {
}

func (a wsApi) Conn(c *gin.Context) {

	m := model.WsConnView{}
	err := c.ShouldBindUri(&m)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	_, err = service.WsService.Conn(m.ID, m.Key, c)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

}
