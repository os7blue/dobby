package config

import (
	"github.com/gin-gonic/gin"
	"message-push/router"
)

func RouterInit(g *gin.Engine) {

	g.GET("/", router.Routers.IndexRouter.ToIndex)

	g.GET("/login", router.Routers.IndexRouter.ToLogin)

}
