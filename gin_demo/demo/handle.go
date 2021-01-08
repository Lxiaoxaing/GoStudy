package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle(r *gin.Engine) {
	redirect(r)
	route(r)
}

//重定向
func redirect(r *gin.Engine) {
	//HTTP重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})

	//路由重定向
	r.GET("/test2", func(c *gin.Context) {
		//指定重定向的URL
		c.Request.URL.Path = "/test"
		r.HandleContext(c)
	})
}

//路由
func route(r *gin.Engine) {
	//普通路由
	//r.GET("/index", func(c *gin.Context) {})
	//r.POST("/index", func(c *gin.Context) {})
	////匹配所有请求的Any方法：
	//r.Any("/index", func(c *gin.Context) {})
	//为没有配置处理函数的路由添加处理程序
	//r.NoRoute(func(c *gin.Context) {
	//	path, _ := os.Getwd()
	//	c.HTML(http.StatusNotFound, path+"\\templates\\posts\\index.html", nil)
	//})
}
