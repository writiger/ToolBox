package utils

import (
	"fmt"
	"github.com/go-redis/redis"
)

// RedisDb 声明一个全局的redisDb变量
var RedisDb *redis.Client

func init() {
	RedisDb = InitRedisClient()
}

func InitRedisClient() *redis.Client {
	c := GetConf()
	// 获取redisClient
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     c.RedisSourceName,
		Password: c.RedisPassword,
		DB:       c.RedisDB,
	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Redis连接成功")
	return RedisDb
}
