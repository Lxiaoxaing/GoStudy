package dbtools

// 数据库驱动
const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
	Utf string
	Charset string
}

//数据库连接配置
var MasterDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   "8400",
	User:   "root",
	Pwd:    "bhjRjxwC8EBqaJC7",
	DbName: "my_test",
}
