package domain

import "time"

const (
	MemoUnfinished = 0
	MemoFinished   = 1
	MemoTimeout    = 2
)

type Memo struct {
	Id      int64
	Info    string
	Owner   int64
	Status  int
	EndTime int64
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
