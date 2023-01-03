package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	Success = "1"
	Error   = "0"
)

func Res(code string, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
func OK(c *gin.Context) {
	Res(Success, "操作成功", map[string]interface{}{}, c)
}
func OKWithData(data interface{}, c *gin.Context) {
	Res(Success, "操作成功", data, c)
}
func OKWithMsg(msg string, c *gin.Context) {
	Res(Success, msg, map[string]interface{}{}, c)
}
func OkWithDetail(msg string, data interface{}, c *gin.Context) {
	Res(Success, msg, data, c)
}
func Fail(c *gin.Context) {
	Res(Error, "操作失败", map[string]interface{}{}, c)
}
func FailWithData(data interface{}, c *gin.Context) {
	Res(Error, "操作失败", data, c)
}
func FailWithMsg(msg string, c *gin.Context) {
	Res(Error, msg, map[string]interface{}{}, c)
}
func FailWithDetail(msg string, data interface{}, c *gin.Context) {
	Res(Error, msg, data, c)
}
