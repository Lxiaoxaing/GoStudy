package dbtools

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
)

/**
MysqlConnectiPool
数据库连接操作库
基于gorm封装开发
*/
var mysqlInstance *MysqlConnectPool
var mysqlOnce sync.Once //在需要的时候执行一次，且只执行一次
var db *gorm.DB
var err_db error

type MysqlConnectPool struct {
}

func GetMysqlInstance() *MysqlConnectPool {
	mysqlOnce.Do(func() {
		mysqlInstance = &MysqlConnectPool{}
	})
	return mysqlInstance
}

/**
初始化数据库连接
*/
func (m *MysqlConnectPool) InitMysqlPool() (issucc bool) {
	dbConf := MasterDbConfig
	db, err_db = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Pwd,
		dbConf.Host,
		dbConf.Port,
		dbConf.DbName))
	db.SingularTable(true)
	fmt.Println(err_db)
	if err_db != nil {
		log.Fatal(err_db)
		return false
	}
	//关闭数据库   db会被多个goroutine共享，可以不调用
	//db.Close()
	return true
}

/**
对外获取数据库连接对象db
*/
func (m *MysqlConnectPool) GetMysqlPool() *gorm.DB {
	return db
}

func GetMysqlDb() (db *gorm.DB) {
	return GetMysqlInstance().GetMysqlPool()
}
