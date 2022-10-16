package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var R = new(r)

type r struct {
}

type Result struct {
	Code  int    `json:"code" default:"0"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
	Count int    `json:"count"`
}

func (r *r) Success(c *gin.Context) {
	c.JSON(http.StatusOK, &Result{
		Code: 1,
		Msg:  "success",
	})
}

func (r *r) SuccessWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &Result{
		Code: 1,
		Msg:  "success",
		Data: data,
	})

}

func (r *r) SuccessWithDataCount(c *gin.Context, data any, count int) {
	c.JSON(http.StatusOK, &Result{
		Code:  1,
		Msg:   "success",
		Data:  data,
		Count: count,
	})

}

func (r *r) Fail(c *gin.Context) {
	c.JSON(http.StatusOK, &Result{
		Code: 0,
		Msg:  "failed",
	})
}

func (r *r) FailWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &Result{
		Code: 0,
		Msg:  msg,
	})
}
