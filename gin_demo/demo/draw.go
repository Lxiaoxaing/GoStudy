package demo

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

/**
页面渲染
*/

//html渲染
func HtmlDraw(r *gin.Engine) {
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

//自定义模板
func CustomTemp(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
	})
}



//静态文件处理
func staticFile(r *gin.Engine) {
	r.Static("/static", "./static")
	HtmlDraw(r)
}



//JSON渲染
func JsonDraw(r *gin.Engine) {
	//gin.H是map[string] interface{}的缩写
	r.GET("/someJSON", func(c *gin.Context) {
		//方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello I’m liuxiaoxiang!"})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		//方法二:使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg.Name = "小向"
		msg.Message = "Hello I’m liuxiaoxiang!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})
}


//XML渲染
func XmlDraw(r *gin.Engine) {
	//gin.H是map[string] interface{}的缩写
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "Hello I’m liuxiaoxiang!"})
	})
	r.GET("moreXML", func(c *gin.Context) {
		//方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Age = 18
		msg.Message = "Hello I’m liuxiaoxiang!"
		msg.Name = "小向"
	})
}
