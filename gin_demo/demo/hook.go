package demo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

/**
中间件注意事项
gin默认中间件
gin.Default()默认使用了Logger和Recovery中间件，其中：

Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

gin中间件中使用goroutine
当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。
*/

//定义中间件
//StatCost 是一个统计耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") //通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		//调用该请求的剩余处理程序
		c.Next()
		//不调用该请求的剩余处理程序
		c.Abort()
		cost := time.Since(start)
		log.Println(cost)
	}
}

//注册中间件
func use(r *gin.Engine) {
	//全局注册中间件
	r.Use(StatCost())

	//单个路由注册
	r.GET("/test", StatCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string) //从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	//为路由组注册中间件
	//写法1
	shop := r.Group("/shop", StatCost())
	{
		shop.GET("index", func(c *gin.Context) {})
	}
	//写法2
	shop2 := r.Group("/shop2")
	shop2.Use(StatCost())
	{
		shop2.Group("index2", func(c *gin.Context) {})
	}

}
