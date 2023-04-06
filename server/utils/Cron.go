package utils

import (
	"github.com/robfig/cron/v3"
	"log"
)

func InitSchedule() *cron.Cron {
	schedule := cron.New()
	// 每月调用一次
	if _, err := schedule.AddFunc("* * * * 1 *", addPay); err != nil {
		log.Println(err.Error())
	}
	return schedule
}

func addPay() {
	log.Println("开始导入工资")
	defer log.Println("导入工资结束")

}
