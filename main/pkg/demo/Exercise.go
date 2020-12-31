package demo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//集合
func Collection() {
	//1、数组 array
	var a [10]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	//2、slice

	//3、map
	ages := map[string]int{
		"lili":  13,
		"nick":  23,
		"jacky": 55,
	}
	fmt.Println(ages)
	fmt.Println(len(ages))
	fmt.Println(ages["lili"])
}

//循环  go中没有do while,循环结构只有for,选择结构有if和switch
func Loop() {
	//循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("i=", i)
		fmt.Println("sum=", sum)
	}
	list := []string{"aaa", "bbb", "ccc"}
	for _, v := range list {
		fmt.Println("%s", v)
	}

	//switch
	a := "test8"
	switch a {
	case "test1":
		fmt.Println("test1")
	case "test2", "test3":
		fmt.Println("testOther")
	default:
		fmt.Println("NoTest")
	}

}

/**
文件操作
*/
func FileOperation() {
	//将字符串写到文件
	WriteStringToFile()
	//将字节写到文件
	WriteByteToFile()
	//将字符串一行行写入文件
	WriteStringToFileByLine()
	//追加到文件
	AddStringToFile()
	//并发写入文件
	ConcurrentWriteToFile()
}

/**
常量赋值
*/
func Constant() {
	fmt.Println("以下是常量赋值**********************")
	//1、itoa使用
	//const中 每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。使用iota能简化定义，在定义枚举时很有用。
	const (
		a1 = iota
		a2
		a3
		a4
	)
	fmt.Println("a1=", a1)
	fmt.Println("a2=", a2)
	fmt.Println("a3=", a3)
	fmt.Println("a4=", a4)

	//2、string
	c := "this is string"
	fmt.Println("c=", c)

	//3、
	d := []byte(c)
	d[0] = 'd'
	e := string(d)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
}

/**
变量赋值函数
*/
func Variable() {
	fmt.Println("以下是变量赋值**********************")
	//第一种方式
	var a = 12
	a = 13
	fmt.Println(a)

	//第二种方式（b的赋值）
	a, b := 12, 23
	fmt.Print(a, b)

	//第三种赋值
	c := 18
	fmt.Println(c)
}

/**
接口
*/
func Imath() {
	// 将字符串输出到控制台
	fmt.Println("hello world")
	fmt.Println(Add(1, 1))
	fmt.Println(Sub(1, 1))

	srv := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(defaultHttp),
	}
	srv.ListenAndServe()
}

/**
接口输出
*/
func defaultHttp(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("wow"))
}

// 自定义返回
type JsonRes struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	TimeStamp int64       `json:"timestmap"`
}

func apiResult(w http.ResponseWriter, code int, data interface{}, msg string) {
	body, _ := json.Marshal(JsonRes{
		Code: code,
		Data: data,
		Msg:  msg,
		// 获取时间戳
		TimeStamp: time.Now().Unix(),
	})
	w.Write(body)
}

//练习题
func Exercise1() {
	//1、创建一个基于for的简单的循环。使用循环10次，并且使用fmt包打印出计数器的值。
	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	}

	//2、用goto修改1的循环，不可使用for
	i := 0
I:
	fmt.Println(i)
	i++
	if i < 5 {
		goto I
	}

	//3、再次改写1的循环，使其遍历一个array，并将array打印到屏幕上
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//方式一
	for _, val := range arr {
		fmt.Println("val=", val)
	}
	//方式二
	for i := 0; i < len(arr); i++ {
		fmt.Println("i=", i)
	}

	//4、编写一个逆转字符的程序，例如："foobar"打印成为raboof
	s := "foobar"
	var result string
	for _, value := range s {
		result = string(value) + result
	}
	fmt.Println(result)

}
