package demo

import (
	"fmt"
	"gin_demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取querystring参数
//querystring指的是URL中?后面携带的参数
func GetQueryString(c *gin.Context) {
	username := c.DefaultQuery("username", "小王子")
	//username:=c.Query("username")
	address := c.Query("address")
	//输出json结果给调用方
	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": username,
		"address":  address,
	})
}

//获取form参数    请求的数据通过form表单来提交
func GetFormString(c *gin.Context) {
	//DefaultPostForm 取不到值时会返回制定的默认值
	//username := c.DefaultPostForm("username", "小王子")
	username := c.PostForm("username")
	address := c.PostForm("address")
	//输出json结果给调用方
	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": username,
		"address":  address,
	})
}

//获取path参数
func GetPathString(c *gin.Context) {
	username := c.Param("username")
	address := c.Param("address")
	//输出json结果给调用方
	c.JSON(http.StatusOK, gin.H{
		"message":  "ok",
		"username": username,
		"address":  address,
	})
}

//参数绑定
//为了更方便获取请求参数，可以基于请求的Content-Type 识别强求数据类型并利用反射机制自动提取请求中 QueryString、form表单、JSON、XML等参数到结构体中。
//绑定JSON的示例（{"user":"lxx","password":"123456"}）
func LoginJson(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBind(&login); err == nil {
		fmt.Printf("login info:%#v\n", login)
		c.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

/**
绑定form表单示例（user=lxx&password=123456）
*/
func LoginForm(c *gin.Context) {
	var login models.Login
	//ShouldBind()会根据请求的Context-Type自行选择绑定器
	if err := c.ShouldBind(&login); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
	}
}

//绑定QueryString示例 （/loginQuery?username=lxx&password=123456）
func LoginQuery(c *gin.Context) {
	var login models.Login
	//ShouldBind()会根据请求的Content-Type自行选择绑定器
	if err := c.ShouldBind(&login); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"username": login.User,
			"password": login.Password,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
	}
}
