package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"message-push/common"
	"message-push/service"
)

type authApi struct {
}

func (a *authApi) Login(c *gin.Context) {
	param := map[string]string{}
	err := c.Bind(&param)
	if err != nil {
		common.R.FailWithMsg(c, "参数有误")
		return
	}

	bl := service.Services.AuthService.Login(param["email"], param["code"])

	if bl {
		uid := uuid.New()
		c.SetCookie(
			"u",
			uid.String(),
			86400,
			"*",
			"*",
			false,
			true,
		)
		common.R.Success(c)
	}

	common.R.Fail(c)

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
