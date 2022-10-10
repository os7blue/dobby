package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var g = gin.Default()

func GinInit() {

	g.LoadHTMLGlob("view/templates/*")
	g.Static("/static", "view/static")
	RouterInit(g)
	ApiInit(g)

	err := g.Run("0.0.0.0:8081")
	if err != nil {
		fmt.Println("gin init failed")
	}
}
