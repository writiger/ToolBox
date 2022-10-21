package dao

import (
	"server/app/domain"

	"xorm.io/xorm"
)

// 声明engine后使用xorm
var engine *xorm.Engine

func init() {
	engine = domain.InitSqlDB()
}

func UserInsert(user domain.User) error {
	engine.Insert(&user)
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
