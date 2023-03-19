package main

import (
	"server/app/dao"
	"testing"
	"time"
)

func TestConsumptionGetMonth(t *testing.T) {
	t.Run("查询一个月的消费记录", func(t *testing.T) {
		_, err := dao.ConsumptionGetByMonth(1, time.Unix(1679197206, 0))
		if err != nil {
			return
		}
	})
}
