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
	create := db.Save(user)
	fmt.Println(create)
	//if err := create; err != nil {
	//	fmt.Println("插入失败")
	//	return false
	//}
	return true
}
