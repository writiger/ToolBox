package domain

import "time"

type Cipher struct {
	Id      int64
	Use     string
	Owner   string
	Info    string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
