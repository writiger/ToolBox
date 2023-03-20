package service

import (
	"server/app/dao"
	"server/app/domain"
	"server/utils"
	"strconv"
	"time"
)

// ConsumptionUpload 上传消费记录
func ConsumptionUpload(inputConsumption *domain.Consumption) error {
	//从数据库中读取并计算
	owner := inputConsumption.Owner
	all, err := dao.ConsumptionAmountAll(owner)
	if err != nil {
		return err
	}
	food, err := dao.ConsumptionFoodAll(owner)
	if err != nil {
		return err
	}
	engel := food / all
	// 每次消费后都更新缓存中的恩格尔系数
	err = utils.RedisDb.Set(strconv.FormatInt(owner, 10)+":engel", engel, 24*7*time.Hour).Err()
	if err != nil {
		return err
	}

	return dao.ConsumptionInsert(inputConsumption)
}

// ConsumptionGetMonth 查询用户一个月的消费记录
func ConsumptionGetMonth(owner int64, start time.Time) ([]domain.Consumption, error) {
	month, _ := time.ParseDuration("720h")
	end := start.Add(month)
	return dao.ConsumptionGetByTime(owner, start, end)
}

// ConsumptionEngel 查询用户的恩格尔系数
func ConsumptionEngel(owner int64) (float64, error) {
	// 1. 尝试从缓存中获取数据
	val, err := utils.RedisDb.Get(strconv.FormatInt(owner, 10) + ":engel").Result()
	if val != "" {
		res, _ := strconv.ParseFloat(val, 64)
		return res, nil
	}
	// 2. 缓存中暂无数据
	//从数据库中读取并计算
	all, err := dao.ConsumptionAmountAll(owner)
	if err != nil {
		return 0, err
	}
	food, err := dao.ConsumptionFoodAll(owner)
	if err != nil {
		return 0, err
	}
	engel := food / all
	//缓存
	err = utils.RedisDb.Set(strconv.FormatInt(owner, 10)+":engel", engel, 24*7*time.Hour).Err()
	return engel, err
}
