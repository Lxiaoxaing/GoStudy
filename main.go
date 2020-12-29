package main

import (
	"./config"
	"./demo"
	"fmt"
)

/**
主函数入口
*/
func main() {
	//exercise.Imath()
	////变量赋值
	//exercise.Variable()
	////常量赋值
	//exercise.Constant()
	////文件操作
	//exercise.FileOperation()
	////循环
	//exercise.Loop()
	////集合
	//exercise.Collection()

	//练习题
	//exercise.Exercise1()

	//连接数据库
	err := config.InitDB()
	if err != nil {
		fmt.Println("init db failed,err:%v\n", err)
		return
	}

	//单行查询
	demo.QueryRowDemo()
	//多行查询
	demo.QueryMultiRowDemo()
	//插入数据
	demo.InsertRowDemo()
	//更新数据
	demo.UpdateRowDemo()
	//删除数据
	demo.DeleteRowDemo()
}
