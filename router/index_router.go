package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToIndex(c *gin.Context) {

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

func ToLogin(c *gin.Context) {

}
