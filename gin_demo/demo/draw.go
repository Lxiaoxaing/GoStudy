package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
页面渲染
*/

//html渲染
func HtmlDrawPost(c *gin.Context) {
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	c.HTML(http.StatusOK, "posts/index.html", gin.H{
		"title": "posts/index",
	})
}

//html渲染
func HtmlDrawGet(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"title": "users/index",
	})
}

//自定义模板
func CustomTemp(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", "<a href='https://liwenzhou.com'>李文周的博客</a>")
}

//静态文件处理
//func staticFile(r *gin.Engine) {
//	r.Static("/static", "./static")
//	HtmlDraw(r)
//}

//JSON渲染
func JsonDraw1(c *gin.Context) {
	//方式一：自己拼接JSON
	c.JSON(http.StatusOK, gin.H{"message": "Hello I’m liuxiaoxiang!"})
}
func JsonDraw2(c *gin.Context) {
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
}


//XML渲染
func XmlDraw1(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "Hello I’m liuxiaoxiang!"})
}
func XmlDraw2(c *gin.Context) {
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
}
