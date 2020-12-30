package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

/**
初始化数据库连接池
*/

/**
初始化sql数据库的配置
*/
//主意结构体的字段名必须大写，不然外部的包会访问不到
type DBServer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Dbname   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

//初始化redis数据库配置
type RedisServer struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int `toml:"db"`
}

//初始化连接数据库的字符串
func (m DBServer) ConnectString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", m.Host, m.Port, m.User, m.Dbname)
}

//初始化pgsql的连接
func (m DBServer) NewPostgrestDB(idleConnection int) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("postgres", m.ConnectString())
	if err != nil {
		return
	}
	//设置最大空间连接数
	db.SetMaxIdleConns(idleConnection)
	return
}

//初始化redis连接池
func (c RedisServer) NewRedisPool(maxIdle int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", c.Addr, redis.DialDatabase(c.DB), redis.DialPassword(c.Password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}


