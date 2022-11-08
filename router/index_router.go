package router

import (
	"github.com/gin-gonic/gin"
	"dobby/common"
	"net/http"
)

type indexRouter struct {
}

func (r *indexRouter) ToIndex(c *gin.Context) {

	ifLogin := common.AuthUtil.Check(c)
	if ifLogin {
		c.Redirect(302, "/admin")
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "消息推送平台登录",
		})
	}

}
