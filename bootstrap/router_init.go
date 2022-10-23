package bootstrap

import (
	"github.com/gin-gonic/gin"
	"message-push/router"
)

func routerInit(g *gin.Engine) {

	g.GET("/", router.Routers.IndexRouter.ToIndex)
	g.GET("/admin", router.Routers.AdminRouter.ToIndex)
	g.GET("/admin/channel", router.Routers.AdminRouter.ToChannel)
	g.GET("/admin/log", router.Routers.AdminRouter.ToLog)
	g.GET("/admin/setting", router.Routers.AdminRouter.ToSetting)

}
