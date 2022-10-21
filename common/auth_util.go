package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"message-push/bootstrap"
	"strings"
)

var AuthUtil = new(authUtil)

type authUtil struct {
}

func (a *authUtil) Check(c *gin.Context) bool {

	u, err := c.Cookie("u")
	if err != nil {
		return false
	}

	_, exist := bootstrap.LocalCache.Get(fmt.Sprintf("auth-%s", u))

	return exist

}

func (a *authUtil) GetTokenValue(c *gin.Context) (interface{}, bool) {

	u, err := c.Cookie("u")
	if err != nil {
		return nil, false
	}

	v, exist := bootstrap.LocalCache.Get(fmt.Sprintf("auth-%s", u))
	if !exist {
		return nil, false
	}

	return v, true

}

// SetToken config token info after login success
func (a *authUtil) SetToken(c *gin.Context, tokenData any) {

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
	bootstrap.LocalCache.Set(fmt.Sprintf("auth-%s", uid), tokenData, 7200)

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
		R.Custom(c, Result{
			Code: 401,
		})
	}
}

// RefreshToken refresh token of every request
func (a *authUtil) RefreshToken(c *gin.Context) {

	u, err := c.Cookie("u")
	if err != nil {
		noAuthRedirect(c)
		return
	}

	v, exist := bootstrap.LocalCache.Get(fmt.Sprintf("auth-%s", u))
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
	bootstrap.LocalCache.Set(fmt.Sprintf("auth-%s", u), v, 7200)

}
