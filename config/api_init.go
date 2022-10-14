package config

import (
	"github.com/gin-gonic/gin"
	"message-push/api"
)

func ApiInit(g *gin.Engine) {

	g.POST("/login/send_code", api.Apis.LoginApi.SendCode)

}
