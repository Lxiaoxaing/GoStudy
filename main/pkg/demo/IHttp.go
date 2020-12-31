package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func IHttp() {
	//GET请求网页
	webGET()
	//带参数的请求
	paramGET()

}

//带参数的GET请求
func paramGET() {

}


//GET 请求网页，将页面HTML、CSS等规则在网页渲染展示出来
func webGET() {
	resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Println("get failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed,err:%v\n", err)
		return
	}
	fmt.Printf(string(body))
}

