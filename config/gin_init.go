package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message-push/bootstrap"
	"message-push/common"
	"net/http"
	"runtime/debug"
	"strings"
)

var G = gin.Default()

func GinInit() {

	G.LoadHTMLGlob("view/templates/*")
	G.Static("/static", "view/static")
	G.Use(Recover)
	G.Use(authCheck())

	RouterInit(G)
	ApiInit(G)

	fmt.Println(bootstrap.Option.Web.Port)
	err := G.Run(fmt.Sprintf("%s:%s", bootstrap.Option.Web.Addr, bootstrap.Option.Web.Port))
	if err != nil {
		fmt.Println("gin init failed")
	}
}

func authCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		uris := strings.Split(uri, "/")
		if uris[0] == "admin" {

			u, err := c.Cookie("u")
			if err != nil {
				noAuthRedirect(c)
			}

			v, exist := bootstrap.LocalCache.Get(u)
			if !exist {
				noAuthRedirect(c)
			}

			refreshAuth(c, u, v.(string))

		}
	}

}

func refreshAuth(c *gin.Context, u string, v string) {
	c.SetCookie(
		"u",
		u,
		86400,
		"*",
		"*",
		false,
		true,
	)
	bootstrap.LocalCache.Set(fmt.Sprintf("auth-%s", u), v, 86400)

}

func noAuthRedirect(c *gin.Context) {
	method := c.Request.Method
	if strings.ToLower(method) == "get" {
		c.Redirect(http.StatusOK, "/login")
	} else {
		common.R.Custom(c, common.Result{
			Code: 401,
		})
	}
}

// Recover gin全局err的处理
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			common.R.FailWithMsg(c, errorToString(r))
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
