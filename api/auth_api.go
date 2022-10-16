package api

import (
	"github.com/gin-gonic/gin"
	"message-push/common"
)

type authApi struct {
}

func (a *authApi) SendCode(c *gin.Context) {

	param := map[string]string{}
	c.Bind(&param)

	common.R.Success(c)

}
