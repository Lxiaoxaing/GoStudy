package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	setting "gorm_demo/pkg/settings"
	"log"
)

type pageInfo struct {
	Page int `json:"page"` //现在页数
	PageSize int `json:"page_size"` //每页条数
	Count int `json:"count"` //总条数
	PageCount float64 `json:"page_count"` //总页数
}

var db *gorm.DB

// Setup initializes the database    instance
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		fmt.Println("err:", err)
		log.Fatalf("models.Setup err: %v", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	if setting.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
