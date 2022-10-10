package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiInit(g *gin.Engine) {

	g.POST("/login/send_code", func(context *gin.Context) {
		context.AsciiJSON(http.StatusOK, map[string]string{
			"data": "1",
			"sord": "2",
		})
	})

}
