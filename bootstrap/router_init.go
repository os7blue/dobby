package bootstrap

import (
	"dobby/router"
	"github.com/gin-gonic/gin"
)

func routerInit(g *gin.Engine) {

	g.GET("/", router.Routers.IndexRouter.ToIndex)
	g.GET("/admin", router.Routers.AdminRouter.ToIndex)
	g.GET("/admin/channel/info", router.Routers.AdminRouter.ToChannelInfo)
	g.GET("/admin/channel/plan", router.Routers.AdminRouter.ToChannelPlan)

	g.GET("/wsp/:id/:key", router.Routers.IndexRouter.ToWsP)

}
