package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//html渲染
	htmlDraw(r)
	//customTemp(r)

	//启动HTTP服务，默认在0.0.0.0:80880启动服务
	r.Run(":8080")
}

//

//自定义模板
func customTemp(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl","<a href='https://liwenzhou.com'>李文周的博客</a>")
	})
}

//html渲染
func htmlDraw(r *gin.Engine) {
	//Gin框架中使用 LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

}
