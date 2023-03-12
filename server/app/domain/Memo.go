package domain

import "time"

const (
	MemoUnfinished = 0
	MemoFinished   = 1
	MemoTimeout    = 2
)

type Memo struct {
	Id      int64     `form:"id"`
	Info    string    `form:"info"`
	Owner   int64     `form:"owner"`
	Status  int       `form:"status"` // 1-未完成 2-已完成
	EndTime int64     `form:"end_time"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
