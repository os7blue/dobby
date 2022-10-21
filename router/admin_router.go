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

func (r *adminRouter) ToChannel(c *gin.Context) {

	c.HTML(http.StatusOK, "channel.html", gin.H{
		"title": "channel",
	})

}

func (r *adminRouter) ToSetting(c *gin.Context) {

	c.HTML(http.StatusOK, "setting.html", gin.H{
		"title": "setting",
	})

}

func (r *adminRouter) ToLog(c *gin.Context) {

	c.HTML(http.StatusOK, "log.html", gin.H{
		"title": "log",
	})

}
