package plain

import (
	"github.com/gin-gonic/gin"
	"gorm_demo/models"
	"gorm_demo/pkg/messge"
	"gorm_demo/pkg/util"
	"time"
)

/**
创建
*/
func Create(c *gin.Context) {
	user := models.User{Name: "张创建", Age: 18, Username: "zhangchuangjian", Updated: time.Now(),Id:18}
	res := models.AddUser(user)
	if res {
		util.ApiReturn(messge.SUCCESS, res, c)
	} else {
		util.ApiReturn(messge.FAIL, res, c)
	}

}
