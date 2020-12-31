package handlers

import (
	"../middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/**
处理用户想滚逻辑的控制器
*/

//用户注册控制器
type UserRegisterReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (c *UserRegisterReq) Validate() error {
	if len(c.Username) < 5 {
		return fmt.Errorf("用户名太短")
	}
	return nil
}

func UserRegister(c *gin.Context) {
	req := c.MustGet(middleware.MiddlewareParam).(*UserRegisterReq)
	_, existsUser, err := getUser(c, req.Username)
	if err != nil {
		log.Fatalf("UserRegister:查询用户信息出错，req:%v,err:%v", req, err)
		middleware.ResponseError(c, middleware.ServerErrorCode, err, "用户查询信息出错")
		return
	}

	//如果用户已经存在
	if existsUser {
		middleware.ResponseError(c, middleware.ServerErrorCode, err, ("y用户已经存在"))
		return
	}

	//如果用户注册失败
	if err = addUser(c, req.Username, req.Password); err != nil {
		log.Fatalf("UserRegister:注册用户信息失败，req:%v,err:%v", req, err)
		middleware.ResponseError(c, middleware.ServerErrorCode, err, "添加用户失败")
	}
	middleware.ResponseSuccess(c, "请求成功", nil)
}

//用户登录注册器
type UserLoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func UserLogin(c *gin.Context) {
	req := c.MustGet(middleware.MiddlewareParam).(*UserLoginReq)

	//查询用户信息
	user, existsUser, err := getUser(c, req.Username)
	if err != nil {
		log.Fatalf("UserLogin:查询用户信息出错，req:%v,err:%v", req, err)
		middleware.ResponseError(c, middleware.ServerErrorCode, err, "查询用户信息出错")
		return
	}

	if !existsUser {
		middleware.ResponseError(c, middleware.ServerErrorCode, err, "用户不存在，请先注册")
		return
	}

	if user.Password != req.Password {
		middleware.ResponseError(c, middleware.ServerErrorCode, err, "密码验证错误")
		return
	}
	middleware.ResponseSuccess(c, "请求成功", user)
}

func UserGetAll(c *gin.Context) {
	users, err := getUsers(c)
	if err != nil {
		middleware.ResponseError(c, middleware.ServerErrorCode, err, "服务器错误")
		return
	}
	middleware.ResponseSuccess(c, "请求成功", users)
	return
}

//用户删除表单
type UserDelreq struct {
	Username string `json:"username" form:"username" binding:"required"`
}

func DelUser(c *gin.Context) {
	req := c.MustGet(middleware.MiddlewareParam).(*UserDelreq)
	err := delUser(c, req.Username)
	if err != nil {
		middleware.ResponseError(c, middleware.ServerErrorCode, err, "删除错误，服务器错误")
		return
	}
	middleware.ResponseSuccess(c, "删除成功", nil)
}

//用户更新表单
type UserUptreq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (c *UserUptreq) Validate() {
	if c.Username == "" || len(c.Username) == 0 {
		fmt.Println("用户名不能为空")
	}
}

func UptUser(c *gin.Context) {
	req := c.MustGet(middleware.MiddlewareParam).(*UserUptreq)
	UsrReq := UserUptreq{
		Username: req.Username,
		Password: req.Password,
	}
	err := uptUser(c, UsrReq)
	if err != nil {
		middleware.ResponseError(c, middleware.ServerErrorCode, err, "更新失败，服务器错误")
		return
	}
	middleware.ResponseSuccess(c, "更新成功", nil)
}
