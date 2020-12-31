package models

import (
	"../utils/config"
	"fmt"
	"../models"
)

func LoginDao(uname, pwd string) *models.User {
	sqlStr := "select id,username,`password`,phone,email,created,updated from user where username=? and password=?"
	rows, err := config.Db.Query(sqlStr, uname, pwd)
	if err != nil {
		fmt.Println("Login failed,err:%v\n", err)
		return nil
	}
	//非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	for rows.Next() {
		user := new(models.User)
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Phone, &user.Email, &user.Created, &user.Updated)
		return user
	}
	return nil
}
