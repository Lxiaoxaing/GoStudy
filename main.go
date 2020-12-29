package main

import (
	"./config"
	"./service"
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

	service.QueryRowDemo()
}
