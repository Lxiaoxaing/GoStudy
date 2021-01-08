package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func UploadFile(c *gin.Context) {
	//单个文件上传
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	path, _ := os.Getwd();
	dst := fmt.Sprintf(path, "/%s", file.Filename)

	//上传文件到指定的目录
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s upload", file.Filename),
	})
}

//文件上传
func UploadFiles(c *gin.Context) {
	//处理multipart form 提交文件时默认的内存时32M
	//可以通过下面的方式修改
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		log.Printf(file.Filename)
		path, _ := os.Getwd(); //获取当前环境
		dst := fmt.Sprintf(path+"/%s", file.Filename)
		//上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploaded!", len(files)),
	})
}
