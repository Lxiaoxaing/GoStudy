package config

import "database/sql"

//定义一个全局对象db
var Db *sql.DB

//定义一个初始化数据库的函数
func InitDB() (err error) {
	//DSN:data source name
	dsn := "root:bhjRjxwC8EBqaJC7@tcp(127.0.0.1:8400)/my_test?charset=utf8mb4&parseTime=true"
	//不会校验账号密码是否正确
	//注意：这里不要使用:=，我们是给全局变量赋值，然后main函数中使用全局变量db
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接（校验dsn是否正确）
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}
