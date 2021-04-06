package main

import (
	"github.com/spf13/viper"
	"os"
)

func main() {
	//demo1
	//r:=routers.SetupRouter()
	//if err :=r.Run();err!=nil {
	//	fmt.Println("startup service failed,err:%v\n",err)
	//}
	InitConfig()

}

func InitConfig() {
	conf, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(conf + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
