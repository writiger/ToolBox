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
	res, err := dao.ClockQueryByOwner(domain.Clock{Owner: owner})
	return res, err
}
