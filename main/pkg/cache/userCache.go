package cache

import "github.com/garyburd/redigo/redis"

/**
用户缓存配置
*/

const (
	//hash storage
	cacheNameUser = "user"
)

//获取缓存中用户
func GetUser(testcache *redis.Pool, username string) (userBytes string, exists bool, err error) {
	conn := testcache.Get()
	defer conn.Close() //延迟关闭数据库连接
	userBytes, err = redis.String(conn.Do("hget", cacheNameUser, username))
	if err != nil {
		if err != redis.ErrNil {
			//如果并不是redis服务器的错误
			return userBytes, false, err
		}
		//如果是服务器的错误
		return userBytes, false, nil
	}
	//返回正确的结果集
	return userBytes, true, nil
}

//在缓存中存入用户信息
func SetUser(testcache *redis.Pool, username string, userInfo string) (err error) {
	conn := testcache.Get()
	defer conn.Close()
	_, err = conn.Do("hset", cacheNameUser, username, userInfo)
	return
}
