package service

import (
	"server/app/dao"
	"server/app/domain"
)

func ClockUpload(inputClock domain.Clock) error {
	err := dao.ClockInsert(inputClock)
	return err
}

func ClockQueryByOwner(owner string) ([]domain.Clock, error) {
	res, err := dao.ClockFindByOwner(domain.Clock{Owner: owner})
	return res, err
}

func ClockClockIn(id int64, owner string) error {
	inputClock := domain.Clock{
		Id:    id,
		Owner: owner,
	}
	// 判断是否达成目标
	isGonnaFinished, err := dao.ClockIsGonnaFinished(id)
	if err != nil {
		return err
	}
	if isGonnaFinished {
		// 将状态修改为已完成
		inputClock.Status = 2
	}
	achieveNow, err := dao.ClockGetAchieveNow(inputClock)
	if err != nil {
		return err
	}
	inputClock.AchieveNow = achieveNow + 1
	err = dao.ClockUpdate(inputClock)
	return err
}

func ClockClockDelete(id int64, owner string) error {
	inputClock := domain.Clock{
		Id:    id,
		Owner: owner,
	}
	err := dao.ClockDelete(inputClock)
	return err
}
