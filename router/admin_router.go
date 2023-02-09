package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type adminRouter struct {
}

func (r *adminRouter) ToIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)

}

func (r *adminRouter) ToChannelInfo(c *gin.Context) {

	c.HTML(http.StatusOK, "channel_info.html", nil)

}

func (r *adminRouter) ToLog(c *gin.Context) {

	c.HTML(http.StatusOK, "log.html", nil)

}

func (r *adminRouter) ToChannelPlan(c *gin.Context) {
	c.HTML(http.StatusOK, "channel_plan.html", nil)

}
