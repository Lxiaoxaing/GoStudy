package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

/**
封装服务器返回码
*/
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//错误状态处理
func ResponseError(c *gin.Context, code int, err error, msg string) {
	resp := &Response{Code: code, Msg: msg, Data: nil}
	c.JSON(code, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(code, err)
}

//正确状态处理
func ResponseSuccess(c *gin.Context, msg string, data interface{}) {
	resp := &Response{Code: 200, Msg: msg, Data: data}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
