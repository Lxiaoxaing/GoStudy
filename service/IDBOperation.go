package service

import (
	"../config"
	"fmt"
)

//单行查询
func QueryRowDemo() {
	sqlStr := "select id,name,age from user where id=?"
	var u user
	//非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := config.Db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}






//用户实体
type user struct {
	id   int
	age  int
	name string
}
