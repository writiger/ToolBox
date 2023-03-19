package domain

import "time"

type User struct {
	Id       int64
	Name     string
	Password string
	Account  string
	Avatar   int
	// 账户余额
	Balance int
	// 每月的入账
	BaseDisk int
	// 每月入账的时间
	// 使用cron进行定时任务
	PayDay  string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
