package domain

import (
	"fmt"
	u "server/utils"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// 使用xorm映射数据库
var engine *xorm.Engine

// 声明一个全局的redisDb变量
var redisDb *redis.Client

func InitRedisClient() *redis.Client {
	c := u.GetConf()
	// 获取redisClient
	redisDb = redis.NewClient(&redis.Options{
		Addr:     c.RedisSourceName,
		Password: c.RedisPassword,
		DB:       c.RedisDB,
	})
	_, err := redisDb.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Redis连接成功")
	return redisDb
}

// InitSqlDB 初始化数据库
func InitSqlDB() *xorm.Engine {
	c := u.GetConf()
	var err error
	engine, err = xorm.NewEngine("mysql", c.DataSourceName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// 测试redis连接
		err2 := engine.Ping()
		if err2 != nil {
			fmt.Println(err2.Error())
		} else {
			print("数据库连接成功")
		}
	}
	// 如果表不存在则创建
	InitTable()
	return engine
}

func InitTable() {
	// 创建表
	err := engine.Sync(new(User), new(Memo), new(Cipher), new(Clock))
	if err != nil {
		panic(err.Error())
	}
}
