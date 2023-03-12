package domain

import "time"

type Cipher struct {
	Id      int64
	Use     string    // 用途
	Owner   string    // 所有者
	Info    string    `xorm:"text info"` // 信息
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
