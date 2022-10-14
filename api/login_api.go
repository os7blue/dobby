package api

import (
	"github.com/gin-gonic/gin"
	R "message-push/model/r"
	"net/http"
)

type loginApi struct {
}

func (a *loginApi) SendCode(c *gin.Context) {

	c.AsciiJSON(http.StatusOK,
		R.Success(),
	)

}
