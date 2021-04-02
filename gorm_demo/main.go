package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gorm_demo/models"
	setting "gorm_demo/pkg/settings"
	"gorm_demo/routers"
	"net/http"
)
var db *gorm.DB
// 初次进入初始化
func init() {
	setting.Setup("")
	// models 数据库初始化
	models.Setup()
}


// 执行
func main() {
	// gin初始化
	gin.SetMode(setting.ServerSetting.RunMode)
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
