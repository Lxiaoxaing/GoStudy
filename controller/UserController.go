package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"../service"
	"../entity"
)

type UserController struct {

}

//注册的路由方法
func (mc * UserController)Router(engine *gin.Engine)  {
	engine.GET("/login",mc.Login)
}

//func UserHandle() {
//	http.HandleFunc("/login", LoginController)
//	http.HandleFunc("/", SayHello)
//}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello girl!")
}

//登录控制器
//func LoginController(w http.ResponseWriter, r *http.Request) {
func (mc *UserController)Login(w http.ResponseWriter, r *http.Request) {
	//参数解析
	uname := r.FormValue("uname")
	pwd := r.FormValue("pwd")

	c := service.LoginService(uname, pwd) //登录逻辑
	b, _ := json.Marshal(c)               //结构转对象
	w.Header().Set("Content-Type", "application/json;charset=uft-8")
	w.Write(b)
}
