package dao

import (
	"server/app/domain"

	"xorm.io/xorm"
)

// 声明engine后使用xorm
var engine *xorm.Engine

func init() {
	engine = domain.InitDB()
}

func UserAdd(user domain.User) error {
	engine.Insert(&user)
	return nil
}
