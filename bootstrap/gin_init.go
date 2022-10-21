package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"message-push/model"
	"net/http"
	"runtime/debug"
	"strings"
)

var G = gin.Default()

func ginInit() {

	G.LoadHTMLGlob("view/templates/*")
	G.Static("/static", "view/static")
	G.Use(Recover)
	G.Use(authCheck())

	routerInit(G)
	apiInit(G)

	fmt.Println(Option.Web.Port)
	err := G.Run(fmt.Sprintf("%s:%s", Option.Web.Addr, Option.Web.Port))
	if err != nil {
		fmt.Println("gin init failed")
	}
}

func authCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
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

func Check(c *gin.Context) bool {

	u, err := c.Cookie("u")
	if err != nil {
		return false
	}

	_, exist := LocalCache.Get(fmt.Sprintf("auth-%s", u))

	return exist

}

func GetTokenValue(c *gin.Context) (interface{}, bool) {

	u, err := c.Cookie("u")
	if err != nil {
		return nil, false
	}

	v, exist := LocalCache.Get(fmt.Sprintf("auth-%s", u))
	if !exist {
		return nil, false
	}

	return v, true

}

// SetToken config token info after login success
func SetToken(c *gin.Context, tokenData any) {

	// set cookie
	uid := uuid.New()
	c.SetCookie(
		"u",
		uid.String(),
		7200,
		"",
		"",
		false,
		true,
	)
	// save token info to local cache
	LocalCache.Set(fmt.Sprintf("auth-%s", uid), tokenData, 7200)

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
	} else {
		c.JSON(http.StatusOK, model.Result{
			Code: 401,
		})
	}
}

// RefreshToken refresh token of every request
func RefreshToken(c *gin.Context) {

	u, err := c.Cookie("u")
	if err != nil {
		noAuthRedirect(c)
		return
	}

	v, exist := LocalCache.Get(fmt.Sprintf("auth-%s", u))
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
	LocalCache.Set(fmt.Sprintf("auth-%s", u), v, 7200)

}
