package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type adminRouter struct {
}

func (r *adminRouter) ToIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "管理页面",
	})

}

func (r *adminRouter) ToChannelInfo(c *gin.Context) {

	c.HTML(http.StatusOK, "channel_info.html", gin.H{
		"title": "通道管理",
	})

}

func (r *adminRouter) ToLog(c *gin.Context) {

	c.HTML(http.StatusOK, "log.html", gin.H{
		"title": "log",
	})

}

func (r *adminRouter) ToChannelPlan(c *gin.Context) {
	c.HTML(http.StatusOK, "channel_plan.html", gin.H{
		"title": "通道方案",
	})

}
