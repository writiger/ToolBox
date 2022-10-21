package dao

import (
	"server/app/domain"
	"server/utils"
	"time"

	"github.com/go-redis/redis"
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
	rc = domain.InitRedisClient()
	c = utils.GetConf()
}

// TODO
func UserInsert(user domain.User) error {
	engine.Insert(&user)
	return nil
}

func UserEmailRedisSave(account string, code int) error {
	err := rc.Set(account, code, time.Duration(c.RegisterCodeFailureTime)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func UserIsExist(account string) (bool, error) {
	has, err := engine.Exist(&domain.User{
		Account: account,
	})
	if err != nil {
		return has, err
	}
	return has, nil
}
