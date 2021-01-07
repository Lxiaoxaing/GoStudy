package main

import (
	"./middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	//html渲染
	htmlDraw(r)

	//模板处理
	customTemp(r)

	//静态文件处理
	//staticFile(r)

	//JSON渲染
	jsonDraw(r)

	//XML渲染
	xmlDraw(r)

	//获取参数
	getQueryString(r)

	//参数绑定
	paramBind(r)

	//文件上传
	fileUpload(r)

	FileUpload(r)

	//启动HTTP服务，默认在0.0.0.0:80880启动服务
	r.Run(":8080")
}

//文件上传
func fileUpload(r *gin.Engine) {
	//处理multipart forms提交文件时默认的内存限制是32M
	//可以通过下面的方式修改
	r.POST("/upload")


}

//参数绑定
//为了更方便获取请求参数，可以基于请求的Content-Type 识别强求数据类型并利用反射机制自动提取请求中 QueryString、form表单、JSON、XML等参数到结构体中。
func paramBind(r *gin.Engine) {

	//绑定JSON的示例（{"user":"lxx","password":"123456"}）
	r.POST("loginJson", func(c *gin.Context) {
		var login Login
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
	})

	//绑定form表单示例（user=lxx&password=123456）
	r.POST("loginForm", func(c *gin.Context) {
		var login Login
		//ShouldBind()会根据请求的Context-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
		}
	})

	//绑定QueryString示例 （/loginQuery?username=lxx&password=123456）
	r.GET("/loginQuery", func(c *gin.Context) {
		var login Login
		//ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
		}
	})
}

//获取querystring参数
//querystring指的是URL中?后面携带的参数
func getQueryString(r *gin.Engine) {
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username:=c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	//获取form参数
	//请求的数据通过form表单来提交
	r.POST("/user/search2", func(c *gin.Context) {
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

	})

	//获取path参数
	r.GET("/user/search3/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
}

//XML渲染
func xmlDraw(r *gin.Engine) {
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

//JSON渲染
func jsonDraw(r *gin.Engine) {
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

//静态文件处理
func staticFile(r *gin.Engine) {
	r.Static("/static", "./static")
	htmlDraw(r)
}

//自定义模板
func customTemp(r *gin.Engine) {
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

func hello(r *gin.Engine) {
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
}

//Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
