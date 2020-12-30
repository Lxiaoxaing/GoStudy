package service

import (
	"../dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginService(uname, pwd string) (c *gin.Context) {
	u := dao.LoginDao(uname, pwd)
	if u != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": u,
			"count":  1,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": u,
			"count":  1,
		})
	}
	return c
}
