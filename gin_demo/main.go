package main

import (
	"gin_demo/demo"
	"github.com/gin-gonic/gin"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()

	//路由组 支持嵌套
	file := r.Group("/file")
	{
		//上传单个文件
		file.POST("/upload", func(c *gin.Context) { demo.UploadFile(c) })
		//上传多个文件
		file.POST("uploadFiles", func(c *gin.Context) { demo.UploadFiles(c) })
	}

	//请求方式
	demo.Hello(r)

	//html渲染
	demo.HtmlDraw(r)

	//模板处理
	demo.CustomTemp(r)

	//静态文件处理
	//staticFile(r)

	//JSON渲染
	demo.JsonDraw(r)

	//XML渲染
	demo.XmlDraw(r)

	//获取参数
	demo.GetQueryString(r)

	//参数绑定
	demo.ParamBind(r)

	//重定向
	demo.Handle(r)

	//启动HTTP服务，默认在0.0.0.0:80880启动服务
	r.Run(":8888")
}
