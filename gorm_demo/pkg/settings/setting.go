package setting

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

//type App struct {
//	RuntimeRootPath string
//	LogPath         string
//	PrePath         string
//}

//var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}
var DatabaseSetting = &Database{}
var ServerSetting = &Server{}
var cfg *ini.File

// Setup initialize the configuration instance
func Setup(content string) {
	var err error
	//destPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	destPath, _ := os.Getwd()
	if content == "" {
		cfg, err = ini.Load(destPath + "/config/app.ini")
	} else {
		cfg, err = ini.Load([]byte(content))
	}
	if err != nil {
		fmt.Println("err", err)
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	//mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	fmt.Println("ServerSetting", ServerSetting)
	fmt.Println("DatabaseSetting", DatabaseSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
