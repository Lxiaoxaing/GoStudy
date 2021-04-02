package hhnacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
	"path/filepath"
	"strings"
)

// nacos公共配置

const nacosDomain = "nacos.hitecloud.cn"
const nacosPath = "/nacos"
const nacosPort = 8848
const nacosDataid = "hitecloud.service-platform.rtc_service"
const nacosGroup = "service-platform"

var nacosUsername string
var nacosPassword string

// GetConfigByEnv 从环境变量获取nacos配置信息
// 格式 namespaceid:用户名:密码
func GetConfigByEnv() string {
	hiteNacos := os.Getenv("HITENACOS")
	if hiteNacos == ""{
		panic("nacos 环境变量异常")
	}
	fmt.Println("nacosDataid-----------",nacosDataid)
	fmt.Println("hiteNacos-----------",hiteNacos)
	hiteNacosArr := strings.Split(hiteNacos,":")
	if len(hiteNacosArr) <3 {
		panic("nacos 环境变量异常")
	}
	namespaceId := hiteNacosArr[0]
	nacosUsername = hiteNacosArr[1]
	nacosPassword = hiteNacosArr[2]
	return GetConfig(nacosDataid, nacosGroup, namespaceId)
}


// GetConfig 从nacos中获取配置
func GetConfig(dataId string, group string, namespaceId string) string {

	destPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.MkdirAll(destPath+"/runtime", os.ModePerm)
	clientConfig := constant.ClientConfig{
		NamespaceId:         namespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              destPath + "/runtime/nacoslog",
		CacheDir:            destPath + "/runtime/nacoscache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
		Username:            nacosUsername,
		Password:            nacosPassword,
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacosDomain,
			ContextPath: nacosPath,
			Port:        nacosPort,
		},
	}
	configClient, errclient := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig": clientConfig,
	})
	if errclient != nil {
		fmt.Println("nacos 客户端初始化错误：", errclient)
		os.Exit(1)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		fmt.Println("nacos 获取配置错误：", err)
		os.Exit(1)
	}
	return content
}
