package main

import "github.com/gin-gonic/gin"

//文件上传
func FileUpload(r *gin.Engine) {
	//处理multipart forms提交文件时默认的内存限制是32M
	//可以通过下面的方式修改
	r.POST("/upload")


}