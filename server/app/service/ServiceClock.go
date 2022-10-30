package service

import (
	"server/app/dao"
	"server/app/domain"
)

func ClockUpload(inputClock domain.Clock) error {
	err := dao.ClockInsert(inputClock)
	return err
}
