package dao

import (
	"server/app/domain"
	"server/error_code"
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

func UserInsert(user domain.User) error {
	has, _ := UserIsExist(user.Account)
	if has {
		return errorcode.GetErr(30105)
	}
	_, err := engine.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

func UserEmailRedisSave(account string, code int) error {
	err := rc.Set(account, code, time.Duration(c.RegisterCodeFailureTime)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func UserEmailCodeGet(account string) string {
	res, _ := rc.Get(account).Result()
	return res
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
