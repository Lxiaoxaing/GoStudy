package util

import (
	"github.com/gin-gonic/gin"
	"gorm_demo/pkg/messge"
	"net/http"
)

type JSONResult struct {
	Success bool        `json:"success" ` // 成功/失败
	Code    int         `json:"code" `    // 错误码：200-成功
	Message string      `json:"message"`  // 消息
	Data    interface{} `json:"data"`     // 详细信息
}

// api返回格式
func ApiReturn(code int, ret interface{}, c *gin.Context) {
	success := true
	if code != messge.SUCCESS {
		success = false
	}
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"code":    code,
		"message": messge.GetMessage(code),
		"data":    ret,
	})
}

// api返回格式,需传接口返回信息
func ApiReturnMsg(code int, ret interface{}, c *gin.Context, msg string) {
	success := true
	if code != messge.SUCCESS {
		success = false
	}
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"code":    code,
		"message": msg,
		"data":    ret,
	})
}
