package routers

import (
	"github.com/gin-gonic/gin"
	"gorm_demo/routers/plain"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/xiang")
	v1.GET("/create", plain.Create)
	v1.GET("/dboperation", plain.DbOperation)
	return r
}
