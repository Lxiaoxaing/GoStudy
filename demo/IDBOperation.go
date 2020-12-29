package demo

import (
	"../config"
	"fmt"
)

//单行查询
//db.QueryRow() 执行一次查询，并期望返回最多一行结果。总是返回非nil的值，直到返回值的Sacn方法被调用时，才会返回被延迟的错误。
func QueryRowDemo() {
	fmt.Println("========QueryRowDemo========")
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

//多行查询
//db.Query()执行一次查询，返回多行结果，一般用于执行select命令。参数ages表示query中的占位参数
func QueryMultiRowDemo() {
	fmt.Println("========QueryMultiRowDemo========")
	sqlStr := "select id,name,age from user where id > 1"
	rows, err := config.Db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	//非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	//循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%V\n", err)
			return
		}
		fmt.Printf("id:%d,name:%s,age:%d", u.id, u.name, u.age)
	}

}

//插入数据   Exec 插入、更新和删除数据
func InsertRowDemo() {
	sqlStr := "insert into user(name,age) values (?,?)"
	ret, err := config.Db.Exec(sqlStr, "小苏", 28)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %d\n", theID)

}

//更新数据
func UpdateRowDemo() {
	sqlStr := "update user set age=? where id=?"
	ret, err := config.Db.Exec(sqlStr, 20, 2)
	if err != nil {
		fmt.Printf("update failed,err%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get rowaffeted failed,err:%v\n", err)
		return
	}
	fmt.Printf("update success,affected row:%d\n", n)

}

//删除数据
func DeleteRowDemo() {
	sqlStr := "delete from user where id=?"
	ret, err := config.Db.Exec(sqlStr, 5)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //影响的行数
	if err != nil {
		fmt.Printf("get rowsaffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete success,affected rows;%d\n", n)
}

//用户实体
type user struct {
	id   int
	age  int
	name string
}
