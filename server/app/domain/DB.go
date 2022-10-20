package domain

import (
	"fmt"
	u "server/utils"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

// 初始化数据库
func InitDB() *xorm.Engine {
	c := u.GetConf()
	var err error
	engine, err = xorm.NewEngine("mysql", c.DataSourceName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err2 := engine.Ping()
		if err2 != nil {
			fmt.Println(err2.Error())
		} else {
			print("数据库连接成功")
		}
	}
	return engine
}

func InitTable() {
	// 创建用户表
	err := engine.Sync(new(User))
	if err != nil {
		fmt.Println(err.Error())
	}
}
