package utils

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/cache"
)

// 定义常量
var (
	RedisClient cache.Cache
)

/*InitRedis 初始化redis*/
func InitRedis() cache.Cache {
	key := beego.AppConfig.String("redis.key")
	conn := beego.AppConfig.String("redis.conn")
	dbNum := beego.AppConfig.String("redis.dbNum")
	// 设置配置参数
	option := map[string]string{
		"key":   key,
		"conn":  conn,
		"dbNum": dbNum,
	}
	optionStr, _ := json.Marshal(option)
	fmt.Println(string(optionStr))
	bm, _ := cache.NewCache("redis", string(optionStr))
	return bm
}
