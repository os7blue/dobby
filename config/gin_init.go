package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var G = gin.Default()

func GinInit() {

	G.LoadHTMLGlob("view/templates/*")
	G.Static("/static", "view/static")

	RouterInit(G)
	ApiInit(G)

	fmt.Println(Option.Web.Port)
	err := G.Run(fmt.Sprintf("%s:%s", Option.Web.Addr, Option.Web.Port))
	if err != nil {
		fmt.Println("gin init failed")
	}
}
