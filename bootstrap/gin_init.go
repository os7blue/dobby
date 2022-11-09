package bootstrap

import (
	"dobby/common"
	"dobby/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

var G = gin.Default()

func ginInit() {

	G.Delims("{[{", "}]}")
	G.LoadHTMLGlob("view/templates/*")
	G.Static("/static", "view/static")

	G.Use(Recover)
	G.Use(authCheck())

	routerInit(G)
	apiInit(G)

	err := G.Run(fmt.Sprintf("%s:%s", common.Option.Web.Addr, common.Option.Web.Port))
	if err != nil {
		fmt.Println("gin init failed")
	}
}

func authCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		//test login token
		//setTestToken(c)
		uri := c.Request.RequestURI
		uris := strings.Split(uri, "/")
		if uris[1] == "admin" {

			RefreshToken(c)

		}
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
			c.JSON(http.StatusOK, &model.Result{
				Code: 0,
				Msg:  errorToString(r),
			})
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

func noAuthRedirect(c *gin.Context) {
	method := c.Request.Method
	c.SetCookie(
		"u",
		"",
		-1,
		"",
		"",
		false,
		true,
	)
	if strings.ToLower(method) == "get" {
		c.Redirect(302, "/")
		c.Abort()
	} else {
		common.AuthUtil.DelToken(c)
		c.JSON(http.StatusOK, model.Result{
			Code: 401,
		})
		c.Abort()
	}
}

func setTestToken(c *gin.Context) {

	c.SetCookie(
		"u",
		"test",
		7200,
		"",
		"",
		false,
		true,
	)
	common.LocalCache.Set(fmt.Sprintf("auth-%s", "test"), "123", 7200)

}

// RefreshToken refresh token of every request
func RefreshToken(c *gin.Context) {

	u, err := c.Cookie("u")
	if err != nil {
		noAuthRedirect(c)
		return
	}

	v, exist := common.LocalCache.Get(fmt.Sprintf("auth-%s", u))
	if !exist {
		noAuthRedirect(c)
		return
	}
	c.SetCookie(
		"u",
		u,
		7200,
		"",
		"",
		false,
		true,
	)
	common.LocalCache.Set(fmt.Sprintf("auth-%s", u), v, 7200)

}
