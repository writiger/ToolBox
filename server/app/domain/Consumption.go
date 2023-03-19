package domain

import "time"

type Consumption struct {
	Id      int64
	Info    string
	Owner   int64
	Amount  float64
	IsFood  bool
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
