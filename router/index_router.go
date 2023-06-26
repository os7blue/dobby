package router

import (
	"dobby/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type indexRouter struct {
}

func (r *indexRouter) ToIndex(c *gin.Context) {

	ifLogin := common.AuthUtil.Check(c)
	if ifLogin {
		c.Redirect(302, "/admin")
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}

}

func (r *indexRouter) ToVoiceChannel(c *gin.Context) {

	c.HTML(http.StatusOK, "voice_channel.html", nil)

}
