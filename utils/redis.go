package utils

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

/*Redis redis*/
type Redis struct {
	pool     *redis.Pool
}

var red *Redis

/*InitRedis 初始化redis*/
func InitRedis() {
	red = new(Redis)
    red.pool = &redis.Pool{
        MaxIdle:     256,
        MaxActive:   0,
        IdleTimeout: time.Duration(120),
        Dial: func() (redis.Conn, error) {
            return redis.Dial(
                "tcp",
                "127.0.0.1:6379",
                redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
                redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
                redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
                redis.DialDatabase(0),
            )
        },
	}
}

/*Set 设置redis*/
func (c *Redis) Set(key string, val string, ex int) {
	bm := red.pool.Get()
	defer bm.Close()
	bm.Do("SET", key, val, "EX", ex)
}

/*GetRedis 获取*/
func (c *Redis) GetRedis(key string) string {
	bm := red.pool.Get()
	defer bm.Close()
	fmt.Println("---------GetRedis----------")
	val, err := redis.String(bm.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} 
	return val
}