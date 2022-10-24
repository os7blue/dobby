package common

import (
	"github.com/gin-gonic/gin"
	"message-push/model"
	"net/http"
)

var R = new(r)

type r struct {
}

func (r *r) Custom(c *gin.Context, result model.Result) {

	c.JSON(http.StatusOK, &result)

}

func (r *r) Success(c *gin.Context) {
	c.JSON(http.StatusOK, &model.Result{
		Code: 1,
		Msg:  "success",
	})
}

func (r *r) SuccessWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &model.Result{
		Code: 1,
		Msg:  "success",
		Data: data,
	})

}

func (r *r) SuccessWithDataCount(c *gin.Context, data any, count int64) {
	c.JSON(http.StatusOK, &model.Result{
		Code:  1,
		Msg:   "success",
		Data:  data,
		Count: count,
	})

}

func (r *r) Fail(c *gin.Context) {
	c.JSON(http.StatusOK, &model.Result{
		Code: 0,
		Msg:  "failed",
	})
}

func (r *r) FailWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &model.Result{
		Code: 0,
		Msg:  msg,
	})
}
