package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

//用户
type User struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Username string `db:"username"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
	Email    string `db:"email"`
	Created  string `db:"created"`
	Updated  string `db:"updated"`
}

//获取所有用户
func GetUsers(testDB *gorm.DB) (users []User, err error) {
	var u []User
	if err = testDB.Table("user").Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

//查找单个用户
func GetUser(testDB *gorm.DB, userName string) (user *User, exists bool, err error) {
	var tempUser User
	if err = testDB.Table("user").Where("username=?", userName).Find(&tempUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			//如果并不是找不到用户的操作
			return nil, false, err
		}
		//如果是找不到的问题，直接返回
		return nil, false, nil
	}
	return &tempUser, true, err
}

//增加用户
func AddUser(testdb *sqlx.DB, user *User) (err error) {
	psql := `insert into user(username,password) values($1,$2)`
	_, err = testdb.Exec(psql, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

//删除用户
func DelUser(testDB *gorm.DB, username string) (err error) {
	if err = testDB.Table("user").Where("username=?", username).Delete(&username).Error; err != nil {
		return err
	}
	return nil
}

//更改用户密码
func UpdUser(testDB *sqlx.DB, user User) (err error) {
	psql := `update user set password = $1 where username = $2`
	testDB.Prepare(psql)
	_, err = testDB.Exec(psql, user.Password, user.Username)
	if err != nil {
		return err
	}
	return nil
}
