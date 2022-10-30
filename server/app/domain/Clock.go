package domain

import "time"

const (
	KindTiming = 1
	KindTarget = 2
	KindTimes  = 3
)

const (
	StatusIng = 1
	StatusEd  = 2
)

type Clock struct {
	Id    int64
	Owner string
	// Kind 打卡类型
	// 1 -> 定时打卡 每次在固定时间之后提示打卡
	// 2 -> 次数打卡 只记录次数并有实现目标
	// 3 -> 记录	 仅仅记录完成次数
	Kind          int
	Start         int
	AchieveTarget int
	AchieveNow    int
	Interval      string
	// Status 打卡状态
	// 1 -> 进行中
	// 2 -> 已完成
	Status  int
	Name    string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
