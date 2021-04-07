package plain

import (
	"github.com/gin-gonic/gin"
	"gorm_demo/models"
	"gorm_demo/pkg/messge"
	"gorm_demo/pkg/util"
	"time"
)

var user = models.User{Name: "张创建", Age: 18, Username: "zhangchuangjian", Updated: time.Now(),
	Password: "123", Phone: "15613282905", Email: "1174755886@qq.com", Created: time.Now().Format("2006-01-02 15:04:05")}

/**
创建
*/
func Create(c *gin.Context) {

	res := models.AddUser(user)
	if res {
		util.ApiReturn(messge.SUCCESS, res, c)
	} else {
		util.ApiReturn(messge.FAIL, res, c)
	}
}

/**
创建
*/
func DbOperation(c *gin.Context) {
	//1、更新操作
	//res := models.UpdateUser(user)

	//2、查询操作
	res := models.SelectUser(user)
	util.ApiReturn(messge.SUCCESS, res, c)

}
