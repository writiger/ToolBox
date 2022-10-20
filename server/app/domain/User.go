package domain

import "time"

type User struct {
	Id       int64
	Name     string
	Password string
	Account  string
	Avatar   int
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}
