package models

import (
	"fmt"
	"time"
)

/**
1、查询
*/
func Select_Simple() {
	var user User
	//1、查询第一条
	db.First(&user)
	fmt.Println("First=======", user)

	//2、随机一条
	user.Id = 0
	db.Take(&user)
	fmt.Println("Take=======", user)

	//3、最后一条，主键排序
	user.Id = 0
	db.Last(&user)
	fmt.Println("Last=========", user)

	//4、查询所有
	var users []User
	db.Find(&users)
	fmt.Println("Find=========", users)

	//5、查指定的模拟一条记录（主键必须为整型）
	user.Id = 0
	db.First(&user, 7)
	fmt.Println("First=========", user)
}

/**
2、带条件查询
*/
func Select_where() {
	var user User
	db.Where("name=?", "xiaosu3").First(&user)
	fmt.Println("Where*****First=======", user)

	var users []User
	db.Where("name=?", "xiaosu3").Find(&users)
	fmt.Println("Where*****Find=======", users)

	//<>
	db.Where("name <>?", "xiaosu3").Find(&users)
	fmt.Println("Where***<>**Find=======", users)

	//IN
	db.Where("name IN (?)", []string{"xiaosu2", "xiaosu3"}).Find(&users)
	fmt.Println("Where***<>**Find=======", users)
	//LIKE
	db.Where("name LIKE ?", "%su%").Find(&users)
	fmt.Println("Where***LIKE**Find=======", users)
	//AND
	db.Where("name = ? AND age=?", "zhangchuangjian", 100).Find(&users)
	fmt.Println("Where***AND**Find=======", users)
	//Time
	db.Where("updated > ?", "2021-04-06 14:24:48").Find(&users)
	fmt.Println("Where***Time**Find=======", users)
	//BETWEEN
	db.Where("updated BETWEEN ? AND ?", "2021-04-06 14:24:48", time.Now()).Find(&users)
	fmt.Println("Where***BETWEEN**Find=======", users)
}

/**
3、Struct & Map
结构体进行查询时，只根据非零字段查询，可以使用指针或Scanner/Valuer接口避免这个问题
*/
func Select_Struct2Map() {
	//Struct
	var user User
	var users []User
	db.Where(&User{Name: "zhangchaungjian", Age: 100}).First(&user)

	//Map
	db.Where(map[string]interface{}{"name": "zhangchuangjian", "age": 100}).Find(&users)

	//主键切片
	db.Where([]int64{33, 34, 35}).Find(&users)

	/*
		// 使用指针
		type User struct {
			gorm.Model
			Name string
			Age  *int
		}

		// 使用 Scanner/Valuer
		type User struct {
			gorm.Model
			Name string
			Age  sql.NullInt64  // sql.NullInt64 实现了 Scanner/Valuer 接口
		}*/
}

/**
4、Not条件
*/
func Select_Not() {
	var users User

	//不在主切片中
	db.Not([]int64{1, 2, 3}).First(&users)
	fmt.Println()
	//其他语法不变
}

/**
5、Or条件
*/
func Select_Or() {
	var users []User
	db.Where("name=?", "zhangchuangjian").Or("name=?", "xiaoxiang").Find(&users)
	fmt.Println("or===============", users)
	db.Where("name=?", "zhangchuangjian").Or(User{Name: "xiaoxiang"}).Find(&users)
	fmt.Println("or======Struct=========", users)
	db.Where("name=?", "zhangchuangjian").Or(map[string]interface{}{"name": "xiaoxiang"}).Find(&users)
	fmt.Println("or======Map=========", users)
}

/**
6、Inline Condition 内联条件
作用与 Where 类似
当内联条件与 多个立即执行方法 一起使用时, 内联条件不会传递给后面的立即执行方法。
*/
func Select_Inline() {
	var user User
	var users [] User
	//1、主键
	db.First(&user, 23)
	fmt.Println("Inline==========", user)

	//2、非整型主键//普通字段
	db.First(&user, "id=?", "23")
	fmt.Println("Inline=====  =====", user)

	//3、Struct & Map
	db.Find(&users, User{Age: 18})
	fmt.Println("Inline===== Map =====", user)
}

/**
7、Extra Querying option 其它查询选项
*/
func Select_Extra() {
	// 为查询 SQL 添加额外的 SQL 操作
	var user User
	db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)
	//// SELECT * FROM users WHERE id = 10 FOR UPDATE;
}

/**
8、FirstOrInit 获取匹配的第一条记录，否则根据给定的条件初始化一个新的对象（仅支持struct和map条件）
*/
func Select_FirstOrInit() {
	var user User
	db.FirstOrInit(&user, User{Name: "liuxiaoxiang_a"})
	fmt.Println("FirstOrInit===== 未找到 ====", user)
	//FirstOrInit===== 未找到 ==== {0 liuxiaoxiang_a 0      0001-01-01 00:00:00 +0000 UTC}

	db.Where(User{Name: "liuxiaoxiang"}).FirstOrInit(&user)
	fmt.Println("FirstOrInit===== 找到 ====", user)
	//FirstOrInit===== 找到 ==== {2 liuxiaoxiang 20 liuxiaoxiang 12346    0001-01-01 00:00:00 +0000 UTC}
}

/**
9、Attrs 如果记录未找到，将使用参数初始化 struct. 三种结构都支持
*/
func Select_Attrs() {
	var user User
	db.Where(User{Name: "liuxiaoxiang3"}).Attrs(User{Age: 800}).FirstOrInit(&user)
	fmt.Println("FirstOrInit===== 未找到 ====", user)
	//FirstOrInit===== 未找到 ==== {0 liuxiaoxiang3 800      0001-01-01 00:00:00 +0000 UTC}

	user.Id = 0
	db.Where(User{Name: "liuxiaoxiang"}).Attrs(User{Age: 800}).FirstOrInit(&user)
	fmt.Println("FirstOrInit===== 找到 ====", user)
}

/**
10、Assign 不管记录是否找到，都将参数赋值给 struct.
*/
func Select_Assign() {
	var user User
	db.Where(User{Name: "liuxiaoxiang3"}).Assign(User{Age: 800}).FirstOrInit(&user)
	fmt.Println("FirstOrInit===== 未找到 ====", user)
	//FirstOrInit===== 未找到 ==== {0 liuxiaoxiang3 800      0001-01-01 00:00:00 +0000 UTC}

	user.Id = 0
	db.Where(User{Name: "liuxiaoxiang"}).Assign(User{Age: 800}).FirstOrInit(&user)
	fmt.Println("FirstOrInit===== 找到 ====", user)
}

/**
11、FirstOrCreate 不存在就创建
*/
func Select_FirstOrCreate() {
	var user User
	//查询 不存在就更新。
	db.FirstOrCreate(&user, User{Name: "xiaokeai"})
	fmt.Println("FirstOrCreate==== 未找到 ==========", user)

	db.FirstOrCreate(&user, User{Name: "liuxiaoxiang"})
	fmt.Println("FirstOrCreate==== 找到 ==========", user)
}
