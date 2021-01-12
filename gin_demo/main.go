package main

import (
	"fmt"
	"gin_demo/routers"
)

func main() {
	//demo1
	r:=routers.SetupRouter()
	if err :=r.Run();err!=nil {
		fmt.Println("startup service failed,err:%v\n",err)
	}
}



