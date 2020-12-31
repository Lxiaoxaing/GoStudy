package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

//配置类
type MyApiConfig struct {
	Listen       string                 `toml:"listen"`
	DBServers    map[string]DBServer    `toml:"dbservers"`
	RedisServers map[string]RedisServer `toml:"redisservers"`
	UserAPI      string                 `toml:"user_api"`
}

//解析toml配置
func UnameshalConfig(tomlfile string) (*MyApiConfig, error) {
	c := &MyApiConfig{}
	if _, err := toml.DecodeFile(tomlfile, c); err != nil {
		//如果解析错误则返回错误
		return c, err
	}
	return c, nil
}

//获取pgsql数据库的配置
func (c MyApiConfig) DBServerConf(key string) (DBServer, bool) {
	s, ok := c.DBServers[key]
	return s, ok
}

//获取redis数据库的配置
func (c MyApiConfig) RedisServerConf(key string) (RedisServer, bool) {
	s, ok := c.RedisServers[key]
	return s, ok
}

//监听地址
func (c MyApiConfig) GetListenAddr() string {
	return c.Listen
}

//Validate 验证配置
func (c *MyApiConfig) Validate() error {
	if c.Listen == "" {
		return fmt.Errorf("listen未配置")
	}
	if c.UserAPI == "" {
		return fmt.Errorf("user_api未配置")
	}
	return nil
}
