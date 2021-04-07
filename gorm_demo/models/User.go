package models

import (
	"fmt"
	"time"
)

type User struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Username string    `json:"username"` // 用户名
	Password string    `json:"password"` // 密码
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Created  string    `json:"created"`
	Updated  time.Time `json:"updated"`

}

func (User) TableName() string {
	return "user"
}

func AddUser(user User) bool {
	//1、create:新增。必须把参数补全，包括id
	create := db.Create(&user)
	fmt.Println("create==================", create)

	//2、newRecord：检测。仅检查当前值是否具有主键 . 所以不要插入记录.创建之前返回true,创建之后返回false
	newRecord := db.NewRecord(&user)
	fmt.Println("newRecord==================", newRecord)

	//3、save:保存，id不存在时为新增
	//        更新，id存在时为更新，默认是更新所有的字段.

	save := db.Save(&user)
	fmt.Println("save==================", save)
	return true
}

func UpdateUser(user User) bool {
	//1、save更新所有字段
	Update_saveUser(user)

	//2、更新修改字段，where条件都包含id
	Update_updateUser(user)

	//3、更新选定字段
	Update_updateSelect(user)

	//4、无Hooks更新
	Update_NoHooks(user)

	//5、批量更新
	Update_bulk(user)
	return true
}


func SelectUser(user User) User {
	//1、简单查询
	Select_Simple()

	//2、带条件查询
	Select_where()

	//3、Struct & Map
	Select_Struct2Map()

	//4、Not条件
	Select_Not()

	//5、Or条件
	Select_Or()

	//6、内连接
	Select_Inline()

	//7、Extra Querying option 其它查询选项
	Select_Extra()

	//8、FirstOrInit 获取匹配的第一条记录，否则根据给定的条件初始化一个新的对象（仅支持struct和map条件）
	Select_FirstOrInit()

	//9、Attrs 如果记录未找到，将使用参数初始化 struct.
	Select_Attrs()

	//10、Assign 不管记录是否找到，都将参数赋值给 struct.
	Select_Assign()

	//11、FirstOrCreate 不存在就创建
	Select_FirstOrCreate()
	return user
}


