package main

import (
	"fmt"
	"server/app/dao"
	"server/utils"
	"strconv"
	"testing"
	"time"
)

func TestConsumptionGetMonth(t *testing.T) {
	t.Run("查询一个月的消费记录", func(t *testing.T) {
		_, err := dao.ConsumptionGetByTime(1, time.Unix(1679197206, 0), time.Unix(1679197206, 0))
		if err != nil {
			return
		}
	})
}

func TestConsumptionGetFoodAmount(t *testing.T) {
	t.Run("查询用户的饮食开销", func(t *testing.T) {
		total, err := dao.ConsumptionFoodAll(1)
		fmt.Println(err)
		fmt.Println(total)
	})
}

func TestConsumptionGetAmount(t *testing.T) {
	t.Run("查询用户的饮食开销", func(t *testing.T) {
		total, err := dao.ConsumptionAmountAll(1)
		fmt.Println(err)
		fmt.Println(total)
	})
}

func Test(t *testing.T) {
	t.Run("redisInsert", func(t *testing.T) {
		err := utils.RedisDb.Set("1:engel", 43.66, 24*7*time.Hour).Err()
		fmt.Println(err)
	})
	t.Run("redis", func(t *testing.T) {
		res, err := utils.RedisDb.Get(strconv.FormatInt(1, 10) + ":engel").Result()
		fmt.Println(err)
		fmt.Println(res)
	})
}
