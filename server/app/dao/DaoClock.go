package dao

import (
	"server/app/domain"
	errorcode "server/error_code"
)

func ClockInsert(inputClock domain.Clock) error {
	_, err := engine.Insert(inputClock)
	return err
}

func ClockFindByOwner(inputClock domain.Clock) ([]domain.Clock, error) {
	clocks := make([]domain.Clock, 0)
	err := engine.Where("owner = ?", inputClock.Owner).Find(&clocks)
	return clocks, err
}

func ClockUpdate(inputClock domain.Clock) error {
	_, err := engine.Where("id = ? and owner = ?", inputClock.Id, inputClock.Owner).Update(inputClock)
	return err
}

func ClockGetAchieveNow(inputClock domain.Clock) (int, error) {
	var ints []int
	err := engine.Table("clock").Where("id = ? and owner = ?", inputClock.Id, inputClock.Owner).
		Cols("achieve_now").Find(&ints)
	if len(ints) == 0 {
		return 0, errorcode.GetErr(errorcode.ErrClockNil)
	}
	return ints[0], err
}

func ClockDelete(inputClock domain.Clock) error {
	has, err := engine.Table("clock").
		Where("id = ? and owner = ?", inputClock.Id, inputClock.Owner).Exist()
	if !has {
		return errorcode.GetErr(errorcode.ErrClockNil)
	}
	_, err = engine.Delete(inputClock)
	return err
}

func ClockIsGonnaFinished(id int64) (bool, error) {
	clock := &domain.Clock{Id: id}
	_, err := engine.Get(clock)
	if err != nil {
		return false, err
	}
	// target为-1则无上限
	if clock.AchieveTarget == -1 {
		return false, nil
	}
	// 即将达成
	if clock.AchieveNow+1 >= clock.AchieveTarget {
		return true, nil
	}
	return false, nil
}
