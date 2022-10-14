package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type indexRouter struct {
}

func (r *indexRouter) ToIndex(c *gin.Context) {

	if true {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "ms主页",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "ms主页",
		})
	}

}

func (r *indexRouter) ToLogin(c *gin.Context) {

}
