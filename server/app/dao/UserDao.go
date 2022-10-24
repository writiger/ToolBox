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
	has, _ := UserIsExist(user)
	if has {
		return errorcode.GetErr(errorcode.ErrUserAlreadyExist)
	}
	_, err := engine.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

func UserEmailRedisSave(inputUser domain.User, code int) error {
	err := rc.Set(inputUser.Account, code, time.Duration(c.RegisterCodeFailureTime)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func UserEmailCodeGet(inputUser domain.User) string {
	res, _ := rc.Get(inputUser.Account).Result()
	return res
}

func UserIsExist(inputUser domain.User) (bool, error) {
	has, err := engine.Exist(&inputUser)
	if err != nil {
		return has, err
	}
	return has, nil
}

func UserLogin(inputUser *domain.User) (domain.User, error) {
	has, err := engine.Get(inputUser)
	if err != nil {
		return domain.User{}, err
	}

	if has {
		return *inputUser, nil
	} else {
		return domain.User{}, errorcode.GetErr(errorcode.ErrUserWrongPassword)
	}
}
