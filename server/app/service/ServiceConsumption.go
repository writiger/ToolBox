package service

import (
	"server/app/dao"
	"server/app/domain"
	"time"
)

func ConsumptionUpload(inputConsumption *domain.Consumption) error {
	return dao.ConsumptionInsert(inputConsumption)
}

func ConsumptionGetMonth(owner int64, start time.Time) ([]domain.Consumption, error) {
	return dao.ConsumptionGetByMonth(owner, start)
}
