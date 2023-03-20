package domain

import (
	"fmt"
	u "server/utils"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// 使用xorm映射数据库
var engine *xorm.Engine

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
	err := engine.Sync(new(User), new(Memo), new(Cipher), new(Clock), new(Consumption))
	if err != nil {
		panic(err.Error())
	}
}
