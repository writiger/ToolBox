package dao

import (
	"server/app/domain"
	"time"
)

// ConsumptionInsert 新增消费记录
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

// ConsumptionGetByTime 查询时间段的消费记录
func ConsumptionGetByTime(owner int64, start, end time.Time) ([]domain.Consumption, error) {
	var res []domain.Consumption
	err := engine.Where("owner = ? and created between ? and ?", owner,
		start.Format("2006-01-02 15:04:05"),
		end.Format("2006-01-02 15:04:05")).Find(&res)
	return res, err
}

// ConsumptionFoodAll 查询用户至今所有在食物上的消费额
func ConsumptionFoodAll(owner int64) (float64, error) {
	var consumption domain.Consumption
	total, err := engine.Where("owner = ? and is_food = 0", owner).Sum(consumption, "amount")
	if err != nil {
		return 0, err
	}
	return total, nil
}

// ConsumptionAmountAll 查询用户至今的所有开销
func ConsumptionAmountAll(owner int64) (float64, error) {
	var consumption domain.Consumption
	total, err := engine.Where("owner = ?", owner).Sum(consumption, "amount")
	if err != nil {
		return 0, err
	}
	return total, nil
}
