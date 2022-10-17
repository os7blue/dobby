package api

import (
	"github.com/gin-gonic/gin"
	"message-push/common"
	"message-push/service"
)

type authApi struct {
}

func (a *authApi) Login(c *gin.Context) {

}

func (a *authApi) SendCode(c *gin.Context) {

	param := map[string]string{}
	err := c.Bind(&param)
	if err != nil {
		common.R.FailWithMsg(c, "参数有误")
		return
	}

	email := param["email"]
	if email == "" {
		common.R.FailWithMsg(c, "email地址格式不正确")
	}

	err = service.Services.AuthService.SendCode(email)
	if err != nil {
		common.R.Fail(c)
	}
	common.R.Success(c)

}
