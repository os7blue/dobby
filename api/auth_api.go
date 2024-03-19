package api

import (
	"dobby/common"
	"dobby/model"
	"dobby/service"
	"github.com/gin-gonic/gin"
)

type authApi struct {
}

func (a *authApi) LogOut(c *gin.Context) {
	common.AuthUtil.DelToken(c)
	common.R.Success(c)
}
func (a *authApi) Login(c *gin.Context) {
	var loginVm model.LoginValidator
	err := c.ShouldBindJSON(&loginVm)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	tokenCode, err := service.AuthService.Login(loginVm.Email, loginVm.Code)

	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	common.AuthUtil.SetToken(c, tokenCode)
	common.R.Success(c)
}

func (a *authApi) SendCode(c *gin.Context) {

	var codeVm model.SendCodeValidator
	err := c.ShouldBindJSON(&codeVm)

	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}

	err = service.AuthService.SendCode(codeVm.Email)
	if err != nil {
		common.R.FailWithMsg(c, err.Error())
		return
	}
	common.R.Success(c)

}
