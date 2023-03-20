package dao

import (
	"github.com/go-redis/redis"
	"server/app/domain"
	"server/utils"
	"xorm.io/xorm"
)

// 声明engine后使用xorm
var engine *xorm.Engine

// 声明redis client
var rc *redis.Client

// 声明配置文件
var c *utils.Conf

func init() {
	engine = domain.InitSqlDB()
	c = utils.GetConf()
}
