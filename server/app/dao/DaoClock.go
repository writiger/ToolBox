package dao

import (
	"server/app/domain"
)

func ClockInsert(inputClock domain.Clock) error {
	_, err := engine.Insert(inputClock)
	return err
}

func ClockQueryByOwner(inputClock domain.Clock) ([]domain.Clock, error) {
	clocks := make([]domain.Clock, 0)
	err := engine.Where("owner = ?", inputClock.Owner).Find(&clocks)
	return clocks, err
}
