package config

import (
	"github.com/gin-gonic/gin"
	"message-push/router"
)

func RouterInit(g *gin.Engine) {

	g.GET("/", router.ToIndex)

	g.GET("/login", router.ToLogin)

}
