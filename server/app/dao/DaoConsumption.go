package dao

import (
	"server/app/domain"
	"time"
)

func ConsumptionInsert(inputConsumption *domain.Consumption) error {
	// 1. 新建事务
	session := engine.NewSession()
	defer session.Close()
	// 2. 记录消费
	if _, err := session.Insert(inputConsumption); err != nil {
		return err
	}
	// 3. 从用户扣款
	sql := "update `user` set balance = balance - ? where id = ?"
	if _, err := session.Exec(sql, inputConsumption.Amount, inputConsumption.Owner); err != nil {
		return err
	}
	// 4. 提交事务
	return session.Commit()
}

func ConsumptionGetByMonth(owner int64, start time.Time) ([]domain.Consumption, error) {
	var res []domain.Consumption
	month, _ := time.ParseDuration("720h")
	end := start.Add(month)
	err := engine.Where("owner = ? and created between ? and ?", owner,
		start.Format("2006-01-02 15:04:05"),
		end.Format("2006-01-02 15:04:05")).Find(&res)
	return res, err
}
