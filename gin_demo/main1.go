package main

import (
	"github.com/gin-gonic/gin"
	"../middleware"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	//GET：请求方式；/hello:请求路径
	//当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON:返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/book", func(c *gin.Context) {
		c.JSON(middleware.SuccessCode, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(middleware.SuccessCode, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(middleware.SuccessCode, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(middleware.SuccessCode, gin.H{
			"message": "DELETE",
		})
	})
	//启动HTTP服务，默认在0.0.0.0:80880启动服务
	r.Run()

}
