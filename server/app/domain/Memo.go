package domain

import "time"

type Memo struct {
	Id      int64
	Info    string
	Belongs int64
	Status  int
	EndTime int64
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
