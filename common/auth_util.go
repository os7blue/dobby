package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var AuthUtil = new(authUtil)

type authUtil struct {
}

func (a *authUtil) Check(c *gin.Context) bool {

	u, err := c.Cookie("u")
	if err != nil {
		return false
	}

	_, exist := LocalCache.Get(fmt.Sprintf("auth-%s", u))

	return exist

}

func (a *authUtil) GetTokenValue(c *gin.Context) (interface{}, bool) {

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
func (a *authUtil) SetToken(c *gin.Context, code string) {

	c.SetCookie(
		"u",
		code,
		7200,
		"",
		"",
		false,
		true,
	)

}
