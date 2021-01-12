package demo

import (
	"gin_demo/middleware"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	// c.JSON:返回JSON格式的数据
	//GET：请求方式；/hello:请求路径
	//当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
func GetHandler(c *gin.Context) {
	// c.JSON:返回JSON格式的数据
	c.JSON(middleware.SuccessCode, gin.H{
		"message": "GET",
	})
}

func PostHandler(c *gin.Context) {
	// c.JSON:返回JSON格式的数据
	c.JSON(middleware.SuccessCode, gin.H{
		"message": "POST",
	})
}

func PutHandler(c *gin.Context) {
	// c.JSON:返回JSON格式的数据
	c.JSON(middleware.SuccessCode, gin.H{
		"message": "PUT",
	})
}

func DeleteHandler(c *gin.Context) {
	// c.JSON:返回JSON格式的数据
	c.JSON(middleware.SuccessCode, gin.H{
		"message": "DELETE",
	})
}
