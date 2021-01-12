package routers

import (
	"gin_demo/demo"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"html/template"
	"log"
	"net/http"
)

func SetupRouter() *gin.Engine {
	//Default返回一个默认的路由引擎
	r := gin.Default()

	//使用全局中间件
	r.Use(demo.StatCost())

	//请求方式
	LoadHello(r)

	//html渲染
	LoadDraw(r)

	//获取参数
	LoadQueryString(r)

	//重定向
	LoadHandle(r)

	//文件上传地址
	LoadFile(r)

	LoadMoreServer()

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	//启动HTTP服务，默认在0.0.0.0:80880启动服务
	r.Run(":8888")

	return r
}


/**
加载不同请求方式
*/
func LoadHello(r *gin.Engine) {
	r.GET("/hello", demo.HelloHandler)
	r.GET("/book", demo.GetHandler)
	r.POST("/book", demo.PostHandler)
	r.PUT("/book", demo.PutHandler)
	r.DELETE("/book", demo.DeleteHandler)
}

/**
加载不同渲染请求
*/
func LoadDraw(r *gin.Engine) {
	//Gin框架中使用 LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染
	r.LoadHTMLGlob("templates/**/*")
	r.POST("/posts/index", demo.HtmlDrawPost)
	r.GET("users/index", demo.HtmlDrawGet)

	//模板处理
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")
	r.GET("/index", demo.CustomTemp)

	//静态文件处理
	r.Static("/static", "./static")
	//按照html渲染

	//JSON渲染
	r.GET("/someJSON", demo.JsonDraw1)
	r.GET("/moreJSON", demo.JsonDraw2)

	//XML渲染
	r.GET("/someXML", demo.XmlDraw1)
	r.GET("/moreXML", demo.XmlDraw2)
}

/**
加载参数解析方式
*/
func LoadQueryString(r *gin.Engine) {
	//get请求接受参数
	r.GET("/user/search", demo.GetQueryString)
	//表单参数
	r.POST("/user/search2", demo.GetFormString)
	//请求地址上带参数
	r.GET("/user/search3/:username/:address", demo.GetPathString)

	//参数绑定，使用实体接受参数
	//json参数
	r.POST("loginJson", demo.LoginJson)
	//表单参数
	r.POST("loginForm", demo.LoginForm)
	//get参数
	r.GET("loginQuery", demo.LoginQuery)
}

/**
重定向
*/
func LoadHandle(r *gin.Engine) {
	//http重定向
	r.GET("/test", demo.HttpRedirect)
	//路由重定向
	//r.GET("/test2",demo.RouterRedirect)

	//为没有配置处理函数的路由添加处理程序
	r.NoRoute(demo.HttpRedirect)
}

/**
文件上传请求地址
*/
func LoadFile(r *gin.Engine) {
	//使用路由组 支持嵌套
	file := r.Group("/file")
	{
		//上传单个文件
		file.POST("/upload", demo.UploadFile)
		//上传多个文件
		file.POST("uploadFiles", demo.UploadFiles)
	}

}


func LoadMoreServer() {
	//运行多个服务
	//server01 := &http.Server{
	//	Addr:         "8881",
	//	Handler:      router01(),
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}
	//server02 := &http.Server{
	//	Addr:         "8882",
	//	Handler:      router02(),
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}

	//借助errgroup.Group或者自行开启两个goroutine分别启动两个服务
	//g.Go(func() error {
	//	return server01.ListenAndServe()
	//})
	//g.Go(func() error {
	//	return server02.ListenAndServe()
	//})

}

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	//路由组 支持嵌套
	file := e.Group("/file")
	{
		//上传单个文件
		file.POST("/upload", func(c *gin.Context) { demo.UploadFile(c) })
		//上传多个文件
		file.POST("uploadFiles", func(c *gin.Context) { demo.UploadFiles(c) })
	}
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	//路由组 支持嵌套
	//参数绑定
	//demo.ParamBind(e)
	return e
}
